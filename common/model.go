package common

import (
	"errors"
	"fmt"
	"os"
)

type License struct {
	LicenseID       string
	ServiceName     string
	Pn              string
	ActiveInfo      string
	Authcode        string
	Number          int64
	ExpireTimestamp int64
}

func NewLicense(serviceName string, pnArr []string, activeInfo string) (*License, error) {
	if len(serviceName) == 0 {
		return nil, errors.New("license service name is empty")
	}
	if len(pnArr) == 0 {
		return nil, errors.New("license pn is empty")
	}

	cluster := os.Getenv("cluster")
	workspace := os.Getenv("workspace")
	namespace := os.Getenv("namespace")

	return &License{
		ServiceName: serviceName,
		ActiveInfo:  activeInfo,
		LicenseID:   fmt.Sprintf("%s%s%s", cluster, workspace, namespace),
	}, nil
}
