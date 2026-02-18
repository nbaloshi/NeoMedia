package utils

import "time"

func HasExpired(expiresAt time.Time) bool {
	return time.Now().After(expiresAt)
}

func IdleTooLong(lastAccessed time.Time, duration time.Duration) bool {
	return time.Since(lastAccessed) > duration
}

func RefreshLastAccessed() time.Time {
	return time.Now()
}