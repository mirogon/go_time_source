package time_source

import "time"

type TimeSourceImpl struct {
}

func (timeSource TimeSourceImpl) Now() time.Time {
	return time.Now()
}

func (timeSource TimeSourceImpl) NowUtcRFC3339() string {
	now := time.Now().UTC()
	return now.Format(time.RFC3339)
}

func (timeSource TimeSourceImpl) FromUtcRFC3339(input string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, input)
	return t, err
}
