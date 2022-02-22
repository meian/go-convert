package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	args := append([]string{}, os.Args...)
	defer func() {
		os.Args = args
	}()
	os.Args = []string{"app", "-t", "uint32"}
	main()
}
