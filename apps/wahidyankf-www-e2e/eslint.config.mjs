import tsParser from "@typescript-eslint/parser";
import jsdoc from "eslint-plugin-jsdoc";

// Browser bindings carry the same commentary contract as production TypeScript
// without duplicating the style and correctness checks already owned by Biome.
export default [
  {
    files: ["tests/**/*.ts", "tools/**/*.ts"],
    languageOptions: { parser: tsParser },
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
