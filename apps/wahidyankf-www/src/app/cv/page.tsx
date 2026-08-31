import { CvContent } from "@/features/cv/shell/cv-content";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "CV | Wahidyan Kresna Fridayoka",
  description:
    "Full curriculum vitae of Wahidyan Kresna Fridayoka — work experience, skills, education, and certifications.",
};

/**
 * Renders the `/cv` route. Next.js binds the default export of this file to
 * that path, so the export name is what the framework reads and not something
 * a caller ever writes. The page holds no logic of its own: everything the CV
 * shows lives in `CvContent`, which can be rendered without a route around it.
 */
export default function CV() {
  return <CvContent />;
}
