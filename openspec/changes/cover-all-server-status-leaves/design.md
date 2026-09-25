## Context

See `proposal.md` for motivation. The active collector retains one `bson.Raw` response and sends it through unchanged legacy decoding plus a static modern module. The current contract has 258 explicit modern mappings but uses broad prefix rules for everything else; those rules cannot prove actual legacy output.

The reviewed fixture union has 1,581 scalar leaves. Exact behavioral probing finds 134 legacy-only leaves, 248 modern-only leaves, 10 duplicates, 1,158 unmapped numeric/boolean leaves, and 31 unmapped string/text leaves.

The deployed baseline is the `0.11.2-evg1` exporter pinned in `evergage-operations/cookbooks/cluster/recipes/_mongodb_exporter.rb`. Its verified Linux release archives contain MongoDB Go driver v1.8.6 and report source revision `8bb62b6` with local build changes. Upstream `ops/W-19159693-update-mongo-driver` at `87209cc` preserves that driver's source and release archive layout; this branch descends directly from it. The pinned release binary is the comparison authority. Retain its dependency and collector contracts; correct only inherited linker flag quoting that otherwise strips candidate version metadata with the local Go toolchain.

## Goals / Non-Goals

**Goals:**

- Give every fixture leaf one explicit, testable disposition.
- Preserve every legacy family unchanged and prefer legacy ownership for duplicates.
- Generate consistent counter/gauge mappings for numeric leaves while preserving raw units.
- Make uncertainty visible as a review gate rather than an omission rule.
- Bound runtime names and labels despite broad observed coverage.
- Keep scrape behavior presence-based and role-independent.

**Non-Goals:**

- Export unknown fields or dynamic keys not represented in the reviewed fixture union.
- Change existing legacy or modern units for consistency.
- Change the deployed MongoDB Go driver v1.8.6, vendored code, collector defaults, or production multi-architecture archive layout. The narrow version-linker flag correction retains the deployed build-info behavior.
- Expose production identities, topology strings, certificate subjects, or opaque diagnostic prose when the reviewed disposition is `Drop`.

## Decisions

### Production compatibility boundary

Chef runs `0.11.2-evg1` with `--suppress.collectshardingstatus`, a five-second Mongo connection and server-selection timeout, and a local monitoring-user X.509 URI. Prometheus and vmagent discover port 9216 as job `mongodb` every 60 seconds; target labels include `instance`, `kind`, optional `cluster`, and Chef tags. Preserve existing legacy families, HELP, TYPE, metric labels, and values for the same MongoDB response. Add only reviewed `mongodb_server_status_*` families on mongod; mongos remains unchanged. Validate actual Prometheus ingestion of all emitted series and a replacement behind the same target identity. Concurrent exporters observe different MongoDB responses, so compare volatile counters and dynamic histogram labels using a captured identical BSON response rather than falsely equating independent scrapes.

The build-specific `mongodb_exporter_build_info` sample necessarily changes version, branch, revision, and build-user label values between releases. Its family, HELP, TYPE, label names, and value `1` remain stable. Build with the deployed Go 1.24.2 toolchain so both `mongodb_exporter_build_info{goversion="go1.24.2"}` and `go_info{version="go1.24.2"}` retain their Go-version labels; the Makefile rejects another compiler to prevent this subtle series split.

### Explicit per-leaf dispositions replace prefix claims

The machine-readable contract will contain one entry per fixture-union scalar leaf. Each entry has:

- source path and primary/secondary presence;
- disposition: `legacy`, `modern`, or `drop`;
- decision source and rationale;
- legacy family evidence or modern metric definition;
- review status for manual decisions.

Prefix rules may group entries for authoring, but they cannot satisfy the audit. The generated expanded contract is the reviewed source of truth.

Alternative rejected: retaining broad `wiredTiger.* -> legacy_metric` rules. They classify 529 WiredTiger leaves as legacy even though exact probing finds only 47 legacy-owned WiredTiger leaves, including six duplicate modern mappings.

### Behavioral legacy inventory

A test-only probe mutates one fixture leaf at a time while preserving its BSON type, gathers legacy output, canonicalizes labels, and records which family changes. A `legacy` disposition must match this derived result. The probe handles the existing global `mongodb_version_info` family explicitly.

This validates source semantics rather than merely checking descriptor names.

### Generated numeric mappings

The mapping generator applies `classification-policy.md` to the 1,158 unmapped numeric/boolean leaves:

- 411 counter proposals;
- 248 gauge proposals;
- 346 dynamic-key leaves assigned to bounded grouped-family vocabularies;
- 151 histogram count/boundary leaves assigned to fixed numeric-boundary counter families.

Every source leaf remains explicit even when a boundary leaf supplies a label rather than its own sample.

New exhaustive mappings retain raw values with conversion `1`. Metric names include raw unit terms already present in the source path or documented unit. Existing mappings are not regenerated under this policy.

