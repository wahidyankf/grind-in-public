// The presentation primitives the app shell renders. These four arrived from a
// shared design-system package during the migration and were inlined, because
// this repository publishes no libraries; this barrel is what keeps their call
// sites importing one path rather than four.
//
// `ScrollToTop` and `ThemeToggle` are default exports in their own files and
// are renamed here, so every consumer imports a named symbol regardless of how
// the component happens to declare itself.
export { HighlightText, highlightText } from "./highlight-text";
export { SearchComponent } from "./search-component";
export { default as ScrollToTop } from "./scroll-to-top";
export { default as ThemeToggle } from "./theme-toggle";
