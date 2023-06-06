package hexconv

import (
	"fmt"
	"math"
	"strings"
)

var toB16 = map[int]rune{
	0:  '0',
	1:  '1',
	2:  '2',
	3:  '3',
	4:  '4',
	5:  '5',
	6:  '6',
	7:  '7',
	8:  '8',
	9:  '9',
	10: 'A',
	11: 'B',
	12: 'C',
	13: 'D',
	14: 'E',
	15: 'F',
}

var toB10 = make(map[rune]int)

func init() {
	for b10, b16 := range toB16 {
		toB10[b16] = b10
	}
}

// "FF" -> 255
func FromHex(b16 string) (int, error) {
	if len(b16) == 0 {
		return 0, fmt.Errorf("cannot convert empty hexidecimal string to decimal value")
	}

	b16 = strings.ToUpper(b16)

	isNegative := b16[0] == '-'
	if isNegative {
		b16 = b16[1:]
	}

	b10 := 0
	exponent := len(b16) - 1

	for _, chrB16 := range b16 {
		digitValueB10, ok := toB10[chrB16]
		if !ok {
			return 0, fmt.Errorf("invalid hexidecimal digit in string: '%v'", chrB16)
		}
		b10 += pow(16, exponent) * digitValueB10
		exponent--
	}

	if isNegative {
		b10 *= -1
	}

	return b10, nil
}

// 255 -> "FF"
func FromDecimal(b10 int) (string, error) {
	var b16 strings.Builder

	isNegative := b10 < 0
	if isNegative {
		b16.WriteRune('-')
		b10 *= -1
	}

	b16Length := Base16Length(b10)
	exponent := b16Length - 1

	for i := 0; i < b16Length; i++ {
		positionalValueB10 := pow(16, exponent)
		digitValueB10 := b10 / positionalValueB10
		b10 %= positionalValueB10

		b16.WriteRune(toB16[digitValueB10])
		exponent--
	}

	return b16.String(), nil
}

func pow(x, y int) int {
	res := 1
	for ; y > 0; y-- {
		res *= x
	}
	return res
}

// take a base 16 logarithm of the decimal value and return the next highest integer
// e.g 255 -> 2
func Base16Length(b10 int) int {
	if b10 == 0 { // can't take log of 0
		return 1
	}
	// COB formula to get log base 16
	// logbase2(x) / logbase2(16)
	// always need to return the next highest integer
	logOfVal := math.Log2(float64(b10)) / 4
	return int(math.Floor(logOfVal)) + 1
}
