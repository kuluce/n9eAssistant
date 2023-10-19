package command

import (
	"fmt"
	"n9eHelper/assets"

	"github.com/spf13/cobra"
)

var (
	Name        = "n9eHelper"
	DisplayName = "n9eHelper"
	Description = "n9eHelper is a tool for n9e"
	VersionCmd  = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of n9eHelper",
		Long:  `All software has versions. This is n9eHelper's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("n9eHelper %s\n", assets.Version)
		},
	}
)
