// Code generated by gen-functions; DO NOT EDIT.

package iterator

import (
	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/classes/cmp"
	"github.com/genkami/dogs/types/pair"
)

// Some packages are unused depending on -include CLI option.
// This prevents compile error when corresponding functions are not defined.
var _ = (algebra.Monoid[int])(nil)
var _ = (cmp.Ord[int])(nil)
var _ = (Iterator[int])(nil)
var _ = (*pair.Pair[int, int])(nil)

// LiftM promotes a function fn to a monad.
func LiftM[T, U any](fn func(T) U) func(Iterator[T]) Iterator[U] {
	return func(mt Iterator[T]) Iterator[U] {
		return AndThen[T, U](mt, func(t T) Iterator[U] {
			return Pure(fn(t))
		})
	}
}
