package retrieve

import (
	"fmt"
	"log"
	"testing"
)

var lincenseUrl = "http://localhost:8080/v1"

func TestRetriever_LicenseIDAndPn(t *testing.T) {
	r := NewRetriever(lincenseUrl)
	lic, err := r.LicenseIDAndPn("dddaadacapcsa", "dd111ddd")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(lic)
}

func TestRetriever_LicenseWithActiveInfoBy(t *testing.T) {
	r := NewRetriever(lincenseUrl)
	lic, err := r.LicenseWithActiveInfoBy("ews02p12lpl", "cccaaa")
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, v := range lic {
		fmt.Println(v)
	}
}

func TestRetriever_LicenseWithoutActiveInfoBy(t *testing.T) {
	r := NewRetriever(lincenseUrl)
	lic, err := r.LicenseWithoutActiveInfoBy("cccaapa", "ttt0212lpl")
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, v := range lic {
		fmt.Println(v)
	}
}
