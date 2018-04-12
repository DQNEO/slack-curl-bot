package main

import (
	"os/exec"
	"log"
	"fmt"
	"strings"
)

func execCurl(input string) string {
	log.Printf("%s\n", input)
	substr := input[5:len(input)]

	cmd := exec.Command("curl", substr)
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

func handle(input string) string {
	text := input
	var output string
	if isCurlCommand(text) {
		log.Printf("it's curl\n")
		body := execCurl(text)
		output = fmt.Sprintf("```\n%s```", body)
	} else {
		log.Printf("it's NOT curl\n")
		output = "(nothing to do)"
	}
	return output
}

