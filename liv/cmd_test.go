package liv

import (
	"fmt"
	"github.com/ensaas/license-sdk/common"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func TestWrapper_GetVersion(t *testing.T) {
	lic := &common.License{
		LicenseID:       "ensaasba529bbb-43c9-45b3-8d70-564a5f61d9e6ensaas-service",
		ServiceName:     "Dashboard",
		Pn:              "980GDSHCP00",
		ActiveInfo:      "",
		Authcode:        "c4b1-7c66-0001",
		Number:          1,
		ExpireTimestamp: 0,
	}
	w := New(lic)
	version, err := w.GetVersion()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(version)
}

func TestWrapper_CheckAuthCode(t *testing.T) {
	lic := &common.License{
		LicenseID:       "ewsb7648b48-fd1d-4ac5-bd40-14bb90536269sd",
		ServiceName:     "Dashboard",
		Pn:              "9806WPDASH",
		ActiveInfo:      "",
		Authcode:        "5083-2101-0001",
		Number:          1,
		ExpireTimestamp: 0,
	}
	w := New(lic)
	isValid, err := w.CheckAuthCode()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isValid)
}

func TestWrapper_CheckAuthCodeWithActiveInfo(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	fmt.Printf(filename, ok)
	lic := &common.License{
		LicenseID:       "ensaase80add40-1f47-4615-9285-0e7a0ad765b7ensaas-service",
		ServiceName:     "Config-Mgmt.",
		Pn:              "980GEKSCM00",
		ActiveInfo:      "sYeDFf0NKw9J9vG4z3oNx6tmojUDL77sUjz7PZwVnxpdytUHdzQD+PW/L9Ygwwf/",
		Authcode:        "5b15-ef02-27e4-7e57",
		Number:          6,
		ExpireTimestamp: 4790890858,
	}
	w := New(lic)
	isValid, err := w.CheckAuthCode()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isValid)
}

func TestWrapper(t *testing.T) {
	absPath, err := os.Getwd()
	if err != nil {
		return
	}
	name, err := exec.LookPath(fmt.Sprintf("%s/bin/%s", strings.TrimRight(absPath, "/liv"), livLinux))
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(name)
}
