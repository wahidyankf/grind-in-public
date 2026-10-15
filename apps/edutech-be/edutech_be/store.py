from dataclasses import dataclass
from datetime import datetime
from typing import Literal, Protocol
from uuid import UUID

TokenKind = Literal["access", "refresh"]


@dataclass(frozen=True)
class User:
    id: UUID
    email: str
    password_hash: str
    created_at: datetime


@dataclass(frozen=True)
class TokenRecord:
    id: UUID
    user_id: UUID
    kind: TokenKind
    token_hash: str
    expires_at: datetime
    revoked_at: datetime | None


class AuthStore(Protocol):
    """Data access for the auth service. Soft-deleted rows are never returned."""

    def find_user_by_email(self, email: str) -> User | None: ...

    def find_user_by_id(self, user_id: UUID) -> User | None: ...

    def insert_user(self, user: User) -> bool:
        """Return False when the email is already taken."""
        ...

    def insert_token(self, token: TokenRecord) -> None: ...

    def find_token_by_hash(self, token_hash: str) -> TokenRecord | None: ...

    def revoke_token(
        self, token_id: UUID, revoked_at: datetime, actor_id: UUID
    ) -> None: ...
