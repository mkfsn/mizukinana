package blog

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mkfsn/mizukinana/cmd/mizukinana/utils"
	"github.com/mkfsn/mizukinana/nanaparty"

	"github.com/spf13/cobra"
)

type Options struct {
	MaxTime time.Duration
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Run() error {
	collection := nanaparty.New()

	ctx, cancel := context.WithTimeout(context.Background(), o.MaxTime)
	defer cancel()

	blogs, err := collection.Blog(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return err
	}

	for _, blog := range blogs {
		fmt.Printf("%v\n", blog.Info())
	}
	return nil
}

func NewCmdBlog() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use: "blog",
		Run: func(cmd *cobra.Command, args []string) {
			_ = o.Run()
		},
	}
	cmd.Flags().DurationVar(&o.MaxTime, "max-time", utils.DefaultMaxTime, "Maximum time allowed for the request")
	return cmd
}
