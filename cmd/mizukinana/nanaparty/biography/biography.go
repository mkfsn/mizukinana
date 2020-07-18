package biography

import (
	"context"
	"fmt"
	"os"
	"time"

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

	biography, err := collection.Biography(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return err
	}

	// TODO(mkfsn): output profile
	biography.Profile()

	return nil
}

func NewCmdBiography() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use: "biography",
		Run: func(cmd *cobra.Command, args []string) {
			_ = o.Run()
		},
	}
	cmd.Flags().DurationVar(&o.MaxTime, "-max-time", 10*time.Second, "Maximum time allowed for the request")
	return cmd
}
