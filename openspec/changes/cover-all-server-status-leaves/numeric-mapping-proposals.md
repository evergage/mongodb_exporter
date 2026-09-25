# Numeric and boolean mapping proposals

This file covers all **1158 currently unmapped numeric and boolean leaves** from the primary/secondary fixture union.

## Applied policy

- Prefer counters for cumulative monotonic event, byte, operation, and elapsed-time statistics.
- Use gauges for current snapshots, populations, capacities, limits, rates, percentages, state values, booleans, timestamps, and generation identifiers.
- Preserve raw MongoDB units. Every generated mapping uses conversion factor `1`.
- Dynamic-key families use bounded grouped labels; histogram ranges use fixed numeric-boundary counter labels; final series ceilings are derived after generation.
- Existing covered mappings are not changed by this proposal.

## Summary

| Decision | Leaves |
|---|---:|
| Generated counter proposals | 411 |
| Generated gauge proposals | 248 |
| Dynamic-key manual review | 346 |
| Histogram manual review | 151 |

### Research basis

- MongoDB `serverStatus` reference: https://www.mongodb.com/docs/v7.0/reference/command/serverstatus/
- MongoDB 5.0 WiredTiger statistics definitions: https://github.com/mongodb/mongo/blob/v5.0/src/third_party/wiredtiger/dist/stat_data.py
- MongoDB 5.0 snapshot formatter: https://github.com/mongodb/mongo/blob/v5.0/src/mongo/bson/timestamp.cpp
- WiredTiger checkpoint architecture: https://source.wiredtiger.com/develop/arch-checkpoint.html
- WiredTiger transaction architecture: https://source.wiredtiger.com/develop/arch-transaction.html

The MongoDB 5.0 WiredTiger source states that statistics are operation-rate values by default; snapshot values are expected to use wording such as `currently` or `in the cache`. The generated decisions use that source rule plus explicit exceptions for state, size, generation, and configuration values.

### Approved dynamic and histogram policy

- Commands, stages, and operator counters use grouped families with fixture-derived bounded label vocabularies.
- Unknown runtime keys emit no sample and increment a diagnostic counter labeled only by dynamic family.
- Histogram bucket counts use grouped counters with one numeric boundary label in the source base unit; no synthetic Prometheus histogram sum is created.
- Query multi-planner histogram `lowerBound` leaves supply the numeric boundary label for their paired count leaf rather than a separate sample.
- Final series ceilings are derived from the approved mapping and reviewed after primary/secondary canaries.

### Researched exceptions

- `db.serverStatus().connections.exhaustHello` is a gauge: current connections whose last request is exhaust-enabled `hello`.
- `db.serverStatus().transportSecurity.*` are counters: cumulative TLS connections by protocol.
- `db.serverStatus().wiredTiger.data-handle['connection data handle size']` is a gauge: current memory size.
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint generation']` is a gauge: an internal generation value.
- `db.serverStatus().wiredTiger.transaction['prepared transactions']` is a counter: cumulative prepare operations.
- `db.serverStatus().wiredTiger.session['tiered storage local retention time (secs)']` is a gauge: configured retention interval.
- `db.serverStatus().wiredTiger.session['table compact timeout']` is a counter: cumulative compactions that reached timeout.

## Generated counter proposals

### asserts.tripwire (1)

- `db.serverStatus().asserts.tripwire` — counter; family: `mongodb_server_status_asserts_tripwire_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic

### connections.awaitingTopologyChanges (1)

- `db.serverStatus().connections.awaitingTopologyChanges` — counter; family: `mongodb_server_status_connections_awaiting_topology_changes_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### connections.threaded (1)

- `db.serverStatus().connections.threaded` — counter; family: `mongodb_server_status_connections_threaded_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### extra_info.input_blocks (1)

- `db.serverStatus().extra_info.input_blocks` — counter; family: `mongodb_server_status_extra_info_input_blocks_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic

### extra_info.involuntary_context_switches (1)

- `db.serverStatus().extra_info.involuntary_context_switches` — counter; family: `mongodb_server_status_extra_info_involuntary_context_switches_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### extra_info.output_blocks (1)

- `db.serverStatus().extra_info.output_blocks` — counter; family: `mongodb_server_status_extra_info_output_blocks_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic

### extra_info.page_reclaims (1)

- `db.serverStatus().extra_info.page_reclaims` — counter; family: `mongodb_server_status_extra_info_page_reclaims_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### extra_info.system_time_us (1)

- `db.serverStatus().extra_info.system_time_us` — counter; family: `mongodb_server_status_extra_info_system_time_us_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic

### extra_info.threads (1)

- `db.serverStatus().extra_info.threads` — counter; family: `mongodb_server_status_extra_info_threads_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### extra_info.user_time_us (1)

- `db.serverStatus().extra_info.user_time_us` — counter; family: `mongodb_server_status_extra_info_user_time_us_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic

### extra_info.voluntary_context_switches (1)

- `db.serverStatus().extra_info.voluntary_context_switches` — counter; family: `mongodb_server_status_extra_info_voluntary_context_switches_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### globalLock.totalTime (1)

- `db.serverStatus().globalLock.totalTime` — counter; family: `mongodb_server_status_global_lock_total_time_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### indexBulkBuilder.count (1)

- `db.serverStatus().indexBulkBuilder.count` — counter; family: `mongodb_server_status_index_bulk_builder_count_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### indexBulkBuilder.filesClosedForExternalSort (1)

- `db.serverStatus().indexBulkBuilder.filesClosedForExternalSort` — counter; family: `mongodb_server_status_index_bulk_builder_files_closed_for_external_sort_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### indexBulkBuilder.filesOpenedForExternalSort (1)

- `db.serverStatus().indexBulkBuilder.filesOpenedForExternalSort` — counter; family: `mongodb_server_status_index_bulk_builder_files_opened_for_external_sort_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### indexBulkBuilder.resumed (1)

- `db.serverStatus().indexBulkBuilder.resumed` — counter; family: `mongodb_server_status_index_bulk_builder_resumed_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic

### indexStats.features (14)

- `db.serverStatus().indexStats.features['2d'].accesses` — counter; family: `mongodb_server_status_index_stats_features_2d_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features['2dsphere'].accesses` — counter; family: `mongodb_server_status_index_stats_features_2dsphere_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.collation.accesses` — counter; family: `mongodb_server_status_index_stats_features_collation_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.compound.accesses` — counter; family: `mongodb_server_status_index_stats_features_compound_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.hashed.accesses` — counter; family: `mongodb_server_status_index_stats_features_hashed_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.id.accesses` — counter; family: `mongodb_server_status_index_stats_features_id_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.normal.accesses` — counter; family: `mongodb_server_status_index_stats_features_normal_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.partial.accesses` — counter; family: `mongodb_server_status_index_stats_features_partial_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.single.accesses` — counter; family: `mongodb_server_status_index_stats_features_single_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.sparse.accesses` — counter; family: `mongodb_server_status_index_stats_features_sparse_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.text.accesses` — counter; family: `mongodb_server_status_index_stats_features_text_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.ttl.accesses` — counter; family: `mongodb_server_status_index_stats_features_ttl_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.unique.accesses` — counter; family: `mongodb_server_status_index_stats_features_unique_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses
- `db.serverStatus().indexStats.features.wildcard.accesses` — counter; family: `mongodb_server_status_index_stats_features_wildcard_accesses_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative index accesses

### lastCommittedOpTime[0] (1)

- `db.serverStatus().lastCommittedOpTime[0]` — counter; family: `mongodb_server_status_last_committed_op_time_0_total`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### lastCommittedOpTime[1] (1)

- `db.serverStatus().lastCommittedOpTime[1]` — counter; family: `mongodb_server_status_last_committed_op_time_1_total`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### locks.Collection (9)

- `db.serverStatus().locks.Collection.acquireCount.R` — counter; family: `mongodb_server_status_locks_collection_acquire_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Collection.acquireCount.W` — counter; family: `mongodb_server_status_locks_collection_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Collection.acquireCount.r` — counter; family: `mongodb_server_status_locks_collection_acquire_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Collection.acquireCount.w` — counter; family: `mongodb_server_status_locks_collection_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Collection.acquireWaitCount.R` — counter; family: `mongodb_server_status_locks_collection_acquire_wait_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Collection.acquireWaitCount.W` — counter; family: `mongodb_server_status_locks_collection_acquire_wait_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Collection.acquireWaitCount.w` — counter; family: `mongodb_server_status_locks_collection_acquire_wait_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Collection.timeAcquiringMicros.R` — counter; family: `mongodb_server_status_locks_collection_time_acquiring_micros_r_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Collection.timeAcquiringMicros.W` — counter; family: `mongodb_server_status_locks_collection_time_acquiring_micros_w_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### locks.Database (6)

- `db.serverStatus().locks.Database.acquireCount.W` — counter; family: `mongodb_server_status_locks_database_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Database.acquireCount.r` — counter; family: `mongodb_server_status_locks_database_acquire_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Database.acquireCount.w` — counter; family: `mongodb_server_status_locks_database_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Database.acquireWaitCount.W` — counter; family: `mongodb_server_status_locks_database_acquire_wait_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Database.acquireWaitCount.w` — counter; family: `mongodb_server_status_locks_database_acquire_wait_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Database.timeAcquiringMicros.W` — counter; family: `mongodb_server_status_locks_database_time_acquiring_micros_w_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### locks.FeatureCompatibilityVersion (2)

- `db.serverStatus().locks.FeatureCompatibilityVersion.acquireCount.r` — counter; family: `mongodb_server_status_locks_feature_compatibility_version_acquire_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.FeatureCompatibilityVersion.acquireCount.w` — counter; family: `mongodb_server_status_locks_feature_compatibility_version_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### locks.Global (6)

- `db.serverStatus().locks.Global.acquireCount.W` — counter; family: `mongodb_server_status_locks_global_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Global.acquireCount.r` — counter; family: `mongodb_server_status_locks_global_acquire_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Global.acquireCount.w` — counter; family: `mongodb_server_status_locks_global_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Global.acquireWaitCount.W` — counter; family: `mongodb_server_status_locks_global_acquire_wait_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Global.acquireWaitCount.r` — counter; family: `mongodb_server_status_locks_global_acquire_wait_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Global.timeAcquiringMicros.W` — counter; family: `mongodb_server_status_locks_global_time_acquiring_micros_w_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### locks.Mutex (5)

- `db.serverStatus().locks.Mutex.acquireCount.W` — counter; family: `mongodb_server_status_locks_mutex_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Mutex.acquireCount.r` — counter; family: `mongodb_server_status_locks_mutex_acquire_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Mutex.acquireWaitCount.W` — counter; family: `mongodb_server_status_locks_mutex_acquire_wait_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Mutex.acquireWaitCount.r` — counter; family: `mongodb_server_status_locks_mutex_acquire_wait_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.Mutex.timeAcquiringMicros.W` — counter; family: `mongodb_server_status_locks_mutex_time_acquiring_micros_w_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### locks.oplog (3)

- `db.serverStatus().locks.oplog.acquireCount.W` — counter; family: `mongodb_server_status_locks_oplog_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.oplog.acquireCount.r` — counter; family: `mongodb_server_status_locks_oplog_acquire_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.oplog.acquireCount.w` — counter; family: `mongodb_server_status_locks_oplog_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### locks.ParallelBatchWriterMode (5)

- `db.serverStatus().locks.ParallelBatchWriterMode.acquireCount.W` — counter; family: `mongodb_server_status_locks_parallel_batch_writer_mode_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.ParallelBatchWriterMode.acquireCount.r` — counter; family: `mongodb_server_status_locks_parallel_batch_writer_mode_acquire_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.ParallelBatchWriterMode.acquireWaitCount.W` — counter; family: `mongodb_server_status_locks_parallel_batch_writer_mode_acquire_wait_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.ParallelBatchWriterMode.acquireWaitCount.r` — counter; family: `mongodb_server_status_locks_parallel_batch_writer_mode_acquire_wait_count_r_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.ParallelBatchWriterMode.timeAcquiringMicros.W` — counter; family: `mongodb_server_status_locks_parallel_batch_writer_mode_time_acquiring_micros_w_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### locks.ReplicationStateTransition (3)

- `db.serverStatus().locks.ReplicationStateTransition.acquireCount.W` — counter; family: `mongodb_server_status_locks_replication_state_transition_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.ReplicationStateTransition.acquireCount.w` — counter; family: `mongodb_server_status_locks_replication_state_transition_acquire_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().locks.ReplicationStateTransition.acquireWaitCount.w` — counter; family: `mongodb_server_status_locks_replication_state_transition_acquire_wait_count_w_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### metrics.changeStreams (1)

