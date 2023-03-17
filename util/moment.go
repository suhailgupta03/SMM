package util

import "time"

// IsDateGreater checks if the first date is greater than the second
// date.
func IsDateGreater(firstDate time.Time, secondDate time.Time) bool {
	if secondDate.Before(firstDate) {
		return true
	}
	return false
}
