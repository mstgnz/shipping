package cargo

import (
	"strings"
)

func ParseAddress(address string) (address1 string, address2 string, address3 string) {
	address1 = ""
	address2 = ""
	address3 = ""

	smash := strings.Fields(address)
	for _, value := range smash {
		if len(address1)+len(value) < 30 {
			address1 += value + " "
		} else if len(address2)+len(value) < 30 {
			address2 += value + " "
		} else if len(address3)+len(value) < 30 {
			address3 += value + " "
		}
	}
	address1 = strings.TrimSpace(address1)
	address2 = strings.TrimSpace(address2)
	address3 = strings.TrimSpace(address3)
	return
}
