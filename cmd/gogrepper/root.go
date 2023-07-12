package gogrepper

import (
	"fmt"
	"os"

	"github.com/TheisFerre/go-grepper/cmd/gogrepper/match"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "go-grepper",
	Short: "go-grepper",
	Long:  `go-grepper`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(match.MatchCmd)
}
