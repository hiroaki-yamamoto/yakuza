package series_test

import (
	"math/big"
	"math/rand"
	"runtime"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hiroaki-yamamoto/yakuza/series"
)

var random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

var _ = Describe("Normal Series", func() {
	Describe("Addition", func() {
		Context("Same size integers", func() {
			generateTests := func(sz int) func() {
				return func() {
					arr1 := make(series.BigIntSeries, sz)
					arr2 := make(series.BigIntSeries, sz)
					sol := make(series.BigIntSeries, sz)
					Context("Full filled with random integers", func() {
						BeforeEach(func() {
							for i := 0; i < sz; i++ {
								arr1[i] = big.NewInt(random.Int63())
								arr2[i] = big.NewInt(random.Int63())
								sol[i] = big.NewInt(0).Add(arr1[i], arr2[i])
							}
						})
						It("should be a correct result.", func() {
							result := series.Add(arr1, arr2, nil)
							Expect(result.(series.BigIntSeries)).To(Equal(sol))
						})
					})

					// Context("filled with nil at several opA's elems", func() {
					// 	BeforeEach(func() {
					// 		for i := 0; i < sz; i++ {
					// 			arr1[i] = big.NewInt(random.Int63())
					// 			if (random.Int() & 0x01) == 0x01 {
					// 				arr1[i] = nil
					// 			} else {
					// 				arr1[i] = big.NewInt(random.Int63())
					// 			}
					// 			arr2[i] = big.NewInt(random.Int63())
					// 			sol[i] = big.NewInt(0).Add(arr1[i], arr2[i])
					// 		}
					// 	})
					// 	Context("without fillValue", func(){
					// 	})
					// })
				}
			}
			Context(
				"Has extra items",
				generateTests(runtime.NumCPU()*100+rand.Intn(runtime.NumCPU()-1)+1),
			)
			Context("Has no extra items", generateTests(runtime.NumCPU()*100))
		})
	})
})
