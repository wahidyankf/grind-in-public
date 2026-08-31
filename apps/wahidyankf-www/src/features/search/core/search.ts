export type SearchTerm = string;
export type SearchResult<T> = T[];

/**
 * Filters items by a case-insensitive substring match against the named keys.
 * A key holding an array matches when any of its string elements matches,
 * which is what lets one search box cover both a project title and its list of
 * skills. An empty term returns the input untouched rather than nothing, so
 * clearing the box restores the full list instead of emptying it.
 */
export function filterItems<T>(
  items: T[],
  searchTerm: string,
  keys: (keyof T)[],
): T[] {
  if (!searchTerm) return items;

  const lowercasedTerm = searchTerm.toLowerCase();

  return items.filter((item) =>
    keys.some((key) => {
      const value = item[key];
      if (typeof value === "string") {
        return value.toLowerCase().includes(lowercasedTerm);
      }
      if (Array.isArray(value)) {
        return value.some((v) =>
          typeof v === "string"
            ? v.toLowerCase().includes(lowercasedTerm)
            : false,
        );
      }
      return false;
    }),
  );
}