- `db.serverStatus().metrics.changeStreams.largeEventsFailed` — counter; family: `mongodb_server_status_metrics_change_streams_large_events_failed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### metrics.dotsAndDollarsFields (2)

- `db.serverStatus().metrics.dotsAndDollarsFields.inserts` — counter; family: `mongodb_server_status_metrics_dots_and_dollars_fields_inserts_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic
- `db.serverStatus().metrics.dotsAndDollarsFields.updates` — counter; family: `mongodb_server_status_metrics_dots_and_dollars_fields_updates_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic

### metrics.getLastError (1)

- `db.serverStatus().metrics.getLastError.default.wtimeouts` — counter; family: `mongodb_server_status_metrics_get_last_error_default_wtimeouts_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### metrics.mongos (2)

- `db.serverStatus().metrics.mongos.cursor.moreThanOneBatch` — counter; family: `mongodb_server_status_metrics_mongos_cursor_more_than_one_batch_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic
- `db.serverStatus().metrics.mongos.cursor.totalOpened` — counter; family: `mongodb_server_status_metrics_mongos_cursor_total_opened_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### metrics.operation (3)

- `db.serverStatus().metrics.operation.transactionTooLargeForCacheErrors` — counter; family: `mongodb_server_status_metrics_operation_transaction_too_large_for_cache_errors_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.operation.transactionTooLargeForCacheErrorsConvertedToWriteConflict` — counter; family: `mongodb_server_status_metrics_operation_transaction_too_large_for_cache_errors_converted_to_write_conflict_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.operation.writeConflicts` — counter; family: `mongodb_server_status_metrics_operation_write_conflicts_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### metrics.query (6)

- `db.serverStatus().metrics.query.multiPlanner.classicCount` — counter; family: `mongodb_server_status_metrics_query_multi_planner_classic_count_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative multi-planner count, work, read, or elapsed-microsecond statistic
- `db.serverStatus().metrics.query.multiPlanner.classicMicros` — counter; family: `mongodb_server_status_metrics_query_multi_planner_classic_micros_total`; raw unit: microseconds; conversion: 1; confidence: high; evidence: cumulative multi-planner count, work, read, or elapsed-microsecond statistic
- `db.serverStatus().metrics.query.multiPlanner.classicWorks` — counter; family: `mongodb_server_status_metrics_query_multi_planner_classic_works_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative multi-planner count, work, read, or elapsed-microsecond statistic
- `db.serverStatus().metrics.query.multiPlanner.sbeCount` — counter; family: `mongodb_server_status_metrics_query_multi_planner_sbe_count_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative multi-planner count, work, read, or elapsed-microsecond statistic
- `db.serverStatus().metrics.query.multiPlanner.sbeMicros` — counter; family: `mongodb_server_status_metrics_query_multi_planner_sbe_micros_total`; raw unit: microseconds; conversion: 1; confidence: high; evidence: cumulative multi-planner count, work, read, or elapsed-microsecond statistic
- `db.serverStatus().metrics.query.multiPlanner.sbeNumReads` — counter; family: `mongodb_server_status_metrics_query_multi_planner_sbe_num_reads_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative multi-planner count, work, read, or elapsed-microsecond statistic

### metrics.queryExecutor (2)

- `db.serverStatus().metrics.queryExecutor.collectionScans.nonTailable` — counter; family: `mongodb_server_status_metrics_query_executor_collection_scans_non_tailable_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.queryExecutor.collectionScans.total` — counter; family: `mongodb_server_status_metrics_query_executor_collection_scans_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### metrics.repl (17)

- `db.serverStatus().metrics.repl.apply.attemptsToBecomeSecondary` — counter; family: `mongodb_server_status_metrics_repl_apply_attempts_to_become_secondary_total`; raw unit: seconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.initialSync.completed` — counter; family: `mongodb_server_status_metrics_repl_initial_sync_completed_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic
- `db.serverStatus().metrics.repl.initialSync.failedAttempts` — counter; family: `mongodb_server_status_metrics_repl_initial_sync_failed_attempts_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.initialSync.failures` — counter; family: `mongodb_server_status_metrics_repl_initial_sync_failures_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.network.getmores.numEmptyBatches` — counter; family: `mongodb_server_status_metrics_repl_network_getmores_num_empty_batches_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.network.notPrimaryLegacyUnacknowledgedWrites` — counter; family: `mongodb_server_status_metrics_repl_network_not_primary_legacy_unacknowledged_writes_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.network.notPrimaryUnacknowledgedWrites` — counter; family: `mongodb_server_status_metrics_repl_network_not_primary_unacknowledged_writes_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.network.oplogGetMoresProcessed.num` — counter; family: `mongodb_server_status_metrics_repl_network_oplog_get_mores_processed_num_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.network.oplogGetMoresProcessed.totalMillis` — counter; family: `mongodb_server_status_metrics_repl_network_oplog_get_mores_processed_total_millis_total`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.network.replSetUpdatePosition.num` — counter; family: `mongodb_server_status_metrics_repl_network_repl_set_update_position_num_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic
- `db.serverStatus().metrics.repl.reconfig.numAutoReconfigsForRemovalOfNewlyAddedFields` — counter; family: `mongodb_server_status_metrics_repl_reconfig_num_auto_reconfigs_for_removal_of_newly_added_fields_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.stateTransition.userOperationsKilled` — counter; family: `mongodb_server_status_metrics_repl_state_transition_user_operations_killed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.syncSource.numSelections` — counter; family: `mongodb_server_status_metrics_repl_sync_source_num_selections_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.syncSource.numSyncSourceChangesDueToSignificantlyCloserNode` — counter; family: `mongodb_server_status_metrics_repl_sync_source_num_sync_source_changes_due_to_significantly_closer_node_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.syncSource.numTimesChoseDifferent` — counter; family: `mongodb_server_status_metrics_repl_sync_source_num_times_chose_different_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.syncSource.numTimesChoseSame` — counter; family: `mongodb_server_status_metrics_repl_sync_source_num_times_chose_same_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().metrics.repl.syncSource.numTimesCouldNotFind` — counter; family: `mongodb_server_status_metrics_repl_sync_source_num_times_could_not_find_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### operationTime[0] (1)

- `db.serverStatus().operationTime[0]` — counter; family: `mongodb_server_status_operation_time_0_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### operationTime[1] (1)

- `db.serverStatus().operationTime[1]` — counter; family: `mongodb_server_status_operation_time_1_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### opLatencies.transactions (2)

- `db.serverStatus().opLatencies.transactions.latency` — counter; family: `mongodb_server_status_op_latencies_transactions_latency_total`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: cumulative event or elapsed-work statistic
- `db.serverStatus().opLatencies.transactions.ops` — counter; family: `mongodb_server_status_op_latencies_transactions_ops_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### repl.lastWrite (6)

- `db.serverStatus().repl.lastWrite.majorityOpTime.t` — counter; family: `mongodb_server_status_repl_last_write_majority_op_time_t_total`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().repl.lastWrite.majorityOpTime.ts[0]` — counter; family: `mongodb_server_status_repl_last_write_majority_op_time_ts_0_total`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().repl.lastWrite.majorityOpTime.ts[1]` — counter; family: `mongodb_server_status_repl_last_write_majority_op_time_ts_1_total`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().repl.lastWrite.opTime.t` — counter; family: `mongodb_server_status_repl_last_write_op_time_t_total`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().repl.lastWrite.opTime.ts[0]` — counter; family: `mongodb_server_status_repl_last_write_op_time_ts_0_total`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().repl.lastWrite.opTime.ts[1]` — counter; family: `mongodb_server_status_repl_last_write_op_time_ts_1_total`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### repl.topologyVersion (1)

- `db.serverStatus().repl.topologyVersion.counter` — counter; family: `mongodb_server_status_repl_topology_version_counter_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### scramCache.SCRAM-SHA-1 (2)

- `db.serverStatus().scramCache['SCRAM-SHA-1'].count` — counter; family: `mongodb_server_status_scram_cache_scram_sha_1_count_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().scramCache['SCRAM-SHA-1'].hits` — counter; family: `mongodb_server_status_scram_cache_scram_sha_1_hits_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### scramCache.SCRAM-SHA-256 (2)

- `db.serverStatus().scramCache['SCRAM-SHA-256'].count` — counter; family: `mongodb_server_status_scram_cache_scram_sha_256_count_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().scramCache['SCRAM-SHA-256'].hits` — counter; family: `mongodb_server_status_scram_cache_scram_sha_256_hits_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### tcmalloc.tcmalloc (1)

- `db.serverStatus().tcmalloc.tcmalloc.spinlock_total_delay_ns` — counter; family: `mongodb_server_status_tcmalloc_tcmalloc_spinlock_total_delay_ns_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### tenantMigrations.totalFailedMigrationsDonated (1)

- `db.serverStatus().tenantMigrations.totalFailedMigrationsDonated` — counter; family: `mongodb_server_status_tenant_migrations_total_failed_migrations_donated_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### tenantMigrations.totalFailedMigrationsReceived (1)

- `db.serverStatus().tenantMigrations.totalFailedMigrationsReceived` — counter; family: `mongodb_server_status_tenant_migrations_total_failed_migrations_received_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### tenantMigrations.totalSuccessfulMigrationsDonated (1)

- `db.serverStatus().tenantMigrations.totalSuccessfulMigrationsDonated` — counter; family: `mongodb_server_status_tenant_migrations_total_successful_migrations_donated_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### tenantMigrations.totalSuccessfulMigrationsReceived (1)

- `db.serverStatus().tenantMigrations.totalSuccessfulMigrationsReceived` — counter; family: `mongodb_server_status_tenant_migrations_total_successful_migrations_received_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### transportSecurity.1 (4)

- `db.serverStatus().transportSecurity['1.0']` — counter; family: `mongodb_server_status_transport_security_1_0_total`; raw unit: count or raw value; conversion: 1; confidence: researched; evidence: cumulative TLS connections by protocol
- `db.serverStatus().transportSecurity['1.1']` — counter; family: `mongodb_server_status_transport_security_1_1_total`; raw unit: count or raw value; conversion: 1; confidence: researched; evidence: cumulative TLS connections by protocol
- `db.serverStatus().transportSecurity['1.2']` — counter; family: `mongodb_server_status_transport_security_1_2_total`; raw unit: count or raw value; conversion: 1; confidence: researched; evidence: cumulative TLS connections by protocol
- `db.serverStatus().transportSecurity['1.3']` — counter; family: `mongodb_server_status_transport_security_1_3_total`; raw unit: count or raw value; conversion: 1; confidence: researched; evidence: cumulative TLS connections by protocol

### transportSecurity.unknown (1)

- `db.serverStatus().transportSecurity.unknown` — counter; family: `mongodb_server_status_transport_security_unknown_total`; raw unit: count or raw value; conversion: 1; confidence: researched; evidence: cumulative TLS connections by protocol

### twoPhaseCommitCoordinator.totalAbortedTwoPhaseCommit (1)

- `db.serverStatus().twoPhaseCommitCoordinator.totalAbortedTwoPhaseCommit` — counter; family: `mongodb_server_status_two_phase_commit_coordinator_total_aborted_two_phase_commit_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### twoPhaseCommitCoordinator.totalCommittedTwoPhaseCommit (1)

- `db.serverStatus().twoPhaseCommitCoordinator.totalCommittedTwoPhaseCommit` — counter; family: `mongodb_server_status_two_phase_commit_coordinator_total_committed_two_phase_commit_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### twoPhaseCommitCoordinator.totalCreated (1)

- `db.serverStatus().twoPhaseCommitCoordinator.totalCreated` — counter; family: `mongodb_server_status_two_phase_commit_coordinator_total_created_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### twoPhaseCommitCoordinator.totalStartedTwoPhaseCommit (1)

