from typing import Annotated, cast

from fastapi import Depends, HTTPException, Request
from fastapi.security import HTTPAuthorizationCredentials, HTTPBearer

from forum_be_python.db import Pool
from forum_be_python.schemas import UserOut
from forum_be_python.tokens import hash_token

# auto_error is off so a missing header and a bad token both leave through the same 401 below.
_bearer = HTTPBearer(auto_error=False)


def _get_pool(request: Request) -> Pool:
    return cast(Pool, request.app.state.pool)


PoolDep = Annotated[Pool, Depends(_get_pool)]


def _current_user(
    credentials: Annotated[HTTPAuthorizationCredentials | None, Depends(_bearer)], pool: PoolDep
) -> UserOut:
    unauthorized = HTTPException(status_code=401, detail="Not authenticated", headers={"WWW-Authenticate": "Bearer"})
    if credentials is None:
        raise unauthorized
    with pool.connection() as conn:
        row = conn.execute(
            "SELECT u.id, u.username FROM auth_tokens t JOIN users u ON u.id = t.user_id"
            " WHERE t.token_hash = %s AND t.expires_at > now()",
            (hash_token(credentials.credentials),),
        ).fetchone()
    if row is None:
        raise unauthorized
    return UserOut(id=row[0], username=row[1])


CurrentUser = Annotated[UserOut, Depends(_current_user)]
