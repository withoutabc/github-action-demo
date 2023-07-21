package demo

import "testing"

func TestLengthSmall(t *testing.T) {
	type test struct {
		input string
		want  bool
	}
	tests := map[string]test{
		"test1": {input: "(())", want: true},
		"test2": {input: "{(]]", want: false},
		"test3": {input: "([])", want: true},
	}
	for name, tc := range tests {
		got := isValid(tc.input)
		t.Run(name, func(t *testing.T) {
			if got != tc.want {
				t.Errorf("expected:%t, got:%t %#v", tc.want, got, tc)
				//	t.Errorf("expected:%d, got:%d", tc.want, got)
			}
		})
	}
}
