from edutech_be.config import (
    DEFAULT_ACCESS_TOKEN_TTL_SECONDS,
    DEFAULT_DATABASE_URL,
    DEFAULT_REFRESH_TOKEN_TTL_SECONDS,
    build_settings,
)


def test_database_url_defaults_to_local_compose_database() -> None:
    settings = build_settings({})

    assert settings.database_url == DEFAULT_DATABASE_URL


def test_database_url_comes_from_env() -> None:
    settings = build_settings(
        {"DATABASE_URL": "postgresql+psycopg://u:p@db:5432/other"}
    )

    assert settings.database_url == "postgresql+psycopg://u:p@db:5432/other"


def test_token_ttls_default_to_15_minutes_and_7_days() -> None:
    settings = build_settings({})

    assert settings.access_token_ttl_seconds == DEFAULT_ACCESS_TOKEN_TTL_SECONDS == 900
    assert (
        settings.refresh_token_ttl_seconds
        == DEFAULT_REFRESH_TOKEN_TTL_SECONDS
        == 604800
    )


def test_token_ttls_come_from_env() -> None:
    settings = build_settings(
        {
            "ACCESS_TOKEN_TTL_SECONDS": "60",
            "REFRESH_TOKEN_TTL_SECONDS": "3600",
        }
    )

    assert settings.access_token_ttl_seconds == 60
    assert settings.refresh_token_ttl_seconds == 3600
