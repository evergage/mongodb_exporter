# Dynamic vocabulary and histogram review

## Approved policy

- Use bounded grouped families with fixture-derived label vocabularies.
- Preserve all observed primary/secondary keys, including internal MongoDB keys.
- Unknown runtime keys emit no sample and increment `mongodb_server_status_unknown_keys_total{family="..."}` without exposing the unknown key.
- Histogram bucket counts use fixed grouped counters. Each bucket has one numeric lower-bound label in the source base unit.
- Query histogram `lowerBound` leaves provide labels for paired `count` leaves and do not create separate samples.
- Derive and review final series ceilings after all mappings are generated.

## Commands

Proposed family: `mongodb_server_status_commands_total{command,outcome}` for `total` and `failed`. Command-specific auxiliary fields use bounded grouped families with `command` and `feature` labels.

Observed command vocabulary (60):

- `_flushDatabaseCacheUpdates` — fields: `failed`, `total` — internal
- `_flushDatabaseCacheUpdatesWithWriteConcern` — fields: `failed`, `total` — internal
- `_flushRoutingTableCacheUpdates` — fields: `failed`, `total` — internal
- `_flushRoutingTableCacheUpdatesWithWriteConcern` — fields: `failed`, `total` — internal
- `_getNextSessionMods` — fields: `failed`, `total` — internal
- `_killOperations` — fields: `failed`, `total` — internal
- `_migrateClone` — fields: `failed`, `total` — internal
- `_recvChunkAbort` — fields: `failed`, `total` — internal
- `_recvChunkCommit` — fields: `failed`, `total` — internal
- `_recvChunkStart` — fields: `failed`, `total` — internal
- `_recvChunkStatus` — fields: `failed`, `total` — internal
- `_shardsvrCreateCollectionParticipant` — fields: `failed`, `total` — internal
- `_shardsvrDropCollection` — fields: `failed`, `total` — internal
- `_shardsvrDropCollectionParticipant` — fields: `failed`, `total` — internal
- `_shardsvrDropDatabase` — fields: `failed`, `total` — internal
- `_shardsvrDropDatabaseParticipant` — fields: `failed`, `total` — internal
- `_shardsvrSetAllowMigrations` — fields: `failed`, `total` — internal
- `_transferMods` — fields: `failed`, `total` — internal
- `aggregate` — fields: `allowDiskUseTrue`, `failed`, `total`
- `authenticate` — fields: `failed`, `total`
- `autoSplitVector` — fields: `failed`, `total`
- `buildInfo` — fields: `failed`, `total`
- `collStats` — fields: `failed`, `total`
- `connectionStatus` — fields: `failed`, `total`
- `count` — fields: `failed`, `total`
- `createIndexes` — fields: `failed`, `total`
- `currentOp` — fields: `failed`, `total`
- `delete` — fields: `failed`, `total`
- `distinct` — fields: `failed`, `total`
- `drop` — fields: `failed`, `total`
- `endSessions` — fields: `failed`, `total`
- `features` — fields: `failed`, `total`
- `find` — fields: `failed`, `total`
- `findAndModify` — fields: `arrayFilters`, `failed`, `pipeline`, `total`
- `getCmdLineOpts` — fields: `failed`, `total`
- `getLog` — fields: `failed`, `total`
- `getMore` — fields: `failed`, `total`
- `getParameter` — fields: `failed`, `total`
- `hello` — fields: `failed`, `total`
- `insert` — fields: `failed`, `total`
- `isMaster` — fields: `failed`, `total`
- `killCursors` — fields: `failed`, `total`
- `listCollections` — fields: `failed`, `total`
- `listDatabases` — fields: `failed`, `total`
- `listIndexes` — fields: `failed`, `total`
- `moveChunk` — fields: `failed`, `total`
- `ping` — fields: `failed`, `total`
- `replSetGetConfig` — fields: `failed`, `total`
- `replSetGetStatus` — fields: `failed`, `total`
- `replSetHeartbeat` — fields: `failed`, `total`
- `replSetReconfig` — fields: `failed`, `total`
- `replSetRequestVotes` — fields: `failed`, `total`
- `replSetStepUp` — fields: `failed`, `total`
- `replSetUpdatePosition` — fields: `failed`, `total`
- `rotateCertificates` — fields: `failed`, `total`
- `serverStatus` — fields: `failed`, `total`
- `setShardVersion` — fields: `failed`, `total`
- `splitChunk` — fields: `failed`, `total`
- `update` — fields: `arrayFilters`, `failed`, `pipeline`, `total`
- `whatsmyuri` — fields: `failed`, `total`

