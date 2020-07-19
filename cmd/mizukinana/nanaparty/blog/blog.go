package blog

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mkfsn/mizukinana/cmd/mizukinana/utils"
	"github.com/mkfsn/mizukinana/cmd/mizukinana/utils/table"
	"github.com/mkfsn/mizukinana/nanaparty"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

type Options struct {
	MaxTime time.Duration
	Latest  int
	Output  string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Run() {
	collection := nanaparty.New()

	ctx, cancel := context.WithTimeout(context.Background(), o.MaxTime)
	defer cancel()

	blogs, err := collection.Blog(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	if o.Latest >= 0 {
		blogs = blogs.Latest(o.Latest)
	}

	var res []byte
	switch o.Output {
	case "", "table":
		res, err = table.Marshal(blogs)
	case "json":
		res, err = json.MarshalIndent(blogs, "", "    ")
	case "yaml":
		res, err = yaml.Marshal(blogs)
	default:
		err = fmt.Errorf("Invalid output: --output=%q is not acceptable\n", o.Output)
	}
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("%s\n", res)
}

func NewCmdBlog() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use: "blog",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run()
		},
	}
	cmd.Flags().DurationVar(&o.MaxTime, "max-time", utils.DefaultMaxTime, "Maximum time allowed for the request")
	cmd.Flags().StringVarP(&o.Output, "output", "o", "", "One of 'yaml', 'json', 'csv' or 'table'")
	cmd.Flags().IntVarP(&o.Latest, "", "n", -1, "Show only latest n blogs")
	return cmd
}
