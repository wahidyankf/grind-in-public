from pathlib import Path

from alembic.config import Config

APP_ROOT = Path(__file__).resolve().parents[1]


def alembic_config() -> Config:
    # Absolute paths let tests and commands run from any directory.
    config = Config(str(APP_ROOT / "alembic.ini"))
    config.set_main_option("script_location", str(APP_ROOT / "migrations"))
    return config
