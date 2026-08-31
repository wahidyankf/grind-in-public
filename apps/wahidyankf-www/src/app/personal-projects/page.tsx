import { PersonalProjectsContent } from "@/features/personal-projects/shell/personal-projects-content";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Independent Projects | Wahidyan Kresna Fridayoka",
  description:
    "Open-source and independent projects by Wahidyan Kresna Fridayoka, including OSE, AyoKoding, OrganicLever, and more.",
};

/**
 * Renders the `/personal-projects` route. As with the other route files here,
 * the page is a binding point for Next.js and the content it shows lives in
 * `PersonalProjectsContent`.
 */
export default function Projects() {
  return <PersonalProjectsContent />;
}
