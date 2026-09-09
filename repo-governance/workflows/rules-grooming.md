---
tldr: "Sweeps the rule corpus for volume carrying no obligation, then hands every approved item to propagation."
when_to_use: "Use only when the owner explicitly directs a grooming sweep of the rule corpus."
---

# Rules Grooming

Run only on explicit owner direction. Grooming is a recurring sweep of the rule corpus for volume that carries no
obligation. It is never the remedy for one file exceeding the
[document word limit](../conventions/document-word-limit-policy.md); progressive disclosure handles that.

Grooming never writes. [Rules propagation](rules/rules-propagation.md) is the sole writer of every rule edit, and
finding an edit does not make grooming an exception. This workflow discovers, ranks, and hands off, which keeps conflict
scanning, placement, and enforcement disposition with the workflow that owns them. It never invokes the
[rules quality gate](rules-quality-gate.md).

## Admitted Classes

- **Cross-surface duplication** — one obligation stated on two or more surfaces with no recorded reason to keep both.
  Keep the canonical home by authority order, then the narrowest binding surface, and replace the rest with links. Every
  candidate carries a target-completeness check: the surviving home must already cover every case the removed text
  covered. A candidate failing it becomes "complete the target first", not a deletion.
- **Non-normative scaffolding** — prose stating no obligation: meta-narration, a preamble restating its own heading, a
  transition adding no condition. It deletes whole sentences and never rewrites one. Expect little.
- **Retirement** — a rule whose subject no longer exists, that a later rule supersedes in fact but not in text, or that
  no surface and no gate reaches. Absent inbound links alone are insufficient evidence. This is the only class that
  removes an obligation.

Nothing else is a candidate. A rule sharded under the [document naming policy](../conventions/document-naming-policy.md)
is judged by these same tests, never by shard count.

## Refused Permanently

- Rewriting prose to save words, or weakening any audience qualifier, scope boundary, exception, or pass condition.
- Trimming a safety guardrail: data safety, commit and push authorization, HIPPO and RHINO pin handling, or the
  drill-practice and integration-path rules.
- Removing or hollowing a document naming a convention, principle, standard, or workflow without per-item authorization
  naming it.
- Raising a word budget, in any class, for any reason.
- Deleting a rule to make room.

## Procedure

1. Record a pre-run obligation inventory under `local-tmp/rules-grooming/`: every distinct obligation with its audience,
   condition, and locations.
2. Discover candidates per admitted class, each carrying its class, affected paths, measured yield in lines, and
   evidence.
3. Rank by yield over risk — duplication and scaffolding low, retirement high — then group by subject. Drop any
   candidate whose yield is not worth one propagation delivery.
4. Present the ranked manifest for approval. Duplication and scaffolding approve as a batch, scaffolding only against
   the verbatim sentences enumerated in the checkpoint. Retirement approves per item with its own evidence; a batch
   approval of retirements is not accepted. An entry-point document needs authorization naming it, whatever class
   proposed it. No answer is not approval. Record every rejection and deferral with a reason, and keep deferrals in the
   manifest so the next run does not rediscover them as new.
5. Hand each approved subject group to propagation once, in ranked order, as a falsifiable rule change with its
   surfaces, canonical home, and evidence. Propagation's blockers bind: record one against its item and continue with
   the next group. Never restate an item more loosely to get it accepted, and never let a merge raise a budget.
6. Re-run the inventory and diff it against the pre-run snapshot. Index entries and routing clauses are navigation, not
   obligation, and are excluded from both sides. The run passes only when no obligation disappeared except approved
   retirements, none changed in audience, condition, qualifier, or exception, and every survivor stays reachable from a
   binding surface. For a verbatim move, confirm every moved line is present in its destination.
7. Record the run, its size delta, and each item's disposition.

## Terminal Contract

The only results are `NO_OP`, `GROOMED`, `PARTIAL`, and `HALTED`. An unanswered checkpoint ends the run `PARTIAL` with
nothing handed off. An unapproved obligation loss halts it: the reverting edit is itself a rule edit and goes back to
propagation, and the loss is reported whether or not that revert lands. Record which class produced it, because a class
that loses an obligation is a candidate for tightening. Passing authorizes neither commit nor push.
