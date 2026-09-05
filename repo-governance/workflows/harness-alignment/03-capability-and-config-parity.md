---
tldr:
  "Compares canonical capabilities with native adapters, then checks project configs and indexes and states what the
  deterministic contract cannot prove."
when_to_use: "Use after checking instruction routing, to verify capability adapters, harness configs, and READMEs."
---

# Capability and Config Parity

## Shared Capabilities

Edit the canonical skill bundle or custom-agent definition first, then reconcile every native adapter as the
[harness capability parity policy](../../conventions/harness-capability-parity-policy.md) requires. Confirm exact
descriptions and routes, complete supporting resources, one adapter per required harness, and the documented native
permission mapping. Remove stale copies rather than carrying two prompt bodies.

## Configs and Indexes

Confirm each project config holds only documented settings, and that each directory is indexed as the
[documentation index policy](../../documentation-index-policy.md) requires, exemptions included.

## Verification Boundary

The deterministic check proves canonical content and configured adapter semantics. It cannot prove that a vendor model
obeys a prompt, a user-global plugin preserves it, or an unavailable native restriction exists. Record those runtime
limits without weakening the repository contract. The read-only [rules quality gate](../rules-quality-gate.md) remains
owner-gated and is not part of routine alignment.
