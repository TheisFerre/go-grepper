package pipe

import (
	"os"

	"github.com/TheisFerre/go-grepper/pkg/gogrepper"
	"github.com/spf13/cobra"
)

var grepMode bool
var PipeCmd = &cobra.Command{
	Use:     "pipe",
	Aliases: []string{"pipe"},
	Short:   "Reads input from a pipe",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		pattern := args[0]

		if gogrepper.IsInputFromPipe() {
			gogrepper.PipeReader(os.Stdin, os.Stdout, pattern)
		} else {
			println("Data is not coming from a pipe")
			os.Exit(1)
		}

	},
}

func init() {
}
