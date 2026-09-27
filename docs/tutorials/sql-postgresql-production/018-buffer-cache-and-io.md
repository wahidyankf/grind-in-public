---
tldr: "Traces page reads and writes through shared buffers, operating-system cache, checkpoints, and PostgreSQL 18 I/O."
when_to_use: "Use when interpreting buffer metrics, cache misses, write spikes, or storage latency."
---

# Buffer Cache and I/O

Backends access table and index pages through shared buffers. A miss reads from storage, usually via the operating
system page cache. Modified pages become dirty; WAL must be durable before the corresponding data page is written.

```text
executor -> shared buffer hit ------------------------> page
             |
             +-> miss -> OS/filesystem -> storage -> buffer

dirty buffer -> background/checkpoint write -> storage
commit WAL ------------------------------------^ must precede data page
```

PostgreSQL uses a clock-sweep replacement strategy rather than a perfect LRU. Each buffer has a usage count; scans age
candidates and reuse pages whose count reaches zero. This approximates recency with bounded coordination cost. Large
sequential scans use a small buffer-access ring so they do not evict the entire working set.

PostgreSQL 18 expands asynchronous I/O support and exposes related statistics. Async reads increase concurrency but do
not make slow storage disappear; queue depth, request size, cache effectiveness, and CPU still matter. Use the views
available in the running minor version rather than copying a dashboard query blindly.

```sql
SELECT backend_type, object, context, reads, read_bytes, read_time,
       writes, write_bytes, write_time
FROM pg_stat_io
ORDER BY read_bytes + write_bytes DESC;
```

`shared_buffers` is not “all available memory.” Leave memory for connections, per-node work memory, maintenance, and the
operating system. Tune from workload evidence and tail latency, then load-test checkpoint behaviour.
