# dist-go

[![Release](https://img.shields.io/github/release/github.com/JaSei.svg?style=flat-square)](https:///github.com/JaSei/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Travis](https://img.shields.io/travis/github.com/JaSei.svg?style=flat-square)](https://travis-ci.org/github.com/JaSei)
[![Go Report Card](https://goreportcard.com/badge//github.com/JaSei?style=flat-square)](https://goreportcard.com/report//github.com/JaSei)
[![GoDoc](https://godoc.org//github.com/JaSei?status.svg&style=flat-square)](http://godoc.org//github.com/JaSei)
[![codecov.io](https://codecov.io//github.com/JaSei/coverage.svg?branch=master)](https://codecov.io//github.com/JaSei?branch=master)
[![Sourcegraph](https://sourcegraph.com//github.com/JaSei/-/badge.svg)](https://sourcegraph.com//github.com/JaSei?badge)

dist-go is tool to make it easier to write and publish new golang software

dist-go has a bold assumption and convention like the followings: * Your project
use [dep](https://github.com/golang/dep) * Your project have VERSION file as
source of version * Your project have tests and is tested in travis (and
appveyour) * README.md is generated from godoc (no more doc duplicity/mismatch)
* Your app project is released to github releases * VCS are git only (yet) *
vendor dir is gitignored (vendor isn't commited -
https://github.com/golang/dep/blob/master/docs/FAQ.md#should-i-commit-my-vendor-directory)
* `-go` suffix in project name is ignored (package called without `-go` suffix)
* run test with coverege (https://github.com/pierrre/gotestcover#deprecated)

### SYNOPSIS

    dist-gp new lib github.com/JaSei/NewLibrary
    chdir $(dist-go pd github.com/JaSei/NewLibrary)
    # chdir $(dist-go pd NewLibrary)
    dist-go test
    dist-go release

    ls -al $(dist-go pd github.com/JaSei/test-go)
    rm -rf $(dist-go pd github.com/JaSei/test-go)

### new

* check project dir * cd to project dir * git init

### test

* generate README.md * run go generate * run tests

### release

* check version * check dirty state of git (goreleaser don't release in dirty
state for example)


### VERSION

it's really important have your code versioned

git tag is nice, but isn't good source of version

version should be part of pull request

In app exists concept of `Version` variable. This variable is automatically?
used in go build, or can be override by ld options.

But in library doesn't exists concepts of versions.

Someone use `version.go` as source of version, for example: *
https://github.com/nsqio/go-nsq/blob/master/version.go *
https://github.com/hashicorp/terraform/blob/master/version/version.go *
https://github.com/tcnksm/ghr/blob/master/version.go

Other possibilities is used `VERSION` file as Version source like: *
https://github.com/ahmetb/govvv

Other is `go generate` from tag: *
https://adrianhesketh.com/2016/09/04/adding-a-version-number-to-go-packages-with-go-generate/

I like `VERSION` file, bacause it's simple to do it, and works it well to app
and lib variant.

## Usage


## Help

HELP_PLACEHOLDER

## Contributing

Contributions are very much welcome.

### Makefile

Makefile provides several handy rules, like README.md `generator` , `setup` for prepare build/dev environment, `test`, `cover`, etc...

Try `make help` for more information.

### Before pull request

please try:
* run tests (`make test`)
* run linter (`make lint`)
* if your IDE don't automaticaly do `make lint`, run `go fmt` (`make fmt`)



### README

README.md are generate from template [.godocdown.tmpl](.godocdown.tmpl) and code documentation via [godocdown](https://github.com/robertkrimen/godocdown).

Never edit README.md direct, because your change will be lost.
.com/robertkrimen/godocdown).

Never edit README.md direct, because your change will be lost.
