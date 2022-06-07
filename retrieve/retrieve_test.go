package retrieve

import (
	"fmt"
	"log"
	"testing"
)

var lincenseUrl = "http://api-license-ensaas.hz.wise-paas.com.cn/v1"

func TestRetriever_LicenseIDAndPn(t *testing.T) {
	r := NewRetriever(lincenseUrl)
	lic, err := r.LicenseIDAndPn("ews0019f9cde69-327b-4b74-af49-f2abafd4529en202202180333317", "980GDSHCS00")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(lic)
}

func TestRetriever_LicenseWithActiveInfoBy(t *testing.T) {
	r := NewRetriever(lincenseUrl)
	lic, err := r.LicenseWithActiveInfoBy("ensaasc1f487fb-b676-47ba-ad21-5a94ab8f853bensaas-service", "EnSaaS-ESM")
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, v := range lic {
		fmt.Println(v)
	}
}

func TestRetriever_LicenseWithoutActiveInfoBy(t *testing.T) {
	r := NewRetriever(lincenseUrl)
	lic, err := r.LicenseWithoutActiveInfoBy("ensaasc1f487fb-b676-47ba-ad21-5a94ab8f853bensaas-service", "EnSaaS-ESM")
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, v := range lic {
		fmt.Println(v)
	}
}
