package main

import (
	"os/exec"
	"log"
	"fmt"
	"strings"
	"regexp"
	"bytes"
)

func execCurl(input string) string {
	cmd := textToCmd(input)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return outStr
}

func isCurlCommand(text string) bool {
	return strings.HasPrefix(text, "curl")
}

var BotId = "@UA5UWB2NB"

func isMentionToMe(input string) bool {
	return strings.HasPrefix(input, fmt.Sprintf("<%s>", BotId))
}

func removeMention(input string) string {
	return input[len("<"+BotId+">"):]
}

func textToCmd(input string) *exec.Cmd {
	log.Printf("text:%s\n", input)
	var substr string
	if len(input) <= 5 {
		substr = "--help"
	} else {
		substr = input[5:len(input)]
	}

	log.Printf("substr:%s\n", substr)
	arg := unwrapUrl(substr)
	args := strings.Split(arg, " ")
	log.Printf("args:%s\n", args)
	cmd := exec.Command("curl", args...)
	return cmd
}

func unwrapUrl(input string) string {
	re2 := regexp.MustCompile("<(https?://.*)>")
	output := re2.ReplaceAllString(input, "$1")
	return output
}

func handle(input string) string {
	var output string
	if ! isMentionToMe(input) {
		log.Printf("not talking to me...\n")
		return ""
	}

	input = removeMention(input)
	log.Printf(input)
	input = strings.TrimLeft(input," ")
	if isCurlCommand(input) {
		log.Printf("it's curl\n")
		body := execCurl(input)
		output = fmt.Sprintf("```\n%s```", body)
	} else {
		log.Printf("it's NOT curl\n")
		output = "(nothing to do)"
	}
	return output
}

