# Arch indexer in Go — foundation

An early-stage foundation for a read-only Arch Network indexer. The current executable checks process configuration; **block indexing is not implemented**. Bump is the intended first protocol adapter, not a dependency of the core or a completed production integration.

## Implemented

- Required configuration fields, bounded timeouts and queue capacity, and a loopback metrics-address default.
- A `--check-config` CLI with explicit success/configuration/usage exit codes and tests that connection strings are not printed.
- Formatting, unit tests, race tests, `go vet` and a Linux build.

The module currently uses the Go standard library only. PostgreSQL and Redis are planned runtime dependencies, not services this executable connects to.

## Try the configuration check

Use the Go toolchain declared in [`go.mod`](go.mod). These illustrative local values do not require running infrastructure:

```bash
DATABASE_URL='postgres://127.0.0.1:5432/arch_indexer?sslmode=disable' \
REDIS_URL='redis://127.0.0.1:6379/0' \
MANIFEST_PATH='deployments/local.json' \
go run ./cmd/indexerd --check-config
```

Output:

```text
configuration valid; indexer runtime is not started
```

This checks required values and configured bounds. It does not verify database/Redis connectivity, validate their URL syntax, read the manifest file or contact Arch RPC. The manifest path above is illustrative, not a supplied deployment manifest.

The binary exits with code `64` when invoked without `--check-config`; it does not start an indexing daemon. Missing required configuration returns `78`. `go run` itself wraps nonzero program exits, so inspect the built binary when testing exact process codes.

## Planned engine

```text
Arch read-only RPC → bounded pipeline → protocol adapters
                    → PostgreSQL ledger + checkpoint + outbox
                    → recoverable at-least-once delivery (Redis)
```

All components in this pipeline are planned. The design makes PostgreSQL authoritative, commits events and progress together, and keeps delivery recoverable. Reorg rollback, deterministic replay and immutable-event conflict checks still need implementation and behavioral proof.

[Planned architecture](docs/architecture.md) · [Phased roadmap](docs/roadmap.md) · [Next work](context/QUEUE.md)

## Verification

```bash
bash scripts/verify.sh
```

The verifier checks formatting, unit/race tests and vet, then builds `bin/indexerd-linux-amd64`. Race tests require CGO and a C compiler; the Linux artifact itself is built with CGO disabled. Use a Go-equipped Linux container if the host lacks the race-test toolchain. A passing foundation verifier does not establish ingestion or recovery correctness.

## Boundaries

No wallets, signing, transaction submission or generic RPC-call surface belong in this project. Public-network checks and any Bump integration require separate approval and evidence. There is no live demo or deployed-indexer claim.

The canonical Lab checkout is `projects/experiments/arch/indexer-go`; the GitHub repository and Go module remain `arch-indexer-go`. Work targets `development`, followed by reviewed `staging` and `main` promotions.
