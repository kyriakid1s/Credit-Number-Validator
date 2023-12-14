package luhnalgor

import (
	"strconv"
)

func Validate(creditNumber string) (bool, error) {
	total := 0
	for i := len(creditNumber) - 1; i >= 0; i-- {
		sum := 0
		digit, err := strconv.Atoi(string(creditNumber[i]))
		if err != nil {
			return false, err
		}
		if i%2 == 0 {
			digit *= 2
		}
		sum = (digit / 10) + (digit % 10)
		total += sum
	}
	return (total % 10) == 0, nil
}
