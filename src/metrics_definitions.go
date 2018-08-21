package main

type serverStatus struct {
	Host                      *string                         `bson:"host"`
	Version                   *string                         `bson:"version"`
	Process                   *string                         `bson:"process"`
	PID                       *int                            `bson:"pid"`
	Uptime                    *int                            `bson:"uptime"`
	Asserts                   *serverStatusAsserts            `bson:"asserts"`
	Connections               *serverStatusConnections        `bson:"connections"`
	LogicalSessionRecordCache *serverStatusLSRC               `bson:"logicalSessionRecordCache"`
	Network                   *serverStatusNetwork            `bson:"network"`
	Opcounters                *serverStatusOpcounters         `bson:"opcounters"`
	OpcountersRepl            *serverStatusOpcountersRepl     `bson:"opcountersRepl"`
	Mem                       *serverStatusMem                `bson:"mem"`
	Metrics                   *serverStatusMetrics            `bson:"metrics"`
	BackgroundFlushing        *serverStatusBackgroundFlushing `bson:"backgroundFlushing"`
	GlobalLock                *serverStatusGlobalLock         `bson:"globalLock"`
	ExtraInfo                 *serverStatusExtraInfo          `bson:"extra_info"`
	WiredTiger                *serverStatusWiredTiger         `bson:"wiredTiger"`
	Locks                     *serverStatusLocks              `bson:"locks"`
	Dur                       *serverStatusDur                `bson:"dur"`
}

type serverStatusAsserts struct {
	Regular   *int `bson:"regular"   metric_name:"asserts.regularPerSecond"   source_type:"rate"`
	Warning   *int `bson:"warning"   metric_name:"asserts.warningPerSecond"   source_type:"rate"`
	Msg       *int `bson:"msg"       metric_name:"asserts.messagesPerSecond"  source_type:"rate"`
	User      *int `bson:"user"      metric_name:"asserts.userPerSecond"      source_type:"rate"`
	Rollovers *int `bson:"rollovers" metric_name:"asserts.rolloversPerSecond" source_type:"rate"`
}

type serverStatusConnections struct {
	Current      *int `bson:"current"      metric_name:"connections.current"      source_type:"gauge"`
	Available    *int `bson:"available"    metric_name:"connections.available"    source_type:"gauge"`
	TotalCreated *int `bson:"totalCreated" metric_name:"connections.totalCreated" source_type:"gauge"`
}

type serverStatusLSRC struct {
	ActiveSessionsCount                       *int `bson:"activeSessionsCount"`
	SessionsCollectionJobCount                *int `bson:"sessionsCollectionJobCount"`
	LastSessionsCollectionJobDurationMillis   *int `bson:"lastSessionsCollectionJobDurationMillis"`
	LastSessionsCollectionJobEntriesRefreshed *int `bson:"lastSessionsCollectionJobEntriesRefreshed"`
	LastSessionsCollectionJobEntriesClosed    *int `bson:"lastSessionCollectionJobEntriesClosed"`
	TransactionReaperJobCount                 *int `bson:"transactionReaperJobCount"`
	LastTransactionReaperJobDurationMillis    *int `bson:"lastTransactionReaperJobDurationMillis"`
	LastTransactionReaperJobEntriesCleanedUp  *int `bson:"lastTransactionReaperJobEntriesCleanedUp"`
}

type serverStatusNetwork struct {
	BytesIn     *int `bson:"bytesIn"     metric_name:"network.bytesIn"  source_type:"gauge"`
	BytesOut    *int `bson:"bytesOut"    metric_name:"network.bytesOut" source_type:"gauge"`
	NumRequests *int `bson:"numRequests" metric_name:"network.requests" source_type:"gauge"`
}

type serverStatusOpcounters struct {
	Insert  *int `bson:"insert"  metric_name:"opcounters.insertPerSecond"  source_type:"rate"`
	Query   *int `bson:"query"   metric_name:"opcounters.queryPerSecond"   source_type:"rate"`
	Update  *int `bson:"update"  metric_name:"opcounters.updatePerSecond"  source_type:"rate"`
	Delete  *int `bson:"delete"  metric_name:"opcounters.deletePerSecond"  source_type:"rate"`
	Getmore *int `bson:"getmore" metric_name:"opcounters.getmorePerSecond" source_type:"rate"`
	Command *int `bson:"command" metric_name:"opcounters.commandPerSecond" source_type:"rate"`
}

