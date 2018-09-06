package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mizukinana",
	Short: "mizukinana is a Command-Line tool for providing some information of Mizuki Nana",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		printBanner()
	},
}

func init() {
	var verbose bool
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(newProfileCommand().Command)
	rootCmd.AddCommand(newVersionCommand().Command)
	rootCmd.AddCommand(newConcertsCommand().Command)
}

func printBanner() {
	banner := []byte(
		`
##     ## #### ######## ##     ## ##    ## ####    ##    ##    ###    ##    ##    ###
###   ###  ##       ##  ##     ## ##   ##   ##     ###   ##   ## ##   ###   ##   ## ##
#### ####  ##      ##   ##     ## ##  ##    ##     ####  ##  ##   ##  ####  ##  ##   ##
## ### ##  ##     ##    ##     ## #####     ##     ## ## ## ##     ## ## ## ## ##     ##
##     ##  ##    ##     ##     ## ##  ##    ##     ##  #### ######### ##  #### #########
##     ##  ##   ##      ##     ## ##   ##   ##     ##   ### ##     ## ##   ### ##     ##
##     ## #### ########  #######  ##    ## ####    ##    ## ##     ## ##    ## ##     ##
		`,
	)
	fmt.Println(string(banner))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
