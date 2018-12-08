// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package main

import (
	"fmt"
	"strings"
)

func printStats(components []*LSTComponent) {
	countTotalPackages := 0
	countNoPackageName := 0
	countNoVersion := 0
	countNoProjectDownloadURL := 0
	countInvalidaProjectDownloadURL := 0
	countNoProjectURL := 0
	countInvalidaProjectURL := 0
	countNoPackageVersionDownloadLink := 0
	countInvalidaPackageVersionDownloadLink := 0
	countNoDescription := 0

	for _, component := range components {
		countTotalPackages++

		if component.PackageName == "" {
			countNoPackageName++
		}

		if component.Version == "" {
			countNoVersion++
			fmt.Printf("===> %s: %s\n", component.PackageName, component.Version)
		}

		if component.ProjectDownloadURL == "" {
			countNoProjectDownloadURL++
		} else if !strings.HasPrefix(component.ProjectDownloadURL, "http") {
			countInvalidaProjectDownloadURL++
		}

		if component.ProjectURL == "" {
			countNoProjectURL++
		} else if !strings.HasPrefix(component.ProjectURL, "http") {
			countInvalidaProjectURL++
		}

		if component.PackageVersionDownloadLink == "" {
			countNoPackageVersionDownloadLink++
		} else if !strings.HasPrefix(component.PackageVersionDownloadLink, "http") {
			countInvalidaPackageVersionDownloadLink++
		}

		if component.Description == "" {
			countNoDescription++
		}
	}

	// now, print the results
	fmt.Printf("\n")
	fmt.Printf("Total packages: %d\n", countTotalPackages)
	fmt.Printf("No Package Name: %d\n", countNoPackageName)
	fmt.Printf("No Version: %d\n", countNoVersion)
	fmt.Printf("No Project Download URL: %d\n", countNoProjectDownloadURL)
	fmt.Printf("Invalid Project Download URL: %d\n", countInvalidaProjectDownloadURL)
	fmt.Printf("No Project URL: %d\n", countNoProjectURL)
	fmt.Printf("Invalid Project URL: %d\n", countInvalidaProjectURL)
	fmt.Printf("No Package Version Download Link: %d\n", countNoPackageVersionDownloadLink)
	fmt.Printf("Invalid Package Version Download Link: %d\n", countInvalidaPackageVersionDownloadLink)
	fmt.Printf("No Description: %d\n", countNoDescription)
}
