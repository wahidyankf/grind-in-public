---
tldr: "Uses hierarchical trees for structured traversal and tries for prefix-based candidate retrieval."
when_to_use: "Use when relationships form a hierarchy or lookup is driven by prefixes rather than whole keys."
---

# Trees and Tries

Trees encode parent-child structure. Binary-search trees add ordering, but their performance depends on balance. Tries
index keys by prefix and trade memory for predictable prefix traversal.

## Production motivation: prefix candidates

A screening pipeline can use a trie to retrieve candidates sharing a normalized prefix before applying an expensive
similarity score. The trie narrows candidates; it does not decide whether two identities match.

```python
from dataclasses import dataclass, field


@dataclass
class TrieNode:
    children: dict[str, "TrieNode"] = field(default_factory=dict)
    values: list[str] = field(default_factory=list)


@dataclass
class PrefixIndex:
    root: TrieNode = field(default_factory=TrieNode)

    def add(self, key: str, value: str) -> None:
        node = self.root
        for character in key:
            node = node.children.setdefault(character, TrieNode())
        node.values.append(value)

    def exact(self, key: str) -> tuple[str, ...]:
        node = self.root
        for character in key:
            next_node = node.children.get(character)
            if next_node is None:
                return ()
            node = next_node
        return tuple(node.values)


index = PrefixIndex()
index.add("al", "ALPHA")
index.add("al", "ALTO")
assert index.exact("al") == ("ALPHA", "ALTO")
```

```text
root
 `- a
     `- l -> [ALPHA, ALTO]
```

Insertion and lookup cost `O(m)` for key length `m`; memory can be large because every node owns a dictionary.
Compressed tries, finite-state structures, or search-engine indexes are better for large immutable corpora.

## Tree traversal in production

Folder trees, organizational hierarchies, and reply threads need DFS or BFS. Recursive traversal risks Python's call
depth on untrusted trees; an explicit stack makes the memory bound visible. Database hierarchies add cycles, missing
parents, authorization, pagination, and concurrent updates.

## Checkpoint

Explain why candidate retrieval and match scoring should be separate stages, and which metrics reveal an overly broad
prefix index.
