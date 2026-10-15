"""Bloom filter.

Used in: skipping disk reads in LSM trees (Cassandra and RocksDB keep one per SSTable), duplicate checks in crawlers,
and "have I seen this?" checks before an expensive lookup.

Complexity: add and might_contain are O(k) for k hash functions. Memory is about 9.6 bits per item at a 1% false
positive rate.

A "no" is always right. A "yes" can be a false positive, and the filter cannot remove items.
"""

import hashlib
import math


class BloomFilter:
    def __init__(self, bit_count: int, hash_count: int) -> None:
        if bit_count < 1 or hash_count < 1:
            raise ValueError("bit_count and hash_count must be at least 1")
        self._bit_count = bit_count
        self._hash_count = hash_count
        self._bits = bytearray((bit_count + 7) // 8)

    @property
    def bit_count(self) -> int:
        return self._bit_count

    @classmethod
    def from_capacity(cls, expected_items: int, false_positive_rate: float) -> "BloomFilter":
        """Size the filter for expected_items at the target false positive rate."""
        if expected_items < 1:
            raise ValueError("expected_items must be at least 1")
        if not 0.0 < false_positive_rate < 1.0:
            raise ValueError("false_positive_rate must be between 0 and 1")
        bit_count = math.ceil(-expected_items * math.log(false_positive_rate) / (math.log(2) ** 2))
        hash_count = max(1, round(bit_count / expected_items * math.log(2)))
        return cls(bit_count, hash_count)

    def add(self, item: str) -> None:
        for index in self._indexes(item):
            self._bits[index >> 3] |= 1 << (index & 7)

    def might_contain(self, item: str) -> bool:
        return all(self._bits[index >> 3] & (1 << (index & 7)) for index in self._indexes(item))

    def _indexes(self, item: str) -> list[int]:
        # Double hashing: two base hashes generate every index, so one digest is enough for any hash count.
        digest = hashlib.sha256(item.encode()).digest()
        first = int.from_bytes(digest[:8], "big")
        second = int.from_bytes(digest[8:16], "big") | 1
        return [(first + i * second) % self._bit_count for i in range(self._hash_count)]
