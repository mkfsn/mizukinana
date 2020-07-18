package top

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

func NewCmdTop() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use: "top",
		Run: func(cmd *cobra.Command, args []string) {
			_ = o.Run()
		},
	}
	return cmd
}
