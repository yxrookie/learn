package trafficlimit

import (
	"log"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity    int        // 桶容量
	tokens      int        // 当前令牌数
	refillRate  float64    // 令牌填充速率
	refillInterval time.Duration //令牌填充间隔
	lastRefill  time.Time  // 上次令牌填充时间
	refillMutex sync.Mutex // 令牌填充互斥锁
}

// 初始化令牌桶
func NewTokenBucket(capacity int, refillRate float64) *TokenBucket {
	t := &TokenBucket{
		capacity: capacity,
		tokens: capacity,
		refillRate: refillRate,
		refillInterval: time.Second,  //默认每秒填充一次
		lastRefill: time.Now(),
	}
	go t.startRefillLoop() // 启动令牌补充循环
	return t
}

// 定时填充令牌
func (tb *TokenBucket) startRefillLoop() {
	for range time.Tick(tb.refillInterval) {
		tb.refill()
	}
}

// 发放令牌
func (t TokenBucket) refill () {
	t.refillMutex.Lock()
	defer t.refillMutex.Unlock()
	now := time.Now()
	newBucket := now.Sub(t.lastRefill).Seconds() * t.refillRate
	t.tokens = min(t.capacity, t.tokens+int(newBucket))
	t.lastRefill = time.Now()
}

// 处理请求进行限流
func (t TokenBucket) Consume(token int) bool {
	t.refillMutex.Lock()
	defer t.refillMutex.Unlock()

	if token > t.tokens {
		log.Println("服务已限流...")
		return false
	}
	t.tokens -= token
	return true
}


func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}