## 1. Exact Coverage Model

- [x] 1.1 Implement a fixture leaf walker that preserves literal object keys, array indices, BSON type, and primary/secondary presence; verify the sanitized fixture union contains exactly 1,581 unique scalar leaves, including dotted and empty-string keys.
- [x] 1.2 Implement the behavioral legacy-ownership probe that mutates one leaf at a time with the same BSON type and compares canonicalized legacy samples; verify all 144 final legacy-owned leaves—including the 10 duplicate resolutions—change their documented legacy family and no other leaf is claimed by a prefix.
- [x] 1.3 Replace prefix classifications with an explicit generated disposition for every leaf: `legacy`, `modern`, or reviewed `drop`; verify generation fails for missing, duplicate, stale, or unreviewed entries.
- [x] 1.4 Import every decision and rationale from `string-leaf-decisions.md`; verify all 31 nonnumeric leaves have no `Pending` or `Research` disposition.

## 2. Generated Numeric Definitions

- [x] 2.1 Implement counter/gauge generation from `classification-policy.md` with raw units and conversion factor `1`; verify counter names end in `_total` and snapshot/state mappings are gauges.
- [x] 2.2 Encode MongoDB and WiredTiger documentation-backed exceptions, including connection exhaust state, TLS protocol counters, WiredTiger data-handle size, checkpoint generation, prepared transactions, tiered retention, and compact timeout; add focused type-policy tests for each exception.
- [x] 2.3 Generate the review-corrected 411 counter and 248 gauge proposals into the expanded contract and static Go definitions; verify each source path, type, raw unit, family, role presence, and decision evidence matches `numeric-mapping-proposals.md`.
- [x] 2.4 Emit pre-split BSON path segments and stable label order in generated Go; benchmark or allocate-test a scrape to verify exhaustive coverage does not add per-definition `strings.Split` allocations.

## 3. Bounded Dynamic Families

- [x] 3.1 Generate the 60-command vocabulary and grouped command families with fixed `command`, `outcome`, and approved feature labels; verify every observed command leaf is owned exactly once, including internal commands.
- [x] 3.2 Generate the 49-stage aggregation vocabulary and grouped stage counter; verify every observed public and internal stage emits with an allowlisted label.
- [x] 3.3 Generate all 210 observed operators across `expressions`, `groupAccumulators`, `match`, and `windowAccumulators`; verify exact fixture enumeration and category/operator label bounds against `dynamic-vocabulary-review.md`.
- [x] 3.4 Add `mongodb_server_status_unknown_keys_total{family}` and ignore unknown dynamic keys; verify injected command, stage, operator, and category keys create no data family or label value and increment only the bounded family diagnostic.

## 4. Fixed Histogram Range Counters

- [x] 4.1 Map query multi-planner histogram count/lower-bound pairs into grouped counters with one numeric lower-bound label in microseconds, plan count, work count, or read count; verify lower-bound leaves are consumed as labels and not emitted separately.
- [x] 4.2 Map resharding latency ranges into grouped counters with `lower_bound_millis` values `0`, `10`, `100`, `1000`, and `10000`; verify source bucket counts remain unscaled.
- [x] 4.3 Parse WiredTiger histogram descriptions into numeric lower-bound labels in their source units; verify every observed bound parses, ambiguous descriptions fail generation, and no synthetic histogram sum or cumulative `le` bucket is produced.
- [x] 4.4 Verify all 151 histogram count/boundary leaves have exactly one contract disposition and every emitted bucket label is a single numeric value.

## 5. Reviewed String and Duplicate Behavior

