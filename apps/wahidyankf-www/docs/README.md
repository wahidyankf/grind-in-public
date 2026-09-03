# Career Evidence

Wahidyan Kresna Fridayoka's career evidence and professional-profile drafts. These documents are the source for CV and
LinkedIn work; read the relevant file before making a claim, an edit, or an export.

They live inside this application because the application holds the repository's single authoritative CV record. That
record is `src/features/cv/core/data.ts`, not a document here: the site renders it, and
`npx nx run wahidyankf-www:generate:cv-pdf` exports it to `public/wahidyankf-kresna-fridayoka-cv.pdf`. Nothing in this
directory is imported by a route, so editing one changes no rendered page.

## Working Rules

Treat `cv-raw.md` as the factual source of truth. Every public-facing claim — in the CV record, on the site, in the
exported PDF, and in LinkedIn copy — stays accurate, contextualized, and consistent with it.

Do not publish or promote material that `cv-raw.md` marks as unsuitable for public use without explicit owner direction.
Where it cites a source that is not in this directory, do not invent that source's contents; add the source, or verify
the affected claims with the owner before relying on them.

Changing a fact means changing it in two places, and in this order: record the evidence in `cv-raw.md`, then update
`src/features/cv/core/data.ts`. Regenerate the PDF afterwards and inspect it before sharing.

These three drafts are exempt from the repository's 120-column prose wrap, because their text is pasted into LinkedIn
and a rewrap would carry its newlines along. `.prettierrc.json` holds that exemption; do not reflow them by hand. This
`README.md` is not exempt.

## Directory Map

- [cv-raw.md](cv-raw.md) — the evidence base, holding dates, metrics, source wording, context, and notes about material
  not suitable for public use.
- [cv-linkedin.md](cv-linkedin.md) — the long-form, copy-ready LinkedIn profile.
- [linkedin-projects.md](linkedin-projects.md) — draft LinkedIn project entries and publishing notes.
