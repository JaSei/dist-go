package dist

import (
	"github.com/JaSei/dist-go/project"
)

func Test() error {
	proj, err := project.LoadCwdProject()
	if err != nil {
		return err
	}

	if err = proj.GenerateReadme(); err != nil {
		return err
	}

	//proj.DepEnsure()
	//proj.GoGenerate()

	if err = proj.GoTestCover(); err != nil {
		return err
	}

	return nil
}

//func DepEnsure() error {
//	if err := GoGet("github.com/golang/dep/cmd/dep"); err != nil {
//		return err
//	}
//	if err := Run("dep", "ensure", "-v"); err != nil {
//		return err
//	}
//
//	return nil
//}
