## Why

The deployed exporter decodes only a legacy subset of MongoDB 5.0.34 `serverStatus`. Production primary and secondary snapshots expose critical replication, flow-control, sharding, command, transaction, security, and storage signals that currently have no Prometheus representation.

## What Changes

- Add complete classification—not automatic export—of the observed MongoDB 5.0 primary/secondary `serverStatus` fixture union; emit only the reviewed, bounded metric subset.
- Use one shared `serverStatus` fetch and one typed union model for both primary and secondary responses; metric emission depends on field presence rather than a role-specific decoder.
- Preserve all existing metric names and collector behavior, including the deployed MongoDB Go driver 1.8.6 CA-rotation compatibility; add new metric families in a separate stable namespace.
- Reconstruct redacted canonical Extended JSON test fixtures from the available normalized primary and secondary snapshots, record the normalized-source provenance and reconstruction rules, and maintain a complete field-classification contract without claiming original BSON type fidelity.
- Add deterministic BSON decoding, metric, cardinality, and same-registry primary/secondary role-transition tests.
- Add an opt-out switch for the new coverage during staged deployment.

## Capabilities

### New Capabilities
- `mongodb-server-status-coverage`: Collects and exports a bounded, documented MongoDB 5.0 `serverStatus` metric contract for both primary and secondary mongod nodes.

### Modified Capabilities
- None.

## Impact

- `collector/mongod/server_status.go` and the mongod server-status collector packages.
- New typed server-status coverage module, descriptors, fixtures, and tests.
- Exporter CLI configuration, build artifacts, dashboards, and alerts that adopt the new metric families.
- Existing legacy metric names and mongos sharding-status behavior remain compatible.
