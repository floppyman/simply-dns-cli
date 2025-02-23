package api

import (
	"encoding/json"
	"fmt"

	apio "github.com/umbrella-sh/simply-dns-cli/internal/api_objects"
	"github.com/umbrella-sh/simply-dns-cli/internal/configs"
	"github.com/umbrella-sh/simply-dns-cli/internal/mocks"
)

// https://www.simply.com/dk/docs/api/

var currentConfig apio.SimplyApiConfig

func Init(config apio.SimplyApiConfig) {
	currentConfig = config
}

func GetDnsRecords(productObject string) ([]*apio.SimplyDnsRecord, error) {
	if configs.IsMocking {
		return mocks.GetDnsRecords()
	}

	res, err := getRequest(fmt.Sprintf("/my/products/%s/dns/records", productObject))
	if err != nil {
		return nil, err
	}

	var records apio.SimplyApiDnsRecords
	err = json.Unmarshal(res, &records)
	if err != nil {
		return nil, err
	}

	return records.Records, nil
}

func CreateDnsRecord(productObject string, obj *apio.SimplyDnsRecord) (*apio.SimplyApiSuccessResponse, error) {
	if configs.IsMocking {
		return mocks.CreateDnsRecord()
	}

	res, err := postRequest(fmt.Sprintf("/my/products/%s/dns/records", productObject), obj)
	if err != nil {
		return nil, err
	}

	var response *apio.SimplyApiSuccessResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func UpdateDnsRecord(productObject string, recordId int64, obj *apio.SimplyDnsRecord) (*apio.SimplyApiSuccessResponse, error) {
	if configs.IsMocking {
		return mocks.UpdateDnsRecord()
	}

	res, err := putRequest(fmt.Sprintf("/my/products/%s/dns/records/%d", productObject, recordId), obj)
	if err != nil {
		return nil, err
	}

	var response *apio.SimplyApiSuccessResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func RemoveDnsRecord(productObject string, recordId int64) (*apio.SimplyApiSuccessResponse, error) {
	if configs.IsMocking {
		return mocks.RemoveDnsRecord()
	}

	res, err := deleteRequest(fmt.Sprintf("/my/products/%s/dns/records/%d", productObject, recordId))
	if err != nil {
		return nil, err
	}

	var response *apio.SimplyApiSuccessResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetProducts() ([]*apio.SimplyProduct, error) {
	if configs.IsMocking {
		return mocks.GetProducts()
	}

	res, err := getRequest("/my/products")
	if err != nil {
		return nil, err
	}

	var records apio.SimplyApiProducts
	err = json.Unmarshal(res, &records)
	if err != nil {
		return nil, err
	}

	return records.Products, nil
}
