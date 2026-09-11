#!/bin/sh
set -eu

repository_root=$(CDPATH='' cd -- "$(dirname -- "$0")/../.." && pwd)
temporary_root=$(mktemp -d)
trap 'rm -rf -- "$temporary_root"' EXIT HUP INT TERM

fixture_root="$temporary_root/work"
fake_bin="$temporary_root/fake-bin"
invocations="$temporary_root/hippo-invocations"
git_calls="$temporary_root/git-calls"
mkdir -p "$fixture_root" "$fake_bin"
cp "$repository_root/.husky/pre-push" "$fixture_root/pre-push"

cat >"$fake_bin/git" <<'EOF'
#!/bin/sh
set -eu
case "${1:-}" in
hash-object)
	printf '%s\n' aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	;;
rev-parse)
	exit 0
	;;
diff)
	printf '%s\n' "$*" >>"$PREPUSH_GIT_CALLS"
	case " $* " in
	*" 1111111111111111111111111111111111111111 "*) exit 1 ;;
	*) exit 0 ;;
	esac
	;;
*)
	printf 'unexpected git invocation: %s\n' "$*" >&2
	exit 96
	;;
esac
EOF
chmod 755 "$fake_bin/git"

cat >"$fixture_root/hippo" <<'EOF'
#!/bin/sh
set -eu
printf 'cloud=%s' "${NX_NO_CLOUD:-unset}" >>"$PREPUSH_INVOCATIONS"
for argument in "$@"; do
	printf ' %s' "$argument" >>"$PREPUSH_INVOCATIONS"
done
printf '\n' >>"$PREPUSH_INVOCATIONS"
case " $* " in
*" ${PREPUSH_FAIL_MATCH:-__never__} "*) exit 42 ;;
*) exit 0 ;;
esac
EOF
chmod 755 "$fixture_root/hippo"

zero_sha=0000000000000000000000000000000000000000
new_sha=1111111111111111111111111111111111111111
existing_sha=2222222222222222222222222222222222222222
remote_sha=3333333333333333333333333333333333333333

(
	cd "$fixture_root"
	# The guard stub records whether the Nx cloud opt-out was set on each
	# invocation, and the assertions below distinguish the lines the hook opts
	# out on from the one it does not. That distinction only exists if the
	# variable is absent to begin with. This file is also run by `test:repo`,
	# which the hook itself invokes as `NX_NO_CLOUD=true ... npm run test:repo`
	# -- so when the push runs, the variable is already exported into
	# everything below, every recorded line reads `cloud=true`, and the
	# assertion becomes a statement about the ambient environment rather than
	# about the hook. Clear it here so the stub observes only what the hook set.
	unset NX_NO_CLOUD
	PATH="$fake_bin:$PATH" PREPUSH_INVOCATIONS="$invocations" PREPUSH_GIT_CALLS="$git_calls" \
		sh ./pre-push <<EOF
refs/heads/new $new_sha refs/heads/new $zero_sha
refs/heads/existing $existing_sha refs/heads/existing $remote_sha
refs/heads/deleted $zero_sha refs/heads/deleted $remote_sha
EOF
)

# Every assertion below says what it expected. A bare test that exits silently
# under `set -e` reports only that something in this file was false, which is
# how a stale expectation here survived a gate migration unnoticed.
fail() {
	printf '%s\n' "$1" >&2
	printf 'recorded guard invocations:\n' >&2
	cat "$invocations" >&2
	exit 1
}

[ "$(wc -l <"$invocations" | tr -d ' ')" -eq 4 ] ||
	fail 'pre-push must make exactly four guarded invocations: the surface screen, one affected run per non-deleted ref, and the repository aggregate'

# The screen runs through the declared surface rather than as a list of
# commands here, so what it enforces is read from repo-config.yml. It is also
# the only one that keeps Nx cloud unset, because no Nx target runs behind it.
grep -Fqx 'cloud=unset run --class ephemeral --disk-path . -- ./rhino gate run --surface pre-push' "$invocations" ||
	fail 'pre-push must screen the refs through the declared pre-push gate surface before anything else'

grep -Fqx "cloud=true run --class ephemeral --disk-path . -- npm exec -- nx affected -t test:quick --base=origin/main --head=$new_sha" "$invocations" ||
	fail 'a newly created ref must be checked against origin/main'
grep -Fqx "cloud=true run --class ephemeral --disk-path . -- npm exec -- nx affected -t test:quick --base=origin/main --head=$existing_sha" "$invocations" ||
	fail 'an existing ref must be checked against origin/main rather than against its own remote tip'
[ "$(grep -Fxc 'cloud=true run --class ephemeral --disk-path . -- npm run test:repo' "$invocations")" -eq 1 ] ||
	fail 'a governance change anywhere in the pushed range must run the repository aggregate exactly once'

# Link validation is what this hook adds over the repository gate: a rename
# breaks links in files the change never touched. It moved onto the declared
# surface, so the assertion moved with it -- the behaviour is still bound, at
# the site that now owns it.
grep -q '^  - id: internal-link$' "$repository_root/repo-config.yml" ||
	fail 'repo-config.yml must still declare an internal-link gate'
awk '/^  - id: internal-link$/ { found = 1 } found && /^ *- pre-push$/ { print; exit }' \
	"$repository_root/repo-config.yml" | grep -q 'pre-push' ||
	fail 'the internal-link gate must stay on the pre-push surface, or a rename can be pushed with broken links'

if grep -Fq -- '--parallel=1' "$invocations" || grep -Fq -- '--parallel=1' "$repository_root/.husky/pre-push"; then
	fail 'pre-push must leave Nx project parallelism allocation-driven'
fi
grep -Fq "$new_sha" "$git_calls" || fail 'a newly created ref must reach hook diff evaluation'
grep -Fq "$existing_sha" "$git_calls" || fail 'an existing ref must reach hook diff evaluation'
if grep -Fq "$zero_sha" "$git_calls"; then
	fail 'deleted refs must not reach hook diff evaluation'
fi

# A failing screen stops the push, and stops it before any Nx work is spent.
: >"$invocations"
set +e
(
	cd "$fixture_root"
	PATH="$fake_bin:$PATH" PREPUSH_INVOCATIONS="$invocations" PREPUSH_GIT_CALLS="$git_calls" \
		PREPUSH_FAIL_MATCH=pre-push sh ./pre-push <<EOF
refs/heads/existing $existing_sha refs/heads/existing $remote_sha
EOF
)
status=$?
set -e
[ "$status" -eq 42 ] || fail "a failing pre-push surface must stop the push, got status $status"
[ "$(wc -l <"$invocations" | tr -d ' ')" -eq 1 ] ||
	fail 'a failing screen must stop the push before any affected run is spent'

printf '%s\n' 'pre-push contract tests passed'
