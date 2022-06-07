package liv

import "os/exec"

func Run(cmd *exec.Cmd) ([]byte, error) {
	return cmd.CombinedOutput()
}

func LookPath(file string) (string, error) {
	return exec.LookPath(file)
}