- `db.serverStatus().twoPhaseCommitCoordinator.totalStartedTwoPhaseCommit` — counter; family: `mongodb_server_status_two_phase_commit_coordinator_total_started_two_phase_commit_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.block-manager (26)

- `db.serverStatus().wiredTiger['block-manager']['block cache cached blocks updated']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_cached_blocks_updated_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger['block-manager']['block cache cached bytes updated']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_cached_bytes_updated_total`; raw unit: bytes; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger['block-manager']['block cache evicted blocks']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_evicted_blocks_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache file size causing bypass']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_file_size_causing_bypass_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger['block-manager']['block cache lookups']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_lookups_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger['block-manager']['block cache number of blocks not evicted due to overhead']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_number_of_blocks_not_evicted_due_to_overhead_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache number of bypasses because no-write-allocate setting was on']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_number_of_bypasses_because_no_write_allocate_setting_was_on_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache number of bypasses due to overhead on put']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_number_of_bypasses_due_to_overhead_on_put_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache number of bypasses on get']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_number_of_bypasses_on_get_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache number of bypasses on put because file is too small']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_number_of_bypasses_on_put_because_file_is_too_small_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache number of eviction passes']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_number_of_eviction_passes_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache number of put bypasses on checkpoint I/O']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_number_of_put_bypasses_on_checkpoint_i_o_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache removed blocks']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_removed_blocks_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache total blocks']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_total_blocks_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache total blocks inserted on read path']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_total_blocks_inserted_on_read_path_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache total blocks inserted on write path']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_total_blocks_inserted_on_write_path_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache total bytes']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_total_bytes_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache total bytes inserted on read path']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_total_bytes_inserted_on_read_path_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['block cache total bytes inserted on write path']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_total_bytes_inserted_on_write_path_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['bytes read via memory map API']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_bytes_read_via_memory_map_api_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['bytes read via system call API']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_bytes_read_via_system_call_api_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['bytes written for checkpoint']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_bytes_written_for_checkpoint_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['bytes written via memory map API']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_bytes_written_via_memory_map_api_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['bytes written via system call API']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_bytes_written_via_system_call_api_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['number of times the file was remapped because it changed size via fallocate or truncate']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_number_of_times_the_file_was_remapped_because_it_changed_size_via_fallocate_or_truncate_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['block-manager']['number of times the region was remapped via write']` — counter; family: `mongodb_server_status_wired_tiger_block_manager_number_of_times_the_region_was_remapped_via_write_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.cache (73)

- `db.serverStatus().wiredTiger.cache['bytes allocated for updates']` — counter; family: `mongodb_server_status_wired_tiger_cache_bytes_allocated_for_updates_total`; raw unit: bytes; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.cache['checkpoint blocked page eviction']` — counter; family: `mongodb_server_status_wired_tiger_cache_checkpoint_blocked_page_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction calls to get a page']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_calls_to_get_a_page_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction calls to get a page found queue empty']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_calls_to_get_a_page_found_queue_empty_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction calls to get a page found queue empty after locking']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_calls_to_get_a_page_found_queue_empty_after_locking_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction passes of a file']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_passes_of_a_file_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction server candidate queue empty when topping up']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_server_candidate_queue_empty_when_topping_up_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction server candidate queue not empty when topping up']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_server_candidate_queue_not_empty_when_topping_up_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction server evicting pages']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_server_evicting_pages_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction server skips pages that previously failed eviction and likely will again']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_server_skips_pages_that_previously_failed_eviction_and_likely_will_again_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction server skips pages that we do not want to evict']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_server_skips_pages_that_we_do_not_want_to_evict_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction server skips trees that are being checkpointed']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_server_skips_trees_that_are_being_checkpointed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction server skips trees that were not useful before']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_server_skips_trees_that_were_not_useful_before_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction server slept, because we did not make progress with eviction']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_server_slept_because_we_did_not_make_progress_with_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction server unable to reach eviction goal']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_server_unable_to_reach_eviction_goal_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction server waiting for a leaf page']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_server_waiting_for_a_leaf_page_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction walk most recent sleeps for checkpoint handle gathering']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_walk_most_recent_sleeps_for_checkpoint_handle_gathering_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction walk target strategy only clean pages']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_walk_target_strategy_only_clean_pages_total`; raw unit: raw rate; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction walks abandoned']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_walks_abandoned_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction walks gave up because they restarted their walk twice']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_walks_gave_up_because_they_restarted_their_walk_twice_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction walks gave up because they saw too many pages and found no candidates']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_walks_gave_up_because_they_saw_too_many_pages_and_found_no_candidates_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction walks gave up because they saw too many pages and found too few candidates']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_walks_gave_up_because_they_saw_too_many_pages_and_found_too_few_candidates_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction walks reached end of tree']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_walks_reached_end_of_tree_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction walks restarted']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_walks_restarted_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction walks started from root of tree']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_walks_started_from_root_of_tree_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction walks started from saved location in tree']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_walks_started_from_saved_location_in_tree_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction worker thread created']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_worker_thread_created_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction worker thread evicting pages']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_worker_thread_evicting_pages_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction worker thread removed']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_worker_thread_removed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['eviction worker thread stable number']` — counter; family: `mongodb_server_status_wired_tiger_cache_eviction_worker_thread_stable_number_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['files with new eviction walks started']` — counter; family: `mongodb_server_status_wired_tiger_cache_files_with_new_eviction_walks_started_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['force re-tuning of eviction workers once in a while']` — counter; family: `mongodb_server_status_wired_tiger_cache_force_re_tuning_of_eviction_workers_once_in_a_while_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['forced eviction - do not retry count to evict pages selected to evict during reconciliation']` — counter; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_do_not_retry_count_to_evict_pages_selected_to_evict_during_reconciliation_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['forced eviction - pages evicted that were clean count']` — counter; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_pages_evicted_that_were_clean_count_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['forced eviction - pages evicted that were clean time (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_pages_evicted_that_were_clean_time_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['forced eviction - pages evicted that were dirty count']` — counter; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_pages_evicted_that_were_dirty_count_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['forced eviction - pages evicted that were dirty time (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_pages_evicted_that_were_dirty_time_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['forced eviction - pages selected because of a large number of updates to a single item']` — counter; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_pages_selected_because_of_a_large_number_of_updates_to_a_single_item_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['forced eviction - pages selected because of too many deleted items count']` — counter; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_pages_selected_because_of_too_many_deleted_items_count_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['forced eviction - pages selected count']` — counter; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_pages_selected_count_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['forced eviction - pages selected unable to be evicted count']` — counter; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_pages_selected_unable_to_be_evicted_count_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['forced eviction - pages selected unable to be evicted time']` — counter; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_pages_selected_unable_to_be_evicted_time_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['hazard pointer blocked page eviction']` — counter; family: `mongodb_server_status_wired_tiger_cache_hazard_pointer_blocked_page_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['hazard pointer check calls']` — counter; family: `mongodb_server_status_wired_tiger_cache_hazard_pointer_check_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['hazard pointer check entries walked']` — counter; family: `mongodb_server_status_wired_tiger_cache_hazard_pointer_check_entries_walked_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['in-memory page passed criteria to be split']` — counter; family: `mongodb_server_status_wired_tiger_cache_in_memory_page_passed_criteria_to_be_split_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.cache['in-memory page splits']` — counter; family: `mongodb_server_status_wired_tiger_cache_in_memory_page_splits_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.cache['internal pages evicted']` — counter; family: `mongodb_server_status_wired_tiger_cache_internal_pages_evicted_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['internal pages queued for eviction']` — counter; family: `mongodb_server_status_wired_tiger_cache_internal_pages_queued_for_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['internal pages seen by eviction walk']` — counter; family: `mongodb_server_status_wired_tiger_cache_internal_pages_seen_by_eviction_walk_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['internal pages seen by eviction walk that are already queued']` — counter; family: `mongodb_server_status_wired_tiger_cache_internal_pages_seen_by_eviction_walk_that_are_already_queued_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['internal pages split during eviction']` — counter; family: `mongodb_server_status_wired_tiger_cache_internal_pages_split_during_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['leaf pages split during eviction']` — counter; family: `mongodb_server_status_wired_tiger_cache_leaf_pages_split_during_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['modified pages evicted by application threads']` — counter; family: `mongodb_server_status_wired_tiger_cache_modified_pages_evicted_by_application_threads_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['operations timed out waiting for space in cache']` — counter; family: `mongodb_server_status_wired_tiger_cache_operations_timed_out_waiting_for_space_in_cache_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['overflow pages read into cache']` — counter; family: `mongodb_server_status_wired_tiger_cache_overflow_pages_read_into_cache_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['page split during eviction deepened the tree']` — counter; family: `mongodb_server_status_wired_tiger_cache_page_split_during_eviction_deepened_the_tree_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages dirtied due to obsolete time window']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_dirtied_due_to_obsolete_time_window_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.cache['pages evicted by application threads']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_evicted_by_application_threads_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages evicted in parallel with checkpoint']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_evicted_in_parallel_with_checkpoint_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages queued for eviction']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_queued_for_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages queued for eviction post lru sorting']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_queued_for_eviction_post_lru_sorting_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages queued for urgent eviction']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_queued_for_urgent_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages queued for urgent eviction during walk']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_queued_for_urgent_eviction_during_walk_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages read into cache after truncate']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_read_into_cache_after_truncate_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages removed from the ordinary queue to be queued for urgent eviction']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_removed_from_the_ordinary_queue_to_be_queued_for_urgent_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages requested from the cache']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_requested_from_the_cache_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages seen by eviction walk']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_seen_by_eviction_walk_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages seen by eviction walk that are already queued']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_seen_by_eviction_walk_that_are_already_queued_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages selected for eviction unable to be evicted']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_selected_for_eviction_unable_to_be_evicted_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages selected for eviction unable to be evicted because of failure in reconciliation']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_selected_for_eviction_unable_to_be_evicted_because_of_failure_in_reconciliation_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages walked for eviction']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_walked_for_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cache['pages written requiring in-memory restoration']` — counter; family: `mongodb_server_status_wired_tiger_cache_pages_written_requiring_in_memory_restoration_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.checkpoint-cleanup (3)

- `db.serverStatus().wiredTiger['checkpoint-cleanup']['pages added for eviction']` — counter; family: `mongodb_server_status_wired_tiger_checkpoint_cleanup_pages_added_for_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['checkpoint-cleanup']['pages removed']` — counter; family: `mongodb_server_status_wired_tiger_checkpoint_cleanup_pages_removed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['checkpoint-cleanup']['pages skipped during tree walk']` — counter; family: `mongodb_server_status_wired_tiger_checkpoint_cleanup_pages_skipped_during_tree_walk_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.connection (13)

- `db.serverStatus().wiredTiger.connection['auto adjusting condition resets']` — counter; family: `mongodb_server_status_wired_tiger_connection_auto_adjusting_condition_resets_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.connection['auto adjusting condition wait calls']` — counter; family: `mongodb_server_status_wired_tiger_connection_auto_adjusting_condition_wait_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.connection['auto adjusting condition wait raced to update timeout and skipped updating']` — counter; family: `mongodb_server_status_wired_tiger_connection_auto_adjusting_condition_wait_raced_to_update_timeout_and_skipped_updating_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.connection['detected system time went backwards']` — counter; family: `mongodb_server_status_wired_tiger_connection_detected_system_time_went_backwards_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.connection['memory allocations']` — counter; family: `mongodb_server_status_wired_tiger_connection_memory_allocations_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.connection['memory frees']` — counter; family: `mongodb_server_status_wired_tiger_connection_memory_frees_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.connection['memory re-allocations']` — counter; family: `mongodb_server_status_wired_tiger_connection_memory_re_allocations_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.connection['pthread mutex condition wait calls']` — counter; family: `mongodb_server_status_wired_tiger_connection_pthread_mutex_condition_wait_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.connection['pthread mutex shared lock read-lock calls']` — counter; family: `mongodb_server_status_wired_tiger_connection_pthread_mutex_shared_lock_read_lock_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.connection['pthread mutex shared lock write-lock calls']` — counter; family: `mongodb_server_status_wired_tiger_connection_pthread_mutex_shared_lock_write_lock_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.connection['total fsync I/Os']` — counter; family: `mongodb_server_status_wired_tiger_connection_total_fsync_i_os_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.connection['total read I/Os']` — counter; family: `mongodb_server_status_wired_tiger_connection_total_read_i_os_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.connection['total write I/Os']` — counter; family: `mongodb_server_status_wired_tiger_connection_total_write_i_os_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.cursor (35)

