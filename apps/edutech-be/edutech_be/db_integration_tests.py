import pytest
from alembic import command
from sqlalchemy import inspect, literal, select

from edutech_be.config import settings
from edutech_be.db import create_db_engine
from edutech_be.migrations import alembic_config

# Needs the Postgres container from `npx nx db:up edutech-be`.
pytestmark = pytest.mark.integration


def test_query_builder_runs_against_postgres() -> None:
    engine = create_db_engine(settings)

    with engine.connect() as connection:
        value = connection.execute(select(literal(1))).scalar_one()

    engine.dispose()
    assert value == 1


def test_migrations_reach_head() -> None:
    command.upgrade(alembic_config(), "head")

    engine = create_db_engine(settings)
    with engine.connect() as connection:
        tables = inspect(connection).get_table_names()
    engine.dispose()

    assert "alembic_version" in tables
