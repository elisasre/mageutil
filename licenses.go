package mageutil

import (
	"context"
	"encoding/csv"
	"flag"
	"io"
	"os"

	"github.com/google/go-licenses/licenses"
	"k8s.io/klog/v2"
)

const (
	UNKNOWN = "Unknown"
)

// LicenseCheck runs github.com/google/go-licenses/licenses for given targets
// and writes toe output into w.
func LicenseCheck(ctx context.Context, w io.Writer, targets ...string) error {
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
	Name        string
	LicenseURL  string
	LicenseName string
	LicensePath string
	Version     string
}

// LicenseText reads and returns the contents of LicensePath, if set
// or an empty string if not.
func (lib libraryData) LicenseText() (string, error) {
	if lib.LicensePath == "" {
		return "", nil
	}
	data, err := os.ReadFile(lib.LicensePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func reportMain(w io.Writer, args ...string) error {
	// Defaults from original source
	var (
		confidenceThreshold          = 0.9
		includeTests                 = false
		ignore              []string = nil
	)

	classifier, err := licenses.NewClassifier(confidenceThreshold)
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
			Name:        lib.Name(),
			Version:     version,
			LicenseURL:  UNKNOWN,
			LicenseName: UNKNOWN,
		}
		if lib.LicensePath != "" {
			libData.LicensePath = lib.LicensePath
			name, _, err := classifier.Identify(lib.LicensePath)
			if err == nil {
				libData.LicenseName = name
			} else {
				klog.Errorf("Error identifying license in %q: %v", lib.LicensePath, err)
			}
			url, err := lib.FileURL(context.Background(), lib.LicensePath)
			if err == nil {
				libData.LicenseURL = url
			} else {
				klog.Warningf("Error discovering license URL: %s", err)
			}
		}
		reportData = append(reportData, libData)
	}

	return reportCSV(reportData, w)
}

func reportCSV(libs []libraryData, w io.Writer) error {
	writer := csv.NewWriter(w)
	for _, lib := range libs {
		if err := writer.Write([]string{lib.Name, lib.LicenseURL, lib.LicenseName}); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}
