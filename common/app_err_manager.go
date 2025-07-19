package common

import (
	"fmt"
	"os"
)

/* ending app with error*/
func EndWithErr(err error) {
	fmt.Fprintln(os.Stderr,err)
	os.Exit(1)
}