package goversion

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/blang/semver"
)

// GoVersion function return version of your go like `go version`
// this method use `go version` command and semver parser
func GoVersion() (*semver.Version, error) {
	outBytes, err := exec.Command("go", "version").Output()
	if err != nil {
		return nil, err
	}

	out := string(outBytes[:])

	//go version go1.9.2 linux/amd64
	strVer := strings.TrimPrefix(out, "go version")
	strVer = strings.TrimPrefix(strVer, " go")
	strVer = strings.TrimSuffix(strVer, " "+runtime.GOOS+"/"+runtime.GOARCH+"\n")

	return semver.New(strVer)
}
