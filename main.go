/*
dist-go is tool to make it easier to write and publish new golang software

dist-go has a bold assumption and convention like the followings:
* Your project use [dep](https://github.com/golang/dep)
* Your project have VERSION file as source of version
* Your project have tests and is tested in travis (and appveyour)
* README.md is generated from godoc (no more doc duplicity/mismatch)
* Your app project is released to github releases
* vendor dir is gitignored (vendor isn't commited - https://github.com/golang/dep/blob/master/docs/FAQ.md#should-i-commit-my-vendor-directory)
* `-go` suffix in project name is ignored (package called without `-go` suffix)

SYNOPSIS

	dist-gp new lib github.com/JaSei/NewLibrary
	chdir $(dist-go pd github.com/JaSei/NewLibrary)
	# chdir $(dist-go pd NewLibrary)
	dist-go test
	dist-go release

	ls -al $(dist-go pd github.com/JaSei/test-go)
	rm -rf $(dist-go pd github.com/JaSei/test-go)


new

* check project dir
* cd to project dir
* git init

release

* check version
* check dirty state of git (goreleaser don't release in dirty state for example)

VERSION

it's really important have your code versioned

git tag is nice, but isn't good source of version

version should be part of pull request

In app exists concept of `Version` variable.
This variable is automatically? used in go build, or can be override by ld options.

But in library doesn't exists concepts of versions.

Someone use `version.go` as source of version, for example:
* https://github.com/nsqio/go-nsq/blob/master/version.go
* https://github.com/hashicorp/terraform/blob/master/version/version.go
* https://github.com/tcnksm/ghr/blob/master/version.go

Other possibilities is used `VERSION` file as Version source like:
* https://github.com/ahmetb/govvv

Other is `go generate` from tag:
* https://adrianhesketh.com/2016/09/04/adding-a-version-number-to-go-packages-with-go-generate/

I like `VERSION` file, bacause it's simple to do it, and works it well to app and lib variant.

*/
package main

import (
	"log"

	"github.com/JaSei/dist-go/dist"
	"github.com/alecthomas/kingpin"
)

var (
	new               = kingpin.Command("new", "create new repo")
	newLib            = new.Command("lib", "create new library repo")
	newLibProjectName = newLib.Arg("project", "project name (eg. github.com/JaSei/test-go").Required().String()
	newApp            = new.Command("app", "create new application repo")

	pd            = kingpin.Command("projectDir", "get project directory").Alias("pd")
	pdProjectName = pd.Arg("project", "project name (eg. github.com/JaSei/test-go or JaSei/test-go or test-go)").Required().String()

	test = kingpin.Command("test", "run test(s)")

	release = kingpin.Command("release", "release new version")
)

func main() {
	switch kingpin.Parse() {
	case newLib.FullCommand():
		handleError(dist.NewLib(*newLibProjectName))
	case newApp.FullCommand():
		handleError(dist.NewApp())
	case pd.FullCommand():
		handleError(dist.PrintProjectDirectory(*pdProjectName))
	case test.FullCommand():
		handleError(dist.Test())
	case release.FullCommand():
		handleError(dist.Release())
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
