package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"text/template"

	"github.com/meian/go-convert/internal/components/name"
	"github.com/meian/go-convert/internal/components/output"
	"github.com/meian/go-convert/internal/components/param"
)

const toolName = "gen-const"

var (
	//go:embed source.tmpl
	tmpl string
)

type Param struct {
	*param.Base
	pkgName  string
	use64Flg bool
}

func main() {
	param.Args = os.Args
	p, err := parseParam()
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	tmpl, err := template.New("").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	tp := map[string]interface{}{
		"command":   p.Command(),
		"package":   p.pkgName,
		"use64":     p.use64Flg,
		"signeds":   name.Signeds,
		"unsigneds": name.Unsigneds,
	}
	if err := tmpl.Execute(buf, tp); err != nil {
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
	fs.StringVar(&p.pkgName, "p", "", "source package name")
	fs.BoolVar(&p.use64Flg, "use64flg", false, "add is64 definition and init")
	if err := fs.Parse(os.Args[1:]); err != nil {
		fs.Usage()
		return nil, err
	}
	re := regexp.MustCompile(`^\pL[_\pL\pN]*$`)
	if !re.MatchString(p.pkgName) {
		fs.Usage()
		return nil, fmt.Errorf(`invalid package name: "%s"`, p.pkgName)
	}
	return p, nil
}