- `db.serverStatus().wiredTiger.cursor['Total number of deleted pages skipped during tree walk']` — counter; family: `mongodb_server_status_wired_tiger_cursor_total_number_of_deleted_pages_skipped_during_tree_walk_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['Total number of entries skipped by cursor next calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_total_number_of_entries_skipped_by_cursor_next_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['Total number of entries skipped by cursor prev calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_total_number_of_entries_skipped_by_cursor_prev_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['Total number of in-memory deleted pages skipped during tree walk']` — counter; family: `mongodb_server_status_wired_tiger_cursor_total_number_of_in_memory_deleted_pages_skipped_during_tree_walk_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['Total number of times a search near has exited due to prefix config']` — counter; family: `mongodb_server_status_wired_tiger_cursor_total_number_of_times_a_search_near_has_exited_due_to_prefix_config_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cached cursor count']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cached_cursor_count_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor bulk loaded cursor insert calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_bulk_loaded_cursor_insert_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor close calls that result in cache']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_close_calls_that_result_in_cache_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor create calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_create_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor insert calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_insert_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor insert key and value bytes']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_insert_key_and_value_bytes_total`; raw unit: bytes; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.cursor['cursor modify calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_modify_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor modify key and value bytes affected']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_modify_key_and_value_bytes_affected_total`; raw unit: bytes; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.cursor['cursor modify value bytes modified']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_modify_value_bytes_modified_total`; raw unit: bytes; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.cursor['cursor next calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_next_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor next calls that skip greater than or equal to 100 entries']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_next_calls_that_skip_greater_than_or_equal_to_100_entries_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor next calls that skip less than 100 entries']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_next_calls_that_skip_less_than_100_entries_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor operation restarted']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_operation_restarted_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor prev calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_prev_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor prev calls that skip greater than or equal to 100 entries']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_prev_calls_that_skip_greater_than_or_equal_to_100_entries_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor prev calls that skip less than 100 entries']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_prev_calls_that_skip_less_than_100_entries_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor remove calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_remove_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor remove key bytes removed']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_remove_key_bytes_removed_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor reserve calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_reserve_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor reset calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_reset_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor search calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_search_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor search near calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_search_near_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor sweep cursors closed']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_sweep_cursors_closed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor sweep cursors examined']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_sweep_cursors_examined_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.cursor['cursor sweeps']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_sweeps_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor truncate calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_truncate_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor update calls']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_update_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.cursor['cursor update key and value bytes']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_update_key_and_value_bytes_total`; raw unit: bytes; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.cursor['cursor update value size change']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursor_update_value_size_change_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.cursor['cursors reused from cache']` — counter; family: `mongodb_server_status_wired_tiger_cursor_cursors_reused_from_cache_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default

### wiredTiger.data-handle (7)

- `db.serverStatus().wiredTiger['data-handle']['connection sweep candidate became referenced']` — counter; family: `mongodb_server_status_wired_tiger_data_handle_connection_sweep_candidate_became_referenced_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger['data-handle']['connection sweep dhandles closed']` — counter; family: `mongodb_server_status_wired_tiger_data_handle_connection_sweep_dhandles_closed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['data-handle']['connection sweep time-of-death sets']` — counter; family: `mongodb_server_status_wired_tiger_data_handle_connection_sweep_time_of_death_sets_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger['data-handle']['connection sweeps']` — counter; family: `mongodb_server_status_wired_tiger_data_handle_connection_sweeps_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['data-handle']['connection sweeps skipped due to checkpoint gathering handles']` — counter; family: `mongodb_server_status_wired_tiger_data_handle_connection_sweeps_skipped_due_to_checkpoint_gathering_handles_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['data-handle']['session dhandles swept']` — counter; family: `mongodb_server_status_wired_tiger_data_handle_session_dhandles_swept_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger['data-handle']['session sweep attempts']` — counter; family: `mongodb_server_status_wired_tiger_data_handle_session_sweep_attempts_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.lock (12)

