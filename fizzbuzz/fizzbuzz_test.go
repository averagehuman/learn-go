package fizzbuzz

import "testing"

type strategy func(uint64) string

type testcase struct {
	in   uint64
	want string
}

// You can't have a constant array nor a function that returns [...]testcase
func tests() []testcase {
	items := [...]testcase{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
		{17, "17"},
		{18, "Fizz"},
		{19, "19"},
		{20, "Buzz"},
		{30, "FizzBuzz"},
	}
	return items[:]
}

// Alternatively, without the explicit strategy type:
//     func baseTest(t *testing.T, f func(uint64) string) {
func baseTest(t *testing.T, f strategy) {
	for _, item := range tests() {
		result := f(item.in)
		if result != item.want {
			t.Errorf("fizzbuzz(%d) == %q but got %q", item.in, item.want, result)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	baseTest(t, fizzbuzz)

}

func TestFizzBuzz2(t *testing.T) {
	var g func(uint64) string = fizzbuzz2
	baseTest(t, g)

}
