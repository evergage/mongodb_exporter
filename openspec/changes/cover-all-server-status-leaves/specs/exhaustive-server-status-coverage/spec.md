## Purpose

Define exact, reviewable ownership for every scalar leaf observed in the sanitized MongoDB 5.0.34 primary and secondary `serverStatus` fixture union, while preserving legacy metrics and preventing duplicate or unbounded output.

## ADDED Requirements

### Requirement: Complete fixture-union disposition

The coverage contract SHALL contain exactly one explicit disposition for every scalar leaf in the union of the supported primary and secondary fixtures. A disposition SHALL be one of: existing legacy metric, modern metric, or reviewed drop with rationale.

No leaf SHALL remain unclassified, implicitly covered by a section prefix, or owned by both collectors.

#### Scenario: Coverage audit
- **WHEN** the fixture-union coverage audit runs
- **THEN** it SHALL report 1,581 uniquely disposed leaves and zero unmapped or duplicate leaves

#### Scenario: Fixture shape changes
- **WHEN** a supported fixture gains, loses, or renames a scalar leaf
- **THEN** the audit SHALL fail until the contract and reviewed inventory are updated

### Requirement: Observable legacy ownership

A legacy ownership claim SHALL be accepted only when changing that source leaf while preserving its BSON type changes an emitted legacy Prometheus sample attributable to that leaf.

Broad section-prefix claims SHALL NOT constitute evidence of legacy coverage.

#### Scenario: False legacy claim
- **WHEN** a leaf classified as legacy is mutated without changing any legacy sample
- **THEN** the audit SHALL fail that ownership claim

### Requirement: Legacy compatibility and duplicate resolution

Existing legacy metric families SHALL remain unchanged. For each source leaf currently consumed by both legacy and modern collectors, the legacy metric SHALL remain the sole owner and the duplicate modern mapping SHALL be removed.

#### Scenario: Existing duplicate
- **WHEN** the contract processes one of the four cursor or six WiredTiger concurrent-transaction duplicate leaves
- **THEN** only its existing legacy metric SHALL be emitted

### Requirement: Counter-first numeric mapping

Each previously unmapped numeric or boolean leaf SHALL receive a proposed modern mapping unless it belongs to an approved manual-review family.

A cumulative value that is monotonic during the lifetime of its MongoDB process or subsystem SHALL be a Prometheus counter. An instantaneous, configured, state, capacity, size, boolean, timestamp, percentage, rate, or generation value SHALL be a gauge.

Counter family names SHALL end in `_total`.

#### Scenario: Cumulative event value
- **WHEN** an observed source field counts events such as evictions, failures, reads, writes, commits, or retries
- **THEN** its generated mapping SHALL use a counter

#### Scenario: Current sampled value
- **WHEN** an observed source field reports current bytes, active work, available capacity, an open population, a limit, or a state
- **THEN** its generated mapping SHALL use a gauge

### Requirement: Raw units for new exhaustive mappings

New exhaustive mappings SHALL preserve the units returned by MongoDB and SHALL use conversion factor `1`.

Microseconds and milliseconds SHALL remain microseconds and milliseconds. Bytes, counts, percentages, rates, timestamps, and logical timestamp components SHALL not be rescaled.

Existing legacy and modern mappings SHALL retain their current units and conversions for compatibility.

#### Scenario: Raw duration
- **WHEN** a newly mapped field is reported by MongoDB in microseconds
- **THEN** the emitted family SHALL identify microseconds and SHALL emit the unscaled source value

### Requirement: Bounded dynamic and histogram mappings

Observed command, aggregation-stage, and operator-counter leaves SHALL use grouped metric families with fixture-derived finite label vocabularies. Internal observed keys SHALL remain in those vocabularies.

Unknown runtime keys SHALL emit no data series and SHALL increment a bounded diagnostic counter labeled only by dynamic family. Unknown keys SHALL NOT appear in metric names or labels.

Observed histogram bucket counts SHALL use grouped counters with one numeric lower-bound label in the source base unit. Boundary source leaves SHALL supply labels for paired count leaves and SHALL remain explicit contract dispositions. The exporter SHALL NOT synthesize histogram sums or cumulative `le` buckets.

Final primary, secondary, union, and per-family series ceilings SHALL be derived from the approved output and confirmed by canary review.

#### Scenario: Unknown dynamic key
- **WHEN** a scrape includes a dynamic key absent from the reviewed fixture vocabulary
- **THEN** the exporter SHALL emit no data sample for that key and SHALL increment the family-only unknown-key diagnostic

