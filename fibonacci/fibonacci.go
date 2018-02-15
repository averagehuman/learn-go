package fibonacci

// Basic recursive
func fib1(n uint) uint64 {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

// Basic iterative
func fib2(n uint) uint64 {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	var i uint
	var a, b uint64 = 0, 1
	for i = 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// Fast recursive
// Given F(k) and F(k+1):
//     F(2k) = F(k) * (2 * F(k+1) - F(k))
//     F(2k+1) = pow(F(k+1), 2) + pow(F(k), 2)

func _fib3(n uint) (uint64, uint64) {
	if n == 0 {
		return 0, 1
	}
	a, b := _fib3(n / 2)
	c := a * (b*2 - a)
	d := a*a + b*b
	if n%2 == 0 {
		return c, d
	} else {
		return d, c + d
	}
}

func fib3(n uint) uint64 {
	ret, _ := _fib3(n)
	return ret
}
