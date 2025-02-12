package push

import "github.com/spf13/cobra"

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

}
