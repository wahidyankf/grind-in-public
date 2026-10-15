from functools import cache
from pathlib import Path

from argon2 import PasswordHasher
from argon2.exceptions import VerifyMismatchError

# NIST SP 800-63B: at least 8 characters, support at least 64.
MIN_LENGTH = 8
MAX_LENGTH = 64

COMMON_PASSWORDS_PATH = Path(__file__).with_name("common_passwords.txt")

_hasher = PasswordHasher()


def hash_password(password: str) -> str:
    return _hasher.hash(password)


def verify_password(password_hash: str, password: str) -> bool:
    try:
        return _hasher.verify(password_hash, password)
    except VerifyMismatchError:
        return False


def is_acceptable_password(password: str) -> bool:
    if not MIN_LENGTH <= len(password) <= MAX_LENGTH:
        return False
    return password.casefold() not in _common_passwords()


@cache
def _common_passwords() -> frozenset[str]:
    lines = COMMON_PASSWORDS_PATH.read_text(encoding="utf-8").splitlines()
    return frozenset(line.casefold() for line in lines if line)
