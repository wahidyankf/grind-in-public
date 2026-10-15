"""Trie (prefix tree).

Used in: autocomplete and search-as-you-type, longest-prefix matching in IP routing tables, and spell checkers.

Complexity: insert, search, and starts_with cost O(L) for a word of length L. autocomplete walks the subtree under the
prefix, so it costs O(size of that subtree) plus a sort of the matches.
"""

class _TrieNode:
    def __init__(self) -> None:
        self.children: dict[str, _TrieNode] = {}
        self.frequency = 0  # 0 means no word ends at this node


class Trie:
    def __init__(self) -> None:
        self._root = _TrieNode()

    def insert(self, word: str, count: int = 1) -> None:
        node = self._root
        for char in word:
            node = node.children.setdefault(char, _TrieNode())
        node.frequency += count

    def search(self, word: str) -> bool:
        node = self._find(word)
        return node is not None and node.frequency > 0

    def starts_with(self, prefix: str) -> bool:
        return self._find(prefix) is not None

    def autocomplete(self, prefix: str, limit: int = 5) -> list[str]:
        """Return up to limit words under the prefix, most frequent first, then alphabetical."""
        start = self._find(prefix)
        if start is None:
            return []
        found: list[tuple[int, str]] = []
        stack: list[tuple[_TrieNode, str]] = [(start, prefix)]
        while stack:
            node, text = stack.pop()
            if node.frequency > 0:
                found.append((node.frequency, text))
            for char, child in node.children.items():
                stack.append((child, text + char))
        found.sort(key=lambda item: (-item[0], item[1]))
        return [word for _, word in found[:limit]]

    def _find(self, text: str) -> _TrieNode | None:
        node = self._root
        for char in text:
            child = node.children.get(char)
            if child is None:
                return None
            node = child
        return node
