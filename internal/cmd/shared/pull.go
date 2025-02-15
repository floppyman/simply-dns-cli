package shared

import (
	"fmt"

	log "github.com/umbrella-sh/um-common/logging/basic"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
)

func PullProductsAndDnsRecords() []*api.SimplyProduct {
	products := PullProducts()
	if products == nil {
		return nil
	}

	log.WaitPrint("getting dns records from each product")
	for _, product := range products {
		product = PullDnsRecordsForProduct(product, "  ")
	}
	log.SuccessPrint("dns records downloaded")

	return products
}

func PullProducts() []*api.SimplyProduct {
	log.WaitPrint("getting products from account")
	products, err := api.GetProducts()
	if err != nil {
		log.FailPrint("failed to get products")
		log.Errorln(err)
		return nil
	}
	log.SuccessPrint("products downloaded")

	return products
}

func PullDnsRecordsForProduct(product *api.SimplyProduct, printPrefix string) *api.SimplyProduct {
	log.BlankPrint(fmt.Sprintf("%s%s", printPrefix, product.Name))
	records, err := api.GetDnsRecords(product.Object)
	if err != nil {
		log.WarnPrint(fmt.Sprintf("%s%s >> %v", printPrefix, product.Name, err.Error()))
		product.DnsRecords = make([]*api.SimplyDnsRecord, 0)
		return product
	}

	product.DnsRecords = records
	return product
}
