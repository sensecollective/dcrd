// Package dbnamespace contains constants that define the database namespaces
// for the purpose of the blockchain, so that external callers may easily access
// this data.
package dbnamespace

import (
	"encoding/binary"
)

var (
	// ByteOrder is the preferred byte order used for serializing numeric
	// fields for storage in the database.
	ByteOrder = binary.LittleEndian

	// BCDBInfoBucketName is the name of the database bucket used to house
	// global versioning and date information for the blockchain database.
	BCDBInfoBucketName = []byte("dbinfo")

	// BCDBInfoVersionKeyName is the name of the database key used to house
	// the database version.  It is itself under the BCDBInfoBucketName
	// bucket.
	BCDBInfoVersionKeyName = []byte("version")

	// BCDBInfoCompressionVersionKeyName is the name of the database key
	// used to house the database compression version.  It is itself under
	// the BCDBInfoBucketName bucket.
	BCDBInfoCompressionVersionKeyName = []byte("compver")

	// BCDBInfoCreatedKeyName is the name of the database key used to house
	// date the database was created.  It is itself under the
	// BCDBInfoBucketName bucket.
	BCDBInfoCreatedKeyName = []byte("created")

	// HashIndexBucketName is the name of the db bucket used to house to the
	// block hash -> block height index.
	HashIndexBucketName = []byte("hashidx")

	// HeightIndexBucketName is the name of the db bucket used to house to
	// the block height -> block hash index.
	HeightIndexBucketName = []byte("heightidx")

	// ChainStateKeyName is the name of the db key used to store the best
	// chain state.
	ChainStateKeyName = []byte("chainstate")

	// SpendJournalBucketName is the name of the db bucket used to house
	// transactions outputs that are spent in each block.
	SpendJournalBucketName = []byte("spendjournal")

	// UtxoSetBucketName is the name of the db bucket used to house the
	// unspent transaction output set.
	UtxoSetBucketName = []byte("utxoset")
)
