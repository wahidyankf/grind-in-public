from dataclasses import replace
from datetime import datetime
from uuid import UUID

from edutech_be.store import TokenRecord, User


class FakeAuthStore:
    """In-memory AuthStore for unit tests. Keeps the same rules as the SQL store."""

    def __init__(self) -> None:
        self.users: dict[UUID, User] = {}
        self.tokens: dict[str, TokenRecord] = {}  # keyed by token_hash

    def find_user_by_email(self, email: str) -> User | None:
        return next((u for u in self.users.values() if u.email == email), None)

    def find_user_by_id(self, user_id: UUID) -> User | None:
        return self.users.get(user_id)

    def insert_user(self, user: User) -> bool:
        if self.find_user_by_email(user.email) is not None:
            return False
        self.users[user.id] = user
        return True

    def insert_token(self, token: TokenRecord) -> None:
        self.tokens[token.token_hash] = token

    def find_token_by_hash(self, token_hash: str) -> TokenRecord | None:
        return self.tokens.get(token_hash)

    def revoke_token(
        self, token_id: UUID, revoked_at: datetime, actor_id: UUID
    ) -> None:
        for token_hash, token in self.tokens.items():
            if token.id == token_id:
                self.tokens[token_hash] = replace(token, revoked_at=revoked_at)
