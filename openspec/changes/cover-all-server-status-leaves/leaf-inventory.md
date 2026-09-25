# MongoDB 5.0.34 serverStatus leaf inventory

Generated from the sanitized primary and secondary fixtures. This inventory contains **1581 unique scalar leaves**.

## Summary

| Bucket | Leaves |
|---|---:|
| Duplicate legacy and modern coverage | 10 |
| Unmapped numeric and boolean leaves | 1158 |
| Unmapped string, identity, enum, and textual leaves | 31 |
| Legacy-only coverage | 134 |
| Modern-only coverage | 248 |

Invariant check: 10 + 1158 + 31 + 134 + 248 = 1581.

### Key findings

- **1189 leaves are currently unmapped**: 1158 numeric/boolean leaves and 31 string/text leaves.
- **10 leaves are emitted twice** by both legacy and modern collectors.
- WiredTiger has 541 leaves: 41 legacy-only, 4 modern-only, 6 duplicated, and **490 unmapped**.
- The largest unmapped top-level sections are `metrics` (507), `wiredTiger` (490), `locks` (39), `indexStats` (29), `repl` (18), `shardingStatistics` (12).

### Interpretation

- `Legacy-only`: changing the source leaf changes at least one legacy Prometheus family.
- `Modern-only`: the current explicit modern contract maps the source leaf and no legacy family consumes it.
- `Duplicate`: both collectors consume the same source leaf; one side must be removed or explicitly designated as a compatibility duplicate.
- `Unmapped numeric and boolean`: no current legacy or modern metric consumes the leaf.
- `Unmapped string`: requires review before choosing bounded state, presence/count, timestamp conversion, or identity handling.

The exact-coverage target is: every leaf has exactly one owner, with both unmapped buckets and the duplicate bucket reduced to zero.

## Review buckets

## Duplicate legacy and modern coverage

### metrics.cursor (4)

- `db.serverStatus().metrics.cursor.open.noTimeout` — legacy: mongodb_mongod_metrics_cursor_open; modern: mongodb_server_status_metrics_cursor_open_no_timeout
- `db.serverStatus().metrics.cursor.open.pinned` — legacy: mongodb_mongod_metrics_cursor_open; modern: mongodb_server_status_metrics_cursor_open_pinned
- `db.serverStatus().metrics.cursor.open.total` — legacy: mongodb_mongod_metrics_cursor_open; modern: mongodb_server_status_metrics_cursor_open
- `db.serverStatus().metrics.cursor.timedOut` — legacy: mongodb_mongod_metrics_cursor_timed_out_total; modern: mongodb_server_status_metrics_cursor_timed_out_total

### wiredTiger.concurrentTransactions (6)

- `db.serverStatus().wiredTiger.concurrentTransactions.read.available` — legacy: mongodb_mongod_wiredtiger_concurrent_transactions_available_tickets; modern: mongodb_server_status_wired_tiger_concurrent_transactions
- `db.serverStatus().wiredTiger.concurrentTransactions.read.out` — legacy: mongodb_mongod_wiredtiger_concurrent_transactions_out_tickets; modern: mongodb_server_status_wired_tiger_concurrent_transactions
- `db.serverStatus().wiredTiger.concurrentTransactions.read.totalTickets` — legacy: mongodb_mongod_wiredtiger_concurrent_transactions_total_tickets; modern: mongodb_server_status_wired_tiger_concurrent_transactions
- `db.serverStatus().wiredTiger.concurrentTransactions.write.available` — legacy: mongodb_mongod_wiredtiger_concurrent_transactions_available_tickets; modern: mongodb_server_status_wired_tiger_concurrent_transactions
- `db.serverStatus().wiredTiger.concurrentTransactions.write.out` — legacy: mongodb_mongod_wiredtiger_concurrent_transactions_out_tickets; modern: mongodb_server_status_wired_tiger_concurrent_transactions
- `db.serverStatus().wiredTiger.concurrentTransactions.write.totalTickets` — legacy: mongodb_mongod_wiredtiger_concurrent_transactions_total_tickets; modern: mongodb_server_status_wired_tiger_concurrent_transactions

## Unmapped numeric and boolean leaves

### $clusterTime.clusterTime[0] (1)

- `db.serverStatus().$clusterTime.clusterTime[0]` — type: number; roles: primary, secondary

### $clusterTime.clusterTime[1] (1)

- `db.serverStatus().$clusterTime.clusterTime[1]` — type: number; roles: primary, secondary

### $clusterTime.signature (1)

- `db.serverStatus().$clusterTime.signature.keyId` — type: number; roles: primary, secondary

### $configServerState.opTime (3)

- `db.serverStatus().$configServerState.opTime.t` — type: number; roles: primary, secondary
- `db.serverStatus().$configServerState.opTime.ts[0]` — type: number; roles: primary, secondary
- `db.serverStatus().$configServerState.opTime.ts[1]` — type: number; roles: primary, secondary

### $gleStats.lastOpTime[0] (1)

- `db.serverStatus().$gleStats.lastOpTime[0]` — type: number; roles: primary, secondary

### $gleStats.lastOpTime[1] (1)

- `db.serverStatus().$gleStats.lastOpTime[1]` — type: number; roles: primary, secondary

### asserts.tripwire (1)

- `db.serverStatus().asserts.tripwire` — type: number; roles: primary, secondary

### catalogStats.capped (1)

- `db.serverStatus().catalogStats.capped` — type: number; roles: primary, secondary

### catalogStats.collections (1)

- `db.serverStatus().catalogStats.collections` — type: number; roles: primary, secondary

### catalogStats.internalCollections (1)

- `db.serverStatus().catalogStats.internalCollections` — type: number; roles: primary, secondary

### catalogStats.internalViews (1)

- `db.serverStatus().catalogStats.internalViews` — type: number; roles: primary, secondary

### catalogStats.timeseries (1)

- `db.serverStatus().catalogStats.timeseries` — type: number; roles: primary, secondary

### catalogStats.views (1)

- `db.serverStatus().catalogStats.views` — type: number; roles: primary, secondary

### connections.awaitingTopologyChanges (1)

- `db.serverStatus().connections.awaitingTopologyChanges` — type: number; roles: primary, secondary

### connections.exhaustHello (1)

- `db.serverStatus().connections.exhaustHello` — type: number; roles: primary, secondary

### connections.exhaustIsMaster (1)

- `db.serverStatus().connections.exhaustIsMaster` — type: number; roles: primary, secondary

### connections.threaded (1)

- `db.serverStatus().connections.threaded` — type: number; roles: primary, secondary

### extra_info.input_blocks (1)

- `db.serverStatus().extra_info.input_blocks` — type: number; roles: primary, secondary

### extra_info.involuntary_context_switches (1)

- `db.serverStatus().extra_info.involuntary_context_switches` — type: number; roles: primary, secondary

### extra_info.maximum_resident_set_kb (1)

- `db.serverStatus().extra_info.maximum_resident_set_kb` — type: number; roles: primary, secondary

### extra_info.output_blocks (1)

- `db.serverStatus().extra_info.output_blocks` — type: number; roles: primary, secondary

### extra_info.page_reclaims (1)

- `db.serverStatus().extra_info.page_reclaims` — type: number; roles: primary, secondary

### extra_info.system_time_us (1)

- `db.serverStatus().extra_info.system_time_us` — type: number; roles: primary, secondary

### extra_info.threads (1)

- `db.serverStatus().extra_info.threads` — type: number; roles: primary, secondary

### extra_info.user_time_us (1)

- `db.serverStatus().extra_info.user_time_us` — type: number; roles: primary, secondary

### extra_info.voluntary_context_switches (1)

- `db.serverStatus().extra_info.voluntary_context_switches` — type: number; roles: primary, secondary

### featureCompatibilityVersion.major (1)

- `db.serverStatus().featureCompatibilityVersion.major` — type: number; roles: primary, secondary

### featureCompatibilityVersion.minor (1)

- `db.serverStatus().featureCompatibilityVersion.minor` — type: number; roles: primary, secondary

### featureCompatibilityVersion.transitioning (1)

- `db.serverStatus().featureCompatibilityVersion.transitioning` — type: number; roles: primary, secondary

### globalLock.activeClients (1)

- `db.serverStatus().globalLock.activeClients.total` — type: number; roles: primary, secondary

### globalLock.currentQueue (1)

- `db.serverStatus().globalLock.currentQueue.total` — type: number; roles: primary, secondary

### globalLock.totalTime (1)

- `db.serverStatus().globalLock.totalTime` — type: number; roles: primary, secondary

### indexBulkBuilder.count (1)

- `db.serverStatus().indexBulkBuilder.count` — type: number; roles: primary, secondary

### indexBulkBuilder.filesClosedForExternalSort (1)

- `db.serverStatus().indexBulkBuilder.filesClosedForExternalSort` — type: number; roles: primary, secondary

### indexBulkBuilder.filesOpenedForExternalSort (1)

- `db.serverStatus().indexBulkBuilder.filesOpenedForExternalSort` — type: number; roles: primary, secondary

### indexBulkBuilder.resumed (1)

- `db.serverStatus().indexBulkBuilder.resumed` — type: number; roles: primary, secondary

### indexStats.count (1)

- `db.serverStatus().indexStats.count` — type: number; roles: primary, secondary

### indexStats.features (28)

- `db.serverStatus().indexStats.features['2d'].accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features['2d'].count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features['2dsphere'].accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features['2dsphere'].count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.collation.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.collation.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.compound.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.compound.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.hashed.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.hashed.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.id.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.id.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.normal.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.normal.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.partial.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.partial.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.single.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.single.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.sparse.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.sparse.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.text.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.text.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.ttl.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.ttl.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.unique.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.unique.count` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.wildcard.accesses` — type: number; roles: primary, secondary
- `db.serverStatus().indexStats.features.wildcard.count` — type: number; roles: primary, secondary

### lastCommittedOpTime[0] (1)

- `db.serverStatus().lastCommittedOpTime[0]` — type: number; roles: primary, secondary

### lastCommittedOpTime[1] (1)

- `db.serverStatus().lastCommittedOpTime[1]` — type: number; roles: primary, secondary

### locks.Collection (9)

- `db.serverStatus().locks.Collection.acquireCount.R` — type: number; roles: primary
- `db.serverStatus().locks.Collection.acquireCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Collection.acquireCount.r` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Collection.acquireCount.w` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Collection.acquireWaitCount.R` — type: number; roles: primary
- `db.serverStatus().locks.Collection.acquireWaitCount.W` — type: number; roles: primary
- `db.serverStatus().locks.Collection.acquireWaitCount.w` — type: number; roles: primary
- `db.serverStatus().locks.Collection.timeAcquiringMicros.R` — type: number; roles: primary
- `db.serverStatus().locks.Collection.timeAcquiringMicros.W` — type: number; roles: primary

### locks.Database (6)

- `db.serverStatus().locks.Database.acquireCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Database.acquireCount.r` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Database.acquireCount.w` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Database.acquireWaitCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Database.acquireWaitCount.w` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Database.timeAcquiringMicros.W` — type: number; roles: primary, secondary

### locks.FeatureCompatibilityVersion (2)

- `db.serverStatus().locks.FeatureCompatibilityVersion.acquireCount.r` — type: number; roles: primary, secondary
- `db.serverStatus().locks.FeatureCompatibilityVersion.acquireCount.w` — type: number; roles: primary, secondary

### locks.Global (6)

- `db.serverStatus().locks.Global.acquireCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Global.acquireCount.r` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Global.acquireCount.w` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Global.acquireWaitCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Global.acquireWaitCount.r` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Global.timeAcquiringMicros.W` — type: number; roles: primary, secondary

### locks.Mutex (5)

- `db.serverStatus().locks.Mutex.acquireCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Mutex.acquireCount.r` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Mutex.acquireWaitCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Mutex.acquireWaitCount.r` — type: number; roles: primary, secondary
- `db.serverStatus().locks.Mutex.timeAcquiringMicros.W` — type: number; roles: primary, secondary

### locks.oplog (3)

- `db.serverStatus().locks.oplog.acquireCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.oplog.acquireCount.r` — type: number; roles: primary, secondary
- `db.serverStatus().locks.oplog.acquireCount.w` — type: number; roles: primary, secondary

### locks.ParallelBatchWriterMode (5)

- `db.serverStatus().locks.ParallelBatchWriterMode.acquireCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.ParallelBatchWriterMode.acquireCount.r` — type: number; roles: primary, secondary
- `db.serverStatus().locks.ParallelBatchWriterMode.acquireWaitCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.ParallelBatchWriterMode.acquireWaitCount.r` — type: number; roles: primary, secondary
- `db.serverStatus().locks.ParallelBatchWriterMode.timeAcquiringMicros.W` — type: number; roles: primary, secondary

### locks.ReplicationStateTransition (3)

- `db.serverStatus().locks.ReplicationStateTransition.acquireCount.W` — type: number; roles: primary, secondary
- `db.serverStatus().locks.ReplicationStateTransition.acquireCount.w` — type: number; roles: primary, secondary
- `db.serverStatus().locks.ReplicationStateTransition.acquireWaitCount.w` — type: number; roles: primary

### mem.bits (1)

- `db.serverStatus().mem.bits` — type: number; roles: primary, secondary

### mem.supported (1)

- `db.serverStatus().mem.supported` — type: boolean; roles: primary, secondary

### metrics.aggStageCounters (39)

