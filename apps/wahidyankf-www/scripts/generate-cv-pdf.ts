import fs from "node:fs";
import path from "node:path";
import { cvData } from "../src/features/cv/core/data";
import { buildCvPdfDocument } from "../src/features/cv/core/pdf";
import { renderCvPdf } from "../src/features/cv/shell/pdf";

const OUTPUT_PATH = path.resolve(
  __dirname,
  "../public/wahidyankf-kresna-fridayoka-cv.pdf",
);

const document = buildCvPdfDocument(cvData);
const pdf = renderCvPdf(document);
pdf.pipe(fs.createWriteStream(OUTPUT_PATH));

// This is a build script invoked by `generate:cv-pdf`. Stdout is how it tells
// the operator which file it wrote, so silencing it would make a successful
// run indistinguishable from one that produced nothing.
// biome-ignore lint/suspicious/noConsole: a build script reports its output on stdout
console.log(`Generated CV PDF at ${OUTPUT_PATH}`);
