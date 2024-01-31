package util

import (
	"strings"
)

// ParseAddress parse the address according to a finite length.
func ParseAddress(address string, length int) (address1 string, address2 string, address3 string) {
	address1 = ""
	address2 = ""
	address3 = ""
	smash := strings.Fields(address)
	for _, value := range smash {
		if len(address1)+len(value) < length {
			address1 += value + " "
		} else if len(address2)+len(value) < length {
			address2 += value + " "
		} else if len(address3)+len(value) < length {
			address3 += value + " "
		}
	}
	address1 = strings.TrimSpace(address1)
	address2 = strings.TrimSpace(address2)
	address3 = strings.TrimSpace(address3)
	return
}
