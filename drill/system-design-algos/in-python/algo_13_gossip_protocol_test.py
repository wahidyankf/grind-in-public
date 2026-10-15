"""Tests for algo_13_gossip_protocol.py."""

from algo_13_gossip_protocol import GossipCluster


def test_news_reaches_every_node() -> None:
    cluster = GossipCluster([f"node-{i}" for i in range(8)], seed=1)
    cluster.nodes[0].write("config", "v1")
    rounds = cluster.run_until_converged()
    assert rounds is not None
    assert all(node.state["config"].value == "v1" for node in cluster.nodes)


def test_higher_version_wins() -> None:
    cluster = GossipCluster(["a", "b"], seed=2)
    cluster.nodes[0].write("key", "old")
    cluster.nodes[1].write("key", "first")
    cluster.nodes[1].write("key", "second")
    assert cluster.run_until_converged() is not None
    assert cluster.nodes[0].state["key"].value == "second"


def test_already_converged_cluster_needs_zero_rounds() -> None:
    cluster = GossipCluster(["a", "b", "c"], seed=3)
    assert cluster.run_until_converged() == 0


def test_same_seed_gives_same_round_count() -> None:
    def run() -> int | None:
        cluster = GossipCluster([f"n{i}" for i in range(6)], seed=7)
        cluster.nodes[0].write("k", "v")
        return cluster.run_until_converged()

    assert run() == run()


def test_single_node_cluster_is_always_converged() -> None:
    cluster = GossipCluster(["solo"])
    cluster.run_round()
    assert cluster.converged() is True
