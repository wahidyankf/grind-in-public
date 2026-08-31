import type { MetadataRoute } from "next";

/**
 * Builds the crawler directives Next.js serves for this site. Every path is
 * allowed because this is a public personal site with no private routes, and
 * the sitemap is named by absolute URL because a crawler resolves that field
 * against the origin rather than against the file it read.
 */
export default function robots(): MetadataRoute.Robots {
  return {
    rules: [{ userAgent: "*", allow: "/" }],
    sitemap: "https://www.wahidyankf.com/sitemap.xml",
  };
}
