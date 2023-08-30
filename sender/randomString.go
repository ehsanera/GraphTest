package main

import (
	"crypto/rand"
	"math/big"
)

const (
	minSizeBytes = int64(50)
	maxSizeBytes = int64(8192) // 8 KB in bytes
)

func generateRandomByteInRange(minByte, maxByte byte) (byte, error) {
	rangeSize := new(big.Int).SetInt64(int64(maxByte - minByte + 1))
	randomIndex, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		return 0, err
	}

	return minByte + byte(randomIndex.Int64()), nil
}

func generateRandomString() string {
	sizeRange := new(big.Int).Sub(big.NewInt(maxSizeBytes), big.NewInt(minSizeBytes))
	size, err := rand.Int(rand.Reader, sizeRange)
	if err != nil {
		return ""
	}

	randomByteSlice := make([]byte, size.Int64())
	for i := range randomByteSlice {
		randomByte, err := generateRandomByteInRange('a', 'z')
		if err != nil {
			return ""
		}
		randomByteSlice[i] = randomByte
	}

	return string(randomByteSlice)
}
