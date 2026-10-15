"""Tests for algo_15_merkle_tree.py."""

import pytest

from algo_15_merkle_tree import MerkleTree


def _blocks(count: int) -> list[bytes]:
    return [f"block-{i}".encode() for i in range(count)]


def test_every_proof_verifies_for_an_odd_leaf_count() -> None:
    blocks = _blocks(5)
    tree = MerkleTree(blocks)
    for index, block in enumerate(blocks):
        assert MerkleTree.verify(block, tree.proof(index), tree.root)


def test_tampered_block_fails_verification() -> None:
    blocks = _blocks(4)
    tree = MerkleTree(blocks)
    assert not MerkleTree.verify(b"forged", tree.proof(1), tree.root)


def test_single_block_tree_root_is_its_leaf_hash() -> None:
    tree = MerkleTree([b"only"])
    assert tree.proof(0) == []
    assert MerkleTree.verify(b"only", [], tree.root)


def test_same_blocks_give_the_same_root() -> None:
    assert MerkleTree(_blocks(7)).root == MerkleTree(_blocks(7)).root


def test_diff_finds_the_one_changed_leaf() -> None:
    original = _blocks(6)
    changed = list(original)
    changed[4] = b"changed"
    assert MerkleTree(original).differing_leaves(MerkleTree(changed)) == [4]


def test_identical_trees_have_no_difference() -> None:
    assert MerkleTree(_blocks(6)).differing_leaves(MerkleTree(_blocks(6))) == []


def test_diff_rejects_different_leaf_counts() -> None:
    with pytest.raises(ValueError):
        MerkleTree(_blocks(3)).differing_leaves(MerkleTree(_blocks(4)))


def test_empty_input_is_rejected() -> None:
    with pytest.raises(ValueError):
        MerkleTree([])


def test_proof_index_out_of_range_raises() -> None:
    with pytest.raises(IndexError):
        MerkleTree(_blocks(2)).proof(2)
