# Planned architecture

This is the target design, not an inventory of running components. Only process-configuration checks and the CLI tracer are implemented. The datasource, canonical pipeline, adapters, PostgreSQL store, outbox and transport are planned work; see the [implementation roadmap](roadmap.md).

## Objective

Provide a reusable Arch Network indexing engine whose correctness does not depend on Bump-specific code. Bump is a compiled-in adapter and the first real parity target.

## Boundaries

```text
cmd/indexerd
  -> config + reviewed release manifest
  -> Arch read-only datasource
  -> bounded pipeline
  -> adapter interface
  -> PostgreSQL block/event/checkpoint transaction
  -> transactional outbox
  -> transport publisher
```

Dependency rules:

- Datasources know nothing about protocol adapters.
- Adapters know nothing about PostgreSQL, Redis, process configuration, or HTTP servers.
- Store packages own persistence but not decoding policy.
- Transport publishers receive immutable outbox envelopes and cannot advance checkpoints.
- `cmd/indexerd` is the composition root.

## Correctness invariants

1. Fetch complete blocks in contiguous height order.
2. Preserve each transaction's absolute index in the full block array.
3. Identify events by network, transaction ID, and canonical numeric instruction path.
4. Atomically persist block lineage, canonical events, outbox intents, and checkpoint.
5. Never advance progress for partial, malformed, missing, or conflicting data.
6. Treat Redis as recoverable at-least-once delivery, not canonical storage.
7. Detect changed block hashes, find a bounded common ancestor, roll back, and replay deterministically.
8. Fail closed when an immutable event identity reappears with conflicting content.

## Security model

The RPC client will expose only explicit read methods. The repository must never contain wallet loading, signing, transaction submission, or a public generic JSON-RPC call method. Public-network validation is read-only and release-manifest bound.

## Initial tracer

The first executable behavior is `indexerd --check-config`. It validates bounded non-secret process configuration and exits. It does not connect to infrastructure or claim runtime readiness.
