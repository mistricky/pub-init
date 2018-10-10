package main

import (
	"./convert_yaml"
	"./error_checker"
	"./path_util"

	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	TARGET_FILE_NAEM = "pubspec.yaml"
)

type Question struct {
	questionString string
	value          string
}

var questions []Question

func printString(content string) {
	fmt.Printf("%c[%d;%dm%s%c[0m", 0x1B, 0, 34, content, 0x1B)
}

func processValue(value string) string {
	if value != "" {
		return fmt.Sprintf("(%s) ", value)
	}

	return ""
}

func getInfo() {
	input := bufio.NewScanner(os.Stdin)

	for i := 0; i < len(questions); i++ {
		printString(fmt.Sprintf("%s %s: ", questions[i].questionString, processValue(questions[i].value)))
		input.Scan()

		if input.Text() != "" {
			questions[i].value = input.Text()
		}
	}
}

func buildFile() {
	f, err := os.Create("./" + TARGET_FILE_NAEM)
	defer f.Close()

	content := make([]string, 0)

	for i := 0; i < len(questions); i++ {
		content = append(content, YAMLConverter.Simple(questions[i].questionString, questions[i].value))
	}

	checker.CheckErr(err)
	f.WriteString(strings.Join(content, "\n"))

}

func main() {
	if !path_util.IsExist(TARGET_FILE_NAEM) {
		getInfo()
		buildFile()
	} else {
		fmt.Println("pubspec is exist")
	}
}

func init() {
	// fill questions
	questions = []Question{
		{
			"name",
			path_util.GetBasePath(path_util.PwdMust()),
		},
		{
			"description",
			"",
		},
		{
			"version",
			"0.0.1",
		},
		{
			"environment",
			"sdk: '>=1.24.0 <=2.0.0'",
		},
		{
			"isRepository",
			"true",
		},
	}
}
