package time_source

import "time"

type DebugTimeSource struct {
}

func (timeSource DebugTimeSource) Now() time.Time {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
	return t
}
func (timeSource DebugTimeSource) NowUtcRFC3339() string {
	now := time.Now().UTC()
	return now.Format(time.RFC3339)
}

func (timeSource DebugTimeSource) FromUtcRFC3339(input string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, input)
	return t, err
}
