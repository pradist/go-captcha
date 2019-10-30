package captcha

import (
	"strconv"
)

func captcha(pattern int, leftOperand int, operator int, rightOperand int) string  {
	numbers := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	return strconv.Itoa(leftOperand) + " + " + numbers[rightOperand - 1]
}
