/*
Package root implements the root command for the whom CLI tool.
*/
package root

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "whom",
		Short: "Whom is a tool to manage Podman Quadlet files",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()
			}
			return nil
		},
	}
}
