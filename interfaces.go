package yakuza

// ICalculable represents the provided structure is calculable.
type ICalculable interface {
	calculable()
}

// ISeries represents the provided structure is Seroes.
type ISeries interface {
	series()
}

// ICalculableSeries represents the provided structure is calculable, and Series.
type ICalculableSeries interface {
	calculable()
	series()
}
