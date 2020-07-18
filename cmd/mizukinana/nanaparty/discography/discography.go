package discography

import (
	"encoding/json"
	"os"

	"github.com/mkfsn/mizukinana/discography"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	Output string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Run(cmd *cobra.Command) {
	all := discography.All

	var result []byte
	var err error

	switch o.Output {
	case "yaml":
		result, err = yaml.Marshal(all)
	case "json":
		result, err = json.MarshalIndent(all, "", "\t")
	case "table":
		result, err = all.MarshalTable()
	default:
		cmd.Println("Error: unsupported output type")
		_ = cmd.Usage()
	}

	if err != nil {
		cmd.Printf("Error:", err.Error())
		os.Exit(-1)
	}

	cmd.Printf("%s\n", string(result))
}

func NewCmdDiscography() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "discography",
		Short: "Print all Mizuki Nana's discography, e.g. Singles, Albums",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run(cmd)
		},
	}
	cmd.Flags().StringVarP(&o.Output, "output", "o", "table", "One of 'table', 'json' or 'yaml'")
	return cmd
}
