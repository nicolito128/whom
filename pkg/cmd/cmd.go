/*
Package cmd is responsible for defining and executing CLI commands for the application.
*/
package cmd

import (
	"os"

	initc "github.com/nicolito128/whom/pkg/cmd/init"
	"github.com/nicolito128/whom/pkg/cmd/pod"
	"github.com/nicolito128/whom/pkg/cmd/root"
)

var rootCmd = root.NewCommand()

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initc.NewCommand())
	rootCmd.AddCommand(pod.NewCommand())
}
