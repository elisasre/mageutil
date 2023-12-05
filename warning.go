package mageutil

import (
	"fmt"
	"os"
)

func deprecated() {
	fmt.Fprintf(os.Stderr, "Warning: functions under mageutil are deprecated, use domain specific packages instead!\n")
}
