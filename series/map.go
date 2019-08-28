package series

import "math/big"

// GeneralMapSeries represents a map with string key and interface value.
type GeneralMapSeries map[string]interface{}

// IntMapSeries represents a map with string key and int value.
type IntMapSeries map[string]int

// Int16MapSeries represents a map with string key and int16 value.
type Int16MapSeries map[string]int16

// Int32MapSeries represents a map with string key and int32 value.
type Int32MapSeries map[string]int32

// Int64MapSeries represents a map with string key and int64 value.
type Int64MapSeries map[string]int64

// BigIntMapSeries represents a map with string key and big.Int value
type BigIntMapSeries map[string]*big.Int

// Float32MapSeries represents a map with string key and float32 value.
type Float32MapSeries map[string]float32

// Float64MapSeries represents a map with string key and float64 value.
type Float64MapSeries map[string]float64

// BigFloatMapSeries represents a map with string key and big.Float value.
type BigFloatMapSeries map[string]*big.Float

// BigRatMapSeries represents a map with string key and big.Rat value.
type BigRatMapSeries map[string]*big.Rat

// Implement ISeries
func (s *GeneralMapSeries) series()  {}
func (s *IntMapSeries) series()      {}
func (s *Int16MapSeries) series()    {}
func (s *Int32MapSeries) series()    {}
func (s *Int64MapSeries) series()    {}
func (s *BigIntMapSeries) series()   {}
func (s *Float32MapSeries) series()  {}
func (s *Float64MapSeries) series()  {}
func (s *BigFloatMapSeries) series() {}
func (s *BigRatMapSeries) series()   {}

// Implement ICalculable
func (s *IntMapSeries) calculable()      {}
func (s *Int16MapSeries) calculable()    {}
func (s *Int32MapSeries) calculable()    {}
func (s *Int64MapSeries) calculable()    {}
func (s *BigIntMapSeries) calculable()   {}
func (s *Float32MapSeries) calculable()  {}
func (s *Float64MapSeries) calculable()  {}
func (s *BigFloatMapSeries) calculable() {}
func (s *BigRatMapSeries) calculable()   {}
