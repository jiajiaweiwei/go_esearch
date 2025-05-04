package inverted

import (
	"hash/fnv"
	"sync"
)

// 手动实现支持并发读写的map 不是简单的加锁变成串行的锁

// 思想：对map进行分片的思想，并加读写锁
type ConcurrentHashMap struct {
	mps   []map[string]any // 全局map的数据由多个小map组成
	seg   int              // 小map的个数
	locks []sync.RWMutex   // 每个小map配一把读写锁，避免全局锁的竞争
	seed  uint32           // 每次执行farmHash的随机种子
}

// NewConcurrentHashMap 获取一个实例 cap预估容纳多少个元素 seg标识内部几个小map
func NewConcurrentHashMap(seg, cap int) *ConcurrentHashMap {
	// 每个小map都要初始化
	mps := make([]map[string]any, seg)

	for i := 0; i < seg; i++ {
		mps[i] = make(map[string]any, cap/seg)
	}
	return &ConcurrentHashMap{
		mps:   mps,
		seg:   seg,
		locks: make([]sync.RWMutex, seg),
		seed:  0,
	}
}

// getIndex 获取键的分区
func (c *ConcurrentHashMap) getIndex(key string) int {
	// 创建一个新的 FNV-32a 哈希对象
	h := fnv.New32a()
	// 将键写入哈希对象
	_, err := h.Write([]byte(key))
	if err != nil {
		panic(err)
	}
	// 计算哈希值
	hashValue := h.Sum32()
	// 将哈希值映射到分区索引上
	return int(hashValue) % c.seg
}

// 插入
func (c *ConcurrentHashMap) Store(key string, value any) {
	// 1.取哈希，取模确认分区
	index := c.getIndex(key)
	// 2.加写锁存数据
	c.locks[index].Lock()
	defer c.locks[index].Unlock()
	c.mps[index][key] = value
}

// 删除
func (c *ConcurrentHashMap) Delete(key string) {
	// 1.取哈希，取模确认分区
	index := c.getIndex(key)
	// 2.加写锁存数据
	c.locks[index].Lock()
	defer c.locks[index].Unlock()
	delete(c.mps[index], key)
}

// 读取
func (c *ConcurrentHashMap) Load(key string) (value any, ok bool) {
	// 1.取哈希，取模确认分区
	index := c.getIndex(key)
	// 2.加读锁查询
	c.locks[index].RLock()
	defer c.locks[index].RUnlock()
	value, ok = c.mps[index][key]
	return
}

// Rows  迭代器模式，要有next方法
type Rows interface {
	Next() bool
}
