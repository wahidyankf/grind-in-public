# Grind in Public

Personal lifelong-learning workspace maintained by **Wahidyan Kresna Fridayoka**.

[LinkedIn](https://www.linkedin.com/in/wahidyan-kresna-fridayoka/)

This repository is where I practice, review, and document the craft of software engineering, in the open, for as long as
I keep building.

This README and [`docs/`](docs/README.md) are for people.

## Hands-On by Design

Unlike my `ose-*` repositories, this repository is intentionally not fully automation-first. I will complete parts of
the work by hand—especially coding exercises and drills—to build muscle memory, reinforce the fundamentals, and practice
recalling and explaining solutions without help.

Tools can support review and feedback, but they should not replace the hands-on practice at the heart of this
repository.

## Focus Areas

- Data structures and algorithms
- Coding exercises and problem-solving patterns
- System design
- Computer science and software engineering fundamentals
- Languages, tools, and engineering craft
- Retrospective notes on what stuck and what did not

## Practice Workflow

1. Choose a topic or problem worth understanding.
2. Work through it by hand under deliberate constraints, without shortcuts.
3. Record the solution and reasoning.
4. Review trade-offs, mistakes, and possible improvements.
5. Revisit the exercise until the approach is clear and repeatable.

## Goal

Build strong fundamentals, explain solutions clearly, and keep compounding skill as a software engineer over years
rather than cramming for a deadline.

## Nx Workspace

The repository uses Nx as a raw task runner for its npm workspaces:

- `apps/` holds runnable applications.
- `libs/` holds reusable packages consumed by applications.

Build all projects with `npm run build` and run quick checks with `npm test`. This workspace deliberately avoids
technology-specific Nx plugins and uses Nx only as a raw task runner.

Repository checks, including Markdown links and line length, are run through
[RHINO](https://github.com/wahidyankf/rhino), a Rust CLI consumed through a checksum-pinned wrapper and configured
entirely by `repo-config.yml`. Run `npm run test:repo` for the deterministic repository checks.

One application lives here. [wahidyankf-www](apps/wahidyankf-www/README.md) is the Next.js personal site, and it holds
the repository's single authoritative CV record — the practice workspace and its public face maintained under one set of
rules. It is held to the repository's gates: a Gherkin corpus under [`specs/`](specs/README.md), and a 99% line-coverage
floor.

## Specs

What the software should do is described as Gherkin in [`specs/`](specs/README.md), alongside each non-drill
application's current as-built architecture and separate from the code that implements them.

## Documentation

Human-facing project documentation is organized with the [Diátaxis framework](docs/README.md).
