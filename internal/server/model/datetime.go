package model

import (
	"strings"
	"time"
)

type DateTime time.Time

const ctLayout = "2006-01-02 15:04:05 Z07:00"

func (d DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format(ctLayout)
	jsonStr := "\"" + formatted + "\""
	return []byte(jsonStr), nil
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	newTime, err := time.Parse(ctLayout, s)
	if err != nil {
		return err
	}
	*d = DateTime(newTime)
	return nil
}
