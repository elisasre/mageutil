// Package golicenses exposes github.com/google/go-licenses/licenses as library.
package golicenses

import (
	"context"
	"encoding/csv"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/google/go-licenses/licenses"
	"k8s.io/klog/v2"
)

const UNKNOWN = "Unknown"

// Run executes github.com/google/go-licenses/licenses for given targets and writes output into w.
func Run(ctx context.Context, w io.Writer, targets ...string) error {
	// Change klog default log level to INFO.
	klog.InitFlags(nil)
	err := flag.Set("logtostderr", "true")
	if err != nil {
		return err
	}
	err = flag.Set("stderrthreshold", "INFO")
	if err != nil {
		return err
	}

	return reportMain(w, targets...)
}

type libraryData struct {
	Name         string
	LicenseURLs  []string
	LicenseNames []string
	LicenseFile  string
	Version      string
}

// LicenseText reads and returns the contents of LicenseFile, if set or an empty string if not.
func (lib libraryData) LicenseText() (string, error) {
	if lib.LicenseFile == "" {
		return "", nil
	}
	data, err := os.ReadFile(lib.LicenseFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func reportMain(w io.Writer, args ...string) error {
	// Defaults from original source
	var (
		includeTests          = false
		ignore       []string = nil
	)

	classifier, err := licenses.NewClassifier()
	if err != nil {
		return err
	}

	libs, err := licenses.Libraries(context.Background(), classifier, includeTests, ignore, args...)
	if err != nil {
		return err
	}

	reportData := make([]libraryData, 0, len(libs))
	for _, lib := range libs {
		version := lib.Version()
		if len(version) == 0 {
			version = UNKNOWN
		}
		libData := libraryData{
			Name:         lib.Name(),
			Version:      version,
			LicenseURLs:  []string{},
			LicenseNames: []string{},
		}
		if lib.LicenseFile != "" {
			libData.LicenseFile = lib.LicenseFile
			licenses, err := classifier.Identify(lib.LicenseFile)
			for _, license := range licenses {
				if err == nil {
					libData.LicenseNames = append(libData.LicenseNames, license.Name)
				} else {
					klog.Errorf("Error identifying license in %q: %v", lib.LicenseFile, err)
				}
				url, err := lib.FileURL(context.Background(), lib.LicenseFile)
				if err == nil {
					libData.LicenseURLs = append(libData.LicenseURLs, url)
				} else {
					klog.Warningf("Error discovering license URL: %s", err)
				}
			}
		}
		reportData = append(reportData, libData)
	}

	return reportCSV(reportData, w)
}

func reportCSV(libs []libraryData, w io.Writer) error {
	writer := csv.NewWriter(w)
	for _, lib := range libs {
		if err := writer.Write([]string{lib.Name, strings.Join(lib.LicenseURLs, ", "), strings.Join(lib.LicenseNames, ", ")}); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}