type serverStatusOpcountersRepl struct {
	Insert  *int `bson:"insert"  metric_name:"opcountersrepl.insertPerSecond"  source_type:"rate"`
	Query   *int `bson:"query"   metric_name:"opcountersrepl.queryPerSecond"   source_type:"rate"`
	Update  *int `bson:"update"  metric_name:"opcountersrepl.updatePerSecond"  source_type:"rate"`
	Delete  *int `bson:"delete"  metric_name:"opcountersrepl.deletePerSecond"  source_type:"rate"`
	Getmore *int `bson:"getmore" metric_name:"opcountersrepl.getmorePerSecond" source_type:"rate"`
	Command *int `bson:"command" metric_name:"opcountersrepl.commandPerSecond" source_type:"rate"`
}

type serverStatusMem struct {
	Bits              *int `bson:"bits"              metric_name:"mem.bits"                     source_type:"gauge"`
	Resident          *int `bson:"resident"          metric_name:"mem.residentInBytes"          source_type:"gauge"`
	Virtual           *int `bson:"virtual"           metric_name:"mem.virtualInBytes"           source_type:"gauge"`
	Mapped            *int `bson:"mapped"            metric_name:"mem.mappedInBytes"            source_type:"gauge"`
	MappedWithJournal *int `bson:"mappedWithJournal" metric_name:"mem.mappedWithJournalInBytes" source_type:"gauge"`
}

type serverStatusMetrics struct {
	Cursor        *serverStatusMetricsCursor        `bson:"cursor"`
	Commands      *serverStatusMetricsCommands      `bson:"commands"`
	GetLastError  *serverStatusMetricsGetLastError  `bson:"getLastError"`
	Operation     *serverStatusMetricsOperation     `bson:"operation"`
	QueryExecutor *serverStatusMetricsQueryExecutor `bson:"queryExecutor"`
	Record        *serverStatusMetricsRecord        `bson:"record"`
	Repl          *serverStatusMetricsRepl          `bson:"repl"`
	TTL           *serverStatusMetricsTTL           `bson:"ttl"`
}

type serverStatusMetricsCursor struct {
	TimedOut *int                           `bson:"timedOut" metric_name:"cursor.timedOutPerSecond" source_type:"rate"`
	Open     *serverStatusMetricsCursorOpen `bson:"open"`
}

type serverStatusMetricsCursorOpen struct {
	Total  *int `bson:"total"  metric_name:"cursor.openTotal"  source_type:"gauge"`
	Pinned *int `bson:"pinned" metric_name:"cursor.openPinned" source_type:"gauge"`
}

type serverStatusMetricsCommands struct {
	Count         *serverStatusMetricsCommandCount         `bson:"count"`
	CreateIndexes *serverStatusMetricsCommandCreateIndexes `bson:"createIndexes"`
	Delete        *serverStatusMetricsCommandDelete        `bson:"delete"`
	Eval          *serverStatusMetricsCommandEval          `bson:"eval"`
	FindAndModify *serverStatusMetricsCommandFindAndModify `bson:"findAndModify"`
	Insert        *serverStatusMetricsCommandInsert        `bson:"insert"`
	Update        *serverStatusMetricsCommandUpdate        `bson:"update"`
}

type serverStatusMetricsCommandCount struct {
	Failed *int `bson:"failed" metric_name:"commands.countFailedPerSecond" source_type:"rate"`
	Total  *int `bson:"total"  metric_name:"commands.countFailedPerSecond" source_type:"rate"`
}

type serverStatusMetricsCommandCreateIndexes struct {
	Failed *int `bson:"failed" metric_name:"commands.createIndexesFailedPerSecond" source_type:"rate"`
	Total  *int `bson:"total"  metric_name:"commands.createIndexesFailedPerSecond" source_type:"rate"`
}

