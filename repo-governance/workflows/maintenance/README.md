---
tldr: "Indexes the upkeep workflows: rules grooming and development artifact clean-up."
when_to_use: "Use when reducing the rule corpus or removing what one piece of work left behind."
---

# Maintenance Workflows

These workflows keep the repository in order between deliveries.

## Directory Map

- [Rules Grooming](rules-grooming.md) runs only on explicit owner direction, never writes, and hands each approved
  reduction to [Rules Propagation](../quality/rules-propagation.md).
- [Dev Artifact Clean-Up](dev-artifact-clean-up.md) removes exactly the development artifacts one piece of work created
  — its regenerable build output — and leaves local `main` level with `origin/main`.
