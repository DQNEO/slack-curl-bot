package main

import (
	"os/exec"
	"log"
	"fmt"
	"strings"
)

func execCurl(input string) string {
	cmd := textToCmd(input)
	byts, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	out := string(byts)
	fmt.Printf(out)
	return out
}

func isCurlCommand(text string) bool {
	return strings.HasPrefix(text, "curl ")
}

var BotId = "@UA5UWB2NB"

func isMentionToMe(input string) bool {
	return strings.HasPrefix(input, fmt.Sprintf("<%s> ", BotId))
}

func textToCmd(input string) *exec.Cmd {
	log.Printf("%s\n", input)
	substr := input[5:len(input)]

	cmd := exec.Command("curl", substr)
	return cmd
}

func handle(input string) string {
	var output string
	if ! isMentionToMe(input) {
		log.Printf("not talking to me...\n")
		return ""
	}

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

