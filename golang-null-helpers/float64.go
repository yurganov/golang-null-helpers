package null

import (
	"database/sql"
)

type Float64 struct {
	sql.NullFloat64
}

func NewFloat64(f float64, v bool) Float64 {
	return Float64{
		NullFloat64: sql.NullFloat64{
			Float64: f, Valid: v,
		},
	}
}

func Float64From(f float64) Float64 {
	var f64 Float64
	f64.SetFrom(f)

	return f64
}

func (f Float64) Get() float64 {
	return f.Float64
}

func (f *Float64) Set(x float64, v bool) {
	f.Float64, f.Valid = x, v
}

func (f *Float64) SetFrom(x float64) {
	f.Float64 = x
	f.Valid = !f.Empty()
}

func (f Float64) Isset() bool {
	return f.Valid
}

func (f Float64) Empty() bool {
	return f.Float64 == 0
}
