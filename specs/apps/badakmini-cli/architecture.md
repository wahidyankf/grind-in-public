# Badak Mini Architecture

This is the canonical as-built C4 model for Badak Mini. Maintain it under the repository
[architecture specification policy](../../../repo-governance/development/architecture-specifications.md).

## System Context

```text
+-------------------------+       command and result       +------------------+
| Person                  | ----------------------------> | Software system  |
| Repository contributor  |                               | Badak Mini CLI   |
+-------------------------+                               +--------+---------+
                                                                    |
+-------------------------+       child process / status            | read-only filesystem
| External system         | ----------------------------------------+ and Git inspection
| Nx targets and Git hooks|                                        |
+-------------------------+                                        v
                                                         +------------------+
                                                         | External data    |
                                                         | Repository tree  |
                                                         +------------------+
```

Badak Mini is a local, network-free governance tool. Contributors invoke it directly, while Nx targets and Git hooks run
the same public command. It reads the repository tree and Git index but does not modify either.

## Container View

```text
+-------------------------+       arguments, text, exit status      +---------------------------+
| Contributor or          | --------------------------------------> | Container                 |
| Nx target / Git hook    |                                         | Badak Mini CLI process    |
+-------------------------+                                         | Go and Cobra              |
                                                                      +------------+--------------+
                                                                                   |
                                                                                   | filesystem APIs and Git subprocess
                                                                                   v
                                                                      +---------------------------+
                                                                      | External data / process   |
                                                                      | repository tree and Git    |
                                                                      +---------------------------+
```

The compiled Go process is the only Badak Mini runtime container. The operating system and Git process form its material
local boundary; there is no network service, database, or application-owned persistent store.

## Component View

```text
+--------------------+      arguments      +-------------------------+
| Caller             | -----------------> | Process entry            |
| contributor / hook |                    | cmd/badak-mini           |
+--------------------+                    +------------+------------+
                                                     |
                                                     | host adapters
                                                     v
                                        +-------------------------+
                                        | Command orchestration   |
                                        | internal/cli, Cobra     |
                                        +------------+------------+
                                                     |
                 +-----------------------------------+-----------------------------------+
                 |                                   |                                   |
                 v                                   v                                   v
    +------------------------+       +------------------------+       +------------------------+
    | Governance validators  |       | Markdown-link validator |       | Parity and rule-change |
    | instruction size       |       | tracked local links     |       | harness rule surfaces  |
    +------------------------+       +------------------------+       +------------------------+
                 \                                   |                                   /
                  \                                  |                                  /
                   +--------------------------------+---------------------------------+
                                                    |
                                                    | injected filesystem, Git, and stream boundary
                                                    v
                                       +-------------------------+
                                       | Repository tree and Git |
                                       +-------------------------+
```

`cmd/badak-mini` owns production filesystem, Git-process, and stream adapters. `internal/cli` owns Cobra command
parsing, dispatch, output, and exit codes; its injected runtime lets unit tests replace the host boundary. The
validation packages own their focused inspections.

## Architectural Constraints

- The CLI is local and network-free, and its validators do not mutate inspected repositories.
- Unit tests replace filesystem, Git, and stream boundaries; integration tests use isolated local Git and filesystem
  fixtures; process E2E tests observe only the built command's public contract.
- Text output is the human interface. Exit code `0` means success, `1` means a validation or operational failure, and
  `2` means invalid invocation.
- Governance documents remain authoritative. Badak Mini reports structural problems and announces required workflows,
  but it does not decide whether a workflow has been followed.

## Behavior Traceability

Executable behavior is specified in [behavior/](behavior/). The unit, local-integration, and process E2E adapters all
consume that same recursive corpus.
