from datetime import datetime
from typing import Literal

from pydantic import BaseModel, Field


class Credentials(BaseModel):
    username: str = Field(pattern=r"^[a-z0-9][a-z0-9_-]{2,31}$")
    password: str = Field(min_length=8, max_length=128)


class UserOut(BaseModel):
    id: int
    username: str


class LoginRequest(BaseModel):
    # No format rules here: a malformed username is just a failed login, not a different kind of error.
    username: str
    password: str


class TokenOut(BaseModel):
    access_token: str
    token_type: Literal["bearer"] = "bearer"
    expires_at: datetime


class TopicIn(BaseModel):
    title: str = Field(min_length=1, max_length=120)
    description: str = Field(default="", max_length=500)


class TopicPatch(BaseModel):
    # A field left out of the request stays as it is, so absence and an empty value are different things.
    title: str | None = Field(default=None, min_length=1, max_length=120)
    description: str | None = Field(default=None, max_length=500)


class TopicOut(BaseModel):
    id: int
    title: str
    description: str
    created_by: str
    created_at: datetime


class ThreadIn(BaseModel):
    title: str = Field(min_length=1, max_length=200)
    body: str = Field(min_length=1, max_length=10_000)


class ThreadPatch(BaseModel):
    title: str = Field(min_length=1, max_length=200)


class ThreadOut(BaseModel):
    id: int
    topic_id: int
    title: str
    created_by: str
    created_at: datetime


class PostIn(BaseModel):
    body: str = Field(min_length=1, max_length=10_000)
    parent_id: int | None = None


class PostPatch(BaseModel):
    body: str = Field(min_length=1, max_length=10_000)


class PostOut(BaseModel):
    id: int
    thread_id: int
    parent_id: int | None
    author: str
    body: str
    created_at: datetime
    updated_at: datetime
    # Pydantic copies a mutable default for every instance, so sharing one list is not a risk here.
    replies: list["PostOut"] = []


class ThreadDetail(ThreadOut):
    posts: list[PostOut]
