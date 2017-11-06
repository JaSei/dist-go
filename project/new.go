package project

import (
	"strconv"
	"strings"
	"time"

	"github.com/JaSei/pathutil-go"
	"github.com/Masterminds/vcs"
	"github.com/nishanths/license/pkg/license"
	"github.com/pkg/errors"
)

func (project project) VCSInit() error {
	projectVcs, err := vcs.NewGitRepo(project.Path().Canonpath(), project.Path().Canonpath())
	if err != nil {
		return err
	}

	return projectVcs.Init()
}

func (project project) MakeGitIgnore() error {
	gitignore := project.gitignorePath()
	if err := gitignore.Spew(`vendor/
`); err != nil {
		return errors.Wrap(err, "Spew(.gitignore)")
	}

	return nil
}

func (project project) MakeVersionFile() error {
	version := project.versionFilePath()
	if err := version.Spew("0.1.0"); err != nil {
		return errors.Wrap(err, "MakeVersionFile(VERSION)")
	}

	return nil
}

func (project project) MakeExampleLib() error {
	pkg := project.gpp.Package()

	if strings.HasSuffix(pkg, "-go") {
		pkg = pkg[0 : len(pkg)-3]
	}

	lib, err := pathutil.NewPath(project.Path().String(), pkg+".go")
	if err != nil {
		return errors.Wrapf(err, "MakeExampleLib(%s.go)", pkg)
	}

	if err = lib.Spew("package " + pkg); err != nil {
		return errors.Wrapf(err, "MakeExampleLib(%s.go)", pkg)
	}

	test, err := pathutil.NewPath(project.Path().String(), pkg+"_test.go")
	if err != nil {
		return errors.Wrapf(err, "MakeExampleLib(%s_test.go)", pkg)
	}

	if err = test.Spew("package " + pkg); err != nil {
		return errors.Wrapf(err, "MakeExampleLib(%s_test.go)", pkg)
	}

	return nil
}

func (project project) MakeDepFiles() error {
	gopkg := project.depGopkgTomlPath()

	return errors.Wrap(gopkg.Spew(""), "MakeDepFiles(Gopkg.toml)")
}

func (project project) MakeLicenseFile() error {
	client := license.NewClient()
	lic, err := client.Info(project.license)
	if err != nil {
		errors.Wrapf(err, "MakeLicenseFile(%s)", project.license)
	}

	currentTime := time.Now()
	licenseText := strings.Replace(lic.Body, "[year]", strconv.Itoa(currentTime.Year()), 1)
	licenseText = strings.Replace(licenseText, "[fullname]", project.author, 1)

	return project.licensePath().Spew(licenseText)
}
