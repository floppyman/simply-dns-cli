package api

import (
	"encoding/json"
	"fmt"

	"github.com/umbrella-sh/um-common/logging/ulog"
)

// https://www.simply.com/dk/docs/api/

var currentConfig SimplyApiConfig

func Init(config SimplyApiConfig) {
	currentConfig = config
}

func GetDnsRecords(productObject string) ([]*SimplyDnsRecord, error) {
	res, err := getRequest(fmt.Sprintf("/my/products/%s/dns/records", productObject))
	if err != nil {
		return nil, err
	}

	var records SimplyApiDnsRecords
	err = json.Unmarshal(res, &records)
	if err != nil {
		return nil, err
	}

	return records.Records, nil
}

func CreateDnsRecord(productObject string, obj *SimplyDnsRecord) error {
	res, err := postRequest(fmt.Sprintf("/my/products/%s/dns/records", productObject), obj)
	if err != nil {
		return err
	}

	ulog.Console.Debug().Str("body", string(res)).Msg("CreateDnsRecord Response")

	// var records SimplyApiDnsRecords
	// err = json.Unmarshal(res, &records)
	// if err != nil {
	// 	return  err
	// }

	return nil
}

func UpdateDnsRecord(productObject string, recordId int64, obj *SimplyDnsRecord) error {
	res, err := putRequest(fmt.Sprintf("/my/products/%s/dns/records/%d", productObject, recordId), obj)
	if err != nil {
		return err
	}

	ulog.Console.Debug().Str("body", string(res)).Msg("UpdateDnsRecord Response")

	// var records SimplyApiDnsRecords
	// err = json.Unmarshal(res, &records)
	// if err != nil {
	// 	return  err
	// }

	return nil
}

func DeleteDnsRecord(productObject string, recordId int64) error {
	res, err := deleteRequest(fmt.Sprintf("/my/products/%s/dns/records/%d", productObject, recordId))
	if err != nil {
		return err
	}

	ulog.Console.Debug().Str("body", string(res)).Msg("DeleteDnsRecord Response")

	// var records SimplyApiDnsRecords
	// err = json.Unmarshal(res, &records)
	// if err != nil {
	// 	return  err
	// }

	return nil
}

func GetProducts() ([]*SimplyProduct, error) {
	res, err := getRequest("/my/products")
	if err != nil {
		ulog.Console.Debug().Msg("1")
		return nil, err
	}

	var records SimplyApiProducts
	err = json.Unmarshal(res, &records)
	if err != nil {
		return nil, err
	}

	return records.Products, nil
}
