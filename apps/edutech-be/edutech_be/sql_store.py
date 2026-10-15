from datetime import datetime
from uuid import UUID

from sqlalchemy import (
    Column,
    ColumnElement,
    Connection,
    DateTime,
    MetaData,
    Table,
    Text,
    Uuid,
    func,
    select,
    update,
)
from sqlalchemy.dialects.postgresql import insert

from edutech_be.store import TokenRecord, User

# Query-builder tables only. Schema changes go in hand-written Alembic revisions.
metadata = MetaData()

users = Table(
    "users",
    metadata,
    Column("id", Uuid, primary_key=True),
    Column("email", Text, nullable=False),
    Column("password_hash", Text, nullable=False),
    Column("created_at", DateTime(timezone=True), nullable=False),
    Column("created_by", Uuid, nullable=False),
    Column("updated_at", DateTime(timezone=True), nullable=False),
    Column("updated_by", Uuid, nullable=False),
    Column("deleted_at", DateTime(timezone=True)),
    Column("deleted_by", Uuid),
)

auth_tokens = Table(
    "auth_tokens",
    metadata,
    Column("id", Uuid, primary_key=True),
    Column("user_id", Uuid, nullable=False),
    Column("kind", Text, nullable=False),
    Column("token_hash", Text, nullable=False, unique=True),
    Column("expires_at", DateTime(timezone=True), nullable=False),
    Column("revoked_at", DateTime(timezone=True)),
    Column("created_at", DateTime(timezone=True), nullable=False),
    Column("created_by", Uuid, nullable=False),
    Column("updated_at", DateTime(timezone=True), nullable=False),
    Column("updated_by", Uuid, nullable=False),
    Column("deleted_at", DateTime(timezone=True)),
    Column("deleted_by", Uuid),
)


class SqlAuthStore:
    """AuthStore on SQLAlchemy Core. Never commits; the caller owns the transaction."""

    def __init__(self, connection: Connection) -> None:
        self._connection = connection

    def find_user_by_email(self, email: str) -> User | None:
        return self._find_user(func.lower(users.c.email) == email.lower())

    def find_user_by_id(self, user_id: UUID) -> User | None:
        return self._find_user(users.c.id == user_id)

    def insert_user(self, user: User) -> bool:
        # No conflict target: the only unique rules are the email index and the id key.
        statement = (
            insert(users)
            .values(
                id=user.id,
                email=user.email,
                password_hash=user.password_hash,
                created_at=user.created_at,
                created_by=user.id,
                updated_at=user.created_at,
                updated_by=user.id,
            )
            .on_conflict_do_nothing()
            # RETURNING gives a row only when the insert happened. rowcount is -1 here.
            .returning(users.c.id)
        )
        return self._connection.execute(statement).first() is not None

    def insert_token(self, token: TokenRecord) -> None:
        # created_at and updated_at come from the database defaults.
        self._connection.execute(
            insert(auth_tokens).values(
                id=token.id,
                user_id=token.user_id,
                kind=token.kind,
                token_hash=token.token_hash,
                expires_at=token.expires_at,
                revoked_at=token.revoked_at,
                created_by=token.user_id,
                updated_by=token.user_id,
            )
        )

    def find_token_by_hash(self, token_hash: str) -> TokenRecord | None:
        row = (
            self._connection.execute(
                select(auth_tokens).where(
                    auth_tokens.c.token_hash == token_hash,
                    auth_tokens.c.deleted_at.is_(None),
                )
            )
            .mappings()
            .first()
        )
        if row is None:
            return None
        return TokenRecord(
            id=row["id"],
            user_id=row["user_id"],
            kind=row["kind"],
            token_hash=row["token_hash"],
            expires_at=row["expires_at"],
            revoked_at=row["revoked_at"],
        )

    def revoke_token(
        self, token_id: UUID, revoked_at: datetime, actor_id: UUID
    ) -> None:
        self._connection.execute(
            update(auth_tokens)
            .where(auth_tokens.c.id == token_id)
            .values(revoked_at=revoked_at, updated_by=actor_id)
        )

    def _find_user(self, condition: ColumnElement[bool]) -> User | None:
        row = (
            self._connection.execute(
                select(users).where(condition, users.c.deleted_at.is_(None))
            )
            .mappings()
            .first()
        )
        if row is None:
            return None
        return User(
            id=row["id"],
            email=row["email"],
            password_hash=row["password_hash"],
            created_at=row["created_at"],
        )
