"""Tests for algo_18_geohash_quadtree.py."""

import random

import pytest

from algo_18_geohash_quadtree import Box, Point, QuadTree, geohash_decode, geohash_encode


def test_known_geohash_example() -> None:
    assert geohash_encode(57.64911, 10.40744, precision=11) == "u4pruydqqvj"


def test_decode_returns_a_point_inside_the_cell() -> None:
    lat, lon = geohash_decode(geohash_encode(57.64911, 10.40744, precision=11))
    assert abs(lat - 57.64911) < 0.0001
    assert abs(lon - 10.40744) < 0.0001


def test_shorter_geohash_is_a_prefix_of_longer_one() -> None:
    long_code = geohash_encode(-33.8688, 151.2093, precision=9)
    assert geohash_encode(-33.8688, 151.2093, precision=5) == long_code[:5]


def test_invalid_coordinates_are_rejected() -> None:
    with pytest.raises(ValueError):
        geohash_encode(91.0, 0.0)


def test_quadtree_query_matches_brute_force() -> None:
    rng = random.Random(42)
    points = [Point(rng.uniform(0, 100), rng.uniform(0, 100), f"p{i}") for i in range(200)]
    tree = QuadTree(Box(0, 0, 100, 100))
    for point in points:
        assert tree.insert(point)
    area = Box(10, 10, 40, 60)
    expected = sorted(point.name for point in points if area.contains(point))
    assert sorted(point.name for point in tree.query(area)) == expected


def test_point_outside_boundary_is_rejected() -> None:
    tree = QuadTree(Box(0, 0, 10, 10))
    assert tree.insert(Point(20, 20, "far")) is False


def test_many_identical_points_do_not_recurse_forever() -> None:
    tree = QuadTree(Box(0, 0, 10, 10))
    for i in range(50):
        assert tree.insert(Point(5, 5, f"same-{i}"))
    assert len(tree.query(Box(4, 4, 6, 6))) == 50
