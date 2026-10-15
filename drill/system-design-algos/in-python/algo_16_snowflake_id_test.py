"""Tests for algo_16_snowflake_id.py."""

import pytest

from algo_16_snowflake_id import EPOCH_MS, SnowflakeGenerator, decode_snowflake


class ScriptedClock:
    """Returns the scripted values in order, then keeps returning the last one."""

    def __init__(self, values: list[int]) -> None:
        self._values = values
        self._index = 0

    def __call__(self) -> int:
        value = self._values[min(self._index, len(self._values) - 1)]
        self._index += 1
        return value


T = EPOCH_MS + 1_000


def test_fields_round_trip_through_decode() -> None:
    generator = SnowflakeGenerator(worker_id=5, clock=ScriptedClock([T]))
    assert decode_snowflake(generator.next_id()) == (T, 5, 0)


def test_same_millisecond_uses_the_sequence() -> None:
    generator = SnowflakeGenerator(worker_id=1, clock=ScriptedClock([T, T, T]))
    sequences = [decode_snowflake(generator.next_id())[2] for _ in range(3)]
    assert sequences == [0, 1, 2]


def test_ids_increase_as_time_moves_forward() -> None:
    generator = SnowflakeGenerator(worker_id=1, clock=ScriptedClock([T, T + 1, T + 2]))
    first, second, third = generator.next_id(), generator.next_id(), generator.next_id()
    assert first < second < third


def test_sequence_overflow_waits_for_the_next_millisecond() -> None:
    # 4096 IDs fill millisecond T. The 4097th ID reads T again, waits, and then gets T + 1.
    generator = SnowflakeGenerator(worker_id=2, clock=ScriptedClock([T] * 4098 + [T + 1]))
    ids = [generator.next_id() for _ in range(4097)]
    assert len(set(ids)) == 4097
    assert decode_snowflake(ids[-1]) == (T + 1, 2, 0)


def test_clock_moving_backwards_raises() -> None:
    generator = SnowflakeGenerator(worker_id=1, clock=ScriptedClock([T + 5, T]))
    generator.next_id()
    with pytest.raises(RuntimeError):
        generator.next_id()


def test_worker_id_out_of_range_is_rejected() -> None:
    with pytest.raises(ValueError):
        SnowflakeGenerator(worker_id=1024)
