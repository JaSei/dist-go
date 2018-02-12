package project

import (
	"os/exec"

	"github.com/JaSei/dist-go/gotool"
	"github.com/pkg/errors"
)

// GoTestCover method do test and calculate coverage
// (ignore vendor file) to coverage.txt file
func (project project) GoTestCover() error {
	coverageTxtPath, _ := project.Path().Child("coverage.txt")

	return gotool.GoTestCover(coverageTxtPath)
}

func (project project) GenerateReadme() error {
	//.godocdown.tmpl
	godocdown := `
# {{ .Name }}

[![Release](https://img.shields.io/github/release/` + project.gpp.UserPackage() + `.svg?style=flat-square)](https://` + project.gpp.FullPackage() + `/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Travis](https://img.shields.io/travis/` + project.gpp.UserPackage() + `.svg?style=flat-square)](https://travis-ci.org/` + project.gpp.UserPackage() + `)
[![Go Report Card](https://goreportcard.com/badge/` + project.gpp.FullPackage() + `?style=flat-square)](https://goreportcard.com/report/` + project.gpp.FullPackage() + `)
[![GoDoc](https://godoc.org/` + project.gpp.FullPackage() + `?status.svg&style=flat-square)](http://godoc.org/` + project.gpp.FullPackage() + `)
[![codecov.io](https://codecov.io/` + project.gpp.FullPackage() + `/coverage.svg?branch=master)](https://codecov.io/` + project.gpp.FullPackage() + `?branch=master)
[![Sourcegraph](https://sourcegraph.com/` + project.gpp.FullPackage() + `/-/badge.svg)](https://sourcegraph.com/` + project.gpp.FullPackage() + `?badge)

{{ .EmitSynopsis }}

{{ .EmitUsage }}


## Help

HELP_PLACEHOLDER

## Contributing

Contributions are very much welcome.

### Makefile

Makefile provides several handy rules, like README.md ` + "`" + `generator` + "`" + ` , ` + "`" + `setup` + "`" + ` for prepare build/dev environment, ` + "`" + `test` + "`" + `, ` + "`" + `cover` + "`" + `, etc...

Try ` + "`" + `make help` + "`" + ` for more information.

### Before pull request

please try:
* run tests (` + "`" + `make test` + "`" + `)
* run linter (` + "`" + `make lint` + "`" + `)
* if your IDE don't automaticaly do ` + "`" + `make lint` + "`" + `, run ` + "`" + `go fmt` + "`" + ` (` + "`" + `make fmt` + "`" + `)



### README

README.md are generate from template [.godocdown.tmpl](.godocdown.tmpl) and code documentation via [godocdown](https://github.com/robertkrimen/godocdown).

Never edit README.md direct, because your change will be lost.
`

	if err := project.goDocDownTmplPath().Spew(godocdown); err != nil {
		return errors.Wrapf(err, "write to .godocdown.tmpl fail")
	}

	out, err := exec.Command("godocdown").Output()
	if err != nil {
		return errors.Wrapf(err, "godocdown utility fail")
	}

	return project.readmePath().SpewBytes(out)
}
