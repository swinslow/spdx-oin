// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package main

import "strings"

func hasAnyPrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}

	return false
}
