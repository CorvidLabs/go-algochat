---
module: algochat
version: 1
status: active
files:
  - doc.go

db_tables: []
depends_on: []
---

# Go AlgoChat Primitives

## Purpose

Provide the existing Go primitives for AlgoChat encryption, envelope encoding, signing, account derivation, message models, offline queues, and in-memory storage while the high-level client remains under development.

## Public API

### Package Interface

The `algochat` package exposes cryptographic key derivation, message encryption and decryption, envelope codecs, key signatures and fingerprints, account primitives, queue operations, and in-memory message and key stores documented by Go package comments and tests.

## Invariants

1. Encryption derives and exchanges X25519 keys and authenticates ciphertext with ChaCha20-Poly1305 according to the existing version-one envelope format.
2. Decoding rejects unsupported versions, protocols, undersized envelopes, and oversized plaintext without returning an apparent success.
3. Queue and in-memory storage operations preserve ordering, retry limits, deduplication, participant isolation, and explicit missing-key behavior.
4. Signing and verification use the expected Ed25519 key and signature sizes and reject mismatches.
5. This migration does not claim the unfinished high-level send/receive client is implemented.

## Behavioral Examples

```
Given two valid AlgoChat accounts and a plaintext within the protocol limit
When one account encrypts and encodes a message for the other
Then the recipient can decode and authenticate the original plaintext while an unrelated key cannot
```

## Error Cases

| Error | When | Behavior |
|-------|------|----------|
| Invalid envelope | Header, version, protocol, or size is invalid | Return an error without plaintext |
| Oversized message | Plaintext exceeds the protocol capacity | Reject encryption |
| Wrong key | Authentication cannot succeed for the supplied recipient | Return an error |
| Full queue | Enqueue would exceed the configured limit | Return `ErrQueueFull` |
| Missing stored key | A requested address has no key | Return `ErrKeyNotFound` |

## Dependencies

- Go 1.25 or newer
- `golang.org/x/crypto` for Curve25519, ChaCha20-Poly1305, and HKDF primitives
- Go standard-library Ed25519, synchronization, and time facilities

## Change Log

| Version | Date | Changes |
|---------|------|---------|
| 1 | 2026-07-12 | Initial spec |
