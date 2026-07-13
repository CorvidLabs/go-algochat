---
spec: algochat.spec.md
---

## Requirements

### REQ-algochat-001

The package SHALL derive keys and encrypt, encode, decode, and decrypt version-one AlgoChat envelopes with authenticated failure behavior.

Acceptance Criteria
- A supported version-one envelope round-trips its original plaintext for the intended recipient.
- A wrong recipient key, tampered ciphertext, unsupported version, or oversized message returns an error.

### REQ-algochat-002

The package SHALL sign and verify encryption-key announcements and produce deterministic key fingerprints.

Acceptance Criteria
- A valid announcement verifies with its signing key, while a mismatched key or signature is rejected.
- Repeated fingerprinting of the same key produces the same value.

### REQ-algochat-003

Queue and storage primitives SHALL preserve their documented limits, transitions, ordering, isolation, retry, and missing-key behavior.

Acceptance Criteria
- Queue tests cover FIFO order, retry limits, deduplication, and terminal failure transitions.
- Storage tests cover participant isolation and explicit missing-key results.

### REQ-algochat-004

Native verification SHALL build, vet, and race-test every Go package without asserting that the unfinished high-level client exists.

Acceptance Criteria
- `go build ./...`, `go vet ./...`, and `go test -race ./...` complete successfully.
- The verification lane does not represent unfinished client send/receive behavior as implemented.

## Constraints

- The module currently requires Go 1.25 or newer.
- The high-level AlgoChat client remains outside the implemented surface.

## Out of Scope

- Implementing send, receive, key discovery, or live Algorand network integration.
