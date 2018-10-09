package YAMLConverter

import (
	"../error_checker"

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

func Array(arr []interface{}) string{
	isArrRegex, err := regexp.Compile("^\\[\\]")
	checker.CheckErr(err)
	targetArr := make([]string, 0)

	for _, ele := range arr {
		if isArrRegex.MatchString(fmt.Sprintf("%T", ele)) {
			//re
			targetArr = append(targetArr, fmt.Sprintf(" \n- %s", Array((ele.([]interface{})))))
		} else {
			targetArr = append(targetArr, fmt.Sprintf(" - %s", ele))
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

