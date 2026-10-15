"""Tests for algo_04_heap.py."""

import pytest
from algo_04_heap import MinHeap, top_k


def test_pop_returns_items_in_ascending_order() -> None:
    heap: MinHeap[int] = MinHeap()
    for value in [5, 1, 9, 3, 7, 2]:
        heap.push(value)
    assert [heap.pop() for _ in range(6)] == [1, 2, 3, 5, 7, 9]


def test_peek_does_not_remove() -> None:
    heap: MinHeap[int] = MinHeap()
    heap.push(4)
    heap.push(2)
    assert heap.peek() == 2
    assert len(heap) == 2


def test_empty_heap_raises() -> None:
    heap: MinHeap[int] = MinHeap()
    with pytest.raises(IndexError):
        heap.pop()
    with pytest.raises(IndexError):
        heap.peek()


def test_top_k_returns_largest_in_descending_order() -> None:
    assert top_k([3, 8, 1, 9, 4, 7], 3) == [9, 8, 7]


def test_top_k_when_k_exceeds_input() -> None:
    assert top_k([2, 1], 5) == [2, 1]


def test_top_k_zero_is_empty() -> None:
    assert top_k([1, 2, 3], 0) == []


def test_top_k_on_strings() -> None:
    assert top_k(["pear", "apple", "zebra"], 1) == ["zebra"]