- `db.serverStatus().wiredTiger.lock['checkpoint lock application thread wait time (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_checkpoint_lock_application_thread_wait_time_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['checkpoint lock internal thread wait time (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_checkpoint_lock_internal_thread_wait_time_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['dhandle lock application thread time waiting (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_dhandle_lock_application_thread_time_waiting_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['dhandle lock internal thread time waiting (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_dhandle_lock_internal_thread_time_waiting_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['metadata lock application thread wait time (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_metadata_lock_application_thread_wait_time_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['metadata lock internal thread wait time (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_metadata_lock_internal_thread_wait_time_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['schema lock application thread wait time (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_schema_lock_application_thread_wait_time_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['schema lock internal thread wait time (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_schema_lock_internal_thread_wait_time_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['table lock application thread time waiting for the table lock (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_table_lock_application_thread_time_waiting_for_the_table_lock_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['table lock internal thread time waiting for the table lock (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_table_lock_internal_thread_time_waiting_for_the_table_lock_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['txn global lock application thread time waiting (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_txn_global_lock_application_thread_time_waiting_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.lock['txn global lock internal thread time waiting (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_lock_txn_global_lock_internal_thread_time_waiting_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.log (31)

- `db.serverStatus().wiredTiger.log['busy returns attempting to switch slots']` — counter; family: `mongodb_server_status_wired_tiger_log_busy_returns_attempting_to_switch_slots_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['force archive time sleeping (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_log_force_archive_time_sleeping_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['log files manually zero-filled']` — counter; family: `mongodb_server_status_wired_tiger_log_log_files_manually_zero_filled_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.log['log force write operations']` — counter; family: `mongodb_server_status_wired_tiger_log_log_force_write_operations_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['log force write operations skipped']` — counter; family: `mongodb_server_status_wired_tiger_log_log_force_write_operations_skipped_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['log records too small to compress']` — counter; family: `mongodb_server_status_wired_tiger_log_log_records_too_small_to_compress_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['log release advances write LSN']` — counter; family: `mongodb_server_status_wired_tiger_log_log_release_advances_write_lsn_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['log server thread advances write LSN']` — counter; family: `mongodb_server_status_wired_tiger_log_log_server_thread_advances_write_lsn_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['log server thread write LSN walk skipped']` — counter; family: `mongodb_server_status_wired_tiger_log_log_server_thread_write_lsn_walk_skipped_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['log sync time duration (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_log_log_sync_time_duration_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['log sync_dir time duration (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_log_log_sync_dir_time_duration_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['logging bytes consolidated']` — counter; family: `mongodb_server_status_wired_tiger_log_logging_bytes_consolidated_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['number of pre-allocated log files to create']` — counter; family: `mongodb_server_status_wired_tiger_log_number_of_pre_allocated_log_files_to_create_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['pre-allocated log files prepared']` — counter; family: `mongodb_server_status_wired_tiger_log_pre_allocated_log_files_prepared_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.log['pre-allocated log files used']` — counter; family: `mongodb_server_status_wired_tiger_log_pre_allocated_log_files_used_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.log['slot close lost race']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_close_lost_race_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.log['slot close unbuffered waits']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_close_unbuffered_waits_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['slot closures']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_closures_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.log['slot join atomic update races']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_join_atomic_update_races_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['slot join calls atomic updates raced']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_join_calls_atomic_updates_raced_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['slot join calls did not yield']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_join_calls_did_not_yield_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['slot join calls slept']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_join_calls_slept_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['slot join calls yielded']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_join_calls_yielded_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['slot joins yield time (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_joins_yield_time_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['slot transitions unable to find free slot']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_transitions_unable_to_find_free_slot_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.log['slot unbuffered writes']` — counter; family: `mongodb_server_status_wired_tiger_log_slot_unbuffered_writes_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['total in-memory size of compressed records']` — counter; family: `mongodb_server_status_wired_tiger_log_total_in_memory_size_of_compressed_records_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['total log buffer size']` — counter; family: `mongodb_server_status_wired_tiger_log_total_log_buffer_size_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['total size of compressed records']` — counter; family: `mongodb_server_status_wired_tiger_log_total_size_of_compressed_records_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.log['written slots coalesced']` — counter; family: `mongodb_server_status_wired_tiger_log_written_slots_coalesced_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.log['yields waiting for previous log file close']` — counter; family: `mongodb_server_status_wired_tiger_log_yields_waiting_for_previous_log_file_close_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.reconciliation (15)

- `db.serverStatus().wiredTiger.reconciliation['approximate byte size of transaction IDs in pages written']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_approximate_byte_size_of_transaction_ids_in_pages_written_total`; raw unit: bytes; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['fast-path pages deleted']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_fast_path_pages_deleted_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['leaf-page overflow keys']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_leaf_page_overflow_keys_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.reconciliation['page reconciliation calls']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_page_reconciliation_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['page reconciliation calls for eviction']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_page_reconciliation_calls_for_eviction_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['page reconciliation calls that resulted in values with prepared transaction metadata']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_page_reconciliation_calls_that_resulted_in_values_with_prepared_transaction_metadata_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['page reconciliation calls that resulted in values with transaction ids']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_page_reconciliation_calls_that_resulted_in_values_with_transaction_ids_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['pages deleted']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_pages_deleted_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated newest stop transaction ID']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_an_aggregated_newest_stop_transaction_id_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated newest transaction ID ']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_an_aggregated_newest_transaction_id_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated prepare']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_an_aggregated_prepare_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one start transaction ID']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_at_least_one_start_transaction_id_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one stop transaction ID']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_at_least_one_stop_transaction_id_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['records written including a start transaction ID']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_records_written_including_a_start_transaction_id_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.reconciliation['records written including a stop transaction ID']` — counter; family: `mongodb_server_status_wired_tiger_reconciliation_records_written_including_a_stop_transaction_id_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.session (25)

- `db.serverStatus().wiredTiger.session['attempts to remove a local object and the object is in use']` — counter; family: `mongodb_server_status_wired_tiger_session_attempts_to_remove_a_local_object_and_the_object_is_in_use_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['flush_tier operation calls']` — counter; family: `mongodb_server_status_wired_tiger_session_flush_tier_operation_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['local objects removed']` — counter; family: `mongodb_server_status_wired_tiger_session_local_objects_removed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table alter failed calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_alter_failed_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table alter successful calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_alter_successful_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table alter triggering checkpoint calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_alter_triggering_checkpoint_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table alter unchanged and skipped']` — counter; family: `mongodb_server_status_wired_tiger_session_table_alter_unchanged_and_skipped_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table compact failed calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_compact_failed_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table compact skipped as process would not reduce file size']` — counter; family: `mongodb_server_status_wired_tiger_session_table_compact_skipped_as_process_would_not_reduce_file_size_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table compact successful calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_compact_successful_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table compact timeout']` — counter; family: `mongodb_server_status_wired_tiger_session_table_compact_timeout_total`; raw unit: count or raw value; conversion: 1; confidence: researched; evidence: cumulative WiredTiger event count
- `db.serverStatus().wiredTiger.session['table create failed calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_create_failed_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table create successful calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_create_successful_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table drop failed calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_drop_failed_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table drop successful calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_drop_successful_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table rename failed calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_rename_failed_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table rename successful calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_rename_successful_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table salvage failed calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_salvage_failed_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table salvage successful calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_salvage_successful_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table truncate failed calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_truncate_failed_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table truncate successful calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_truncate_successful_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table verify failed calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_verify_failed_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['table verify successful calls']` — counter; family: `mongodb_server_status_wired_tiger_session_table_verify_successful_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['tiered operations dequeued and processed']` — counter; family: `mongodb_server_status_wired_tiger_session_tiered_operations_dequeued_and_processed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.session['tiered operations scheduled']` — counter; family: `mongodb_server_status_wired_tiger_session_tiered_operations_scheduled_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.thread-yield (14)

- `db.serverStatus().wiredTiger['thread-yield']['application thread time evicting (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_application_thread_time_evicting_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['application thread time waiting for cache (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_application_thread_time_waiting_for_cache_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['connection close blocked waiting for transaction state stabilization']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_connection_close_blocked_waiting_for_transaction_state_stabilization_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['connection close yielded for lsm manager shutdown']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_connection_close_yielded_for_lsm_manager_shutdown_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['data handle lock yielded']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_data_handle_lock_yielded_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['get reference for page index and slot time sleeping (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_get_reference_for_page_index_and_slot_time_sleeping_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['page access yielded due to prepare state change']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_page_access_yielded_due_to_prepare_state_change_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['page acquire busy blocked']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_page_acquire_busy_blocked_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['page acquire eviction blocked']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_page_acquire_eviction_blocked_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['page acquire locked blocked']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_page_acquire_locked_blocked_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['page acquire read blocked']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_page_acquire_read_blocked_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['page acquire time sleeping (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_page_acquire_time_sleeping_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['page delete rollback time sleeping for state change (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_page_delete_rollback_time_sleeping_for_state_change_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger['thread-yield']['page reconciliation yielded due to child modification']` — counter; family: `mongodb_server_status_wired_tiger_thread_yield_page_reconciliation_yielded_due_to_child_modification_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

### wiredTiger.transaction (31)

- `db.serverStatus().wiredTiger.transaction['Number of prepared updates']` — counter; family: `mongodb_server_status_wired_tiger_transaction_number_of_prepared_updates_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['Number of prepared updates committed']` — counter; family: `mongodb_server_status_wired_tiger_transaction_number_of_prepared_updates_committed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['Number of prepared updates repeated on the same key']` — counter; family: `mongodb_server_status_wired_tiger_transaction_number_of_prepared_updates_repeated_on_the_same_key_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['Number of prepared updates rolled back']` — counter; family: `mongodb_server_status_wired_tiger_transaction_number_of_prepared_updates_rolled_back_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['prepared transactions']` — counter; family: `mongodb_server_status_wired_tiger_transaction_prepared_transactions_total`; raw unit: count or raw value; conversion: 1; confidence: researched; evidence: cumulative WiredTiger event count
- `db.serverStatus().wiredTiger.transaction['prepared transactions committed']` — counter; family: `mongodb_server_status_wired_tiger_transaction_prepared_transactions_committed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['prepared transactions rolled back']` — counter; family: `mongodb_server_status_wired_tiger_transaction_prepared_transactions_rolled_back_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.transaction['race to read prepared update retry']` — counter; family: `mongodb_server_status_wired_tiger_transaction_race_to_read_prepared_update_retry_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable calls']` — counter; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_calls_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable keys removed']` — counter; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_keys_removed_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable keys restored']` — counter; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_keys_restored_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable skipping delete rle']` — counter; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_skipping_delete_rle_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable skipping stable rle']` — counter; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_skipping_stable_rle_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable tree walk skipping pages']` — counter; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_tree_walk_skipping_pages_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable updates aborted']` — counter; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_updates_aborted_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent duration for gathering all handles (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_most_recent_duration_for_gathering_all_handles_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent duration for gathering applied handles (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_most_recent_duration_for_gathering_applied_handles_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent duration for gathering skipped handles (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_most_recent_duration_for_gathering_skipped_handles_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent handles applied']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_most_recent_handles_applied_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent handles skipped']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_most_recent_handles_skipped_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent handles walked']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_most_recent_handles_walked_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint most recent time (msecs)']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_most_recent_time_msecs_total`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint prepare most recent time (msecs)']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_prepare_most_recent_time_msecs_total`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint prepare total time (msecs)']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_prepare_total_time_msecs_total`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint scrub dirty target']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_scrub_dirty_target_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint scrub time (msecs)']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_scrub_time_msecs_total`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoints due to obsolete pages']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoints_due_to_obsolete_pages_total`; raw unit: count or raw value; conversion: 1; confidence: source-rule; evidence: WiredTiger 5.0 statistic without snapshot wording; upstream stat definitions specify operation-rate semantics by default
- `db.serverStatus().wiredTiger.transaction['transaction checkpoints skipped because database was clean']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoints_skipped_because_database_was_clean_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction fsync calls for checkpoint after allocating the transaction ID']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_fsync_calls_for_checkpoint_after_allocating_the_transaction_id_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['transaction fsync duration for checkpoint after allocating the transaction ID (usecs)']` — counter; family: `mongodb_server_status_wired_tiger_transaction_transaction_fsync_duration_for_checkpoint_after_allocating_the_transaction_id_usecs_total`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming
- `db.serverStatus().wiredTiger.transaction['update conflicts']` — counter; family: `mongodb_server_status_wired_tiger_transaction_update_conflicts_total`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: cumulative event, time, or byte naming

## Generated gauge proposals

### $clusterTime.clusterTime[0] (1)

- `db.serverStatus().$clusterTime.clusterTime[0]` — gauge; family: `mongodb_server_status_cluster_time_cluster_time_0`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### $clusterTime.clusterTime[1] (1)

- `db.serverStatus().$clusterTime.clusterTime[1]` — gauge; family: `mongodb_server_status_cluster_time_cluster_time_1`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### $clusterTime.signature (1)

- `db.serverStatus().$clusterTime.signature.keyId` — gauge; family: `mongodb_server_status_cluster_time_signature_key_id`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### $configServerState.opTime (3)

- `db.serverStatus().$configServerState.opTime.t` — gauge; family: `mongodb_server_status_config_server_state_op_time_t`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot
- `db.serverStatus().$configServerState.opTime.ts[0]` — gauge; family: `mongodb_server_status_config_server_state_op_time_ts_0`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot
- `db.serverStatus().$configServerState.opTime.ts[1]` — gauge; family: `mongodb_server_status_config_server_state_op_time_ts_1`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### $gleStats.lastOpTime[0] (1)

- `db.serverStatus().$gleStats.lastOpTime[0]` — gauge; family: `mongodb_server_status_gle_stats_last_op_time_0`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### $gleStats.lastOpTime[1] (1)

- `db.serverStatus().$gleStats.lastOpTime[1]` — gauge; family: `mongodb_server_status_gle_stats_last_op_time_1`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### catalogStats.capped (1)

- `db.serverStatus().catalogStats.capped` — gauge; family: `mongodb_server_status_catalog_stats_capped`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current catalog population

### catalogStats.collections (1)

- `db.serverStatus().catalogStats.collections` — gauge; family: `mongodb_server_status_catalog_stats_collections`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current catalog population

### catalogStats.internalCollections (1)

- `db.serverStatus().catalogStats.internalCollections` — gauge; family: `mongodb_server_status_catalog_stats_internal_collections`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current catalog population

### catalogStats.internalViews (1)

- `db.serverStatus().catalogStats.internalViews` — gauge; family: `mongodb_server_status_catalog_stats_internal_views`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current catalog population

### catalogStats.timeseries (1)

- `db.serverStatus().catalogStats.timeseries` — gauge; family: `mongodb_server_status_catalog_stats_timeseries`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current catalog population

### catalogStats.views (1)

- `db.serverStatus().catalogStats.views` — gauge; family: `mongodb_server_status_catalog_stats_views`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current catalog population

### connections.exhaustHello (1)

- `db.serverStatus().connections.exhaustHello` — gauge; family: `mongodb_server_status_connections_exhaust_hello`; raw unit: count or raw value; conversion: 1; confidence: researched; evidence: current connections whose last request is exhaust hello

### connections.exhaustIsMaster (1)

- `db.serverStatus().connections.exhaustIsMaster` — gauge; family: `mongodb_server_status_connections_exhaust_is_master`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### extra_info.maximum_resident_set_kb (1)

- `db.serverStatus().extra_info.maximum_resident_set_kb` — gauge; family: `mongodb_server_status_extra_info_maximum_resident_set_kb`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### featureCompatibilityVersion.major (1)

- `db.serverStatus().featureCompatibilityVersion.major` — gauge; family: `mongodb_server_status_feature_compatibility_version_major`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### featureCompatibilityVersion.minor (1)

- `db.serverStatus().featureCompatibilityVersion.minor` — gauge; family: `mongodb_server_status_feature_compatibility_version_minor`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### featureCompatibilityVersion.transitioning (1)

- `db.serverStatus().featureCompatibilityVersion.transitioning` — gauge; family: `mongodb_server_status_feature_compatibility_version_transitioning`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### globalLock.activeClients (1)

- `db.serverStatus().globalLock.activeClients.total` — gauge; family: `mongodb_server_status_global_lock_active_clients_total_value`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### globalLock.currentQueue (1)

- `db.serverStatus().globalLock.currentQueue.total` — gauge; family: `mongodb_server_status_global_lock_current_queue_total_value`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### indexStats.count (1)

- `db.serverStatus().indexStats.count` — gauge; family: `mongodb_server_status_index_stats_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population

### indexStats.features (14)

- `db.serverStatus().indexStats.features['2d'].count` — gauge; family: `mongodb_server_status_index_stats_features_2d_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features['2dsphere'].count` — gauge; family: `mongodb_server_status_index_stats_features_2dsphere_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.collation.count` — gauge; family: `mongodb_server_status_index_stats_features_collation_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.compound.count` — gauge; family: `mongodb_server_status_index_stats_features_compound_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.hashed.count` — gauge; family: `mongodb_server_status_index_stats_features_hashed_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.id.count` — gauge; family: `mongodb_server_status_index_stats_features_id_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.normal.count` — gauge; family: `mongodb_server_status_index_stats_features_normal_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.partial.count` — gauge; family: `mongodb_server_status_index_stats_features_partial_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.single.count` — gauge; family: `mongodb_server_status_index_stats_features_single_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.sparse.count` — gauge; family: `mongodb_server_status_index_stats_features_sparse_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.text.count` — gauge; family: `mongodb_server_status_index_stats_features_text_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.ttl.count` — gauge; family: `mongodb_server_status_index_stats_features_ttl_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.unique.count` — gauge; family: `mongodb_server_status_index_stats_features_unique_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population
- `db.serverStatus().indexStats.features.wildcard.count` — gauge; family: `mongodb_server_status_index_stats_features_wildcard_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current index population

### mem.bits (1)

- `db.serverStatus().mem.bits` — gauge; family: `mongodb_server_status_mem_bits`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### mem.supported (1)

- `db.serverStatus().mem.supported` — gauge; family: `mongodb_server_status_mem_supported`; raw unit: boolean; conversion: 1; confidence: high; evidence: boolean snapshot

### metrics.getLastError (1)

- `db.serverStatus().metrics.getLastError.default.unsatisfiable` — gauge; family: `mongodb_server_status_metrics_get_last_error_default_unsatisfiable`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### metrics.repl (6)

- `db.serverStatus().metrics.repl.apply.batchSize` — gauge; family: `mongodb_server_status_metrics_repl_apply_batch_size`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current/configured snapshot
- `db.serverStatus().metrics.repl.executor.pool.inProgressCount` — gauge; family: `mongodb_server_status_metrics_repl_executor_pool_in_progress_count`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current/configured snapshot
- `db.serverStatus().metrics.repl.executor.shuttingDown` — gauge; family: `mongodb_server_status_metrics_repl_executor_shutting_down`; raw unit: boolean; conversion: 1; confidence: high; evidence: boolean snapshot
- `db.serverStatus().metrics.repl.stateTransition.userOperationsRunning` — gauge; family: `mongodb_server_status_metrics_repl_state_transition_user_operations_running`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current/configured snapshot
- `db.serverStatus().metrics.repl.waiters.opTime` — gauge; family: `mongodb_server_status_metrics_repl_waiters_op_time`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: current/configured snapshot
- `db.serverStatus().metrics.repl.waiters.replication` — gauge; family: `mongodb_server_status_metrics_repl_waiters_replication`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current/configured snapshot

### ok (1)

- `db.serverStatus().ok` — gauge; family: `mongodb_server_status_ok`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### pid (1)

- `db.serverStatus().pid` — gauge; family: `mongodb_server_status_pid`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### scramCache.SCRAM-SHA-1 (1)

- `db.serverStatus().scramCache['SCRAM-SHA-1'].misses` — gauge; family: `mongodb_server_status_scram_cache_scram_sha_1_misses`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### scramCache.SCRAM-SHA-256 (1)

- `db.serverStatus().scramCache['SCRAM-SHA-256'].misses` — gauge; family: `mongodb_server_status_scram_cache_scram_sha_256_misses`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### sharding.lastSeenConfigServerOpTime (3)

- `db.serverStatus().sharding.lastSeenConfigServerOpTime.t` — gauge; family: `mongodb_server_status_sharding_last_seen_config_server_op_time_t`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot
- `db.serverStatus().sharding.lastSeenConfigServerOpTime.ts[0]` — gauge; family: `mongodb_server_status_sharding_last_seen_config_server_op_time_ts_0`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot
- `db.serverStatus().sharding.lastSeenConfigServerOpTime.ts[1]` — gauge; family: `mongodb_server_status_sharding_last_seen_config_server_op_time_ts_1`; raw unit: raw timestamp component; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### sharding.maxChunkSizeInBytes (1)

- `db.serverStatus().sharding.maxChunkSizeInBytes` — gauge; family: `mongodb_server_status_sharding_max_chunk_size_in_bytes`; raw unit: bytes; conversion: 1; confidence: high; evidence: current/configured snapshot

### storageEngine.oldestRequiredTimestampForCrashRecovery[0] (1)

- `db.serverStatus().storageEngine.oldestRequiredTimestampForCrashRecovery[0]` — gauge; family: `mongodb_server_status_storage_engine_oldest_required_timestamp_for_crash_recovery_0`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### storageEngine.oldestRequiredTimestampForCrashRecovery[1] (1)

- `db.serverStatus().storageEngine.oldestRequiredTimestampForCrashRecovery[1]` — gauge; family: `mongodb_server_status_storage_engine_oldest_required_timestamp_for_crash_recovery_1`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### tcmalloc.tcmalloc (1)

- `db.serverStatus().tcmalloc.tcmalloc.release_rate` — gauge; family: `mongodb_server_status_tcmalloc_tcmalloc_release_rate`; raw unit: raw rate; conversion: 1; confidence: high; evidence: current/configured snapshot

### tenantMigrations.currentMigrationsDonating (1)

- `db.serverStatus().tenantMigrations.currentMigrationsDonating` — gauge; family: `mongodb_server_status_tenant_migrations_current_migrations_donating`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### tenantMigrations.currentMigrationsReceiving (1)

- `db.serverStatus().tenantMigrations.currentMigrationsReceiving` — gauge; family: `mongodb_server_status_tenant_migrations_current_migrations_receiving`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### trafficRecording.running (1)

- `db.serverStatus().trafficRecording.running` — gauge; family: `mongodb_server_status_traffic_recording_running`; raw unit: boolean; conversion: 1; confidence: high; evidence: boolean snapshot

### twoPhaseCommitCoordinator.currentInSteps (5)

- `db.serverStatus().twoPhaseCommitCoordinator.currentInSteps.deletingCoordinatorDoc` — gauge; family: `mongodb_server_status_two_phase_commit_coordinator_current_in_steps_deleting_coordinator_doc`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current coordinator population
- `db.serverStatus().twoPhaseCommitCoordinator.currentInSteps.waitingForDecisionAcks` — gauge; family: `mongodb_server_status_two_phase_commit_coordinator_current_in_steps_waiting_for_decision_acks`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current coordinator population
- `db.serverStatus().twoPhaseCommitCoordinator.currentInSteps.waitingForVotes` — gauge; family: `mongodb_server_status_two_phase_commit_coordinator_current_in_steps_waiting_for_votes`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current coordinator population
- `db.serverStatus().twoPhaseCommitCoordinator.currentInSteps.writingDecision` — gauge; family: `mongodb_server_status_two_phase_commit_coordinator_current_in_steps_writing_decision`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current coordinator population
- `db.serverStatus().twoPhaseCommitCoordinator.currentInSteps.writingParticipantList` — gauge; family: `mongodb_server_status_two_phase_commit_coordinator_current_in_steps_writing_participant_list`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current coordinator population

### uptimeEstimate (1)

- `db.serverStatus().uptimeEstimate` — gauge; family: `mongodb_server_status_uptime_estimate`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### uptimeMillis (1)

- `db.serverStatus().uptimeMillis` — gauge; family: `mongodb_server_status_uptime_millis`; raw unit: milliseconds; conversion: 1; confidence: high; evidence: clock, identity, configuration, or state snapshot

### wiredTiger.block-manager (2)

- `db.serverStatus().wiredTiger['block-manager']['block cache number of hits including existence checks']` — gauge; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_number_of_hits_including_existence_checks`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger['block-manager']['block cache number of misses including existence checks']` — gauge; family: `mongodb_server_status_wired_tiger_block_manager_block_cache_number_of_misses_including_existence_checks`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.cache (59)

- `db.serverStatus().wiredTiger.cache['application threads page read from disk to cache count']` — gauge; family: `mongodb_server_status_wired_tiger_cache_application_threads_page_read_from_disk_to_cache_count`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['application threads page read from disk to cache time (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_cache_application_threads_page_read_from_disk_to_cache_time_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['application threads page write from cache to disk count']` — gauge; family: `mongodb_server_status_wired_tiger_cache_application_threads_page_write_from_cache_to_disk_count`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['application threads page write from cache to disk time (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_cache_application_threads_page_write_from_cache_to_disk_time_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['bytes belonging to page images in the cache']` — gauge; family: `mongodb_server_status_wired_tiger_cache_bytes_belonging_to_page_images_in_the_cache`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['bytes belonging to the history store table in the cache']` — gauge; family: `mongodb_server_status_wired_tiger_cache_bytes_belonging_to_the_history_store_table_in_the_cache`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['bytes dirty in the cache cumulative']` — gauge; family: `mongodb_server_status_wired_tiger_cache_bytes_dirty_in_the_cache_cumulative`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['bytes not belonging to page images in the cache']` — gauge; family: `mongodb_server_status_wired_tiger_cache_bytes_not_belonging_to_page_images_in_the_cache`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['cache overflow score']` — gauge; family: `mongodb_server_status_wired_tiger_cache_cache_overflow_score`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['checkpoint of history store file blocked non-history store page eviction']` — gauge; family: `mongodb_server_status_wired_tiger_cache_checkpoint_of_history_store_file_blocked_non_history_store_page_eviction`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction currently operating in aggressive mode']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_currently_operating_in_aggressive_mode`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction empty score']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_empty_score`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction gave up due to detecting an out of order on disk value behind the last update on the chain']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_gave_up_due_to_detecting_an_out_of_order_on_disk_value_behind_the_last_update_on_the_chain`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction gave up due to detecting an out of order tombstone ahead of the selected on disk update']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_gave_up_due_to_detecting_an_out_of_order_tombstone_ahead_of_the_selected_on_disk_update`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction gave up due to detecting an out of order tombstone ahead of the selected on disk update after validating the update chain']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_gave_up_due_to_detecting_an_out_of_order_tombstone_ahead_of_the_selected_on_disk_update_after_validating_the_update_chain`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction gave up due to detecting out of order timestamps on the update chain after the selected on disk update']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_gave_up_due_to_detecting_out_of_order_timestamps_on_the_update_chain_after_the_selected_on_disk_update`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction gave up due to needing to remove a record from the history store but checkpoint is running']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_gave_up_due_to_needing_to_remove_a_record_from_the_history_store_but_checkpoint_is_running`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction server skips dirty pages during a running checkpoint']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_server_skips_dirty_pages_during_a_running_checkpoint`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction server skips metadata pages with history']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_server_skips_metadata_pages_with_history`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction server skips pages that are written with transactions greater than the last running']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_server_skips_pages_that_are_written_with_transactions_greater_than_the_last_running`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction server skips trees because there are too many active walks']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_server_skips_trees_because_there_are_too_many_active_walks`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction server skips trees that are configured to stick in cache']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_server_skips_trees_that_are_configured_to_stick_in_cache`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction server skips trees that disable eviction']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_server_skips_trees_that_disable_eviction`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction state']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_state`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction walk target pages reduced due to history store cache pressure']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_walk_target_pages_reduced_due_to_history_store_cache_pressure`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction walk target strategy both clean and dirty pages']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_walk_target_strategy_both_clean_and_dirty_pages`; raw unit: raw rate; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction walk target strategy only dirty pages']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_walk_target_strategy_only_dirty_pages`; raw unit: raw rate; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['eviction worker thread active']` — gauge; family: `mongodb_server_status_wired_tiger_cache_eviction_worker_thread_active`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['files with active eviction walks']` — gauge; family: `mongodb_server_status_wired_tiger_cache_files_with_active_eviction_walks`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['forced eviction - history store pages failed to evict while session has history store cursor open']` — gauge; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_history_store_pages_failed_to_evict_while_session_has_history_store_cursor_open`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['forced eviction - history store pages selected while session has history store cursor open']` — gauge; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_history_store_pages_selected_while_session_has_history_store_cursor_open`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['forced eviction - history store pages successfully evicted while session has history store cursor open']` — gauge; family: `mongodb_server_status_wired_tiger_cache_forced_eviction_history_store_pages_successfully_evicted_while_session_has_history_store_cursor_open`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['hazard pointer maximum array length']` — gauge; family: `mongodb_server_status_wired_tiger_cache_hazard_pointer_maximum_array_length`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store score']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_score`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table insert calls']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_insert_calls`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table insert calls that returned restart']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_insert_calls_that_returned_restart`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table max on-disk size']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_max_on_disk_size`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table on-disk size']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_on_disk_size`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table out-of-order resolved updates that lose their durable timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_out_of_order_resolved_updates_that_lose_their_durable_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table out-of-order updates that were fixed up by reinserting with the fixed timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_out_of_order_updates_that_were_fixed_up_by_reinserting_with_the_fixed_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table reads']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_reads`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table reads missed']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_reads_missed`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table reads requiring squashed modifies']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_reads_requiring_squashed_modifies`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table truncation by rollback to stable to remove an unstable update']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_truncation_by_rollback_to_stable_to_remove_an_unstable_update`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table truncation by rollback to stable to remove an update']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_truncation_by_rollback_to_stable_to_remove_an_update`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table truncation to remove an update']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_truncation_to_remove_an_update`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table truncation to remove range of updates due to key being removed from the data page during reconciliation']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_truncation_to_remove_range_of_updates_due_to_key_being_removed_from_the_data_page_during_reconciliation`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table truncation to remove range of updates due to out-of-order timestamp update on data page']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_truncation_to_remove_range_of_updates_due_to_out_of_order_timestamp_update_on_data_page`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['history store table writes requiring squashed modifies']` — gauge; family: `mongodb_server_status_wired_tiger_cache_history_store_table_writes_requiring_squashed_modifies`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['maximum milliseconds spent at a single eviction']` — gauge; family: `mongodb_server_status_wired_tiger_cache_maximum_milliseconds_spent_at_a_single_eviction`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['maximum page size seen at eviction']` — gauge; family: `mongodb_server_status_wired_tiger_cache_maximum_page_size_seen_at_eviction`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['page written requiring history store records']` — gauge; family: `mongodb_server_status_wired_tiger_cache_page_written_requiring_history_store_records`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['pages queued for urgent eviction from history store due to high dirty content']` — gauge; family: `mongodb_server_status_wired_tiger_cache_pages_queued_for_urgent_eviction_from_history_store_due_to_high_dirty_content`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['pages read into cache after truncate in prepare state']` — gauge; family: `mongodb_server_status_wired_tiger_cache_pages_read_into_cache_after_truncate_in_prepare_state`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['pages selected for eviction unable to be evicted because of active children on an internal page']` — gauge; family: `mongodb_server_status_wired_tiger_cache_pages_selected_for_eviction_unable_to_be_evicted_because_of_active_children_on_an_internal_page`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['pages selected for eviction unable to be evicted because of race between checkpoint and out of order timestamps handling']` — gauge; family: `mongodb_server_status_wired_tiger_cache_pages_selected_for_eviction_unable_to_be_evicted_because_of_race_between_checkpoint_and_out_of_order_timestamps_handling`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['the number of times full update inserted to history store']` — gauge; family: `mongodb_server_status_wired_tiger_cache_the_number_of_times_full_update_inserted_to_history_store`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['the number of times reverse modify inserted to history store']` — gauge; family: `mongodb_server_status_wired_tiger_cache_the_number_of_times_reverse_modify_inserted_to_history_store`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cache['total milliseconds spent inside reentrant history store evictions in a reconciliation']` — gauge; family: `mongodb_server_status_wired_tiger_cache_total_milliseconds_spent_inside_reentrant_history_store_evictions_in_a_reconciliation`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.capacity (14)