- `db.serverStatus().metrics.aggStageCounters.$_internalApplyOplogUpdate` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalBoundedSort` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalConvertBucketIndexStats` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalFindAndModifyImageLookup` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalInhibitOptimization` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalReshardingIterateTransaction` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalReshardingOwnershipMatch` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalSetWindowFields` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalShredDocuments` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalSplitPipeline` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalUnpackBucket` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_unpackBucket` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$addFields` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$bucket` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$bucketAuto` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$changeStream` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$collStats` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$currentOp` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$documents` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$facet` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$geoNear` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$graphLookup` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$listLocalSessions` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$listSessions` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$merge` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$mergeCursors` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$operationMetrics` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$out` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$planCacheStats` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$queue` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$redact` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$replaceRoot` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$replaceWith` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$setWindowFields` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$skip` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$sortByCount` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$unionWith` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$unset` — type: number; roles: primary, secondary

### metrics.changeStreams (1)

- `db.serverStatus().metrics.changeStreams.largeEventsFailed` — type: number; roles: primary, secondary

### metrics.commands (107)

- `db.serverStatus().metrics.commands._flushDatabaseCacheUpdates.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._flushDatabaseCacheUpdates.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._flushDatabaseCacheUpdatesWithWriteConcern.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._flushDatabaseCacheUpdatesWithWriteConcern.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._flushRoutingTableCacheUpdates.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._flushRoutingTableCacheUpdates.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._flushRoutingTableCacheUpdatesWithWriteConcern.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._flushRoutingTableCacheUpdatesWithWriteConcern.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._getNextSessionMods.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._getNextSessionMods.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._killOperations.failed` — type: number; roles: secondary
- `db.serverStatus().metrics.commands._killOperations.total` — type: number; roles: secondary
- `db.serverStatus().metrics.commands._migrateClone.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._migrateClone.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._recvChunkAbort.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._recvChunkAbort.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._recvChunkCommit.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._recvChunkCommit.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._recvChunkStart.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._recvChunkStart.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._recvChunkStatus.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._recvChunkStatus.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrCreateCollectionParticipant.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrCreateCollectionParticipant.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropCollection.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropCollection.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropCollectionParticipant.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropCollectionParticipant.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropDatabase.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropDatabase.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropDatabaseParticipant.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropDatabaseParticipant.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrSetAllowMigrations.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._shardsvrSetAllowMigrations.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands._transferMods.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands._transferMods.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.aggregate.allowDiskUseTrue` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.authenticate.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.authenticate.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.autoSplitVector.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.autoSplitVector.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.buildInfo.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.buildInfo.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.collStats.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.collStats.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.connectionStatus.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.connectionStatus.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.count.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.count.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.createIndexes.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.createIndexes.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.currentOp.failed` — type: number; roles: secondary
- `db.serverStatus().metrics.commands.currentOp.total` — type: number; roles: secondary
- `db.serverStatus().metrics.commands.distinct.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.distinct.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.drop.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.drop.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.endSessions.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.endSessions.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.features.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.features.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.findAndModify.arrayFilters` — type: number; roles: primary
- `db.serverStatus().metrics.commands.findAndModify.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.findAndModify.pipeline` — type: number; roles: primary
- `db.serverStatus().metrics.commands.findAndModify.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.getCmdLineOpts.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.getCmdLineOpts.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.getLog.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.getLog.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.getParameter.failed` — type: number; roles: secondary
- `db.serverStatus().metrics.commands.getParameter.total` — type: number; roles: secondary
- `db.serverStatus().metrics.commands.hello.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.hello.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.isMaster.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.isMaster.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.killCursors.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.killCursors.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.listCollections.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.listCollections.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.listDatabases.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.listDatabases.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.listIndexes.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.listIndexes.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.ping.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.ping.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetGetConfig.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetGetConfig.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetGetStatus.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetGetStatus.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetHeartbeat.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetHeartbeat.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetReconfig.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.replSetReconfig.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.replSetRequestVotes.failed` — type: number; roles: secondary
- `db.serverStatus().metrics.commands.replSetRequestVotes.total` — type: number; roles: secondary
- `db.serverStatus().metrics.commands.replSetStepUp.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.replSetStepUp.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.replSetUpdatePosition.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetUpdatePosition.total` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.setShardVersion.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.setShardVersion.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.splitChunk.failed` — type: number; roles: primary
- `db.serverStatus().metrics.commands.splitChunk.total` — type: number; roles: primary
- `db.serverStatus().metrics.commands.update.arrayFilters` — type: number; roles: primary
- `db.serverStatus().metrics.commands.update.pipeline` — type: number; roles: primary
- `db.serverStatus().metrics.commands.whatsmyuri.failed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.commands.whatsmyuri.total` — type: number; roles: primary, secondary

### metrics.dotsAndDollarsFields (2)

- `db.serverStatus().metrics.dotsAndDollarsFields.inserts` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.dotsAndDollarsFields.updates` — type: number; roles: primary, secondary

### metrics.getLastError (2)

- `db.serverStatus().metrics.getLastError.default.unsatisfiable` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.getLastError.default.wtimeouts` — type: number; roles: primary, secondary

### metrics.mongos (2)

- `db.serverStatus().metrics.mongos.cursor.moreThanOneBatch` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.mongos.cursor.totalOpened` — type: number; roles: primary, secondary

### metrics.operation (3)

- `db.serverStatus().metrics.operation.transactionTooLargeForCacheErrors` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operation.transactionTooLargeForCacheErrorsConvertedToWriteConflict` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operation.writeConflicts` — type: number; roles: primary, secondary

### metrics.operatorCounters (200)

- `db.serverStatus().metrics.operatorCounters.expressions.$_internalJsEmit` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$_internalKeyStringValue` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$abs` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$acos` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$acosh` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$allElementsTrue` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$and` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$anyElementTrue` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$arrayElemAt` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$arrayToObject` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$asin` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$asinh` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$atan` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$atan2` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$atanh` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$avg` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$binarySize` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$bsonSize` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$ceil` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$cmp` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$concat` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$concatArrays` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$cond` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$const` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$convert` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$cos` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$cosh` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateAdd` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateDiff` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateFromParts` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateFromString` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateSubtract` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateToParts` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateToString` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateTrunc` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dayOfMonth` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dayOfWeek` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dayOfYear` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$degreesToRadians` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$divide` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$exp` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$filter` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$first` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$floor` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$function` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$getField` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$gte` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$hour` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$indexOfArray` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$indexOfBytes` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$indexOfCP` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$isArray` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$isNumber` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$isoDayOfWeek` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$isoWeek` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$isoWeekYear` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$last` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$let` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$literal` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$ln` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$log` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$log10` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$lt` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$lte` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$ltrim` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$map` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$max` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$mergeObjects` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$meta` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$millisecond` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$min` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$minute` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$mod` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$month` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$ne` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$not` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$objectToArray` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$or` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$pow` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$radiansToDegrees` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$rand` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$range` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$regexFind` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$regexFindAll` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$regexMatch` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$replaceAll` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$replaceOne` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$reverseArray` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$round` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$rtrim` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$second` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setDifference` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setEquals` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setField` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setIntersection` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setIsSubset` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setUnion` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$sin` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$sinh` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$slice` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$split` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$sqrt` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$stdDevPop` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$stdDevSamp` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$strLenBytes` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$strLenCP` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$strcasecmp` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$substr` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$substrBytes` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$substrCP` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$sum` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$switch` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$tan` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$tanh` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toBool` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toDate` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toDecimal` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toDouble` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toHashedIndexKey` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toInt` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toLong` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toLower` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toObjectId` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toString` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toUpper` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$trim` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$trunc` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$type` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$unsetField` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$week` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$year` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$zip` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$_internalJsReduce` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$accumulator` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$addToSet` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$avg` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$first` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$last` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$max` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$mergeObjects` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$min` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$push` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$stdDevPop` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$stdDevSamp` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$sum` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$all` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$alwaysFalse` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$alwaysTrue` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$and` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$bitsAllClear` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$bitsAllSet` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$bitsAnyClear` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$bitsAnySet` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$comment` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$elemMatch` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$exists` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$expr` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$geoIntersects` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$geoWithin` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$gt` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$gte` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$in` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$jsonSchema` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$lt` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$lte` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$mod` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$ne` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$near` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$nearSphere` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$nin` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$nor` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$not` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$or` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$regex` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$sampleRate` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$size` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$text` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$type` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$where` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$addToSet` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$avg` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$covariancePop` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$covarianceSamp` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$denseRank` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$derivative` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$documentNumber` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$expMovingAvg` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$first` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$integral` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$last` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$max` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$min` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$push` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$rank` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$shift` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$stdDevPop` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$stdDevSamp` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$sum` — type: number; roles: primary, secondary

### metrics.query (118)

- `db.serverStatus().metrics.query.multiPlanner.classicCount` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.classicMicros` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.classicWorks` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[0].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[0].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[10].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[10].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[11].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[11].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[1].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[1].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[2].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[2].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[3].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[3].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[4].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[4].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[5].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[5].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[6].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[6].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[7].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[7].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[8].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[8].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[9].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[9].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[0].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[0].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[1].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[1].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[2].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[2].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[3].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[3].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[4].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[4].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[5].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[5].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[0].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[0].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[1].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[1].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[2].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[2].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[3].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[3].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[4].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[4].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[5].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[5].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[6].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[6].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[7].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[7].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[8].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[8].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[9].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[9].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[0].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[0].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[10].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[10].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[11].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[11].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[1].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[1].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[2].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[2].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[3].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[3].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[4].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[4].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[5].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[5].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[6].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[6].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[7].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[7].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[8].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[8].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[9].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[9].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[0].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[0].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[1].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[1].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[2].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[2].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[3].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[3].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[4].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[4].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[5].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[5].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[0].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[0].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[1].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[1].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[2].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[2].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[3].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[3].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[4].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[4].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[5].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[5].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[6].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[6].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[7].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[7].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[8].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[8].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[9].count` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[9].lowerBound` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.sbeCount` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.sbeMicros` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.sbeNumReads` — type: number; roles: primary, secondary

### metrics.queryExecutor (2)

