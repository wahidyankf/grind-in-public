import os
from collections.abc import Mapping
from dataclasses import dataclass

# Matches the defaults in the root compose.yaml, so db:up works without setup.
DEFAULT_DATABASE_URL = "postgresql+psycopg://edutech_be:edutech_be@localhost:5132/edutech_be"

# Short access tokens limit the harm from a leak; refresh tokens keep users signed in.
DEFAULT_ACCESS_TOKEN_TTL_SECONDS = 900
DEFAULT_REFRESH_TOKEN_TTL_SECONDS = 604800


@dataclass(frozen=True)
class Settings:
    database_url: str
    access_token_ttl_seconds: int
    refresh_token_ttl_seconds: int


def build_settings(env: Mapping[str, str]) -> Settings:
    return Settings(
        database_url=env.get("DATABASE_URL", DEFAULT_DATABASE_URL),
        access_token_ttl_seconds=int(
            env.get("ACCESS_TOKEN_TTL_SECONDS", DEFAULT_ACCESS_TOKEN_TTL_SECONDS)
        ),
        refresh_token_ttl_seconds=int(
            env.get("REFRESH_TOKEN_TTL_SECONDS", DEFAULT_REFRESH_TOKEN_TTL_SECONDS)
        ),
    )


# Built once at import time, so the rest of the app reads a single object.
settings = build_settings(os.environ)
