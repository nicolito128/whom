/*
Package cli
*/
package cli

import (
	"fmt"
	"log"
	"os"
)

var subcommands = map[string]Command{
	initCmd.Name(): initCmd,
}

func Run() {
	if len(os.Args[1:]) < 1 {
		log.Println("[INFO] You need to specify a command.")
		printUsage()
		return
	}

	if os.Args[1] == "help" {
		printUsage()
		return
	}

	cmd := subcommands[os.Args[1]]
	if cmd == nil {
		log.Fatalf("[ERROR] unknown subcommand '%s', see help for more details.", os.Args[1])
	}
	cmd.Setup()

	if err := cmd.FlagSet().Parse(os.Args[2:]); err != nil {
		log.Fatalf("[ERROR] %v", err)
	}

	if err := cmd.Exec(); err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
}

func printUsage() {
	s := "\nUsage:"
	s += `
	whom [ COMMAND ] [ OPTIONS ]
	`
	fmt.Println(s)
}
