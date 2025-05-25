package types

type HelloSearchReq struct {
	QueryCond   string `json:"queryCond"`
	ResponseCtx string `json:"response"`
}

type HelloSearchRsp struct {
	Response string `json:"response"`
}
