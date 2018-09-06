package commands

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/mkfsn/mizukinana/concerts"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	ErrUnsupportedOutputType = errors.New("Unsupported output type")
)

type concertsCommand struct {
	*cobra.Command
	outputFlag string
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

	command.Flags().StringVarP(&command.outputFlag, "output", "o", "table", "output format: table, json, and yaml")

	return &command
}

func (c *concertsCommand) print(cmd *cobra.Command) {
	var result []byte
	var err error

	switch c.outputFlag {
	case "yaml":
		result, err = yaml.Marshal(concerts.Concerts)
	case "json":
		result, err = json.MarshalIndent(concerts.Concerts, "", "\t")
	case "table":
		result, err = concerts.Concerts.MarshalTable()
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
