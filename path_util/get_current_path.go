package path_util

import (
	"os"
	"../error_checker"
	"regexp"
)

func PwdMust() string {
	dir, err := os.Getwd()
	checker.CheckErr(err)
	return dir
}

func GetBasePath(path string) string {
	reg := regexp.MustCompile("[\\/]([^\\/]+)$")
	params := reg.FindStringSubmatch(path)
	return params[1]
}