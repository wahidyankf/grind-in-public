import { HomeContent } from "@/features/home/shell/home-content";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title:
    "Wahidyan Kresna Fridayoka | Engineering Leader — Digital Banking, Fintech & RegTech",
};

/**
 * Renders the site root. The page delegates to `HomeContent` so the rendering
 * logic stays in a component that can be exercised without a route around it.
 */
export default function Home() {
  return <HomeContent />;
}
