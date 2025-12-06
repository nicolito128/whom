/*
Package pod provides commands to manage pods related services and configurations.
*/
package pod

import (
	newc "github.com/nicolito128/whom/pkg/cmd/pod/new"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pod",
		Short: "Manage pods related services and configurations",
	}
	cmd.AddCommand(newc.NewCommand())
	return cmd
}
