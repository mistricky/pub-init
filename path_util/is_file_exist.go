package path

import (
	"os"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return true
	} else if (os.IsNotExist(err)) {
		return false
	} else {
		return false
	}
}
