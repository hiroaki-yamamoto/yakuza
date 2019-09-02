package series

import "math/big"

// GeneralSeries represents a simple alias of the array of interface
type GeneralSeries []interface{}

// BigIntSeries represents a multiple precision integer series
type BigIntSeries []*big.Int

// BigFloatSeries represents a multiple precision float series
type BigFloatSeries []*big.Float

// BigRatSeries represents a multiple precision quotient series
type BigRatSeries []*big.Rat

// All Series must implements series to follow ISeries
func (s GeneralSeries) series()  {}
func (s BigIntSeries) series()   {}
func (s BigFloatSeries) series() {}
func (s BigRatSeries) series()   {}

// Integer related series must implement calculable to follow ICalculable.
func (s BigIntSeries) calculable()   {}
func (s BigFloatSeries) calculable() {}
func (s BigRatSeries) calculable()   {}
