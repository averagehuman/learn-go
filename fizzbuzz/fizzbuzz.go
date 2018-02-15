package fizzbuzz

import "os"
import "fmt"
import "strconv"

const FIZZ = "Fizz"
const BUZZ = "Buzz"
const FIZZBUZZ = FIZZ + BUZZ

func fizzbuzz(x uint64) string {
	s := ""
	if x%3 == 0 {
		s += FIZZ
	}
	if x%5 == 0 {
		s += BUZZ
	}
	if s == "" {
		// Not divisible by 3 or 5
		return strconv.FormatUint(x, 10)
	}
	return s
}

func fizzbuzz2(x uint64) string {
	switch {
	case x%15 == 0:
		return FIZZBUZZ
	case x%5 == 0:
		return BUZZ
	case x%3 == 0:
		return FIZZ
	default:
		return strconv.FormatUint(x, 10)
	}
}

func main() {
	args := os.Args[1:]
	n, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		fmt.Println("Invalid input %s", args[0])
	}
	fmt.Println(fizzbuzz2(n))
}
