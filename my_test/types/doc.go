package types

func (kw Keyword) ToString() string {
	if len(kw.Word) > 0 {
		// 拼接分隔符，用于构建高效的符合键
		/*
			为什么选择 \001 作为分隔符？
			不可见字符：\001（ASCII 1）是控制字符，极少出现在正常文本中，避免与关键词内容冲突。
			有序性：在字节序中，\001 位于普通字符之前（如 \001 < A < a），确保复合键按字段字典序排列。
			高效解析：在反序列化时，可以用 strings.SplitN(s, "\001", 2) 快速还原为两个字段
		*/
		return kw.Field + "\001" + kw.Word
	} else {
		return ""
	}
}
