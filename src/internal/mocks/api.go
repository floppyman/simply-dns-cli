package mocks

import (
	"time"

	"github.com/floppyman/um-common/jsons"

	"github.com/floppyman/simply-dns-cli/internal/objects"
)

func GetDnsRecords() ([]*objects.SimplyDnsRecord, error) {
	return []*objects.SimplyDnsRecord{
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
			Type:     "A",
			Priority: nil,
			Comment:  "",
		},
		{
			RecordId: 6,
			Name:     "mail",
			TTL:      21600,
			Data:     "127.0.0.1",
			Type:     "MX",
			Priority: jsons.NewJsonInt32(10),
			Comment:  "",
		},
	}, nil
}

func CreateDnsRecord() (*objects.SimplyApiSuccessResponse, error) {
	return &objects.SimplyApiSuccessResponse{
		Record: struct {
			Id int `json:"id"`
		}{
			Id: 7,
		},
		Status:  200,
		Message: "",
	}, nil
}

func UpdateDnsRecord() (*objects.SimplyApiSuccessResponse, error) {
	return &objects.SimplyApiSuccessResponse{
		Status:  200,
		Message: "",
	}, nil
}

func RemoveDnsRecord() (*objects.SimplyApiSuccessResponse, error) {
	return &objects.SimplyApiSuccessResponse{
		Status:  200,
		Message: "",
	}, nil
}

func GetProducts() ([]*objects.SimplyProduct, error) {
	return []*objects.SimplyProduct{
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

func LoadBackup() *objects.RestoreFile {
	return &objects.RestoreFile{
		TimeStamp: time.Now(),
		Items: map[string]*objects.SimplyProduct{
			"domain.com": {
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
				DnsRecords: []*objects.SimplyDnsRecord{
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
				},
			},
		},
	}
}
