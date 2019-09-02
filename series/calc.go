package series

import (
	"math"
	"math/big"
	"runtime"
	"sync"
)

// Calculation code for series

type fragment struct {
	Offset int
	End    int
	Items  ICalculable
}

// Add calculates the addtion of a and b and returns the result.
func Add(
	a, b ICalculable,
	fillValue interface{},
) ICalculable {
	opA := a.(BigIntSeries)
	opB := b.(BigIntSeries)
	exLen := int(math.Min(float64(len(opA)), float64(len(opB))))
	res := make(BigIntSeries, exLen)
	numSeq := runtime.NumCPU() + 1
	seqLen := exLen / numSeq
	fch := make(chan fragment, numSeq)
	// defer close(fch)
	var w sync.WaitGroup
	w.Add(exLen / seqLen)
	for i := 0; i < exLen; i += seqLen {
		_end := int(math.Min(float64(i+seqLen), float64(exLen)))
		go func(
			pa, pb ICalculable,
			offset, end int,
			wg *sync.WaitGroup,
			fret chan<- fragment,
		) {
			defer wg.Done()
			opPA := pa.(BigIntSeries)
			opPB := pb.(BigIntSeries)
			pres := make(BigIntSeries, len(opPA))
			for j := 0; j < len(pres); j++ {
				pres[j] = big.NewInt(0).Add(opPA[j], opPB[j])
			}
			fret <- fragment{
				Offset: offset,
				End:    end,
				Items:  &pres,
			}
		}(opA[i:_end], opB[i:_end], i, _end, &w, fch)
	}
	go func(wg *sync.WaitGroup, ch chan fragment) {
		wg.Wait()
		close(ch)
	}(&w, fch)
	for v := range fch {
		copy(res[v.Offset:v.End], *v.Items.(*BigIntSeries))
	}
	return res
}