type serverStatusMetricsCommandDelete struct {
	Failed *int `bson:"failed" metric_name:"commands.deleteFailedPerSecond" source_type:"rate"`
	Total  *int `bson:"total"  metric_name:"commands.deleteFailedPerSecond" source_type:"rate"`
}

type serverStatusMetricsCommandEval struct {
	Failed *int `bson:"failed" metric_name:"commands.evalFailedPerSecond" source_type:"rate"`
	Total  *int `bson:"total"  metric_name:"commands.evalFailedPerSecond" source_type:"rate"`
}

type serverStatusMetricsCommandFindAndModify struct {
	Failed *int `bson:"failed" metric_name:"commands.modifyFailedPerSecond" source_type:"rate"`
	Total  *int `bson:"total"  metric_name:"commands.modifyFailedPerSecond" source_type:"rate"`
}

type serverStatusMetricsCommandInsert struct {
	Failed *int `bson:"failed" metric_name:"commands.insertFailedPerSecond" source_type:"rate"`
	Total  *int `bson:"total"  metric_name:"commands.insertFailedPerSecond" source_type:"rate"`
}

type serverStatusMetricsCommandUpdate struct {
	Failed *int `bson:"failed" metric_name:"commands.updateFailedPerSecond" source_type:"rate"`
	Total  *int `bson:"total"  metric_name:"commands.updateFailedPerSecond" source_type:"rate"`
}

type serverStatusMetricsGetLastError struct {
	Wtime     *serverStatusMetricsGetLastErrorWtime `bson:"wtime"`
	Wtimeouts *int                                  `bson:"wtimeouts" metric_name:"getlasterror.wtimeoutsPerSecond" source_type:"rate"`
}

type serverStatusMetricsGetLastErrorWtime struct {
	TotalMillis *int `bson:"totalMillis" metric_name:"getlasterror.wtimeMillisPerSecond" source_type:"rate"`
}

type serverStatusMetricsOperation struct {
	ScanAndOrder   *int `bson:"scanAndOrder"   metric_name:"operation.scanAndOrderPerSecond"   source_type:"rate"`
	WriteConflicts *int `bson:"writeConflicts" metric_name:"operation.writeConflictsPerSecond" source_type:"rate"`
}

type serverStatusMetricsQueryExecutor struct {
	Scanned *int `bson:"scanned" metric_name:"queryexecutor.scannedPerSecond" source_type:"rate"`
}

type serverStatusMetricsRecord struct {
	Moves *int `bson:"moves" metric_name:"record.movesPerSecond" source_type:"rate"`
}

type serverStatusMetricsRepl struct {
	Apply   *serverStatusMetricsReplApply   `bson:"apply"`
	Buffer  *serverStatusMetricsReplBuffer  `bson:"buffer"`
	Network *serverStatusMetricsReplNetwork `bson:"network"`
	Preload *serverStatusMetricsReplPreload `bson:"preload"`
}

type serverStatusMetricsReplApply struct {
	Ops     *int                                 `bson:"ops" metric_name:"repl.apply.operationsPerSecond" source_type:"rate"`
	Batches *serverStatusMetricsReplApplyBatches `bson:"batches"`
}

type serverStatusMetricsReplApplyBatches struct {
	Num *int `bson:"num" metric_name:"repl.apply.batchesPerSec" source_type:"rate"`
}

type serverStatusMetricsReplBuffer struct {
	Count        *int `bson:"count"        metric_name:"repl.buffer.count"          source_type:"gauge"`
	MaxSizeBytes *int `bson:"maxSizeBytes" metric_name:"repl.buffer.maxSizeInBytes" source_type:"gauge"`
	SizeBytes    *int `bson:"sizeBytes"    metric_name:"repl.buffer.sizeInBytes"    source_type:"gauge"`
}

type serverStatusMetricsReplNetwork struct {
	Bytes          *int                                    `bson:"bytes"          metric_name:"repl.network.bytesPerSecond"          source_type:"rate"`
	Ops            *int                                    `bson:"ops"            metric_name:"repl.network.operationPerSecond"      source_type:"rate"`
	ReadersCreated *int                                    `bson:"readersCreated" metric_name:"repl.network.readersCreatedPerSecond" source_type:"rate"`
	Getmores       *serverStatusMetricsReplNetworkGetmores `bson:"getmores"`
}