- [x] 5.1 Remove the four cursor and six WiredTiger concurrent-transaction modern definitions; verify only the unchanged legacy families consume those leaves.
- [x] 5.2 Add the oplog-truncation processing-method info family with the raw `method` label; verify present, absent, and changed method values.
- [x] 5.3 Add raw BSON millisecond gauges for replication `lastWriteDate` and `majorityWriteDate`; verify conversion factor `1` and absent-field omission.
- [x] 5.4 Implement the MongoDB 5 `MMM DD HH:mm:ss:increment` parser and anchored candidate selection using `localTime` plus response logical timestamps; verify year boundaries, timezone offsets, valid primary/secondary fixtures, ambiguity, malformed input, and zero sentinels.
- [x] 5.5 Emit separate Unix-second and increment gauges for latest and oldest WiredTiger majority-snapshot timestamps; verify both gauges become zero when anchoring fails and other families continue.
- [x] 5.6 Verify every reviewed drop emits no sample and exposes no production identity or opaque string label while retaining its explicit contract rationale.

## 6. Exhaustive Verification and Canary

- [x] 6.1 Extend contract tests to prove all 1,581 leaves have exactly one disposition, all ownership claims are observable, no source leaf is consumed twice, and no unreviewed omission remains.
- [x] 6.2 Regenerate primary and secondary Prometheus goldens; verify family names, counter/gauge types, raw units, labels, explicit zeroes, role-specific absence, and malformed-family isolation.
- [x] 6.3 Derive exact per-family, primary, secondary, and union series ceilings from generated output; review and record those ceilings, then verify any additional family or label value fails policy tests.
- [x] 6.4 Run generation checks and the serverstatusv5 race-enabled suite; verify deterministic fixture reconstruction, provenance digests, legacy compatibility, role transitions, unknown-key diagnostics, and zero stale samples.
- [x] 6.5 Build the candidate on the deployed `0.11.2-evg1` dependency baseline (MongoDB Go driver v1.8.6, upstream ops branch); compare the exact pinned production binary and candidate against MongoDB 5 standalone, primary, secondary, config/shard and mongos nodes, including direct `serverStatus` values, scrape latency, and opt-out behavior.
- [x] 6.6 Update the changelog and PR description with final family/series counts and review boundary; verify the diff from upstream ops baseline retains driver, vendor, collector defaults, and release archive layout. Only the feature, local Nix development, OpenSpec, docs, candidate version, and linker flag correction that preserves deployed build-info metadata are approved.


## 7. Review Remediation

- [x] 7.1 Reconstruct all logical-time pairs as wire-level BSON Timestamp values and generate gauge seconds/increment transforms.
- [x] 7.2 Bound oplog processing methods and drop signing-key and process identifiers.
- [x] 7.3 Diagnose unknown empty dynamic objects without exposing their keys.
- [x] 7.4 Enforce fixture-derived path segments, roles, and histogram consumer links.
- [x] 7.5 Run generated projection freshness under normal Go tests.
- [x] 7.6 Probe every declared role for each legacy ownership claim.
- [x] 7.7 Remove automatic golden replacement and require explicit semantic review.

## 8. Production Baseline and Ingestion Gate

- [x] 8.1 Verify the pinned `0.11.2-evg1` archive checksums and build provenance against Evergage Operations; exercise that exact Linux artifact alongside a Linux build of this branch. Record unchanged driver v1.8.6, candidate Go 1.24.2 and stamped version/build labels, `--suppress.collectshardingstatus`, URI timeouts, and collector defaults.
- [x] 8.2 For standalone, PRIMARY, SECONDARY, and sharded MongoDB 5 roles, compare all legacy HELP/TYPE/label signatures and same-response values, then show every candidate sample survives a real Prometheus scrape with the declared type and finite value; require only reviewed modern families as additions.
- [x] 8.3 Replace the exporter behind one stable production-shaped Prometheus target and verify historical series label identity, resumed legacy ingestion, new-series ingestion, and error/health alerts. Exercise `--suppress.collectserverstatusv5` and the monitoring-user X.509 URI. Treat externally owned alert expressions and vmagent remote-write as explicit test-environment gates if unavailable locally.
- [x] 8.4 Record reports and honest version/corpus/cardinality limits outside the repository; audit that no disposable harness, private key, fixture capture, or production configuration change enters the branch.
