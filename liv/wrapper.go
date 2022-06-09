package liv

import (
	"fmt"
	"github.com/ensaas/license-sdk/models"
	"os/exec"
	"strconv"
)

// Wrapper is wrap methods for check license auth code
type Wrapper interface {
	CheckAuthCode(lic *models.License) (bool, error)
	GetVersion() (string, error)
}

type wrapper struct{}

func New() Wrapper {
	return &wrapper{}
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

func (w *wrapper) CheckAuthCode(lic *models.License) (bool, error) {
	cmd, err := w.prepareCheckCmd(lic)
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

func (w *wrapper) prepareCheckCmd(lic *models.License) (*exec.Cmd, error) {
	args := []string{
		"validate",
		"--PN", lic.Pn,
		"--LicenseID", lic.LicenseID,
		"--Authcode", lic.Authcode,
		"--Number", strconv.Itoa(lic.Number),
	}

	if len(lic.ActiveInfo) != 0 {
		args = append([]string{"--ActiveInfo", lic.ActiveInfo}, args...)
	}
	if lic.ExpireTimestamp != 0 {
		args = append([]string{"--Expire", strconv.FormatInt(lic.ExpireTimestamp, 10)}, args...)
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
