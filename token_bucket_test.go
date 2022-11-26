package ratelimiter

import (
	"fmt"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tokenBucket := NewTokenBucket(2, 10)
	generateNumbers(tokenBucket)
}

func generateNumbers(tb *TokenBucket) {
	n := 1
	for n < 50 {
		allowed := tb.IsRequestAllowed()
		if !allowed {
			time.Sleep(1 * time.Second)
			fmt.Println("Not Allowed, waiting for a second")
			continue
		}
		fmt.Println(n)
		n += 1
	}
}
