package list

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/umbrella-sh/simply-dns-cli/internal/collectors"
	"github.com/umbrella-sh/simply-dns-cli/internal/objects"
	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

func cmdRun(_ *cobra.Command, _ []string) {
	styles.Println(styles.Info("Listing DNS Record from Domain"))
	styles.Blank()

	cancelled, domain := collectors.CollectDomain(options.Domain)
	if cancelled {
		printCancelText()
		return
	}
	styles.Blank()

	records := shared.PullDnsRecords(domain, "")
	if records == nil {
		return
	}
	styles.Blank()

	rows := make([][]string, 0)
	for _, v := range records {
		v1 := strconv.FormatInt(v.RecordId, 10)
		v2 := objects.DnsTypeToText(v.Type)
		v3 := v.Name
		v4 := v.Data
		v5 := objects.DnsTTLToNumberText(v.TTL)
		v6 := ""
		if v.Priority != nil {
			v6 = v.Priority.ToString()
		}
		v7 := v.Comment

		rows = append(rows, []string{v1, v2, v3, v4, v5, v6, v7})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"#",
		"Type",
		"Name",
		"Data",
		"TTL",
		"Priority",
		"Comment",
	})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetTablePadding(" ")
	table.SetCenterSeparator(styles.Graphic("|"))
	table.SetColumnSeparator(styles.Graphic("|"))
	table.SetRowSeparator(styles.Graphic("-"))
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(rows) // Add Bulk Data
	table.Render()
}

func printCancelText() { styles.Println(styles.Warn("\nList was cancelled\n")) }