- `db.serverStatus().metrics.queryExecutor.collectionScans.nonTailable` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.queryExecutor.collectionScans.total` — type: number; roles: primary, secondary

### metrics.repl (23)

- `db.serverStatus().metrics.repl.apply.attemptsToBecomeSecondary` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.apply.batchSize` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.executor.pool.inProgressCount` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.executor.shuttingDown` — type: boolean; roles: primary, secondary
- `db.serverStatus().metrics.repl.initialSync.completed` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.initialSync.failedAttempts` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.initialSync.failures` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.network.getmores.numEmptyBatches` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.network.notPrimaryLegacyUnacknowledgedWrites` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.network.notPrimaryUnacknowledgedWrites` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.network.oplogGetMoresProcessed.num` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.network.oplogGetMoresProcessed.totalMillis` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.network.replSetUpdatePosition.num` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.reconfig.numAutoReconfigsForRemovalOfNewlyAddedFields` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.stateTransition.userOperationsKilled` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.stateTransition.userOperationsRunning` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.syncSource.numSelections` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.syncSource.numSyncSourceChangesDueToSignificantlyCloserNode` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.syncSource.numTimesChoseDifferent` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.syncSource.numTimesChoseSame` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.syncSource.numTimesCouldNotFind` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.waiters.opTime` — type: number; roles: primary, secondary
- `db.serverStatus().metrics.repl.waiters.replication` — type: number; roles: primary, secondary

### ok (1)

- `db.serverStatus().ok` — type: number; roles: primary, secondary

### operationTime[0] (1)

- `db.serverStatus().operationTime[0]` — type: number; roles: primary, secondary

### operationTime[1] (1)

- `db.serverStatus().operationTime[1]` — type: number; roles: primary, secondary

### opLatencies.transactions (2)

- `db.serverStatus().opLatencies.transactions.latency` — type: number; roles: primary, secondary
- `db.serverStatus().opLatencies.transactions.ops` — type: number; roles: primary, secondary

### pid (1)

- `db.serverStatus().pid` — type: number; roles: primary, secondary

### repl.lastWrite (6)

- `db.serverStatus().repl.lastWrite.majorityOpTime.t` — type: number; roles: primary, secondary
- `db.serverStatus().repl.lastWrite.majorityOpTime.ts[0]` — type: number; roles: primary, secondary
- `db.serverStatus().repl.lastWrite.majorityOpTime.ts[1]` — type: number; roles: primary, secondary
- `db.serverStatus().repl.lastWrite.opTime.t` — type: number; roles: primary, secondary
- `db.serverStatus().repl.lastWrite.opTime.ts[0]` — type: number; roles: primary, secondary
- `db.serverStatus().repl.lastWrite.opTime.ts[1]` — type: number; roles: primary, secondary

### repl.topologyVersion (1)

- `db.serverStatus().repl.topologyVersion.counter` — type: number; roles: primary, secondary

### scramCache.SCRAM-SHA-1 (3)

- `db.serverStatus().scramCache['SCRAM-SHA-1'].count` — type: number; roles: primary, secondary
- `db.serverStatus().scramCache['SCRAM-SHA-1'].hits` — type: number; roles: primary, secondary
- `db.serverStatus().scramCache['SCRAM-SHA-1'].misses` — type: number; roles: primary, secondary

### scramCache.SCRAM-SHA-256 (3)

- `db.serverStatus().scramCache['SCRAM-SHA-256'].count` — type: number; roles: primary, secondary
- `db.serverStatus().scramCache['SCRAM-SHA-256'].hits` — type: number; roles: primary, secondary
- `db.serverStatus().scramCache['SCRAM-SHA-256'].misses` — type: number; roles: primary, secondary

### sharding.lastSeenConfigServerOpTime (3)

- `db.serverStatus().sharding.lastSeenConfigServerOpTime.t` — type: number; roles: primary, secondary
- `db.serverStatus().sharding.lastSeenConfigServerOpTime.ts[0]` — type: number; roles: primary, secondary
- `db.serverStatus().sharding.lastSeenConfigServerOpTime.ts[1]` — type: number; roles: primary, secondary

### sharding.maxChunkSizeInBytes (1)

- `db.serverStatus().sharding.maxChunkSizeInBytes` — type: number; roles: primary, secondary

### shardingStatistics.resharding (12)

- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis['(-inf, 10)'].count` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis['[10, 100)'].count` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis['[100, 1000)'].count` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis['[1000, 10000)'].count` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis['[10000, inf)'].count` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis.totalCount` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis['(-inf, 10)'].count` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis['[10, 100)'].count` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis['[100, 1000)'].count` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis['[1000, 10000)'].count` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis['[10000, inf)'].count` — type: number; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis.totalCount` — type: number; roles: primary, secondary

### storageEngine.oldestRequiredTimestampForCrashRecovery[0] (1)

- `db.serverStatus().storageEngine.oldestRequiredTimestampForCrashRecovery[0]` — type: number; roles: primary, secondary

### storageEngine.oldestRequiredTimestampForCrashRecovery[1] (1)

- `db.serverStatus().storageEngine.oldestRequiredTimestampForCrashRecovery[1]` — type: number; roles: primary, secondary

### tcmalloc.tcmalloc (2)

- `db.serverStatus().tcmalloc.tcmalloc.release_rate` — type: number; roles: primary, secondary
- `db.serverStatus().tcmalloc.tcmalloc.spinlock_total_delay_ns` — type: number; roles: primary, secondary

### tenantMigrations.currentMigrationsDonating (1)

- `db.serverStatus().tenantMigrations.currentMigrationsDonating` — type: number; roles: primary, secondary

### tenantMigrations.currentMigrationsReceiving (1)

- `db.serverStatus().tenantMigrations.currentMigrationsReceiving` — type: number; roles: primary, secondary

### tenantMigrations.totalFailedMigrationsDonated (1)

- `db.serverStatus().tenantMigrations.totalFailedMigrationsDonated` — type: number; roles: primary, secondary

### tenantMigrations.totalFailedMigrationsReceived (1)

- `db.serverStatus().tenantMigrations.totalFailedMigrationsReceived` — type: number; roles: primary, secondary

### tenantMigrations.totalSuccessfulMigrationsDonated (1)

- `db.serverStatus().tenantMigrations.totalSuccessfulMigrationsDonated` — type: number; roles: primary, secondary

### tenantMigrations.totalSuccessfulMigrationsReceived (1)

- `db.serverStatus().tenantMigrations.totalSuccessfulMigrationsReceived` — type: number; roles: primary, secondary

### trafficRecording.running (1)

- `db.serverStatus().trafficRecording.running` — type: boolean; roles: primary, secondary

### transportSecurity.1 (4)

- `db.serverStatus().transportSecurity['1.0']` — type: number; roles: primary, secondary
- `db.serverStatus().transportSecurity['1.1']` — type: number; roles: primary, secondary
- `db.serverStatus().transportSecurity['1.2']` — type: number; roles: primary, secondary
- `db.serverStatus().transportSecurity['1.3']` — type: number; roles: primary, secondary

### transportSecurity.unknown (1)

- `db.serverStatus().transportSecurity.unknown` — type: number; roles: primary, secondary

### twoPhaseCommitCoordinator.currentInSteps (5)

- `db.serverStatus().twoPhaseCommitCoordinator.currentInSteps.deletingCoordinatorDoc` — type: number; roles: primary, secondary
- `db.serverStatus().twoPhaseCommitCoordinator.currentInSteps.waitingForDecisionAcks` — type: number; roles: primary, secondary
- `db.serverStatus().twoPhaseCommitCoordinator.currentInSteps.waitingForVotes` — type: number; roles: primary, secondary
- `db.serverStatus().twoPhaseCommitCoordinator.currentInSteps.writingDecision` — type: number; roles: primary, secondary
- `db.serverStatus().twoPhaseCommitCoordinator.currentInSteps.writingParticipantList` — type: number; roles: primary, secondary

### twoPhaseCommitCoordinator.totalAbortedTwoPhaseCommit (1)

- `db.serverStatus().twoPhaseCommitCoordinator.totalAbortedTwoPhaseCommit` — type: number; roles: primary, secondary

### twoPhaseCommitCoordinator.totalCommittedTwoPhaseCommit (1)

- `db.serverStatus().twoPhaseCommitCoordinator.totalCommittedTwoPhaseCommit` — type: number; roles: primary, secondary

### twoPhaseCommitCoordinator.totalCreated (1)

- `db.serverStatus().twoPhaseCommitCoordinator.totalCreated` — type: number; roles: primary, secondary

### twoPhaseCommitCoordinator.totalStartedTwoPhaseCommit (1)

- `db.serverStatus().twoPhaseCommitCoordinator.totalStartedTwoPhaseCommit` — type: number; roles: primary, secondary

### uptimeEstimate (1)

- `db.serverStatus().uptimeEstimate` — type: number; roles: primary, secondary

### uptimeMillis (1)

- `db.serverStatus().uptimeMillis` — type: number; roles: primary, secondary

### wiredTiger.block-manager (28)

- `db.serverStatus().wiredTiger['block-manager']['block cache cached blocks updated']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache cached bytes updated']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache evicted blocks']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache file size causing bypass']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache lookups']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache number of blocks not evicted due to overhead']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache number of bypasses because no-write-allocate setting was on']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache number of bypasses due to overhead on put']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache number of bypasses on get']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache number of bypasses on put because file is too small']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache number of eviction passes']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache number of hits including existence checks']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache number of misses including existence checks']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache number of put bypasses on checkpoint I/O']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache removed blocks']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache total blocks']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache total blocks inserted on read path']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache total blocks inserted on write path']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache total bytes']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache total bytes inserted on read path']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['block cache total bytes inserted on write path']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['bytes read via memory map API']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['bytes read via system call API']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['bytes written for checkpoint']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['bytes written via memory map API']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['bytes written via system call API']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['number of times the file was remapped because it changed size via fallocate or truncate']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['block-manager']['number of times the region was remapped via write']` — type: number; roles: primary, secondary

### wiredTiger.cache (137)

- `db.serverStatus().wiredTiger.cache['application threads page read from disk to cache count']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['application threads page read from disk to cache time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['application threads page write from cache to disk count']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['application threads page write from cache to disk time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['bytes allocated for updates']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['bytes belonging to page images in the cache']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['bytes belonging to the history store table in the cache']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['bytes dirty in the cache cumulative']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['bytes not belonging to page images in the cache']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['cache overflow score']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['checkpoint blocked page eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['checkpoint of history store file blocked non-history store page eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction calls to get a page']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction calls to get a page found queue empty']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction calls to get a page found queue empty after locking']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction currently operating in aggressive mode']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction empty score']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction gave up due to detecting an out of order on disk value behind the last update on the chain']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction gave up due to detecting an out of order tombstone ahead of the selected on disk update']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction gave up due to detecting an out of order tombstone ahead of the selected on disk update after validating the update chain']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction gave up due to detecting out of order timestamps on the update chain after the selected on disk update']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction gave up due to needing to remove a record from the history store but checkpoint is running']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction passes of a file']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server candidate queue empty when topping up']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server candidate queue not empty when topping up']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server evicting pages']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server skips dirty pages during a running checkpoint']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server skips metadata pages with history']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server skips pages that are written with transactions greater than the last running']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server skips pages that previously failed eviction and likely will again']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server skips pages that we do not want to evict']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server skips trees because there are too many active walks']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server skips trees that are being checkpointed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server skips trees that are configured to stick in cache']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server skips trees that disable eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server skips trees that were not useful before']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server slept, because we did not make progress with eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server unable to reach eviction goal']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction server waiting for a leaf page']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction state']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk most recent sleeps for checkpoint handle gathering']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target pages histogram - 0-9']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target pages histogram - 10-31']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target pages histogram - 128 and higher']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target pages histogram - 32-63']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target pages histogram - 64-128']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target pages reduced due to history store cache pressure']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target strategy both clean and dirty pages']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target strategy only clean pages']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target strategy only dirty pages']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walks abandoned']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walks gave up because they restarted their walk twice']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walks gave up because they saw too many pages and found no candidates']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walks gave up because they saw too many pages and found too few candidates']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walks reached end of tree']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walks restarted']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walks started from root of tree']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walks started from saved location in tree']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction worker thread active']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction worker thread created']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction worker thread evicting pages']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction worker thread removed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction worker thread stable number']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['files with active eviction walks']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['files with new eviction walks started']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['force re-tuning of eviction workers once in a while']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - do not retry count to evict pages selected to evict during reconciliation']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - history store pages failed to evict while session has history store cursor open']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - history store pages selected while session has history store cursor open']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - history store pages successfully evicted while session has history store cursor open']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - pages evicted that were clean count']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - pages evicted that were clean time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - pages evicted that were dirty count']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - pages evicted that were dirty time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - pages selected because of a large number of updates to a single item']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - pages selected because of too many deleted items count']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - pages selected count']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - pages selected unable to be evicted count']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['forced eviction - pages selected unable to be evicted time']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['hazard pointer blocked page eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['hazard pointer check calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['hazard pointer check entries walked']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['hazard pointer maximum array length']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store score']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table insert calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table insert calls that returned restart']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table max on-disk size']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table on-disk size']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table out-of-order resolved updates that lose their durable timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table out-of-order updates that were fixed up by reinserting with the fixed timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table reads']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table reads missed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table reads requiring squashed modifies']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table truncation by rollback to stable to remove an unstable update']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table truncation by rollback to stable to remove an update']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table truncation to remove an update']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table truncation to remove range of updates due to key being removed from the data page during reconciliation']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table truncation to remove range of updates due to out-of-order timestamp update on data page']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['history store table writes requiring squashed modifies']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['in-memory page passed criteria to be split']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['in-memory page splits']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['internal pages evicted']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['internal pages queued for eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['internal pages seen by eviction walk']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['internal pages seen by eviction walk that are already queued']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['internal pages split during eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['leaf pages split during eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['maximum milliseconds spent at a single eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['maximum page size seen at eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['modified pages evicted by application threads']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['operations timed out waiting for space in cache']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['overflow pages read into cache']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['page split during eviction deepened the tree']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['page written requiring history store records']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages dirtied due to obsolete time window']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages evicted by application threads']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages evicted in parallel with checkpoint']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages queued for eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages queued for eviction post lru sorting']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages queued for urgent eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages queued for urgent eviction during walk']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages queued for urgent eviction from history store due to high dirty content']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages read into cache after truncate']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages read into cache after truncate in prepare state']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages removed from the ordinary queue to be queued for urgent eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages requested from the cache']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages seen by eviction walk']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages seen by eviction walk that are already queued']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages selected for eviction unable to be evicted']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages selected for eviction unable to be evicted because of active children on an internal page']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages selected for eviction unable to be evicted because of failure in reconciliation']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages selected for eviction unable to be evicted because of race between checkpoint and out of order timestamps handling']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages walked for eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['pages written requiring in-memory restoration']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['the number of times full update inserted to history store']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['the number of times reverse modify inserted to history store']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['total milliseconds spent inside reentrant history store evictions in a reconciliation']` — type: number; roles: primary, secondary

### wiredTiger.capacity (14)

- `db.serverStatus().wiredTiger.capacity['background fsync file handles considered']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['background fsync file handles synced']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['background fsync time (msecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['bytes read']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['bytes written for checkpoint']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['bytes written for eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['bytes written for log']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['bytes written total']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['threshold to call fsync']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['time waiting due to total capacity (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['time waiting during checkpoint (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['time waiting during eviction (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['time waiting during logging (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.capacity['time waiting during read (usecs)']` — type: number; roles: primary, secondary

### wiredTiger.checkpoint-cleanup (4)

- `db.serverStatus().wiredTiger['checkpoint-cleanup']['pages added for eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['checkpoint-cleanup']['pages removed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['checkpoint-cleanup']['pages skipped during tree walk']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['checkpoint-cleanup']['pages visited']` — type: number; roles: primary, secondary

### wiredTiger.connection (16)

- `db.serverStatus().wiredTiger.connection['auto adjusting condition resets']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['auto adjusting condition wait calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['auto adjusting condition wait raced to update timeout and skipped updating']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['detected system time went backwards']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['files currently open']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['hash bucket array size for data handles']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['hash bucket array size general']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['memory allocations']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['memory frees']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['memory re-allocations']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['pthread mutex condition wait calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['pthread mutex shared lock read-lock calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['pthread mutex shared lock write-lock calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['total fsync I/Os']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['total read I/Os']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.connection['total write I/Os']` — type: number; roles: primary, secondary

### wiredTiger.cursor (41)

- `db.serverStatus().wiredTiger.cursor['Total number of deleted pages skipped during tree walk']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['Total number of entries skipped by cursor next calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['Total number of entries skipped by cursor prev calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['Total number of entries skipped to position the history store cursor']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['Total number of in-memory deleted pages skipped during tree walk']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['Total number of times a search near has exited due to prefix config']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cached cursor count']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor bulk loaded cursor insert calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor close calls that result in cache']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor create calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor insert calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor insert key and value bytes']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor modify calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor modify key and value bytes affected']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor modify value bytes modified']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor next calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor next calls that skip due to a globally visible history store tombstone']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor next calls that skip greater than or equal to 100 entries']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor next calls that skip less than 100 entries']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor operation restarted']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor prev calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor prev calls that skip due to a globally visible history store tombstone']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor prev calls that skip greater than or equal to 100 entries']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor prev calls that skip less than 100 entries']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor remove calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor remove key bytes removed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor reserve calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor reset calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor search calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor search history store calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor search near calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor sweep buckets']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor sweep cursors closed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor sweep cursors examined']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor sweeps']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor truncate calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor update calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor update key and value bytes']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursor update value size change']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['cursors reused from cache']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.cursor['open cursor count']` — type: number; roles: primary, secondary

### wiredTiger.data-handle (10)

