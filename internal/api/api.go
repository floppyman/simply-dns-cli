package api

import (
	"encoding/json"
	"fmt"
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

func CreateDnsRecord(productObject string, obj *SimplyDnsRecord) (*SimplyApiSuccessResponse, error) {
	res, err := postRequest(fmt.Sprintf("/my/products/%s/dns/records", productObject), obj)
	if err != nil {
		return nil, err
	}

	var response *SimplyApiSuccessResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func UpdateDnsRecord(productObject string, recordId int64, obj *SimplyDnsRecord) (*SimplyApiSuccessResponse, error) {
	res, err := putRequest(fmt.Sprintf("/my/products/%s/dns/records/%d", productObject, recordId), obj)
	if err != nil {
		return nil, err
	}

	var response *SimplyApiSuccessResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func DeleteDnsRecord(productObject string, recordId int64) (*SimplyApiSuccessResponse, error) {
	res, err := deleteRequest(fmt.Sprintf("/my/products/%s/dns/records/%d", productObject, recordId))
	if err != nil {
		return nil, err
	}

	var response *SimplyApiSuccessResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetProducts() ([]*SimplyProduct, error) {
	res, err := getRequest("/my/products")
	if err != nil {
		return nil, err
	}

	var records SimplyApiProducts
	err = json.Unmarshal(res, &records)
	if err != nil {
		return nil, err
	}

	return records.Products, nil
}
