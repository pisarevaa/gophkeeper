package model

import (
	"time"
)

type DateTime time.Time

func (d DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format("2006-01-02T15:04:05Z07:00")
	jsonStr := "\"" + formatted + "\""
	return []byte(jsonStr), nil
}
