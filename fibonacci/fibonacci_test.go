package fibonacci

import "testing"

type strategy func(uint) uint64

type testcase struct {
	in   uint
	want uint64
}

// You can't have a constant array nor a function that returns [...]testcase
func tests() []testcase {
	items := [...]testcase{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{13, 233},
		{14, 377},
		{15, 610},
		{16, 987},
		{17, 1597},
		{18, 2584},
		{19, 4181},
		{20, 6765},
	}
	return items[:]
}

// Alternatively, without the explicit strategy type:
//     func baseTest(t *testing.T, f func(uint64) string) {
func baseTest(t *testing.T, f strategy) {
	for _, item := range tests() {
		result := f(item.in)
		if result != item.want {
			t.Errorf("fibonacci(%d) == %d but got %d", item.in, item.want, result)
		}
	}
}

func TestFib1(t *testing.T) {
	baseTest(t, fib1)
}

func TestFib2(t *testing.T) {
	baseTest(t, fib2)
}

func TestFib3(t *testing.T) {
	baseTest(t, fib3)
}
