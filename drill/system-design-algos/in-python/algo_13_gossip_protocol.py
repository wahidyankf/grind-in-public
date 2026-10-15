"""Gossip protocol: nodes share state with random peers until they agree.

Used in: cluster membership and failure detection (Cassandra, Consul's Serf, and Redis Cluster's bus). Each round, every
node swaps state with one random peer, so news reaches the whole cluster in about log(n) rounds.

Complexity: one round costs one exchange per node. Rounds to converge grow about as log(n).

Each key carries a version, and the higher version wins. This version has no failure timers. It only shows how state
spreads.
"""

import random
from dataclasses import dataclass


@dataclass(frozen=True)
class Entry:
    version: int
    value: str


class GossipNode:
    def __init__(self, name: str) -> None:
        self.name = name
        self.state: dict[str, Entry] = {}

    def write(self, key: str, value: str) -> None:
        current = self.state.get(key)
        version = 1 if current is None else current.version + 1
        self.state[key] = Entry(version, value)

    def merge(self, incoming: dict[str, Entry]) -> None:
        for key, entry in incoming.items():
            current = self.state.get(key)
            if current is None or entry.version > current.version:
                self.state[key] = entry


class GossipCluster:
    def __init__(self, names: list[str], seed: int = 0) -> None:
        self.nodes = [GossipNode(name) for name in names]
        # A fixed seed makes every run repeat, so tests can count rounds exactly.
        self._random = random.Random(seed)

    def run_round(self) -> None:
        if len(self.nodes) < 2:
            return
        for node in self.nodes:
            peer = self._random.choice([other for other in self.nodes if other is not node])
            node_snapshot = dict(node.state)
            peer_snapshot = dict(peer.state)
            node.merge(peer_snapshot)
            peer.merge(node_snapshot)

    def converged(self) -> bool:
        if not self.nodes:
            return True
        first = self.nodes[0].state
        return all(node.state == first for node in self.nodes)

    def run_until_converged(self, max_rounds: int = 100) -> int | None:
        """Return how many rounds it took, or None if the cluster has not converged by max_rounds."""
        for round_number in range(max_rounds + 1):
            if self.converged():
                return round_number
            self.run_round()
        return None
