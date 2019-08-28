package series

import "math/big"

// GeneralMapSeries represents a map with string key and interface value.
type GeneralMapSeries map[string]interface{}

// BigIntMapSeries represents a map with string key and big.Int value
type BigIntMapSeries map[string]*big.Int

// BigFloatMapSeries represents a map with string key and big.Float value.
type BigFloatMapSeries map[string]*big.Float

// BigRatMapSeries represents a map with string key and big.Rat value.
type BigRatMapSeries map[string]*big.Rat

// Implement ISeries
func (s *GeneralMapSeries) series()  {}
func (s *BigIntMapSeries) series()   {}
func (s *BigFloatMapSeries) series() {}
func (s *BigRatMapSeries) series()   {}

// Implement ICalculable
func (s *BigIntMapSeries) calculable()   {}
func (s *BigFloatMapSeries) calculable() {}
func (s *BigRatMapSeries) calculable()   {}
