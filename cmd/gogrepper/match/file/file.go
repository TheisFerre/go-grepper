package file

import (
	"github.com/TheisFerre/go-grepper/pkg/gogrepper"
	"github.com/spf13/cobra"
)

var grepMode bool
var FipeCmd = &cobra.Command{
	Use:     "file",
	Aliases: []string{"file"},
	Short:   "Reads file",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		fileName := args[0]
		pattern := args[1]

		gogrepper.ReadFile(fileName, pattern)

	},
}

func init() {

}
