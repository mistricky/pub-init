package YAMLConverter

import (
	"../error_checker"
	"../parser"

	"fmt"
	"regexp"
	"strings"
)

func SingleObject(key string, val string) string {
	return fmt.Sprintf("%s: %s\n", key, val)
}

func MultipleObject(keys []string, vals []string) string {
	objContent := make([]string, 0)

	for index, key := range keys {
		objContent = append(objContent, fmt.Sprintf(" %s: %s", key, vals[index]))
	}

	return strings.Join(objContent, "\n")
}

func Array(arr interface{}, layer int) string{
	slice := parser.Slice(arr)
	isArrRegex, err := regexp.Compile("^\\[\\]")
	checker.CheckErr(err)
	targetArr := make([]string, 0)

	for _, ele := range slice {
		if isArrRegex.MatchString(fmt.Sprintf("%T", ele)) {
			targetArr = append(targetArr, fmt.Sprintf("%s", Array((ele.([]interface{})), layer + 1)))
		} else {
			targetArr = append(targetArr, fmt.Sprintf("%s- %s", repeatSpace(layer), ele))
		}
	}

	return strings.Join(targetArr, "\n")
}

func Complex(key string, val string) string {
	return fmt.Sprintf("%s:\n%s", key, val)
}

func Simple(key string, val string) string {
	return fmt.Sprintf("%s: %s", key, val)
}

func repeatSpace(count int) string {
	spaces := make([]string, 0)

	for i := 0;i < count;i++ {
		spaces = append(spaces, " ")
	}

	return strings.Join(spaces, "")
}
