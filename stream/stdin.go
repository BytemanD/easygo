package stream

import (
	"io"
	"os"
)

func ReadStringFromStdin() (string, error) {
	if bytes, err := io.ReadAll(os.Stdin); err == nil {
		return string(bytes), nil
	} else {
		return "", err
	}
}
