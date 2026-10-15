import pytest

from edutech_be.passwords import hash_password, is_acceptable_password, verify_password


def test_verify_accepts_the_password_it_hashed() -> None:
    password_hash = hash_password("correct horse")

    assert verify_password(password_hash, "correct horse") is True


def test_verify_rejects_a_wrong_password() -> None:
    password_hash = hash_password("correct horse")

    assert verify_password(password_hash, "wrong horse") is False


def test_hash_does_not_contain_the_plain_password() -> None:
    password_hash = hash_password("correct horse")

    assert "correct horse" not in password_hash


def test_password_shorter_than_8_characters_is_rejected() -> None:
    assert is_acceptable_password("abcdefg") is False


def test_password_of_8_characters_is_accepted() -> None:
    assert is_acceptable_password("abcdefgh") is True


def test_password_of_64_characters_is_accepted() -> None:
    assert is_acceptable_password("a" * 64) is True


def test_password_longer_than_64_characters_is_rejected() -> None:
    assert is_acceptable_password("a" * 65) is False


def test_length_counts_characters_not_bytes() -> None:
    # 8 characters, but 24 bytes in UTF-8.
    assert is_acceptable_password("日本語のパスワード") is True


def test_spaces_and_symbols_are_allowed() -> None:
    assert is_acceptable_password("        ") is True
    assert is_acceptable_password("!@#$%^&*") is True


@pytest.mark.parametrize("password", ["password", "12345678", "PASSWORD", "Qwerty123"])
def test_common_passwords_are_rejected_ignoring_case(password: str) -> None:
    assert is_acceptable_password(password) is False
