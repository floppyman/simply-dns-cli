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

	styles.WaitPrint("Getting DNS records from each product")
	for _, product := range products {
		product = PullDnsRecordsForProduct(product, "  ")
	}
	styles.SuccessPrint("DNS records downloaded")

	return products
}

func PullProducts() []*api.SimplyProduct {
	styles.WaitPrint("Getting products from account")
	products, err := api.GetProducts()
	if err != nil {
		styles.FailPrint("Failed to get products")
		styles.FailPrint("Error: %v", err)
		return nil
	}
	styles.SuccessPrint("Products downloaded")

	return products
}

func PullDnsRecords(productObject string, printPrefix string) []*api.SimplyDnsRecord {
	styles.WaitPrint("Getting dns records for '%s'", productObject)
	records, err := api.GetDnsRecords(productObject)
	if err != nil {
		styles.WarnPrint(fmt.Sprintf("%s%s >> %v", printPrefix, productObject, err.Error()))
		return nil
	}

	return records
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
