# shellcheck shell=bash
# A pushed range is screened commit by commit. A value one commit adds and the
# next deletes is still in history for every clone, so the range's final files
# being clean is not enough. Content the range did not add is not rescreened:
# the rule binds from its adoption onward, not retroactively.
run() {
	local repo home_path base clean_head leaky_head clean_tip leaky_tip merge_head out rc
	repo="$CASE_TMP/repo"
	mkdir -p "$repo/scripts/public-safety" || return 1
	cp "$PUBLIC_SAFETY_ROOT/scripts/public-safety/check.sh" \
		"$PUBLIC_SAFETY_ROOT/scripts/public-safety/outbound-preflight.sh" \
		"$PUBLIC_SAFETY_ROOT/scripts/public-safety/shape-terms.txt" \
		"$repo/scripts/public-safety/" || return 1
	chmod +x "$repo/scripts/public-safety/check.sh" "$repo/scripts/public-safety/outbound-preflight.sh"

	# Assembled at run time so this file never carries the shape it tests.
	home_path=$(printf '/%s/%s/notes' Users fixtureuser)

	fixture_commit() {
		git -C "$repo" add -A &&
			git -C "$repo" -c user.name=fixture -c user.email=fixture@example.invalid \
				commit -q -m "$1"
	}

	# One ref-update line, as Git hands it to the pre-push hook.
	push_update() {
		printf 'refs/heads/main %s refs/heads/main %s\n' "$2" "$1"
	}

	{
		git init -q "$repo" &&
			printf 'kept from before adoption: %s\nsecond line\n' "$home_path" >"$repo/old.md" &&
			fixture_commit "base" &&
			base=$(git -C "$repo" rev-parse HEAD) &&
			printf 'kept from before adoption: %s\nedited second line\n' "$home_path" >"$repo/old.md" &&
			printf 'see ~/notes for details\n' >"$repo/new.md" &&
			fixture_commit "portable change" &&
			clean_head=$(git -C "$repo" rev-parse HEAD) &&
			printf 'see ~/notes for details\nscratch %s\n' "$home_path" >"$repo/new.md" &&
			fixture_commit "add scratch" &&
			printf 'see ~/notes for details\n' >"$repo/new.md" &&
			fixture_commit "drop scratch" &&
			leaky_head=$(git -C "$repo" rev-parse HEAD) &&
			# The tree screen would flag the pre-adoption file on its own, so the
			# fixture's final tree drops it: what remains to judge is history.
			git -C "$repo" rm -q old.md &&
			fixture_commit "retire old notes" &&
			git -C "$repo" checkout -q -b clean "$clean_head" &&
			git -C "$repo" rm -q old.md &&
			fixture_commit "retire old notes on the clean line" &&
			git -C "$repo" checkout -q main
	} >/dev/null 2>&1 || {
		echo "    cannot build the fixture history" >&2
		return 1
	}
	clean_tip=$(git -C "$repo" rev-parse clean) || return 1
	leaky_tip=$(git -C "$repo" rev-parse main) || return 1

	# Untouched pre-existing content and a `~/` path both pass.
	out=$(cd "$repo" && OSE_GATE_SURFACE=ci PUBLIC_SAFETY_BASE="$base" \
		PUBLIC_SAFETY_HEAD="$clean_head" bash scripts/public-safety/check.sh 2>&1)
	rc=$?
	assert_exit 0 "$rc" "exit code for a range adding only portable content" || return 1

	# The final files are clean, but one commit in the range added the path.
	out=$(cd "$repo" && OSE_GATE_SURFACE=ci PUBLIC_SAFETY_BASE="$base" \
		PUBLIC_SAFETY_HEAD="$leaky_head" bash scripts/public-safety/check.sh 2>&1)
	rc=$?
	assert_exit 1 "$rc" "exit code for a path added then deleted inside the range" || return 1
	assert_contains "maintainer-path" "$out" "diagnostic" || return 1
	assert_contains "/new.md:2" "$out" "finding location" || return 1
	assert_absent "fixtureuser" "$out" "diagnostic" || return 1

	# A half-declared range is a scan error, never a quiet fall back to the tree.
	out=$(cd "$repo" && OSE_GATE_SURFACE=ci PUBLIC_SAFETY_BASE="$base" \
		bash scripts/public-safety/check.sh 2>&1)
	rc=$?
	assert_exit 2 "$rc" "exit code for an incomplete declared range" || return 1

	# The pre-push hook screens the same history from Git's ref-update lines.
	out=$(cd "$repo" && push_update "$base" "$clean_tip" |
		OSE_GATE_SURFACE=pre-push bash scripts/public-safety/check.sh 2>&1)
	rc=$?
	assert_exit 0 "$rc" "pre-push exit code for a clean pushed range" || return 1

	out=$(cd "$repo" && push_update "$base" "$leaky_tip" |
		OSE_GATE_SURFACE=pre-push bash scripts/public-safety/check.sh 2>&1)
	rc=$?
	assert_exit 1 "$rc" "pre-push exit code for a path added then deleted in the push" || return 1
	assert_contains "/new.md:2" "$out" "pre-push finding location" || return 1
	assert_absent "fixtureuser" "$out" "pre-push diagnostic" || return 1

	# A merge commit contributes only what it resolved beyond the automatic
	# merge, so merging a clean side branch keeps a clean range clean.
	{
		git -C "$repo" checkout -q -b side "$clean_tip" &&
			printf 'side note\n' >"$repo/side.md" &&
			fixture_commit "side note" &&
			git -C "$repo" checkout -q clean &&
			git -C "$repo" -c user.name=fixture -c user.email=fixture@example.invalid \
				merge -q --no-ff -m "merge side" side &&
			merge_head=$(git -C "$repo" rev-parse HEAD)
	} >/dev/null 2>&1 || {
		echo "    cannot build the fixture merge" >&2
		return 1
	}
	out=$(cd "$repo" && push_update "$clean_tip" "$merge_head" |
		OSE_GATE_SURFACE=pre-push bash scripts/public-safety/check.sh 2>&1)
	rc=$?
	assert_exit 0 "$rc" "pre-push exit code for a range ending in a clean merge" || return 1

	# A ref update that does not carry commit IDs is refused, not guessed at.
	out=$(cd "$repo" && printf 'refs/heads/main %s refs/heads/main not-a-commit\n' "$clean_tip" |
		OSE_GATE_SURFACE=pre-push bash scripts/public-safety/check.sh 2>&1)
	rc=$?
	assert_exit 2 "$rc" "pre-push exit code for a malformed ref update" || return 1
}
