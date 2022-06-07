package liv

import (
	"fmt"
	"github.com/ensaas/license-sdk/common"
	"os/exec"
	"strconv"
)

// Wrapper is wrap methods for check license auth code
type Wrapper interface {
	CheckAuthCode() (bool, error)
	GetVersion() (string, error)
}

type wrapper struct {
	license *common.License
}

func New(lic *common.License) Wrapper {
	return &wrapper{license: lic}
}

func (w *wrapper) GetVersion() (string, error) {
	cmd, err := w.prepareVersionCmd()
	if err != nil {
		return "", fmt.Errorf("failed preparing liv version command:%w", err)
	}

	output, err := Run(cmd)
	if err != nil {
		return "", fmt.Errorf("failed running liv version command: %w:%v", err, string(output))
	}

	return string(output), nil
}

func (w *wrapper) CheckAuthCode() (bool, error) {
	cmd, err := w.prepareCheckCmd()
	if err != nil {
		return false, fmt.Errorf("failed preparing liv check command: %w", err)
	}

	output, err := Run(cmd)
	if err != nil {
		return false, fmt.Errorf("failed run liv check command: %w:%v", err, string(output))
	}
	b, err := strconv.ParseBool(string(output))
	if err != nil {
		return false, fmt.Errorf("failed parse output of liv check command:%w", err)
	}
	return b, nil
}

func (w *wrapper) prepareCheckCmd() (*exec.Cmd, error) {
	args := []string{
		"validate",
		"--PN", w.license.Pn,
		"--LicenseID", w.license.LicenseID,
		"--Authcode", w.license.Authcode,
		"--Number", strconv.FormatInt(w.license.Number, 10),
	}

	if len(w.license.ActiveInfo) != 0 {
		args = append([]string{"--ActiveInfo", w.license.ActiveInfo}, args...)
	}
	if w.license.ExpireTimestamp != 0 {
		args = append([]string{"--Expire", strconv.FormatInt(w.license.ExpireTimestamp, 10)}, args...)
	}

	name, err := LookExec()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(name, args...)

	return cmd, nil
}

func (w *wrapper) prepareVersionCmd() (*exec.Cmd, error) {
	args := []string{
		"version",
	}

	name, err := LookExec()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(name, args...)
	return cmd, nil
}
