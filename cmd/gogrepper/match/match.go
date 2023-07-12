package match

import (
	"github.com/TheisFerre/go-grepper/cmd/gogrepper/match/file"
	"github.com/TheisFerre/go-grepper/cmd/gogrepper/match/pipe"
	"github.com/spf13/cobra"
)

var grepMode bool
var MatchCmd = &cobra.Command{
	Use:     "match",
	Aliases: []string{"match"},
	Short:   "grep like tool. Use either 'pipe' or 'file' for different behaviours",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	MatchCmd.Flags().BoolVarP(&grepMode, "grep-mode", "g", false, "Pipe input to tool")
	MatchCmd.AddCommand(pipe.PipeCmd)
	MatchCmd.AddCommand(file.FipeCmd)
}
