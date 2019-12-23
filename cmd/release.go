package cmd

import (
	"komutan/release"

	"github.com/spf13/cobra"
)

var (
	Patch bool
	Minor bool
	Major bool
)

func init() {
	releaseAsCmd.Flags().BoolVarP(
		&Patch,
		"patch",
		"",
		false,
		"release as patch",
	)
	releaseAsCmd.Flags().BoolVarP(
		&Minor,
		"minor",
		"",
		false,
		"release as minor",
	)
	releaseAsCmd.Flags().BoolVarP(
		&Major,
		"major",
		"",
		false,
		"release as major",
	)
	rootCmd.AddCommand(releaseAsCmd)
}

var releaseAsCmd = &cobra.Command{
	Use:     "release-as",
	Short:   "create new tag at head of repo",
	Example: "komutan release-as patch # tags: v1.0.1 -> v1.0.1, v1.0.2",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if Patch {
			err := release.Patch()
			if err != nil {
				log.Errorf(err.Error())
			}
		} else if Minor {
			return
		} else if Major {
			return
		}

	},
}
