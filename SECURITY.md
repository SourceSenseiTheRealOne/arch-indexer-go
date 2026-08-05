# Security Policy

## Scope

`arch-indexer-go` is a read-only Arch Network indexing engine. It must not sign or submit transactions, load wallet material, or expose arbitrary JSON-RPC methods.

## Reporting

Report suspected vulnerabilities privately to the repository owner. Do not open a public issue containing exploit details, credentials, private endpoints, or production data.

Include:

- affected commit and component;
- reproducible synthetic/local steps;
- expected and observed behavior;
- impact and proposed mitigation;
- whether the issue involves event loss, ordering, rollback, identity conflict, or secret exposure.

## Secrets

Never commit database credentials, Redis credentials, wallet material, RPC tokens, session state, or raw private payloads. Use environment-driven values and sanitized fixtures.

## Current status

The repository is a foundation scaffold. It does not yet ingest blocks or connect to PostgreSQL/Redis. Production deployment is not approved.
