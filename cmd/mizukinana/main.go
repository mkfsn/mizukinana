package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mizukinana",
	Short: "mizukinana is a Command-Line tool for providing some information of Mizuki Nana",
	Run: func(cmd *cobra.Command, args []string) {
		printBanner()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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