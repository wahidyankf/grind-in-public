---
tldr: "Maps the curriculum to primary papers and official operational documentation for deeper study."
when_to_use: "Use after a lesson or case to verify a concept at its source instead of memorizing a summary."
---

# Primary References

Read references to answer a specific design question. Start with the lesson and case, identify the uncertain mechanism,
then return with a one-paragraph summary of the guarantee, assumption, operating cost, and rejection condition.

## Reliability and operations

- [Site Reliability Engineering](https://sre.google/sre-book/table-of-contents/) covers SLOs, monitoring, automation,
  overload, change, and incident response.
- [The Site Reliability Workbook](https://sre.google/workbook/table-of-contents/) turns those principles into practices.
- [Timeouts, retries, and backoff with jitter](https://aws.amazon.com/builders-library/timeouts-retries-and-backoff-with-jitter/)
  explains retry amplification and bounded client behaviour.
- [Kubernetes concepts](https://kubernetes.io/docs/concepts/) defines workload, scheduling, configuration, security, and
  cluster primitives used in the deployment lessons.

## Distributed storage and consistency

- [Dynamo: a highly available key-value store](https://www.amazon.science/publications/dynamo-amazons-highly-available-key-value-store)
  motivates consistent hashing, object versions, quorums, and application conflict handling.
- [Bigtable](https://research.google/pubs/bigtable-a-distributed-storage-system-for-structured-data/) connects a sorted,
  distributed map to large structured workloads.
- [Spanner](https://research.google/pubs/spanner-googles-globally-distributed-database-2/) explains synchronous
  replication, global transactions, and clock uncertainty.
- [Kafka protocol and design](https://kafka.apache.org/documentation/) documents partition ordering, consumer groups,
  retention, replication, and delivery semantics.

These systems made different trade-offs for different workloads. Do not copy one paper's architecture without its
assumptions.

## Datastores and serving

- [PostgreSQL](https://www.postgresql.org/docs/current/), [Cassandra](https://cassandra.apache.org/doc/latest/),
  [HBase](https://hbase.apache.org/book.html), and [MongoDB](https://www.mongodb.com/docs/manual/) are the canonical
  references for the datastore families used in this course.
- [Redis](https://redis.io/docs/latest/) documents memory-oriented structures and persistence/replication choices.
- [OpenSearch](https://docs.opensearch.org/latest/) documents text indexes, mappings, aliases, and distributed search.

## Migration and evolution

- [Transparent migration of Datastore to Firestore](https://research.google/pubs/transparent-migration-of-datastore-to-firestore/)
  is a production account of a large, non-disruptive datastore migration.
- [Expand and contract pattern](https://martinfowler.com/bliki/ParallelChange.html) explains compatible interface
  evolution through parallel change.

## Study method

For each source, record:

```text
problem -> workload/failure assumptions -> mechanism -> guarantee -> cost -> where it breaks
```

That note is interview-ready because it connects the source to judgment rather than to a product name.
