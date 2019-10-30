package captcha

import "testing"

func Test_Captcha(t *testing.T) {
	type args struct {
		pattern int
		leftOperand int
		operator int
		rightOperand int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "captcha_whenPatternIs1_leftOperandIs1_operatorIs1_rightOperandIs1_ShouldReturn1PlusOne", args: args{1, 1, 1, 1}, want: "1 + One"},
		{name: "captcha_whenPatternIs1_leftOperandIs2_operatorIs1_rightOperandIs1_ShouldReturn2PlusOne", args: args{1, 2, 1, 1}, want: "2 + One"},
		{name: "captcha_whenPatternIs1_leftOperandIs3_operatorIs1_rightOperandIs1_ShouldReturn3PlusOne", args: args{1, 3, 1, 1}, want: "3 + One"},
		{name: "captcha_whenPatternIs1_leftOperandIs3_operatorIs1_rightOperandIs2_ShouldReturn3PlusTwo", args: args{1, 3, 1, 2}, want: "3 + Two"},
		{name: "captcha_whenPatternIs1_leftOperandIs3_operatorIs1_rightOperandIs3_ShouldReturn3PlusThree", args: args{1, 3, 1, 3}, want: "3 + Three"},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captcha(tt.args.pattern, tt.args.leftOperand, tt.args.operator, tt.args.rightOperand); got != tt.want {
				t.Errorf("Test_Captcha() = %v, want %v", got, tt.want)
			}
		})
	}
}
