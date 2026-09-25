# Metric classification policy

## Scope

Apply this policy to the 1,581 scalar leaves in `leaf-inventory.md`. The end state assigns every fixture-union leaf exactly one explicit disposition: legacy metric, modern metric, or user-reviewed drop with rationale. Unmapped and duplicate buckets must both reach zero.

## Existing coverage

- Retain all 134 legacy-only mappings unchanged.
- Retain all 248 modern-only mappings unchanged.
- For the 10 duplicate legacy/modern mappings, keep the legacy metric as the owner and remove the modern mapping. Legacy compatibility takes precedence over the new namespace.

## Numeric and boolean leaves

Generate proposed mappings for all currently unmapped numeric and boolean leaves.

### Counter

Prefer a Prometheus counter when the MongoDB value is cumulative and monotonic for the life of the process or subsystem. Typical evidence includes:

- event or operation totals;
- failures, successes, attempts, calls, reads, writes, commits, aborts, evictions, and retries accumulated over time;
- cumulative bytes or records processed;
- cumulative elapsed time spent in an operation;
- names documented as `total`, `count`, `num`, `operations`, or equivalent, unless documentation identifies the value as a current population or configured limit.

Counter names must end in `_total`. A process restart may reset the source value and is valid Prometheus counter behavior.

### Gauge

Use a Prometheus gauge when the MongoDB value is an instantaneous sample or can move both upward and downward. Typical evidence includes:

- current, active, available, open, queued, running, pinned, dirty, used, free, resident, or allocated values;
- configured limits, capacities, rates, percentages, scores, and sizes;
- booleans and state codes;
- timestamps and logical clock positions;
- the latest, minimum, maximum, or most recent observation when it is not itself a cumulative total;
- histogram bucket populations captured as a snapshot, pending manual histogram policy.

### Units and conversions

Preserve the raw units exposed by `serverStatus`.

- Do not convert microseconds or milliseconds to seconds.
- Do not rescale bytes, counts, rates, percentages, timestamps, or logical clock components.
- Encode the raw unit in the metric family name where MongoDB exposes it in the field name or documentation.
- Set the runtime conversion factor to `1` for newly generated mappings.

This policy applies to newly added exhaustive mappings. Existing legacy-only and modern-only mappings remain unchanged for compatibility.

## Manual-review buckets

Do not auto-finalize these categories:

1. All 31 string, identity, enum, timestamp-text, and opaque-text leaves. Review each leaf individually.
2. Dynamic map keys, including commands, query shapes, aggregation stages, and operators. Review family shape and bounded labels.
3. Histogram bucket fields. Review whether to expose buckets as Prometheus histograms, fixed labeled counters/gauges, or another bounded representation.
4. Series ceilings. Derive ceilings after all family and label decisions are complete.
5. Numeric leaves whose monotonicity remains ambiguous after source-name and documentation research.

## Generated mapping requirements

Each proposed mapping must record:

- exact `db.serverStatus()` source path;
- owning collector (`legacy` or `modern`);
- Prometheus family;
- counter or gauge;
- raw unit;
- fixed and value-derived labels;
- evidence for monotonicity when classified as a counter;
- primary, secondary, or shared presence;
- whether the decision was automatic, researched, or manually approved.

## Verification invariant

Tests must derive the fixture-union leaf set and prove:

- every leaf has exactly one owner;
- every legacy ownership claim changes an actual emitted legacy sample when the source leaf changes;
- every modern ownership claim emits the documented family, value, and labels;
- no leaf is owned by both collectors;
- no unreviewed omission remains; every dropped leaf has an explicit user-approved rationale;
- unknown runtime fields cannot create metric names or unbounded label values.
