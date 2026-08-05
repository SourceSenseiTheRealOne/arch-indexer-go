# Implementation Roadmap

The detailed implementation plan is maintained in the coding lab's local Hermes plan store. This repository tracks the durable, non-sensitive phase order and verified status.

## Phase 1 — Foundation

- [x] Go module and toolchain policy
- [x] Strict process configuration loader
- [x] Safe `--check-config` CLI tracer
- [x] Unit, race, vet, format, and build gates
- [x] Architecture, security, and contributor rules
- [x] Local repository foundation prepared for public bootstrap
- [x] Coding-lab-compatible project context and task ledgers

## Phase 2 — Canonical domain

- [ ] Immutable event envelope v1
- [ ] Canonical instruction-path parser and numeric ordering
- [ ] Canonical JSON payload hashing
- [ ] JSON Schema contract

## Phase 3 — Engine boundaries

- [ ] Narrow datasource, adapter, store, and publisher interfaces
- [ ] Typed block and commit-batch models

## Phase 4 — Runtime identity

- [ ] Strict reviewed release-manifest loader
- [ ] Arch v0.6.7 fixture parity
- [ ] Public-network override rejection

## Phase 5 — Arch datasource

- [ ] Typed read-only JSON-RPC client
- [ ] Complete-block validation
- [ ] Gap-free single-cursor polling
- [ ] Bounded retry and cancellation

## Phase 6 — Durable state

- [ ] PostgreSQL migrations
- [ ] Canonical blocks and events
- [ ] Atomic checkpoints and transactional outbox
- [ ] Claim leases and crash-window tests

## Phase 7 — Bump adapter

- [ ] Reviewed protocol fixtures
- [ ] Initialize/create/buy/sell/graduate decoding
- [ ] Nested instruction paths and fuzz tests
- [ ] Versioned historical compatibility

## Phase 8 — Lifecycle correctness

- [ ] Common-ancestor discovery
- [ ] Rollback and reapply transitions
- [ ] Deterministic replay and conflict rejection

## Phase 9 — Delivery and runtime

- [ ] Recoverable Redis publisher
- [ ] Bounded pipeline and graceful shutdown
- [ ] Liveness, readiness, metrics, and safe logs

## Phase 10 — Adoption proof

- [ ] Containerized local integration stack
- [ ] Deterministic parity harness
- [ ] Separate Bump shadow-mode integration
- [ ] Performance/security review and immutable release evidence
- [ ] Explicit human gate before any production cutover
