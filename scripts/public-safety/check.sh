#!/usr/bin/env bash
# ==============================================================================
# check.sh — this repository's public-safety gate
# ==============================================================================
# Usage: OSE_GATE_SURFACE=<surface> scripts/public-safety/check.sh [hook args]
#
#   commit-msg   $1 is the file holding the message being written
#   pre-commit   no arguments; the staged tree is the subject
#   pre-push     ref updates arrive on stdin, as Git supplies them; each
#                pushed range is screened commit by commit
#   ci           no arguments; the checked-out tree is the subject, unless
#                PUBLIC_SAFETY_BASE and PUBLIC_SAFETY_HEAD declare one range,
#                which is then screened commit by commit instead
#
# The surface arrives in the environment and nowhere else. It is never inferred
# from an argument's filename, from which hook happens to be running, or from
# whether a remote is reachable — a gate that guesses its own surface will
# eventually guess a weaker one, and that is precisely the case where guessing
# is expensive.
#
# This file decides *what is outbound* at each surface. `outbound-preflight.sh`
# decides whether any of it is prohibited. Keeping those apart is what lets the
# leaf be tested against synthetic inputs with no repository at all.
#
# Exit codes pass through from the leaf: 0 clean, 1 blocked, 2 scan error.
# ==============================================================================

set -uo pipefail

here=$(CDPATH='' cd -- "$(dirname -- "$0")" && pwd)
root=$(git -C "$here" rev-parse --show-toplevel 2>/dev/null) || {
	printf '[public-safety] blocked scan-error not inside a Git repository\n' >&2
	exit 2
}
leaf="$here/outbound-preflight.sh"
[[ -x "$leaf" ]] || {
	printf '[public-safety] blocked scan-error the leaf wrapper is missing or not executable\n' >&2
	exit 2
}

surface="${OSE_GATE_SURFACE:-${RHINO_GATE_SURFACE:-}}"
case "$surface" in
main | scheduled | manual | pull-request) surface=ci ;;
esac
case "$surface" in
commit-msg | pre-commit | pre-push | ci) ;;
"")
	printf '[public-safety] blocked scan-error OSE_GATE_SURFACE is unset\n' >&2
	exit 2
	;;
*)
	printf '[public-safety] blocked scan-error OSE_GATE_SURFACE is not a known surface\n' >&2
	exit 2
	;;
esac

cd "$root" || exit 2

# ------------------------------------------------------------------------------
# Input collection. Every helper appends to `args`, an argument vector — never a
# string. A path containing a space or a shell metacharacter is data here, and
# building a command line out of it would make it an instruction.
# ------------------------------------------------------------------------------

declare -a args=()
# Inputs staged by add_range_history, screened from inside `$history`.
declare -a history_args=()
history=""

# Paths reach the leaf as NUL-delimited list files, never as one argument per
# path: a large tracked tree overflows the host's argument limit, and a gate
# that cannot start screens nothing.
lists=$(mktemp -d "${TMPDIR:-/tmp}/public-safety-lists.XXXXXX") || {
	printf '[public-safety] blocked scan-error cannot create a working directory\n' >&2
	exit 2
}
trap 'rm -rf "$lists"' EXIT INT TERM
list_count=0

add_paths() {
	# add_paths <command...>: the command prints NUL-delimited paths. Every
	# existing file is content, and every path is a name: a name is outbound
	# material too, and a directory named after something private leaks whether
	# or not any file inside it does.
	local files names f
	list_count=$((list_count + 1))
	files="$lists/files-$list_count"
	names="$lists/names-$list_count"
	while IFS= read -r -d '' f; do
		[[ -f "$f" ]] && printf '%s\0' "$f" >&3
		printf '%s\0' "$f" >&4
	done < <("$@") 3>"$files" 4>"$names"
	[[ -s "$files" ]] && args+=(--file-list "$files")
	[[ -s "$names" ]] && args+=(--names-list "$names")
}

add_tracked_tree() {
	add_paths git ls-files -z
}

add_staged() {
	add_paths git diff --cached --name-only -z --diff-filter=ACMR
}

range_error() {
	printf '[public-safety] blocked scan-error %s\n' "$1" >&2
	exit 2
}

is_commit_id() {
	[[ "$1" =~ ^[0-9a-fA-F]{7,64}$ ]]
}

