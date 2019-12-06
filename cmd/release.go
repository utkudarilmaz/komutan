package cmd

import (
	"komutan/release"

	"github.com/spf13/cobra"
)

var (
	Pre   bool
	Patch bool
	Minor bool
	Major bool
)

func init() {
	releaseCmd.Flags().BoolVarP(
		&Patch,
		"patch",
		"",
		false,
		"patch release",
	)
	releaseCmd.Flags().BoolVarP(
		&Minor,
		"minor",
		"",
		false,
		"minor release",
	)
	releaseCmd.Flags().BoolVarP(
		&Major,
		"major",
		"",
		false,
		"major release",
	)
	rootCmd.AddCommand(releaseCmd)
}

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "create new release at current commit",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := release.Patch()
		if err != nil {
			log.Errorf(err.Error())
		}
	},
}
