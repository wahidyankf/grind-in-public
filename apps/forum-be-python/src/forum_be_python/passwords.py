import hashlib
import hmac
import secrets

_SCHEME = "scrypt"
_N = 2**14
_R = 8
_P = 1
_KEY_BYTES = 32
_SALT_BYTES = 16


def _derive_key(password: str, salt: bytes) -> bytes:
    return hashlib.scrypt(password.encode(), salt=salt, n=_N, r=_R, p=_P, dklen=_KEY_BYTES)


def hash_password(password: str) -> str:
    salt = secrets.token_bytes(_SALT_BYTES)
    return f"{_SCHEME}${_N}${_R}${_P}${salt.hex()}${_derive_key(password, salt).hex()}"


def verify_password(password: str, stored: str) -> bool:
    parts = stored.split("$")
    if len(parts) != 6 or parts[0] != _SCHEME:
        return False
    try:
        salt = bytes.fromhex(parts[4])
        expected = bytes.fromhex(parts[5])
    except ValueError:
        return False
    return hmac.compare_digest(_derive_key(password, salt), expected)
