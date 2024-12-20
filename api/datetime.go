package api

import (
	"encoding/json"
	"strings"
	"time"
)

var _ json.Marshaler = (*DateTime)(nil)
var _ json.Unmarshaler = (*DateTime)(nil)

type DateTime time.Time

var TimeFormats = []string{
	"2006-01-02T15:04:05.999",
	"2006-01-02T15:04:05.999Z",
}

// Implement Marshaler and Unmarshaler interface
func (d *DateTime) UnmarshalJSON(b []byte) error {
	var (
		t   time.Time
		err error
	)
	s := strings.Trim(string(b), "\"")
	for _, tf := range TimeFormats {
		t, err = time.Parse(tf, s)
		if err == nil {
			*d = DateTime(t)
			return nil
		}

	}
	return err
}

func (j DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.AsTime())
}

func (d DateTime) String() string {
	return d.AsTime().Format("2006-01-02T15:04:05.999")
}

func (j DateTime) AsTime() time.Time {
	return time.Time(j)
}

var _ json.Marshaler = (*Date)(nil)
var _ json.Unmarshaler = (*Date)(nil)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.AsTime().Format(time.DateOnly))
}

func (d Date) String() string {
	return d.AsTime().Format(time.DateOnly)
}

func (d Date) AsTime() time.Time {
	return time.Time(d)
}

func NewDate(year, month, day int) Date {
	return Date(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}
