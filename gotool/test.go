package gotool

import (
	"fmt"
	"os/exec"

	"github.com/JaSei/pathutil-go"
	"github.com/pkg/errors"
)

// GoTestCover method do test and calculate coverage
// (ignore vendor file)
// same as
//	echo 'mode: atomic' > $coveragePath && go list ./... | grep -v vendor | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> $coveragePath && rm coverage.tmp
func GoTestCover(coveragePath pathutil.Path) error {
	if coveragePath.Exists() {
		if err := coveragePath.Remove(); err != nil {
			return errors.Wrap(err, "Remove previous coverage.txt")
		}
	}

	if err := coveragePath.Append("mode: atomic\n"); err != nil {
		return errors.Wrap(err, "Append to coverage.txt fail")
	}

	depList, err := GoListWithoutVendor()
	if err != nil {
		return errors.Wrap(err, "go list")
	}

	for _, dep := range depList {
		_ = goTest(coveragePath, dep)
	}

	return nil
}

func goTest(coverprofile pathutil.Path, dep string) (err error) {
	temp, err := pathutil.NewTempFile(pathutil.TempOpt{})
	defer func() {
		err = errors.Wrap(temp.Remove(), "remove temp file fail")
	}()

	out, err := exec.Command("go", "test", "-covermode=atomic", "-coverprofile="+temp.Canonpath(), dep).Output()

	if err != nil {
		return err
	}

	fmt.Print(string(out))

	lines, err := temp.Lines()
	if err != nil {
		return err
	}

	if len(lines) > 2 {
		for _, line := range lines[1:] {
			coverprofile.Append(line + "\n")
		}
	}

	return nil
}
