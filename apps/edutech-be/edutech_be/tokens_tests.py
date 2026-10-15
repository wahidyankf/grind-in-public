from edutech_be.tokens import generate_token, hash_token


def test_generated_tokens_are_long_and_unique() -> None:
    first = generate_token()
    second = generate_token()

    assert len(first) >= 43  # 32 random bytes, base64url without padding
    assert first != second


def test_hash_is_64_hex_characters() -> None:
    token_hash = hash_token("some-token")

    assert len(token_hash) == 64
    assert set(token_hash) <= set("0123456789abcdef")


def test_hash_is_stable_for_the_same_token() -> None:
    assert hash_token("some-token") == hash_token("some-token")


def test_hash_differs_from_the_raw_token() -> None:
    token = generate_token()

    assert hash_token(token) != token
    assert token not in hash_token(token)
