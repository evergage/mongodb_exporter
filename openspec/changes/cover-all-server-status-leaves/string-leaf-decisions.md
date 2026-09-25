# String leaf decisions

Each fixture-union string leaf receives an explicit reviewed disposition. `Drop` means no Prometheus sample is emitted, but the source path remains in the contract as a reviewed omission with rationale.

| # | Source leaf | Category | Decision | Rationale |
|---:|---|---|---|---|
| 1 | `db.serverStatus().$clusterTime.signature.hash` | Identity and topology | Drop | Opaque cluster signature hash has no diagnostic value. |
| 2 | `db.serverStatus().$gleStats.electionId` | Identity and topology | Drop | Opaque election ObjectID duplicates more useful role and election counters. |
| 3 | `db.serverStatus().extra_info.note` | Opaque diagnostic text | Drop | Static explanatory prose is not runtime diagnostic state. |
| 4 | `db.serverStatus().host` | Identity and topology | Drop | Prometheus target labels already identify the host; exporting it would duplicate topology identity. |
| 5 | `db.serverStatus().metrics.apiVersions['MongoD Async'][0]` | Finite enum or info | Drop | Internal-client API mode is low-value metadata. |
| 6 | `db.serverStatus().metrics.apiVersions['MongoD Sync'][0]` | Finite enum or info | Drop | Internal-client API mode is low-value metadata. |
| 7 | `db.serverStatus().metrics.apiVersions['MongoDB Shell'][0]` | Finite enum or info | Drop | Client API-mode metadata is dynamic and low-value. |
| 8 | `db.serverStatus().metrics.apiVersions.OplogFetcher[0]` | Finite enum or info | Drop | Internal replication-client API mode is low-value metadata. |
| 9 | `db.serverStatus().metrics.apiVersions[''][0]` | Finite enum or info | Drop | Unnamed API-version client cannot support a stable useful metric. |
| 10 | `db.serverStatus().metrics.apiVersions.mongodb_exporter[0]` | Finite enum or info | Drop | Exporter API mode is deployment metadata, not server health. |
| 11 | `db.serverStatus().metrics.repl.executor.networkInterface` | Opaque diagnostic text | Drop | Deprecated diagnostic prose is not runtime state. |
| 12 | `db.serverStatus().metrics.repl.stateTransition.lastStateTransition` | Finite enum or info | Drop | Existing role and election metrics are sufficient. |
| 13 | `db.serverStatus().oplogTruncation.processingMethod` | Finite enum or info | Bounded method label | Expose only reviewed processing-method values as a label; suppress unknown values with a bounded diagnostic. |
| 14 | `db.serverStatus().process` | Finite enum or info | Drop | Static process type is redundant on the mongod collector path. |
| 15 | `db.serverStatus().repl.electionId` | Identity and topology | Drop | Opaque election identity duplicates role and election counters. |
| 16 | `db.serverStatus().repl.hosts[0]` | Identity and topology | Drop | Service discovery owns replica-set member identity. |
| 17 | `db.serverStatus().repl.hosts[1]` | Identity and topology | Drop | Service discovery owns replica-set member identity. |
| 18 | `db.serverStatus().repl.hosts[2]` | Identity and topology | Drop | Service discovery owns replica-set member identity. |
| 19 | `db.serverStatus().repl.hosts[3]` | Identity and topology | Drop | Service discovery owns replica-set member identity. |
| 20 | `db.serverStatus().repl.lastWrite.lastWriteDate` | Date or timestamp text | Raw date gauge | Emit native BSON DateTime milliseconds since epoch with conversion factor 1. |
| 21 | `db.serverStatus().repl.lastWrite.majorityWriteDate` | Date or timestamp text | Raw date gauge | Emit native BSON DateTime milliseconds since epoch with conversion factor 1. |
| 22 | `db.serverStatus().repl.me` | Identity and topology | Drop | Prometheus target labels already identify this member. |
| 23 | `db.serverStatus().repl.primary` | Identity and topology | Drop | Role metrics identify primary state without hostname churn. |
| 24 | `db.serverStatus().repl.setName` | Identity and topology | Drop | Deployment configuration owns replica-set identity. |
| 25 | `db.serverStatus().repl.topologyVersion.processId` | Identity and topology | Drop | Opaque process ObjectID has no useful aggregate semantics. |
| 26 | `db.serverStatus().security.SSLServerSubjectName` | Identity and topology | Drop | Certificate identity is sensitive metadata; expiration and TLS counters provide operational coverage. |
| 27 | `db.serverStatus().sharding.configsvrConnectionString` | Identity and topology | Drop | Connection string is sensitive high-cardinality topology metadata. |
| 28 | `db.serverStatus().tcmalloc.tcmalloc.formattedString` | Opaque diagnostic text | Drop | Formatted prose duplicates individual numeric tcmalloc leaves. |
| 29 | `db.serverStatus().wiredTiger['snapshot-window-settings']['latest majority snapshot timestamp available']` | Date or timestamp text | Anchored date gauges | Infer Unix seconds from response `localTime`; emit a separate increment gauge; emit 0 sentinels when parsing or plausibility checks fail. |
| 30 | `db.serverStatus().wiredTiger['snapshot-window-settings']['oldest majority snapshot timestamp available']` | Date or timestamp text | Anchored date gauges | Reuse the latest timestamp's inferred timezone/year anchor; emit separate Unix-seconds and increment gauges with 0 sentinels on failure. |
| 31 | `db.serverStatus().wiredTiger.uri` | Finite enum or info | Drop | Static `statistics:` URI has no diagnostic value. |
