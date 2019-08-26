package romannumerals

import (
	"errors"
)

type number struct {
	thousands int
	hundreds  int
	tens      int
	ones      int
}

func isInvalid(n *int) bool {
	return *n <= 0 || *n > 3000
}

func (n *number) resolveOnes() string {
	mapping := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return mapping[n.ones]
}

func (n *number) resolveTens() string {
	mapping := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	return mapping[n.tens]
}

func (n *number) resolveHundreds() string {
	mapping := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	return mapping[n.hundreds]
}

func (n *number) resolveThousands() string {
	mapping := [4]string{"", "M", "MM", "MMM"}
	return mapping[n.thousands]
}

func (n *number) convert() string {
	return n.resolveThousands() + n.resolveHundreds() + n.resolveTens() + n.resolveOnes()
}

// ToRomanNumeral converts arabic numbers to roman numbers.
func ToRomanNumeral(n int) (string, error) {
	if isInvalid(&n) {
		return "", errors.New("invalid number")
	}
	thousands := n / 1000
	hundreds := (n - (thousands * 1000)) / 100
	tens := (n - ((thousands * 1000) + (hundreds * 100))) / 10
	ones := n - ((thousands * 1000) + (hundreds * 100) + (tens * 10))
	num := &number{thousands, hundreds, tens, ones}
	return num.convert(), nil
}
