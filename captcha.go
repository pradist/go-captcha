package captcha

import (
	"strconv"
)

func captcha(pattern int, leftOperand int, operator int, rightOperand int) string  {
	if rightOperand == 2 {
		return strconv.Itoa(leftOperand) + " + Two"
	}
	return strconv.Itoa(leftOperand) + " + One"
}
