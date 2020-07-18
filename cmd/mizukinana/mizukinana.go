package main

import (
	"fmt"
	"os"

	"github.com/mkfsn/mizukinana/cmd/mizukinana/concerts"
	"github.com/spf13/cobra"

	"github.com/mkfsn/mizukinana/cmd/mizukinana/nanaparty"
	"github.com/mkfsn/mizukinana/cmd/mizukinana/version"
)

const (
	banner = `
##     ## #### ######## ##     ## ##    ## ####    ##    ##    ###    ##    ##    ###
###   ###  ##       ##  ##     ## ##   ##   ##     ###   ##   ## ##   ###   ##   ## ##
#### ####  ##      ##   ##     ## ##  ##    ##     ####  ##  ##   ##  ####  ##  ##   ##
## ### ##  ##     ##    ##     ## #####     ##     ## ## ## ##     ## ## ## ## ##     ##
##     ##  ##    ##     ##     ## ##  ##    ##     ##  #### ######### ##  #### #########
##     ##  ##   ##      ##     ## ##   ##   ##     ##   ### ##     ## ##   ### ##     ##
##     ## #### ########  #######  ##    ## ####    ##    ## ##     ## ##    ## ##     ##
`
)

type Options struct {
	Verbose bool
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Run() error {
	fmt.Printf("%s\n", banner)
	return nil
}

func NewCmdMizukinana() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "mizukinana",
		Short: "mizukinana is a Command-Line tool for providing some information of Mizuki Nana",
		Run: func(cmd *cobra.Command, args []string) {
			_ = o.Run()
			_ = cmd.Usage()
		},
	}
	cmd.PersistentFlags().BoolVarP(&o.Verbose, "verbose", "v", false, "verbose output")
	cmd.AddCommand(version.NewCmdVersion())
	cmd.AddCommand(nanaparty.NewCmdNanaparty())
	cmd.AddCommand(concerts.NewCmdConcerts())
	return cmd
}

func main() {
	cmd := NewCmdMizukinana()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
