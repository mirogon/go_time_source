package time_source

import "time"

type TimeSource interface {
	Now() time.Time
	NowUtcRFC3339() string
	FromUtcRFC3339(string) (time.Time, error)
}

var TIME TimeSource = TimeSourceImpl{}
