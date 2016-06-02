package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of the Kubernetes Cloud Foundry cc-uploader.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ccUploaderVersion)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
