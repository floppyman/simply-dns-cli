package push

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/umbrella-sh/um-common/logging/ulog"
	"github.com/umbrella-sh/um-common/utils"

	"github.com/umbrella-sh/simply-dns-sync/internal/api"
	"github.com/umbrella-sh/simply-dns-sync/internal/configs"
)

//goland:noinspection GoNameStartsWithPackageName
var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes all the DNS Record changes to Simply.com DNS",
	Args:  handleCommandArguments,
	Run:   pushCmdRun,
}

func handleCommandArguments(cmd *cobra.Command, args []string) error {
	return nil
}

func pushCmdRun(_ *cobra.Command, _ []string) {
	currentState := pullCurrentState()
	changes := compareCurrentWithLocalState(currentState, configs.Sync.Products)
	for _, change := range changes {
		printChange(change)
	}
}

func printChange(change DnsChange) {
	switch change.Type {
	case CtAdded:
		ulog.Console.Debug().Msgf("ADDED   | %s | %s | %s | %d", change.NewRecord.Type, change.NewRecord.Name, change.NewRecord.Data, change.NewRecord.Ttl)
		break
	case CtUpdated:
		ulog.Console.Debug().Msgf("UPDATED | %s | %s | %s | %d", change.NewRecord.Type, change.NewRecord.Name, change.NewRecord.Data, change.NewRecord.Ttl)
		break
	case CtDeleted:
		ulog.Console.Debug().Msgf("DELETED | %s | %s | %s | %d", change.NewRecord.Type, change.NewRecord.Name, change.NewRecord.Data, change.NewRecord.Ttl)
		break
	}
}

func pullCurrentState() []*api.SimplyProduct {
	ulog.Console.Info().Msg("Pulling products from account...")
	products, err := api.GetProducts()
	if err != nil {
		ulog.Console.Err(err).Msg("Failed to get products")
		return nil
	}
	ulog.Console.Info().Msg(utils.Green("Done"))

	ulog.Console.Info().Msg("Pulling dns records from each product...")
	for _, product := range products {
		ulog.Console.Info().Msgf("  %s", product.Name)
		records, err := api.GetDnsRecords(product.Object)
		if err != nil {
			ulog.Console.Err(err).Msgf("  Failed to get dns records for %s", product.Name)
			product.DnsRecords = make([]*api.SimplyDnsRecord, 0)
			continue
		}

		product.DnsRecords = records
	}
	ulog.Console.Info().Msg(utils.Green("Done"))

	return products
}

func compareCurrentWithLocalState(remote []*api.SimplyProduct, local []*api.SimplyProduct) []DnsChange {
	remoteMap := make(map[string][]*api.SimplyDnsRecord)
	localMap := make(map[string][]*api.SimplyDnsRecord)

	for _, remoteItem := range remote {
		if _, ok := remoteMap[remoteItem.Object]; !ok {
			remoteMap[remoteItem.Object] = make([]*api.SimplyDnsRecord, 0)
		}
		remoteMap[remoteItem.Object] = remoteItem.DnsRecords
	}

	for _, localItem := range local {
		if _, ok := localMap[localItem.Object]; !ok {
			localMap[localItem.Object] = make([]*api.SimplyDnsRecord, 0)
		}
		localMap[localItem.Object] = localItem.DnsRecords
	}

	res := make([]DnsChange, 0)

	// Added
	for key, localItems := range localMap {
		if _, ok := remoteMap[key]; !ok {
			continue
		}

		for _, localItem := range localItems {
			localUniq := fmt.Sprintf("%s %s %s", localItem.Type, localItem.Name, localItem.Data)

			found := false
			remoteItems := remoteMap[key]
			for _, remoteItem := range remoteItems {
				remoteUniq := fmt.Sprintf("%s %s %s", remoteItem.Type, remoteItem.Name, remoteItem.Data)

				if localUniq == remoteUniq {
					found = true
				}
			}

			if !found {
				res = append(res, DnsChange{
					Type:          CtAdded,
					ProductObject: key,
					OldRecord:     nil,
					NewRecord:     localItem,
				})
			}
		}
	}

	// Updated

	// Deleted

	return res
}

func matchRecords(remote *api.SimplyDnsRecord, local *api.SimplyDnsRecord) bool {
	if remote.Name != local.Name {
		return true
	}
	if remote.Ttl != local.Ttl {
		return true
	}
	if remote.Data != local.Data {
		return true
	}
	if remote.Type != local.Type {
		return true
	}
	if remote.Priority != local.Priority {
		return true
	}
	if remote.Comment != local.Comment {
		return true
	}

	return false
}

type DnsChange struct {
	Type          ChangeType
	ProductObject string
	OldRecord     *api.SimplyDnsRecord
	NewRecord     *api.SimplyDnsRecord
}

type ChangeType int

const (
	CtAdded ChangeType = iota
	CtUpdated
	CtDeleted
)
