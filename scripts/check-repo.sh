#!/usr/bin/env bash
# ==============================================================================
# check-repo.sh -- every deterministic repository mechanism, in order
# ==============================================================================
# This is what `npm run test:repo` runs.
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

# Markdown and config checks. RHINO runs the Markdown link and Mermaid checks
# and reads their policy from repo-config.yml.
run ./rhino repo-config validate
run ./rhino md internal-link validate
run ./rhino md mermaid validate
run ./node_modules/.bin/markdownlint-cli2

# British spelling in repository-owned terms. The two exclusions are a React
# component and its test, where the American form is part of a third-party API
# name rather than this repository's own prose.
printf '\n[repo] British spelling\n'
if rg -n -i 'behavio[r]' \
	apps specs scripts .github .husky \
	--glob '!apps/wahidyankf-www/src/features/ui/shell/scroll-to-top.tsx' \
	--glob '!apps/wahidyankf-www/src/features/ui/shell/scroll-to-top.unit.test.tsx'; then
	echo 'ERROR: repository-owned terms must use the British spelling behaviour' >&2
	exit 1
fi

printf '\n[repo] all checks passed\n'
