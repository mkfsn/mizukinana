package commands

import (
	"encoding/json"
	"os"

	"github.com/mkfsn/mizukinana/concerts"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type concertsCommand struct {
	*cobra.Command
	output string
	filter string
}

func newConcertsCommand() *concertsCommand {
	var command concertsCommand

	command.Command = &cobra.Command{
		Use:   "concerts",
		Short: "Print all Mizuki Nana's concerts",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			command.print(cmd)
		},
	}

	command.Flags().StringVarP(&command.output, "output", "o", "table", "output format: table, json, and yaml")
	command.Flags().StringVarP(&command.filter, "filter", "f", "", "filtering the concerts")

	return &command
}

func (c *concertsCommand) print(cmd *cobra.Command) {
	concerts := concerts.PersonalConcerts
	if c.filter != "" {
		concerts = concerts.Filter(c.filter)
	}

	var result []byte
	var err error

	switch c.output {
	case "yaml":
		result, err = yaml.Marshal(concerts)
	case "json":
		result, err = json.MarshalIndent(concerts, "", "\t")
	case "table":
		result, err = concerts.MarshalTable()
	default:
		cmd.Printf("Error: %s\n", ErrUnsupportedOutputType.Error())
		cmd.Usage()
	}

	if err != nil {
		cmd.Printf("Error:", err.Error())
		os.Exit(-1)
	}

	cmd.Printf("%s\n", string(result))
}
