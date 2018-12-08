// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package main

import (
	"fmt"
	"strconv"

	"github.com/anaskhan96/soup"
)

func parseRow(trElt soup.Root) (*LSTComponent, error) {
	tdElts := trElt.FindAll("td")
	if len(tdElts) < 1 {
		return nil, fmt.Errorf("Error while parsing HTML, no columns found in table row: %#v", trElt)
	}

	// build and fill in new component
	component := LSTComponent{}
	sn, err := strconv.Atoi(tdElts[0].Text())
	if err != nil {
		return nil, fmt.Errorf("Error while parsing HTML, could not parse serial number %s as integer", tdElts[0].Text())
	}
	component.SerialNumber = sn

	component.PackageName = tdElts[1].Text()
	component.Version = tdElts[2].Text()
	component.ProjectDownloadURL = tdElts[3].Find("a").Text()
	component.ProjectURL = tdElts[4].Find("a").Text()
	component.PackageVersionDownloadLink = tdElts[5].Find("a").Text()
	component.Description = tdElts[6].Text()

	return &component, nil
}

func parseRawHTML(raw string, tableNumber int) ([]*LSTComponent, error) {
	// build the parsed soup
	doc := soup.HTMLParse(raw)

	// get pointers to all <tr> elements
	trElts := doc.FindAll("tr")
	if len(trElts) < 1 {
		return nil, fmt.Errorf("Error while parsing HTML for table %d, no table rows found", tableNumber)
	}

	// the first <tr> should be the header row
	// let's make sure it really is -- check the text of the first column
	trHeader := trElts[0]
	tdHeader := trHeader.Find("td")
	if tdHeader.Text() != "S.no" {
		return nil, fmt.Errorf("Error while parsing HTML for table %d, expected header row text to be 'S.no', got %s", tableNumber, tdHeader.Text())
	}

	// We're good, so start parsing remainder and build components list
	components := []*LSTComponent{}
	for _, trElt := range trElts[1:] {
		component, err := parseRow(trElt)
		if err != nil {
			return nil, err
		}

		component.TableNumber = tableNumber
		components = append(components, component)
	}

	// and we're done!
	return components, nil
}
