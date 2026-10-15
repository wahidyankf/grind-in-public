from edutech_be.config import build_settings
from edutech_be.db import create_db_engine


def test_engine_uses_database_url_from_settings() -> None:
    engine = create_db_engine(
        build_settings({"DATABASE_URL": "postgresql+psycopg://u:p@h:5432/d"})
    )

    assert engine.url.drivername == "postgresql+psycopg"
    assert engine.url.database == "d"
    engine.dispose()
