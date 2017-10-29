package dist

import (
	"os"
	"os/exec"

	"github.com/JaSei/pathutil-go"
	"github.com/pkg/errors"
)

func CheckGoPath() error {
	if os.Getenv("GOPATH") == "" {
		return errors.New("Environment variable GOPATH isn't set")
	}

	return nil
}

func GoSrcPath() (pathutil.Path, error) {
	return pathutil.NewPath(os.Getenv("GOPATH"), "src")
}

func ProjectPath(projectName string) (pathutil.Path, error) {
	return pathutil.NewPath(os.Getenv("GOPATH"), "src", projectName)
}

func NewApp() error {
	return nil
}

func Release() error {
	return nil
}

func GoGet(repo string) error {
	goGet := exec.Command("go", "get", "-u", "-v", repo)
	if err := goGet.Run(); err != nil {
		return errors.Wrap(err, "go get")
	}

	return nil
}

func Run(command string, commandArgs ...string) error {
	cmd := exec.Command(command, commandArgs...)
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "Run")
	}

	return nil
}
