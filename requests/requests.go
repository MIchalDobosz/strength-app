package requests

import (
	"encoding/json"
	"io"
)

func Read(reader io.Reader, dest any) error {
	return json.NewDecoder(reader).Decode(&dest)
}

func validateRequired[T comparable](name string, value T, valid *bool, errors map[string]string) {
	if value == *new(T) {
		errors[name] = name + " is required."
		*valid = false
	}
}
