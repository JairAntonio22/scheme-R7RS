package read

import (
	"strings"
)

type ReadError struct {
	errs []error
}

func (e *ReadError) Error() string {
	var builder strings.Builder

	for i, err := range e.errs {
		if i > 0 {
			builder.WriteString("\n")
		}

		builder.WriteString(err.Error())
	}

	return builder.String()
}
