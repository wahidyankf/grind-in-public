"""HyperLogLog: estimate how many distinct items you have seen, in a few kilobytes.

Used in: Redis PFCOUNT, unique-visitor counts, and distinct-value counts in analytics. With 2^14 registers the sketch
uses about 16 KB and has roughly 0.8% standard error.

Complexity: add is O(1). estimate and merge are O(m) for m registers.

Each item hashes to 64 bits. The top p bits pick a register. The register keeps the largest rank it has seen, where the
rank is the position of the first 1 bit in the remaining bits.
"""

import hashlib
import math

HASH_BITS = 64


def _alpha(register_count: int) -> float:
    """Bias correction constant from the HyperLogLog paper."""
    if register_count == 16:
        return 0.673
    if register_count == 32:
        return 0.697
    if register_count == 64:
        return 0.709
    return 0.7213 / (1 + 1.079 / register_count)


class HyperLogLog:
    def __init__(self, precision: int = 12) -> None:
        if not 4 <= precision <= 16:
            raise ValueError("precision must be between 4 and 16")
        self._precision = precision
        self._register_count = 1 << precision
        self._registers: list[int] = [0] * self._register_count

    def add(self, item: str) -> None:
        digest = hashlib.sha256(item.encode()).digest()
        value = int.from_bytes(digest[:8], "big")
        suffix_bits = HASH_BITS - self._precision
        index = value >> suffix_bits
        rest = value & ((1 << suffix_bits) - 1)
        rank = suffix_bits - rest.bit_length() + 1
        if rank > self._registers[index]:
            self._registers[index] = rank

    def estimate(self) -> int:
        m = self._register_count
        raw = _alpha(m) * m * m / sum(2.0**-register for register in self._registers)
        empty_registers = self._registers.count(0)
        if raw <= 2.5 * m and empty_registers > 0:
            # Linear counting is more accurate than the raw formula for small counts.
            return round(m * math.log(m / empty_registers))
        return round(raw)

    def merge(self, other: "HyperLogLog") -> None:
        """Fold another sketch of the same precision into this one. The result counts the union."""
        if other._precision != self._precision:
            raise ValueError("sketches must have the same precision")
        self._registers = [max(mine, theirs) for mine, theirs in zip(self._registers, other._registers)]