type serverStatusMetricsReplNetworkGetmores struct {
	Num *int `bson:"num" metric_name:"repl.network.getmoresPerSecond" source_type:"rate"`
}

type serverStatusMetricsReplPreload struct {
	Docs    *serverStatusMetricsReplPreloadDocs    `bson:"docs"`
	Indexes *serverStatusMetricsReplPreloadIndexes `bson:"indexes"`
}

type serverStatusMetricsReplPreloadDocs struct {
	Num         *int `bson:"num"         metric_name:"repl.docsLoadedPrefetch"        source_type:"gauge"`
	TotalMillis *int `bson:"totalMillis" metric_name:"repl.docsPreloadInMilliseconds" source_type:"gauge"`
}

type serverStatusMetricsReplPreloadIndexes struct {
	Num         *int `bson:"num"         metric_name:"repl.IndexLoadedPrefetch"        source_type:"gauge"`
	TotalMillis *int `bson:"totalMillis" metric_name:"repl.indexPreloadInMilliseconds" source_type:"gauge"`
}

type serverStatusMetricsTTL struct {
	DeletedDocuments *int `bson:"deletedDocuments" metric_name:"ttl.deleteDocumentsPerSecond" source_type:"rate"`
	Passes           *int `bson:"passes"           metric_name:"ttl.removeDocumentPerSecond"  source_type:"rate"`
}

type serverStatusBackgroundFlushing struct {
	Flushes   *int     `bson:"flushes"    metric_name:"flush.flushesDisk"           source_type:"gauge"`
	TotalMs   *float64 `bson:"total_ms"   metric_name:"flush.totalInMillisends"     source_type:"gauge"`
	AverageMs *float64 `bson:"average_ms" metric_name:"flush.averageInMilliseconds" source_type:"gauge"`
	LastMs    *float64 `bson:"last_ms"    metric_name:"flush.lastInMilliseconds"    source_type:"gauge"`
}

type serverStatusGlobalLock struct {
	TotalTime     *float32                             `bson:"totalTime" metric_name:"globallock.totaltime" source_type:"gauge"`
	CurrentQueue  *serverStatusGlobalLockCurrentQueue  `bson:"currentQueue"`
	ActiveClients *serverStatusGlobalLockActiveClients `bson:"activeClients"`
}

type serverStatusGlobalLockCurrentQueue struct {
	Total   *int `bson:"totalTime" metric_name:"globallock.currentQueueTotal"   source_type:"gauge"`
	Readers *int `bson:"readers"   metric_name:"globallock.currentQueueReaders" source_type:"gauge"`
	Writers *int `bson:"writers"   metric_name:"globallock.currentQueueWriters" source_type:"gauge"`
}

type serverStatusGlobalLockActiveClients struct {
	Total   *int `bson:"totalTime" metric_name:"globallock.activeClientsTotal"   source_type:"gauge"`
	Readers *int `bson:"readers"   metric_name:"globallock.activeClientsReaders" source_type:"gauge"`
	Writers *int `bson:"writers"   metric_name:"globallock.activeClientsWriters" source_type:"gauge"`
}

type serverStatusExtraInfo struct {
	PageFaults *int `bson:"page_faults" metric_name:"pageFaultsPerSecond" source_type:"rate"`
}

type serverStatusWiredTiger struct {
	Cache                  *serverStatusWiredTigerCache                  `bson:"cache"`
	ConcurrentTransactions *serverStatusWiredTigerConcurrentTransactions `bson:"concurrentTransactions"`
}

