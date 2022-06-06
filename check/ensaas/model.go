package ensaas

import (
	"errors"
	"fmt"
	"github.com/ensaas/license-sdk/check"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type license struct {
	licenseID       string
	pn              string
	authCode        string
	activeInfo      string
	expireTimestamp int64
	number          int
}

func newLicense(licenseID, pn, authcode, activeInfo string, timestamp int64, number int) *license {
	return &license{
		licenseID:       licenseID,
		pn:              pn,
		authCode:        authcode,
		activeInfo:      activeInfo,
		expireTimestamp: timestamp,
		number:          number,
	}
}

func (lic *license) validateAuthCode() (err error) {
	defer func() {
		if val := recover(); val != nil {
			err = val.(error)
			logrus.Println(err)
		}
	}()

	ss := fmt.Sprintf("%s+%s+%s", lic.pn, strconv.Itoa(lic.number), lic.activeInfo)
	md5Str1, err := check.Md5SumString(ss)
	if err != nil {
		logrus.Printf("generate auth code failed:%s", err.Error())
		return err
	}

	randStr1 := strings.Split(lic.authCode, "-")[0][3:4]
	randNum1, err := strconv.Atoi(randStr1)
	if err != nil {
		logrus.Printf("parse rand number failed")
		return err
	}

	p1 := md5Str1[randNum1 : randNum1+3]
	p1 = fmt.Sprintf("%s%d", p1, randNum1)

	randStr2 := strings.Split(lic.authCode, "-")[1][2:3]
	randNum2, err := strconv.Atoi(randStr2)
	if err != nil {
		logrus.Printf("parse rand number failed")
		return err
	}

	p2 := md5Str1[randNum2 : randNum2+2]
	p2 = fmt.Sprintf("%s%d%d", p2, randNum2, len(md5Str1[randNum2:])%10)

	s2 := fmt.Sprintf("%s+%s+%s", lic.licenseID, strconv.FormatInt(lic.expireTimestamp, 10), lic.activeInfo)
	md5Str2, err := check.Md5SumString(s2)
	if err != nil {
		logrus.Printf("generate auth code failed:%s", err.Error())
		return err
	}

	randStr1 = strings.Split(lic.authCode, "-")[2][3:4]
	randNum1, err = strconv.Atoi(randStr1)
	if err != nil {
		logrus.Printf("parse rand number failed")
		return err
	}

	p3 := md5Str2[randNum1 : randNum1+3]
	p3 = fmt.Sprintf("%s%d", p3, randNum1)

	randStr2 = strings.Split(lic.authCode, "-")[3][2:3]
	randNum2, err = strconv.Atoi(randStr2)
	if err != nil {
		logrus.Printf("parse rand number failed")
		return err
	}

	p4 := md5Str2[randNum2 : randNum2+2]
	p4 = fmt.Sprintf("%s%d%d", p4, randNum2, len(md5Str2[randNum2:])%10)

	authcodeStr := fmt.Sprintf("%s-%s-%s-%s", p1, p2, p3, p4)
	if lic.authCode != authcodeStr {
		return errors.New("auth code not correct")
	}

	return nil
}
