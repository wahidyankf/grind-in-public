import "./src/env-loader.ts";
import "./src/env.ts";
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // `next dev` detects an AI coding agent and, by default, writes its own AGENTS.md and CLAUDE.md
  // into this directory so the agent reads version-matched Next docs instead of stale training data.
  // This repository already answers that question its own way: AGENTS.md at the root is authoritative
  // and CLAUDE.md derives from it, under the agent instruction alignment policy, and the harness
  // parity and governance checks are scoped to those files. A generated pair here would be two more
  // instruction files no policy owns and no check reads, and the generated text tells an agent to
  // commit them. Disabled rather than gitignored, so the files are never written at all.
  agentRules: false,
  transpilePackages: ["@t3-oss/env-nextjs", "@t3-oss/env-core"],
  images: {
    unoptimized: true,
  },
};

export default nextConfig;
