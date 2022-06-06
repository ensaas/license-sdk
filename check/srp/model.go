package srp

import (
	"fmt"
	"github.com/ensaas/license-sdk/check"
	"strconv"
	"strings"
)

// license srp check model
type license struct {
	licenseID string
	pn        string
	authCode  string
	number    int
}

func newLicense(licenseID, pn, authcode string, number int) *license {
	return &license{
		licenseID: licenseID,
		pn:        pn,
		authCode:  authcode,
		number:    number,
	}
}

//validate authcode
func (lic *license) validateAuthCode() error {
	var err error
	defer func() {
		if val := recover(); val != nil {
			err = val.(error)
		}
	}()

	ch1 := fmt.Sprintf("%s", lic.authCode[0:4])
	ch1First := ch1[3:4]
	ch1Start := check.BHex2Num(ch1First, 10)
	ch2 := fmt.Sprintf("%s", lic.authCode[5:9])
	ch2First := ch2[3:4]
	ch2Start := check.BHex2Num(ch2First, 10)
	str := fmt.Sprintf("%s+%s+%d+%s", lic.pn, lic.licenseID, lic.number, "")
	fmt.Println(str)
	md5StrEncode, err := check.Md5SumString(str)
	if err != nil {
		return err
	}
	ch1new := md5StrEncode[ch1Start : ch1Start+3]
	ch1new = ch1new + strconv.Itoa(ch1Start)
	ch2new := md5StrEncode[ch2Start : ch2Start+2]
	ch2new = ch2new + ch2[2:3] + strconv.Itoa(ch2Start)
	ch3new := check.Num2BHex(lic.number, 36)
	ch3new = check.Lpad(ch3new, 4, '0')
	authcodenew := fmt.Sprintf("%s-%s-%s", ch1new, ch2new, ch3new)
	if strings.EqualFold(lic.authCode, authcodenew) {
		return nil
	} else {
		return nil
	}
}
