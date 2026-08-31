/**
 * Wired as the very first import in `next.config.ts`, before `./env.ts`, so every other config
 * module — including `./env.ts`'s `createEnv()` validation — observes the tier file's values.
 *
 * The actual loader logic — tier resolution, the stray-file guard, and the process-env-wins
 * application — lives in `./features/env/core/tier-env`, which holds the full contract. This
 * file's only job is the explicit call: that module never loads its own tier file on import,
 * because which loader won would then depend on import order, so the composition root calls
 * `loadTierEnv()` itself.
 *
 * The specifier below is relative rather than the `@/` alias every other module here uses, and
 * that is deliberate. Next compiles `next.config.ts` through its own config transpiler, which
 * does not read this project's `paths` mapping — an aliased import from this file resolves
 * against the wrong base and the build dies with `Cannot find module` before it starts. This file
 * is the one module loaded from inside that pipeline, so it is the one module that cannot rely on
 * the alias.
 */
import { loadTierEnv } from "./features/env/core/tier-env";

export {
  loadTierEnv,
  resolveTier,
  tierEnvFilePath,
} from "./features/env/core/tier-env";

loadTierEnv();
