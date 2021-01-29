package arrays

import (
	"reflect"
	"testing"
)

func assert_eq_slice(t testing.TB, got, want []int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("array of 5 nums", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given: %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{2, 3})
	want := []int{3, 5}

	assert_eq_slice(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{2, 3})
		want := []int{2, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("sum empty", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 18})
		want := []int{0, 21}

		assert_eq_slice(t, got, want)
	})

}
