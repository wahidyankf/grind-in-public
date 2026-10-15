# 08. LSM tree

A log-structured merge (LSM) tree turns random writes into sequential ones. Writes go to memory first. When memory is
full, the data is written to disk as one sorted run. Reads check the newest data first.

Code: [algo_08_lsm_tree.py](algo_08_lsm_tree.py). Tests: [algo_08_lsm_tree_test.py](algo_08_lsm_tree_test.py).

## The idea

- The **memtable** is an in-memory map. Every write goes here first.
- When the memtable is full, it is **flushed** as a **sorted run**. Runs never change after they are written.
- A **delete** writes a **tombstone** (`None`). It hides older values in the runs below it.
- A read checks the memtable first, then the runs from newest to oldest. The first match wins.
- **Compaction** merges all runs into one. It keeps only the newest value per key and drops tombstones.

Writing sorted files in one go is much faster on disk than updating values in place.

## How it works

Put and delete:

1. Store the key and value (or tombstone) in the memtable.
2. If the memtable has reached its limit, flush it as a new sorted run.

Get:

1. If the key is in the memtable, return its value (`None` means deleted).
2. Otherwise, for each run from newest to oldest, binary-search for the key.
3. The first run that has the key decides the answer. If no run has it, return `None`.

Compact: merge runs from oldest to newest into one dictionary, so newer values overwrite older ones. Then keep only the
non-tombstone entries, sorted.

## Flow

```text
put(key, value)  or  delete(key)  (delete stores None)
      |
      v
memtable[key] = value
      |
      v
memtable has reached the limit? --no--> done
      |
     yes
      v
flush: write memtable as a sorted run, then clear memtable


get(key)
   |
   v
key in memtable? --yes--> return memtable[key]
   | no
   v
for each run, newest first:
   key in run? --yes--> return its value (None if tombstone)
   |
   v
return None
```

## Example

`memtable_limit = 3`.

```text
put a=1, b=2, c=3           memtable full -> flush
                            run 1: [a=1, b=2, c=3]

put a=10, delete b, put d=4 memtable full -> flush
                            run 2: [a=10, b=None, d=4]

get("a")   -> not in memtable, run 2 has a=10          -> 10
get("b")   -> run 2 has a tombstone                    -> None
get("c")   -> run 2 has no c, run 1 has c=3            -> 3

compact    -> merge run 1 and run 2, newest wins, drop tombstone
              run 1: [a=10, c=3, d=4]
```

## Complexity

- `put`, `delete`: O(1) amortised, plus the cost of a flush.
- Flush: O(m log m) to sort m items.
- `get`: O(r log n) for r runs and n keys per run. Bloom filters reduce the runs that get searched.

## Why it is useful

Disks are much faster at sequential writes than random ones. An LSM tree batches writes in memory and writes them out in
order, so a database can take a very high write rate. Reads are slower than in a B+ tree, so this design suits
write-heavy workloads such as logs, time series, and event data.

## Where it shows up

Cassandra, RocksDB, LevelDB, HBase, and ScyllaDB. Most write-heavy databases use this design.

## Simplifications in this code

- Runs live in memory. A real engine writes them to files and merges them in levels.
- Compaction merges every run at once. Real engines merge small runs often, so no single merge is very large.
- No write-ahead log. A crash loses the memtable.
