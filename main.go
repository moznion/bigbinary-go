/*
Package bigbinary provides some utility functions to convert a byte slice to a big integer.
*/
package bigbinary

import (
	"math/big"
)

// ConvertWithBigEndian converts byte slice to a big integer according to big-endian.
func ConvertWithBigEndian(bytes []byte) *big.Int {
	num := big.NewInt(0)
	l := len(bytes) - 1
	for i, b := range bytes {
		n := big.NewInt(int64(b))
		num = num.Add(num, n.Lsh(n, uint((l-i)*8)))
	}
	return num
}

// ConvertWithLittleEndian converts byte slice to a big integer according to little-endian.
func ConvertWithLittleEndian(bytes []byte) *big.Int {
	num := big.NewInt(0)
	for i, b := range bytes {
		n := big.NewInt(int64(b))
		num = num.Add(num, n.Lsh(n, uint(i*8)))
	}
	return num
}
