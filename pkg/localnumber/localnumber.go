package localnumber

import (
	"strings"

	"github.com/nyaruka/phonenumbers"
)

func Format(input string) (output string, err error) {
	number, err := phonenumbers.Parse(input, "GB")
	if err != nil {
		return
	}
	output = phonenumbers.Format(number, phonenumbers.NATIONAL)
	output = strings.ReplaceAll(output, " ", "")
	if number.GetCountryCode() == 91 && len(output) > 10 && output[0:1] == "0" {
		output = output[1:]
	}
	return
}
