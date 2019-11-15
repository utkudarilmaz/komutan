package cmd

import (
	"komutan/commit"
	"os"

	"github.com/spf13/cobra"
)

var (
	Message string
)

func init() {
	validateCmd.Flags().StringVarP(&Message, "message", "m", "", "commit message")
	validateCmd.MarkFlagRequired("message")
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate the given commit message",
	Run: func(cmd *cobra.Command, args []string) {
		err := commit.Validate(Message)
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
	},
}
