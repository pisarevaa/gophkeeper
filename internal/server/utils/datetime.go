package utils

import (
	"time"
)

type Datetime time.Time

func (d Datetime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format("2006-01-02T15:04:05Z07:00")
	jsonStr := "\"" + formatted + "\""
	return []byte(jsonStr), nil
}
