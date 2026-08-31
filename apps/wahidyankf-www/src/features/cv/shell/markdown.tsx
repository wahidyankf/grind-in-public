import type React from "react";
import { HighlightText } from "@/features/ui/shell";

export const parseMarkdownLinks = (
  text: string,
  searchTerm: string,
): React.ReactNode => {
  const linkRegex = /\[([^\]]+)\]\(([^)]+)\)/g;
  const parts = text.split(linkRegex);

  return parts.map((part, index) => {
    if (index % 3 === 1) {
      // This is the link text
      const linkText = part;
      const linkUrl = parts[index + 1];
      return (
        <a
          // biome-ignore lint/suspicious/noArrayIndexKey: split() parts are positional, and the same link text may appear twice
          key={index}
          href={linkUrl}
          target="_blank"
          rel="noopener noreferrer"
          className="content-link" // Changed from about-me-link to content-link
        >
          <HighlightText text={linkText} searchTerm={searchTerm} />
        </a>
      );
    } else if (index % 3 === 0) {
      // This is regular text
      // biome-ignore lint/suspicious/noArrayIndexKey: split() parts are positional, and the same text may appear twice
      return <HighlightText key={index} text={part} searchTerm={searchTerm} />;
    }
    return null;
  });
};
