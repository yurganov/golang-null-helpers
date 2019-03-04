package null

import (
	"database/sql"
)

type Int64 struct {
	sql.NullInt64
}

func NewInt64(i int64, v bool) Int64 {
	return Int64{
		NullInt64: sql.NullInt64{
			Int64: i,
			Valid: v,
		},
	}
}

func Int64From(i int64) Int64 {
	var i64 Int64
	i64.SetFrom(i)

	return i64
}

func (i Int64) Get() int64 {
	return i.Int64
}

func (i *Int64) Set(x int64, v bool) {
	i.Int64, i.Valid = x, v
}

func (i *Int64) SetFrom(x int64) {
	i.Int64 = x
	i.Valid = !i.Empty()
}

func (i Int64) Isset() bool {
	return i.Valid
}

func (i Int64) Empty() bool {
	return i.Int64 == 0
}