#### Scenario: Fixed histogram range
- **WHEN** a reviewed raw bucket count is present
- **THEN** the exporter SHALL emit its unscaled cumulative count with one numeric lower-bound label in the documented source unit

### Requirement: Reviewed string dispositions

Every string, identity, enum, timestamp-text, and opaque-text fixture leaf SHALL use the disposition recorded in `string-leaf-decisions.md`.

Reviewed drops SHALL emit no sample and SHALL retain their rationale in the contract. The oplog-truncation processing method SHALL be emitted only when it belongs to the reviewed finite method vocabulary; unknown values SHALL emit no data sample and SHALL increment a bounded diagnostic. Replication last-write dates SHALL be emitted as raw BSON DateTime millisecond gauges.

#### Scenario: Sensitive identity field
- **WHEN** a reviewed hostname, connection string, certificate subject, ObjectID, or replica-set identity leaf is present
- **THEN** the exporter SHALL emit no sample or raw label for that leaf

#### Scenario: Oplog processing method
- **WHEN** `oplogTruncation.processingMethod` is present
- **THEN** the exporter SHALL emit the reviewed method value as the method label on the approved info family

#### Scenario: Replication write date
- **WHEN** a supported replication write-date BSON value is present
- **THEN** the exporter SHALL emit its native milliseconds-since-epoch value with no conversion

### Requirement: Anchored WiredTiger snapshot timestamps

The two WiredTiger majority-snapshot timestamp strings SHALL produce separate Unix-second and logical-increment gauges.

The parser SHALL recognize MongoDB 5's `MMM DD HH:mm:ss:increment` representation, infer missing year and timezone using the response `localTime` and available logical timestamps, and accept a result only when it passes documented plausibility and uniqueness checks.

A parse failure or ambiguous result SHALL emit `0` for both the Unix-second and increment gauges.

#### Scenario: Valid snapshot timestamp
- **WHEN** the timestamp string can be uniquely anchored to the current response
- **THEN** the exporter SHALL emit the inferred Unix seconds and the parsed increment as separate gauges

#### Scenario: Invalid or ambiguous snapshot timestamp
- **WHEN** parsing fails or more than one plausible absolute instant remains
- **THEN** both gauges SHALL emit `0` and the scrape SHALL continue

### Requirement: Generated mapping evidence

Each generated mapping SHALL record its exact source path, owning collector, Prometheus family, type, raw unit, labels, role presence, decision source, and monotonicity evidence for counters.

#### Scenario: Mapping review
- **WHEN** a reviewer inspects a generated numeric mapping
- **THEN** the mapping SHALL state whether its decision came from an explicit rule, MongoDB or WiredTiger documentation, or a user-approved manual decision

### Requirement: Deployed exporter compatibility and Prometheus ingestion

The candidate SHALL derive from the upstream source tree used for the `0.11.2-evg1` exporter pinned in Evergage Operations, retaining MongoDB Go driver v1.8.6 and all existing collector defaults. For an identical MongoDB response, every legacy MongoDB-derived metric name, HELP, TYPE, label name and value, and sample value SHALL be identical to the deployed exporter. Existing process/runtime metric families and label identities SHALL also be preserved, though their sampled values vary with process state. Only reviewed `mongodb_server_status_*` metric families MAY be added to mongod output; mongos output SHALL remain compatible with the existing `--suppress.collectshardingstatus` setting.

The version-bearing `mongodb_exporter_build_info` sample MAY have different release version, branch, revision, and build-user label values. Its family name, HELP, TYPE, label names, numeric value `1`, and Go-version label SHALL be preserved. Release builds SHALL use the production Go 1.24.2 toolchain so `go_info{version="go1.24.2"}` retains its exact series identity.

All exposed legacy and modern samples SHALL be valid Prometheus text exposition and SHALL be accepted by a running Prometheus scraper with matching metric identity and type. A binary replacement at the same job, target, and port SHALL resume the existing legacy time series with the same target and metric labels while ingesting the new series. A bounded restart scrape gap is permitted; a changed series identity is not.

#### Scenario: Deployed binary comparison
- **WHEN** the exact checksum-pinned production `0.11.2-evg1` artifact and the candidate collect an identical MongoDB 5 response with the production collector flags
- **THEN** every original metric family, HELP, TYPE, label set, and sample value SHALL match and candidate additions SHALL use reviewed modern families

#### Scenario: Prometheus target replacement
- **WHEN** Prometheus has scraped the deployed exporter and then the candidate on one unchanged production-shaped target
- **THEN** previously present legacy series SHALL resume under their original labels and new series SHALL be queryable with valid types and finite values, with both exporter health and target health successful
