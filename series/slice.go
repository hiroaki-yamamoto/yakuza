package series

import "math/big"

// GeneralSeries represents a simple alias of the array of interface
type GeneralSeries []interface{}

// IntSeries represents a simple alias of the array of int
type IntSeries []int

// Int16Series represents a simple alias of the array of int16
type Int16Series []int16

// Int32Series represents a simple alias of the array of int32
type Int32Series []int16

// Int64Series represents a simple alias of the array of int64
type Int64Series []int16

// BigIntSeries represents a multiple precision integer series
type BigIntSeries []big.Int

// Float32Series represents a simple alias of the array of float32
type Float32Series []int16

// Float64Series represents a simple alias of the array of float64
type Float64Series []int16

// BigFloatSeries represents a multiple precision float series
type BigFloatSeries []big.Float

// BigRatSeries represents a multiple precision quotient series
type BigRatSeries []big.Float

// All Series must implements series to follow ISeries
func (s GeneralSeries) series()  {}
func (s IntSeries) series()      {}
func (s Int16Series) series()    {}
func (s Int32Series) series()    {}
func (s Int64Series) series()    {}
func (s BigIntSeries) series()   {}
func (s Float32Series) series()  {}
func (s Float64Series) series()  {}
func (s BigFloatSeries) series() {}
func (s BigRatSeries) series()   {}

// Integer related series must implement calculable to follow ICalculable.
func (s IntSeries) calculable()      {}
func (s Int16Series) calculable()    {}
func (s Int32Series) calculable()    {}
func (s Int64Series) calculable()    {}
func (s BigIntSeries) calculable()   {}
func (s Float32Series) calculable()  {}
func (s Float64Series) calculable()  {}
func (s BigFloatSeries) calculable() {}
func (s BigRatSeries) calculable()   {}
