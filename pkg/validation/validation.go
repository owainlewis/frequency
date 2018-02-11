package validation

import (
	"strings"
)

func ConsolidateErrors(errs []error) string {
	var errorMessages []string
	for _, err := range errs {
		errorMessages = append(errorMessages, err.Error())
	}

	return strings.Join(errorMessages, ", ")
}
