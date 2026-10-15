"""Simplified Raft consensus.

Used in: etcd, Consul, CockroachDB, and TiKV. Raft picks one leader per term, and an entry counts as committed only
after a majority of nodes store it, so a committed command survives a leader change.

Complexity: one replication round sends one call per follower. Each entry is stored on every node.

This version drops timers, persistence, the network, snapshots, and membership changes. Method calls stand in for RPCs,
so the caller drives elections and heartbeats by hand.
"""

from dataclasses import dataclass
from enum import Enum


class Role(Enum):
    FOLLOWER = "follower"
    CANDIDATE = "candidate"
    LEADER = "leader"


@dataclass(frozen=True)
class LogEntry:
    term: int
    command: str


class RaftNode:
    def __init__(self, node_id: int, cluster_size: int) -> None:
        self.node_id = node_id
        self._cluster_size = cluster_size
        self.current_term = 0
        self.voted_for: int | None = None
        self.leader_id: int | None = None
        self.log: list[LogEntry] = []
        self.commit_index = -1
        self.last_applied = -1
        self.applied: list[str] = []
        self.role = Role.FOLLOWER
        self.next_index: dict[int, int] = {}
        self.match_index: dict[int, int] = {}

    def is_leader(self) -> bool:
        return self.role is Role.LEADER

    def handle_request_vote(
        self,
        term: int,
        candidate_id: int,
        last_log_index: int,
        last_log_term: int,
    ) -> tuple[int, bool]:
        if term < self.current_term:
            return self.current_term, False
        if term > self.current_term:
            self._become_follower(term)
        # A voter only backs a candidate whose log is at least as up to date as its own.
        candidate_is_current = (last_log_term, last_log_index) >= (self._last_term(), len(self.log) - 1)
        if self.voted_for in (None, candidate_id) and candidate_is_current:
            self.voted_for = candidate_id
            return self.current_term, True
        return self.current_term, False

    def handle_append_entries(
        self,
        term: int,
        leader_id: int,
        prev_log_index: int,
        prev_log_term: int,
        entries: list[LogEntry],
        leader_commit: int,
    ) -> tuple[int, bool]:
        if term < self.current_term:
            return self.current_term, False
        self._become_follower(term)
        self.leader_id = leader_id
        if prev_log_index >= 0 and (
            prev_log_index >= len(self.log) or self.log[prev_log_index].term != prev_log_term
        ):
            return self.current_term, False
        start = prev_log_index + 1
        for offset, entry in enumerate(entries):
            position = start + offset
            if position < len(self.log):
                if self.log[position].term == entry.term:
                    continue
                # A conflicting entry means this log has diverged from the leader's, so drop the rest.
                del self.log[position:]
            self.log.append(entry)
        if leader_commit > self.commit_index:
            self.commit_index = min(leader_commit, start + len(entries) - 1)
        self._apply_committed()
        return self.current_term, True

    def start_election(self, peers: list["RaftNode"]) -> bool:
        """Run one election. Returns True if this node won a majority of votes."""
        self.role = Role.CANDIDATE
        self.current_term += 1
        self.voted_for = self.node_id
        votes = 1
        for peer in peers:
            term, granted = peer.handle_request_vote(
                self.current_term,
                self.node_id,
                len(self.log) - 1,
                self._last_term(),
            )
            if term > self.current_term:
                self._become_follower(term)
                return False
            if granted:
                votes += 1
        if votes * 2 <= self._cluster_size:
            self.role = Role.FOLLOWER
            return False
        self.role = Role.LEADER
        self.leader_id = self.node_id
        self.next_index = {peer.node_id: len(self.log) for peer in peers}
        self.match_index = {peer.node_id: -1 for peer in peers}
        return True

    def propose(self, command: str) -> int:
        if not self.is_leader():
            raise RuntimeError("only the leader accepts commands")
        self.log.append(LogEntry(self.current_term, command))
        return len(self.log) - 1

    def replicate(self, peers: list["RaftNode"]) -> None:
        """Send entries to every peer, then commit and apply what a majority stores."""
        if not self.is_leader():
            return
        for peer in peers:
            self._send_to(peer)
        if self.is_leader():
            self._advance_commit(peers)
            self._apply_committed()

    def _send_to(self, peer: "RaftNode") -> None:
        next_index = self.next_index[peer.node_id]
        prev_index = next_index - 1
        prev_term = self.log[prev_index].term if prev_index >= 0 else 0
        entries = self.log[next_index:]
        term, success = peer.handle_append_entries(
            self.current_term,
            self.node_id,
            prev_index,
            prev_term,
            entries,
            self.commit_index,
        )
        if term > self.current_term:
            self._become_follower(term)
            return
        if success:
            self.match_index[peer.node_id] = prev_index + len(entries)
            self.next_index[peer.node_id] = self.match_index[peer.node_id] + 1
        else:
            self.next_index[peer.node_id] = max(0, next_index - 1)

    def _advance_commit(self, peers: list["RaftNode"]) -> None:
        for index in range(len(self.log) - 1, self.commit_index, -1):
            # A leader commits only entries from its own term. Older entries commit along with them.
            if self.log[index].term != self.current_term:
                continue
            replicas = 1 + sum(1 for peer in peers if self.match_index[peer.node_id] >= index)
            if replicas * 2 > self._cluster_size:
                self.commit_index = index
                return

    def _apply_committed(self) -> None:
        while self.last_applied < self.commit_index:
            self.last_applied += 1
            self.applied.append(self.log[self.last_applied].command)

    def _last_term(self) -> int:
        return self.log[-1].term if self.log else 0

    def _become_follower(self, term: int) -> None:
        if term > self.current_term:
            self.current_term = term
            self.voted_for = None
        self.role = Role.FOLLOWER


class RaftCluster:
    def __init__(self, size: int) -> None:
        if size < 1:
            raise ValueError("a cluster needs at least one node")
        self.nodes = [RaftNode(node_id, size) for node_id in range(size)]

    def elect(self, node_id: int) -> bool:
        node = self.nodes[node_id]
        return node.start_election(self._peers_of(node))

    def leader(self) -> RaftNode | None:
        leaders = [node for node in self.nodes if node.is_leader()]
        return max(leaders, key=lambda node: node.current_term, default=None)

    def submit(self, command: str) -> int:
        leader = self.leader()
        if leader is None:
            raise RuntimeError("no leader to accept the command")
        index = leader.propose(command)
        leader.replicate(self._peers_of(leader))
        return index

    def heartbeat(self) -> None:
        """Send an empty append round so followers learn the newest commit index."""
        leader = self.leader()
        if leader is not None:
            leader.replicate(self._peers_of(leader))

    def _peers_of(self, node: RaftNode) -> list[RaftNode]:
        return [other for other in self.nodes if other is not node]
