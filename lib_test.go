package main

import "testing"

func Test_isCurlCommand(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"is curl",
		args{input:"curl foo",},
		true,
		},
		{
			"is not curl",
			args{input:"hello"},
			false,
		},
		{
			"is not curl",
			args{input:"curls"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCurlCommand(tt.args.input); got != tt.want {
				t.Errorf("isCurlCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
