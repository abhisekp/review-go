package hashtable

import (
	"review-go/utils"
	"slices"
	"unicode"
)

func Hash1(key string, hashSize int) int {
	return 1 % hashSize
}

func Hash2(key string, hashSize int) int {
	return len(key) % hashSize
}

func Hash3(key string, hashSize int) int {
	return int(key[0]) % hashSize
}

func Hash4(key string, hashSize int) int {
	primes := slices.Collect(utils.Take(hashSize, utils.Primes()))
	// fmt.Println(primes)
	sum := 0
	for _, ch := range key {
		idx := int(unicode.ToLower(ch) - 'a')
		sum += primes[idx]
	}
	return sum % hashSize
}
