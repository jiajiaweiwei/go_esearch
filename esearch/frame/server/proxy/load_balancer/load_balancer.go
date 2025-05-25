package load_balancer

// LoadBalancer 负载均衡接口，定义选择Endpoint的方法
type LoadBalancer interface {
	// Take 从给定的端点列表中选择一个
	Take(endpoints []string) string
}

// 该接口一共有两个实现，也就是有两个负载均衡算法 一个轮询 一个随机
