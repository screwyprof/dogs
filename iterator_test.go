package dogs_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/genkami/dogs"
)

func TestSliceFromIterator(t *testing.T) {
	subject := func(xs []int) []int {
		return dogs.SliceFromIterator(dogs.Slice[int](xs).Iter())
	}
	assert.Equal(t, subject([]int{}), []int{})
	assert.Equal(t, subject([]int{1}), []int{1})
	assert.Equal(t, subject([]int{1, 2, 3}), []int{1, 2, 3})
}

func TestMap(t *testing.T) {
	subject := func(xs []int) []string {
		it := dogs.Slice[int](xs).Iter()
		mapped := dogs.Map[int, string](it, func(x int) string {
			return strconv.FormatInt(int64(x), 10)
		})
		return dogs.SliceFromIterator(mapped)
	}

	assert.Equal(t, subject([]int{}), []string{})
	assert.Equal(t, subject([]int{1}), []string{"1"})
	assert.Equal(t, subject([]int{1, 2}), []string{"1", "2"})
	assert.Equal(t, subject([]int{1, 2, 3}), []string{"1", "2", "3"})
}

func TestFold(t *testing.T) {
	add := func(x string, y int) string {
		return x + strconv.FormatInt(int64(y), 10)
	}
	subject := func(x string, xs []int) string {
		it := dogs.Slice[int](xs).Iter()
		return dogs.Fold[string, int](x, it, add)
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject("", []int{}), "")
		assert.Equal(t, subject("a", []int{}), "a")
		assert.Equal(t, subject("xyzzy", []int{}), "xyzzy")
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, subject("", []int{1}), "1")
		assert.Equal(t, subject("a", []int{2}), "a2")
	})

	t.Run("multiple elements", func(t *testing.T) {
		assert.Equal(t, subject("", []int{1, 2, 3}), "123")
		assert.Equal(t, subject("hoge", []int{3, 2, 1, 0}), "hoge3210")
	})
}

func TestZip(t *testing.T) {
	type Pair = dogs.Pair[int, string]
	subject := func(xs []int, ys []string) []Pair {
		xit := dogs.Slice[int](xs).Iter()
		yit := dogs.Slice[string](ys).Iter()
		zipped := dogs.Zip(xit, yit)
		return dogs.SliceFromIterator(zipped)
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject([]int{}, []string{}),[]Pair{})
		assert.Equal(t, subject([]int{1, 2}, []string{}), []Pair{})
		assert.Equal(t, subject([]int{}, []string{"a", "b"}), []Pair{})
	})

	t.Run("same length", func(t *testing.T) {
		assert.Equal(t, subject([]int{1}, []string{"a"}), []Pair{{1, "a"}})
		assert.Equal(t, subject([]int{1, 2}, []string{"a", "b"}), []Pair{{1, "a"}, {2, "b"}})
		assert.Equal(t, subject([]int{1, 2, 3}, []string{"a", "b", "c"}), []Pair{{1, "a"}, {2, "b"}, {3, "c"}})
	})

	t.Run("different length", func(t *testing.T) {
		assert.Equal(t, subject([]int{1, 2, 3}, []string{"a", "b"}), []Pair{{1, "a"}, {2, "b"}})
		assert.Equal(t, subject([]int{1, 2}, []string{"a", "b", "c"}), []Pair{{1, "a"}, {2, "b"}})
		assert.Equal(t, subject([]int{1, 2, 3}, []string{"a"}), []Pair{{1, "a"}})
	})
}

func TestSumWithInit(t *testing.T) {
	subject := func(x int, xs []int) int {
		return dogs.SumWithInit(x, dogs.Slice[int](xs).Iter(), intSemigroup)
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{}), 0)
		assert.Equal(t, subject(1, []int{}), 1)
		assert.Equal(t, subject(999, []int{}), 999)
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{1}), 1)
		assert.Equal(t, subject(1, []int{2}), 3)
		assert.Equal(t, subject(123, []int{456}), 579)
	})

	t.Run("multiple elements", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{1, 2, 3}), 6)
		assert.Equal(t, subject(1, []int{10, 100, 1000}), 1111)
	})
}

func TestSum(t *testing.T) {
	subject := func(xs []int) int {
		return dogs.Sum(dogs.Slice[int](xs).Iter(), intMonoid)
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject([]int{}), 0)
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, subject([]int{0}), 0)
		assert.Equal(t, subject([]int{1}), 1)
		assert.Equal(t, subject([]int{999}), 999)
	})

	t.Run("multiple elements", func(t *testing.T) {
		assert.Equal(t, subject([]int{1, 2}), 3)
		assert.Equal(t, subject([]int{1, 2, 3}), 6)
		assert.Equal(t, subject([]int{1, 10, 100, 1000}), 1111)
	})
}
