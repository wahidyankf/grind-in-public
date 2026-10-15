from sqlalchemy import Engine, create_engine

from edutech_be.config import Settings


def create_db_engine(settings: Settings) -> Engine:
    # pool_pre_ping drops dead connections, such as after the container restarts.
    return create_engine(settings.database_url, pool_pre_ping=True)
