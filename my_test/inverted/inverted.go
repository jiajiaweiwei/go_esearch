package inverted

// Doc 文档分词
type Doc struct {
	Id       int
	Keywords []string // 提取关键词后的分词
}

// BuildInvertIndex 构建倒排索引
func BuildInvertIndex(docs []*Doc) map[string][]int {
	// 返回值就是构建好的倒排索引
	index := make(map[string][]int, 100)
	for _, doc := range docs {
		for _, keyword := range doc.Keywords {
			index[keyword] = append(index[keyword], doc.Id)
		}
	}
	return index
}

// 倒排索引优化，加候选字段 使用比特位
