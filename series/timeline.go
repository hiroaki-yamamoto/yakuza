package series

import (
	"math/big"
	"time"
)

// GeneralTimeline represents a time-index series with interface value
type GeneralTimeline map[time.Time]interface{}

// BigIntTimeline represents a time-index series with big.Int value
type BigIntTimeline map[time.Time]*big.Int

// BigFloatTimeline represents a time-index series with big.Float value
type BigFloatTimeline map[time.Time]*big.Float

// BigRatTimeline represents a time-index series with big.Rat value
type BigRatTimeline map[time.Time]*big.Rat

// Implement ISeries
func (s GeneralTimeline) series()  {}
func (s BigIntTimeline) series()   {}
func (s BigFloatTimeline) series() {}
func (s BigRatTimeline) series()   {}

// Implement ICalculable
func (s BigIntTimeline) calculable()   {}
func (s BigFloatTimeline) calculable() {}
func (s BigRatTimeline) calculable()   {}
