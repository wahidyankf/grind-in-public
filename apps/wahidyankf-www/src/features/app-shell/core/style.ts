import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

/**
 * Joins class names and then resolves the Tailwind conflicts between them.
 * Order matters here: `clsx` first flattens the conditional and array forms
 * into one string, and `twMerge` then drops the earlier of any two classes
 * that set the same Tailwind property, so the last one written wins.
 */
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}
