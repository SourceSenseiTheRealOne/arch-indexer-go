# arch-indexer-go

A reusable, read-only Arch Network indexing engine written in Go. Bump will be its first protocol adapter, but the engine core will remain protocol-neutral.

## Status

Foundation scaffold only:

- strict environment configuration with safe defaults and bounds;
- a `--check-config` CLI tracer that never prints connection strings;
- unit/race/vet/build quality gates;
- architecture and phased roadmap documentation.

Block ingestion, persistence, Bump decoding, Redis delivery, health endpoints, and production integration are not implemented yet. The executable refuses to start an indexer runtime rather than pretending the scaffold is operational.

## Architecture

```text
Arch JSON-RPC
    |
    v
Read-only datasource
    |
    v
Bounded canonical pipeline
    |
    v
Compiled-in protocol adapters (Bump first)
    |
    v
PostgreSQL canonical ledger + checkpoint + transactional outbox
    |
    v
Recoverable at-least-once delivery (Redis first)
```

PostgreSQL will be authoritative. Redis will never own canonical state. No wallet, signer, or transaction-submission code belongs in this repository.

See `docs/architecture.md` and `docs/roadmap.md`.

## Requirements

- Go 1.26.5 toolchain (the module can auto-select it through the Go toolchain directive)
- Git
- Docker Desktop later, for PostgreSQL/Redis integration slices

## Configuration tracer

The binary currently validates process configuration only:

```bash
export DATABASE_URL='postgres://127.0.0.1:5432/arch_indexer?sslmode=disable'
export REDIS_URL='redis://127.0.0.1:6379/0'
export MANIFEST_PATH='deployments/local.json'
go run ./cmd/indexerd --check-config
```

Expected output:

```text
configuration valid; indexer runtime is not started
```

Running without `--check-config` exits with a usage error because ingestion is intentionally not implemented yet.

## Quality gates

```bash
bash scripts/verify.sh
```

On hosts with GNU Make, `make verify` delegates to the same script. The script recursively checks formatting, runs native unit and race tests, runs `go vet`, and cross-builds the deployable Linux artifact at `bin/indexerd-linux-amd64`. Override `BUILD_GOOS` and `BUILD_GOARCH` only for an explicitly reviewed deployment target.

## Repository state

The public repository uses `development`, `staging`, and `main` environment branches. New work targets `development`; promotion to `staging` and `main` requires reviewed pull requests. Parent-lab registration, Bump integration, and production configuration are separate gated changes.
