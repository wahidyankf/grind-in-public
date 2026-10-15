"""baseline

Revision ID: c6352f4dab8c
Revises:
Create Date: 2026-10-15 10:17:15.431810

"""

from collections.abc import Sequence

# revision identifiers, used by Alembic.
revision: str = "c6352f4dab8c"
down_revision: str | Sequence[str] | None = None
branch_labels: str | Sequence[str] | None = None
depends_on: str | Sequence[str] | None = None


def upgrade() -> None:
    """Upgrade schema."""
    pass


def downgrade() -> None:
    """Downgrade schema."""
    pass
