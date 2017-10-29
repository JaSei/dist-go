package dist

import (
	"strings"

	"github.com/JaSei/pathutil-go"
	"github.com/pkg/errors"
)

// NewLib create new project
func NewLib(projectName string) (err error) {
	if err = CheckGoPath(); err != nil {
		return
	}

	project, err := NewProject(projectName)
	if err != nil {
		return errors.Wrap(err, "NewLib failed")
	}

	// catch panic, remove uncomplete dir and set error
	defer func() {
		if r := recover(); r != nil {
			if err = project.Path().RemoveTree(); err != nil {
				err = errors.Wrapf(err, "Remove uncomplete path %s fail", project.Path())
			} else {
				err = errors.Wrap(r.(error), "NewLib failed")
			}
		}
	}()

	if err := project.GitInit(); err != nil {
		panic(err)
	}
	if err := project.MakeGitIgnore(); err != nil {
		panic(err)
	}
	if err := project.MakeVersionFile(); err != nil {
		panic(err)
	}
	if err := project.MakeExampleLib(); err != nil {
		panic(err)
	}
	if err := project.MakeDepFiles(); err != nil {
		panic(err)
	}
	//project.GenerateDistConf()
	//project.GenerateTravis()
	//project.GenerateAppVeyor()
	//project.MakeLicenseFile()

	return nil
}

func (project project) GitInit() error {
	return Run("git", "init")
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
	pkg := project.gpp.Project()

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
	gopkg, err := pathutil.NewPath(project.Path().String(), "Gopkg.toml")
	if err != nil {
		return errors.Wrapf(err, "MakeDepFiles(Gopkg.toml)")
	}

	return errors.Wrap(gopkg.Spew(""), "MakeDepFiles(Gopkg.toml)")
}
