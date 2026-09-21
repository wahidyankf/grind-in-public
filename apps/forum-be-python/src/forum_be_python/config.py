import os

# The credential-free URL matches the trust-auth container in compose.yaml, which only listens on loopback.
_DEFAULT_DATABASE_URL = "postgresql://postgres@127.0.0.1:54329/forum"

TOKEN_TTL_HOURS = 24


def database_url() -> str:
    return os.environ.get("FORUM_DATABASE_URL", _DEFAULT_DATABASE_URL)
