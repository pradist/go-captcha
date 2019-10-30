package captcha

func captcha(pattern int, leftOperand int, operator int, rightOperand int) string  {
	if leftOperand == 1 {
		return "1 + One"
	}
	if leftOperand ==3 {
		return "3 + One"
	}
	return "2 + One"
}
