package concerts

import (
	"encoding/json"
	"os"

	"github.com/mkfsn/mizukinana/concerts"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	Output string
	Filter string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Run(cmd *cobra.Command) {
	concerts := concerts.PersonalConcerts
	if o.Filter != "" {
		concerts = concerts.Filter(o.Filter)
	}

	var result []byte
	var err error
	switch o.Output {
	case "yaml":
		result, err = yaml.Marshal(concerts)
	case "json":
		result, err = json.MarshalIndent(concerts, "", "\t")
	case "table":
		result, err = concerts.MarshalTable()
	default:
		cmd.Printf("Error: %s\n", errUnsupportedOutputType.Error())
		cmd.Usage()
	}
	if err != nil {
		cmd.Printf("Error: %s", err)
		os.Exit(-1)
	}
	cmd.Printf("%s\n", string(result))
}

func NewCmdConcerts() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "concerts",
		Short: "Print all Mizuki Nana's concerts",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run(cmd)
		},
	}
	cmd.Flags().StringVarP(&o.Output, "output", "o", "table", "output format: table, json, and yaml")
	cmd.Flags().StringVarP(&o.Filter, "filter", "f", "", "filtering the concerts")
	return cmd
}
