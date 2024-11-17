package main

import (
	"slices"
	"testing"
)

func checkSum(t testing.TB, got []int, want []int) {
	if !slices.Equal(want, got) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		want := 6
		if want != sum {
			t.Errorf("got %d want %d given, %v", sum, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("slice of any size", func(t *testing.T) {
		sum := SumAll([]int{1, 2}, []int{0, 9})

		want := []int{3, 9}
		checkSum(t, sum, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("many slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSum(t, got, want)
	})
	t.Run("safely slice with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}
