# Learnings

The Markdown-link checker enumerates Git-tracked Markdown paths, so a plan lifecycle move must be staged before link verification. Running the check first makes it inspect the deleted source path even when the destination exists; stage the move, then validate links before committing.
