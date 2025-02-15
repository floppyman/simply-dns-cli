package add

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/umbrella-sh/um-common/ext"
	log "github.com/umbrella-sh/um-common/logging/basic"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/forms"
)

//goland:noinspection GoNameStartsWithPackageName
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new DNS record to a Domain",
	Args:  handleArgs,
	Run:   cmdRun,
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func cmdRun(_ *cobra.Command, _ []string) {
	log.Infoln("Add new dns record")

	form := forms.New()
	form.Add(forms.NewTextField("name", "Name", true, validateName), "")
	form.Add(forms.NewIntegerField("ttl", "TTL", true, validateTtl), int((1 * time.Hour).Seconds()))
	form.Render()

	values := form.GetValues()
	log.Info(values)
}

func validateName(val string, isRequired bool) (ok bool, msg string, res string) {
	if !isRequired && val == "" {
		return true, "", ""
	}

	if val == "" {
		return false, "A name is required", ""
	}

	//goland:noinspection HttpUrlsUsage
	if strings.Contains(val, "http://") || strings.Contains(val, "https://") {
		return false, "Name should be the domain only, without 'http(s)://'", ""
	}

	if valid, _ := ext.IsValidUrl(val); !valid {
		return false, "Name is not a valid url", ""
	}

	return true, "", val
}

func validateTtl(val string, isRequired bool, converter forms.IntegerConverter) (ok bool, msg string, res int) {
	converted, cVal := converter(val)
	if !converted {
		return false, "TTL is not a valid integer", 0
	}

	dnsTtl := api.DnsRecordTTL(cVal)
	if !(dnsTtl == api.DnsRecTTLMin10 || dnsTtl == api.DnsRecTTLHour1 || dnsTtl == api.DnsRecTTLHours6 || dnsTtl == api.DnsRecTTLHours12 || dnsTtl == api.DnsRecTTLHours24) {
		return false, fmt.Sprintf("TTL must be one of these values, 10 min:%d, 1 hour:%d, 6 hours:%d, 12 hours:%d, 24 hours:%d", api.DnsRecTTLMin10, api.DnsRecTTLHour1, api.DnsRecTTLHours6, api.DnsRecTTLHours12, api.DnsRecTTLHours24), 0
	}

	return true, "", int(dnsTtl)
}
