package index_service

import (
	"framework/frame/util_types"
)

// Indexer Sentinel（分布式grpc的哨兵）和 LocalIndexer（单机索引）都实现了该接口
type Indexer interface {
	AddDoc(doc util_types.Document) (int, error)
	DeleteDoc(docId string) int
	Search(query *util_types.TermQuery, onFlag uint64, offFlag uint64, orFlags []uint64) []*util_types.Document
	Count() int
	Close() error
}
