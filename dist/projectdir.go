package dist

import (
	"fmt"

	"github.com/JaSei/dist-go/utils"
	"github.com/JaSei/pathutil-go"
	"github.com/pkg/errors"
)

type projectMapper map[string][]pathutil.Path

func PrintProjectDirectory(projectName string) error {
	path, err := utils.ProjectPath(projectName)

	if err != nil {
		errors.Wrap(err, "PrintProjectDirectory failed")
	}

	if !path.Exists() {
		mapper, err := findProject()
		if err != nil {
			return err
		}

		if projectPaths, ok := mapper[projectName]; ok {
			if len(projectPaths) == 1 {
				fmt.Println(projectPaths[0].Canonpath())
				return nil
			} else {
				for i, projectPath := range projectPaths {
					fmt.Printf("#%d: %s\n", i, projectPath)
				}

				return errors.New("Too many repositories with this name, please use full name")
			}
		}

		return errors.Errorf("Project %s not exists", projectName)
	}

	fmt.Println(path.Canonpath())
	return nil
}

func findProject() (projectMapper, error) {
	srcPath, err := utils.GoSrcPath()
	if err != nil {
		return nil, err
	}

	mapper := make(projectMapper)

	children, err := srcPath.Children()
	for _, repo := range children {
		repoChildren, _ := repo.Children()
		for _, user := range repoChildren {
			projectChildren, _ := user.Children()
			for _, project := range projectChildren {
				if project.IsDir() {
					mapper[project.Basename()] = append(mapper[project.Basename()], project)
					mapper[user.Basename()+"/"+project.Basename()] = append(mapper[user.Basename()+"/"+project.Basename()], project)
				}
			}
		}
	}

	return mapper, nil
}
