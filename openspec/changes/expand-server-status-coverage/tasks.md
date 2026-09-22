## 1. Evidence Fixtures and Coverage Contract

- [x] 1.1 Deterministically sanitize the available normalized primary and secondary MongoDB 5.0.34 server-status snapshots and reconstruct canonical Extended JSON fixtures using reviewed type assignments; add tracked metadata with role, version, known capture time, source/reconstructed digests, and redaction/reconstruction/normalization versions; verify no production topology or certificate identity remains and record that original BSON type fidelity is unavailable.
- [x] 1.2 Implement the versioned reconstruction and normalization tool plus fixture-integrity test; verify normalized sources deterministically produce the reconstructed canonical Extended JSON fixtures and normalized path views, while source-independent focused BSON fixtures test each supported BSON representation.
- [x] 1.3 Add the versioned coverage contract and classify every scalar leaf in the primary/secondary normalized fixture union exactly once; verify the audit fails for missing, duplicate, or stale classifications.
- [x] 1.4 Generate or maintain an inventory of existing server-status descriptors; verify each proposed new descriptor avoids collisions with legacy metric names.

- [x] 1.5 Define the mandatory emitted mappings, explicit omissions, finite dynamic-key vocabularies, and numeric per-family plus primary/secondary/union series ceilings in the reviewed coverage contract; verify contract-derived ceiling tests enforce every limit.
## 2. Shared Collection and Typed Data Model

- [x] 2.1 Reconcile the MongoDB Go driver, dependency locks, and vendor tree with the deployed 1.8.6 CA-rotation-compatible build; verify a TLS connection scenario before collector changes.
- [x] 2.2 Refactor mongod server-status acquisition to retain one raw BSON command response, preserve legacy decode/output, and pass retained BSON to modern section-local decoding; verify one command serves both paths.
- [x] 2.3 Implement presence-aware optional BSON fields with null treated as absent; verify present zero/false, absence, int32, int64, double, date, timestamp, and boolean semantics.
- [x] 2.4 Implement family-local modern decoding and structured emission errors; verify a malformed modern field suppresses only its family, increments the diagnostic counter, and leaves legacy/other modern samples available.

## 3. Priority Primary and Replication Coverage

- [x] 3.1 Add typed role, election, and primary-only-service coverage with bounded state representation; verify primary and secondary fixtures emit only their present samples.
- [x] 3.2 Add flow-control metrics with correct seconds conversion and boolean gauges; verify the primary fixture’s acquisition, rate, and lag fields.
- [x] 3.3 Add sharding-statistics and move-chunk outcome metrics with fixed direction/outcome labels; verify donor/recipient successes, failures, byte totals, and stale-config counters.
- [x] 3.4 Add logical-session-cache, transaction, read-concern, oplog-truncation, and supported replication extensions; verify absent role-specific subdocuments emit no samples.

## 4. Modern Workload and Resource Coverage

- [x] 4.1 Add bounded command, aggregation-stage, operator, query, and cursor metrics according to the coverage contract; verify unknown dynamic keys cannot create metric names or label values.
- [x] 4.2 Add approved network, connection, service-executor, security, and storage-engine fields with documented units and types; verify new series stay within the contract label vocabulary.
- [x] 4.3 Expand selected WiredTiger coverage using typed fixed fields and bounded labels; verify every remaining fixture path is explicitly classified rather than silently unhandled.

## 5. Fixture-Based Metric Verification

- [x] 5.1 Decode reconstructed canonical Extended JSON through the production-equivalent BSON path and assert every exported contract path reaches its typed field; verify fixture shape or reviewed type-assignment changes fail the test.
- [x] 5.2 Collect each decoded fixture through fresh registries and compare reviewed golden samples; verify family names, types, labels, conversions, explicit zeroes, omissions, and contract series ceilings.
- [x] 5.3 Add sparse-document, null, and malformed-value tests; verify absent/null fields emit no sample and invalid modern data isolates to its metric family.
- [x] 5.4 Add primary-to-secondary and secondary-to-primary tests using one long-lived collector/module and one pedantic registry; verify both gathers succeed and absent-role samples disappear without state leakage.
- [x] 5.5 Add cardinality-policy tests for all dynamic map collectors; verify unknown keys are aggregated or omitted according to the contract and cannot add label values.

## 6. Build, Canary, and Adoption

- [x] 6.1 Define the supported compiler, OS, and architecture release matrix; update the Nix shell and release configuration for every published target, including amd64 and arm64, and verify required build/test tools are available.
- [x] 6.2 Run fixture/coverage tests and the full supported project test command under each required Go compiler; verify both reference roles, contract audit, legacy compatibility, and CA-rotation TLS behaviour.
- [x] 6.3 Build each published architecture artifact with existing release metadata and execute it on native hardware or an emulator; verify expected version/build information and changed scrape output.
- [x] 6.4 Canary on one MongoDB 5.0 primary and one secondary, compare output to direct server-status responses, and verify scrape duration and series count against the pre-approved contract ceilings.
- [x] 6.5 Update dashboards and alerts to consume new metric families after parallel observation; verify legacy dashboards continue working until migration completes.
