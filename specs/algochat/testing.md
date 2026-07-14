---
spec: algochat.spec.md
---

## Test Plan

- Run all existing cryptography, envelope, account, signature, queue, model, and storage tests with the race detector.

- Build, vet, and race-test all packages on Go 1.25 in the Trust lane.
- Preserve the repository's existing Go-version compatibility matrix independently.
