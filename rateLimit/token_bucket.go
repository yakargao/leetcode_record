package rateLimit

import (
	"sync"
	"time"
)

// TokenBucket 令牌桶算法
type TokenBucket struct {
	Capacity      int64
	Tokens        int64
	Rate          int64
	LastTokenTime int64
	mutex         sync.Mutex
}

func NewTokenBucket(capacity, rate int64) *TokenBucket {
	return &TokenBucket{
		Capacity:      capacity,
		Tokens:        capacity,
		Rate:          rate,
		LastTokenTime: time.Now().UnixNano(),
	}
}

func (tb *TokenBucket) Consume(tokens int64) bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	tb.refill()
	if tb.Tokens >= tokens {
		tb.Tokens -= tokens
		return true
	}
	return false
}
func (tb *TokenBucket) refill() {
	now := time.Now().UnixNano()
	elapsed := now - tb.LastTokenTime
	newTokens := elapsed * tb.Rate / 1e9
	if newTokens > 0 {
		tb.Tokens = tb.min(newTokens+tb.Tokens, tb.Capacity)
		tb.LastTokenTime = now
	}
}
func (tb *TokenBucket) min(a, b int64) int64 {
	if a >= b {
		return b
	}
	return a
}
