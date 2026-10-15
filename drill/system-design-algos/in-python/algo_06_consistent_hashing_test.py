"""Tests for algo_06_consistent_hashing.py."""

import pytest
from algo_06_consistent_hashing import ConsistentHashRing


def _keys(count: int) -> list[str]:
    return [f"user:{i}" for i in range(count)]


def test_empty_ring_raises() -> None:
    ring = ConsistentHashRing()
    with pytest.raises(LookupError):
        ring.get_node("x")


def test_same_key_always_maps_to_same_node() -> None:
    ring = ConsistentHashRing()
    for name in ["cache-a", "cache-b", "cache-c"]:
        ring.add_node(name)
    first = [ring.get_node(key) for key in _keys(200)]
    second = [ring.get_node(key) for key in _keys(200)]
    assert first == second


def test_removing_a_node_moves_only_its_keys() -> None:
    ring = ConsistentHashRing()
    for name in ["cache-a", "cache-b", "cache-c"]:
        ring.add_node(name)
    before = {key: ring.get_node(key) for key in _keys(2000)}
    ring.remove_node("cache-b")
    for key, owner in before.items():
        if owner != "cache-b":
            assert ring.get_node(key) == owner


def test_adding_a_node_moves_a_minority_of_keys() -> None:
    ring = ConsistentHashRing()
    for name in ["cache-a", "cache-b", "cache-c"]:
        ring.add_node(name)
    before = {key: ring.get_node(key) for key in _keys(2000)}
    ring.add_node("cache-d")
    moved = sum(1 for key, owner in before.items() if ring.get_node(key) != owner)
    assert 0 < moved < 2000 * 0.4


def test_virtual_nodes_spread_keys_across_nodes() -> None:
    ring = ConsistentHashRing()
    names = ["cache-a", "cache-b", "cache-c"]
    for name in names:
        ring.add_node(name)
    counts = {name: 0 for name in names}
    for key in _keys(3000):
        counts[ring.get_node(key)] += 1
    assert all(0.2 * 3000 < count < 0.5 * 3000 for count in counts.values())


def test_len_counts_distinct_nodes() -> None:
    ring = ConsistentHashRing(virtual_nodes=10)
    ring.add_node("a")
    ring.add_node("b")
    assert len(ring) == 2
    ring.remove_node("a")
    assert len(ring) == 1
