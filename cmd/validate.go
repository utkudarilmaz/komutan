package cmd

import (
	"komutan/commit"
	"os"

	"github.com/spf13/cobra"
)

var (
	Message string
	File    string
)

func init() {
	validate.Flags().StringVarP(
		&Message,
		"message",
		"m",
		"",
		"commit message",
	)
	validate.Flags().StringVarP(
		&File,
		"file",
		"f",
		"",
		"commit message file",
	)
	rootCmd.AddCommand(validate)
}

var validate = &cobra.Command{
	Use:   "validate",
	Short: "validate the given commit message",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		if len(Message) > 0 {
			err := commit.ValidateCommitMsgString(Message)
			if err != nil {
				log.Errorf(err.Error())
				os.Exit(1)
			}
		} else {
			err := commit.ValidateCommitMsgFile(File)
			if err != nil {
				log.Errorf(err.Error())
				os.Exit(1)

			}
		}
	},
}
