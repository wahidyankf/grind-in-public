import hashlib
import secrets


def generate_token() -> str:
    return secrets.token_urlsafe(32)


def hash_token(token: str) -> str:
    # Only the hash is stored, so a DB leak does not reveal usable tokens.
    return hashlib.sha256(token.encode("utf-8")).hexdigest()
