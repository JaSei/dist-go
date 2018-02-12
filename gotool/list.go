// gotool package wrap go cmd for better use in code
package gotool

import (
	"strings"

	"github.com/JaSei/dist-go/executil"
)

// GoListWithoutVendor return recursive list of packages in current working directory and ignore vendor directory
// same as
//	go list ./... | grep -v vendor
func GoListWithoutVendor() ([]string, error) {
	listDependency := make([]string, 0)

	if err := executil.RunLines(func(line string) {
		if !strings.Contains(line, "/vendor/") {
			listDependency = append(listDependency, line)
		}
	}, "go", "list", "./..."); err != nil {
		return nil, err
	}

	return listDependency, nil
}
