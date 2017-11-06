/*
gopackagepath is abstraction of (full) package name

(full) package name contains 3 part of name:

	repository/user/package

e.g.

	github.com/JaSei/pathutil-go

more about package convention in [package names article](https://blog.golang.org/package-names)
*/
package gopackagepath

import (
	"strings"

	"github.com/pkg/errors"
)

type GoPackagePath interface {
	Repo() string
	User() string
	Package() string
	SubPackage() string
	UserPackage() string
}

type gopackagepath struct {
	repo        string
	user        string
	packageName string
	subPackage  string
}

// New create new instance of `GoPackagePath` interface
func New(fullPackageName string) (GoPackagePath, error) {
	packageParts := strings.SplitN(fullPackageName, "/", 4)

	if len(packageParts) < 3 {
		return nil, errors.New("Invalid package name")
	}

	newPackageName := gopackagepath{
		repo:        packageParts[0],
		user:        packageParts[1],
		packageName: packageParts[2],
	}

	if len(packageParts) == 4 {
		newPackageName.subPackage = packageParts[3]
	}

	return newPackageName, nil
}

// Repo return repo (first) part of full package name
func (pkg gopackagepath) Repo() string {
	return pkg.repo
}

// User return user (second) part of full package name
func (pkg gopackagepath) User() string {
	return pkg.user
}

// User return package (third) part of full package name
func (pkg gopackagepath) Package() string {
	return pkg.packageName
}

// SubPackage return next levels (last) part of full package name
func (pkg gopackagepath) SubPackage() string {
	return pkg.subPackage
}

// UserPackage return user + package - should be be relative uniqiue identificator
func (pkg gopackagepath) UserPackage() string {
	return pkg.user + "/" + pkg.packageName
}
