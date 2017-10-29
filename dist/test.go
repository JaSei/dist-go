package dist

func Test() error {
	project, err := LoadCwdProject()
	if err != nil {
		return err
	}
	_ = project.Path()
	//project.GenerateReadme()
	//project.DepEnsure()
	//project.GoTestCover()

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
