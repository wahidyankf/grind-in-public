---
tldr: "Teaches interview algorithms in Python by connecting invariants and complexity to production problems."
when_to_use: "Use after Python foundations to learn problem-solving patterns, then practise without reading solutions."
---

# Python Algorithms Interview

This course teaches algorithms as tools for choosing operations under constraints. Each lesson starts with a production
problem, derives the data structure or pattern, proves the invariant, and only then writes Python. The same reasoning is
what an interviewer needs: clarify the contract, choose an operation, explain correctness, and account for cost.

```text
production need -> required operations -> invariant -> data structure -> Python -> tests -> trade-offs
```

## Reading order

1. [Interview Method and Complexity](001-interview-method-and-complexity.md)
2. [Arrays, Strings, and Hashing](002-arrays-strings-and-hashing.md)
3. [Linked Structures, Stacks, and Queues](003-linked-structures-stacks-and-queues.md)
4. [Pointers, Windows, and Prefix Sums](004-pointers-windows-and-prefix-sums.md)
5. [Binary Search, Sorting, and Selection](005-binary-search-sorting-and-selection.md)
6. [Intervals and Sweep Lines](006-intervals-and-sweep-lines.md)
7. [Trees and Tries](007-trees-and-tries.md)
8. [Heaps and Streaming Statistics](008-heaps-and-streaming-statistics.md)
9. [Graph Traversal and Ordering](009-graph-traversal-and-ordering.md)
10. [Connectivity and Shortest Paths](010-connectivity-and-shortest-paths.md)
11. [Recursion and Backtracking](011-recursion-and-backtracking.md)
12. [Greedy and Dynamic Programming](012-greedy-and-dynamic-programming.md)
13. [Production Algorithms](013-production-algorithms.md)

After the lessons, use the [drill prompts](drills/README.md) without opening the paired
[solutions](solutions/README.md). The drill policy still applies: solve and explain the exercise yourself before using
the solution for review.

All code targets Python 3.14 and uses the standard library. Examples are deliberately small enough to read in NVIM, but
they define every referenced name and state their expected output.
