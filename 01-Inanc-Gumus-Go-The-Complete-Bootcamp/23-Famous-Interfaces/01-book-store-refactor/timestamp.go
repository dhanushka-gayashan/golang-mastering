package main

import (
	"strconv"
	"time"
)

// timestamp stores, formats and automatically prints a timestamp.
type timestamp struct {
	// timestamp anonymously embeds a time.
	// no need to convert a time value to a timestamp value to use the methods of the time type.
	time.Time
}

func (ts timestamp) string() string {
	if ts.IsZero() { // same as: ts.Time.IsZero()
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"

	// same as: ts.Time.Format(layout)
	return ts.Format(layout)
}

// toTimestamp returns a timestamp value depending on the type of `v`.
func toTimestamp(v interface{}) (ts timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
