package main

import (
	"./path_util"
	"./convert_yaml"

	"bufio"
	"fmt"
	"os"
)

const (
	TARGET_FILE_NAEM = "pubspec.yaml"
)

type Question struct {
	questionString string
	value string
}

var questions []Question

func printString(content string){
	fmt.Printf("%c[%d;%dm%s%c[0m", 0x1B, 0, 34, content, 0x1B)
}

func processValue(value string) string{
	if value != "" {
		return fmt.Sprintf("(%s) ",value)
	}

	return ""
}

func getInfo() {
	input := bufio.NewScanner(os.Stdin)

	for i := 0;i < len(questions);i++ {
		printString(fmt.Sprintf("%s %s: ",questions[i].questionString, processValue(questions[i].value)))
		input.Scan()

		if input.Text() != "" {
			questions[i].value = input.Text()
		}
	}
}

// convert question array to yaml
//func covertYAML(){
//
//}
//
//func buildFile() (ok bool) {
//	isExist := path_util.IsExist(TARGET_FILE_NAEM)
//	if(isExist) {
//		return false
//	} else {
//		f, err := os.Create(TARGET_FILE_NAEM)
//		defer f.Close()
//		checker.CheckErr(err)
//		f.WriteString(strings.Join([]string{
//
//		}, "\n"))
//	}
//}

func main(){
	//name := YAMLConverter.String("name", "zhangsan")
	//fmt.Println(name)

	//obj := YAMLConverter.MultipleObject([]string{
	//	"name",
	//	"age",
	//}, []string{
	//	"zhangsan",
	//	"19",
	//})
	//fmt.Println(YAMLConverter.Complex("person", obj))

	arr := YAMLConverter.Array([]interface{}{"a", []interface{}{
		"e",
		"f",
	}, "c"})

	fmt.Println(arr)
}

func init(){
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

