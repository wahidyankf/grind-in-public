import { createLucideIcon } from "lucide-react";

// Lucide deprecated every brand icon in its set and has scheduled them for
// removal in v1.0 (lucide-icons/lucide#670), because the project no longer
// wants to carry third-party marks. Importing `Github` or `Linkedin` from
// `lucide-react` therefore raises a deprecation hint today and breaks the
// build on the next major.
//
// The glyphs themselves are still the right ones for this site, so rather than
// swapping in a second icon dependency, this module rebuilds them from the
// same path data through `createLucideIcon` — the package's supported public
// factory. The result renders identically and accepts the same props as any
// other Lucide icon, so call sites keep passing `className`, `size`, and the
// rest unchanged, and the upgrade to v1.0 leaves them untouched.
//
// The path data is Lucide's own, under its ISC license.

const GithubIcon = createLucideIcon("github", [
  [
    "path",
    {
      d: "M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4",
      key: "tonef",
    },
  ],
  ["path", { d: "M9 18c-4.51 2-5-2-7-2", key: "9comsn" }],
]);

const LinkedinIcon = createLucideIcon("linkedin", [
  [
    "path",
    {
      d: "M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z",
      key: "c2jq9f",
    },
  ],
  ["rect", { width: "4", height: "12", x: "2", y: "9", key: "mk3on5" }],
  ["circle", { cx: "4", cy: "4", r: "2", key: "bt5ra8" }],
]);

export { GithubIcon, LinkedinIcon };
