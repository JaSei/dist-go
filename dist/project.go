package dist

import (
	"strings"

	"github.com/JaSei/dist-go/goprojectpath"
	"github.com/JaSei/pathutil-go"
	"github.com/pkg/errors"
)

type Project interface {
	Path() pathutil.Path
	GitInit() error
	MakeGitIgnore() error
	MakeVersionFile() error
	MakeExampleLib() error
	MakeDepFiles() error
	CheckIntegrity() error
}

// implementation
type project struct {
	path pathutil.Path
	gpp  goprojectpath.GoProjectPath
}

func NewProject(projectName string) (Project, error) {
	project, err := newProject(projectName)
	if err != nil {
		return nil, errors.Wrap(err, "NewProject")
	}

	if project.Path().Exists() {
		return nil, errors.Errorf("Project dir %s already exists", project.Path())
	}

	if err = project.Path().MakePath(); err != nil {
		return nil, errors.Wrapf(err, "NewProject(%s) fail")
	}

	project.Path().Chdir()

	return project, nil
}

// try load project in current working directory
func LoadCwdProject() (Project, error) {
	cwd, err := pathutil.Cwd()
	if err != nil {
		return nil, errors.Wrap(err, "LoadCwdProject")
	}

	goSrc, err := GoSrcPath()
	if err != nil {
		return nil, errors.Wrap(err, "LoadCwdProject")
	}

	if !strings.HasPrefix(cwd.String(), goSrc.String()) {
		return nil, errors.Errorf("You arn't in GOPATH/src directory")
	}

	return LoadProject(strings.TrimPrefix(cwd.String(), goSrc.String()))
}

func LoadProject(projectName string) (Project, error) {
	project, err := newProject(projectName)
	if err != nil {
		return nil, errors.Wrap(err, "LoadProject")
	}

	if err = project.CheckIntegrity(); err != nil {
		return nil, errors.Wrap(err, "LoadProject")
	}

	return project, nil
}

func newProject(projectName string) (Project, error) {
	gpp, err := goprojectpath.New(projectName)
	if err != nil {
		return nil, err
	}

	if err := CheckGoPath(); err != nil {
		return nil, errors.Wrapf(err, "Project(%s) fail", projectName)
	}

	projectPath, err := ProjectPath(projectName)
	if err != nil {
		return nil, errors.Wrapf(err, "Project(%s) fail", projectName)
	}

	return project{path: projectPath, gpp: gpp}, nil
}

func (project project) Path() pathutil.Path {
	return project.path
}

func (project project) CheckIntegrity() error {
	shouldBeExistsPaths := []pathutil.Path{
		project.gitignorePath(),
		project.versionFilePath(),
	}

	for _, path := range shouldBeExistsPaths {
		if !path.Exists() {
			return errors.Errorf("No exists %s", path)
		}
	}

	return nil
}

func (project project) gitignorePath() pathutil.Path {
	// project.Path() must be defined, error could be ignored
	path, _ := pathutil.NewPath(project.Path().String(), ".gitignore")
	return path
}

func (project project) versionFilePath() pathutil.Path {
	// project.Path() must be defined, error could be ignored
	path, _ := pathutil.NewPath(project.Path().String(), "VERSION")
	return path
}
