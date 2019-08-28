package series

import (
	"math/big"
	"time"
)

// GeneralTimeline represents a time-index series with interface value
type GeneralTimeline map[time.Time]interface{}

// IntTimeline represents a time-index series with int value
type IntTimeline map[time.Time]int

// Int16Timeline represents a time-index series with int16 value
type Int16Timeline map[time.Time]int16

// Int32Timeline represents a time-index series with int32 value
type Int32Timeline map[time.Time]int32

// Int64Timeline represents a time-index series with int64 value
type Int64Timeline map[time.Time]int64

// BigIntTimeline represents a time-index series with big.Int value
type BigIntTimeline map[time.Time]*big.Int

// Float32Timeline represents a time-index series with int64 value
type Float32Timeline map[time.Time]float32

// Float64Timeline represents a time-index series with int64 value
type Float64Timeline map[time.Time]float64

// BigFloatTimeline represents a time-index series with big.Float value
type BigFloatTimeline map[time.Time]*big.Float

// BigRatTimeline represents a time-index series with big.Rat value
type BigRatTimeline map[time.Time]*big.Rat

// Implement ISeries
func (s *GeneralTimeline) series()  {}
func (s *IntTimeline) series()      {}
func (s *Int16Timeline) series()    {}
func (s *Int32Timeline) series()    {}
func (s *Int64Timeline) series()    {}
func (s *BigIntTimeline) series()   {}
func (s *Float32Timeline) series()  {}
func (s *Float64Timeline) series()  {}
func (s *BigFloatTimeline) series() {}
func (s *BigRatTimeline) series()   {}

// Implement ICalculable
func (s *IntTimeline) calculable()      {}
func (s *Int16Timeline) calculable()    {}
func (s *Int32Timeline) calculable()    {}
func (s *Int64Timeline) calculable()    {}
func (s *BigIntTimeline) calculable()   {}
func (s *Float32Timeline) calculable()  {}
func (s *Float64Timeline) calculable()  {}
func (s *BigFloatTimeline) calculable() {}
func (s *BigRatTimeline) calculable()   {}