## Aggregation stages

Proposed family: `mongodb_server_status_aggregation_stage_total{stage}`.

Observed stage vocabulary (49):

- `$_internalApplyOplogUpdate` — internal
- `$_internalBoundedSort` — internal
- `$_internalConvertBucketIndexStats` — internal
- `$_internalFindAndModifyImageLookup` — internal
- `$_internalInhibitOptimization` — internal
- `$_internalReshardingIterateTransaction` — internal
- `$_internalReshardingOwnershipMatch` — internal
- `$_internalSetWindowFields` — internal
- `$_internalShredDocuments` — internal
- `$_internalSplitPipeline` — internal
- `$_internalUnpackBucket` — internal
- `$_unpackBucket` — internal
- `$addFields`
- `$bucket`
- `$bucketAuto`
- `$changeStream`
- `$collStats`
- `$count`
- `$currentOp`
- `$documents`
- `$facet`
- `$geoNear`
- `$graphLookup`
- `$group`
- `$indexStats`
- `$limit`
- `$listLocalSessions`
- `$listSessions`
- `$lookup`
- `$match`
- `$merge`
- `$mergeCursors`
- `$operationMetrics`
- `$out`
- `$planCacheStats`
- `$project`
- `$queue`
- `$redact`
- `$replaceRoot`
- `$replaceWith`
- `$sample`
- `$set`
- `$setWindowFields`
- `$skip`
- `$sort`
- `$sortByCount`
- `$unionWith`
- `$unset`
- `$unwind`

## Operator counters

Proposed family: `mongodb_server_status_operator_total{category,operator}`. Categories and operator values are independently allowlisted.

### expressions (141)

- `$_internalJsEmit` — internal
- `$_internalKeyStringValue` — internal
- `$abs`
- `$acos`
- `$acosh`
- `$add`
- `$allElementsTrue`
- `$and`
- `$anyElementTrue`
- `$arrayElemAt`
- `$arrayToObject`
- `$asin`
- `$asinh`
- `$atan`
- `$atan2`
- `$atanh`
- `$avg`
- `$binarySize`
- `$bsonSize`
- `$ceil`
- `$cmp`
- `$concat`
- `$concatArrays`
- `$cond`
- `$const`
- `$convert`
- `$cos`
- `$cosh`
- `$dateAdd`
- `$dateDiff`
- `$dateFromParts`
- `$dateFromString`
- `$dateSubtract`
- `$dateToParts`
- `$dateToString`
- `$dateTrunc`
- `$dayOfMonth`
- `$dayOfWeek`
- `$dayOfYear`
- `$degreesToRadians`
- `$divide`
- `$eq`
- `$exp`
- `$filter`
- `$first`
- `$floor`
- `$function`
- `$getField`
- `$gt`
- `$gte`
- `$hour`
- `$ifNull`
- `$in`
- `$indexOfArray`
- `$indexOfBytes`
- `$indexOfCP`
- `$isArray`
- `$isNumber`
- `$isoDayOfWeek`
- `$isoWeek`
- `$isoWeekYear`
- `$last`
- `$let`
- `$literal`
- `$ln`
- `$log`
- `$log10`
- `$lt`
- `$lte`
- `$ltrim`
- `$map`
- `$max`
- `$mergeObjects`
- `$meta`
- `$millisecond`
- `$min`
- `$minute`
- `$mod`
- `$month`
- `$multiply`
- `$ne`
- `$not`
- `$objectToArray`
- `$or`
- `$pow`
- `$radiansToDegrees`
- `$rand`
- `$range`
- `$reduce`
- `$regexFind`
- `$regexFindAll`
- `$regexMatch`
- `$replaceAll`
- `$replaceOne`
- `$reverseArray`
- `$round`
- `$rtrim`
- `$second`
- `$setDifference`
- `$setEquals`
- `$setField`
- `$setIntersection`
- `$setIsSubset`
- `$setUnion`
- `$sin`
- `$sinh`
- `$size`
- `$slice`
- `$split`
- `$sqrt`
- `$stdDevPop`
- `$stdDevSamp`
- `$strLenBytes`
- `$strLenCP`
- `$strcasecmp`
- `$substr`
- `$substrBytes`
- `$substrCP`
- `$subtract`
- `$sum`
- `$switch`
- `$tan`
- `$tanh`
- `$toBool`
- `$toDate`
- `$toDecimal`
- `$toDouble`
- `$toHashedIndexKey`
- `$toInt`
- `$toLong`
- `$toLower`
- `$toObjectId`
- `$toString`
- `$toUpper`
- `$trim`
- `$trunc`
- `$type`
- `$unsetField`
- `$week`
- `$year`
- `$zip`

