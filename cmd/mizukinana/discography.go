package main

import (
	"encoding/json"
	"os"

	"github.com/mkfsn/mizukinana/discography"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type discographyCommand struct {
	*cobra.Command
	output string
}

func newDiscographyCommand() *discographyCommand {
	var command discographyCommand

	command.Command = &cobra.Command{
		Use:   "discography",
		Short: "Print all Mizuki Nana's discography, e.g. Singles, Albums",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			command.print(cmd)
		},
	}

	command.Flags().StringVarP(&command.output, "output", "o", "table", "output format: table, json, and yaml")

	return &command
}

func (d *discographyCommand) print(cmd *cobra.Command) {
	all := discography.All

	var result []byte
	var err error

	switch d.output {
	case "yaml":
		result, err = yaml.Marshal(all)
	case "json":
		result, err = json.MarshalIndent(all, "", "\t")
	case "table":
		result, err = all.MarshalTable()
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
