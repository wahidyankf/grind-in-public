"""Tests for algo_03_queue_and_append_only_log.py."""

from algo_03_queue_and_append_only_log import WorkQueue, AppendOnlyLog


def test_work_queue_is_first_in_first_out() -> None:
    queue: WorkQueue[str] = WorkQueue()
    queue.enqueue("a")
    queue.enqueue("b")
    assert len(queue) == 2
    assert queue.dequeue() == "a"
    assert queue.dequeue() == "b"
    assert len(queue) == 0


def test_work_queue_empty_dequeue_raises() -> None:
    queue: WorkQueue[int] = WorkQueue()
    try:
        queue.dequeue()
    except IndexError:
        return
    raise AssertionError("expected IndexError")


def test_log_offsets_increase_from_zero() -> None:
    log: AppendOnlyLog[str] = AppendOnlyLog()
    assert log.append("x") == 0
    assert log.append("y") == 1
    assert log.read(0, 10) == ["x", "y"]


def test_consumer_groups_keep_independent_offsets() -> None:
    log: AppendOnlyLog[str] = AppendOnlyLog()
    for name in ["a", "b", "c"]:
        log.append(name)
    assert log.poll("billing", 2) == ["a", "b"]
    log.commit("billing", 2)
    assert log.poll("billing", 2) == ["c"]
    assert log.poll("search", 10) == ["a", "b", "c"]


def test_replay_by_moving_committed_offset_back() -> None:
    log: AppendOnlyLog[int] = AppendOnlyLog()
    for n in range(5):
        log.append(n)
    log.commit("reporting", 5)
    log.commit("reporting", 0)
    assert log.poll("reporting", 3) == [0, 1, 2]
