package api

import (
	"encoding/json"
	"fmt"

	"github.com/floppyman/simply-dns-cli/internal/objects"
	"github.com/floppyman/simply-dns-cli/internal/configs"
	"github.com/floppyman/simply-dns-cli/internal/mocks"
)

// https://www.simply.com/dk/docs/api/

var currentConfig objects.SimplyApiConfig

func Init(config objects.SimplyApiConfig) {
	currentConfig = config
}

func GetDnsRecords(productObject string) ([]*objects.SimplyDnsRecord, error) {
	if configs.IsMocking {
		return mocks.GetDnsRecords()
	}

	res, err := getRequest(fmt.Sprintf("/my/products/%s/dns/records", productObject))
	if err != nil {
		return nil, err
	}

	var records objects.SimplyApiDnsRecords
	err = json.Unmarshal(res, &records)
	if err != nil {
		return nil, err
	}

	return records.Records, nil
}

func CreateDnsRecord(productObject string, obj *objects.SimplyDnsRecord) (*objects.SimplyApiSuccessResponse, error) {
	if configs.IsMocking {
		return mocks.CreateDnsRecord()
	}

	res, err := postRequest(fmt.Sprintf("/my/products/%s/dns/records", productObject), obj)
	if err != nil {
		return nil, err
	}

	var response *objects.SimplyApiSuccessResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func UpdateDnsRecord(productObject string, recordId int64, obj *objects.SimplyDnsRecord) (*objects.SimplyApiSuccessResponse, error) {
	if configs.IsMocking {
		return mocks.UpdateDnsRecord()
	}

	res, err := putRequest(fmt.Sprintf("/my/products/%s/dns/records/%d", productObject, recordId), obj)
	if err != nil {
		return nil, err
	}

	var response *objects.SimplyApiSuccessResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func RemoveDnsRecord(productObject string, recordId int64) (*objects.SimplyApiSuccessResponse, error) {
	if configs.IsMocking {
		return mocks.RemoveDnsRecord()
	}

	res, err := deleteRequest(fmt.Sprintf("/my/products/%s/dns/records/%d", productObject, recordId))
	if err != nil {
		return nil, err
	}

	var response *objects.SimplyApiSuccessResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetProducts() ([]*objects.SimplyProduct, error) {
	if configs.IsMocking {
		return mocks.GetProducts()
	}

	res, err := getRequest("/my/products")
	if err != nil {
		return nil, err
	}

	var records objects.SimplyApiProducts
	err = json.Unmarshal(res, &records)
	if err != nil {
		return nil, err
	}

	return records.Products, nil
}
