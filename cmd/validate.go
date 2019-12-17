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
	validateCmd.Flags().StringVarP(
		&Message,
		"message",
		"m",
		"",
		"commit message",
	)
	validateCmd.Flags().StringVarP(
		&File,
		"file",
		"f",
		"",
		"commit message file",
	)
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate the given commit message",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		if len(Message) > 0 {
			err := commit.ValidateCommitMsg(Message)
			if err != nil {
				log.Errorf(err.Error())
				os.Exit(1)
			}
		} else {
			err := commit.ValidateCommitMsgFromFile(File)
			if err != nil {
				log.Errorf(err.Error())
				os.Exit(1)

			}
		}
	},
}
