package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "latest"

func NewVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	return cmd
}
