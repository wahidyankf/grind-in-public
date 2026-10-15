"""Tests for algo_08_lsm_tree.py."""

from algo_08_lsm_tree import LSMTree


def test_put_then_get() -> None:
    tree = LSMTree()
    tree.put("a", "1")
    assert tree.get("a") == "1"
    assert tree.get("missing") is None


def test_newest_value_wins_across_flushes() -> None:
    tree = LSMTree(memtable_limit=100)
    tree.put("a", "old")
    tree.flush()
    tree.put("a", "new")
    tree.flush()
    assert tree.run_count == 2
    assert tree.get("a") == "new"


def test_tombstone_hides_older_value_in_earlier_run() -> None:
    tree = LSMTree(memtable_limit=100)
    tree.put("a", "1")
    tree.flush()
    tree.delete("a")
    assert tree.get("a") is None
    tree.flush()
    assert tree.get("a") is None


def test_full_memtable_flushes_automatically() -> None:
    tree = LSMTree(memtable_limit=3)
    for i in range(7):
        tree.put(f"k{i}", str(i))
    assert tree.run_count == 2
    assert all(tree.get(f"k{i}") == str(i) for i in range(7))


def test_flushed_run_is_sorted() -> None:
    tree = LSMTree(memtable_limit=100)
    for key in ["c", "a", "b"]:
        tree.put(key, key.upper())
    tree.flush()
    assert tree.runs[0] == [("a", "A"), ("b", "B"), ("c", "C")]


def test_compaction_keeps_latest_and_drops_tombstones() -> None:
    tree = LSMTree(memtable_limit=100)
    tree.put("a", "1")
    tree.put("b", "2")
    tree.flush()
    tree.put("a", "3")
    tree.delete("b")
    tree.flush()
    tree.compact()
    assert tree.run_count == 1
    assert tree.get("a") == "3"
    assert tree.get("b") is None


def test_invalid_memtable_limit_is_rejected() -> None:
    try:
        LSMTree(memtable_limit=0)
    except ValueError:
        return
    raise AssertionError("expected ValueError")