type serverStatusWiredTigerCache struct {
	Size                   *int `bson:"bytes currently in the cache"                                 metric_name:"wiredtiger.cacheInBytes"                                 source_type:"gauge"`
	FailedEvictions        *int `bson:"failed eviction of pages that exceeded the in-memory maximum" metric_name:"wiredtiger.failedEvictionsPagesPerSecond"                source_type:"gauge"`
	PageSplits             *int `bson:"in-memory page splits"                                        metric_name:"cacheInMemoryPageSplits"                                 source_type:"gauge"`
	MaxSize                *int `bson:"maximum bytes configured"                                     metric_name:"wiredtiger.cacheMaxInBytes"                              source_type:"gauge"`
	MaxPageSize            *int `bson:"maximum page size at eviction"                                metric_name:"wiredtiger.cacheMaxPageSizeEvictionInBytes"              source_type:"gauge"`
	ModifiedPagesEvicted   *int `bson:"modified pages evicted"                                       metric_name:"wiredtiger.cacheModifiedPagesEvicted"                    source_type:"gauge"`
	PagesHeld              *int `bson:"pages currently held in the cache"                            metric_name:"wiredtiger.cachePagesHeld"                               source_type:"gauge"`
	PagesEvictedThreads    *int `bson:"pages evicted by application threads"                         metric_name:"wiredtiger.cachePagesEvictedApplicationThreadsPerSecond" source_type:"gauge"`
	PagesEvictedMax        *int `bson:"pages evicted because they exceeded the in-memory maximum"    metric_name:"wiredtiger.cachePagesEvictedInMemoryMaxPerSecond"        source_type:"rate"`
	DirtyData              *int `bson:"tracked dirty bytes in the cache"                             metric_name:"wiredtiger.cacheDirtyDataInBytes"                        source_type:"gauge"`
	UnmodifiedPagesEvicted *int `bson:"unmodified pages evicted"                                     metric_name:"wiredtiger.cacheUnmodifiedPagesEvicted"                  source_type:"gauge"`
}

type serverStatusWiredTigerConcurrentTransactions struct {
	Write *serverStatusWiredTigerConcurrentTransactionsWrite `bson:"write"`
	Read  *serverStatusWiredTigerConcurrentTransactionsRead  `bson:"read"`
}

type serverStatusWiredTigerConcurrentTransactionsWrite struct {
	Out          *int `bson:"out"          metric_name:"wiredtiger.concurrentTransactions.WriteRemaining" source_type:"gauge"`
	Available    *int `bson:"available"    metric_name:"wiredtiger.concurrentTransactions.WriteAvailable" source_type:"gauge"`
	TotalTickets *int `bson:"totalTickets" metric_name:"wiredtiger.concurrentTransactions.WriteTotal"     source_type:"gauge"`
}

type serverStatusWiredTigerConcurrentTransactionsRead struct {
	Out          *int `bson:"out"          metric_name:"wiredtiger.concurrentTransactions.ReadRemaining" source_type:"gauge"`
	Available    *int `bson:"available"    metric_name:"wiredtiger.concurrentTransactions.ReadAvailable" source_type:"gauge"`
	TotalTickets *int `bson:"totalTickets" metric_name:"wiredtiger.concurrentTransactions.ReadTotal"     source_type:"gauge"`
}

type serverStatusLocks struct {
	Collection    *serverStatusLocksCollection    `bson:"Collection"`
	Database      *serverStatusLocksDatabase      `bson:"Database"`
	Global        *serverStatusLocksGlobal        `bson:"Global"`
	Metadata      *serverStatusLocksMetadata      `bson:"Metadata"`
	MMAPV1Journal *serverStatusLocksMMAPV1Journal `bson:"MMAPV1Journal"`
	Oplog         *serverStatusLocksOplog         `bson:"oplog"`
}

type serverStatusLocksCollection struct {
	AcquireCount        *serverStatusLocksCollectionAcquireCount        `bson:"acquireCount"`
	AcquireWaitCount    *serverStatusLocksCollectionAcquireWaitCount    `bson:"acquireWaitCount"`
	TimeAcquiringMicros *serverStatusLocksCollectionTimeAcquiringMicros `bson:"timeAcquiringMicros"`
}

type serverStatusLocksCollectionAcquireCount struct {
	Shared          *int `bson:"R" metric_name:"locks.collectionAcquiredShared" source_type:"gauge"`
	Exclusive       *int `bson:"W" metric_name:"locks.collectionAcquiredShared" source_type:"gauge"`
	IntentShared    *int `bson:"r" metric_name:"locks.collectionAcquiredShared" source_type:"gauge"`
	IntentExclusive *int `bson:"w" metric_name:"locks.collectionAcquiredShared" source_type:"gauge"`
}

