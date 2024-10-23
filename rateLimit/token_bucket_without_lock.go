package rateLimit

import (
	"sync/atomic"
	"time"
)

type TokenBucketWithoutLock struct {
	Capacity      int64
	Tokens        int64
	Rate          int64
	LastTokenTime int64
}

func NewTokenBucketWithoutLock(capacity, rate int64) *TokenBucketWithoutLock {
	return &TokenBucketWithoutLock{
		Capacity:      capacity,
		Tokens:        capacity,
		Rate:          rate,
		LastTokenTime: time.Now().UnixNano(),
	}
}
func (tb *TokenBucketWithoutLock) Consume(tokens int64) bool {
	for {
		tb.refill()
		tokensBefore := atomic.LoadInt64(&tb.Tokens)
		if tokensBefore >= tokens {
			if atomic.CompareAndSwapInt64(&tb.Tokens, tokensBefore, tokensBefore-tokens) {
				return true
			}
		} else {
			return false
		}
	}
}
func (tb *TokenBucketWithoutLock) refill() {
	now := time.Now().UnixNano()
	lastTokenTime := atomic.LoadInt64(&tb.LastTokenTime)
	elapsed := now - lastTokenTime
	newTokens := elapsed * tb.Rate / 1e9
	if newTokens > 0 {
		tokensBefore := atomic.LoadInt64(&tb.Tokens)
		tokensAfter := tb.min(tokensBefore+newTokens, tb.Capacity)
		if atomic.CompareAndSwapInt64(&tb.Tokens, tokensBefore, tokensAfter) {
			atomic.StoreInt64(&tb.LastTokenTime, lastTokenTime)
		}
	}
}
func (tb *TokenBucketWithoutLock) min(a, b int64) int64 {
	if a >= b {
		return b
	}
	return a
}
