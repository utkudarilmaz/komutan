package cmd

import (
	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var (
	log = logging.MustGetLogger("base")
)

var rootCmd = &cobra.Command{
	Use:   "komutan",
	Short: "Make your commit messages efficient",
	Long:  `Order the commit messages`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err.Error())
	}
}
