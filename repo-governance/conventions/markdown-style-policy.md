---
tldr: "Defines a repository-wide source style for Markdown prose and diagrams."
when_to_use: "Use when creating, editing, reviewing, or formatting any Markdown file."
---

# Markdown Style Policy

## Scope

This policy applies to every committed `*.md` file in this repository, including human documentation, agent
instructions, governance, project READMEs, CV material, and scripts documentation.

## Paragraphs

Hard-wrap prose at 120 columns, and separate each paragraph from the next by one blank line. Prettier applies the wrap,
so it is not maintained by hand and not a matter of judgment at the keyboard.

This repository is read in a terminal. Soft wrap solves overflow but not measure: it breaks at whatever the window
happens to be, so a wide pane yields a 180-column line that is hard to track back to the left margin. A fixed 120-column
wrap gives the same measure on every screen. The cost is that editing a word reflows the rest of its paragraph, which
makes a one-word change a several-line diff; that is accepted in exchange for prose that reads the same everywhere.

One kind of document is exempt: a copy-paste target, whose text a downstream consumer receives verbatim. A rewrap
inserts newlines that travel with the pasted text and arrive as ragged breaks in the destination, so there the line
break is content rather than presentation. Exempt such a file with a Prettier `overrides` entry setting
`proseWrap: "preserve"`, never by adding it to `.prettierignore`: the override hands back line-break control alone,
while an ignore also drops the file's tables, lists, and `format:check` coverage. `apps/wahidyankf-www/docs/` holds the
current examples, its CV and LinkedIn drafts; its own `README.md` is an index like any other and stays wrapped.

Use structural Markdown — headings, lists, tables, blockquotes, and fenced code blocks — when content is structurally
distinct, rather than splitting a paragraph for appearance.

## Line Length

No line exceeds 120 characters: prose, headings, list items, table rows, and fenced code all count, so a document reads
in a terminal editor with neither a renderer nor horizontal scrolling. Prettier breaks prose; fix what it cannot break
by hand:

- shorten the cells of a table whose aligned width exceeds 120, split it into narrower tables, or turn it into a list;
- move a long link target into a reference-style definition, which the limit exempts;
- continue a long shell command on the next line with `\`, first moving it from inline code into a fenced block; and
- wrap a long line of example output only after a sentence saying the real output is one line.

Exempt are the copy-paste targets above, completed plans under `plans/done/`, fixture bytes under `specs/fixtures/`, and
YAML front matter, whose values RHINO reads one per line.

## Diagrams and Schemas

For diagrams, schemas, flows, and similar visual models in Markdown, prefer ASCII art in a fenced `text` block. Use only
ASCII characters such as `+`, `-`, `|`, and `>` with clear labels so the model remains legible in a terminal, NVIM, code
review, and a rendered Markdown view. Prefer a Markdown table when a tabular schema communicates the relationship more
directly.

Do not use Mermaid: it is unreadable without a renderer, and the `mermaid` gate declares the plain-text authoring rule,
so it rejects any Mermaid block.

## Related

Filenames are governed separately by the [document naming policy](document-naming-policy.md).

## Enforcement

Prettier owns this style. `.prettierrc.json` sets `proseWrap: "always"` with a `printWidth` of 120, and pins
non-Markdown files back to 80 through an `overrides` entry so the code formatting width is unaffected. Prettier pads
every table into aligned columns, so a table fits only when its aligned width does. The `markdown-line-length` gate runs
`markdownlint-cli2` with rule `MD013` alone, configured with its exemptions in `.markdownlint-cli2.jsonc`, on pre-push
and in `npm run check:hygiene`.

Run `rtk npm run format` after changing Markdown; use `rtk npm run format:check` to verify it before committing.
Pre-commit formats staged files, so the wrap is applied whether or not it was run by hand.
