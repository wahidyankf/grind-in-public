from collections.abc import AsyncGenerator
from contextlib import asynccontextmanager

from fastapi import FastAPI

from edutech_be.config import settings
from edutech_be.db import create_db_engine
from edutech_be.routes import router


@asynccontextmanager
async def lifespan(app: FastAPI) -> AsyncGenerator[None]:
    app.state.engine = create_db_engine(settings)
    yield
    app.state.engine.dispose()


# Routes are defined in routes.py. This file only builds the app.
app = FastAPI(title="edutech-be", lifespan=lifespan)
app.include_router(router)
