## Context

The existing collector already provides the correct outer seam: `collector/mongodb_collector.go` obtains one cached MongoDB client and dispatches by node type; `collector/mongod/server_status.go` issues `admin.serverStatus`; typed child structures emit explicit Prometheus descriptors. The current mongod model covers an older subset and must remain intact because deployed dashboards consume its metric names.

The normalized production MongoDB 5.0.34 primary and secondary fixtures define the supported response union. They contain the same 47 top-level sections but differ in field presence and operational values. The primary additionally demonstrates role/election, flow-control, sharding-migration, and command-outcome signals. `serverStatus` contains nested maps whose keys can be unbounded across MongoDB versions; generic BSON flattening would violate Prometheus cardinality constraints.

## Goals / Non-Goals

**Goals:**
- One binary and one mongod `serverStatus` command per scrape for primary and secondary nodes.
- Preserve existing descriptors and collection behavior exactly.
- Make the MongoDB 5.0 primary/secondary fixture union a finite, testable coverage contract.
- Add modern metrics through typed optional fields and static descriptors with bounded labels.
- Permit a staged rollback of only new coverage.

**Non-Goals:**
- Generic exporting of all future or unknown MongoDB fields.
- Exporting unbounded identity or topology strings as labels.
- Changing existing metric names, labels, types, or dashboard contracts.
- Changing the separately configured mongos sharding-status collector.
- Supporting MongoDB versions outside the declared fixture contract in the initial change.

## Decisions

### Preserve the acquisition seam and add a decoded-document seam

`GetServerStatus` remains the sole command issuer. It SHALL obtain raw BSON once, decode the existing `mongod.ServerStatus` model unchanged, and retain the same `bson.Raw` for section-local decoding by a new `collector/serverstatusv5` module. Existing model failures retain existing scrape semantics.

The new module receives `bson.Raw`, never `*mongo.Client`. It decodes reviewed sections independently and returns an emission report containing family-local decode errors. A malformed modern section SHALL suppress only that family, increment a bounded diagnostic counter, and leave all legacy and successfully decoded modern families collectible. This is a deep module: it hides BSON optionality, units, descriptor construction, and cardinality policy behind one input and report interface.

### Decode the primary/secondary union; emit by presence

The new root model is a typed union of both fixtures. Optional documents and fields use presence-aware section lookup; BSON null and absence have the same chosen external meaning: neither emits a metric. Optional scalar wrappers preserve present zero and false independently from absence. The module shall not promise to distinguish null from absence after decoding.

Earlier `isMaster` classification continues to select mongod versus mongos collection, but it SHALL NOT select a primary or secondary server-status schema. A role can change between commands; presence in the collected response is authoritative.

### Add a finite coverage contract

Add a version-controlled contract file under `collector/serverstatusv5/` that classifies every fixture leaf exactly once:

- `legacy_metric`
- `new_metric`
- `metadata_omitted`
- `dynamic_aggregated`
- `dynamic_omitted`
- `structural`

New-metric entries define BSON path, static metric family, Prometheus type, units/conversion, labels, and help text. Omitted entries record the reason. The contract is test input and review evidence, not runtime configuration; the binary uses compiled typed Go and static descriptors.

### Keep legacy metrics; isolate modern families

Existing descriptors remain authoritative for their source fields. New families use `mongodb_server_status_...` to avoid collisions and ambiguous source semantics. A test inventories existing descriptors and rejects collisions.

Metric conversion rules are explicit:

| Source | Metric representation |
|---|---|
| Monotonic count | Counter with `_total` |
| Current count, configured rate, limit | Gauge |
| Milliseconds / microseconds / nanoseconds | Seconds, divided explicitly by 1e3 / 1e6 / 1e9 |
| Boolean | Gauge 0 or 1 |
| Date | Unix seconds gauge when operationally useful |
| Enum | Allowlisted one-hot gauge only |
| Array | Count/aggregate only when useful |

### Bound labels and dynamic maps

New labels are static, reviewed dimensions only: for example `direction`, `outcome`, `state`, or a fixed command/operator/stage vocabulary. Runtime host names, database names, collection names, shard names, arbitrary command keys, and arbitrary WiredTiger paths SHALL NOT become labels.

