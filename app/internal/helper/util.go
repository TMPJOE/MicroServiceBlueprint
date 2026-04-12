// Package helper provides common utility functions used across the application.
// This includes helper methods for data transformation, validation, formatting,
// and other reusable utility operations.
package helper

import "os"

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// Custom errors
