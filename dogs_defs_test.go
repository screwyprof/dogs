package dogs_test

import "github.com/genkami/dogs"

var (
	intSemigroup dogs.Semigroup[int] = &dogs.DefaultSemigroup[int]{
		CombineImpl: func(x, y int) int { return x + y },
	}
	intMonoid = &dogs.Monoid[int]{
		Semigroup: intSemigroup,
		Empty: func() int { return 0 },
	}
	intEq dogs.Eq[int] = &dogs.DefaultEq[int]{
		EqualImpl: func(x, y int) bool { return x == y },
	}
	intOrd dogs.Ord[int] = &dogs.DefaultOrd[int]{
		CompareImpl: func(x, y int) dogs.Ordering {
			if x < y {
				return dogs.LT
			} else if x == y {
				return dogs.EQ
			} else {
				return dogs.GT
			}
		},
	}

	stringSemigroup dogs.Semigroup[string] = &dogs.DefaultSemigroup[string]{
		CombineImpl: func(x, y string) string { return x + y },
	}
	stringMonoid = &dogs.Monoid[string]{
		Semigroup: stringSemigroup,
		Empty: func() string { return "" },
	}
	stringEq dogs.Eq[string] = &dogs.DefaultEq[string]{
		EqualImpl: func(x, y string) bool { return x == y },
	}
	stringOrd dogs.Ord[string] = &dogs.DefaultOrd[string]{
		CompareImpl: func(x, y string) dogs.Ordering {
			if x < y {
				return dogs.LT
			} else if x == y {
				return dogs.EQ
			} else {
				return dogs.GT
			}
		},
	}
)