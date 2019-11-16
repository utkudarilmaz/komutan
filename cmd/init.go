package cmd

import (
	"komutan/initialize"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initilize the project.",
	Run: func(cmd *cobra.Command, args []string) {
		err := initialize.Init()
		if err != nil {
			log.Errorf(err.Error())
			os.Exit(1)
		}
		log.Notice("Project initialized")
	},
}