type serverStatusLocksCollectionAcquireWaitCount struct {
	Shared    *int `bson:"R" metric_name:"locks.collectionAcquireWaitCountShared"    source_type:"gauge"`
	Exclusive *int `bson:"W" metric_name:"locks.collectionAcquireWaitCountExclusive" source_type:"gauge"`
}

type serverStatusLocksCollectionTimeAcquiringMicros struct {
	Shared    *int `bson:"R" metric_name:"locks.collectionTimeAcquiringMicrosShared"    source_type:"gauge"`
	Exclusive *int `bson:"W" metric_name:"locks.collectionTimeAcquiringMicrosExclusive" source_type:"gauge"`
}

type serverStatusLocksDatabase struct {
	AcquireCount        *serverStatusLocksDatabaseAcquireCount        `bson:"acquireCount"`
	AcquireWaitCount    *serverStatusLocksDatabaseAcquireWaitCount    `bson:"acquireWaitCount"`
	TimeAcquiringMicros *serverStatusLocksDatabaseTimeAcquiringMicros `bson:"timeAcquiringMicros"`
}

type serverStatusLocksDatabaseAcquireCount struct {
	Shared          *int `bson:"R" metric_name:"locks.databaseAcquiredShared" source_type:"gauge"`
	Exclusive       *int `bson:"W" metric_name:"locks.databaseAcquiredShared" source_type:"gauge"`
	IntentShared    *int `bson:"r" metric_name:"locks.databaseAcquiredShared" source_type:"gauge"`
	IntentExclusive *int `bson:"w" metric_name:"locks.databaseAcquiredShared" source_type:"gauge"`
}

type serverStatusLocksDatabaseAcquireWaitCount struct {
	Shared          *int `bson:"R" metric_name:"locks.databaseAcquireWaitCountShared"          source_type:"gauge"`
	Exclusive       *int `bson:"W" metric_name:"locks.databaseAcquireWaitCountExclusive"       source_type:"gauge"`
	IntentShared    *int `bson:"r" metric_name:"locks.databaseAcquireWaitCountIntentShared"    source_type:"gauge"`
	IntentExclusive *int `bson:"w" metric_name:"locks.databaseAcquireWaitCountIntentExclusive" source_type:"gauge"`
}

type serverStatusLocksDatabaseTimeAcquiringMicros struct {
	Shared          *int `bson:"R" metric_name:"locks.databaseTimeAcquiringMicrosShared"          source_type:"gauge"`
	Exclusive       *int `bson:"W" metric_name:"locks.databaseTimeAcquiringMicrosExclusive"       source_type:"gauge"`
	IntentShared    *int `bson:"r" metric_name:"locks.databaseTimeAcquiringMicrosIntentShared"    source_type:"gauge"`
	IntentExclusive *int `bson:"w" metric_name:"locks.databaseTimeAcquiringMicrosIntentExclusive" source_type:"gauge"`
}

type serverStatusLocksGlobal struct {
	AcquireCount        *serverStatusLocksGlobalAcquireCount        `bson:"acquireCount"`
	AcquireWaitCount    *serverStatusLocksGlobalAcquireWaitCount    `bson:"acquireWaitCount"`
	TimeAcquiringMicros *serverStatusLocksGlobalTimeAcquiringMicros `bson:"timeAcquiringMicros"`
}

type serverStatusLocksGlobalAcquireCount struct {
	Shared          *int `bson:"R" metric_name:"locks.globalAcquiredShared" source_type:"gauge"`
	Exclusive       *int `bson:"W" metric_name:"locks.globalAcquiredShared" source_type:"gauge"`
	IntentShared    *int `bson:"r" metric_name:"locks.globalAcquiredShared" source_type:"gauge"`
	IntentExclusive *int `bson:"w" metric_name:"locks.globalAcquiredShared" source_type:"gauge"`
}

