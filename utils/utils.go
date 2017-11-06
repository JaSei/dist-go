package utils

import (
	"os"

	"github.com/JaSei/pathutil-go"
	"github.com/pkg/errors"
)

func GoSrcPath() (pathutil.Path, error) {
	return pathutil.NewPath(os.Getenv("GOPATH"), "src")
}

func CheckGoPath() error {
	if os.Getenv("GOPATH") == "" {
		return errors.New("Environment variable GOPATH isn't set")
	}

	return nil
}

func ProjectPath(projectName string) (pathutil.Path, error) {
	return pathutil.NewPath(os.Getenv("GOPATH"), "src", projectName)
}
