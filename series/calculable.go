package series

func (s IntSeries) calculable()     {}
func (s Int16Series) calculable()   {}
func (s Int32Series) calculable()   {}
func (s Int64Series) calculable()   {}
func (s Float32Series) calculable() {}
func (s Float64Series) calculable() {}

func (s IntMapSeries) calculable()     {}
func (s Int16MapSeries) calculable()   {}
func (s Int32MapSeries) calculable()   {}
func (s Int64MapSeries) calculable()   {}
func (s Float32MapSeries) calculable() {}
func (s Float64MapSeries) calculable() {}

func (s IntTimeline) calculable()     {}
func (s Int16Timeline) calculable()   {}
func (s Int32Timeline) calculable()   {}
func (s Float32Timeline) calculable() {}
func (s Float64Timeline) calculable() {}