type serverStatusLocksGlobalAcquireWaitCount struct {
	Shared          *int `bson:"R" metric_name:"locks.globalAcquireWaitCountShared"          source_type:"gauge"`
	Exclusive       *int `bson:"W" metric_name:"locks.globalAcquireWaitCountExclusive"       source_type:"gauge"`
	IntentShared    *int `bson:"r" metric_name:"locks.globalAcquireWaitCountIntentShared"    source_type:"gauge"`
	IntentExclusive *int `bson:"w" metric_name:"locks.globalAcquireWaitCountIntentExclusive" source_type:"gauge"`
}

type serverStatusLocksGlobalTimeAcquiringMicros struct {
	Shared          *int `bson:"R" metric_name:"locks.globalTimeAcquiringMicrosShared"          source_type:"gauge"`
	Exclusive       *int `bson:"W" metric_name:"locks.globalTimeAcquiringMicrosExclusive"       source_type:"gauge"`
	IntentShared    *int `bson:"r" metric_name:"locks.globalTimeAcquiringMicrosIntentShared"    source_type:"gauge"`
	IntentExclusive *int `bson:"w" metric_name:"locks.globalTimeAcquiringMicrosIntentExclusive" source_type:"gauge"`
}

type serverStatusLocksMetadata struct {
	AcquireCount *serverStatusLocksMetadataAcquireCount `bson:"acquireCount"`
}

type serverStatusLocksMetadataAcquireCount struct {
	Shared    *int `bson:"R" metric_name:"locks.metadataAcquiredShared" source_type:"gauge"`
	Exclusive *int `bson:"W" metric_name:"locks.metadataAcquiredShared" source_type:"gauge"`
}

type serverStatusLocksMMAPV1Journal struct {
	AcquireCount        *serverStatusLocksMMAPV1JournalAcquireCount        `bson:"acquireCount"`
	TimeAcquiringMicros *serverStatusLocksMMAPV1JournalTimeAcquiringMicros `bson:"timeAcquiringMicros"`
}

type serverStatusLocksMMAPV1JournalAcquireCount struct {
	Shared          *int `bson:"R" metric_name:"locks.mmapv1journalAcquiredShared"          source_type:"gauge"`
	Exclusive       *int `bson:"W" metric_name:"locks.mmapv1journalAcquiredExclusive"       source_type:"gauge"`
	IntentShared    *int `bson:"r" metric_name:"locks.mmapv1journalAcquiredIntentShared"    source_type:"gauge"`
	IntentExclusive *int `bson:"w" metric_name:"locks.mmapv1journalAcquiredIntentExclusive" source_type:"gauge"`
}

type serverStatusLocksOplog struct {
	AcquireCount        *serverStatusLocksOplogAcquireCount        `bson:"acquireCount"`
	TimeAcquiringMicros *serverStatusLocksOplogTimeAcquiringMicros `bson:"timeAcquiringMicros"`
}

type serverStatusLocksOplogAcquireCount struct {
	Shared          *int `bson:"R" metric_name:"locks.oplogAcquiredShared"          source_type:"gauge"`
	Exclusive       *int `bson:"W" metric_name:"locks.oplogAcquiredExclusive"       source_type:"gauge"`
	IntentShared    *int `bson:"r" metric_name:"locks.oplogAcquiredIntentShared"    source_type:"gauge"`
	IntentExclusive *int `bson:"w" metric_name:"locks.oplogAcquiredIntentExclusive" source_type:"gauge"`
}

type serverStatusLocksOplogTimeAcquiringMicros struct {
	IntentShared    *int `bson:"r" metric_name:"locks.oplogTimeAcquiringMicrosIntentShared"    source_type:"gauge"`
	IntentExclusive *int `bson:"w" metric_name:"locks.oplogTimeAcquiringMicrosIntentExclusive" source_type:"gauge"`
}

type serverStatusLocksMMAPV1JournalTimeAcquiringMicros struct {
	Shared    *int `bson:"R" metric_name:"locks.mmapv1journalTimeAcquiringMicrosShared"          source_type:"gauge"`
	Exclusive *int `bson:"W" metric_name:"locks.mmapv1journalTimeAcquiringMicrosExclusive"       source_type:"gauge"`
}

