from edutech_be.main import app

# Every HTTP route the API serves. Add a route here when you add it to routes.py.
EXPECTED_ROUTES = {
    ("post", "/auth/register"),
    ("post", "/auth/login"),
    ("post", "/auth/refresh"),
    ("post", "/auth/logout"),
    ("get", "/auth/me"),
    ("get", "/health"),
}


def test_app_serves_exactly_the_expected_routes() -> None:
    paths = app.openapi()["paths"]

    served = {
        (method, path) for path, operations in paths.items() for method in operations
    }

    assert served == EXPECTED_ROUTES
