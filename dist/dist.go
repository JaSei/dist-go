package dist

import (
	"os/exec"

	"github.com/pkg/errors"
)

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