For dynamic map sections, the contract chooses one of three explicit policies: fixed allowlist; a bounded `other` aggregate; or omission. The initial implementation prioritizes fixed fields and operationally meaningful, bounded map dimensions.

### Stage section modules in priority order

Implement typed collectors by section, all supplied by the shared document:

1. Primary role/election and primary-only services.
2. Flow control.
3. Sharding statistics and move-chunk outcomes.
4. Command outcomes plus bounded aggregation/operator/query/cursor extensions.
5. Logical sessions, transactions, read concern, and oplog truncation.
6. Network/service-executor and connection extensions.
7. Storage-engine, remaining bounded WiredTiger, and replication extensions.

This ordering exposes primary write pressure and migration failures early while preserving one coherent data model.

### Fixtures are tracked reconstructed validation artifacts

The current reference JSON files are lossy analysis views and contain production topology and certificate identity data. The original raw BSON captures are unavailable. A versioned fixture tool SHALL sanitize those normalized snapshots and deterministically reconstruct canonical Extended JSON for each role using reviewed default and path-specific BSON type assignments. The reconstructed fixtures preserve observed document shape and values for coverage and scrape tests, but the project SHALL NOT claim that their BSON types reproduce the unavailable source documents.

Do not commit the current references verbatim. The tracked fixture manifest records role, MongoDB version, capture time when known, redaction version, reconstruction version, normalization version, the digest of each normalized source, and the digest of each checked-in reconstructed fixture. Focused synthetic BSON documents, independent of the reconstructed production snapshots, test `$numberLong`, `$numberInt`, `$numberDouble`, `$date`, `$timestamp`, `$oid`, `$binary`, null, and boolean decoding semantics.

Use the fixture set at four levels:

1. **Schema audit:** traverse all scalar normalized paths; join each to exactly one coverage-contract entry; fail on uncovered, duplicate, or stale paths.
2. **Decode audit:** decode reconstructed canonical Extended JSON through the same BSON path used in production; assert every contract path classified as exported reaches its typed field. Test numeric, temporal, binary, boolean, and nullable semantics with source-independent focused BSON fixtures.
3. **Scrape audit:** run the emitter against each decoded fixture through fresh Prometheus registries; compare static sample sets to reviewed golden output. Assert names, types, labels, conversions, explicit zeroes, omissions, and contract series ceilings.
4. **Transition audit:** use one module/collector instance and one pedantic registry for primary→secondary and secondary→primary sequences; require both gathers to succeed and role-absent samples to disappear.

Golden output is appropriate for static sample sets. Focused tests own conversion, family-local decode failure, descriptor/help/label consistency, collision, and cardinality-policy failures.

### Deployment baseline and release matrix

Before collector work, reconcile this source checkout with the deployed MongoDB Go driver 1.8.6 CA-rotation build: update the authoritative dependency, lock, and vendor records together, then prove TLS compatibility. Define the supported compiler and release OS/architecture matrix explicitly. Update release configuration for every published architecture; test each artifact on compatible native hardware or an emulator. The Nix shell is development support, not proof of legacy Go compiler compatibility; if source compatibility with Go 1.13/1.14 is required, run that compiler in the release test matrix.

### Controlled rollout
The new coverage flag defaults enabled after release and disables only the modern module. Legacy server-status collection continues unchanged. Canary promotion requires successful direct comparison with both node roles and compliance with contract-defined series and scrape-duration budgets.

## Risks / Trade-offs

- The observed fixture union defines complete classification, not automatic metric export. The reviewed contract must list required emitted mappings, omissions, finite label vocabularies, and primary/secondary/union series ceilings before implementation starts.
- Prometheus represents numbers as float64; large MongoDB int64 counters may lose integer precision above `2^53`. Preserve existing semantics, document the limitation, and use counters for monotonic values.
- Typed structs can lag new MongoDB releases. Unknown fields are deliberately ignored; adding a MongoDB version requires an explicit contract/fixture expansion.
- Section-local decoding from retained BSON avoids a second MongoDB command but must be measured against the existing decode; do not optimize before scrape-duration evidence.
- Enabling modern metrics increases scrape payload and series count. Contract ceilings are enforced in tests; canaries confirm the approved budget rather than defining it after deployment.
