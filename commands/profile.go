package commands

import (
	"encoding/json"
	"os"

	"github.com/mkfsn/mizukinana/profile"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type profileCommand struct {
	*cobra.Command
	outputFlag string
}

func newProfileCommand() *profileCommand {
	var command profileCommand

	command.Command = &cobra.Command{
		Use:   "profile",
		Short: "Print Mizuki Nana's profile",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			command.print(cmd)
		},
	}

	command.Flags().StringVarP(&command.outputFlag, "output", "o", "table", "output format: table, json, and yaml")

	return &command
}

func (p *profileCommand) print(cmd *cobra.Command) {
	var result []byte
	var err error

	data := profile.MizukiNana
	switch p.outputFlag {
	case "yaml":
		result, err = yaml.Marshal(data)
	case "json":
		result, err = json.MarshalIndent(data, "", "\t")
	case "table":
		result, err = data.MarshalTable()
	default:
		cmd.Printf("Error: %s\n", errUnsupportedOutputType.Error())
		cmd.Usage()
	}

	if err != nil {
		cmd.Printf("Error:", err.Error())
		os.Exit(-1)
	}

	cmd.Printf("%s\n", string(result))
}
