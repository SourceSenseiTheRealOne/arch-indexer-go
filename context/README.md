# Project Context

Sanitized, durable operating context for `arch-indexer-go`.

Load in this order:

1. `stack.md` — pinned stack, runtime, and storage decisions.
2. `QUEUE.md` — pending bounded work.
3. `IN_PROGRESS.md` — currently claimed work.
4. `TASK_TRACKER.md` — verified completed work.
5. `../docs/architecture.md` and `../docs/roadmap.md` — design and delivery phases.

Secrets, credentials, private endpoints, wallet material, production payloads, `.codegraph/`, and raw evidence never belong here.
