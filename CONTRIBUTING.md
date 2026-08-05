# Contributing

## Workflow

1. Start from the latest `development` branch.
2. Create one bounded `feature/*`, `fix/*`, `test/*`, or `docs/*` branch.
3. Use RED → GREEN → REFACTOR for behavioral changes.
4. Run `bash scripts/verify.sh` before requesting review.
5. Open a pull request into `development` and obtain review.
6. Promote only through reviewed pull requests: `development → staging → main`.
7. Do not push directly to environment branches after repository bootstrap.

## Correctness requirements

- Preserve complete-block transaction ordering and canonical numeric instruction paths.
- Advance checkpoints only with atomic canonical persistence.
- Treat Redis delivery as at-least-once and PostgreSQL as authoritative.
- Fail closed on partial blocks, immutable identity conflicts, or unverifiable runtime identity.
- Add synthetic/local fixtures and failure-path tests for every behavioral change.

## Security requirements

- Never add wallet loading, signing, transaction submission, or a generic RPC-call surface.
- Never commit credentials, private endpoints, sessions, raw private payloads, generated CodeGraph state, build output, or local evidence.
- Public-network testing is read-only and must use reviewed release evidence.
- Report vulnerabilities privately as described in `SECURITY.md`.

## Commits

Use concise Conventional Commit subjects without tool or AI attribution. Keep commits independently reviewable.
