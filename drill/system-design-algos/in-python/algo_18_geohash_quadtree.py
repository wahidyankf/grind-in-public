"""Geohash and point quadtree for location search.

Used in: "find nearby" queries for ride-hailing and delivery apps, map tiles, and geo indexes (Redis GEO stores geohash
style scores, and MongoDB's 2dsphere index uses a similar tree).

Complexity: geohash encode and decode are O(precision). A quadtree query costs about O(log n + k) for k results when
points are spread out. Clusters make the tree deeper, so depth is capped.

A geohash turns (lat, lon) into a string. Nearby places usually share a prefix. The quadtree splits a box into four
quadrants when a cell gets full.
"""

from dataclasses import dataclass

GEOHASH_ALPHABET = "0123456789bcdefghjkmnpqrstuvwxyz"
MAX_DEPTH = 16


def geohash_encode(lat: float, lon: float, precision: int = 9) -> str:
    if not -90.0 <= lat <= 90.0 or not -180.0 <= lon <= 180.0:
        raise ValueError("latitude must be in [-90, 90] and longitude in [-180, 180]")
    if precision < 1:
        raise ValueError("precision must be at least 1")
    lat_low, lat_high = -90.0, 90.0
    lon_low, lon_high = -180.0, 180.0
    chars: list[str] = []
    bits = 0
    bit_count = 0
    use_longitude = True  # bits alternate, starting with longitude
    while len(chars) < precision:
        if use_longitude:
            mid = (lon_low + lon_high) / 2
            if lon >= mid:
                bits = bits * 2 + 1
                lon_low = mid
            else:
                bits = bits * 2
                lon_high = mid
        else:
            mid = (lat_low + lat_high) / 2
            if lat >= mid:
                bits = bits * 2 + 1
                lat_low = mid
            else:
                bits = bits * 2
                lat_high = mid
        use_longitude = not use_longitude
        bit_count += 1
        if bit_count == 5:
            chars.append(GEOHASH_ALPHABET[bits])
            bits = 0
            bit_count = 0
    return "".join(chars)


def geohash_decode(code: str) -> tuple[float, float]:
    """Return the centre (lat, lon) of the geohash cell."""
    lat_low, lat_high = -90.0, 90.0
    lon_low, lon_high = -180.0, 180.0
    use_longitude = True
    for char in code:
        value = GEOHASH_ALPHABET.index(char)
        for shift in range(4, -1, -1):
            bit = (value >> shift) & 1
            if use_longitude:
                mid = (lon_low + lon_high) / 2
                if bit:
                    lon_low = mid
                else:
                    lon_high = mid
            else:
                mid = (lat_low + lat_high) / 2
                if bit:
                    lat_low = mid
                else:
                    lat_high = mid
            use_longitude = not use_longitude
    return (lat_low + lat_high) / 2, (lon_low + lon_high) / 2


@dataclass(frozen=True)
class Point:
    x: float
    y: float
    name: str


@dataclass(frozen=True)
class Box:
    min_x: float
    min_y: float
    max_x: float
    max_y: float

    def contains(self, point: Point) -> bool:
        return self.min_x <= point.x <= self.max_x and self.min_y <= point.y <= self.max_y

    def intersects(self, other: "Box") -> bool:
        return not (
            other.min_x > self.max_x
            or other.max_x < self.min_x
            or other.min_y > self.max_y
            or other.max_y < self.min_y
        )


class QuadTree:
    def __init__(self, boundary: Box, capacity: int = 4, depth: int = 0) -> None:
        self._boundary = boundary
        self._capacity = capacity
        self._depth = depth
        self._points: list[Point] = []
        self._children: list[QuadTree] | None = None

    def insert(self, point: Point) -> bool:
        if not self._boundary.contains(point):
            return False
        children = self._children
        if children is None:
            if len(self._points) < self._capacity or self._depth >= MAX_DEPTH:
                self._points.append(point)
                return True
            children = self._split()
        return _insert_into(children, point)

    def query(self, area: Box) -> list[Point]:
        """Return every stored point inside area."""
        if not area.intersects(self._boundary):
            return []
        found = [point for point in self._points if area.contains(point)]
        children = self._children
        if children is not None:
            for child in children:
                found.extend(child.query(area))
        return found

    def _split(self) -> list["QuadTree"]:
        box = self._boundary
        mid_x = (box.min_x + box.max_x) / 2
        mid_y = (box.min_y + box.max_y) / 2
        quadrants = [
            Box(box.min_x, box.min_y, mid_x, mid_y),
            Box(mid_x, box.min_y, box.max_x, mid_y),
            Box(box.min_x, mid_y, mid_x, box.max_y),
            Box(mid_x, mid_y, box.max_x, box.max_y),
        ]
        children = [QuadTree(quadrant, self._capacity, self._depth + 1) for quadrant in quadrants]
        self._children = children
        for point in self._points:
            _insert_into(children, point)
        self._points = []
        return children


def _insert_into(children: list[QuadTree], point: Point) -> bool:
    return any(child.insert(point) for child in children)
