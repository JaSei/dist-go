package dist

import (
	"github.com/JaSei/pathutil-go"
	"github.com/pkg/errors"
)

// NewLib create new project
func NewLib(projectName string) error {
	if err := CheckGoPath(); err != nil {
		return err
	}

	project, err := NewProject(projectName)
	if err != nil {
		return errors.Wrap(err, "NewLib failed")
	}

	if err := project.GitInit(); err != nil {
		return errors.Wrap(err, "NewLib failed")
	}
	if err := project.MakeGitIgnore(); err != nil {
		return errors.Wrap(err, "NewLib failed")
	}
	if err := project.MakeVersionFile(); err != nil {
		return errors.Wrap(err, "NewLib failed")
	}
	//project.GenerateExampleLib()
	//project.GenerateDepFiles()
	//project.GenerateDistConf()
	//project.GenerateTravis()
	//project.GenerateAppVeyor()

	return nil
}

func (project project) GitInit() error {
	return Run("git", "init")
}

func (project project) MakeGitIgnore() error {
	gitignore, err := pathutil.NewPath(project.Path().String(), ".gitignore")
	if err != nil {
		return errors.Wrap(err, "NewPath(.gitignore)")
	}

	if err := gitignore.Spew(`vendor/
`); err != nil {
		return errors.Wrap(err, "Spew(.gitignore)")
	}

	return nil
}

func (project project) MakeVersionFile() error {
	gitignore, err := pathutil.NewPath(project.Path().String(), "VERSION")
	if err != nil {
		return errors.Wrap(err, "MakeVersionFile(VERSION)")
	}

	if err := gitignore.Spew("0.1.0"); err != nil {
		return errors.Wrap(err, "MakeVersionFile(VERSION)")
	}

	return nil
}
