package mageutil

import (
	"log"
	"os"
	"strconv"

	"github.com/magefile/mage/mg"
)

// init sets sane defaults for mage and go.
func init() {
	CGO(false)
	Verbose(true)
}

// Verbose can be used to control mage's verbose state. By default this package will set mage in verbose state.
func Verbose(enabled bool) {
	MustSetEnv(mg.VerboseEnv, strconv.FormatBool(enabled))
}

// CGO can be used to enable of disable CGO. By default this package will disable CGO.
func CGO(enabled bool) {
	if enabled {
		MustSetEnv("CGO_ENABLED", "1")
	} else {
		MustSetEnv("CGO_ENABLED", "0")
	}
}

func MustSetEnv(k, v string) {
	err := os.Setenv(k, v)
	if err != nil {
		log.Fatalf("Failed to set '%s' to '%s': %s", k, v, err)
	}
}
