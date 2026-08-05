# Arch Indexer Go Agent Rules

## Mission

Build a reusable, read-only Arch Network indexing engine. Keep protocol behavior in compiled-in adapters; Bump is the first adapter, not a dependency of the engine core.

## Required workflow

1. Read `README.md`, `docs/architecture.md`, and `docs/roadmap.md` before planning or editing.
2. Run `codegraph status .`; initialize or synchronize the child index explicitly.
3. Use strict RED → GREEN → REFACTOR for behavioral code.
4. Keep queues bounded, preserve deterministic ordering, and fail closed on incomplete blocks or immutable identity conflicts.
5. Run `make verify` before reporting a slice complete.
6. Do not commit, push, create remotes, or modify Bump unless Sensei explicitly authorizes it.

## Security boundaries

- No wallet, signer, private-key, transaction-submission, or generic RPC-call surface.
- Never print database/Redis URLs, credentials, raw private data, or request bodies.
- PostgreSQL will own canonical blocks, events, checkpoints, and outbox state.
- Redis is delivery-only and recoverable at-least-once.
- Mainnet activity is read-only and separately evidence-gated.

## Branching

After repository bootstrap, feature branches target `development`; promotions are `development → staging → main`. Never push directly to environment branches after bootstrap.
