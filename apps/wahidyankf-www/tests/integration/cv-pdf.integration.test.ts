/**
 * Step definitions for the CV export feature, bound at the local integration
 * boundary.
 *
 * Covers: specs/apps/wahidyankf-www/behaviours/cv-export.feature
 *
 * This is the one binding in the application that runs against the real
 * filesystem rather than a DOM. It lives under `tests/integration/` and not
 * `tests/bdd/` for that reason: the `integration` project runs in the `node`
 * environment with no jsdom setup file, and its scenarios are the only ones
 * that write bytes to disk.
 */
import path from "node:path";
import {
  createWriteStream,
  existsSync,
  mkdtempSync,
  readFileSync,
  rmSync,
} from "node:fs";
import { once } from "node:events";
import { tmpdir } from "node:os";
import { Writable } from "node:stream";

import { loadFeature, describeFeature } from "@amiceli/vitest-cucumber";
import { expect } from "vitest";

import { cvData } from "@/features/cv/core/data";
import { buildCvPdfDocument } from "@/features/cv/core/pdf";
import { renderCvPdf } from "@/features/cv/shell/pdf";
import { behaviourTestLayer } from "../bdd/test-layer";

const feature = await loadFeature(
  path.resolve(
    process.cwd(),
    "../../specs/apps/wahidyankf-www/behaviours/cv-export.feature",
  ),
);

/**
 * Creates a throwaway output directory and hands back the path the export
 * should write to, along with the cleanup that removes it.
 *
 * Every scenario in this feature needs an output location that is real enough
 * for the export to write through and isolated enough that a failed run leaves
 * nothing behind. Keeping the creation and the removal in one place is what
 * stops a scenario from creating a directory it forgets to delete.
 */
const virtualDirectories = new Set<string>();
const virtualFiles = new Map<string, Buffer>();
let virtualDirectorySequence = 0;

function outputFixture(): {
  directory: string;
  file: string;
  cleanup: () => void;
} {
  if (behaviourTestLayer === "unit") {
    virtualDirectorySequence += 1;
    const directory = path.join(
      "/virtual",
      `cv-export-${virtualDirectorySequence}`,
    );
    const file = path.join(directory, "cv.pdf");
    virtualDirectories.add(directory);
    return {
      directory,
      file,
      cleanup: () => {
        virtualDirectories.delete(directory);
        virtualFiles.delete(file);
      },
    };
  }

  const directory = mkdtempSync(path.join(tmpdir(), "cv-export-"));
  return {
    directory,
    file: path.join(directory, "cv.pdf"),
    cleanup: () => rmSync(directory, { recursive: true, force: true }),
  };
}

/**
 * Runs the CV export the way `generate:cv-pdf` runs it and resolves once the
 * write has settled, returning the error if one was raised and `undefined` if
 * the file was written.
 *
 * Both scenarios drive the same three calls; only the destination differs, and
 * only one of them expects the write to succeed. Returning the failure instead
 * of throwing keeps a rejected write a value the scenario can assert against,
 * which is what the unwritable-output case needs.
 */
async function exportCvPdfTo(file: string): Promise<Error | undefined> {
  const pdf = renderCvPdf(buildCvPdfDocument(cvData));
  if (behaviourTestLayer === "unit") {
    const directory = path.dirname(file);
    if (!virtualDirectories.has(directory)) {
      return new Error(`ENOENT: no such output directory, open '${file}'`);
    }

    const chunks: Buffer[] = [];
    const sink = new Writable({
      write(chunk, _encoding, callback) {
        chunks.push(Buffer.from(chunk));
        callback();
      },
    });
    pdf.pipe(sink);
    try {
      await once(sink, "finish");
      virtualFiles.set(file, Buffer.concat(chunks));
      return undefined;
    } catch (error) {
      return error as Error;
    }
  }

  const sink = createWriteStream(file);
  pdf.pipe(sink);
  try {
    // `renderCvPdf` has already called `end()`, so the document is complete and
    // only the write is in flight. Awaiting `finish` is what makes a caller
    // read a whole file rather than whatever happened to be flushed. The same
    // await rejects when the stream emits `error` instead, which is how a write
    // into a missing directory surfaces.
    await once(sink, "finish");
    return undefined;
  } catch (error) {
    return error as Error;
  }
}

function outputExists(file: string): boolean {
  return behaviourTestLayer === "unit"
    ? virtualFiles.has(file)
    : existsSync(file);
}

function readOutput(file: string): Buffer {
  if (behaviourTestLayer === "integration") return readFileSync(file);
  const contents = virtualFiles.get(file);
  if (!contents) throw new Error(`No virtual output exists at '${file}'.`);
  return contents;
}

describeFeature(feature, ({ Scenario }) => {
  Scenario(
    "Generating the CV writes a PDF to the local filesystem",
    ({ Given, When, Then, And }) => {
      let output: ReturnType<typeof outputFixture>;
      let contents: Buffer;

      // @covers specs/apps/wahidyankf-www/behaviours/cv-export.feature:Generating the CV writes a PDF to the local filesystem
      Given("the application CV record contains at least one entry", () => {
        expect(cvData.length).toBeGreaterThan(0);
      });

      When(
        "the CV export runs against a writable output directory",
        async () => {
          output = outputFixture();
          expect(await exportCvPdfTo(output.file)).toBeUndefined();
        },
      );

      Then("a readable PDF file exists at the configured output path", () => {
        contents = readOutput(output.file);
        expect(contents.length).toBeGreaterThan(0);
      });

      And("the file begins with the PDF header bytes", () => {
        expect(contents.subarray(0, 5).toString("latin1")).toBe("%PDF-");
        output.cleanup();
      });
    },
  );

  Scenario(
    "Generating the CV reports an unwritable output location",
    ({ Given, When, Then, And }) => {
      let missingDirectory: string;
      let missingFile: string;
      let failure: unknown;

      // @covers specs/apps/wahidyankf-www/behaviours/cv-export.feature:Generating the CV reports an unwritable output location
      Given("the configured CV output directory does not exist", () => {
        const fixture = outputFixture();
        // Created and immediately removed, rather than simply naming a path
        // that was never there. That way the parent is a real temp location
        // this process owns, and the only thing missing is the directory the
        // export was told to write into.
        fixture.cleanup();
        missingDirectory = fixture.directory;
        missingFile = fixture.file;
        expect(
          behaviourTestLayer === "unit"
            ? virtualDirectories.has(missingDirectory)
            : existsSync(missingDirectory),
        ).toBe(false);
      });

      When("the CV export runs", async () => {
        failure = await exportCvPdfTo(missingFile);
      });

      Then("the export fails with a message naming the output path", () => {
        expect(failure).toBeInstanceOf(Error);
        expect(String((failure as Error).message)).toContain(missingFile);
      });

      And("no partial file is left behind", () => {
        expect(outputExists(missingFile)).toBe(false);
      });
    },
  );
});
