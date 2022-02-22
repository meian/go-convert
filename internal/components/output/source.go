package output

import (
	"fmt"
	"go/format"
	"os"
)

func Source(filename string, source []byte) error {
	out := os.Stdout
	if len(filename) > 0 {
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("cannot create file: name=%s, %w", filename, err)
		}
		out = f
	}
	fmtd, err := format.Source(source)
	if err != nil {
		return fmt.Errorf("cannot format source: %w\n\n%s", err, string(source))
	}
	_, err = out.Write(fmtd)
	if err != nil {
		return fmt.Errorf("cannot write source: %w", err)
	}
	return nil
}