- `db.serverStatus().wiredTiger['data-handle']['connection data handle size']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['data-handle']['connection data handles currently active']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['data-handle']['connection sweep candidate became referenced']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['data-handle']['connection sweep dhandles closed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['data-handle']['connection sweep dhandles removed from hash list']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['data-handle']['connection sweep time-of-death sets']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['data-handle']['connection sweeps']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['data-handle']['connection sweeps skipped due to checkpoint gathering handles']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['data-handle']['session dhandles swept']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['data-handle']['session sweep attempts']` — type: number; roles: primary, secondary

### wiredTiger.lock (29)

- `db.serverStatus().wiredTiger.lock['checkpoint lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['checkpoint lock application thread wait time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['checkpoint lock internal thread wait time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['dhandle lock application thread time waiting (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['dhandle lock internal thread time waiting (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['dhandle read lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['dhandle write lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['durable timestamp queue lock application thread time waiting (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['durable timestamp queue lock internal thread time waiting (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['durable timestamp queue read lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['durable timestamp queue write lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['metadata lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['metadata lock application thread wait time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['metadata lock internal thread wait time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['read timestamp queue lock application thread time waiting (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['read timestamp queue lock internal thread time waiting (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['read timestamp queue read lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['read timestamp queue write lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['schema lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['schema lock application thread wait time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['schema lock internal thread wait time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['table lock application thread time waiting for the table lock (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['table lock internal thread time waiting for the table lock (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['table read lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['table write lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['txn global lock application thread time waiting (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['txn global lock internal thread time waiting (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['txn global read lock acquisitions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.lock['txn global write lock acquisitions']` — type: number; roles: primary, secondary

### wiredTiger.log (35)

- `db.serverStatus().wiredTiger.log['busy returns attempting to switch slots']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['force archive time sleeping (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['log files manually zero-filled']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['log force write operations']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['log force write operations skipped']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['log records too small to compress']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['log release advances write LSN']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['log server thread advances write LSN']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['log server thread write LSN walk skipped']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['log sync time duration (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['log sync_dir time duration (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['logging bytes consolidated']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['maximum log file size']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['number of pre-allocated log files to create']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['pre-allocated log files not ready and missed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['pre-allocated log files prepared']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['pre-allocated log files used']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot close lost race']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot close unbuffered waits']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot closures']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot join atomic update races']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot join calls atomic updates raced']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot join calls did not yield']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot join calls found active slot closed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot join calls slept']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot join calls yielded']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot join found active slot closed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot joins yield time (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot transitions unable to find free slot']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['slot unbuffered writes']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['total in-memory size of compressed records']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['total log buffer size']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['total size of compressed records']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['written slots coalesced']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.log['yields waiting for previous log file close']` — type: number; roles: primary, secondary

### wiredTiger.oplog (2)

- `db.serverStatus().wiredTiger.oplog['visibility timestamp'][0]` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.oplog['visibility timestamp'][1]` — type: number; roles: primary, secondary

### wiredTiger.perf (22)

- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 1) - 10-49ms']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 2) - 50-99ms']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 3) - 100-249ms']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 4) - 250-499ms']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 5) - 500-999ms']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 6) - 1000ms+']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 1) - 10-49ms']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 2) - 50-99ms']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 3) - 100-249ms']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 4) - 250-499ms']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 5) - 500-999ms']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 6) - 1000ms+']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation read latency histogram (bucket 1) - 100-249us']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation read latency histogram (bucket 2) - 250-499us']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation read latency histogram (bucket 3) - 500-999us']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation read latency histogram (bucket 4) - 1000-9999us']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation read latency histogram (bucket 5) - 10000us+']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation write latency histogram (bucket 1) - 100-249us']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation write latency histogram (bucket 2) - 250-499us']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation write latency histogram (bucket 3) - 500-999us']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation write latency histogram (bucket 4) - 1000-9999us']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation write latency histogram (bucket 5) - 10000us+']` — type: number; roles: primary, secondary

### wiredTiger.reconciliation (36)

- `db.serverStatus().wiredTiger.reconciliation['approximate byte size of timestamps in pages written']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['approximate byte size of transaction IDs in pages written']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['fast-path pages deleted']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['leaf-page overflow keys']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['maximum milliseconds spent in a reconciliation call']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['maximum milliseconds spent in building a disk image in a reconciliation']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['maximum milliseconds spent in moving updates to the history store in a reconciliation']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['page reconciliation calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['page reconciliation calls for eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['page reconciliation calls that resulted in values with prepared transaction metadata']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['page reconciliation calls that resulted in values with timestamps']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['page reconciliation calls that resulted in values with transaction ids']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages deleted']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated newest start durable timestamp ']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated newest stop durable timestamp ']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated newest stop timestamp ']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated newest stop transaction ID']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated newest transaction ID ']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated oldest start timestamp ']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated prepare']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one prepare state']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one start durable timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one start timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one start transaction ID']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one stop durable timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one stop timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one stop transaction ID']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['records written including a prepare state']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['records written including a start durable timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['records written including a start timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['records written including a start transaction ID']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['records written including a stop durable timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['records written including a stop timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['records written including a stop transaction ID']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['split bytes currently awaiting free']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.reconciliation['split objects currently awaiting free']` — type: number; roles: primary, secondary

### wiredTiger.session (29)

- `db.serverStatus().wiredTiger.session['attempts to remove a local object and the object is in use']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['flush_tier operation calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['local objects removed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['session query timestamp calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table alter failed calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table alter successful calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table alter triggering checkpoint calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table alter unchanged and skipped']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table compact failed calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table compact failed calls due to cache pressure']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table compact running']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table compact skipped as process would not reduce file size']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table compact successful calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table compact timeout']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table create failed calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table create successful calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table drop failed calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table drop successful calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table rename failed calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table rename successful calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table salvage failed calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table salvage successful calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table truncate failed calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table truncate successful calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table verify failed calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['table verify successful calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['tiered operations dequeued and processed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['tiered operations scheduled']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.session['tiered storage local retention time (secs)']` — type: number; roles: primary, secondary

### wiredTiger.snapshot-window-settings (2)

