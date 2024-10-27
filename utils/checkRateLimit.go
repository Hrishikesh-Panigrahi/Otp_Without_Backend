package utils

import (
	"sync"
	"time"
)

type RateLimiter struct {
	Attempts  int
	ResetTime time.Time
}

var (
	rateLimits     = make(map[string]*RateLimiter)
	rateLimitMutex = &sync.Mutex{}
	maxAttempts    = 5
	resetInterval  = 1 * time.Minute
)

// CheckRateLimit checks if the client IP has exceeded the allowed number of attempts.
// It returns true if the request is allowed, and false if the limit is exceeded.
func CheckRateLimit(clientIP string) bool {
	rateLimitMutex.Lock()
	defer rateLimitMutex.Unlock()

	limiter, exists := rateLimits[clientIP]

	// If no limiter exists for this IP or the reset time has passed, reset the counter
	if !exists || time.Now().After(limiter.ResetTime) {
		rateLimits[clientIP] = &RateLimiter{
			Attempts:  1,
			ResetTime: time.Now().Add(resetInterval),
		}
		return true
	}

	// If max attempts are reached, deny the request
	if limiter.Attempts >= maxAttempts {
		return false
	}

	// Increment attempt count and allow the request
	limiter.Attempts++
	return true
}

// CleanRateLimits periodically removes expired entries to free up memory.
func CleanRateLimits() {
	rateLimitMutex.Lock()
	defer rateLimitMutex.Unlock()

	for ip, limiter := range rateLimits {
		if time.Now().After(limiter.ResetTime) {
			delete(rateLimits, ip)
		}
	}
}
