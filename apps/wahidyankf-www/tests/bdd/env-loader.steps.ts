import path from "node:path";
import { mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { loadFeature, describeFeature } from "@amiceli/vitest-cucumber";
import { afterAll, expect } from "vitest";
import { loadTierEnv } from "@/features/env/core/tier-env";
import { behaviourTestLayer } from "./test-layer";

// Unit execution drives loadTierEnv() through an injected in-memory filesystem. Integration uses
// throwaway temp directories. Both use an isolated env record and never mutate process.env.
//
// @amiceli/vitest-cucumber schedules every Given/When/Then/And of a Scenario as its own vitest
// `test()` (see node_modules/@amiceli/vitest-cucumber's describe-scenario.js `test.for(...)`), so
// an `afterEach` here would tear down a scenario's tmp directory between its own Given and When
// steps. Cleanup is deferred to a single `afterAll` for the whole file instead.
const feature = await loadFeature(
  path.resolve(
    process.cwd(),
    "../../specs/apps/wahidyankf-www/behaviours/env-loader.feature",
  ),
);

type EnvRecord = Record<string, string | undefined>;

const tmpDirs: string[] = [];
const virtualFiles = new Map<string, string>();
let virtualDirectorySequence = 0;

function makeTmpAppDir(): string {
  if (behaviourTestLayer === "unit") {
    virtualDirectorySequence += 1;
    return path.join("/virtual", `env-loader-${virtualDirectorySequence}`);
  }
  const dir = mkdtempSync(
    path.join(tmpdir(), "wahidyankf-www-env-loader-test-"),
  );
  tmpDirs.push(dir);
  return dir;
}

function writeEnvFile(
  appDir: string,
  fileName: string,
  contents: string,
): void {
  if (behaviourTestLayer === "unit") {
    virtualFiles.set(path.join(appDir, fileName), contents);
    return;
  }
  writeFileSync(path.join(appDir, fileName), contents, "utf-8");
}

function loadAtCurrentLayer(appDir: string, env: EnvRecord): void {
  if (behaviourTestLayer === "integration") {
    loadTierEnv({ appDir, env });
    return;
  }

  loadTierEnv({
    appDir,
    env,
    fileExists: (filePath) => virtualFiles.has(filePath),
    loadFile: (filePath, targetEnv) => {
      for (const line of (virtualFiles.get(filePath) ?? "").split("\n")) {
        const separator = line.indexOf("=");
        if (separator < 1) continue;
        const key = line.slice(0, separator);
        if (targetEnv[key] === undefined) {
          targetEnv[key] = line.slice(separator + 1);
        }
      }
    },
  });
}

afterAll(() => {
  for (const dir of tmpDirs.splice(0)) {
    rmSync(dir, { recursive: true, force: true });
  }
  virtualFiles.clear();
});

describeFeature(feature, ({ Scenario, ScenarioOutline }) => {
  Scenario(
    "wahidyankf-www builds against the staging tier",
    ({ Given, When, Then }) => {
      let appDir: string;
      let env: EnvRecord;

      Given('only ".env.stag" exists in the app directory', () => {
        appDir = makeTmpAppDir();
        writeEnvFile(
          appDir,
          ".env.stag",
          "SHARED_VAR=stag-value\nSTAG_ONLY_VAR=stag-only\n",
        );
      });

      When('"next build" runs with APP_ENV set to "stag"', () => {
        env = { APP_ENV: "stag" };
        loadAtCurrentLayer(appDir, env);
      });

      // @covers specs/apps/wahidyankf-www/behaviours/env-loader.feature:wahidyankf-www builds against the staging tier
      Then(
        'every variable consumed by the build resolves to its ".env.stag" value',
        () => {
          expect(env.SHARED_VAR).toBe("stag-value");
          expect(env.STAG_ONLY_VAR).toBe("stag-only");
        },
      );
    },
  );

  Scenario(
    "wahidyankf-www process env wins over the local tier file",
    ({ Given, When, Then, And }) => {
      let appDir: string;
      let env: EnvRecord;

      Given('".env.local" sets an app variable to a file value', () => {
        appDir = makeTmpAppDir();
        writeEnvFile(appDir, ".env.local", "SOME_VAR=file-value\n");
      });

      When(
        'the process starts with that variable already exported at tier "local"',
        () => {
          env = { APP_ENV: "local", SOME_VAR: "process-value" };
          loadAtCurrentLayer(appDir, env);
        },
      );

      Then("the exported process value is used", () => {
        expect(env.SOME_VAR).toBe("process-value");
      });

      // @covers specs/apps/wahidyankf-www/behaviours/env-loader.feature:wahidyankf-www process env wins over the local tier file
      And('the ".env.local" value is not applied over it', () => {
        expect(env.SOME_VAR).not.toBe("file-value");
      });
    },
  );

  Scenario(
    "wahidyankf-www tolerates a missing tier file",
    ({ Given, When, Then, And }) => {
      let appDir: string;
      let env: EnvRecord;
      let thrown: unknown;

      Given('no ".env.stag" file exists in the app directory', () => {
        appDir = makeTmpAppDir();
      });

      When('the loader runs with APP_ENV set to "stag"', () => {
        env = { APP_ENV: "stag", EXISTING_VAR: "already-set" };
        try {
          loadAtCurrentLayer(appDir, env);
        } catch (error) {
          thrown = error;
        }
      });

      Then("the loader does not throw", () => {
        expect(thrown).toBeUndefined();
      });

      // @covers specs/apps/wahidyankf-www/behaviours/env-loader.feature:wahidyankf-www tolerates a missing tier file
      And(
        "startup proceeds using whatever the process environment already supplies",
        () => {
          expect(env.EXISTING_VAR).toBe("already-set");
        },
      );
    },
  );

  ScenarioOutline(
    "wahidyankf-www fails loudly on a stray auto-loaded env file",
    ({ Given, When, Then }, variables) => {
      let appDir: string;
      let thrown: unknown;

      Given('a stray "<file>" sits beside the app\'s tier file', () => {
        appDir = makeTmpAppDir();
        writeEnvFile(appDir, ".env.stag", "VAR=value\n");
        writeEnvFile(appDir, String(variables.file), "VAR=other\n");
      });

      When("the loader runs with APP_ENV set to a non-local tier", () => {
        const env: EnvRecord = { APP_ENV: "stag" };
        try {
          loadAtCurrentLayer(appDir, env);
        } catch (error) {
          thrown = error;
        }
      });

      // @covers specs/apps/wahidyankf-www/behaviours/env-loader.feature:wahidyankf-www fails loudly on a stray auto-loaded env file
      Then(
        'the loader throws, naming "<file>" and the correct ".env.<tier>" replacement',
        () => {
          expect(thrown).toBeInstanceOf(Error);
          expect((thrown as Error).message).toContain(String(variables.file));
          expect((thrown as Error).message).toContain(".env.stag");
        },
      );
    },
  );
});
