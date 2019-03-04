package null

import (
	"database/sql/driver"
	"time"
)

type Time struct {
	Time  time.Time
	Valid bool
}

func NewTime(t time.Time, v bool) Time {
	return Time{t, v}
}

func TimeFrom(t time.Time) Time {
	var tm Time
	tm.SetFrom(t)
	return tm
}

func (t Time) Get() time.Time {
	return t.Time
}

func (t *Time) Set(x time.Time, v bool) {
	t.Time, t.Valid = x, v
}

func (t *Time) SetFrom(x time.Time) {
	t.Time = x
	t.Valid = !t.Empty()
}

func (t Time) Isset() bool {
	return t.Valid
}

func (t Time) Empty() bool {
	return t.Time.Equal(time.Time{})
}

func (t *Time) Scan(value interface{}) error {
	t.Time, t.Valid = value.(time.Time)
	return nil
}

func (t Time) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}
