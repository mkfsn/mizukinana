package version

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const (
	// TODO(mkfsn) use commit version or tag version
	version = "0.1"
)

type Version struct {
	Version string `json:"version" yaml:"version"`
}

func NewVersion() Version {
	return Version{Version: version}
}

func (v Version) Short() string {
	return fmt.Sprintf("v%s", version)
}

func (v Version) Long() string {
	return fmt.Sprintf("mizukinana version v%s\n", version)
}

type Options struct {
	Short  bool
	Output string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Run() error {
	v := NewVersion()
	switch o.Output {
	case "":
		if o.Short {
			fmt.Printf("%s", v.Short())
		} else {
			fmt.Printf("%s", v.Long())
		}
	case "json":
		result, err := json.Marshal(v)
		if err != nil {
			return err
		}
		fmt.Printf("%s", result)
	case "yaml":
		result, err := yaml.Marshal(v)
		if err != nil {
			return err
		}
		fmt.Printf("%s", result)
	default:
		return fmt.Errorf("VersionOptions is invalid: --output:%q is not accpectable", o.Output)
	}
	return nil
}

func NewCmdVersion() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of this tool (mizukinana)",
		Long:  `All software has versions. This is mizukinana's`,
		Run: func(cmd *cobra.Command, args []string) {
			o.Run()
		},
	}
	cmd.Flags().BoolVarP(&o.Short, "short", "s", false, "If true, only print version number.")
	cmd.Flags().StringVarP(&o.Output, "output", "o", "", "One of 'yaml' or 'json'")
	return cmd
}
