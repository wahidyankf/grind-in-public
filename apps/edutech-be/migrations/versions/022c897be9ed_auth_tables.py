"""auth tables

Revision ID: 022c897be9ed
Revises: c6352f4dab8c
Create Date: 2026-10-15 11:50:00.398919

"""

from collections.abc import Sequence

from alembic import op

# revision identifiers, used by Alembic.
revision: str = "022c897be9ed"
down_revision: str | Sequence[str] | None = "c6352f4dab8c"
branch_labels: str | Sequence[str] | None = None
depends_on: str | Sequence[str] | None = None

# One statement per entry, because psycopg runs one statement per execute.
# No "%" or ":name" in these strings: they go through SQLAlchemy's text().
UPGRADE_STATEMENTS = (
    """
    CREATE TABLE users (
        id            uuid PRIMARY KEY DEFAULT gen_random_uuid(),
        email         text NOT NULL,
        password_hash text NOT NULL,
        created_at    timestamptz NOT NULL DEFAULT now(),
        created_by    uuid NOT NULL REFERENCES users (id),
        updated_at    timestamptz NOT NULL DEFAULT now(),
        updated_by    uuid NOT NULL REFERENCES users (id),
        deleted_at    timestamptz,
        deleted_by    uuid REFERENCES users (id),
        CHECK ((deleted_at IS NULL) = (deleted_by IS NULL))
    )
    """,
    "CREATE UNIQUE INDEX users_email_lower_key ON users (lower(email))",
    """
    CREATE TABLE auth_tokens (
        id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
        user_id    uuid NOT NULL REFERENCES users (id),
        kind       text NOT NULL CHECK (kind IN ('access', 'refresh')),
        token_hash text NOT NULL UNIQUE,
        expires_at timestamptz NOT NULL,
        revoked_at timestamptz,
        created_at timestamptz NOT NULL DEFAULT now(),
        created_by uuid NOT NULL REFERENCES users (id),
        updated_at timestamptz NOT NULL DEFAULT now(),
        updated_by uuid NOT NULL REFERENCES users (id),
        deleted_at timestamptz,
        deleted_by uuid REFERENCES users (id),
        CHECK ((deleted_at IS NULL) = (deleted_by IS NULL))
    )
    """,
    "CREATE INDEX auth_tokens_user_id_idx ON auth_tokens (user_id)",
    """
    CREATE FUNCTION set_updated_at() RETURNS trigger AS $$
    BEGIN
        NEW.updated_at = now();
        RETURN NEW;
    END;
    $$ LANGUAGE plpgsql
    """,
    """
    CREATE FUNCTION reject_delete() RETURNS trigger AS $$
    BEGIN
        RAISE EXCEPTION 'hard delete is not allowed';
    END;
    $$ LANGUAGE plpgsql
    """,
    """
    CREATE TRIGGER users_set_updated_at BEFORE UPDATE ON users
        FOR EACH ROW EXECUTE FUNCTION set_updated_at()
    """,
    """
    CREATE TRIGGER users_reject_delete BEFORE DELETE ON users
        FOR EACH ROW EXECUTE FUNCTION reject_delete()
    """,
    """
    CREATE TRIGGER auth_tokens_set_updated_at BEFORE UPDATE ON auth_tokens
        FOR EACH ROW EXECUTE FUNCTION set_updated_at()
    """,
    """
    CREATE TRIGGER auth_tokens_reject_delete BEFORE DELETE ON auth_tokens
        FOR EACH ROW EXECUTE FUNCTION reject_delete()
    """,
)

# Tables go first: they depend on the trigger functions only through their triggers,
# and auth_tokens references users, so it must be dropped before users.
DOWNGRADE_STATEMENTS = (
    "DROP TABLE auth_tokens",
    "DROP TABLE users",
    "DROP FUNCTION reject_delete()",
    "DROP FUNCTION set_updated_at()",
)


def upgrade() -> None:
    """Upgrade schema."""
    for statement in UPGRADE_STATEMENTS:
        op.execute(statement)


def downgrade() -> None:
    """Downgrade schema."""
    for statement in DOWNGRADE_STATEMENTS:
        op.execute(statement)
