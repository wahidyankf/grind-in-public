import pytest

from forum_be_python.passwords import hash_password, verify_password


def test_verify_accepts_the_password_that_was_hashed() -> None:
    stored = hash_password("correct horse battery staple")

    assert verify_password("correct horse battery staple", stored) is True


def test_verify_rejects_a_different_password() -> None:
    stored = hash_password("correct horse battery staple")

    assert verify_password("wrong horse battery staple", stored) is False


def test_hashing_the_same_password_twice_gives_different_hashes() -> None:
    first = hash_password("correct horse battery staple")
    second = hash_password("correct horse battery staple")

    assert first != second


@pytest.mark.parametrize(
    "stored",
    ["", "not-a-hash", "scrypt$16384$8$1$zz$zz", "bcrypt$16384$8$1$00$00"],
)
def test_verify_returns_false_for_a_malformed_stored_hash(stored: str) -> None:
    assert verify_password("correct horse battery staple", stored) is False
