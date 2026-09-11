#!/usr/bin/env bash
# ==============================================================================
# check-repo.sh -- every deterministic repository mechanism, in order
# ==============================================================================
# This is what `npm run test:repo` runs. It used to be an Nx target inside the
# Badak Mini project, which meant the whole aggregate depended on a Go build
# for the sake of one Go command among eleven. Nine of the others were node,
# shell, or RHINO and never needed a project at all.
#
# Ordered cheapest-first so an obvious failure arrives before an expensive one.
# No parallelism: the output is read by a human deciding what to fix, and
# interleaved output from ten checks is not read, it is scrolled past.
# ==============================================================================

set -euo pipefail

root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
cd "$root"

run() {
	printf '\n[repo] %s\n' "$*"
	"$@"
}

# The declared gates, dispatched by surface so what runs is read from
# repo-config.yml rather than listed twice.
run ./rhino gate run --surface ci

# The repository's own mechanisms: each validator paired with the unit tests
# that prove the validator itself still discriminates. The rule-change notice
# is the exception: it is a reporting check, so running it here would announce
# this run's own staged tree. Its entrypoint is driven by a test instead.
run node --test scripts/project-contract.test.mjs
run node scripts/check-project-contract.mjs
run node --test scripts/governance-structure.test.mjs
run node scripts/check-governance-structure.mjs
run node --test scripts/workflow-contract.test.mjs
run node scripts/check-workflow-contract.mjs
run node --test scripts/rule-change.test.mjs
run node --test scripts/check-rule-change.test.mjs

# The two contracts that can only be checked by running the real thing.
run ./.github/scripts/test-hippo-bootstrap.sh
run ./.github/scripts/test-pre-push-contract.sh

# British spelling in repository-owned terms. The two exclusions are a React
# component and its test, where the American form is part of a third-party API
# name rather than this repository's own prose.
printf '\n[repo] British spelling\n'
if rg -n -i 'behavio[r]' \
	AGENTS.md CLAUDE.md RTK.md \
	repo-governance apps specs scripts .github .husky .agents .claude .codex .opencode \
	--glob '!apps/wahidyankf-www/src/features/ui/shell/scroll-to-top.tsx' \
	--glob '!apps/wahidyankf-www/src/features/ui/shell/scroll-to-top.unit.test.tsx'; then
	echo 'ERROR: repository-owned terms must use the British spelling behaviour' >&2
	exit 1
fi

printf '\n[repo] all checks passed\n'
