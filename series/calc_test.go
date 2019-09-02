package series_test

import (
	"math/big"
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hiroaki-yamamoto/yakuza/series"
)

var random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

var _ = Describe("Normal Series", func() {
	Describe("Addition", func() {
		Context("with same size integers", func() {
			const len = 100
			arr1 := make(series.BigIntSeries, len)
			arr2 := make(series.BigIntSeries, len)
			sol := make(series.BigIntSeries, len)
			Context("full filled with random integers", func() {
				for i := 0; i < len; i++ {
					arr1[i] = big.NewInt(random.Int63())
					arr2[i] = big.NewInt(random.Int63())
					sol[i] = big.NewInt(0).Add(arr1[i], arr2[i])
				}
				It("should successfully and correct result.", func() {
					result := series.Add(arr1, arr2, nil)
					Expect(result.(series.BigIntSeries)).To(Equal(sol))
				})
			})
		})
	})
})
