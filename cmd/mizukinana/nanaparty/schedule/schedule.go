package schedule

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Options struct{}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Run() error {
	fmt.Println("Not implemented yet")
	return nil
}

func NewCmdSchedule() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use: "schedule",
		Run: func(cmd *cobra.Command, args []string) {
			_ = o.Run()
		},
	}
	return cmd
}
