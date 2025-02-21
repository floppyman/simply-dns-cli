package shared

import (
	"fmt"

	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

func PrintValue(header string, val string) {
	fmt.Printf("%s %s\n", styles.Header(header), val)
}
