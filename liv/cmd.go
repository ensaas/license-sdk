package liv

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	livLinux   = "Liv"
	livWindows = "Liv.exe"
)

func Run(cmd *exec.Cmd) ([]byte, error) {
	return cmd.CombinedOutput()
}

// LookExec look executable file for license
func LookExec() (name string, err error) {
	var absPath string
	// for linux
	if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
		name, err = exec.LookPath(livLinux)
		if err == nil {
			return
		}

		absPath, err = os.Getwd()
		if err != nil {
			return
		}
		name, err = exec.LookPath(fmt.Sprintf("%s/bin/%s", strings.TrimRight(absPath, "/liv"), livLinux))
		return
	}
	// for windows
	if runtime.GOOS == "windows" {
		name, err = exec.LookPath(livWindows)
		if err == nil {
			return
		}
		absPath, err = os.Getwd()
		if err != nil {
			return
		}
		name, err = exec.LookPath(fmt.Sprintf("%s/bin/%s", strings.TrimRight(absPath, "/liv"), livWindows))
		return
	}

	err = fmt.Errorf("executable Liv not found,GOOS %s and arch %s not supported", runtime.GOOS, runtime.GOARCH)
	return
}
