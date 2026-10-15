from collections.abc import Iterator
from datetime import UTC, datetime
from typing import Annotated, Literal
from uuid import UUID

from fastapi import APIRouter, Depends, HTTPException, status
from fastapi import Request as HTTPRequest
from fastapi.security import HTTPAuthorizationCredentials, HTTPBearer
from pydantic import BaseModel

from edutech_be import auth
from edutech_be.config import Settings, settings
from edutech_be.sql_store import SqlAuthStore
from edutech_be.store import AuthStore

# Every HTTP route lives in this file, with its full path, so the list is easy to scan.
router = APIRouter()
_bearer = HTTPBearer(auto_error=False)


class RegisterRequest(BaseModel):
    email: str
    password: str


class RegisterResponse(BaseModel):
    email: str


class LoginRequest(BaseModel):
    email: str
    password: str


class LoginResponse(BaseModel):
    access_token: str
    refresh_token: str
    token_type: Literal["bearer"] = "bearer"
    expires_in: int


class RefreshRequest(BaseModel):
    refresh_token: str


class AccessTokenResponse(BaseModel):
    access_token: str
    token_type: Literal["bearer"] = "bearer"
    expires_in: int


class MeResponse(BaseModel):
    id: UUID
    email: str


def get_store(request: HTTPRequest) -> Iterator[AuthStore]:
    # One transaction per request, so a failed request leaves no partial rows.
    with request.app.state.engine.begin() as connection:
        yield SqlAuthStore(connection)


def get_settings() -> Settings:
    return settings


def _now() -> datetime:
    return datetime.now(UTC)


def _unauthorized(detail: str) -> HTTPException:
    return HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail=detail,
        headers={"WWW-Authenticate": "Bearer"},
    )


def _bearer_token(
    credentials: Annotated[HTTPAuthorizationCredentials | None, Depends(_bearer)],
) -> str:
    if credentials is None:
        raise _unauthorized("Not authenticated")
    return credentials.credentials


@router.post(
    "/auth/register",
    status_code=status.HTTP_201_CREATED,
    response_model=RegisterResponse,
)
def register(
    body: RegisterRequest, store: Annotated[AuthStore, Depends(get_store)]
) -> RegisterResponse:
    try:
        email = auth.register(store, body.email, body.password, _now())
    except auth.InvalidEmail:
        raise HTTPException(
            status.HTTP_422_UNPROCESSABLE_CONTENT, "Invalid email"
        ) from None
    except auth.WeakPassword:
        raise HTTPException(
            status.HTTP_422_UNPROCESSABLE_CONTENT,
            "Password must be 8 to 64 characters and not a common password",
        ) from None
    except auth.EmailTaken:
        raise HTTPException(
            status.HTTP_409_CONFLICT, "Email is already registered"
        ) from None
    return RegisterResponse(email=email)


@router.post("/auth/login", response_model=LoginResponse)
def login(
    body: LoginRequest,
    store: Annotated[AuthStore, Depends(get_store)],
    current_settings: Annotated[Settings, Depends(get_settings)],
) -> LoginResponse:
    try:
        pair = auth.login(store, current_settings, body.email, body.password, _now())
    except auth.InvalidCredentials:
        raise _unauthorized("Invalid email or password") from None
    return LoginResponse(
        access_token=pair.access_token,
        refresh_token=pair.refresh_token,
        expires_in=pair.access_expires_in,
    )


@router.post("/auth/refresh", response_model=AccessTokenResponse)
def refresh(
    body: RefreshRequest,
    store: Annotated[AuthStore, Depends(get_store)],
    current_settings: Annotated[Settings, Depends(get_settings)],
) -> AccessTokenResponse:
    try:
        token = auth.refresh(store, current_settings, body.refresh_token, _now())
    except auth.InvalidToken:
        raise _unauthorized("Invalid or expired token") from None
    return AccessTokenResponse(access_token=token.token, expires_in=token.expires_in)


@router.post("/auth/logout", status_code=status.HTTP_204_NO_CONTENT)
def logout(
    body: RefreshRequest,
    access_token: Annotated[str, Depends(_bearer_token)],
    store: Annotated[AuthStore, Depends(get_store)],
) -> None:
    try:
        auth.logout(store, access_token, body.refresh_token, _now())
    except auth.InvalidToken:
        raise _unauthorized("Invalid or expired token") from None


@router.get("/auth/me", response_model=MeResponse)
def me(
    access_token: Annotated[str, Depends(_bearer_token)],
    store: Annotated[AuthStore, Depends(get_store)],
) -> MeResponse:
    try:
        user = auth.get_current_user(store, access_token, _now())
    except auth.InvalidToken:
        raise _unauthorized("Invalid or expired token") from None
    return MeResponse(id=user.id, email=user.email)


@router.get("/health")
def health() -> dict[str, str]:
    # Load balancers and uptime checks call this, so keep it free of I/O.
    return {"status": "ok"}
