package searcher

import (
	"framework/demo/proto/prototypes"
	"framework/demo/types"
	indexservice "framework/frame/server/worker/index"
	"framework/frame/util_types"
	"github.com/gogo/protobuf/proto"
)

type ReCaller interface {
	ReCall(*types.HelloSearchReq, indexservice.Indexer) []*types.HelloSearchRsp
}

type Filter interface {
	Apply([]*types.HelloSearchRsp, indexservice.Indexer) []*types.HelloSearchRsp
}

type RecallMod1 struct {
}

type FilterMod1 struct {
}

func (s *RecallMod1) ReCall(req *types.HelloSearchReq, index indexservice.Indexer) []*types.HelloSearchRsp {
	// 这里调用框架内部的方法
	// grpc远程调用
	query := &util_types.TermQuery{}
	// grpc调用
	documents := index.Search(query, 1, 1, []uint64{1})
	// 序列化grpc的结果，并返回
	// 创建一个用于存储匹配视频的切片
	helloSearchRsp := make([]*types.HelloSearchRsp, 0, len(documents))
	// 遍历匹配的文档，反序列化为视频对象，加入到视频列表中
	for _, doc := range documents {
		var rsp prototypes.HelloSearchRsp
		if err := proto.Unmarshal(doc.Bytes, &rsp); err == nil {
			var temp types.HelloSearchRsp
			temp.Response = rsp.Response
			helloSearchRsp = append(helloSearchRsp, &temp)
		}
	}
	return helloSearchRsp
}

func (s *FilterMod1) Apply(rsp []*types.HelloSearchRsp, index indexservice.Indexer) []*types.HelloSearchRsp {
	// 这里调用框架内部的方法
	// todo 省略过滤
	return rsp

}
