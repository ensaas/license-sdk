package liv

import (
	"fmt"
	"github.com/ensaas/license-sdk/models"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestWrapper_GetVersion(t *testing.T) {
	w := New()
	version, err := w.GetVersion()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(version)
}

func TestWrapper_CheckAuthCode(t *testing.T) {
	lic := &models.License{
		LicenseID:       "ewsb7648b48-fd1d-4ac5-bd40-14bb90536269sd",
		Pn:              "9806WPDASH",
		ActiveInfo:      "",
		Authcode:        "5083-2101-0001",
		Number:          1,
		ExpireTimestamp: 0,
	}
	w := New()
	isValid, err := w.CheckAuthCode(lic)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isValid)
}

func TestWrapper_CheckAuthCodeWithActiveInfo(t *testing.T) {
	lic := &models.License{
		LicenseID:       "ensaas6756e64d-3b9d-43ae-b6f6-1c6df0edb4bfensaas-service",
		Pn:              "980GSSOPS00",
		ActiveInfo:      "c86b5f1e867cde5ad4020360c03552c7",
		Authcode:        "92c8-4539-5d11-c848",
		Number:          1,
		ExpireTimestamp: 1686182400,
	}
	w := New()
	isValid, err := w.CheckAuthCode(lic)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", isValid)
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
