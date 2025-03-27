// Package tool allows running `go tool` commands as mage targets.
package tool

import (
	"context"
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
	"golang.org/x/mod/modfile"
)

// Exec checks if the tool is installed and runs it with the given arguments.
// If tool isn't installed, it will be installed first.
func Exec(_ context.Context, name string, args ...string) error {
	if err := VerifyInstallation(name); err != nil {
		return fmt.Errorf("verify installation: %w", err)
	}
	args = append([]string{"tool", name}, args...)
	return sh.RunV("go", args...)
}

// VerifyInstallation checks if the tool is installed and installs it if not.
func VerifyInstallation(name string) error {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return fmt.Errorf("read go.mod: %w", err)
	}

	f, err := modfile.ParseLax("go.mod", data, nil)
	if err != nil {
		return fmt.Errorf("parse go.mod: %w", err)
	}

	// Check if the tool is in the go.mod file
	for _, t := range f.Tool {
		if t.Path == name {
			return nil
		}
	}

	// If not, install the tool
	return sh.RunV("go", "get", "-tool", name+"@latest")
}
