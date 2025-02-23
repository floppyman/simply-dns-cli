package update

import "github.com/spf13/cobra"

//goland:noinspection GoNameStartsWithPackageName
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing DNS record on a Domain.",
	Args:  handleArgs,
	Run:   cmdRun,
}

type updateOptions struct {
	Domain   string
	RecordId int64
	Type     string
	TTL      int
	Name     string
	Data     string
	Priority int
	Comment  string
}

var options = updateOptions{
	Domain:   "",
	RecordId: 0,
	Type:     "",
	TTL:      0,
	Name:     "",
	Data:     "",
	Priority: 0,
	Comment:  "",
}

const NoCommentValue = "<{NO_COMMENT}>"

func init() {
	UpdateCmd.Flags().StringVarP(&options.Domain, "domain", "d", "", "TLD name to remove the record from, ex: domain.com")
	UpdateCmd.Flags().Int64VarP(&options.RecordId, "record", "r", 0, "Id of the record to remove")
	UpdateCmd.Flags().StringVarP(&options.Type, "type", "t", "", "Type of record to create")
	UpdateCmd.Flags().IntVarP(&options.TTL, "ttl", "l", 0, "TTL of record (time to live)")
	UpdateCmd.Flags().StringVarP(&options.Name, "name", "n", "", "Name of the record (sub domain), ex: 'example.domain.com' but without '.domain.com'")
	UpdateCmd.Flags().StringVarP(&options.Data, "data", "v", "", "Data of the record, ex: Destination IP")
	UpdateCmd.Flags().IntVarP(&options.Priority, "priority", "p", 0, "Priority of the MX record, only used with MX type else ignored")
	UpdateCmd.Flags().StringVarP(&options.Comment, "comment", "c", NoCommentValue, "Comment for the record")
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}
