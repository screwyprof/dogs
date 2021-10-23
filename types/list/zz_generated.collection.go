// Code generated by gen-functions; DO NOT EDIT.

package list

import (
	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/classes/cmp"
	"github.com/genkami/dogs/types/iterator"
	"github.com/genkami/dogs/types/pair"
)

// Some packages are unused depending on -include CLI option.
// This prevents compile error when corresponding functions are not defined.
var _ = (algebra.Monoid[int])(nil)
var _ = (cmp.Ord[int])(nil)
var _ = (iterator.Iterator[int])(nil)
var _ = (*pair.Pair[int, int])(nil)

// Filter returns a collection that only returns elements that satisfies given predicate.
func Filter[T any](xs *List[T], fn func(T) bool) *List[T] {
	return FromIterator[T](iterator.Filter[T](xs.Iter(), fn))
}

// Find returns a first element in xs that satisfies the given predicate fn.
// It returns false as a second return value if no elements are found.
func Find[T any](xs *List[T], fn func(T) bool) (T, bool) {
	return iterator.Find[T](xs.Iter(), fn)
}

// FindElem returns a first element in xs that equals to e in the sense of given Eq.
// It returns false as a second return value if no elements are found.
func FindElem[T any](eq cmp.Eq[T]) func(xs *List[T], e T) (T, bool) {
	return func(xs *List[T], e T) (T, bool) {
		return iterator.FindElem[T](eq)(xs.Iter(), e)
	}
}

// FindElemIndex returns a first index of an element in xs that equals to e in the sense of given Eq.
// It returns negative value if no elements are found.
func FindElemIndex[T any](eq cmp.Eq[T]) func(xs *List[T], e T) int {
	return func(xs *List[T], e T) int {
		return iterator.FindElemIndex[T](eq)(xs.Iter(), e)
	}
}

// FindIndex returns a first index of an element in xs that satisfies the given predicate fn.
// It returns negative value if no elements are found.
func FindIndex[T any](xs *List[T], fn func(T) bool) int {
	return iterator.FindIndex[T](xs.Iter(), fn)
}

// Fold accumulates every element in a collection by applying fn.
func Fold[T any, U any](init T, xs *List[U], fn func(T, U) T) T {
	return iterator.Fold[T, U](init, xs.Iter(), fn)
}

// ForEach applies fn to each element in xs.
func ForEach[T any](xs *List[T], fn func(T)) {
	iterator.ForEach[T](xs.Iter(), fn)
}

// Map returns a collection that applies fn to each element of xs.
func Map[T, U any](xs *List[T], fn func(T) U) *List[U] {
	return FromIterator[U](iterator.Map[T, U](xs.Iter(), fn))
}

// Sum sums up all values in xs.
// It returns m.Empty() when xs is empty.
func Sum[T any](m algebra.Monoid[T]) func(xs *List[T]) T {
	return func(xs *List[T]) T {
		var s algebra.Semigroup[T] = m
		return SumWithInit[T](s)(m.Empty(), xs)
	}
}

// SumWithInit sums up init and all values in xs.
func SumWithInit[T any](s algebra.Semigroup[T]) func(init T, xs *List[T]) T {
	return func(init T, xs *List[T]) T {
		return Fold[T, T](init, xs, s.Combine)
	}
}

// Zip combines two collections into one that contains pairs of corresponding elements.
func Zip[T, U any](a *List[T], b *List[U]) *List[pair.Pair[T, U]] {
	return FromIterator[pair.Pair[T, U]](iterator.Zip(a.Iter(), b.Iter()))
}
