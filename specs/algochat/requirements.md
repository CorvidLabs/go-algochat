---
spec: algochat.spec.md
---

## Requirements

- **REQ-go-algochat-001** (stable): The package shall derive keys and encrypt, encode, decode, and decrypt version-one AlgoChat envelopes with authenticated failure behavior.
- **REQ-go-algochat-002** (stable): The package shall sign and verify encryption-key announcements and produce deterministic key fingerprints.
- **REQ-go-algochat-003** (stable): Queue and storage primitives shall preserve their documented limits, transitions, ordering, isolation, retry, and missing-key behavior.
- **REQ-go-algochat-004** (stable): Native verification shall build, vet, and race-test every Go package without asserting that the unfinished high-level client exists.

## Constraints

- The module currently requires Go 1.25 or newer.
- The high-level AlgoChat client remains outside the implemented surface.

## Out of Scope

- Implementing send, receive, key discovery, or live Algorand network integration.
