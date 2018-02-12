package project

import (
	"strings"

	"github.com/JaSei/dist-go/gopackagepath"
	"github.com/JaSei/dist-go/utils"
	"github.com/JaSei/pathutil-go"
	"github.com/pkg/errors"
)

type Project interface {
	Path() pathutil.Path
	VCSInit() error
	MakeGitIgnore() error
	MakeVersionFile() error
	MakeExampleLib() error
	MakeDepFiles() error
	CheckIntegrity() error
	MakeLicenseFile() error
	SaveConfig() error
	LoadConfig() error
	GoTestCover() error
	GenerateReadme() error
}

// implementation
type project struct {
	path    pathutil.Path
	gpp     gopackagepath.GoPackagePath
	author  string
	license string
}

func New(projectName, author, license string) (Project, error) {
	project, err := newProject(projectName, author, license)
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

	goSrc, err := utils.GoSrcPath()
	if err != nil {
		return nil, errors.Wrap(err, "LoadCwdProject")
	}

	if !strings.HasPrefix(cwd.String(), goSrc.String()) {
		return nil, errors.Errorf("You arn't in GOPATH/src directory")
	}

	return LoadProject(strings.TrimPrefix(cwd.String(), goSrc.String()), "", "")
}

func LoadProject(projectName, author, license string) (Project, error) {
	project, err := newProject(projectName, author, license)
	if err != nil {
		return nil, errors.Wrap(err, "LoadProject")
	}

	if author == "" && license == "" {
		if err = project.LoadConfig(); err != nil {
			return nil, errors.Wrap(err, "LoadProject")
		}
	}

	if err = project.CheckIntegrity(); err != nil {
		return nil, errors.Wrap(err, "LoadProject")
	}

	return project, nil
}

func newProject(projectName, author, license string) (Project, error) {
	gpp, err := gopackagepath.New(projectName)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckGoPath(); err != nil {
		return nil, errors.Wrapf(err, "Project(%s) fail", projectName)
	}

	projectPath, err := utils.ProjectPath(projectName)
	if err != nil {
		return nil, errors.Wrapf(err, "Project(%s) fail", projectName)
	}

	return &project{path: projectPath, gpp: gpp, author: author, license: license}, nil
}

func (project project) Path() pathutil.Path {
	return project.path
}

func (project project) CheckIntegrity() error {
	shouldBeExistsPaths := []pathutil.Path{
		project.versionFilePath(),
		project.depGopkgTomlPath(),
		project.distGoTomlPath(),
	}

	for _, path := range shouldBeExistsPaths {
		if !path.Exists() {
			return errors.Errorf("No exists %s", path)
		}
	}

	if project.author == "" {
		return errors.Errorf("Author isn't set")
	}

	if project.license == "" {
		return errors.Errorf("License isn't set")
	}

	return nil
}

func (project project) gitignorePath() pathutil.Path {
	// project.Path() must be defined, error could be ignored
	path, _ := pathutil.New(project.Path().String(), ".gitignore")
	return path
}

func (project project) versionFilePath() pathutil.Path {
	// project.Path() must be defined, error could be ignored
	path, _ := pathutil.New(project.Path().String(), "VERSION")
	return path
}

func (project project) depGopkgTomlPath() pathutil.Path {
	// project.Path() must be defined, error could be ignored
	path, _ := pathutil.New(project.Path().String(), "Gopkg.toml")
	return path
}

func (project project) licensePath() pathutil.Path {
	// project.Path() must be defined, error could be ignored
	path, _ := pathutil.New(project.Path().String(), "LICENSE")
	return path
}

func (project project) distGoTomlPath() pathutil.Path {
	// project.Path() must be defined, error could be ignored
	path, _ := pathutil.New(project.Path().String(), "dist-go.toml")
	return path
}

func (project project) goDocDownTmplPath() pathutil.Path {
	// project.Path() must be defined, error could be ignored
	path, _ := pathutil.New(project.Path().String(), ".godocdown.tmpl")
	return path
}

func (project project) readmePath() pathutil.Path {
	// project.Path() must be defined, error could be ignored
	path, _ := pathutil.New(project.Path().String(), "README.md")
	return path
}
