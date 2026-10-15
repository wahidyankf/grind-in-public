"""Tests for algo_11_raft.py."""

from algo_11_raft import RaftCluster


def test_elected_leader_is_the_candidate_that_won() -> None:
    cluster = RaftCluster(3)
    assert cluster.elect(0) is True
    leader = cluster.leader()
    assert leader is not None
    assert leader.node_id == 0


def test_candidate_needs_a_majority() -> None:
    cluster = RaftCluster(5)
    # Only node 1 is reachable, so 2 of 5 votes is not a majority.
    assert cluster.nodes[0].start_election([cluster.nodes[1]]) is False
    # Nodes 1 and 2 are reachable, so 3 of 5 votes is a majority.
    assert cluster.nodes[0].start_election([cluster.nodes[1], cluster.nodes[2]]) is True


def test_committed_commands_reach_every_node() -> None:
    cluster = RaftCluster(3)
    cluster.elect(0)
    cluster.submit("a")
    cluster.submit("b")
    cluster.heartbeat()
    assert all(node.applied == ["a", "b"] for node in cluster.nodes)


def test_new_leader_keeps_committed_entries() -> None:
    cluster = RaftCluster(3)
    cluster.elect(0)
    cluster.submit("a")
    assert cluster.elect(1) is True
    cluster.submit("b")
    cluster.heartbeat()
    assert cluster.nodes[1].applied == ["a", "b"]
    assert cluster.nodes[2].applied == ["a", "b"]


def test_stale_log_does_not_win_a_vote() -> None:
    cluster = RaftCluster(3)
    cluster.elect(0)
    cluster.submit("a")
    # Node 2 has an empty log, so node 1, which holds "a", must refuse its vote.
    _, granted = cluster.nodes[1].handle_request_vote(
        term=3,
        candidate_id=2,
        last_log_index=-1,
        last_log_term=0,
    )
    assert granted is False


def test_append_from_an_old_term_is_rejected() -> None:
    cluster = RaftCluster(3)
    follower = cluster.nodes[1]
    follower.current_term = 5
    term, success = follower.handle_append_entries(
        term=2,
        leader_id=0,
        prev_log_index=-1,
        prev_log_term=0,
        entries=[],
        leader_commit=-1,
    )
    assert (term, success) == (5, False)


def test_submit_without_a_leader_raises() -> None:
    cluster = RaftCluster(3)
    try:
        cluster.submit("x")
    except RuntimeError:
        return
    raise AssertionError("expected RuntimeError")
