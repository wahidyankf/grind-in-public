import "./globals.css";
import { Inter } from "next/font/google";
import type { Metadata } from "next";
import { ScrollToTop, ThemeToggle } from "@/features/ui/shell";
import { GoogleAnalytics, GoogleTagManager } from "@next/third-parties/google";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title:
    "Wahidyan Kresna Fridayoka | Engineering Leader — Digital Banking, Fintech & RegTech",
  description:
    "Portfolio and CV of Wahidyan Kresna Fridayoka, an engineering leader with nearly nine years of experience in Sharia-compliant digital banking, fintech, and RegTech, leading cross-functional engineering teams and open-source work.",
  icons: [
    { rel: "icon", url: "/favicon.ico" },
    { rel: "apple-touch-icon", url: "/favicon.ico" },
  ],
  keywords: [
    "Engineering Leadership",
    "Engineering Manager",
    "Digital Banking",
    "Sharia-Compliant Fintech",
    "RegTech",
    "Software System Designer",
    "Financing Engineer",
    "Software Testing",
    "Core Banking",
    "ISO 27001",
  ],
  authors: [{ name: "Wahidyan Kresna Fridayoka" }],
  creator: "Wahidyan Kresna Fridayoka",
  openGraph: {
    type: "website",
    locale: "en_US",
    url: "https://www.wahidyankf.com",
    siteName: "Wahidyan Kresna Fridayoka",
    title:
      "Wahidyan Kresna Fridayoka | Engineering Leader — Digital Banking, Fintech & RegTech",
    description:
      "Portfolio and CV of Wahidyan Kresna Fridayoka, an engineering leader with nearly nine years of experience in Sharia-compliant digital banking, fintech, and RegTech, leading cross-functional engineering teams and open-source work.",
    images: [
      {
        url: "https://www.wahidyankf.com/og-image.jpg",
        width: 1200,
        height: 630,
      },
    ],
  },
  twitter: {
    card: "summary_large_image",
    title:
      "Wahidyan Kresna Fridayoka | Engineering Leader — Digital Banking, Fintech & RegTech",
    description:
      "Portfolio and CV of Wahidyan Kresna Fridayoka, an engineering leader with nearly nine years of experience in Sharia-compliant digital banking, fintech, and RegTech, leading cross-functional engineering teams and open-source work.",
    images: ["https://www.wahidyankf.com/og-image.jpg"],
  },
  robots: {
    index: true,
    follow: true,
    googleBot: {
      index: true,
      follow: true,
      "max-video-preview": -1,
      "max-image-preview": "large",
      "max-snippet": -1,
    },
  },
};

/**
 * Wraps every page in the document shell: the language and font on the root
 * element, the pre-paint theme script in the head, the theme toggle and
 * scroll-to-top controls around the page body, and the analytics tags last.
 * Next.js renders this around all other routes, so anything that must exist on
 * every page belongs here rather than being repeated in each one.
 */
export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className={inter.className}>
      <head>
        {/* This script must run before the first paint, or the page renders in
            the wrong theme and then corrects itself visibly. It reads one
            localStorage key and sets one class; no user input reaches it, which
            is why injecting it as markup is safe here. */}
        <script
          // biome-ignore lint/security/noDangerouslySetInnerHtml: a blocking pre-paint theme script with no user input
          dangerouslySetInnerHTML={{
            __html: `
              (function() {
                var theme = localStorage.getItem('theme');
                if (theme === 'light') {
                  document.documentElement.classList.add('light-theme');
                }
              })();
            `,
          }}
        />
      </head>
      <body className="root-layout flex flex-col lg:flex-row">
        <div className="body-content flex-grow">
          <div className="fixed top-4 right-4 z-50">
            <ThemeToggle />
          </div>
          {children}
          <ScrollToTop />
        </div>
      </body>
      <GoogleAnalytics gaId="G-0F62KPYGM3" />
      <GoogleTagManager gtmId="GTM-THSDKWSZ" />
    </html>
  );
}