- `db.serverStatus().wiredTiger['snapshot-window-settings']['min pinned timestamp'][0]` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['snapshot-window-settings']['min pinned timestamp'][1]` — type: number; roles: primary, secondary

### wiredTiger.thread-state (3)

- `db.serverStatus().wiredTiger['thread-state']['active filesystem fsync calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-state']['active filesystem read calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-state']['active filesystem write calls']` — type: number; roles: primary, secondary

### wiredTiger.thread-yield (14)

- `db.serverStatus().wiredTiger['thread-yield']['application thread time evicting (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['application thread time waiting for cache (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['connection close blocked waiting for transaction state stabilization']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['connection close yielded for lsm manager shutdown']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['data handle lock yielded']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['get reference for page index and slot time sleeping (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['page access yielded due to prepare state change']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['page acquire busy blocked']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['page acquire eviction blocked']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['page acquire locked blocked']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['page acquire read blocked']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['page acquire time sleeping (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['page delete rollback time sleeping for state change (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger['thread-yield']['page reconciliation yielded due to child modification']` — type: number; roles: primary, secondary

### wiredTiger.transaction (65)

- `db.serverStatus().wiredTiger.transaction['Number of prepared updates']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['Number of prepared updates committed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['Number of prepared updates repeated on the same key']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['Number of prepared updates rolled back']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['oldest pinned transaction ID rolled back for eviction']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['prepared transactions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['prepared transactions committed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['prepared transactions currently active']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['prepared transactions rolled back']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['query timestamp calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['race to read prepared update retry']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable history store records with stop timestamps older than newer records']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable inconsistent checkpoint']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable keys removed']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable keys restored']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable pages visited']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable restored tombstones from history store']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable restored updates from history store']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable skipping delete rle']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable skipping stable rle']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable sweeping history store keys']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable tree walk skipping pages']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable updates aborted']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['rollback to stable updates removed from history store']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['sessions scanned in each walk of concurrent sessions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['set timestamp calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['set timestamp durable calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['set timestamp durable updates']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['set timestamp oldest calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['set timestamp oldest updates']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['set timestamp stable calls']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['set timestamp stable updates']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint currently running for history store file']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint generation']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint history store file duration (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent duration for gathering all handles (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent duration for gathering applied handles (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent duration for gathering skipped handles (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent handles applied']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent handles skipped']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent handles walked']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent time (msecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint prepare currently running']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint prepare max time (msecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint prepare min time (msecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint prepare most recent time (msecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint prepare total time (msecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint scrub dirty target']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint scrub time (msecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint stop timing stress active']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoints due to obsolete pages']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction checkpoints skipped because database was clean']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction fsync calls for checkpoint after allocating the transaction ID']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction fsync duration for checkpoint after allocating the transaction ID (usecs)']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction range of IDs currently pinned']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction range of IDs currently pinned by a checkpoint']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction range of timestamps currently pinned']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction range of timestamps pinned by a checkpoint']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction range of timestamps pinned by the oldest active read timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction range of timestamps pinned by the oldest timestamp']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction read timestamp of the oldest active reader']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction rollback to stable currently running']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['transaction walk of concurrent sessions']` — type: number; roles: primary, secondary
- `db.serverStatus().wiredTiger.transaction['update conflicts']` — type: number; roles: primary, secondary

## Unmapped string, identity, enum, and textual leaves

### Identity and topology (14)

- `db.serverStatus().$clusterTime.signature.hash` — type: string; roles: primary, secondary
- `db.serverStatus().$gleStats.electionId` — type: string; roles: primary, secondary
- `db.serverStatus().host` — type: string; roles: primary, secondary
- `db.serverStatus().repl.electionId` — type: string; roles: primary
- `db.serverStatus().repl.hosts[0]` — type: string; roles: primary, secondary
- `db.serverStatus().repl.hosts[1]` — type: string; roles: primary, secondary
- `db.serverStatus().repl.hosts[2]` — type: string; roles: primary, secondary
- `db.serverStatus().repl.hosts[3]` — type: string; roles: primary, secondary
- `db.serverStatus().repl.me` — type: string; roles: primary, secondary
- `db.serverStatus().repl.primary` — type: string; roles: primary, secondary
- `db.serverStatus().repl.setName` — type: string; roles: primary, secondary
- `db.serverStatus().repl.topologyVersion.processId` — type: string; roles: primary, secondary
- `db.serverStatus().security.SSLServerSubjectName` — type: string; roles: primary, secondary
- `db.serverStatus().sharding.configsvrConnectionString` — type: string; roles: primary, secondary

### Date or timestamp text (4)

- `db.serverStatus().repl.lastWrite.lastWriteDate` — type: string; roles: primary, secondary
- `db.serverStatus().repl.lastWrite.majorityWriteDate` — type: string; roles: primary, secondary
- `db.serverStatus().wiredTiger['snapshot-window-settings']['latest majority snapshot timestamp available']` — type: string; roles: primary, secondary
- `db.serverStatus().wiredTiger['snapshot-window-settings']['oldest majority snapshot timestamp available']` — type: string; roles: primary, secondary

### Opaque diagnostic text (3)

- `db.serverStatus().extra_info.note` — type: string; roles: primary, secondary
- `db.serverStatus().metrics.repl.executor.networkInterface` — type: string; roles: primary, secondary
- `db.serverStatus().tcmalloc.tcmalloc.formattedString` — type: string; roles: primary, secondary

### Finite enum or info (10)

- `db.serverStatus().metrics.apiVersions['MongoD Async'][0]` — type: string; roles: primary, secondary
- `db.serverStatus().metrics.apiVersions['MongoD Sync'][0]` — type: string; roles: primary, secondary
- `db.serverStatus().metrics.apiVersions['MongoDB Shell'][0]` — type: string; roles: primary, secondary
- `db.serverStatus().metrics.apiVersions.OplogFetcher[0]` — type: string; roles: primary
- `db.serverStatus().metrics.apiVersions[''][0]` — type: string; roles: primary, secondary
- `db.serverStatus().metrics.apiVersions.mongodb_exporter[0]` — type: string; roles: primary, secondary
- `db.serverStatus().metrics.repl.stateTransition.lastStateTransition` — type: string; roles: primary, secondary
- `db.serverStatus().oplogTruncation.processingMethod` — type: string; roles: primary, secondary
- `db.serverStatus().process` — type: string; roles: primary, secondary
- `db.serverStatus().wiredTiger.uri` — type: string; roles: primary, secondary

## Legacy-only coverage

### asserts.msg (1)

- `db.serverStatus().asserts.msg` — legacy: mongodb_asserts_total

### asserts.regular (1)

- `db.serverStatus().asserts.regular` — legacy: mongodb_asserts_total

### asserts.rollovers (1)

- `db.serverStatus().asserts.rollovers` — legacy: mongodb_asserts_total

### asserts.user (1)

- `db.serverStatus().asserts.user` — legacy: mongodb_asserts_total

### asserts.warning (1)

- `db.serverStatus().asserts.warning` — legacy: mongodb_asserts_total

### connections.active (1)

- `db.serverStatus().connections.active` — legacy: mongodb_connections

### connections.available (1)

- `db.serverStatus().connections.available` — legacy: mongodb_connections

### connections.current (1)

- `db.serverStatus().connections.current` — legacy: mongodb_connections

### connections.totalCreated (1)

- `db.serverStatus().connections.totalCreated` — legacy: mongodb_connections_metrics_created_total

### extra_info.page_faults (1)

- `db.serverStatus().extra_info.page_faults` — legacy: mongodb_extra_info_page_faults_total

### globalLock.activeClients (2)

- `db.serverStatus().globalLock.activeClients.readers` — legacy: mongodb_mongod_global_lock_client
- `db.serverStatus().globalLock.activeClients.writers` — legacy: mongodb_mongod_global_lock_client

### globalLock.currentQueue (2)

- `db.serverStatus().globalLock.currentQueue.readers` — legacy: mongodb_mongod_global_lock_current_queue
- `db.serverStatus().globalLock.currentQueue.writers` — legacy: mongodb_mongod_global_lock_current_queue

### localTime (1)

- `db.serverStatus().localTime` — legacy: mongodb_instance_local_time

### locks.Collection (1)

- `db.serverStatus().locks.Collection.timeAcquiringMicros.w` — legacy: mongodb_mongod_locks_time_acquiring_global_microseconds_total

### locks.Database (1)

- `db.serverStatus().locks.Database.timeAcquiringMicros.w` — legacy: mongodb_mongod_locks_time_acquiring_global_microseconds_total

### locks.Global (1)

- `db.serverStatus().locks.Global.timeAcquiringMicros.r` — legacy: mongodb_mongod_locks_time_acquiring_global_microseconds_total

### locks.Mutex (1)

- `db.serverStatus().locks.Mutex.timeAcquiringMicros.r` — legacy: mongodb_mongod_locks_time_acquiring_global_microseconds_total

### locks.ParallelBatchWriterMode (1)

- `db.serverStatus().locks.ParallelBatchWriterMode.timeAcquiringMicros.r` — legacy: mongodb_mongod_locks_time_acquiring_global_microseconds_total

### locks.ReplicationStateTransition (1)

- `db.serverStatus().locks.ReplicationStateTransition.timeAcquiringMicros.w` — legacy: mongodb_mongod_locks_time_acquiring_global_microseconds_total

### mem.resident (1)

- `db.serverStatus().mem.resident` — legacy: mongodb_memory

### mem.virtual (1)

- `db.serverStatus().mem.virtual` — legacy: mongodb_memory

### metrics.document (4)

- `db.serverStatus().metrics.document.deleted` — legacy: mongodb_mongod_metrics_document_total
- `db.serverStatus().metrics.document.inserted` — legacy: mongodb_mongod_metrics_document_total
- `db.serverStatus().metrics.document.returned` — legacy: mongodb_mongod_metrics_document_total
- `db.serverStatus().metrics.document.updated` — legacy: mongodb_mongod_metrics_document_total

### metrics.getLastError (3)

- `db.serverStatus().metrics.getLastError.wtime.num` — legacy: mongodb_mongod_metrics_get_last_error_wtime_num_total
- `db.serverStatus().metrics.getLastError.wtime.totalMillis` — legacy: mongodb_mongod_metrics_get_last_error_wtime_total_milliseconds
- `db.serverStatus().metrics.getLastError.wtimeouts` — legacy: mongodb_mongod_metrics_get_last_error_wtimeouts_total

### metrics.operation (1)

- `db.serverStatus().metrics.operation.scanAndOrder` — legacy: mongodb_mongod_metrics_operation_total

### metrics.queryExecutor (2)

- `db.serverStatus().metrics.queryExecutor.scanned` — legacy: mongodb_mongod_metrics_query_executor_total
- `db.serverStatus().metrics.queryExecutor.scannedObjects` — legacy: mongodb_mongod_metrics_query_executor_total

### metrics.record (1)

- `db.serverStatus().metrics.record.moves` — legacy: mongodb_mongod_metrics_record_moves_total

### metrics.repl (14)

- `db.serverStatus().metrics.repl.apply.batches.num` — legacy: mongodb_mongod_metrics_repl_apply_batches_num_total
- `db.serverStatus().metrics.repl.apply.batches.totalMillis` — legacy: mongodb_mongod_metrics_repl_apply_batches_total_milliseconds
- `db.serverStatus().metrics.repl.apply.ops` — legacy: mongodb_mongod_metrics_repl_apply_ops_total
- `db.serverStatus().metrics.repl.buffer.count` — legacy: mongodb_mongod_metrics_repl_buffer_count
- `db.serverStatus().metrics.repl.buffer.maxSizeBytes` — legacy: mongodb_mongod_metrics_repl_buffer_max_size_bytes
- `db.serverStatus().metrics.repl.buffer.sizeBytes` — legacy: mongodb_mongod_metrics_repl_buffer_size_bytes
- `db.serverStatus().metrics.repl.executor.queues.networkInProgress` — legacy: mongodb_mongod_metrics_repl_executor_queue
- `db.serverStatus().metrics.repl.executor.queues.sleepers` — legacy: mongodb_mongod_metrics_repl_executor_queue
- `db.serverStatus().metrics.repl.executor.unsignaledEvents` — legacy: mongodb_mongod_metrics_repl_executor_unsignaled_events
- `db.serverStatus().metrics.repl.network.bytes` — legacy: mongodb_mongod_metrics_repl_network_bytes_total
- `db.serverStatus().metrics.repl.network.getmores.num` — legacy: mongodb_mongod_metrics_repl_network_getmores_num_total
- `db.serverStatus().metrics.repl.network.getmores.totalMillis` — legacy: mongodb_mongod_metrics_repl_network_getmores_total_milliseconds
- `db.serverStatus().metrics.repl.network.ops` — legacy: mongodb_mongod_metrics_repl_network_ops_total
- `db.serverStatus().metrics.repl.network.readersCreated` — legacy: mongodb_mongod_metrics_repl_network_readers_created_total

### metrics.ttl (2)

- `db.serverStatus().metrics.ttl.deletedDocuments` — legacy: mongodb_mongod_metrics_ttl_deleted_documents_total
- `db.serverStatus().metrics.ttl.passes` — legacy: mongodb_mongod_metrics_ttl_passes_total

### network.bytesIn (1)

- `db.serverStatus().network.bytesIn` — legacy: mongodb_network_bytes_total

### network.bytesOut (1)

- `db.serverStatus().network.bytesOut` — legacy: mongodb_network_bytes_total

### network.numRequests (1)

- `db.serverStatus().network.numRequests` — legacy: mongodb_network_metrics_num_requests_total

### opcounters.command (1)

- `db.serverStatus().opcounters.command` — legacy: mongodb_op_counters_total

### opcounters.delete (1)

- `db.serverStatus().opcounters.delete` — legacy: mongodb_op_counters_total

### opcounters.getmore (1)

- `db.serverStatus().opcounters.getmore` — legacy: mongodb_op_counters_total

### opcounters.insert (1)

- `db.serverStatus().opcounters.insert` — legacy: mongodb_op_counters_total

### opcounters.query (1)

- `db.serverStatus().opcounters.query` — legacy: mongodb_op_counters_total

### opcounters.update (1)

- `db.serverStatus().opcounters.update` — legacy: mongodb_op_counters_total

### opcountersRepl.command (1)

- `db.serverStatus().opcountersRepl.command` — legacy: mongodb_op_counters_repl_total

### opcountersRepl.delete (1)

- `db.serverStatus().opcountersRepl.delete` — legacy: mongodb_op_counters_repl_total

### opcountersRepl.getmore (1)

- `db.serverStatus().opcountersRepl.getmore` — legacy: mongodb_op_counters_repl_total

### opcountersRepl.insert (1)

- `db.serverStatus().opcountersRepl.insert` — legacy: mongodb_op_counters_repl_total

### opcountersRepl.query (1)

- `db.serverStatus().opcountersRepl.query` — legacy: mongodb_op_counters_repl_total

### opcountersRepl.update (1)

- `db.serverStatus().opcountersRepl.update` — legacy: mongodb_op_counters_repl_total

### opLatencies.commands (2)

- `db.serverStatus().opLatencies.commands.latency` — legacy: mongodb_mongod_op_latencies_latency_total
- `db.serverStatus().opLatencies.commands.ops` — legacy: mongodb_mongod_op_latencies_ops_total

### opLatencies.reads (2)

- `db.serverStatus().opLatencies.reads.latency` — legacy: mongodb_mongod_op_latencies_latency_total
- `db.serverStatus().opLatencies.reads.ops` — legacy: mongodb_mongod_op_latencies_ops_total

### opLatencies.writes (2)

- `db.serverStatus().opLatencies.writes.latency` — legacy: mongodb_mongod_op_latencies_latency_total
- `db.serverStatus().opLatencies.writes.ops` — legacy: mongodb_mongod_op_latencies_ops_total

### storageEngine.name (1)

- `db.serverStatus().storageEngine.name` — legacy: mongodb_mongod_storage_engine

### tcmalloc.generic (2)

- `db.serverStatus().tcmalloc.generic.current_allocated_bytes` — legacy: mongodb_tcmalloc_generic_heap
- `db.serverStatus().tcmalloc.generic.heap_size` — legacy: mongodb_tcmalloc_generic_heap

### tcmalloc.tcmalloc (17)

- `db.serverStatus().tcmalloc.tcmalloc.aggressive_memory_decommit` — legacy: mongodb_tcmalloc_aggressive_memory_decommit
- `db.serverStatus().tcmalloc.tcmalloc.central_cache_free_bytes` — legacy: mongodb_tcmalloc_cache_bytes
- `db.serverStatus().tcmalloc.tcmalloc.current_total_thread_cache_bytes` — legacy: mongodb_tcmalloc_cache_bytes
- `db.serverStatus().tcmalloc.tcmalloc.max_total_thread_cache_bytes` — legacy: mongodb_tcmalloc_cache_bytes
- `db.serverStatus().tcmalloc.tcmalloc.pageheap_commit_count` — legacy: mongodb_tcmalloc_pageheap_count
- `db.serverStatus().tcmalloc.tcmalloc.pageheap_committed_bytes` — legacy: mongodb_tcmalloc_pageheap_bytes
- `db.serverStatus().tcmalloc.tcmalloc.pageheap_decommit_count` — legacy: mongodb_tcmalloc_pageheap_count
- `db.serverStatus().tcmalloc.tcmalloc.pageheap_free_bytes` — legacy: mongodb_tcmalloc_pageheap_bytes
- `db.serverStatus().tcmalloc.tcmalloc.pageheap_reserve_count` — legacy: mongodb_tcmalloc_pageheap_count
- `db.serverStatus().tcmalloc.tcmalloc.pageheap_scavenge_count` — legacy: mongodb_tcmalloc_pageheap_count
- `db.serverStatus().tcmalloc.tcmalloc.pageheap_total_commit_bytes` — legacy: mongodb_tcmalloc_pageheap_bytes
- `db.serverStatus().tcmalloc.tcmalloc.pageheap_total_decommit_bytes` — legacy: mongodb_tcmalloc_pageheap_bytes
- `db.serverStatus().tcmalloc.tcmalloc.pageheap_total_reserve_bytes` — legacy: mongodb_tcmalloc_pageheap_bytes
- `db.serverStatus().tcmalloc.tcmalloc.pageheap_unmapped_bytes` — legacy: mongodb_tcmalloc_pageheap_bytes
- `db.serverStatus().tcmalloc.tcmalloc.thread_cache_free_bytes` — legacy: mongodb_tcmalloc_cache_bytes
- `db.serverStatus().tcmalloc.tcmalloc.total_free_bytes` — legacy: mongodb_tcmalloc_free_bytes
- `db.serverStatus().tcmalloc.tcmalloc.transfer_cache_free_bytes` — legacy: mongodb_tcmalloc_cache_bytes

### uptime (1)

- `db.serverStatus().uptime` — legacy: mongodb_instance_uptime_estimate_seconds, mongodb_instance_uptime_seconds

### version (1)

- `db.serverStatus().version` — legacy: mongodb_version_info

### wiredTiger.block-manager (7)

- `db.serverStatus().wiredTiger['block-manager']['blocks pre-loaded']` — legacy: mongodb_mongod_wiredtiger_blockmanager_blocks_total
- `db.serverStatus().wiredTiger['block-manager']['blocks read']` — legacy: mongodb_mongod_wiredtiger_blockmanager_blocks_total
- `db.serverStatus().wiredTiger['block-manager']['blocks written']` — legacy: mongodb_mongod_wiredtiger_blockmanager_blocks_total
- `db.serverStatus().wiredTiger['block-manager']['bytes read']` — legacy: mongodb_mongod_wiredtiger_blockmanager_bytes_total
- `db.serverStatus().wiredTiger['block-manager']['bytes written']` — legacy: mongodb_mongod_wiredtiger_blockmanager_bytes_total
- `db.serverStatus().wiredTiger['block-manager']['mapped blocks read']` — legacy: mongodb_mongod_wiredtiger_blockmanager_blocks_total
- `db.serverStatus().wiredTiger['block-manager']['mapped bytes read']` — legacy: mongodb_mongod_wiredtiger_blockmanager_bytes_total

### wiredTiger.cache (14)

- `db.serverStatus().wiredTiger.cache['bytes currently in the cache']` — legacy: mongodb_mongod_wiredtiger_cache_bytes
- `db.serverStatus().wiredTiger.cache['bytes read into cache']` — legacy: mongodb_mongod_wiredtiger_cache_bytes_total
- `db.serverStatus().wiredTiger.cache['bytes written from cache']` — legacy: mongodb_mongod_wiredtiger_cache_bytes_total
- `db.serverStatus().wiredTiger.cache['maximum bytes configured']` — legacy: mongodb_mongod_wiredtiger_cache_max_bytes
- `db.serverStatus().wiredTiger.cache['modified pages evicted']` — legacy: mongodb_mongod_wiredtiger_cache_evicted_total
- `db.serverStatus().wiredTiger.cache['pages currently held in the cache']` — legacy: mongodb_mongod_wiredtiger_cache_pages
- `db.serverStatus().wiredTiger.cache['pages read into cache']` — legacy: mongodb_mongod_wiredtiger_cache_pages_total
- `db.serverStatus().wiredTiger.cache['pages written from cache']` — legacy: mongodb_mongod_wiredtiger_cache_pages_total
- `db.serverStatus().wiredTiger.cache['percentage overhead']` — legacy: mongodb_mongod_wiredtiger_cache_overhead_percent
- `db.serverStatus().wiredTiger.cache['tracked bytes belonging to internal pages in the cache']` — legacy: mongodb_mongod_wiredtiger_cache_bytes
- `db.serverStatus().wiredTiger.cache['tracked bytes belonging to leaf pages in the cache']` — legacy: mongodb_mongod_wiredtiger_cache_bytes
- `db.serverStatus().wiredTiger.cache['tracked dirty bytes in the cache']` — legacy: mongodb_mongod_wiredtiger_cache_bytes
- `db.serverStatus().wiredTiger.cache['tracked dirty pages in the cache']` — legacy: mongodb_mongod_wiredtiger_cache_pages
- `db.serverStatus().wiredTiger.cache['unmodified pages evicted']` — legacy: mongodb_mongod_wiredtiger_cache_evicted_total

### wiredTiger.log (11)

- `db.serverStatus().wiredTiger.log['log bytes of payload data']` — legacy: mongodb_mongod_wiredtiger_log_bytes_total
- `db.serverStatus().wiredTiger.log['log bytes written']` — legacy: mongodb_mongod_wiredtiger_log_bytes_total
- `db.serverStatus().wiredTiger.log['log flush operations']` — legacy: mongodb_mongod_wiredtiger_log_operations_total
- `db.serverStatus().wiredTiger.log['log records compressed']` — legacy: mongodb_mongod_wiredtiger_log_records_total
- `db.serverStatus().wiredTiger.log['log records not compressed']` — legacy: mongodb_mongod_wiredtiger_log_records_total
- `db.serverStatus().wiredTiger.log['log scan operations']` — legacy: mongodb_mongod_wiredtiger_log_operations_total
- `db.serverStatus().wiredTiger.log['log scan records requiring two reads']` — legacy: mongodb_mongod_wiredtiger_log_operations_total
- `db.serverStatus().wiredTiger.log['log sync operations']` — legacy: mongodb_mongod_wiredtiger_log_operations_total
- `db.serverStatus().wiredTiger.log['log sync_dir operations']` — legacy: mongodb_mongod_wiredtiger_log_operations_total
- `db.serverStatus().wiredTiger.log['log write operations']` — legacy: mongodb_mongod_wiredtiger_log_operations_total
- `db.serverStatus().wiredTiger.log['records processed by log scan']` — legacy: mongodb_mongod_wiredtiger_log_records_scanned_total

### wiredTiger.session (1)

- `db.serverStatus().wiredTiger.session['open session count']` — legacy: mongodb_mongod_wiredtiger_session_open_sessions_total

### wiredTiger.transaction (8)

- `db.serverStatus().wiredTiger.transaction['transaction begins']` — legacy: mongodb_mongod_wiredtiger_transactions_total
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint currently running']` — legacy: mongodb_mongod_wiredtiger_transactions_running_checkpoints
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint max time (msecs)']` — legacy: mongodb_mongod_wiredtiger_transactions_checkpoint_milliseconds
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint min time (msecs)']` — legacy: mongodb_mongod_wiredtiger_transactions_checkpoint_milliseconds
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint total time (msecs)']` — legacy: mongodb_mongod_wiredtiger_transactions_checkpoint_milliseconds_total
- `db.serverStatus().wiredTiger.transaction['transaction checkpoints']` — legacy: mongodb_mongod_wiredtiger_transactions_total
- `db.serverStatus().wiredTiger.transaction['transactions committed']` — legacy: mongodb_mongod_wiredtiger_transactions_total
- `db.serverStatus().wiredTiger.transaction['transactions rolled back']` — legacy: mongodb_mongod_wiredtiger_transactions_total

## Modern-only coverage

### electionMetrics.averageCatchUpOps (1)

- `db.serverStatus().electionMetrics.averageCatchUpOps` — modern: mongodb_server_status_election_metrics_average_catch_up_ops__total

### electionMetrics.catchUpTakeover (2)

- `db.serverStatus().electionMetrics.catchUpTakeover.called` — modern: mongodb_server_status_election_calls_total
- `db.serverStatus().electionMetrics.catchUpTakeover.successful` — modern: mongodb_server_status_election_calls_total

### electionMetrics.electionTimeout (2)

- `db.serverStatus().electionMetrics.electionTimeout.called` — modern: mongodb_server_status_election_calls_total
- `db.serverStatus().electionMetrics.electionTimeout.successful` — modern: mongodb_server_status_election_calls_total

### electionMetrics.freezeTimeout (2)

- `db.serverStatus().electionMetrics.freezeTimeout.called` — modern: mongodb_server_status_election_calls_total
- `db.serverStatus().electionMetrics.freezeTimeout.successful` — modern: mongodb_server_status_election_calls_total

### electionMetrics.numCatchUps (1)

- `db.serverStatus().electionMetrics.numCatchUps` — modern: mongodb_server_status_election_metrics_num_catch_ups__total

### electionMetrics.numCatchUpsAlreadyCaughtUp (1)

- `db.serverStatus().electionMetrics.numCatchUpsAlreadyCaughtUp` — modern: mongodb_server_status_election_metrics_num_catch_ups_already_caught_up__total

### electionMetrics.numCatchUpsFailedWithError (1)

- `db.serverStatus().electionMetrics.numCatchUpsFailedWithError` — modern: mongodb_server_status_election_metrics_num_catch_ups_failed_with_error__total

### electionMetrics.numCatchUpsFailedWithNewTerm (1)

- `db.serverStatus().electionMetrics.numCatchUpsFailedWithNewTerm` — modern: mongodb_server_status_election_metrics_num_catch_ups_failed_with_new_term__total

### electionMetrics.numCatchUpsFailedWithReplSetAbortPrimaryCatchUpCmd (1)

- `db.serverStatus().electionMetrics.numCatchUpsFailedWithReplSetAbortPrimaryCatchUpCmd` — modern: mongodb_server_status_election_metrics_num_catch_ups_failed_with_repl_set_abort_primary_catch_up_cmd__total

### electionMetrics.numCatchUpsSkipped (1)

- `db.serverStatus().electionMetrics.numCatchUpsSkipped` — modern: mongodb_server_status_election_metrics_num_catch_ups_skipped__total

### electionMetrics.numCatchUpsSucceeded (1)

- `db.serverStatus().electionMetrics.numCatchUpsSucceeded` — modern: mongodb_server_status_election_metrics_num_catch_ups_succeeded__total

### electionMetrics.numCatchUpsTimedOut (1)

- `db.serverStatus().electionMetrics.numCatchUpsTimedOut` — modern: mongodb_server_status_election_metrics_num_catch_ups_timed_out__total

### electionMetrics.numStepDownsCausedByHigherTerm (1)

- `db.serverStatus().electionMetrics.numStepDownsCausedByHigherTerm` — modern: mongodb_server_status_election_metrics_num_step_downs_caused_by_higher_term__total

### electionMetrics.priorityTakeover (2)

- `db.serverStatus().electionMetrics.priorityTakeover.called` — modern: mongodb_server_status_election_calls_total
- `db.serverStatus().electionMetrics.priorityTakeover.successful` — modern: mongodb_server_status_election_calls_total

### electionMetrics.stepUpCmd (2)

- `db.serverStatus().electionMetrics.stepUpCmd.called` — modern: mongodb_server_status_election_calls_total
- `db.serverStatus().electionMetrics.stepUpCmd.successful` — modern: mongodb_server_status_election_calls_total

### flowControl.enabled (1)

- `db.serverStatus().flowControl.enabled` — modern: mongodb_server_status_flow_control_enabled

### flowControl.isLagged (1)

- `db.serverStatus().flowControl.isLagged` — modern: mongodb_server_status_flow_control_is_lagged

### flowControl.isLaggedCount (1)

- `db.serverStatus().flowControl.isLaggedCount` — modern: mongodb_server_status_flow_control_is_lagged_count_total

### flowControl.isLaggedTimeMicros (1)

- `db.serverStatus().flowControl.isLaggedTimeMicros` — modern: mongodb_server_status_flow_control_is_lagged_time_seconds

### flowControl.locksPerKiloOp (1)

- `db.serverStatus().flowControl.locksPerKiloOp` — modern: mongodb_server_status_flow_control_locks_per_kilo_op

### flowControl.sustainerRate (1)

- `db.serverStatus().flowControl.sustainerRate` — modern: mongodb_server_status_flow_control_sustainer_rate

### flowControl.targetRateLimit (1)

- `db.serverStatus().flowControl.targetRateLimit` — modern: mongodb_server_status_flow_control_target_rate_limit

### flowControl.timeAcquiringMicros (1)

- `db.serverStatus().flowControl.timeAcquiringMicros` — modern: mongodb_server_status_flow_control_time_acquiring_seconds

### logicalSessionRecordCache.activeSessionsCount (1)

- `db.serverStatus().logicalSessionRecordCache.activeSessionsCount` — modern: mongodb_server_status_logical_session_record_cache_active_sessions_count

### logicalSessionRecordCache.lastSessionsCollectionJobCursorsClosed (1)

- `db.serverStatus().logicalSessionRecordCache.lastSessionsCollectionJobCursorsClosed` — modern: mongodb_server_status_logical_session_record_cache_last_sessions_collection_job_cursors_closed

### logicalSessionRecordCache.lastSessionsCollectionJobDurationMillis (1)

- `db.serverStatus().logicalSessionRecordCache.lastSessionsCollectionJobDurationMillis` — modern: mongodb_server_status_logical_session_record_cache_last_sessions_collection_job_duration_seconds

### logicalSessionRecordCache.lastSessionsCollectionJobEntriesEnded (1)

- `db.serverStatus().logicalSessionRecordCache.lastSessionsCollectionJobEntriesEnded` — modern: mongodb_server_status_logical_session_record_cache_last_sessions_collection_job_entries_ended

### logicalSessionRecordCache.lastSessionsCollectionJobEntriesRefreshed (1)

- `db.serverStatus().logicalSessionRecordCache.lastSessionsCollectionJobEntriesRefreshed` — modern: mongodb_server_status_logical_session_record_cache_last_sessions_collection_job_entries_refreshed

### logicalSessionRecordCache.lastSessionsCollectionJobTimestamp (1)

- `db.serverStatus().logicalSessionRecordCache.lastSessionsCollectionJobTimestamp` — modern: mongodb_server_status_logical_session_record_cache_last_sessions_collection_job_timestamp_seconds

### logicalSessionRecordCache.lastTransactionReaperJobDurationMillis (1)

- `db.serverStatus().logicalSessionRecordCache.lastTransactionReaperJobDurationMillis` — modern: mongodb_server_status_logical_session_record_cache_last_transaction_reaper_job_duration_seconds

### logicalSessionRecordCache.lastTransactionReaperJobEntriesCleanedUp (1)

- `db.serverStatus().logicalSessionRecordCache.lastTransactionReaperJobEntriesCleanedUp` — modern: mongodb_server_status_logical_session_record_cache_last_transaction_reaper_job_entries_cleaned_up

### logicalSessionRecordCache.lastTransactionReaperJobTimestamp (1)

- `db.serverStatus().logicalSessionRecordCache.lastTransactionReaperJobTimestamp` — modern: mongodb_server_status_logical_session_record_cache_last_transaction_reaper_job_timestamp_seconds

### logicalSessionRecordCache.sessionCatalogSize (1)

- `db.serverStatus().logicalSessionRecordCache.sessionCatalogSize` — modern: mongodb_server_status_logical_session_record_cache_session_catalog_size

### logicalSessionRecordCache.sessionsCollectionJobCount (1)

- `db.serverStatus().logicalSessionRecordCache.sessionsCollectionJobCount` — modern: mongodb_server_status_logical_session_record_cache_sessions_collection_job_count_total

### logicalSessionRecordCache.transactionReaperJobCount (1)

- `db.serverStatus().logicalSessionRecordCache.transactionReaperJobCount` — modern: mongodb_server_status_logical_session_record_cache_transaction_reaper_job_count_total

### metrics.aggStageCounters (10)

- `db.serverStatus().metrics.aggStageCounters.$group` — modern: mongodb_server_status_aggregation_stage_total
- `db.serverStatus().metrics.aggStageCounters.$indexStats` — modern: mongodb_server_status_aggregation_stage_total
- `db.serverStatus().metrics.aggStageCounters.$limit` — modern: mongodb_server_status_aggregation_stage_total
- `db.serverStatus().metrics.aggStageCounters.$lookup` — modern: mongodb_server_status_aggregation_stage_total
- `db.serverStatus().metrics.aggStageCounters.$match` — modern: mongodb_server_status_aggregation_stage_total
- `db.serverStatus().metrics.aggStageCounters.$project` — modern: mongodb_server_status_aggregation_stage_total
- `db.serverStatus().metrics.aggStageCounters.$sample` — modern: mongodb_server_status_aggregation_stage_total
- `db.serverStatus().metrics.aggStageCounters.$set` — modern: mongodb_server_status_aggregation_stage_total
- `db.serverStatus().metrics.aggStageCounters.$sort` — modern: mongodb_server_status_aggregation_stage_total
- `db.serverStatus().metrics.aggStageCounters.$unwind` — modern: mongodb_server_status_aggregation_stage_total

### metrics.commands (18)

- `db.serverStatus().metrics.commands.aggregate.failed` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.aggregate.total` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.delete.failed` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.delete.total` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.find.failed` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.find.total` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.getMore.failed` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.getMore.total` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.insert.failed` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.insert.total` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.moveChunk.failed` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.moveChunk.total` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.rotateCertificates.failed` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.rotateCertificates.total` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.serverStatus.failed` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.serverStatus.total` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.update.failed` — modern: mongodb_server_status_commands_total
- `db.serverStatus().metrics.commands.update.total` — modern: mongodb_server_status_commands_total

### metrics.cursor (9)

- `db.serverStatus().metrics.cursor.lifespan.greaterThanOrEqual10Minutes` — modern: mongodb_server_status_metrics_cursor_lifespan_greater_than_or_equal10_minutes
- `db.serverStatus().metrics.cursor.lifespan.lessThan10Minutes` — modern: mongodb_server_status_metrics_cursor_lifespan_less_than10_minutes
- `db.serverStatus().metrics.cursor.lifespan.lessThan15Seconds` — modern: mongodb_server_status_metrics_cursor_lifespan_less_than15_seconds
- `db.serverStatus().metrics.cursor.lifespan.lessThan1Minute` — modern: mongodb_server_status_metrics_cursor_lifespan_less_than1_minute
- `db.serverStatus().metrics.cursor.lifespan.lessThan1Second` — modern: mongodb_server_status_metrics_cursor_lifespan_less_than1_second
- `db.serverStatus().metrics.cursor.lifespan.lessThan30Seconds` — modern: mongodb_server_status_metrics_cursor_lifespan_less_than30_seconds
- `db.serverStatus().metrics.cursor.lifespan.lessThan5Seconds` — modern: mongodb_server_status_metrics_cursor_lifespan_less_than5_seconds
- `db.serverStatus().metrics.cursor.moreThanOneBatch` — modern: mongodb_server_status_metrics_cursor_more_than_one_batch
- `db.serverStatus().metrics.cursor.totalOpened` — modern: mongodb_server_status_metrics_cursor_total_opened_total

### metrics.operatorCounters (10)

- `db.serverStatus().metrics.operatorCounters.expressions.$add` — modern: mongodb_server_status_operator_total
- `db.serverStatus().metrics.operatorCounters.expressions.$eq` — modern: mongodb_server_status_operator_total
- `db.serverStatus().metrics.operatorCounters.expressions.$gt` — modern: mongodb_server_status_operator_total
- `db.serverStatus().metrics.operatorCounters.expressions.$ifNull` — modern: mongodb_server_status_operator_total
- `db.serverStatus().metrics.operatorCounters.expressions.$in` — modern: mongodb_server_status_operator_total
- `db.serverStatus().metrics.operatorCounters.expressions.$multiply` — modern: mongodb_server_status_operator_total
- `db.serverStatus().metrics.operatorCounters.expressions.$reduce` — modern: mongodb_server_status_operator_total
- `db.serverStatus().metrics.operatorCounters.expressions.$size` — modern: mongodb_server_status_operator_total
- `db.serverStatus().metrics.operatorCounters.expressions.$subtract` — modern: mongodb_server_status_operator_total
- `db.serverStatus().metrics.operatorCounters.match.$eq` — modern: mongodb_server_status_operator_total

### metrics.query (10)

- `db.serverStatus().metrics.query.deleteManyCount` — modern: mongodb_server_status_metrics_query_delete_many_count_total
- `db.serverStatus().metrics.query.externalRetryableWriteCount` — modern: mongodb_server_status_metrics_query_external_retryable_write_count_total
- `db.serverStatus().metrics.query.internalRetryableWriteCount` — modern: mongodb_server_status_metrics_query_internal_retryable_write_count_total
- `db.serverStatus().metrics.query.planCacheTotalSizeEstimateBytes` — modern: mongodb_server_status_metrics_query_plan_cache_total_size_estimate_bytes
- `db.serverStatus().metrics.query.updateDeleteManyDocumentsMaxCount` — modern: mongodb_server_status_metrics_query_update_delete_many_documents_max_count_total
- `db.serverStatus().metrics.query.updateDeleteManyDocumentsTotalCount` — modern: mongodb_server_status_metrics_query_update_delete_many_documents_total_count_total
- `db.serverStatus().metrics.query.updateDeleteManyDurationMaxMs` — modern: mongodb_server_status_metrics_query_update_delete_many_duration_max_ms
- `db.serverStatus().metrics.query.updateDeleteManyDurationTotalMs` — modern: mongodb_server_status_metrics_query_update_delete_many_duration_total_ms_total
- `db.serverStatus().metrics.query.updateManyCount` — modern: mongodb_server_status_metrics_query_update_many_count_total
- `db.serverStatus().metrics.query.updateOneOpStyleBroadcastWithExactIDCount` — modern: mongodb_server_status_metrics_query_update_one_op_style_broadcast_with_exact_idcount_total

### network.compression (8)

- `db.serverStatus().network.compression.snappy.compressor.bytesIn` — modern: mongodb_server_status_network_compression_bytes_total
- `db.serverStatus().network.compression.snappy.compressor.bytesOut` — modern: mongodb_server_status_network_compression_bytes_total
- `db.serverStatus().network.compression.snappy.decompressor.bytesIn` — modern: mongodb_server_status_network_compression_bytes_total
- `db.serverStatus().network.compression.snappy.decompressor.bytesOut` — modern: mongodb_server_status_network_compression_bytes_total
- `db.serverStatus().network.compression.zstd.compressor.bytesIn` — modern: mongodb_server_status_network_compression_bytes_total
- `db.serverStatus().network.compression.zstd.compressor.bytesOut` — modern: mongodb_server_status_network_compression_bytes_total
- `db.serverStatus().network.compression.zstd.decompressor.bytesIn` — modern: mongodb_server_status_network_compression_bytes_total
- `db.serverStatus().network.compression.zstd.decompressor.bytesOut` — modern: mongodb_server_status_network_compression_bytes_total

### network.numSlowDNSOperations (1)

- `db.serverStatus().network.numSlowDNSOperations` — modern: mongodb_server_status_network_num_slow_dnsoperations_total

### network.numSlowSSLOperations (1)

- `db.serverStatus().network.numSlowSSLOperations` — modern: mongodb_server_status_network_num_slow_ssloperations_total

### network.physicalBytesIn (1)

- `db.serverStatus().network.physicalBytesIn` — modern: mongodb_server_status_network_physical_bytes_in_total

### network.physicalBytesOut (1)

- `db.serverStatus().network.physicalBytesOut` — modern: mongodb_server_status_network_physical_bytes_out_total

### network.serviceExecutors (8)

- `db.serverStatus().network.serviceExecutors.fixed.clientsInTotal` — modern: mongodb_server_status_network_service_executor
- `db.serverStatus().network.serviceExecutors.fixed.clientsRunning` — modern: mongodb_server_status_network_service_executor
- `db.serverStatus().network.serviceExecutors.fixed.clientsWaitingForData` — modern: mongodb_server_status_network_service_executor
- `db.serverStatus().network.serviceExecutors.fixed.threadsRunning` — modern: mongodb_server_status_network_service_executor
- `db.serverStatus().network.serviceExecutors.passthrough.clientsInTotal` — modern: mongodb_server_status_network_service_executor
- `db.serverStatus().network.serviceExecutors.passthrough.clientsRunning` — modern: mongodb_server_status_network_service_executor
- `db.serverStatus().network.serviceExecutors.passthrough.clientsWaitingForData` — modern: mongodb_server_status_network_service_executor
- `db.serverStatus().network.serviceExecutors.passthrough.threadsRunning` — modern: mongodb_server_status_network_service_executor

### network.tcpFastOpen (4)

- `db.serverStatus().network.tcpFastOpen.accepted` — modern: mongodb_server_status_network_tcp_fast_open_accepted_total
- `db.serverStatus().network.tcpFastOpen.clientSupported` — modern: mongodb_server_status_network_tcp_fast_open_client_supported
- `db.serverStatus().network.tcpFastOpen.kernelSetting` — modern: mongodb_server_status_network_tcp_fast_open_kernel_setting
- `db.serverStatus().network.tcpFastOpen.serverSupported` — modern: mongodb_server_status_network_tcp_fast_open_server_supported

### oplogTruncation.totalTimeProcessingMicros (1)

- `db.serverStatus().oplogTruncation.totalTimeProcessingMicros` — modern: mongodb_server_status_oplog_truncation_total_time_processing_seconds

### oplogTruncation.totalTimeTruncatingMicros (1)

- `db.serverStatus().oplogTruncation.totalTimeTruncatingMicros` — modern: mongodb_server_status_oplog_truncation_total_time_truncating_seconds

### oplogTruncation.truncateCount (1)

- `db.serverStatus().oplogTruncation.truncateCount` — modern: mongodb_server_status_oplog_truncation_truncate_count_total

### readConcernCounters.nonTransactionOps (12)

- `db.serverStatus().readConcernCounters.nonTransactionOps.available` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_available_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.linearizable` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_linearizable_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.local` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_local_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.majority` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_majority_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.none` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_none_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.noneInfo.CWRC.available` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_none_info_cwrc_available_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.noneInfo.CWRC.local` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_none_info_cwrc_local_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.noneInfo.CWRC.majority` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_none_info_cwrc_majority_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.noneInfo.implicitDefault.available` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_none_info_implicit_default_available_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.noneInfo.implicitDefault.local` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_none_info_implicit_default_local_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.snapshot.withClusterTime` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_snapshot_with_cluster_time_total
- `db.serverStatus().readConcernCounters.nonTransactionOps.snapshot.withoutClusterTime` — modern: mongodb_server_status_read_concern_counters_non_transaction_ops_snapshot_without_cluster_time_total

### readConcernCounters.transactionOps (8)

- `db.serverStatus().readConcernCounters.transactionOps.local` — modern: mongodb_server_status_read_concern_counters_transaction_ops_local_total
- `db.serverStatus().readConcernCounters.transactionOps.majority` — modern: mongodb_server_status_read_concern_counters_transaction_ops_majority_total
- `db.serverStatus().readConcernCounters.transactionOps.none` — modern: mongodb_server_status_read_concern_counters_transaction_ops_none_total
- `db.serverStatus().readConcernCounters.transactionOps.noneInfo.CWRC.local` — modern: mongodb_server_status_read_concern_counters_transaction_ops_none_info_cwrc_local_total
- `db.serverStatus().readConcernCounters.transactionOps.noneInfo.CWRC.majority` — modern: mongodb_server_status_read_concern_counters_transaction_ops_none_info_cwrc_majority_total
- `db.serverStatus().readConcernCounters.transactionOps.noneInfo.implicitDefault.local` — modern: mongodb_server_status_read_concern_counters_transaction_ops_none_info_implicit_default_local_total
- `db.serverStatus().readConcernCounters.transactionOps.snapshot.withClusterTime` — modern: mongodb_server_status_read_concern_counters_transaction_ops_snapshot_with_cluster_time_total
- `db.serverStatus().readConcernCounters.transactionOps.snapshot.withoutClusterTime` — modern: mongodb_server_status_read_concern_counters_transaction_ops_snapshot_without_cluster_time_total

### repl.isImplicitDefaultMajorityWC (1)

- `db.serverStatus().repl.isImplicitDefaultMajorityWC` — modern: mongodb_server_status_repl_is_implicit_default_majority_wc

### repl.isWritablePrimary (1)

- `db.serverStatus().repl.isWritablePrimary` — modern: mongodb_server_status_repl_role

### repl.primaryOnlyServices (8)

- `db.serverStatus().repl.primaryOnlyServices.RenameCollectionParticipantService.numInstances` — modern: mongodb_server_status_repl_primary_only_service_instances
- `db.serverStatus().repl.primaryOnlyServices.RenameCollectionParticipantService.state` — modern: mongodb_server_status_repl_primary_only_service_state
- `db.serverStatus().repl.primaryOnlyServices.ReshardingDonorService.numInstances` — modern: mongodb_server_status_repl_primary_only_service_instances
- `db.serverStatus().repl.primaryOnlyServices.ReshardingDonorService.state` — modern: mongodb_server_status_repl_primary_only_service_state
- `db.serverStatus().repl.primaryOnlyServices.ReshardingRecipientService.numInstances` — modern: mongodb_server_status_repl_primary_only_service_instances
- `db.serverStatus().repl.primaryOnlyServices.ReshardingRecipientService.state` — modern: mongodb_server_status_repl_primary_only_service_state
- `db.serverStatus().repl.primaryOnlyServices.ShardingDDLCoordinator.numInstances` — modern: mongodb_server_status_repl_primary_only_service_instances
- `db.serverStatus().repl.primaryOnlyServices.ShardingDDLCoordinator.state` — modern: mongodb_server_status_repl_primary_only_service_state

### repl.rbid (1)

- `db.serverStatus().repl.rbid` — modern: mongodb_server_status_repl_rbid

### repl.secondary (1)

- `db.serverStatus().repl.secondary` — modern: mongodb_server_status_repl_role

### repl.setVersion (1)

- `db.serverStatus().repl.setVersion` — modern: mongodb_server_status_repl_set_version

### security.authentication (19)

- `db.serverStatus().security.authentication.mechanisms['MONGODB-X509'].authenticate.received` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['MONGODB-X509'].authenticate.successful` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['MONGODB-X509'].clusterAuthenticate.received` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['MONGODB-X509'].clusterAuthenticate.successful` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['MONGODB-X509'].speculativeAuthenticate.received` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['MONGODB-X509'].speculativeAuthenticate.successful` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-1'].authenticate.received` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-1'].authenticate.successful` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-1'].clusterAuthenticate.received` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-1'].clusterAuthenticate.successful` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-1'].speculativeAuthenticate.received` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-1'].speculativeAuthenticate.successful` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-256'].authenticate.received` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-256'].authenticate.successful` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-256'].clusterAuthenticate.received` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-256'].clusterAuthenticate.successful` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-256'].speculativeAuthenticate.received` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.mechanisms['SCRAM-SHA-256'].speculativeAuthenticate.successful` — modern: mongodb_server_status_security_authentication_total
- `db.serverStatus().security.authentication.saslSupportedMechsReceived` — modern: mongodb_server_status_security_authentication_sasl_supported_mechs_received_total

### security.SSLServerCertificateExpirationDate (1)

- `db.serverStatus().security.SSLServerCertificateExpirationDate` — modern: mongodb_server_status_security_sslserver_certificate_expiration_date_seconds

### shardingStatistics.catalogCache (9)

- `db.serverStatus().shardingStatistics.catalogCache.countFailedRefreshes` — modern: mongodb_server_status_sharding_statistics_catalog_cache_count_failed_refreshes_total
- `db.serverStatus().shardingStatistics.catalogCache.countFullRefreshesStarted` — modern: mongodb_server_status_sharding_statistics_catalog_cache_count_full_refreshes_started_total
- `db.serverStatus().shardingStatistics.catalogCache.countIncrementalRefreshesStarted` — modern: mongodb_server_status_sharding_statistics_catalog_cache_count_incremental_refreshes_started_total
- `db.serverStatus().shardingStatistics.catalogCache.countStaleConfigErrors` — modern: mongodb_server_status_sharding_statistics_catalog_cache_count_stale_config_errors_total
- `db.serverStatus().shardingStatistics.catalogCache.numActiveFullRefreshes` — modern: mongodb_server_status_sharding_statistics_catalog_cache_num_active_full_refreshes
- `db.serverStatus().shardingStatistics.catalogCache.numActiveIncrementalRefreshes` — modern: mongodb_server_status_sharding_statistics_catalog_cache_num_active_incremental_refreshes
- `db.serverStatus().shardingStatistics.catalogCache.numCollectionEntries` — modern: mongodb_server_status_sharding_statistics_catalog_cache_num_collection_entries
- `db.serverStatus().shardingStatistics.catalogCache.numDatabaseEntries` — modern: mongodb_server_status_sharding_statistics_catalog_cache_num_database_entries
- `db.serverStatus().shardingStatistics.catalogCache.totalRefreshWaitTimeMicros` — modern: mongodb_server_status_sharding_statistics_catalog_cache_total_refresh_wait_time_seconds

### shardingStatistics.chunkMigrationConcurrency (1)

- `db.serverStatus().shardingStatistics.chunkMigrationConcurrency` — modern: mongodb_server_status_sharding_statistics_chunk_migration_concurrency

### shardingStatistics.countBytesClonedOnCatchUpOnRecipient (1)

- `db.serverStatus().shardingStatistics.countBytesClonedOnCatchUpOnRecipient` — modern: mongodb_server_status_sharding_statistics_count_bytes_cloned_on_catch_up_on_recipient_total

### shardingStatistics.countBytesClonedOnDonor (1)

- `db.serverStatus().shardingStatistics.countBytesClonedOnDonor` — modern: mongodb_server_status_sharding_statistics_count_bytes_cloned_on_donor_total

### shardingStatistics.countBytesClonedOnRecipient (1)

- `db.serverStatus().shardingStatistics.countBytesClonedOnRecipient` — modern: mongodb_server_status_sharding_statistics_count_bytes_cloned_on_recipient_total

### shardingStatistics.countBytesDeletedByRangeDeleter (1)

- `db.serverStatus().shardingStatistics.countBytesDeletedByRangeDeleter` — modern: mongodb_server_status_sharding_statistics_count_bytes_deleted_by_range_deleter_total

### shardingStatistics.countDocsClonedOnCatchUpOnRecipient (1)

- `db.serverStatus().shardingStatistics.countDocsClonedOnCatchUpOnRecipient` — modern: mongodb_server_status_sharding_statistics_count_docs_cloned_on_catch_up_on_recipient_total

### shardingStatistics.countDocsClonedOnDonor (1)

- `db.serverStatus().shardingStatistics.countDocsClonedOnDonor` — modern: mongodb_server_status_sharding_statistics_count_docs_cloned_on_donor_total

### shardingStatistics.countDocsClonedOnRecipient (1)

- `db.serverStatus().shardingStatistics.countDocsClonedOnRecipient` — modern: mongodb_server_status_sharding_statistics_count_docs_cloned_on_recipient_total

### shardingStatistics.countDocsDeletedByRangeDeleter (1)

- `db.serverStatus().shardingStatistics.countDocsDeletedByRangeDeleter` — modern: mongodb_server_status_sharding_statistics_count_docs_deleted_by_range_deleter_total

### shardingStatistics.countDonorMoveChunkAbortConflictingIndexOperation (1)

- `db.serverStatus().shardingStatistics.countDonorMoveChunkAbortConflictingIndexOperation` — modern: mongodb_server_status_sharding_statistics_count_donor_move_chunk_abort_conflicting_index_operation_total

### shardingStatistics.countDonorMoveChunkAborted (1)

- `db.serverStatus().shardingStatistics.countDonorMoveChunkAborted` — modern: mongodb_server_status_sharding_move_chunk_total

### shardingStatistics.countDonorMoveChunkCommitted (1)

- `db.serverStatus().shardingStatistics.countDonorMoveChunkCommitted` — modern: mongodb_server_status_sharding_move_chunk_total

### shardingStatistics.countDonorMoveChunkLockTimeout (1)

- `db.serverStatus().shardingStatistics.countDonorMoveChunkLockTimeout` — modern: mongodb_server_status_sharding_statistics_count_donor_move_chunk_lock_timeout_total

### shardingStatistics.countDonorMoveChunkStarted (1)

- `db.serverStatus().shardingStatistics.countDonorMoveChunkStarted` — modern: mongodb_server_status_sharding_move_chunk_total

### shardingStatistics.countRecipientMoveChunkStarted (1)

- `db.serverStatus().shardingStatistics.countRecipientMoveChunkStarted` — modern: mongodb_server_status_sharding_move_chunk_total

### shardingStatistics.countStaleConfigErrors (1)

- `db.serverStatus().shardingStatistics.countStaleConfigErrors` — modern: mongodb_server_status_sharding_statistics_count_stale_config_errors_total

### shardingStatistics.rangeDeleterTasks (1)

- `db.serverStatus().shardingStatistics.rangeDeleterTasks` — modern: mongodb_server_status_sharding_statistics_range_deleter_tasks

### shardingStatistics.resharding (15)

- `db.serverStatus().shardingStatistics.resharding.bytesCopied` — modern: mongodb_server_status_sharding_statistics_resharding_bytes_copied
- `db.serverStatus().shardingStatistics.resharding.countReshardingCanceled` — modern: mongodb_server_status_sharding_statistics_resharding_count_resharding_canceled_total
- `db.serverStatus().shardingStatistics.resharding.countReshardingFailures` — modern: mongodb_server_status_sharding_statistics_resharding_count_resharding_failures_total
- `db.serverStatus().shardingStatistics.resharding.countReshardingOperations` — modern: mongodb_server_status_sharding_statistics_resharding_count_resharding_operations_total
- `db.serverStatus().shardingStatistics.resharding.countReshardingSuccessful` — modern: mongodb_server_status_sharding_statistics_resharding_count_resharding_successful_total
- `db.serverStatus().shardingStatistics.resharding.countWritesDuringCriticalSection` — modern: mongodb_server_status_sharding_statistics_resharding_count_writes_during_critical_section_total
- `db.serverStatus().shardingStatistics.resharding.documentsCopied` — modern: mongodb_server_status_sharding_statistics_resharding_documents_copied
- `db.serverStatus().shardingStatistics.resharding.lastOpEndingChunkImbalance` — modern: mongodb_server_status_sharding_statistics_resharding_last_op_ending_chunk_imbalance
- `db.serverStatus().shardingStatistics.resharding.maxShardRemainingOperationTimeEstimatedMillis` — modern: mongodb_server_status_sharding_statistics_resharding_max_shard_remaining_operation_time_estimated_seconds
- `db.serverStatus().shardingStatistics.resharding.minShardRemainingOperationTimeEstimatedMillis` — modern: mongodb_server_status_sharding_statistics_resharding_min_shard_remaining_operation_time_estimated_seconds
- `db.serverStatus().shardingStatistics.resharding.opcounters.delete` — modern: mongodb_server_status_sharding_statistics_resharding_opcounters_delete_total
- `db.serverStatus().shardingStatistics.resharding.opcounters.insert` — modern: mongodb_server_status_sharding_statistics_resharding_opcounters_insert_total
- `db.serverStatus().shardingStatistics.resharding.opcounters.update` — modern: mongodb_server_status_sharding_statistics_resharding_opcounters_update_seconds
- `db.serverStatus().shardingStatistics.resharding.oplogEntriesApplied` — modern: mongodb_server_status_sharding_statistics_resharding_oplog_entries_applied
- `db.serverStatus().shardingStatistics.resharding.oplogEntriesFetched` — modern: mongodb_server_status_sharding_statistics_resharding_oplog_entries_fetched

### shardingStatistics.totalCriticalSectionCommitTimeMillis (1)

- `db.serverStatus().shardingStatistics.totalCriticalSectionCommitTimeMillis` — modern: mongodb_server_status_sharding_statistics_total_critical_section_commit_time_seconds

### shardingStatistics.totalCriticalSectionTimeMillis (1)

- `db.serverStatus().shardingStatistics.totalCriticalSectionTimeMillis` — modern: mongodb_server_status_sharding_statistics_total_critical_section_time_seconds

### shardingStatistics.totalDonorChunkCloneTimeMillis (1)

- `db.serverStatus().shardingStatistics.totalDonorChunkCloneTimeMillis` — modern: mongodb_server_status_sharding_statistics_total_donor_chunk_clone_time_seconds

### shardingStatistics.totalDonorMoveChunkTimeMillis (1)

- `db.serverStatus().shardingStatistics.totalDonorMoveChunkTimeMillis` — modern: mongodb_server_status_sharding_statistics_total_donor_move_chunk_time_seconds

### shardingStatistics.unfinishedMigrationFromPreviousPrimary (1)

- `db.serverStatus().shardingStatistics.unfinishedMigrationFromPreviousPrimary` — modern: mongodb_server_status_sharding_statistics_unfinished_migration_from_previous_primary

### storageEngine.backupCursorOpen (1)

- `db.serverStatus().storageEngine.backupCursorOpen` — modern: mongodb_server_status_storage_engine_backup_cursor_open

### storageEngine.dropPendingIdents (1)

- `db.serverStatus().storageEngine.dropPendingIdents` — modern: mongodb_server_status_storage_engine_drop_pending_idents

### storageEngine.persistent (1)

- `db.serverStatus().storageEngine.persistent` — modern: mongodb_server_status_storage_engine_persistent

### storageEngine.readOnly (1)

- `db.serverStatus().storageEngine.readOnly` — modern: mongodb_server_status_storage_engine_read_only

### storageEngine.supportsCommittedReads (1)

- `db.serverStatus().storageEngine.supportsCommittedReads` — modern: mongodb_server_status_storage_engine_supports_committed_reads

### storageEngine.supportsPendingDrops (1)

- `db.serverStatus().storageEngine.supportsPendingDrops` — modern: mongodb_server_status_storage_engine_supports_pending_drops

### storageEngine.supportsResumableIndexBuilds (1)

- `db.serverStatus().storageEngine.supportsResumableIndexBuilds` — modern: mongodb_server_status_storage_engine_supports_resumable_index_builds

### storageEngine.supportsSnapshotReadConcern (1)

- `db.serverStatus().storageEngine.supportsSnapshotReadConcern` — modern: mongodb_server_status_storage_engine_supports_snapshot_read_concern

### transactions.currentActive (1)

- `db.serverStatus().transactions.currentActive` — modern: mongodb_server_status_transactions_current_active

### transactions.currentInactive (1)

- `db.serverStatus().transactions.currentInactive` — modern: mongodb_server_status_transactions_current_inactive

### transactions.currentOpen (1)

- `db.serverStatus().transactions.currentOpen` — modern: mongodb_server_status_transactions_current_open

### transactions.currentPrepared (1)

- `db.serverStatus().transactions.currentPrepared` — modern: mongodb_server_status_transactions_current_prepared

### transactions.retriedCommandsCount (1)

- `db.serverStatus().transactions.retriedCommandsCount` — modern: mongodb_server_status_transactions_retried_commands_count_total

### transactions.retriedStatementsCount (1)

- `db.serverStatus().transactions.retriedStatementsCount` — modern: mongodb_server_status_transactions_retried_statements_count_total

### transactions.totalAborted (1)

- `db.serverStatus().transactions.totalAborted` — modern: mongodb_server_status_transactions_total_aborted_total

### transactions.totalCommitted (1)

- `db.serverStatus().transactions.totalCommitted` — modern: mongodb_server_status_transactions_total_committed_total

### transactions.totalPrepared (1)

- `db.serverStatus().transactions.totalPrepared` — modern: mongodb_server_status_transactions_total_prepared_total

### transactions.totalPreparedThenAborted (1)

- `db.serverStatus().transactions.totalPreparedThenAborted` — modern: mongodb_server_status_transactions_total_prepared_then_aborted_total

### transactions.totalPreparedThenCommitted (1)

- `db.serverStatus().transactions.totalPreparedThenCommitted` — modern: mongodb_server_status_transactions_total_prepared_then_committed_total

### transactions.totalStarted (1)

- `db.serverStatus().transactions.totalStarted` — modern: mongodb_server_status_transactions_total_started_total

### transactions.transactionsCollectionWriteCount (1)

- `db.serverStatus().transactions.transactionsCollectionWriteCount` — modern: mongodb_server_status_transactions_transactions_collection_write_count_total

### wiredTiger.snapshot-window-settings (4)

- `db.serverStatus().wiredTiger['snapshot-window-settings']['current available snapshot window size in seconds']` — modern: mongodb_server_status_wired_tiger_snapshot_window_settings_current_available_snapshot_window_size_in_seconds
- `db.serverStatus().wiredTiger['snapshot-window-settings']['minimum target snapshot window size in seconds']` — modern: mongodb_server_status_wired_tiger_snapshot_window_settings_minimum_target_snapshot_window_size_in_seconds
- `db.serverStatus().wiredTiger['snapshot-window-settings']['pinned timestamp requests']` — modern: mongodb_server_status_wired_tiger_snapshot_window_settings_pinned_timestamp_requests
- `db.serverStatus().wiredTiger['snapshot-window-settings']['total number of SnapshotTooOld errors']` — modern: mongodb_server_status_wired_tiger_snapshot_window_settings_total_number_of_snapshot_too_old_errors_total
