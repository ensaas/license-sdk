package common

import (
	"errors"
	"fmt"
	"os"
)

type License struct {
	LicenseID       string
	Pn              string
	ActiveInfo      string
	Authcode        string
	Number          int64
	ExpireTimestamp int64
}

func NewLicense(pnArr []string, activeInfo string) (*License, error) {
	if len(pnArr) == 0 {
		return nil, errors.New("license pn is empty")
	}

	cluster := os.Getenv("cluster")
	workspace := os.Getenv("workspace")
	namespace := os.Getenv("namespace")

	return &License{
		ActiveInfo: activeInfo,
		LicenseID:  fmt.Sprintf("%s%s%s", cluster, workspace, namespace),
	}, nil
}
