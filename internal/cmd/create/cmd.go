package create

import "github.com/spf13/cobra"

//goland:noinspection GoNameStartsWithPackageName
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new DNS record to a Domain",
	Args:  handleArgs,
	Run:   cmdRun,
}

type createOptions struct {
	Domain   string
	Type     string
	TTL      int
	Name     string
	Data     string
	Priority int
	Comment  string
}

var options = createOptions{
	Domain:   "",
	Type:     "",
	TTL:      0,
	Name:     "",
	Data:     "",
	Priority: 0,
	Comment:  "",
}

const NoCommentValue = "<{NO_COMMENT}>"

func init() {
	CreateCmd.Flags().StringVarP(&options.Domain, "domain", "d", "", "TLD name to create record under, ex: domain.com")
	CreateCmd.Flags().StringVarP(&options.Type, "type", "t", "", "Type of record to create")
	CreateCmd.Flags().IntVarP(&options.TTL, "ttl", "l", 0, "TTL of record (time to live)")
	CreateCmd.Flags().StringVarP(&options.Name, "name", "n", "", "Name of the record (sub domain), ex: 'example.domain.com' but without '.domain.com'")
	CreateCmd.Flags().StringVarP(&options.Data, "data", "v", "", "Data of the record, ex: Destination IP")
	CreateCmd.Flags().IntVarP(&options.Priority, "priority", "p", 0, "Priority of the MX record, only used with MX type else ignored")
	CreateCmd.Flags().StringVarP(&options.Comment, "comment", "c", NoCommentValue, "Comment for the record")
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}
