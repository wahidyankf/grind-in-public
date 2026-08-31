import type { MetadataRoute } from "next";

const siteUrl = "https://www.wahidyankf.com";

/**
 * Lists the three routes this site publishes, with the priorities that rank
 * them against one another. The list is written out rather than derived from
 * the route tree: three static pages gain nothing from a crawler of their own
 * directory, and deriving them would lose the ability to weight them by hand.
 */
export default function sitemap(): MetadataRoute.Sitemap {
  return [
    { url: siteUrl, changeFrequency: "monthly", priority: 1 },
    { url: `${siteUrl}/cv`, changeFrequency: "monthly", priority: 0.8 },
    {
      url: `${siteUrl}/personal-projects`,
      changeFrequency: "monthly",
      priority: 0.8,
    },
  ];
}
