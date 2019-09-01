package series

import (
	"math"
	"math/big"
	"runtime"
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
	exLen := int(math.Min(float64(len(a)), float64(len(b))))
	res := make(BigIntSeries, exLen)
	fch := make(chan fragment, runtime.NumCPU()+1)
	for i := 0; i < runtime.NumCPU()+1; i++ {
		end := (i + 1) * exLen / (runtime.NumCPU() + 1)
		go func(pa, pb ICalculable, i, end int, fret chan fragment) {
			pres := make(BigIntSeries, len(pa))
			for j := 0; j < len(pres); j++ {
				pres[j] = big.NewInt(0).Add(pa[j]).Add(pb[j])
			}
			fret <- fragment{
				Offset: i,
				End:    end,
				Items:  &pres,
			}
			close(fret)
		}(a[i:end], b[i:end], i, end, fch)
	}
	for range fch {
		select {
		case v := <-fch:
			copy(res[v.Offset:v.End], v.Items)
		default:
		}
	}
}
