package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/meian/go-convert/internal/components/name"
	"github.com/meian/go-convert/internal/components/output"
	"github.com/meian/go-convert/internal/components/param"
)

const toolName = "gen-int"

var (
	//go:embed tmpls
	tmplRoot embed.FS
)

type Param struct {
	*param.Base
	target string
	name   name.Name
}

func main() {
	param.Args = os.Args
	p, err := parseParam()
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	tmpls := []string{
		"tmpls/file.tmpl",
		"tmpls/int/value.tmpl",
		"tmpls/int/pointer.tmpl",
	}
	if !p.name.IsSigned() {
		for i := range tmpls {
			tmpls[i] = strings.Replace(tmpls[i], "int", "uint", 1)
		}
	}
	template, err := template.New("file.tmpl").ParseFS(tmplRoot, tmpls...)
	if err != nil {
		panic(err)
	}

	tp := map[string]interface{}{
		"command":   p.Command(),
		"package":   "convert",
		"target":    p.name,
		"signeds":   name.Signeds,
		"unsigneds": name.Unsigneds,
	}
	if err := template.Execute(buf, tp); err != nil {
		panic(err)
	}

	if err := output.Source(p.OutFile, buf.Bytes()); err != nil {
		panic(err)
	}
}

func parseParam() (*Param, error) {
	p := &Param{
		Base: param.New(toolName),
	}
	fs := p.FlagSet
	fs.StringVar(&p.target, "t", "", "target type")
	if err := fs.Parse(os.Args[1:]); err != nil {
		fs.Usage()
		return nil, err
	}
	n, ok := name.ByName(p.target)
	if !ok {
		fs.Usage()
		return nil, fmt.Errorf("invalid type: type=%s", p.target)
	}
	p.name = n
	return p, nil
}
