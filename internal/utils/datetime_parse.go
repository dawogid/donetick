package utils

import "time"

// ParseFlexibleDateTime attempts multiple layouts:
// 1. RFC3339
// 2. DateTimeLayout (DD/MM/YYYY 24h)
// 3. DateLayout (DD/MM/YYYY) -> midnight UTC
func ParseFlexibleDateTime(raw string) (time.Time, error) {
	if raw == "" {
		return time.Time{}, ErrEmptyDate
	}
	if t, err := time.Parse(time.RFC3339, raw); err == nil {
		return t, nil
	}
	if t, err := time.Parse(DateTimeLayout, raw); err == nil {
		return t, nil
	}
	if d, err := time.Parse(DateLayout, raw); err == nil {
		return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC), nil
	}
	return time.Time{}, ErrInvalidDateFormat
}

var (
	ErrEmptyDate        = &DateParseError{Message: "empty date string"}
	ErrInvalidDateFormat = &DateParseError{Message: "invalid date format (expected RFC3339 or DD/MM/YYYY [with optional time])"}
)

type DateParseError struct{ Message string }

func (e *DateParseError) Error() string { return e.Message }