### WiredTiger source semantics

MongoDB 5.0's vendored WiredTiger `stat_data.py` is the authoritative source for ambiguous WiredTiger statistics. Its default is operation-rate/cumulative semantics; snapshot statistics are identified by descriptions such as `currently`, `in the cache`, limits, sizes, state, or source flags. Path-specific researched exceptions are recorded in `numeric-mapping-proposals.md`.

### Dynamic-key families

Commands, operator counters, and aggregation-stage counters use fixture-derived bounded labels documented in `dynamic-vocabulary-review.md`: 60 commands, 49 aggregation stages, and 210 operators across four operator categories. Internal observed keys remain explicit because the goal is complete fixture coverage.

Unknown runtime keys emit no sample and increment `mongodb_server_status_unknown_keys_total{family="..."}`. The diagnostic label identifies only the bounded dynamic family, never the unknown key.

The expanded contract still contains one entry for every observed leaf, even when multiple leaves feed one labeled family.

### Fixed histogram range counters

Histogram bucket counts use bounded grouped counter families rather than synthesized Prometheus histograms. Each series has one numeric lower-bound label in its raw base unit.

- Query multi-planner `lowerBound` leaves provide the label for their paired count leaves.
- Resharding millisecond ranges use `lower_bound_millis` values `0`, `10`, `100`, `1000`, and `10000`; nonnegative latency makes `0` the normalized lower bound for the source `(-inf, 10)` range.
- WiredTiger textual ranges are parsed into a single numeric lower bound in the unit implied by the source statistic.
- Generation fails when a bound cannot be parsed unambiguously.

No histogram sum or cumulative `le` series is synthesized. Raw bucket counts and units are preserved.

The four cursor and six WiredTiger concurrent-transaction duplicate leaves retain their legacy owners. Their modern definitions are deleted. This prevents double emission while preserving dashboards.

### Reviewed string behavior

`string-leaf-decisions.md` is authoritative for all 31 nonnumeric leaves:

- 28 identity, API-mode, topology, static, process, signing-key, or opaque fields are reviewed drops;
- `oplogTruncation.processingMethod` becomes an info family with a reviewed finite method vocabulary;
- two replication write dates become raw BSON millisecond gauges;
- two WiredTiger snapshot timestamp strings each produce Unix-second and increment gauges.

A reviewed drop remains in the contract with rationale and is not considered unmapped.

### Anchored timestamp parser

MongoDB 5.0 constructs the WiredTiger strings with `Timestamp::toStringPretty()`, which uses `time_t_to_String_short()` plus `:<increment>`. The format is `MMM DD HH:mm:ss:increment`, produced in the mongod process's local timezone and without a year.

The parser will:

1. parse and validate the textual components and unsigned increment;
2. use response `localTime` as the year and wall-clock anchor;
3. enumerate plausible year-boundary and timezone-offset candidates;
4. compare candidates with BSON logical timestamps available in the same response;
5. accept only one uniquely closest candidate inside a documented plausibility bound;
6. reuse the accepted anchor for the oldest timestamp;
7. emit zero seconds and zero increment when parsing, anchoring, or plausibility validation fails.

The increment is always a separate gauge. Zero is an explicit unknown sentinel for these two generated families.

### Generated runtime definitions

The generator emits static Go definitions from approved modern dispositions. Paths are emitted as pre-split segments so scrapes do not repeatedly call `strings.Split`. Fixed labels are emitted in stable order. The runtime does not parse the coverage contract.

### Verification and cardinality

Tests derive:

- the fixture leaf union;
- behavioral legacy ownership;
- modern definitions and observable output;
- reviewed drops;
- duplicate ownership;
- family schemas and bounded labels;
- primary, secondary, and union series counts.

Final series ceilings are written only after dynamic and histogram mappings are approved. Canary acceptance compares both roles against those derived ceilings and the existing scrape-duration ceiling.

## Risks / Trade-offs

- Exhaustive observed coverage creates substantially more series. Derived ceilings and canary evidence are mandatory before merge.
- Some MongoDB/WiredTiger fields may change semantics across versions. The contract targets the two MongoDB 5.0.34 fixtures and ignores unknown runtime keys.
- Counter/gauge classification inferred from descriptive names can be wrong. Documentation-backed exceptions and generated evidence reduce but do not eliminate this risk.
- Anchoring timezone-less snapshot timestamp strings is inherently uncertain. Unique-candidate checks and zero sentinels prevent silently publishing ambiguous dates.
- A 1,581-entry expanded contract is larger than prefix rules, but it is mechanically generated and reviewable through the grouped Markdown inventories. Correctness takes precedence over a smaller policy file.
- Hundreds of static BSON lookups increase scrape work. Pre-split paths, one retained raw response, and canary latency ceilings constrain the cost.
