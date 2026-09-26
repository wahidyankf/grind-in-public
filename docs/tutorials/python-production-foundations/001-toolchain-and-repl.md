---
tldr: "Starts a disposable Python 3.14 project and uses the REPL for small, repeatable experiments."
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

The packaged application template already creates `src/balance_report/__init__.py`. Replace its generated greeting with
a module docstring so the package exposes no accidental public behaviour:

```python
"""The balance-report package."""
```

## Checkpoint

Run `uv run python -c "import balance_report; print(balance_report.__doc__)"`. It prints the package docstring, proving
that `uv` installed the `src/` package into the project environment. Do not commit the scratch project to this
repository. It is your own environment for typing and running the course.
