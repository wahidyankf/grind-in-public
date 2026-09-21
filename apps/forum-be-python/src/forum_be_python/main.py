from collections.abc import AsyncGenerator
from contextlib import asynccontextmanager

from fastapi import FastAPI

from forum_be_python.config import database_url
from forum_be_python.db import create_pool, init_schema
from forum_be_python.deps import PoolDep
from forum_be_python.routes_auth import router as auth_router
from forum_be_python.routes_posts import router as posts_router
from forum_be_python.routes_threads import router as threads_router
from forum_be_python.routes_topics import router as topics_router


def create_app(url: str | None = None) -> FastAPI:
    @asynccontextmanager
    async def lifespan(app: FastAPI) -> AsyncGenerator[None]:
        pool = create_pool(url or database_url())
        init_schema(pool)
        app.state.pool = pool
        yield
        pool.close()

    app = FastAPI(title="Forum API", lifespan=lifespan)
    app.include_router(auth_router)
    app.include_router(topics_router)
    app.include_router(threads_router)
    app.include_router(posts_router)

    @app.get("/health")
    def health(pool: PoolDep) -> dict[str, str]:
        with pool.connection() as conn:
            conn.execute("SELECT 1")
        return {"status": "ok"}

    return app


app = create_app()
