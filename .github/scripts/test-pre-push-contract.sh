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
	PATH="$fake_bin:$PATH" PREPUSH_INVOCATIONS="$invocations" PREPUSH_GIT_CALLS="$git_calls" \
		sh ./pre-push <<EOF
refs/heads/new $new_sha refs/heads/new $zero_sha
refs/heads/existing $existing_sha refs/heads/existing $remote_sha
refs/heads/deleted $zero_sha refs/heads/deleted $remote_sha
EOF
)

[ "$(wc -l <"$invocations" | tr -d ' ')" -eq 4 ]
grep -Fqx "cloud=true run --class ephemeral --disk-path . -- npm exec -- nx affected -t test:quick --base=origin/main --head=$new_sha" "$invocations"
grep -Fqx "cloud=true run --class ephemeral --disk-path . -- npm exec -- nx affected -t test:quick --base=origin/main --head=$existing_sha" "$invocations"
[ "$(grep -Fxc 'cloud=true run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t test:repo' "$invocations")" -eq 1 ]
[ "$(grep -Fxc 'cloud=true run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t markdown-links' "$invocations")" -eq 1 ]
if grep -Fq -- '--parallel=1' "$invocations" || grep -Fq -- '--parallel=1' "$repository_root/.husky/pre-push"; then
	printf '%s\n' 'pre-push must leave Nx project parallelism allocation-driven' >&2
	exit 1
fi
grep -Fq "$new_sha" "$git_calls"
grep -Fq "$existing_sha" "$git_calls"
if grep -Fq "$zero_sha" "$git_calls"; then
	printf '%s\n' 'deleted refs must not reach hook diff evaluation' >&2
	exit 1
fi

: >"$invocations"
set +e
(
	cd "$fixture_root"
	PATH="$fake_bin:$PATH" PREPUSH_INVOCATIONS="$invocations" PREPUSH_GIT_CALLS="$git_calls" \
		PREPUSH_FAIL_MATCH=markdown-links sh ./pre-push <<EOF
refs/heads/existing $existing_sha refs/heads/existing $remote_sha
EOF
)
status=$?
set -e
[ "$status" -eq 42 ]

printf '%s\n' 'pre-push contract tests passed'
