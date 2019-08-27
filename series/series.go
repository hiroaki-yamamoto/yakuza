package series

func (s GeneralSeries) series() {}
func (s IntSeries) series()     {}
func (s Int16Series) series()   {}
func (s Int32Series) series()   {}
func (s Int64Series) series()   {}
func (s Float32Series) series() {}
func (s Float64Series) series() {}

func (s GeneralMapSeries) series() {}
func (s IntMapSeries) series()     {}
func (s Int16MapSeries) series()   {}
func (s Int32MapSeries) series()   {}
func (s Int64MapSeries) series()   {}
func (s Float32MapSeries) series() {}
func (s Float64MapSeries) series() {}

func (s GeneralTimeline) series() {}
func (s IntTimeline) series()     {}
func (s Int16Timeline) series()   {}
func (s Int32Timeline) series()   {}
func (s Float32Timeline) series() {}
func (s Float64Timeline) series() {}
