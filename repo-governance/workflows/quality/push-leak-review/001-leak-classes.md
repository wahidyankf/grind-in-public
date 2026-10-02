---
tldr: "Defines the three leak classes, what is never a leak, and why every outbound commit is in scope."
when_to_use: "Use when deciding whether a candidate value in a commit, file name, or message is a leak."
---

# Leak Classes

A leak review judges three classes and no others. The
[commit hook policy](../../../development/commit-hook-policy.md#public-repository-safety) owns what may never be
committed; these classes are how a review names what it found.

| Class                            | A finding is                                                                   |
| -------------------------------- | ------------------------------------------------------------------------------ |
| `secret_or_private_value`        | a real credential or other value that grants access, in any environment        |
| `protected_environment_property` | a value that belongs in environment or secret storage, not in a tracked file   |
| `machine_specific_absolute_path` | anything identifying the machine it came from, an absolute home path above all |

- `secret_or_private_value` names staging and production credentials explicitly, because they reach real data and real
  users.
- `protected_environment_property` covers a connection string, or a non-public environment's endpoint or account
  identifier.
- `machine_specific_absolute_path` covers `/Users/<name>/`, `/home/<name>/`, and `C:\Users\<name>\`, a username, a
  hostname, and a private address.

## Real Means Real

Any real credential is a finding, whichever environment it belongs to: a development key that works is still a key, and
deciding which environment a token reaches is guesswork the review does not attempt. Only a value that is unmistakably
synthetic passes.

## What Is Not a Leak

- A home-relative path such as `~/notes`. It names no account and resolves on every machine.
- A repository-relative path, or a documented placeholder such as `<name>` or `<repository-path>`.
- Public identifiers, documented public values, and loopback addresses in test configuration that targets a local
  service.
- Synthetic fixtures that are unmistakably synthetic.

A name containing `key`, `token`, `secret`, or `prod` is not evidence alone. A candidate is a finding only when shape
and context show the value is real.

## History Is the Subject

Every commit bound for the remote is outbound on its own. A value added by one commit and deleted by the next is
published with both, and every clone keeps it. A review therefore reads each commit's additions, its file names, and its
message, never only the range's final files.

The review binds from adoption onward. Content a range does not add is not judged again, and history published before
adoption is out of scope for the review; a leak found there is handled under the commit hook policy instead.

## The Review Stays Private

Nothing about the review is posted, and its notes are outbound the moment they reach a commit, a log, or a report. A
path pasted into them is the finding the review exists to catch, and a quoted secret is disclosed again in the record of
its discovery. Anything found is treated as already disclosed.

It is not a security or semantic review. A screen matches shapes; this review reads context, and three classes keep it
small enough for every push.
