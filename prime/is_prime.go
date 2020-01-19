package prime

import "math/big"

// IsPrime .
func IsPrime(x int64) bool {
	i := big.NewInt(x)
	return i.ProbablyPrime(0)
}

// GetPrimeNumbers .
func GetPrimeNumbers(begin, limit uint64) []uint64 {
	primeNumbers := []uint64{}

	for ; begin < limit; begin++ {
		if IsPrime(int64(begin)) {
			primeNumbers = append(primeNumbers, begin)
		}
	}

	return primeNumbers
}
