package dist

import (
	"github.com/JaSei/dist-go/project"
)

func Test() error {
	proj, err := project.LoadCwdProject()
	if err != nil {
		return err
	}
	_ = proj.Path()
	//proj.GenerateReadme()
	//proj.DepEnsure()
	//proj.GoTestCover()

	return nil
}

func DepEnsure() error {
	if err := GoGet("github.com/golang/dep/cmd/dep"); err != nil {
		return err
	}
	if err := Run("dep", "ensure", "-v"); err != nil {
		return err
	}

	return nil
}
