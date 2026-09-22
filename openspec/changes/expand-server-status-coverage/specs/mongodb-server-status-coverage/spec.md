## Purpose

Expose a bounded, documented MongoDB 5.0 `serverStatus` metric contract from one exporter binary for both primary and secondary mongod nodes while preserving existing exporter output.

## ADDED Requirements

### Requirement: Shared primary and secondary server-status collection

The exporter SHALL collect the MongoDB `serverStatus` response once per mongod or replica-set scrape and SHALL support the union of fields present in the supported primary and secondary MongoDB 5.0 reference responses.

The exporter SHALL determine metric emission from a field's presence in that response, not from a role-specific `serverStatus` decoder. A missing field SHALL emit no sample; a present zero or false value SHALL emit a zero-valued sample.

#### Scenario: Primary-only field is present
- **WHEN** a primary `serverStatus` response contains a supported primary-only field
- **THEN** the exporter SHALL emit that field's documented metric family

#### Scenario: Primary-only field is absent on a secondary
- **WHEN** a secondary `serverStatus` response omits a supported primary-only field
- **THEN** the exporter SHALL emit no sample for that field and SHALL continue collecting all other supported metrics

#### Scenario: Role changes between scrapes
- **WHEN** a node changes role and its subsequent `serverStatus` response changes field presence
- **THEN** the exporter SHALL reflect the new response without retaining samples from absent fields

### Requirement: Stable legacy metric compatibility

The exporter SHALL preserve the names, labels, values, and collection behavior of existing server-status metric families.

New coverage SHALL use documented metric families that do not collide with existing metric names. The exporter SHALL continue its existing mongos sharding-status behavior independently of the new mongod server-status coverage.

#### Scenario: Existing dashboard metric
- **WHEN** a supported legacy server-status field is scraped before and after this change
- **THEN** its existing Prometheus metric name and labels SHALL remain available with equivalent source semantics

### Requirement: Bounded metric contract

The exporter SHALL define a finite, version-controlled classification for every leaf in the supported primary-and-secondary reference response union. Classification SHALL NOT imply export: the reviewed contract SHALL identify the required emitted subset and explicit omission rationale for every remaining leaf.

Each leaf SHALL be classified as an existing metric, a new metric, metadata intentionally omitted, dynamic data intentionally aggregated, dynamic data intentionally omitted, or structural data. Every new metric entry SHALL document its metric family, unit, Prometheus type, conversion, complete finite label vocabulary, and the per-family and total primary/secondary/union series ceilings. The test suite SHALL derive and enforce those ceilings from the contract.

The exporter SHALL NOT derive metric names or label values from arbitrary BSON map keys, host names, database names, collection names, shard names, or other unbounded runtime strings.

#### Scenario: Unknown response field
- **WHEN** MongoDB returns a field outside the documented supported contract
- **THEN** the exporter SHALL ignore that field without creating a new metric or label value

#### Scenario: Dynamic command key
- **WHEN** a supported dynamic server-status map contains an unrecognised command or operator key
- **THEN** the exporter SHALL apply its documented bounded aggregation or omission policy

### Requirement: Modern primary and secondary operational coverage

The exporter SHALL provide documented coverage for the supported contract's primary and secondary operational signals, including role/election state, flow control, sharding statistics, command outcomes, logical session cache activity, transaction state, read concern counters, oplog truncation, and selected modern network, storage, WiredTiger, query, cursor, aggregation, operator, and replication metrics.

Durations SHALL be exported in seconds. Boolean values SHALL be exported as zero-or-one gauges. Date values SHALL be exported as Unix-time gauges when they are operationally useful.

#### Scenario: Primary flow-control pressure
- **WHEN** a primary `serverStatus` response reports flow-control acquisition time and lag state
- **THEN** the exporter SHALL expose documented flow-control metrics with seconds-based duration units

#### Scenario: Sharded migration failures
- **WHEN** a primary response reports supported donor, recipient, or command failure counters
- **THEN** the exporter SHALL expose them with documented bounded outcome and direction labels

### Requirement: Fixture-backed coverage verification

The project SHALL retain redacted reconstructed canonical Extended JSON fixtures for the primary and secondary MongoDB 5.0 server-status documents. The fixtures SHALL be deterministically derived from the available normalized production snapshots using reviewed path-based BSON type assignments. They SHALL preserve the observed document shape and values but SHALL NOT be represented as source-faithful BSON captures.

The fixture manifest SHALL record node role, MongoDB version, capture timestamp when known, redaction version, reconstruction version, normalization version, the digest of each normalized source snapshot, and the digest of each checked-in reconstructed Extended JSON fixture. The project SHALL explicitly record that the original BSON types and raw-capture provenance are unavailable.

The fixture suite SHALL validate the contract at three levels:

- fixture integrity: each reconstructed Extended JSON fixture SHALL identify its normalized source and matching source digest, reconstruction version, and fixture digest;
- document coverage: every scalar leaf in the supported fixture union SHALL have exactly one contract classification;
- observable export: each classified exported leaf SHALL produce its documented Prometheus family, type, unit conversion, bounded labels, and no more than the documented series ceiling.

Tests SHALL verify metric emission for present values, zero values, absent optional fields, descriptor uniqueness, bounded labels, and primary-to-secondary transitions using the same registered collector. Reconstructed canonical Extended JSON SHALL supply the full-document decode and scrape fixtures; dedicated source-independent BSON decoding tests SHALL prove null and each supported numeric/date/timestamp/boolean representation because the original BSON types are unavailable.

#### Scenario: Capture provenance
- **WHEN** a fixture-backed test loads a reconstructed primary or secondary document
- **THEN** the test SHALL verify its normalized source digest, reconstruction version, and reconstructed-fixture digest

#### Scenario: Reference coverage audit
- **WHEN** the fixture coverage verification runs
- **THEN** it SHALL fail if any scalar leaf in either supported reference fixture lacks exactly one documented classification

#### Scenario: Export contract
- **WHEN** a classified source leaf is present in either reference fixture
- **THEN** the test SHALL assert its expected Prometheus descriptor, sample value after documented conversion, and complete label set

#### Scenario: Explicit null
- **WHEN** a supported field has BSON null
- **THEN** the exporter SHALL treat it as absent and SHALL emit no sample

#### Scenario: Sparse document
- **WHEN** a supported nested document is absent from a valid server-status response
- **THEN** the exporter SHALL not fail the scrape and SHALL not emit that document's metrics

### Requirement: Deployment compatibility

The implementation SHALL preserve the deployed MongoDB Go driver 1.8.6 CA-rotation behavior while adding server-status coverage. Release verification SHALL define the supported Go compiler, operating-system, and architecture matrix, build every published artifact for that matrix, and exercise each artifact on a compatible native runner or emulator.

#### Scenario: Deployed TLS connection
- **WHEN** the exporter connects to a deployment using the CA-rotation-compatible TLS configuration
- **THEN** the exporter SHALL retain its existing successful connection behavior while collecting legacy and new metrics

### Requirement: Controlled rollout

The exporter SHALL provide an opt-out configuration switch for the new server-status coverage while retaining legacy collection. The switch SHALL be enabled by default after the feature is released for normal use.

#### Scenario: Coverage opt-out
- **WHEN** the new-coverage switch is disabled
- **THEN** the exporter SHALL continue emitting all legacy metrics and SHALL omit only the new coverage families
