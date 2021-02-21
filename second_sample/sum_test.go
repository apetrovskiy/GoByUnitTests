package second_sample

import (
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoByUnitTests Suite", func() {
	var (
		numbers01 []int
		numbers02 []int
	)

	BeforeEach(func() {
		numbers01 = []int{1, 2, 3, 4, 5}

		numbers02 = []int{5}
	})

	Describe("Summarization", func() {
		Context("Non-empty array", func() {
			It("should be 15", func() {
				Expect(Sum(numbers01)).To(Equal(15))
			})
		})

		Context("Empty array", func() {
			It("should be 0", func() {
				Expect(Sum(numbers02)).To(Equal(0))
			})
		})
	})
})

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
