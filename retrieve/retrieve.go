package retrieve

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"net/url"
	"strings"
	"time"
)

/**
  Retriever means license retriever,which get license by restful api
*/
type Retriever interface {
	LicenseIDAndPn(id, pn string) (*LicenseResponse, error)
	LicenseWithActiveInfoBy(id, service string) ([]*LicenseResponse, error)
	LicenseWithoutActiveInfoBy(id, service string) ([]*LicenseResponse, error)
}

type retriever struct {
	licenseUrl string
	httpClient *resty.Client
}

func NewRetriever(licenseUrl string) Retriever {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(20 * time.Second)
	client.SetCloseConnection(true)

	return &retriever{httpClient: client, licenseUrl: licenseUrl}
}

type licenseResource struct {
	TotalCount int         `json:"total"`
	Resources  []*resource `json:"resources"`
}

type resource struct {
	Pn        string `json:"pn"`
	LicenseId string `json:"id"`
	Authcode  string `json:"authcode"`
	Number    int    `json:"number"`
}

//LicenseIDAndPn get license by license id and pn
func (r *retriever) LicenseIDAndPn(id, pn string) (*LicenseResponse, error) {
	licenseUrl := fmt.Sprintf("%s/%s?pn=%s&id=%s", strings.TrimRight(r.licenseUrl, "/"), "api/partNum/licenseQty",
		url.QueryEscape(pn), url.QueryEscape(id))
	logrus.Infof("license url is %s", licenseUrl)

	resp, err := r.httpClient.R().
		SetHeader("Content-Type", "application/json").
		Get(licenseUrl)
	if err != nil {
		logrus.Errorf("Failed get license, url is [%s],err:[%v]", licenseUrl, err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		logrus.Errorf("Call license API [%s] failed,status:[%d],response:[%s]", licenseUrl, resp.StatusCode(), string(resp.Body()))
		return nil, fmt.Errorf("Call license API [%s] failed,status:[%d],response:[%s]", licenseUrl, resp.StatusCode(), string(resp.Body()))
	}

	//check license response
	isVailid, err := validateResponse(resp)
	if err != nil {
		return nil, err
	}
	if !isVailid {
		logrus.Errorf("data from License API illegal,please check it,License url is %s, response is %s", licenseUrl, string(resp.Body()))
		return nil, fmt.Errorf("data from License API illegal,please check it,License url is %s, response is %s", licenseUrl, string(resp.Body()))
	}

	var data = &resource{}
	if err := json.Unmarshal(resp.Body(), data); err != nil {
		logrus.Errorf("unmarshal license data failed:%s", err.Error())
		return nil, fmt.Errorf("unmarshal license data failed:%s", err.Error())
	}

	return toLicense(data), nil
}

// LicenseWithoutActiveInfoBy retrieve license data by license id and service name,returned data dont have activeInfo
func (r *retriever) LicenseWithoutActiveInfoBy(id, service string) ([]*LicenseResponse, error) {
	licenseUrl := fmt.Sprintf("%s/api/serviceName/%s/serviceInstanceId/%s", strings.TrimRight(r.licenseUrl, "/"), url.QueryEscape(service), url.QueryEscape(id))
	logrus.Infof("license url is %s", licenseUrl)

	resp, err := r.httpClient.R().
		SetHeader("Content-Type", "application/json").
		Get(licenseUrl)
	if err != nil {
		logrus.Errorf("Failed get license, url is [%s],err:[%v]", licenseUrl, err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		logrus.Errorf("Call license API [%s] failed,status:[%d],response:[%s]", licenseUrl, resp.StatusCode(), string(resp.Body()))
		return nil, fmt.Errorf("Call license API [%s] failed,status:[%d],response:[%s]", licenseUrl, resp.StatusCode(), string(resp.Body()))
	}

	//check license response
	isVailid, err := validateResponse(resp)
	if err != nil {
		return nil, err
	}
	if !isVailid {
		logrus.Errorf("data from License API illegal,please check it,License url is %s, response is %s", licenseUrl, string(resp.Body()))
		return nil, fmt.Errorf("data from License API illegal,please check it,License url is %s, response is %s", licenseUrl, string(resp.Body()))
	}

	var data = &licenseResource{}
	if err := json.Unmarshal(resp.Body(), data); err != nil {
		logrus.Errorf("unmarshal license data failed:%s", err.Error())
		return nil, fmt.Errorf("unmarshal license data failed:%s", err.Error())
	}
	licList := make([]*LicenseResponse, 0)
	for _, v := range data.Resources {
		licList = append(licList, toLicense(v))
	}

	return licList, nil
}

func toLicense(lic *resource) *LicenseResponse {
	return &LicenseResponse{
		Pn:        lic.Pn,
		LicenseId: lic.LicenseId,
		AuthCode:  lic.Authcode,
		Number:    lic.Number,
	}
}

type LicenseData struct {
	TotalCount  int                `json:"total"`
	LicenseList []*LicenseResponse `json:"resources"`
}

type LicenseResponse struct {
	Pn              string `json:"pn"`
	AuthCode        string `json:"authCode"`
	LicenseId       string `json:"licenseId"`
	ExpireTimestamp int64  `json:"expireTimestamp"`
	Number          int    `json:"number"`
}

// LicenseWithActiveInfoBy retrieve license data by license id and service name,returned data has activeInfo
func (r *retriever) LicenseWithActiveInfoBy(id, service string) ([]*LicenseResponse, error) {
	licenseUrl := fmt.Sprintf("%s/api/ensaasService/%s/licenseId/%s", strings.TrimRight(r.licenseUrl, "/"), url.QueryEscape(service), url.QueryEscape(id))
	logrus.Infof("license url is %s,service name is %s,licenseId is %s", licenseUrl, service, id)

	resp, err := r.httpClient.R().
		SetHeader("Content-Type", "application/json").
		Get(licenseUrl)
	if err != nil {
		logrus.Errorf("Failed get license, url is [%s],err:[%v]", licenseUrl, err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		logrus.Errorf("Call license API [%s] failed,status:[%d],response:[%s]", licenseUrl, resp.StatusCode(), string(resp.Body()))
		return nil, fmt.Errorf("Call license API [%s] failed,status:[%d],response:[%s]", licenseUrl, resp.StatusCode(), string(resp.Body()))
	}

	//check license response
	isVailid, err := validateResponse(resp)
	if err != nil {
		return nil, err
	}
	if !isVailid {
		logrus.Errorf("data from License API illegal,please check it,License url is %s, response is %s", licenseUrl, string(resp.Body()))
		return nil, fmt.Errorf("data from License API illegal,please check it,License url is %s, response is %s", licenseUrl, string(resp.Body()))
	}

	var data = &LicenseData{}
	if err := json.Unmarshal(resp.Body(), data); err != nil {
		logrus.Errorf("unmarshal license data failed:%s", err.Error())
		return nil, fmt.Errorf("unmarshal license data failed:%s", err.Error())
	}
	// set license id
	for _, v := range data.LicenseList {
		v.LicenseId = id
	}

	return data.LicenseList, nil
}

// validateResponse validate reponse from server is valid
func validateResponse(response *resty.Response) (bool, error) {
	checksum := response.Header().Get("checksum")
	if len(checksum) == 0 {
		logrus.Warningf("License Response: checksum not found in header")
		return true, nil
	}

	dataBytes, err := base64.StdEncoding.DecodeString(checksum)
	if err != nil {
		logrus.Errorf("[License API Response] decoding checksum failed:[%v]", err)
		return false, fmt.Errorf("decoding checksum failed:[%v]", err)
	}
	isValid, err := RsaVerySignWithSha256(response.Body(), dataBytes)
	if err != nil {
		logrus.Errorf("[License API Response] validate checksum failed:%v", err)
		return false, fmt.Errorf("validate checksum failed:%v", err)
	}
	return isValid, nil
}
