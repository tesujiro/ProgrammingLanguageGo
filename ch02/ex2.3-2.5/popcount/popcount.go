package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// excercise 2.3
func PopCount2(x uint64) int {
	result := 0
	for i := uint(0); i < 9; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

// excercise 2.4
func PopCount3(x uint64) int {
	result := 0
	for i := uint(0); i <= 64; i++ {
		result += int(x) & 1
		x = uint64(x >> 1)
	}
	return result
}

// excercise 2.5
func PopCount4(x uint64) int {
	c := 0
	for ; x > 0; x = x & (x - 1) {
		c++
	}
	return c
}
