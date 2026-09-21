import hashlib
import secrets

_TOKEN_BYTES = 32


def new_token() -> str:
    return secrets.token_urlsafe(_TOKEN_BYTES)


def hash_token(token: str) -> str:
    # Only this digest is stored, so a leaked database cannot be replayed as a bearer token.
    return hashlib.sha256(token.encode()).hexdigest()
