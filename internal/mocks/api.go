package mocks

import (
	"github.com/umbrella-sh/simply-dns-cli/internal/api_objects"
)

func GetDnsRecords() ([]*api_objects.SimplyDnsRecord, error) {
	return []*api_objects.SimplyDnsRecord{
		{
			RecordId: 1,
			Name:     "@",
			TTL:      3600,
			Data:     "ns1.simply.com",
			Type:     "NS",
			Priority: nil,
			Comment:  "",
		},
		{
			RecordId: 2,
			Name:     "@",
			TTL:      3600,
			Data:     "ns2.simply.com",
			Type:     "NS",
			Priority: nil,
			Comment:  "",
		},
		{
			RecordId: 3,
			Name:     "@",
			TTL:      3600,
			Data:     "ns3.simply.com",
			Type:     "NS",
			Priority: nil,
			Comment:  "",
		},
		{
			RecordId: 4,
			Name:     "@",
			TTL:      3600,
			Data:     "127.0.0.1",
			Type:     "A",
			Priority: nil,
			Comment:  "",
		},
		{
			RecordId: 5,
			Name:     "test",
			TTL:      600,
			Data:     "127.0.0.1",
			Type:     "NS",
			Priority: nil,
			Comment:  "",
		},
	}, nil
}

func CreateDnsRecord() (*api_objects.SimplyApiSuccessResponse, error) {
	return &api_objects.SimplyApiSuccessResponse{
		Record: struct {
			Id int `json:"id"`
		}{
			Id: 6,
		},
		Status:  200,
		Message: "",
	}, nil
}

func UpdateDnsRecord() (*api_objects.SimplyApiSuccessResponse, error) {
	return &api_objects.SimplyApiSuccessResponse{
		Status:  200,
		Message: "",
	}, nil
}

func RemoveDnsRecord() (*api_objects.SimplyApiSuccessResponse, error) {
	return &api_objects.SimplyApiSuccessResponse{
		Status:  200,
		Message: "",
	}, nil
}

func GetProducts() ([]*api_objects.SimplyProduct, error) {
	return []*api_objects.SimplyProduct{
		{
			Object:    "domain.com",
			Name:      "domain.com",
			AutoRenew: true,
			Cancelled: false,
			Domain: struct {
				Name          string `json:"name"`
				NameIdn       string `json:"name_idn"`
				Managed       bool   `json:"managed"`
				DateRenewDate int    `json:"date_renewdate"`
			}{
				Name:          "domain.com",
				NameIdn:       "domain.com",
				Managed:       false,
				DateRenewDate: 1769562000,
			},
			Product: struct {
				Id          int         `json:"id"`
				Name        string      `json:"name"`
				DateCreated int         `json:"date_created"`
				DateExpire  interface{} `json:"date_expire"`
			}{
				Id:          1,
				Name:        "dnsservice",
				DateCreated: 1738241585,
				DateExpire:  nil,
			},
		},
	}, nil
}
