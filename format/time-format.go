package format

import "time"

const (
	TimeFormat = "2006-01-02 15:04:05"
	DayFormat  = "2006-01-02"
)

var (
	beijing = time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))
)
