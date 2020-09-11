package main

import "testing"

func TestInfix2Postfix(t *testing.T) {
	type args struct {
		input string
		ops   map[rune]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "arithmetic  1",
			args: args{
				input: "a+b",
				ops:   map[rune]int{'+': 1},
			},
			want: "ab+",
		},
		{
			name: "arithmetic 2",
			args: args{
				input: "(a+b)*(c-d)",
				ops:   map[rune]int{'+': 1, '-': 1, '*': 2},
			},
			want: "ab+cd-*",
		},
		{
			name: "arithmetic 3",
			args: args{
				input: "a+b*(c^d-e)^(f+g*h)-i",
				ops:   map[rune]int{'+': 1, '-': 1, '*': 2, '/': 2, '^': 3},
			},
			want: "abcd^e-fgh*+^*+i-",
		},
		{
			name: "regex 1",
			args: args{
				input: "a.b",
				ops:   map[rune]int{'.': 1},
			},
			want: "ab.",
		},
		{
			name: "regex 2",
			args: args{
				input: "a?(a+b)*?b",
				ops:   map[rune]int{'+': 1, '?': 2, '*': 3},
			},
			want: "aab+*?b?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Infix2Postfix(tt.args.input, tt.args.ops); got != tt.want {
				t.Errorf("Infix2Postfix() = %v, want %v", got, tt.want)
			}
		})
	}
}
