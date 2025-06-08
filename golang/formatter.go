package golang

import (
	"go/format"
)

// FormatNode formats the given Go code.
// It uses go/format to format the code.
func FormatNode(code string) (string, error) {
	formattedCode, err := format.Source([]byte(code))
	if err != nil {
		return "", err
	}
	return string(formattedCode), nil
}
