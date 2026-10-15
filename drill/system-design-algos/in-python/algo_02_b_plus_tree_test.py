"""Tests for algo_02_b_plus_tree.py."""

from algo_02_b_plus_tree import BPlusTree


def _build(keys: list[int], order: int = 4) -> BPlusTree[int, str]:
    tree: BPlusTree[int, str] = BPlusTree(order=order)
    for key in keys:
        tree.insert(key, f"v{key}")
    return tree


def test_search_finds_every_inserted_key() -> None:
    keys = [7, 3, 9, 1, 5, 11, 2, 8, 4, 10, 6, 0]
    tree = _build(keys)
    assert len(tree) == len(keys)
    assert all(tree.search(key) == f"v{key}" for key in keys)


def test_search_missing_key_returns_none() -> None:
    tree = _build([1, 2, 3])
    assert tree.search(99) is None


def test_insert_overwrites_existing_key() -> None:
    tree = _build([1, 2, 3])
    tree.insert(2, "new")
    assert tree.search(2) == "new"
    assert len(tree) == 3


def test_range_scan_is_inclusive_and_ordered() -> None:
    tree = _build(list(range(0, 50)))
    result = tree.range(10, 19)
    assert [key for key, _ in result] == list(range(10, 20))


def test_items_cover_all_leaves_in_order() -> None:
    keys = list(range(40, 0, -1))
    tree = _build(keys)
    assert [key for key, _ in tree.items()] == sorted(keys)


def test_splits_up_to_a_new_root_keep_every_key_findable() -> None:
    tree = _build(list(range(100)), order=3)
    assert len(tree) == 100
    assert tree.range(0, 99)[-1] == (99, "v99")
    assert tree.search(57) == "v57"


def test_rejects_tiny_order() -> None:
    try:
        BPlusTree[int, int](order=2)
    except ValueError:
        return
    raise AssertionError("expected ValueError")