added_lines() {
	# added_lines: reads one file's zero-context diff on standard input and
	# prints the lines it adds, each at the line number it occupies after the
	# change with blank lines between, so a finding names a real location in
	# that commit's file. Exits 3 for a binary change, whose added content a
	# text diff cannot show.
	awk '
		/^Binary files / { binary = 1; exit }
		/^@@ / {
			header = $0
			sub(/^@@ -[0-9,]+ \+/, "", header)
			split(header, at, /[, ]/)
			line = at[1] + 0
			in_hunk = 1
			next
		}
		in_hunk && /^\+/ {
			added[line] = substr($0, 2)
			if (line > last) last = line
			line++
		}
		END {
			if (binary) exit 3
			for (i = 1; i <= last; i++) print ((i in added) ? added[i] : "")
		}
	'
}

add_range_history() {
	# add_range_history <rev-list arguments...>: every commit in the range is
	# outbound on its own. History keeps what a later commit deletes, so
	# screening only the range's final files would pass a value that the next
	# commit removed while every clone still carries it. Each commit contributes
	# only the lines it added, staged as `<commit>/<path>` under `$history`;
	# content the range did not touch is never rescreened, so the rule binds
	# from its adoption onward. A merge contributes only what it resolved
	# beyond the automatic merge. The caller validates every commit ID first.
	local commits commit parent empty_tree status old new out
	local -a diff_cmd statuses
	commits=$(git rev-list --reverse "$@") || range_error "cannot list the range's commits"
	empty_tree=$(git hash-object -t tree /dev/null) || range_error "cannot resolve the empty tree"

	history="$lists/history"
	list_count=$((list_count + 1))
	local files="$lists/files-$list_count" names="$lists/names-$list_count"
	: >"$files"
	: >"$names"

	for commit in $commits; do
		# Two refs pushed together can share commits; stage each one once.
		[[ -d "$history/${commit:0:12}" ]] && continue
		mkdir -p "$history/${commit:0:12}" || range_error "cannot stage a commit's additions"
		if git rev-parse --quiet --verify "$commit^2" >/dev/null; then
			diff_cmd=(show --format= --remerge-diff "$commit")
		else
			parent=$(git rev-parse --quiet --verify "$commit^1") || parent=$empty_tree
			diff_cmd=(diff "$parent" "$commit")
		fi
		git "${diff_cmd[@]}" -M --name-status -z --diff-filter=ACMR >"$lists/changed" ||
			range_error "cannot list a commit's changes"
		while IFS= read -r -d '' status; do
			IFS= read -r -d '' old || range_error "a commit's change list is truncated"
			new=$old
			if [[ "$status" == R* ]]; then
				IFS= read -r -d '' new || range_error "a commit's change list is truncated"
			fi
			printf '%s\0%s\0' "$old" "$new" >>"$names"
			out="$history/${commit:0:12}/$new"
			mkdir -p "$(dirname -- "$out")" || range_error "cannot stage a commit's additions"
			git "${diff_cmd[@]}" -M -U0 --no-color --no-ext-diff --no-textconv \
				-- ":(literal)$old" ":(literal)$new" | added_lines >"$out"
			statuses=("${PIPESTATUS[@]}")
			[[ ${statuses[0]} -eq 0 ]] || range_error "cannot read a commit's additions"
			if [[ ${statuses[1]} -eq 3 ]]; then
				git cat-file blob "$commit:$new" >"$out" || range_error "cannot read a binary addition"
			elif [[ ${statuses[1]} -ne 0 ]]; then
				range_error "cannot stage a commit's additions"
			fi
			if [[ -s "$out" ]]; then
				printf '%s\0' "${commit:0:12}/$new" >>"$files"
			else
				rm -f -- "$out"
			fi
		done <"$lists/changed"
	done

	[[ -s "$files" ]] && history_args+=(--file-list "$files")
	[[ -s "$names" ]] && args+=(--names-list "$names")
	return 0
}

