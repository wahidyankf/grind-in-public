import tsParser from "@typescript-eslint/parser";
import jsdoc from "eslint-plugin-jsdoc";

// This configuration owns commentary and nothing else. Biome owns style and
// correctness for this project, so a rule that both could report is a finding
// delivered twice and is deliberately absent here. That split is what the code
// commentary policy asks for, and it is the TypeScript counterpart of what
// Revive does for badakmini-cli under golangci-lint.
export default [
  {
    files: ["src/**/*.ts", "src/**/*.tsx"],
    // ESLint's own parser reads JavaScript only, so it stops at the first type
    // annotation or JSX tag and reports a parse error instead of a commentary
    // finding. The TypeScript parser is loaded for its reader alone: no
    // `@typescript-eslint` rule is enabled here, because every rule that plugin
    // offers is one Biome already reports for this project.
    languageOptions: {
      parser: tsParser,
      parserOptions: {
        ecmaFeatures: { jsx: true },
      },
    },
    plugins: { jsdoc },
    rules: {
      // A named executable declaration is one a future reader will look up by
      // name, so it owes them a sentence saying what it is for. Anonymous and
      // inline functions are excluded: naming a callback's purpose belongs in
      // the call site that reads it, not in a comment above it.
      "jsdoc/require-jsdoc": [
        "error",
        {
          require: {
            FunctionDeclaration: true,
            ClassDeclaration: true,
            MethodDefinition: true,
            ArrowFunctionExpression: false,
            FunctionExpression: false,
          },
        },
      ],
      // A summary that is not a sentence is a label, and a label repeats the
      // name rather than explaining it.
      "jsdoc/require-description": ["error", { descriptionStyle: "body" }],
      // The rule ends a sentence at every period, so an abbreviation reads as a
      // full stop and whatever follows it reads as a sentence starting in
      // lowercase. Listing the abbreviations this codebase actually writes
      // keeps the rule reporting real findings instead of punctuation it
      // misread.
      "jsdoc/require-description-complete-sentence": [
        "error",
        { abbreviations: ["e.g", "i.e", "etc", "vs"] },
      ],
    },
  },
  {
    // The moved Playwright step files. They carry the same three commentary
    // rules as `src`, because a step definition is read by the same people for
    // the same reason. The JSX parser feature is deliberately absent: a step
    // file drives a browser and contains no JSX, so enabling it would be
    // configuration the code cannot exercise.
    files: ["tests/e2e/steps/**/*.ts"],
    languageOptions: {
      parser: tsParser,
    },
    plugins: { jsdoc },
    rules: {
      "jsdoc/require-jsdoc": [
        "error",
        {
          require: {
            FunctionDeclaration: true,
            ClassDeclaration: true,
            MethodDefinition: true,
            ArrowFunctionExpression: false,
            FunctionExpression: false,
          },
        },
      ],
      "jsdoc/require-description": ["error", { descriptionStyle: "body" }],
      "jsdoc/require-description-complete-sentence": [
        "error",
        { abbreviations: ["e.g", "i.e", "etc", "vs"] },
      ],
    },
  },
];
