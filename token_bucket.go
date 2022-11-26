package ratelimiter

import (
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	rate          int64
	capacity      int64
	currentTokens int64
	lastRefillTs  time.Time
	mutex         sync.Mutex
}

func NewTokenBucket(rate int64, capacity int64) *TokenBucket {
	return &TokenBucket{
		rate:          rate,
		capacity:      capacity,
		currentTokens: capacity,
		lastRefillTs:  time.Now(),
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	end := time.Since(tb.lastRefillTs).Seconds()
	tokensTobeAdded := end * float64(tb.rate)
	tb.currentTokens = int64(math.Min(float64(tb.currentTokens)+tokensTobeAdded, float64(tb.capacity)))
	tb.lastRefillTs = now
}

func (tb *TokenBucket) IsRequestAllowed() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	tb.refill()
	if tb.currentTokens > 0 {
		tb.currentTokens -= 1
		return true
	}
	return false
}
