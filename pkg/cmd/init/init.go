/*
Package init provides a command to initialize a new repository.
*/
package init

import (
	"fmt"
	"strings"

	"github.com/nicolito128/whom/internal/gens"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [repository name]",
		Short: "Initialize a new Podman Quadlet repository",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			if strings.TrimSpace(name) == "" {
				return fmt.Errorf("repository name cannot be empty")
			}

			fmt.Printf("Initializing new repository: %s\n", name)
			if err := gens.GenerateBaseProject(name); err != nil {
				return err
			}

			fmt.Println("Repository initialized successfully!")
			return nil
		},
	}
	return cmd
}
