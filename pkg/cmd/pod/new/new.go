/*
Package new implements the 'pod new' command for creating new pods
*/
package new

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/nicolito128/whom/pkg/whom"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "new [pod name]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := whom.IsValidRepository(); err != nil {
				return err
			}

			podName := args[0]
			if strings.TrimSpace(podName) == "" {
				return fmt.Errorf("pod name cannot be empty")
			}

			typ, err := handlePodTypeFlags(cmd)
			if err != nil {
				return err
			}

			fmt.Printf("Creating new pod: %s\n", podName)
			if err = whom.CreatePod(podName, typ); err != nil {
				return err
			}
			fmt.Println("Pod created successfully!")

			if err = handleEdit(cmd, podName, typ); err != nil {
				return err
			}

			return nil
		},
	}
	cmd.Flags().BoolP("command", "c", false, "Create a command pod")
	cmd.Flags().BoolP("compose", "m", false, "Create a compose pod")
	cmd.Flags().BoolP("edit", "e", false, "If the pod is created, open it with $EDITOR")
	return cmd
}

func handlePodTypeFlags(cmd *cobra.Command) (whom.PodType, error) {
	var typ whom.PodType

	commandFlag, err := cmd.Flags().GetBool("command")
	if err != nil {
		return typ, err
	}

	composeFlag, err := cmd.Flags().GetBool("compose")
	if err != nil {
		return typ, err
	}

	if !commandFlag && !composeFlag {
		return typ, fmt.Errorf("must specify either --command or --compose flag")
	}
	if commandFlag && composeFlag {
		return typ, fmt.Errorf("cannot specify both --command and --compose flags")
	}

	if commandFlag {
		typ = whom.CommandPod
	} else {
		typ = whom.ComposePod
	}
	return typ, nil
}

func handleEdit(cmd *cobra.Command, podName string, typ whom.PodType) error {
	editFlag, err := cmd.Flags().GetBool("edit")
	if err != nil {
		return err
	}
	if !editFlag {
		return nil
	}

	editorEnv, ok := os.LookupEnv("EDITOR")
	if ok && editorEnv != "" {
		file := "command"
		if typ == whom.ComposePod {
			file = "compose.yml"
		}

		editorCmd := strings.Fields(editorEnv)
		if len(editorCmd) == 0 {
			return fmt.Errorf("invalid EDITOR environment variable")
		}

		_, err := exec.LookPath(editorCmd[0])
		if err != nil {
			return fmt.Errorf("editor '%s' not found in PATH", editorCmd[0])
		}

		cmdArgs := append(editorCmd[1:], fmt.Sprintf("pods/%s/%s", podName, file))
		cmd := exec.Command(editorCmd[0], cmdArgs...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			return fmt.Errorf("failed to start editor: %w", err)
		}

		// Wait for the exit signal (blocking call)
		if err := cmd.Wait(); err != nil {
			return fmt.Errorf("editor exited with error: %w", err)
		}
	}
	return nil
}
