---
tldr: "Starts a disposable Python 3.13 project and uses the REPL for small, repeatable experiments."
when_to_use: "Use when beginning the course or checking a language behaviour before putting it in an application."
---

# Toolchain and REPL

Use a project environment rather than whichever `python` happens to be global on your machine. From the scratch project
described in the [course entry point](README.md), run a one-line program with `uv run python -c` and start a REPL with
`uv run python`. `uv run` creates and uses the project environment consistently.

At the REPL prompt, try a value and an f-string:

```python
name = "Ada"
greeting = f"Hello, {name}!"
print(greeting)
```

The result is `Hello, Ada!`. A variable names an object; it is not a typed box. Use small REPL experiments to confirm
what the language does, then move the behaviour into a function and a test when it matters to the application.

## Checkpoint

Create `src/balance_report/__init__.py` and run the following with `uv run python -m balance_report` after the CLI
lesson. For now, the package marker may be empty; creating it tells Python that this directory is a package.

```python
"""The balance-report package."""
```

Do not commit the scratch project to this repository. It is your own environment for typing and running the course.