### groupAccumulators (14)

- `$_internalJsReduce` — internal
- `$accumulator`
- `$addToSet`
- `$avg`
- `$count`
- `$first`
- `$last`
- `$max`
- `$mergeObjects`
- `$min`
- `$push`
- `$stdDevPop`
- `$stdDevSamp`
- `$sum`

### match (35)

- `$all`
- `$alwaysFalse`
- `$alwaysTrue`
- `$and`
- `$bitsAllClear`
- `$bitsAllSet`
- `$bitsAnyClear`
- `$bitsAnySet`
- `$comment`
- `$elemMatch`
- `$eq`
- `$exists`
- `$expr`
- `$geoIntersects`
- `$geoWithin`
- `$gt`
- `$gte`
- `$in`
- `$jsonSchema`
- `$lt`
- `$lte`
- `$mod`
- `$ne`
- `$near`
- `$nearSphere`
- `$nin`
- `$nor`
- `$not`
- `$or`
- `$regex`
- `$sampleRate`
- `$size`
- `$text`
- `$type`
- `$where`

### windowAccumulators (20)

- `$addToSet`
- `$avg`
- `$count`
- `$covariancePop`
- `$covarianceSamp`
- `$denseRank`
- `$derivative`
- `$documentNumber`
- `$expMovingAvg`
- `$first`
- `$integral`
- `$last`
- `$max`
- `$min`
- `$push`
- `$rank`
- `$shift`
- `$stdDevPop`
- `$stdDevSamp`
- `$sum`

## Query multi-planner histograms

Proposed family shapes:

- `mongodb_server_status_query_multi_planner_duration_bucket_total{engine,lower_bound_microseconds}`
- `mongodb_server_status_query_multi_planner_plan_count_bucket_total{engine,lower_bound_plans}`
- `mongodb_server_status_query_multi_planner_work_bucket_total{engine,lower_bound_works}`
- `mongodb_server_status_query_multi_planner_read_bucket_total{engine,lower_bound_reads}`

Observed lower bounds:

- `classicMicros`: 0, 1024, 4096, 16384, 65536, 262144, 1048576, 4194304, 16777216, 67108864, 268435456, 1073741824
- `classicNumPlans`: 0, 2, 4, 8, 16, 32
- `classicWorks`: 0, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768
- `sbeMicros`: 0, 1024, 4096, 16384, 65536, 262144, 1048576, 4194304, 16777216, 67108864, 268435456, 1073741824
- `sbeNumPlans`: 0, 2, 4, 8, 16, 32
- `sbeNumReads`: 0, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768

## Resharding latency ranges

Use raw milliseconds and a single `lower_bound_millis` label. Since latency is non-negative, the first `(-inf, 10)` source bucket is normalized to lower bound `0`. Approved lower bounds: `0`, `10`, `100`, `1000`, `10000`.

Proposed families:

- `mongodb_server_status_resharding_coll_cloner_fill_batch_latency_bucket_total{lower_bound_millis}`
- `mongodb_server_status_resharding_oplog_applier_apply_batch_latency_bucket_total{lower_bound_millis}`

## WiredTiger histogram ranges

Use one numeric lower-bound label in the unit implied by each WiredTiger statistic. Normalize textual ranges such as `0-9`, `10-31`, and `128 and higher` to lower bounds `0`, `10`, and `128`. Preserve raw cumulative counts.

The generator must reject a histogram description whose numeric lower bound cannot be extracted unambiguously.
