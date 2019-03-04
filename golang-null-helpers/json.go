package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Json struct {
	Blob  []byte
	Valid bool
}

var _ sql.Scanner = (*Json)(nil)
var _ driver.Valuer = (*Json)(nil)

func NewJson(b []byte, v bool) Json {
	return Json{
		Blob:  b,
		Valid: v,
	}
}

func JsonFrom(b []byte) Json {
	var js Json
	js.SetFrom(b)

	return js
}

func (j Json) Get() []byte {
	return j.Blob
}

func (j *Json) Set(x []byte, v bool) {
	j.Blob, j.Valid = x, v
}

func (j *Json) SetFrom(x []byte) {
	j.Blob = x
	j.Valid = !j.Empty()
}

func (j Json) Isset() bool {
	return j.Valid
}

func (j Json) Empty() bool {
	return len(j.Blob) == 0 || string(j.Blob) == "null"
}

func (j Json) Value() (driver.Value, error) {
	if !j.Valid {
		return nil, nil
	}

	return j.Blob, nil
}

func (j *Json) Scan(value interface{}) error {
	if value == nil {
		j.Blob, j.Valid = []byte("null"), false

		return nil
	}

	switch x := value.(type) {
	case []byte:
		j.Blob = x
	case string:
		j.Blob = []byte(x)
	default:
		return errors.New("Incompatible type for null.NullJSON")
	}

	j.Valid = true

	return nil
}

func (j *Json) MarshalJSON(v interface{}) *Json {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic("can't marshaled json for database")
	}
	j.SetFrom(bytes)

	return j
}

func (j Json) UnmarshalJSON(v interface{}) {
	if err := json.Unmarshal(j.Blob, v); err != nil {
		panic("can't unmarshaled json from database")
	}
}
