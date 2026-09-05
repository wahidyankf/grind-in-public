# GitHub Workflow Scripts

- [`test-hippo-bootstrap.sh`](test-hippo-bootstrap.sh) verifies exact lock identity, checksum and warm-cache rejection,
  concurrent and stale installer ownership, unsupported-platform failure, cache reuse, and Nx/Go environment forwarding
  with a synthetic local release.
- [`test-pre-push-contract.sh`](test-pre-push-contract.sh) verifies multi-ref/deletion handling, exact guarded Nx
  arguments, adaptive project selection, shared repository/link gates, and failure propagation without pushing.
