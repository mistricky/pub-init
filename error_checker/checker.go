package checker

import (
	"os"
	"fmt"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
