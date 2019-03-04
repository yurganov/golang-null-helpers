package null

import (
	"database/sql"
)

type Bool struct {
	sql.NullBool
}

func NewBool(b bool, v bool) Bool {
	return Bool{
		NullBool: sql.NullBool{
			Bool:  b,
			Valid: v,
		},
	}
}

func BoolFrom(b bool) Bool {
	var boo Bool
	boo.SetFrom(b)

	return boo
}

func (b Bool) Get() bool {
	return b.Bool
}

func (b *Bool) Set(x bool, v bool) {
	b.Bool, b.Valid = x, v
}

func (b *Bool) SetFrom(x bool) {
	b.Bool = x
	b.Valid = !b.Empty()
}

func (b Bool) Isset() bool {
	return b.Valid
}

func (b Bool) Empty() bool {
	return !b.Bool
}
