/*
Package gens provides code generation utilities for creating project structures.
*/
package gens

import (
	_ "embed"
	"fmt"
	"log"
	"os"
)

//go:embed resources/config.toml
var baseConfigToml string

//go:embed resources/pods.toml
var basePodsToml string

func GenerateBaseProject(name string) error {
	// Checking if directory already exists and is not empty
	if info, err := os.Stat(name); err == nil && info.IsDir() {
		if entries, err := os.ReadDir(name); err == nil && len(entries) > 0 {
			return fmt.Errorf("directory %s already exists and is not empty", name)
		}
	}

	log.Println("Generating base project:", name)

	dir := fmt.Sprintf("./%s/pods", name)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	whomPath := fmt.Sprintf("./%s/.whom", name)
	configPath := fmt.Sprintf("./%s/config.toml", name)
	podsPath := fmt.Sprintf("./%s/pods.toml", name)

	if err := os.WriteFile(whomPath, []byte{}, 0o644); err != nil {
		return fmt.Errorf("failed to write .whom file: %w", err)
	}
	if err := os.WriteFile(configPath, []byte(baseConfigToml), 0o644); err != nil {
		return fmt.Errorf("failed to write config.toml: %w", err)
	}
	if err := os.WriteFile(podsPath, []byte(basePodsToml), 0o644); err != nil {
		return fmt.Errorf("failed to write pods.toml: %w", err)
	}

	return nil
}
