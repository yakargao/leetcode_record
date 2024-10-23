package rateLimit

import "testing"

// 基准测试有锁版本
func BenchmarkTokenBucketWithLock(b *testing.B) {
	tb := NewTokenBucket(100, 50) // 容量为100，每秒产生50个令牌

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tb.Consume(1)
	}
}

// 基准测试无锁版本
func BenchmarkTokenBucketWithoutLock(b *testing.B) {
	tb := NewTokenBucketWithoutLock(100, 50) // 容量为100，每秒产生50个令牌

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tb.Consume(1)
	}
}
