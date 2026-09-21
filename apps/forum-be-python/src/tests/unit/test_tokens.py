import re

from forum_be_python.tokens import hash_token, new_token


def test_new_token_is_url_safe_and_differs_on_every_call() -> None:
    first = new_token()
    second = new_token()

    assert re.fullmatch(r"[A-Za-z0-9_-]{43}", first)
    assert first != second


def test_hash_token_is_a_stable_digest_that_does_not_contain_the_token() -> None:
    token = "synthetic-token-value"

    digest = hash_token(token)

    assert digest == hash_token(token)
    assert re.fullmatch(r"[0-9a-f]{64}", digest)
    assert token not in digest
