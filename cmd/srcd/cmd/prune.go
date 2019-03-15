package cmd

import (
	"github.com/spf13/cobra"
	"github.com/src-d/engine/components"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Removes all resources used by engine.",
	Run: func(cmd *cobra.Command, args []string) {
		withImages, _ := cmd.Flags().GetBool("with-images")

		if err := components.Prune(withImages); err != nil {
			fatal(err, "cound not prune components")
		}
	},
}

func init() {
	rootCmd.AddCommand(pruneCmd)

	pruneCmd.Flags().Bool("with-images", false, "remove docker images")
}
