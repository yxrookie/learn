package trafficlimit

import (
	"fmt"
	"testing"
)

func TestTokenbucket(t *testing.T) {
	tb := NewTokenBucket(10, 1)
	for i := 0; i < 15; i++ {
		if tb.Consume(i) {
			fmt.Println("第", i, "轮", "服务正常调用")
		}
	}
}