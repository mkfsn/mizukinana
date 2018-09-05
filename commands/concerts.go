package commands

import (
	"encoding/json"
	"errors"
	"fmt"
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
	switch c.outputFlag {
	case "yaml":
		c.printYAML(cmd)
	case "json":
		c.printJSON(cmd)
	case "table":
		c.printTable(cmd)
	default:
		cmd.Printf("Error: %s\n", ErrUnsupportedOutputType.Error())
		cmd.Usage()
	}
}

func (c *concertsCommand) printJSON(cmd *cobra.Command) {
	bytes, err := json.MarshalIndent(concerts.Concerts, "", "\t")
	if err != nil {
		cmd.Printf("Error:", err.Error())
		cmd.Usage()
		os.Exit(-1)
	}
	fmt.Printf("%s", string(bytes))
}

func (c *concertsCommand) printYAML(cmd *cobra.Command) {
	result, err := yaml.Marshal(concerts.Concerts)
	if err != nil {
		cmd.Printf("Error:", err.Error())
		cmd.Usage()
		os.Exit(-1)
	}
	fmt.Printf("%s", string(result))
}

func (l *concertsCommand) printTable(cmd *cobra.Command) {
	bytes, err := concerts.Concerts.MarshalTable()
	if err != nil {
		cmd.Printf("Error:", err.Error())
		cmd.Usage()
		os.Exit(-1)
	}
	fmt.Printf("%s", string(bytes))
}
