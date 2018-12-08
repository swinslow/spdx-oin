// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/swinslow/spdx-go/v0/tvsaver"
)

// LSTComponent represents a single component on the Linux System Table.
type LSTComponent struct {
	TableNumber                string
	SerialNumber               int
	PackageName                string
	Version                    string
	ProjectDownloadURL         string
	ProjectURL                 string
	PackageVersionDownloadLink string
	Description                string
}

func loadRawHTML(filename string) (string, error) {
	r, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("Error while opening %v for reading: %v", filename, err)
	}
	defer r.Close()

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("Error while reading %v: %v", filename, err)
	}

	return string(bytes), nil
}

func main() {
	tables := map[string]string{
		"1-2": "scratch/table-1-2.html",
		"3":   "scratch/table-3.html",
		"4":   "scratch/table-4.html",
		"5":   "scratch/table-5.html",
		"6":   "scratch/table-6.html",
		"7":   "scratch/table-7.html",
		"8":   "scratch/table-8.html",
		"9":   "scratch/table-9.html",
	}

	for tableNum, htmlPath := range tables {
		s, err := loadRawHTML(htmlPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		components, err := parseRawHTML(s, tableNum)
		if err != nil {
			fmt.Println(err)
			return
		}

		doc, err := createDocument(components, tableNum)
		if err != nil {
			fmt.Println(err)
			return
		}

		// create a new file for writing
		fileOut := fmt.Sprintf("spdxdocs/table-%s.spdx", tableNum)
		w, err := os.Create(fileOut)
		if err != nil {
			fmt.Printf("Error while opening %v for writing: %v", fileOut, err)
			return
		}
		defer w.Close()

		// try to save the document to disk as a tag-value file
		err = tvsaver.Save2_1(doc, w)
		if err != nil {
			fmt.Printf("Error while saving %v: %v", fileOut, err)
			return
		}

		// it worked
		fmt.Printf("Successfully saved %s\n", fileOut)
	}

}
