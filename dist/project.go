package dist

import (
	"github.com/JaSei/pathutil-go"
	"github.com/pkg/errors"
)

type Project interface {
	Path() pathutil.Path
	GitInit() error
	MakeGitIgnore() error
	MakeVersionFile() error
}

// implementation
type project struct {
	path pathutil.Path
}

func NewProject(projectName string) (Project, error) {
	if err := CheckGoPath(); err != nil {
		return nil, errors.Wrapf(err, "NewProject(%s) fail", projectName)
	}

	projectPath, err := ProjectPath(projectName)
	if err != nil {
		return nil, errors.Wrapf(err, "NewProject(%s) fail", projectName)
	}

	if projectPath.Exists() {
		return nil, errors.Errorf("Project dir %s already exists", projectPath)
	}

	if err = projectPath.MakePath(); err != nil {
		return nil, errors.Wrapf(err, "NewProject(%s) fail")
	}

	projectPath.Chdir()

	return project{path: projectPath}, nil
}

func (project project) Path() pathutil.Path {
	return project.path
}