- `db.serverStatus().wiredTiger.capacity['background fsync file handles considered']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_background_fsync_file_handles_considered`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['background fsync file handles synced']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_background_fsync_file_handles_synced`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['background fsync time (msecs)']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_background_fsync_time_msecs`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['bytes read']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_bytes_read`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['bytes written for checkpoint']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_bytes_written_for_checkpoint`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['bytes written for eviction']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_bytes_written_for_eviction`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['bytes written for log']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_bytes_written_for_log`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['bytes written total']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_bytes_written_total_value`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['threshold to call fsync']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_threshold_to_call_fsync`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['time waiting due to total capacity (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_time_waiting_due_to_total_capacity_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['time waiting during checkpoint (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_time_waiting_during_checkpoint_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['time waiting during eviction (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_time_waiting_during_eviction_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['time waiting during logging (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_time_waiting_during_logging_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.capacity['time waiting during read (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_capacity_time_waiting_during_read_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.checkpoint-cleanup (1)

- `db.serverStatus().wiredTiger['checkpoint-cleanup']['pages visited']` — gauge; family: `mongodb_server_status_wired_tiger_checkpoint_cleanup_pages_visited`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.connection (1)

- `db.serverStatus().wiredTiger.connection['files currently open']` — gauge; family: `mongodb_server_status_wired_tiger_connection_files_currently_open`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.connection['hash bucket array size for data handles']` — gauge; family: `mongodb_server_status_wired_tiger_connection_hash_bucket_array_size_for_data_handles`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current hash-table structure size, not a histogram bucket
- `db.serverStatus().wiredTiger.connection['hash bucket array size general']` — gauge; family: `mongodb_server_status_wired_tiger_connection_hash_bucket_array_size_general`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current hash-table structure size, not a histogram bucket
- `db.serverStatus().wiredTiger.cursor['cursor sweep buckets']` — gauge; family: `mongodb_server_status_wired_tiger_cursor_cursor_sweep_buckets`; raw unit: count or raw value; conversion: 1; confidence: high; evidence: current cursor-sweep bucket configuration, not a histogram bucket

### wiredTiger.cursor (5)

- `db.serverStatus().wiredTiger.cursor['Total number of entries skipped to position the history store cursor']` — gauge; family: `mongodb_server_status_wired_tiger_cursor_total_number_of_entries_skipped_to_position_the_history_store_cursor`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cursor['cursor next calls that skip due to a globally visible history store tombstone']` — gauge; family: `mongodb_server_status_wired_tiger_cursor_cursor_next_calls_that_skip_due_to_a_globally_visible_history_store_tombstone`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cursor['cursor prev calls that skip due to a globally visible history store tombstone']` — gauge; family: `mongodb_server_status_wired_tiger_cursor_cursor_prev_calls_that_skip_due_to_a_globally_visible_history_store_tombstone`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cursor['cursor search history store calls']` — gauge; family: `mongodb_server_status_wired_tiger_cursor_cursor_search_history_store_calls`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.cursor['open cursor count']` — gauge; family: `mongodb_server_status_wired_tiger_cursor_open_cursor_count`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.data-handle (3)

- `db.serverStatus().wiredTiger['data-handle']['connection data handle size']` — gauge; family: `mongodb_server_status_wired_tiger_data_handle_connection_data_handle_size`; raw unit: count or raw value; conversion: 1; confidence: researched; evidence: current data-handle memory size
- `db.serverStatus().wiredTiger['data-handle']['connection data handles currently active']` — gauge; family: `mongodb_server_status_wired_tiger_data_handle_connection_data_handles_currently_active`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger['data-handle']['connection sweep dhandles removed from hash list']` — gauge; family: `mongodb_server_status_wired_tiger_data_handle_connection_sweep_dhandles_removed_from_hash_list`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.lock (17)

- `db.serverStatus().wiredTiger.lock['checkpoint lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_checkpoint_lock_acquisitions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['dhandle read lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_dhandle_read_lock_acquisitions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['dhandle write lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_dhandle_write_lock_acquisitions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['durable timestamp queue lock application thread time waiting (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_lock_durable_timestamp_queue_lock_application_thread_time_waiting_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['durable timestamp queue lock internal thread time waiting (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_lock_durable_timestamp_queue_lock_internal_thread_time_waiting_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['durable timestamp queue read lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_durable_timestamp_queue_read_lock_acquisitions`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['durable timestamp queue write lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_durable_timestamp_queue_write_lock_acquisitions`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['metadata lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_metadata_lock_acquisitions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['read timestamp queue lock application thread time waiting (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_lock_read_timestamp_queue_lock_application_thread_time_waiting_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['read timestamp queue lock internal thread time waiting (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_lock_read_timestamp_queue_lock_internal_thread_time_waiting_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['read timestamp queue read lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_read_timestamp_queue_read_lock_acquisitions`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['read timestamp queue write lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_read_timestamp_queue_write_lock_acquisitions`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['schema lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_schema_lock_acquisitions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['table read lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_table_read_lock_acquisitions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['table write lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_table_write_lock_acquisitions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['txn global read lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_txn_global_read_lock_acquisitions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.lock['txn global write lock acquisitions']` — gauge; family: `mongodb_server_status_wired_tiger_lock_txn_global_write_lock_acquisitions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.log (4)

- `db.serverStatus().wiredTiger.log['maximum log file size']` — gauge; family: `mongodb_server_status_wired_tiger_log_maximum_log_file_size`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.log['pre-allocated log files not ready and missed']` — gauge; family: `mongodb_server_status_wired_tiger_log_pre_allocated_log_files_not_ready_and_missed`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.log['slot join calls found active slot closed']` — gauge; family: `mongodb_server_status_wired_tiger_log_slot_join_calls_found_active_slot_closed`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.log['slot join found active slot closed']` — gauge; family: `mongodb_server_status_wired_tiger_log_slot_join_found_active_slot_closed`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.oplog (2)

- `db.serverStatus().wiredTiger.oplog['visibility timestamp'][0]` — gauge; family: `mongodb_server_status_wired_tiger_oplog_visibility_timestamp_0`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.oplog['visibility timestamp'][1]` — gauge; family: `mongodb_server_status_wired_tiger_oplog_visibility_timestamp_1`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.reconciliation (21)

- `db.serverStatus().wiredTiger.reconciliation['approximate byte size of timestamps in pages written']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_approximate_byte_size_of_timestamps_in_pages_written`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['maximum milliseconds spent in a reconciliation call']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_maximum_milliseconds_spent_in_a_reconciliation_call`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['maximum milliseconds spent in building a disk image in a reconciliation']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_maximum_milliseconds_spent_in_building_a_disk_image_in_a_reconciliation`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['maximum milliseconds spent in moving updates to the history store in a reconciliation']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_maximum_milliseconds_spent_in_moving_updates_to_the_history_store_in_a_reconciliation`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['page reconciliation calls that resulted in values with timestamps']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_page_reconciliation_calls_that_resulted_in_values_with_timestamps`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated newest start durable timestamp ']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_an_aggregated_newest_start_durable_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated newest stop durable timestamp ']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_an_aggregated_newest_stop_durable_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated newest stop timestamp ']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_an_aggregated_newest_stop_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including an aggregated oldest start timestamp ']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_an_aggregated_oldest_start_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one prepare state']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_at_least_one_prepare_state`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one start durable timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_at_least_one_start_durable_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one start timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_at_least_one_start_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one stop durable timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_at_least_one_stop_durable_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['pages written including at least one stop timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_pages_written_including_at_least_one_stop_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['records written including a prepare state']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_records_written_including_a_prepare_state`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['records written including a start durable timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_records_written_including_a_start_durable_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['records written including a start timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_records_written_including_a_start_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['records written including a stop durable timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_records_written_including_a_stop_durable_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['records written including a stop timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_records_written_including_a_stop_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['split bytes currently awaiting free']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_split_bytes_currently_awaiting_free`; raw unit: bytes; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.reconciliation['split objects currently awaiting free']` — gauge; family: `mongodb_server_status_wired_tiger_reconciliation_split_objects_currently_awaiting_free`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.session (4)

- `db.serverStatus().wiredTiger.session['session query timestamp calls']` — gauge; family: `mongodb_server_status_wired_tiger_session_session_query_timestamp_calls`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.session['table compact failed calls due to cache pressure']` — gauge; family: `mongodb_server_status_wired_tiger_session_table_compact_failed_calls_due_to_cache_pressure`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.session['table compact running']` — gauge; family: `mongodb_server_status_wired_tiger_session_table_compact_running`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.session['tiered storage local retention time (secs)']` — gauge; family: `mongodb_server_status_wired_tiger_session_tiered_storage_local_retention_time_secs`; raw unit: seconds; conversion: 1; confidence: researched; evidence: configured retention interval

### wiredTiger.snapshot-window-settings (2)

- `db.serverStatus().wiredTiger['snapshot-window-settings']['min pinned timestamp'][0]` — gauge; family: `mongodb_server_status_wired_tiger_snapshot_window_settings_min_pinned_timestamp_0`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger['snapshot-window-settings']['min pinned timestamp'][1]` — gauge; family: `mongodb_server_status_wired_tiger_snapshot_window_settings_min_pinned_timestamp_1`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.thread-state (3)

- `db.serverStatus().wiredTiger['thread-state']['active filesystem fsync calls']` — gauge; family: `mongodb_server_status_wired_tiger_thread_state_active_filesystem_fsync_calls`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger['thread-state']['active filesystem read calls']` — gauge; family: `mongodb_server_status_wired_tiger_thread_state_active_filesystem_read_calls`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger['thread-state']['active filesystem write calls']` — gauge; family: `mongodb_server_status_wired_tiger_thread_state_active_filesystem_write_calls`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

### wiredTiger.transaction (34)

- `db.serverStatus().wiredTiger.transaction['oldest pinned transaction ID rolled back for eviction']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_oldest_pinned_transaction_id_rolled_back_for_eviction`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['prepared transactions currently active']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_prepared_transactions_currently_active`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['query timestamp calls']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_query_timestamp_calls`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable history store records with stop timestamps older than newer records']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_history_store_records_with_stop_timestamps_older_than_newer_records`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable inconsistent checkpoint']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_inconsistent_checkpoint`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable pages visited']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_pages_visited`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable restored tombstones from history store']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_restored_tombstones_from_history_store`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable restored updates from history store']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_restored_updates_from_history_store`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable sweeping history store keys']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_sweeping_history_store_keys`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['rollback to stable updates removed from history store']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_rollback_to_stable_updates_removed_from_history_store`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['sessions scanned in each walk of concurrent sessions']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_sessions_scanned_in_each_walk_of_concurrent_sessions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['set timestamp calls']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_set_timestamp_calls`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['set timestamp durable calls']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_set_timestamp_durable_calls`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['set timestamp durable updates']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_set_timestamp_durable_updates`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['set timestamp oldest calls']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_set_timestamp_oldest_calls`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['set timestamp oldest updates']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_set_timestamp_oldest_updates`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['set timestamp stable calls']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_set_timestamp_stable_calls`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['set timestamp stable updates']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_set_timestamp_stable_updates`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint currently running for history store file']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_currently_running_for_history_store_file`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint generation']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_generation`; raw unit: count or raw value; conversion: 1; confidence: researched; evidence: checkpoint generation state value
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint history store file duration (usecs)']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_history_store_file_duration_usecs`; raw unit: microseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint prepare currently running']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_prepare_currently_running`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint prepare max time (msecs)']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_prepare_max_time_msecs`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint prepare min time (msecs)']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_prepare_min_time_msecs`; raw unit: milliseconds; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction checkpoint stop timing stress active']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_checkpoint_stop_timing_stress_active`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction range of IDs currently pinned']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_range_of_ids_currently_pinned`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction range of IDs currently pinned by a checkpoint']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_range_of_ids_currently_pinned_by_a_checkpoint`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction range of timestamps currently pinned']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_range_of_timestamps_currently_pinned`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction range of timestamps pinned by a checkpoint']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_range_of_timestamps_pinned_by_a_checkpoint`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction range of timestamps pinned by the oldest active read timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_range_of_timestamps_pinned_by_the_oldest_active_read_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction range of timestamps pinned by the oldest timestamp']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_range_of_timestamps_pinned_by_the_oldest_timestamp`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction read timestamp of the oldest active reader']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_read_timestamp_of_the_oldest_active_reader`; raw unit: raw timestamp component; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction rollback to stable currently running']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_rollback_to_stable_currently_running`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming
- `db.serverStatus().wiredTiger.transaction['transaction walk of concurrent sessions']` — gauge; family: `mongodb_server_status_wired_tiger_transaction_transaction_walk_of_concurrent_sessions`; raw unit: count or raw value; conversion: 1; confidence: medium; evidence: snapshot, limit, or state naming

## Dynamic-key manual review

### metrics.aggStageCounters (39)

- `db.serverStatus().metrics.aggStageCounters.$_internalApplyOplogUpdate` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalBoundedSort` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalConvertBucketIndexStats` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalFindAndModifyImageLookup` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalInhibitOptimization` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalReshardingIterateTransaction` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalReshardingOwnershipMatch` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalSetWindowFields` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalShredDocuments` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalSplitPipeline` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_internalUnpackBucket` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$_unpackBucket` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$addFields` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$bucket` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$bucketAuto` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$changeStream` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$collStats` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$count` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$currentOp` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$documents` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$facet` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$geoNear` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$graphLookup` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$listLocalSessions` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$listSessions` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$merge` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$mergeCursors` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$operationMetrics` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$out` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$planCacheStats` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$queue` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$redact` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$replaceRoot` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$replaceWith` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$setWindowFields` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$skip` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$sortByCount` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$unionWith` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.aggStageCounters.$unset` — bounded grouped-family vocabulary review required; roles: primary, secondary

### metrics.commands (107)

- `db.serverStatus().metrics.commands._flushDatabaseCacheUpdates.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._flushDatabaseCacheUpdates.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._flushDatabaseCacheUpdatesWithWriteConcern.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._flushDatabaseCacheUpdatesWithWriteConcern.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._flushRoutingTableCacheUpdates.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._flushRoutingTableCacheUpdates.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._flushRoutingTableCacheUpdatesWithWriteConcern.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._flushRoutingTableCacheUpdatesWithWriteConcern.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._getNextSessionMods.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._getNextSessionMods.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._killOperations.failed` — bounded grouped-family vocabulary review required; roles: secondary
- `db.serverStatus().metrics.commands._killOperations.total` — bounded grouped-family vocabulary review required; roles: secondary
- `db.serverStatus().metrics.commands._migrateClone.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._migrateClone.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._recvChunkAbort.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._recvChunkAbort.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._recvChunkCommit.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._recvChunkCommit.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._recvChunkStart.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._recvChunkStart.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._recvChunkStatus.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._recvChunkStatus.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrCreateCollectionParticipant.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrCreateCollectionParticipant.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropCollection.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropCollection.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropCollectionParticipant.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropCollectionParticipant.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropDatabase.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropDatabase.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropDatabaseParticipant.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrDropDatabaseParticipant.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrSetAllowMigrations.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._shardsvrSetAllowMigrations.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._transferMods.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands._transferMods.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.aggregate.allowDiskUseTrue` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.authenticate.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.authenticate.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.autoSplitVector.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.autoSplitVector.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.buildInfo.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.buildInfo.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.collStats.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.collStats.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.connectionStatus.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.connectionStatus.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.count.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.count.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.createIndexes.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.createIndexes.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.currentOp.failed` — bounded grouped-family vocabulary review required; roles: secondary
- `db.serverStatus().metrics.commands.currentOp.total` — bounded grouped-family vocabulary review required; roles: secondary
- `db.serverStatus().metrics.commands.distinct.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.distinct.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.drop.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.drop.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.endSessions.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.endSessions.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.features.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.features.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.findAndModify.arrayFilters` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.findAndModify.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.findAndModify.pipeline` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.findAndModify.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.getCmdLineOpts.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.getCmdLineOpts.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.getLog.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.getLog.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.getParameter.failed` — bounded grouped-family vocabulary review required; roles: secondary
- `db.serverStatus().metrics.commands.getParameter.total` — bounded grouped-family vocabulary review required; roles: secondary
- `db.serverStatus().metrics.commands.hello.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.hello.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.isMaster.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.isMaster.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.killCursors.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.killCursors.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.listCollections.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.listCollections.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.listDatabases.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.listDatabases.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.listIndexes.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.listIndexes.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.ping.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.ping.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetGetConfig.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetGetConfig.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetGetStatus.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetGetStatus.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetHeartbeat.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetHeartbeat.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetReconfig.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.replSetReconfig.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.replSetRequestVotes.failed` — bounded grouped-family vocabulary review required; roles: secondary
- `db.serverStatus().metrics.commands.replSetRequestVotes.total` — bounded grouped-family vocabulary review required; roles: secondary
- `db.serverStatus().metrics.commands.replSetStepUp.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.replSetStepUp.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.replSetUpdatePosition.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.replSetUpdatePosition.total` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.setShardVersion.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.setShardVersion.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.splitChunk.failed` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.splitChunk.total` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.update.arrayFilters` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.update.pipeline` — bounded grouped-family vocabulary review required; roles: primary
- `db.serverStatus().metrics.commands.whatsmyuri.failed` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.commands.whatsmyuri.total` — bounded grouped-family vocabulary review required; roles: primary, secondary

### metrics.operatorCounters (200)

- `db.serverStatus().metrics.operatorCounters.expressions.$_internalJsEmit` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$_internalKeyStringValue` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$abs` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$acos` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$acosh` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$allElementsTrue` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$and` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$anyElementTrue` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$arrayElemAt` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$arrayToObject` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$asin` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$asinh` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$atan` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$atan2` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$atanh` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$avg` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$binarySize` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$bsonSize` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$ceil` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$cmp` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$concat` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$concatArrays` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$cond` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$const` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$convert` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$cos` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$cosh` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateAdd` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateDiff` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateFromParts` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateFromString` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateSubtract` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateToParts` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateToString` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dateTrunc` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dayOfMonth` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dayOfWeek` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$dayOfYear` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$degreesToRadians` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$divide` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$exp` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$filter` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$first` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$floor` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$function` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$getField` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$gte` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$hour` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$indexOfArray` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$indexOfBytes` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$indexOfCP` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$isArray` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$isNumber` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$isoDayOfWeek` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$isoWeek` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$isoWeekYear` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$last` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$let` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$literal` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$ln` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$log` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$log10` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$lt` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$lte` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$ltrim` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$map` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$max` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$mergeObjects` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$meta` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$millisecond` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$min` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$minute` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$mod` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$month` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$ne` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$not` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$objectToArray` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$or` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$pow` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$radiansToDegrees` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$rand` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$range` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$regexFind` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$regexFindAll` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$regexMatch` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$replaceAll` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$replaceOne` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$reverseArray` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$round` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$rtrim` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$second` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setDifference` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setEquals` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setField` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setIntersection` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setIsSubset` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$setUnion` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$sin` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$sinh` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$slice` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$split` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$sqrt` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$stdDevPop` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$stdDevSamp` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$strLenBytes` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$strLenCP` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$strcasecmp` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$substr` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$substrBytes` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$substrCP` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$sum` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$switch` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$tan` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$tanh` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toBool` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toDate` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toDecimal` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toDouble` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toHashedIndexKey` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toInt` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toLong` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toLower` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toObjectId` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toString` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$toUpper` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$trim` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$trunc` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$type` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$unsetField` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$week` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$year` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.expressions.$zip` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$_internalJsReduce` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$accumulator` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$addToSet` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$avg` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$count` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$first` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$last` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$max` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$mergeObjects` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$min` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$push` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$stdDevPop` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$stdDevSamp` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.groupAccumulators.$sum` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$all` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$alwaysFalse` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$alwaysTrue` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$and` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$bitsAllClear` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$bitsAllSet` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$bitsAnyClear` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$bitsAnySet` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$comment` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$elemMatch` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$exists` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$expr` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$geoIntersects` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$geoWithin` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$gt` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$gte` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$in` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$jsonSchema` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$lt` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$lte` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$mod` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$ne` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$near` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$nearSphere` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$nin` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$nor` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$not` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$or` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$regex` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$sampleRate` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$size` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$text` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$type` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.match.$where` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$addToSet` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$avg` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$count` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$covariancePop` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$covarianceSamp` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$denseRank` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$derivative` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$documentNumber` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$expMovingAvg` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$first` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$integral` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$last` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$max` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$min` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$push` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$rank` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$shift` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$stdDevPop` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$stdDevSamp` — bounded grouped-family vocabulary review required; roles: primary, secondary
- `db.serverStatus().metrics.operatorCounters.windowAccumulators.$sum` — bounded grouped-family vocabulary review required; roles: primary, secondary

## Histogram manual review

### metrics.query (112)

- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[0].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[0].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[10].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[10].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[11].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[11].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[1].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[1].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[2].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[2].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[3].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[3].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[4].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[4].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[5].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[5].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[6].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[6].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[7].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[7].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[8].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[8].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[9].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicMicros[9].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[0].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[0].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[1].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[1].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[2].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[2].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[3].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[3].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[4].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[4].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[5].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicNumPlans[5].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[0].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[0].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[1].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[1].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[2].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[2].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[3].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[3].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[4].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[4].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[5].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[5].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[6].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[6].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[7].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[7].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[8].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[8].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[9].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.classicWorks[9].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[0].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[0].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[10].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[10].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[11].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[11].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[1].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[1].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[2].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[2].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[3].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[3].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[4].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[4].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[5].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[5].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[6].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[6].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[7].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[7].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[8].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[8].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[9].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeMicros[9].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[0].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[0].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[1].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[1].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[2].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[2].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[3].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[3].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[4].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[4].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[5].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumPlans[5].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[0].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[0].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[1].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[1].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[2].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[2].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[3].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[3].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[4].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[4].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[5].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[5].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[6].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[6].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[7].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[7].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[8].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[8].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[9].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().metrics.query.multiPlanner.histograms.sbeNumReads[9].lowerBound` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary

### shardingStatistics.resharding (12)

- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis['(-inf, 10)'].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis['[10, 100)'].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis['[100, 1000)'].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis['[1000, 10000)'].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis['[10000, inf)'].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.collClonerFillBatchForInsertLatencyMillis.totalCount` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis['(-inf, 10)'].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis['[10, 100)'].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis['[100, 1000)'].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis['[1000, 10000)'].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis['[10000, inf)'].count` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().shardingStatistics.resharding.oplogApplierApplyBatchLatencyMillis.totalCount` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary

### wiredTiger.cache (5)

- `db.serverStatus().wiredTiger.cache['eviction walk target pages histogram - 0-9']` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target pages histogram - 10-31']` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target pages histogram - 128 and higher']` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target pages histogram - 32-63']` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary
- `db.serverStatus().wiredTiger.cache['eviction walk target pages histogram - 64-128']` — histogram representation approved as fixed numeric-boundary counters; roles: primary, secondary


### wiredTiger.perf (22)

- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 1) - 10-49ms']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 2) - 50-99ms']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 3) - 100-249ms']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 4) - 250-499ms']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 5) - 500-999ms']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system read latency histogram (bucket 6) - 1000ms+']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 1) - 10-49ms']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 2) - 50-99ms']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 3) - 100-249ms']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 4) - 250-499ms']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 5) - 500-999ms']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['file system write latency histogram (bucket 6) - 1000ms+']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation read latency histogram (bucket 1) - 100-249us']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation read latency histogram (bucket 2) - 250-499us']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation read latency histogram (bucket 3) - 500-999us']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation read latency histogram (bucket 4) - 1000-9999us']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation read latency histogram (bucket 5) - 10000us+']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation write latency histogram (bucket 1) - 100-249us']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation write latency histogram (bucket 2) - 250-499us']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation write latency histogram (bucket 3) - 500-999us']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation write latency histogram (bucket 4) - 1000-9999us']` — histogram representation deferred; roles: primary, secondary
- `db.serverStatus().wiredTiger.perf['operation write latency histogram (bucket 5) - 10000us+']` — histogram representation deferred; roles: primary, secondary
