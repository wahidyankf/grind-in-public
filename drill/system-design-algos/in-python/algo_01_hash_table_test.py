"""Tests for algo_01_hash_table.py."""

from algo_01_hash_table import HashTable, INITIAL_BUCKETS


class _Clash:
    """Every instance hashes to the same bucket, so the tests exercise collision handling."""

    def __init__(self, name: str) -> None:
        self.name = name

    def __hash__(self) -> int:
        return 1

    def __eq__(self, other: object) -> bool:
        return isinstance(other, _Clash) and other.name == self.name


def test_put_and_get() -> None:
    table: HashTable[str, int] = HashTable()
    table.put("a", 1)
    assert table.get("a") == 1
    assert table.get("missing") is None
    assert len(table) == 1


def test_put_overwrites_without_growing() -> None:
    table: HashTable[str, int] = HashTable()
    table.put("a", 1)
    table.put("a", 2)
    assert table.get("a") == 2
    assert len(table) == 1


def test_delete_returns_whether_key_existed() -> None:
    table: HashTable[str, int] = HashTable()
    table.put("a", 1)
    assert table.delete("a") is True
    assert table.delete("a") is False
    assert table.get("a") is None
    assert len(table) == 0


def test_collisions_are_kept_apart() -> None:
    table: HashTable[_Clash, int] = HashTable()
    a, b, c = _Clash("a"), _Clash("b"), _Clash("c")
    table.put(a, 1)
    table.put(b, 2)
    table.put(c, 3)
    assert table.delete(b) is True
    assert table.get(a) == 1
    assert table.get(b) is None
    assert table.get(c) == 3


def test_resize_keeps_every_key() -> None:
    table: HashTable[int, str] = HashTable()
    for i in range(100):
        table.put(i, str(i))
    assert table.bucket_count > INITIAL_BUCKETS
    assert len(table) == 100
    assert all(table.get(i) == str(i) for i in range(100))
