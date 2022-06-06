package datasource

type EnSaaSLicenseList []*EnSaaSLicense

type EnSaaSLicense struct {
	LicenseID       string `json:"licenseId"`
	Pn              string `json:"pn"`
	AuthCode        string `json:"authCode"`
	ActiveInfo      string `json:"activeInfo"`
	ExpireTimestamp int64  `json:"expireTimestamp"`
	Number          int    `json:"number"`
}

type SrpLicenseList []*SrpLicense

type SrpLicense struct {
	LicenseID          string `json:"id"`
	PN                 string `json:"pn"`
	SubscriptionId     string `json:"subscriptionId"`
	DatacenterCode     string `json:"datacenterCode"`
	IsValidTransaction bool   `json:"isValidTransaction"`
	Number             int    `json:"number"`
	AuthCode           string `json:"authcode"`
	Company            string `json:"company"`
	CreateTime         string `json:"createTime"`
	ExpireTime         string `json:"expireTime"`
}
