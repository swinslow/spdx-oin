// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/spdx/tools-golang/spdx"
)

func createPackage(component *LSTComponent) (*spdx.Package2_2, error) {
	// fill in initial mandatory values
	// for now we'll leave empty any mandatory values that require more processing
	pkg := spdx.Package2_2{
		PackageName:           component.PackageName,
		PackageSPDXIdentifier: spdx.ElementID(fmt.Sprintf("%d", component.SerialNumber)),
		// fill in PackageDownloadLocation below
		FilesAnalyzed:             false,
		IsFilesAnalyzedTagPresent: true,
		PackageLicenseConcluded:   "NOASSERTION",
		PackageLicenseDeclared:    "NOASSERTION",
		PackageCopyrightText:      "NOASSERTION",
	}

	// fill in data as appropriate, if it is present.
	// where not present, note in package comment string.
	var pkgComment strings.Builder

	prefixes := []string{
		"http",
		"ftp",
		"www",
	}

	if component.Version == "" {
		pkgComment.WriteString("Version field was empty.\n")
	} else {
		pkg.PackageVersion = component.Version
	}

	if component.ProjectDownloadURL == "" {
		pkg.PackageSourceInfo = "Project Download URL field was empty"
	} else if !hasAnyPrefix(component.ProjectDownloadURL, prefixes) {
		pkg.PackageSourceInfo = fmt.Sprintf("Project Download URL field was invalid format: %s", component.ProjectDownloadURL)
	} else {
		pkg.PackageSourceInfo = fmt.Sprintf("Project Download URL was %s", component.ProjectDownloadURL)
	}

	if component.ProjectURL == "" {
		pkgComment.WriteString("Project URL field was empty.\n")
		pkg.PackageHomePage = "NONE"
	} else if !hasAnyPrefix(component.ProjectURL, prefixes) {
		pkgComment.WriteString(fmt.Sprintf("Project URL field was invalid format: %s\n", component.ProjectURL))
		pkg.PackageHomePage = "NOASSERTION"
	} else {
		pkg.PackageHomePage = component.ProjectURL
	}

	if component.PackageVersionDownloadLink == "" {
		pkgComment.WriteString("Package Version Download Link field was empty.\n")
		pkg.PackageDownloadLocation = "NONE"
	} else if !hasAnyPrefix(component.PackageVersionDownloadLink, prefixes) {
		pkgComment.WriteString(fmt.Sprintf("Package Version Download Link field was invalid format: %s\n", component.PackageVersionDownloadLink))
		pkg.PackageDownloadLocation = "NOASSERTION"
	} else {
		pkg.PackageDownloadLocation = component.PackageVersionDownloadLink
	}

	if component.Description == "" {
		pkgComment.WriteString("Description field was empty.\n")
	} else {
		pkg.PackageDescription = component.Description
	}

	// finally, add comments if any
	if pkgComment.Len() > 0 {
		pkg.PackageComment = pkgComment.String()
	}

	return &pkg, nil
}

func createDocument(components []*LSTComponent, tableNumber string) (*spdx.Document2_2, error) {
	// build main document
	doc := spdx.Document2_2{}

	// build creation info section
	doc.CreationInfo = &spdx.CreationInfo2_2{
		SPDXVersion:       "SPDX-2.2",
		DataLicense:       "CC0-1.0",
		SPDXIdentifier:    spdx.ElementID("DOCUMENT"),
		DocumentName:      fmt.Sprintf("Linux System Table %s", tableNumber),
		DocumentNamespace: fmt.Sprintf("https://github.com/swinslow/spdx-oin/spdxdocs/table-%s.spdx", tableNumber),
		CreatorTools:      []string{"github.com/swinslow/spdx-oin-0.0.2"},
		Created:           time.Now().Format("2006-01-02T15:04:05Z"),
		DocumentComment:   fmt.Sprintf("Automatically generated from parsing HTML for Linux System Table %s from Open Invention Network website.\nNo attempt has been made to analyze the files, licenses or copyright statements for these packages.", tableNumber),
	}

	// create and add packages from components
	doc.Packages = map[spdx.ElementID]*spdx.Package2_2{}
	for _, component := range components {
		pkg, err := createPackage(component)
		if err != nil {
			return nil, err
		}
		doc.Packages[pkg.PackageSPDXIdentifier] = pkg
	}

	return &doc, nil
}
