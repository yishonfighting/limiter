package limiter

// Limiter 限流器接口
// 在高并发业务场景下，保护系统时，常用的"三板斧"有："熔断、降级和限流"。
// 限流算法常用的几种实现方式有如下四种：
// 1. 计数器
// 2. 滑动窗口
// 3. 漏桶
// 4. 令牌桶
// 本接口定义了限流器的基本行为，即判断是否达到限流阀值。
type Limiter interface {
	IsLimit() bool
}
