"""Log-structured merge tree (LSM tree).

Used in: write-heavy storage engines: Cassandra, RocksDB, LevelDB, HBase, and ScyllaDB.

Complexity: put and delete are O(1) amortised, plus the cost of a flush. A flush sorts the memtable, O(m log m) for m
items. A get checks the memtable and then each sorted run, O(r log n) for r runs.

Writes go to an in-memory memtable. A full memtable is flushed as one sorted run. Deletes write a tombstone (None).
Compaction merges all runs into one and drops the tombstones.
"""

from bisect import bisect_left

type Run = list[tuple[str, str | None]]


class LSMTree:
    def __init__(self, memtable_limit: int = 4) -> None:
        if memtable_limit < 1:
            raise ValueError("memtable_limit must be at least 1")
        self._memtable_limit = memtable_limit
        self._memtable: dict[str, str | None] = {}
        self._runs: list[Run] = []

    @property
    def runs(self) -> list[Run]:
        return [list(run) for run in self._runs]

    @property
    def run_count(self) -> int:
        return len(self._runs)

    def put(self, key: str, value: str) -> None:
        self._write(key, value)

    def delete(self, key: str) -> None:
        self._write(key, None)

    def get(self, key: str) -> str | None:
        if key in self._memtable:
            return self._memtable[key]
        for run in reversed(self._runs):  # newest run first, so its value wins
            found, value = _search_run(run, key)
            if found:
                return value
        return None

    def flush(self) -> None:
        if not self._memtable:
            return
        self._runs.append(sorted(self._memtable.items()))
        self._memtable = {}

    def compact(self) -> None:
        """Merge every run into one. Tombstones are dropped, because nothing older is left to hide."""
        merged: dict[str, str | None] = {}
        for run in self._runs:
            merged.update(run)
        self._runs = [sorted((key, value) for key, value in merged.items() if value is not None)]

    def _write(self, key: str, value: str | None) -> None:
        self._memtable[key] = value
        if len(self._memtable) >= self._memtable_limit:
            self.flush()


def _first(pair: tuple[str, str | None]) -> str:
    return pair[0]


def _search_run(run: Run, key: str) -> tuple[bool, str | None]:
    i = bisect_left(run, key, key=_first)
    if i < len(run) and run[i][0] == key:
        return True, run[i][1]
    return False, None
