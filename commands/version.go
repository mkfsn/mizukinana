package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "0.1"
)

type versionCommand struct {
	*cobra.Command
	shortFlag bool
}

func newVersionCommand() *versionCommand {
	var command versionCommand

	command.Command = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of mizukinana",
		Long:  `All software has versions. This is mizukinana's`,
		Run: func(cmd *cobra.Command, args []string) {
			command.print()
		},
	}

	command.Flags().BoolVarP(&command.shortFlag, "short", "s", false, "only print version")

	return &command
}

func (v *versionCommand) print() {
	if v.shortFlag {
		fmt.Printf("v%s", version)
		return
	}
	fmt.Printf("mizukinana version v%s\n", version)
}
