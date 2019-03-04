package null

import (
	"database/sql"
)

type String struct {
	sql.NullString
}

func NewString(s string, v bool) String {
	return String{
		NullString: sql.NullString{
			String: s,
			Valid:  v,
		},
	}
}

func StringFrom(s string) String {
	var str String
	str.SetFrom(s)

	return str
}

func (s String) Get() string {
	return s.String
}

func (s *String) Set(x string, v bool) {
	s.String, s.Valid = x, v
}

func (s *String) SetFrom(x string) {
	s.String = x
	s.Valid = !s.Empty()
}

func (s String) Isset() bool {
	return s.Valid
}

func (s String) Empty() bool {
	return len(s.String) == 0
}
