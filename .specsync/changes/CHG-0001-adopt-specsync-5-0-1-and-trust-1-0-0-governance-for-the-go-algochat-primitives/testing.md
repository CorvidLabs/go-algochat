---
change: CHG-0001-adopt-specsync-5-0-1-and-trust-1-0-0-governance-for-the-go-algochat-primitives
artifact: testing
---

# Testing

- Strict SpecSync at advisory threshold zero
- All four agent integrations and Trust doctor
- `go build ./...`
- `go vet ./...`
- `go test -v -race -count=1 ./...`
- Existing hosted Go compatibility matrix
