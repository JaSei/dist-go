package dist

import (
	"github.com/JaSei/dist-go/project"
	"github.com/JaSei/dist-go/utils"
	"github.com/pkg/errors"
)

// NewLib create new project
func NewLib(projectName, author, license string) (err error) {
	if err = utils.CheckGoPath(); err != nil {
		return
	}

	proj, err := project.New(projectName, author, license)
	if err != nil {
		return errors.Wrap(err, "NewLib failed")
	}

	// catch panic, remove uncomplete dir and set error
	defer func() {
		if r := recover(); r != nil {
			if err = proj.Path().RemoveTree(); err != nil {
				err = errors.Wrapf(err, "Remove uncomplete path %s fail", proj.Path())
			} else {
				err = errors.Wrap(r.(error), "NewLib failed")
			}
		}
	}()

	if err := proj.VCSInit(); err != nil {
		panic(err)
	}
	if err := proj.MakeGitIgnore(); err != nil {
		panic(err)
	}
	if err := proj.MakeVersionFile(); err != nil {
		panic(err)
	}
	if err := proj.MakeExampleLib(); err != nil {
		panic(err)
	}
	if err := proj.MakeDepFiles(); err != nil {
		panic(err)
	}
	//proj.GenerateDistConf()
	//proj.GenerateTravis()
	//proj.GenerateAppVeyor()
	proj.MakeLicenseFile()

	return nil
}
