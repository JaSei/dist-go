package executil

import (
	"bufio"
	"os/exec"

	"github.com/pkg/errors"
)

// RunLines exec run command with commandArgs
// and stdout is parsed line by line and send to lineFunc
func RunLines(lineFunc func(string), command string, commandArgs ...string) error {
	cmd := exec.Command(command, commandArgs...)

	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "Open stdout pipe fail")
	}
	if err = cmd.Start(); err != nil {
		return errors.Wrapf(err, "Start process %s fail", command)
	}

	scanner := bufio.NewScanner(cmdOut)

	for scanner.Scan() {
		lineFunc(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return errors.Wrap(err, "Scanner fail")
	}

	if err = cmd.Wait(); err != nil {
		return errors.Wrapf(err, "Process %s fail", command)
	}

	return nil
}

//func Run(command string, commandArgs ...string) error {
//	cmd := exec.Command(command, commandArgs...)
//	if err := cmd.Run(); err != nil {
//		return errors.Wrap(err, "Run")
//	}
//
//	return nil
//}
