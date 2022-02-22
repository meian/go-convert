package param

import (
	"flag"
	"strings"
)

var Args []string

type Base struct {
	OutFile string
	FlagSet *flag.FlagSet
}

func New(name string) *Base {
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	b := &Base{
		FlagSet: fs,
	}
	fs.StringVar(&b.OutFile, "o", "", "file name to write output")
	return b
}

func (b *Base) Name() string {
	return b.FlagSet.Name()
}

func (b *Base) Command() string {
	return strings.Join(append([]string{b.Name()}, Args[1:]...), " ")
}
