package searcher

import (
	"framework/demo/types"
	index_service "framework/frame/server/worker/index"
	"sync"
)

// StringSearcher 业务代码 策略模式中的搜索器 与框架无关
type StringSearcher struct {
	ReCallers []ReCaller
	Filters   []Filter
}

func NewStringSearcherByMod1() *StringSearcher {
	// 使用模式1
	searcher := &StringSearcher{}
	searcher.ReCallers = []ReCaller{&RecallMod1{}}
	searcher.Filters = []Filter{&FilterMod1{}}
	return searcher
}

func (s *StringSearcher) WithReCallers(reCallers ...ReCaller) {
	s.ReCallers = append(s.ReCallers, reCallers...)
}

func (s *StringSearcher) WithFilters(filters ...Filter) {
	s.Filters = append(s.Filters, filters...)
}

func (s *StringSearcher) ReCalls(req *types.HelloSearchReq, index index_service.Indexer) []*types.HelloSearchRsp {

	// 1.收集结果
	collection := make(chan *types.HelloSearchRsp, 100)
	var wg sync.WaitGroup
	wg.Add(len(s.ReCallers))
	for _, reCaller := range s.ReCallers {
		go func(reCaller ReCaller) {
			wg.Add(1)
			defer wg.Done()
			searchRsps := reCaller.ReCall(req, index)
			for _, searchRsp := range searchRsps {
				collection <- searchRsp
			}
		}(reCaller)
	}

	// 启动另一个协程收集结果
	// 手机完成的关闭信号
	signalChan := make(chan struct{})
	// 用于合并手机的结果
	// 用于合并多路召回的视频结果
	helloSearchRsps := make([]*types.HelloSearchRsp, 0, 100)
	go func() {

		for {
			result, ok := <-collection
			if !ok {
				// 收集完成就退出循环
				break
			}
			// 有数据则收集
			helloSearchRsps = append(helloSearchRsps, result)
		}
		// 发送任务通知收集完成
		signalChan <- struct{}{}

	}()
	// 等待所有任务完成
	wg.Wait()
	// 结果发送完后就可以直接关闭结果集通道，关闭了不影响接收，避免内存泄露
	close(collection)
	// 等待结果集接收合并完成
	<-signalChan
	// 收集完成，返回数据
	return helloSearchRsps
}

func (s *StringSearcher) Applies(rsp []*types.HelloSearchRsp, index index_service.Indexer) []*types.HelloSearchRsp {
	for _, filter := range s.Filters {
		filter.Apply(rsp, index)
	}
	return nil
}

// Search 业务搜索
func (s *StringSearcher) Search(req *types.HelloSearchReq, index index_service.Indexer) []*types.HelloSearchRsp {
	return s.Applies(s.ReCalls(req, index), index)
}
