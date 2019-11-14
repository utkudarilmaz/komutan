package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initilize the project.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init")
		return
	},
}
