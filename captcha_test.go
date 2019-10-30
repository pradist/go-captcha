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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captcha(tt.args.pattern, tt.args.leftOperand, tt.args.operator, tt.args.rightOperand); got != tt.want {
				t.Errorf("Test_Captcha() = %v, want %v", got, tt.want)
			}
		})
	}
}
