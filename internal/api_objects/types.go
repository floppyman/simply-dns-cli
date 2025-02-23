package api_objects

type SimplyApiConfig struct {
	Url           string `json:"url"`
	AccountNumber string `json:"account_number"`
	AccountApiKey string `json:"account_api_key"`
}

type SimplyApiDnsRecords struct {
	Records []*SimplyDnsRecord `json:"records"`
}

type SimplyApiProducts struct {
	Products []*SimplyProduct `json:"products"`
}

type SimplyApiSuccessResponse struct {
	Record struct {
		Id int `json:"id"`
	} `json:"record"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}
