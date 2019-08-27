package series

import "time"

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

// Float32Series represents a simple alias of the array of float32
type Float32Series []int16

// Float64Series represents a simple alias of the array of float64
type Float64Series []int16

// GeneralMapSeries represents a simple map with string key and interface value.
type GeneralMapSeries map[string]interface{}

// IntMapSeries represents a simple map with string key and int value.
type IntMapSeries map[string]int

// Int16MapSeries represents a simple map with string key and int16 value.
type Int16MapSeries map[string]int16

// Int32MapSeries represents a simple map with string key and int32 value.
type Int32MapSeries map[string]int32

// Int64MapSeries represents a simple map with string key and int64 value.
type Int64MapSeries map[string]int64

// Float32MapSeries represents a simple map with string key and float32 value.
type Float32MapSeries map[string]float32

// Float64MapSeries represents a simple map with string key and float64 value.
type Float64MapSeries map[string]float64

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

// Float32Timeline represents a time-index series with int64 value
type Float32Timeline map[time.Time]float32

// Float64Timeline represents a time-index series with int64 value
type Float64Timeline map[time.Time]float64