type serverStatusDur struct {
	Commits               *int                   `bson:"commits"                                       metric_name:"dur.commits"            source_type:"gauge"`
	CommitsInWriteLock    *int                   `bson:"commitsInWriteLock"                            metric_name:"dur.commitsInWriteLock" source_type:"gauge"`
	Compression           *float32               `bson:"compression"                                   metric_name:"dur.compression"        source_type:"gauge"`
	EarlyCommits          *int                   `bson:"earlyCommits"                                  metric_name:"dur.earlyCommits"       source_type:"gauge"`
	JournaledBytes        *float32               `metric_name:"dur.journaledInBytes"                   source_type:"gauge"`
	WriteToDataFilesBytes *float32               `metric_name:"dur.dataWrittenJournalDataFilesInBytes" source_type:"gauge"`
	JournaledMB           *float32               `bson:"journaledMB"`
	WriteToDataFilesMB    *float32               `bson:"writeToDataFilesMB"`
	TimeMS                *serverStatusDurTimeMS `bson:"timeMs"`
}

type serverStatusDurTimeMS struct {
	Commits            *int `bson:"commits"            metric_name:"dur.commitsInMilliseconds"              source_type:"gauge"`
	CommitsInWriteLock *int `bson:"commitsInWriteLock" metric_name:"dur.commitsInWriteLockInMilliseconds"   source_type:"gauge"`
	Dt                 *int `bson:"dt"                 metric_name:"dur.timeCollectedCommitsInMilliseconds" source_type:"gauge"`
	PrepLogBuffer      *int `bson:"prepLogBuffer"      metric_name:"dur.preparingInMilliseconds"            source_type:"gauge"`
	RemapPrivateView   *int `bson:"remapPrivateView"   metric_name:"dur.remappingInMilliseconds"            source_type:"gauge"`
	WriteToDataFiles   *int `bson:"writeToDataFiles"   metric_name:"dur.writingDataFilesInMilliseconds"     source_type:"gauge"`
	WriteToJournal     *int `bson:"writeToJournal"     metric_name:"dur.writingJournalInMilliseconds"       source_type:"gauge"`
}

type collStats struct {
	Size        *int  `bson:"size"        metric_name:"collection.sizeInBytes"        source_type:"gauge"`
	AvgObjSize  *int  `bson:"avgObjSize"  metric_name:"collection.avgObjSizeInBytes"  source_type:"gauge"`
	Count       *int  `bson:"count"       metric_name:"collection.count"              source_type:"gauge"`
	Capped      *bool `bson:"capped"      metric_name:"collection.capped"             source_type:"attribute"`
	Max         *int  `bson:"max"         metric_name:"collection.max"                source_type:"gauge"`
	MaxSize     *int  `bson:"maxSize"     metric_name:"collection.maxSizeInBytes"     source_type:"gauge"`
	StorageSize *int  `bson:"storageSize" metric_name:"collection.storageSizeInBytes" source_type:"gauge"`
	Nindexes    *int  `bson:"nindexes"    metric_name:"collection.nindexes"           source_type:"gauge"`
}

type dbStats struct {
	Objects     *int `bson:"objects"     metric_name:"stats.objects"        source_type:"gauge"`
	StorageSize *int `bson:"storageSize" metric_name:"stats.storageInBytes" source_type:"gauge"`
	IndexSize   *int `bson:"indexSize"   metric_name:"stats.indexInBytes"   source_type:"gauge"`
	Indexes     *int `bson:"indexes"     metric_name:"stats.indexes"        source_type:"gauge"`
	DataSize    *int `bson:"dataSize"    metric_name:"stats.dataInBytes"    source_type:"gauge"`
}

type replSetGetStatus struct {
	Members []replSetGetStatusMember `bson:"members"`
}

type replSetGetStatusMember struct {
	Name     *string `bson:"name"     metric_name:"replset.name"                 source_type:"attribute"`
	Health   *int    `bson:"health"   metric_name:"replset.health"               source_type:"gauge"`
	StateStr *string `bson:"stateStr" metric_name:"replset.state"                source_type:"attribute"`
	Uptime   *int    `bson:"uptime"   metric_name:"replset.uptimeInMilliseconds" source_type:"gauge"`
}