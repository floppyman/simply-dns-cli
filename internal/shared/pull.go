package shared

import (
	"fmt"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

func PullProductsAndDnsRecords() []*api.SimplyProduct {
	products := PullProducts()
	if products == nil {
		return nil
	}

	styles.WaitPrint("getting dns records from each product")
	for _, product := range products {
		product = PullDnsRecordsForProduct(product, "  ")
	}
	styles.SuccessPrint("dns records downloaded")

	return products
}

func PullProducts() []*api.SimplyProduct {
	styles.WaitPrint("getting products from account")
	products, err := api.GetProducts()
	if err != nil {
		styles.FailPrint("failed to get products")
		styles.FailPrint("error: %v", err)
		return nil
	}
	styles.SuccessPrint("products downloaded")

	return products
}

func PullDnsRecordsForProduct(product *api.SimplyProduct, printPrefix string) *api.SimplyProduct {
	styles.BlankPrint(fmt.Sprintf("%s%s", printPrefix, product.Name))
	records, err := api.GetDnsRecords(product.Object)
	if err != nil {
		styles.WarnPrint(fmt.Sprintf("%s%s >> %v", printPrefix, product.Name, err.Error()))
		product.DnsRecords = make([]*api.SimplyDnsRecord, 0)
		return product
	}

	product.DnsRecords = records
	return product
}
