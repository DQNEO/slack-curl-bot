package main

import (
	"os/exec"
	"reflect"
	"testing"
)

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
			args{input: "curl foo"},
			true,
		},
		{
			"is not curl",
			args{input: "hello"},
			false,
		},
		{
			"is not curl",
			args{input: "curls"},
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

func Test_textToCmd(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		{
			"arg 1",
			args{input:"curl --help"},
			&exec.Cmd{
				Path:"curl",
				Args:[]string{"--help"},
			},
		},
		{
			"arg 1",
			args{input:"curl -X get https://google.com/"},
			&exec.Cmd{
				Path:"curl",
				Args:[]string{" -X get https://google.com/"},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := textToCmd(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("textToCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
