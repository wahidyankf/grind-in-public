---
tldr: "Fix the responsible cause of a problem instead of suppressing its symptoms."
when_to_use: "Use when diagnosing errors, regressions, mistakes, or recurring failures."
---

# Root Cause Orientation

When an error, mistake, or failure is found, investigate its root cause and fix it at the responsible layer. Do not
ignore, suppress, or leave the cause in place merely to make the immediate symptom disappear.

Trace the failing flow end to end. Before changing a shared function, search for every caller and inspect sibling paths
that exercise the same responsibility. When they share the cause, repair the shared function once instead of adding a
guard to each caller or patching only the path named in the report.

If a durable fix requires owner direction or exceeds the task scope, report the evidence, impact, and recommended next
step.
