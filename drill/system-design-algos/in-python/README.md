# System design algorithms in Python

Twenty data structures and algorithms that come up in system design, each with a small Python implementation, its tests,
and a guide that explains how it works. This is hands-on practice, so the code is meant to be read, typed out, and
broken on purpose.

The numbers follow the most used to least used order, based on how often each one appears in real systems and in system
design discussions. That order is a judgment, not measured data.

## Layout

- `algo_NN_<name>.md` explains how the algorithm works, with Mermaid diagrams, an example, and why it is useful.
- `algo_NN_<name>.py` holds the implementation.
- `algo_NN_<name>_test.py` holds its pytest tests, in the same folder so the two files open side by side.
- The `algo_` prefix is there because Python cannot import a module whose name starts with a digit. Without it, the
  tests could not import the code with full type checking.

Each implementation has a module docstring that says what it is, where it is used, and its complexity.

## Run

From this folder:

```bash
uv sync          # create .venv and install the dev tools
uv run pytest    # run every *_test.py file
uv run pyright   # type-check in strict mode
```

Python 3.13 or newer is required. The versions of pytest and pyright match `apps/edutech-be`, and every package is more
than a week old (`exclude-newer` in `pyproject.toml`).

## The algorithms

Each item links to its guide (`.md`), its code, and its tests.

1. **Hash table**: [guide](algo_01_hash_table.md), [code](algo_01_hash_table.py), [tests](algo_01_hash_table_test.py).
   O(1) on average. Separate chaining with resizing. Used in caches, session stores, and key-to-server maps.
2. **B+ tree**: [guide](algo_02_b_plus_tree.md), [code](algo_02_b_plus_tree.py), [tests](algo_02_b_plus_tree_test.py).
   O(log n) search and insert, and linked leaves for range scans. Used in database indexes (PostgreSQL, MySQL InnoDB).
3. **Queue and append-only log**: [guide](algo_03_queue_and_append_only_log.md),
   [code](algo_03_queue_and_append_only_log.py), [tests](algo_03_queue_and_append_only_log_test.py). O(1) enqueue and
   dequeue. A log keeps records, and each consumer group tracks its own offset. Used in message brokers such as Kafka.
4. **Heap**: [guide](algo_04_heap.md), [code](algo_04_heap.py), [tests](algo_04_heap_test.py). O(log n) push and pop,
   and top-k in O(n log k). Used in schedulers, timers, and "top N" queries.
5. **LRU cache**: [guide](algo_05_lru_cache.md), [code](algo_05_lru_cache.py), [tests](algo_05_lru_cache_test.py). O(1)
   get and put, built from a hash map and a doubly linked list. Used in application caches and cache eviction.
6. **Consistent hashing**: [guide](algo_06_consistent_hashing.md), [code](algo_06_consistent_hashing.py),
   [tests](algo_06_consistent_hashing_test.py). O(log v) lookup with virtual nodes. Adding or removing a node moves few
   keys. Used in Dynamo, Cassandra, and CDN routing.
7. **Bloom filter**: [guide](algo_07_bloom_filter.md), [code](algo_07_bloom_filter.py),
   [tests](algo_07_bloom_filter_test.py). O(k) add and lookup. "No" is always right, and "yes" can be a false positive.
   Used in LSM tree SSTables and duplicate checks.
8. **LSM tree**: [guide](algo_08_lsm_tree.md), [code](algo_08_lsm_tree.py), [tests](algo_08_lsm_tree_test.py). Writes go
   to memory, then flush to sorted runs. Deletes are tombstones, and compaction merges the runs. Used in Cassandra,
   RocksDB, and LevelDB.
9. **Inverted index**: [guide](algo_09_inverted_index.md), [code](algo_09_inverted_index.py),
   [tests](algo_09_inverted_index_test.py). Each word maps to a sorted posting list. An AND query intersects the lists.
   Used in full-text search (Lucene, Elasticsearch).
10. **Token bucket**: [guide](algo_10_token_bucket.md), [code](algo_10_token_bucket.py),
    [tests](algo_10_token_bucket_test.py). O(1) per request. Allows short bursts while holding the long-run rate. Used
    in API rate limiting.
11. **Raft (simplified)**: [guide](algo_11_raft.md), [code](algo_11_raft.py), [tests](algo_11_raft_test.py). Leader
    election and majority commit, run in memory with method calls instead of a network. Used in etcd, Consul, and
    CockroachDB. No timers, persistence, or membership changes.
12. **Graph search**: [guide](algo_12_graph_search.md), [code](algo_12_graph_search.py),
    [tests](algo_12_graph_search_test.py). BFS, topological order, and Dijkstra. O((V + E) log V) for Dijkstra. Used in
    routing, social graphs, and build order.
13. **Gossip protocol**: [guide](algo_13_gossip_protocol.md), [code](algo_13_gossip_protocol.py),
    [tests](algo_13_gossip_protocol_test.py). Random pairwise exchanges spread state in about log(n) rounds. Used in
    cluster membership (Cassandra, Serf).
14. **Trie**: [guide](algo_14_trie.md), [code](algo_14_trie.py), [tests](algo_14_trie_test.py). O(L) per word of length
    L, and autocomplete ranks matches by frequency. Used in search-as-you-type and longest-prefix routing.
15. **Merkle tree**: [guide](algo_15_merkle_tree.md), [code](algo_15_merkle_tree.py),
    [tests](algo_15_merkle_tree_test.py). Compares data by hash, so only the differing branches are checked. Used in
    Dynamo anti-entropy and Git.
16. **Snowflake-style IDs**: [guide](algo_16_snowflake_id.md), [code](algo_16_snowflake_id.py),
    [tests](algo_16_snowflake_id_test.py). 64-bit IDs made of time, worker, and sequence bits, unique across machines
    and sortable by time. Used in large services.
17. **HyperLogLog**: [guide](algo_17_hyperloglog.md), [code](algo_17_hyperloglog.py),
    [tests](algo_17_hyperloglog_test.py). Estimates distinct items with a few kilobytes. Used in Redis PFCOUNT and
    unique-visitor counts.
18. **Geohash and quadtree**: [guide](algo_18_geohash_quadtree.md), [code](algo_18_geohash_quadtree.py),
    [tests](algo_18_geohash_quadtree_test.py). Geohash turns a location into a prefix string, and the quadtree answers
    box queries. Used in "find nearby" features.
19. **Leaky bucket**: [guide](algo_19_leaky_bucket.md), [code](algo_19_leaky_bucket.py),
    [tests](algo_19_leaky_bucket_test.py). A bounded queue drained at a fixed rate. Used in traffic shaping and
    protecting slow services from bursts.
20. **Count-min sketch**: [guide](algo_20_count_min_sketch.md), [code](algo_20_count_min_sketch.py),
    [tests](algo_20_count_min_sketch_test.py). Approximate counts in fixed memory that never undercount, plus
    heavy-hitter tracking. Used in streaming analytics and hot keys.
