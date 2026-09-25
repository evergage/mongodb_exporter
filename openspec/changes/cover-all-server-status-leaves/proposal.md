## Why

The current MongoDB 5 `serverStatus` change classifies all fixture leaves, but actual emission covers only 392 leaves uniquely, duplicates 10 leaves, and leaves 1,189 leaves without a metric. Broad prefix claims—especially for WiredTiger—mistake classification for real legacy coverage and prevent reviewers from seeing the gap.

## What Changes

- Replace broad legacy-coverage claims with an exact, behaviorally verified source-leaf inventory for the 1,581-leaf primary/secondary fixture union.
- Preserve 134 legacy-only and 248 modern-only mappings unchanged.
- Remove 10 duplicate modern mappings and retain their existing legacy metrics.
- Generate modern mappings for unmapped numeric and boolean leaves using a counter-first policy: cumulative monotonic values become counters; instantaneous, configured, state, boolean, timestamp, size, capacity, and generation values become gauges.
- Preserve raw MongoDB units and use conversion factor `1` for newly generated mappings.
- Use fixture-bounded grouped families for dynamic keys, fixed numeric lower-bound counter labels for raw histogram buckets, a family-only diagnostic for unknown keys, and derived series ceilings reviewed after canary.
- Apply the reviewed dispositions in `string-leaf-decisions.md`: most identity and opaque text fields are explicitly dropped; oplog-truncation processing method is exposed as a label; replication dates retain raw BSON milliseconds; WiredTiger majority-snapshot timestamps produce anchored Unix-second and separate increment gauges with zero sentinels on parse failure.
- Strengthen contract verification so every fixture leaf has exactly one explicit disposition and every legacy or modern ownership claim corresponds to observable emitted behavior.

## Capabilities

### New Capabilities

- `exhaustive-server-status-coverage`: Exact source-leaf ownership, generated numeric mapping policy, reviewed nonnumeric dispositions, duplicate prevention, and observable verification for the complete MongoDB 5.0.34 primary/secondary `serverStatus` fixture union.

### Modified Capabilities

- None. The current branch does not retain a repository OpenSpec capability to modify; this change supersedes the incomplete coverage assumptions in the active implementation contract.

## Impact

- Production baseline: the pinned `0.11.2-evg1` exporter built with MongoDB Go driver v1.8.6; this branch is based on the upstream `ops/W-19159693-update-mongo-driver` source tree, not the older `v0.11.2-evg` driver v1.3.2 tag.
- Affected runtime code: `collector/serverstatusv5`, its generated definitions, and contract generator.
- Affected compatibility surface: additive modern metric families only; existing legacy families remain untouched.
- Affected tests: fixture-union audit, exact legacy ownership probe, generated mapping checks, duplicate audit, goldens, cardinality, role transitions, and malformed-family isolation.
- Expected series count increases substantially; exact primary, secondary, union, and per-family ceilings are derived from generated output and confirmed by canary review.
- No change to the production MongoDB driver, vendor tree, multi-architecture archive layout, or mongos metrics. The Nix development shell supports local MongoDB 5 and Prometheus; the inherited Makefile's linker flag quoting is corrected so the candidate retains deployed build-info version labels.
- Planning evidence: `leaf-inventory.md`, `classification-policy.md`, `numeric-mapping-proposals.md`, `dynamic-vocabulary-review.md`, and `string-leaf-decisions.md`.
