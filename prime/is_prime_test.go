package prime

import (
	"fmt"
	"testing"
)

// 2^64 == 18446744073709551616

// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

func TestIsPrime(t *testing.T) {
	primeNumbers := []int64{2, 3, 5, 7, 11, 13}
	for _, num := range primeNumbers {
		if !IsPrime(num) {
			t.Error()
		}
	}
}

func TestGetPrimeNumbers(t *testing.T) {
	var begin uint64 = 0
	var limit uint64 = 10000
	// var limit uint64 = math.MaxUint64

	primeNumbers := GetPrimeNumbers(begin, limit)

	fmt.Println(primeNumbers)
}
