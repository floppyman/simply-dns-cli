package pull

import "github.com/spf13/cobra"

//goland:noinspection GoNameStartsWithPackageName
var PullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls all the current Products and DNS Records from Simply.com DNS",
	Args:  handleCommandArguments,
	Run:   pullCmdRun,
}

func handleCommandArguments(cmd *cobra.Command, args []string) error {
	return nil
}

func pullCmdRun(_ *cobra.Command, _ []string) {

}
