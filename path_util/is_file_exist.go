package path_util

import (
	"os"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	} else if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
