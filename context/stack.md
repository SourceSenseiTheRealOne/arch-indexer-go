# Stack

- Language: Go 1.26.0 with toolchain `go1.26.5`.
- Runtime target: Linux AMD64, statically built with CGO disabled.
- Authority: PostgreSQL for canonical events, checkpoints, projections, and transactional outbox state.
- Delivery: Redis as a recoverable at-least-once transport, never canonical storage.
- Protocol boundary: strict read-only Arch JSON-RPC allowlist.
- Initial adapter: Bump, isolated from the reusable engine core.
- Branches: `development`, `staging`, and `main`, promoted in that order.
- Full local gate: `bash scripts/verify.sh`.