run_leaf() {
	# run_leaf <leaf-surface>; consumes and clears `args`.
	local leaf_surface=$1 rc
	if [[ ${#args[@]} -eq 0 ]]; then
		args=()
		return 0
	fi
	"$leaf" --surface "$leaf_surface" "${args[@]}"
	rc=$?
	args=()
	return $rc
}

run_history_leaf() {
	# run_history_leaf <leaf-surface>: screens what add_range_history staged.
	# The leaf runs inside the staging directory so every finding is labelled
	# `<commit>/<path>:<line>` rather than with a temporary absolute path.
	local leaf_surface=$1
	[[ ${#history_args[@]} -eq 0 ]] && return 0
	(cd "$history" && "$leaf" --surface "$leaf_surface" "${history_args[@]}")
}

current_ref() {
	git symbolic-ref --quiet --short HEAD 2>/dev/null || git rev-parse --short HEAD
}

# ------------------------------------------------------------------------------
# Surfaces
# ------------------------------------------------------------------------------

case "$surface" in
commit-msg)
	[[ $# -ge 1 && -r "$1" ]] || {
		printf '[public-safety] blocked scan-error commit-msg received no readable message file\n' >&2
		exit 2
	}
	args+=(--file "$1" --text "$(current_ref)")
	run_leaf commit || exit $?
	;;

pre-commit)
	# The tracked tree first: a leak that is already committed does not become
	# safe because this particular change did not introduce it.
	add_tracked_tree
	run_leaf baseline || exit $?
	add_staged
	run_leaf diff || exit $?
	;;

pre-push)
	# Git supplies `<local-ref> <local-sha> <remote-ref> <remote-sha>` per line.
	# Deletions carry an all-zero local sha and send nothing outbound.
	local_refs=()
	ranges=()
	while read -r local_ref local_sha remote_ref remote_sha; do
		[[ -z "${local_ref:-}" ]] && continue
		[[ "$local_sha" =~ ^0+$ ]] && continue
		# Both IDs become revision arguments below, so anything else is refused
		# rather than handed to Git.
		is_commit_id "$local_sha" && is_commit_id "${remote_sha:-}" ||
			range_error "a ref update does not carry commit IDs"
		local_refs+=("${local_ref#refs/heads/}" "${remote_ref#refs/heads/}")
		if [[ "$remote_sha" =~ ^0+$ ]]; then
			ranges+=("$local_sha --not --remotes")
		else
			ranges+=("$remote_sha..$local_sha")
		fi
	done

	# Invoked outside a hook, with nothing on stdin: screen what is here now
	# rather than reporting a vacuous pass.
	if [[ ${#local_refs[@]} -eq 0 ]]; then
		local_refs=("$(current_ref)")
		ranges=("HEAD --not --remotes")
	fi

	for r in "${local_refs[@]}"; do
		[[ -n "$r" ]] && args+=(--text "$r")
	done
	run_leaf ref || exit $?

	for range in "${ranges[@]}"; do
		# shellcheck disable=SC2086
		while IFS= read -r sha; do
			[[ -n "$sha" ]] && args+=(--text "$(git log -1 --format=%B "$sha")")
		done < <(git rev-list $range 2>/dev/null)
	done
	run_leaf commit || exit $?

	# Every pushed commit's own additions and file names, so a value one
	# commit adds and the next deletes cannot pass on a clean final tree.
	for range in "${ranges[@]}"; do
		# shellcheck disable=SC2086
		add_range_history $range
	done
	run_leaf diff || exit $?
	run_history_leaf diff || exit $?

	add_tracked_tree
	run_leaf baseline || exit $?
	;;

ci)
	if [[ -n "${PUBLIC_SAFETY_BASE:-}" || -n "${PUBLIC_SAFETY_HEAD:-}" ]]; then
		# A hosted replay of one pushed range, screened exactly as pre-push
		# screens it. The tree is a separate gate there.
		is_commit_id "${PUBLIC_SAFETY_BASE:-}" && is_commit_id "${PUBLIC_SAFETY_HEAD:-}" ||
			range_error "the declared range is incomplete or not commit IDs"
		args+=(--text "$PUBLIC_SAFETY_BASE" --text "$PUBLIC_SAFETY_HEAD")
		run_leaf ref || exit $?
		args+=(--text "$(git log --format=%B "$PUBLIC_SAFETY_BASE..$PUBLIC_SAFETY_HEAD")")
		run_leaf commit || exit $?
		add_range_history "$PUBLIC_SAFETY_BASE..$PUBLIC_SAFETY_HEAD"
		run_leaf diff || exit $?
		run_history_leaf diff || exit $?
	else
		args+=(--text "$(current_ref)" --text "$(git log -1 --format=%B HEAD)")
		run_leaf ref || exit $?
		add_tracked_tree
		run_leaf baseline || exit $?
	fi
	;;
esac

printf '[public-safety] %s: clean\n' "$surface"
exit 0
