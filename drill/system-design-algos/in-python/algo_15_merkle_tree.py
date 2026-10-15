"""Merkle tree.

Used in: replica repair (Dynamo and Cassandra anti-entropy), Git commits and trees, certificate transparency, and
blockchains. Two replicas compare root hashes first, and only the differing branches get checked further.

Complexity: building costs O(n) hashes. A membership proof is O(log n) hashes. Finding d differing leaves costs about
O(d log n) hash comparisons.

When a level has an odd count, its last hash pairs with itself. No padding is stored.
"""

import hashlib


def _leaf_hash(block: bytes) -> str:
    return hashlib.sha256(block).hexdigest()


def _pair_hash(left: str, right: str) -> str:
    return hashlib.sha256(bytes.fromhex(left) + bytes.fromhex(right)).hexdigest()


class MerkleTree:
    def __init__(self, blocks: list[bytes]) -> None:
        if not blocks:
            raise ValueError("a Merkle tree needs at least one block")
        # levels[0] holds the leaves, and the last level holds the root.
        self._levels: list[list[str]] = [[_leaf_hash(block) for block in blocks]]
        while len(self._levels[-1]) > 1:
            level = self._levels[-1]
            parents = [
                _pair_hash(level[i], level[i + 1] if i + 1 < len(level) else level[i])
                for i in range(0, len(level), 2)
            ]
            self._levels.append(parents)

    @property
    def root(self) -> str:
        return self._levels[-1][0]

    @property
    def leaf_count(self) -> int:
        return len(self._levels[0])

    def proof(self, index: int) -> list[tuple[str, bool]]:
        """Return (sibling hash, sibling is on the left) pairs, from the leaf up to the root."""
        if not 0 <= index < self.leaf_count:
            raise IndexError("leaf index out of range")
        steps: list[tuple[str, bool]] = []
        position = index
        for level in self._levels[:-1]:
            sibling = position ^ 1
            if sibling < len(level):
                steps.append((level[sibling], sibling < position))
            else:
                steps.append((level[position], False))  # an odd node was paired with itself
            position //= 2
        return steps

    @staticmethod
    def verify(block: bytes, proof: list[tuple[str, bool]], root: str) -> bool:
        current = _leaf_hash(block)
        for sibling, sibling_is_left in proof:
            if sibling_is_left:
                current = _pair_hash(sibling, current)
            else:
                current = _pair_hash(current, sibling)
        return current == root

    def differing_leaves(self, other: "MerkleTree") -> list[int]:
        """Return the leaf indexes where the two trees disagree, checking only the branches that differ."""
        if self.leaf_count != other.leaf_count:
            raise ValueError("trees must have the same number of leaves")
        if self.root == other.root:
            return []
        return self._diff(other, len(self._levels) - 1, 0)

    def _diff(self, other: "MerkleTree", level: int, index: int) -> list[int]:
        if self._levels[level][index] == other._levels[level][index]:
            return []
        if level == 0:
            return [index]
        found: list[int] = []
        for child in (2 * index, 2 * index + 1):
            if child < len(self._levels[level - 1]):
                found.extend(self._diff(other, level - 1, child))
        return found
