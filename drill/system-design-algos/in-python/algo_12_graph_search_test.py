"""Tests for algo_12_graph_search.py."""

import pytest

from algo_12_graph_search import Graph, bfs_shortest_path, dijkstra, topological_order


def test_bfs_takes_the_fewest_edges() -> None:
    graph = Graph()
    graph.add_edge("a", "b")
    graph.add_edge("b", "c")
    graph.add_edge("c", "d")
    graph.add_edge("a", "d")
    assert bfs_shortest_path(graph, "a", "d") == ["a", "d"]


def test_bfs_returns_none_when_unreachable() -> None:
    graph = Graph()
    graph.add_edge("a", "b")
    assert bfs_shortest_path(graph, "b", "a") is None


def test_topological_order_puts_prerequisites_first() -> None:
    graph = Graph()
    graph.add_edge("wake", "shower")
    graph.add_edge("shower", "dress")
    graph.add_edge("wake", "breakfast")
    graph.add_edge("breakfast", "dress")
    order = topological_order(graph)
    assert order.index("wake") < order.index("shower") < order.index("dress")
    assert order.index("wake") < order.index("breakfast") < order.index("dress")


def test_topological_order_rejects_a_cycle() -> None:
    graph = Graph()
    graph.add_edge("a", "b")
    graph.add_edge("b", "a")
    with pytest.raises(ValueError):
        topological_order(graph)


def test_dijkstra_prefers_the_cheaper_longer_path() -> None:
    graph = Graph()
    graph.add_edge("a", "b", 1.0)
    graph.add_edge("b", "c", 1.0)
    graph.add_edge("a", "c", 5.0)
    distance = dijkstra(graph, "a")
    assert distance["c"] == 2.0
    assert distance["a"] == 0.0


def test_dijkstra_leaves_out_unreachable_nodes() -> None:
    graph = Graph()
    graph.add_edge("a", "b")
    graph.add_node("z")
    assert "z" not in dijkstra(graph, "a")
