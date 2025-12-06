/*
Package whom provides functionalities to manage whom repositories.
*/
package whom

import (
	"fmt"
	"os"
)

type Repository struct {
	Name        string   `toml:"name"`
	Description string   `toml:"description"`
	URL         string   `toml:"url"`
	Author      string   `toml:"author"`
	Maintainers []string `toml:"maintainers"`
}

type Config struct {
	Repo Repository `toml:"repository"`
}

// IsValidRepository checks if the current directory is a valid whom repository.
func IsValidRepository() error {
	if _, err := os.Stat(".whom"); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("not a valid whom repository: .whom file not found")
		}
	}
	if stat, err := os.Stat("pods"); err != nil || !stat.IsDir() {
		return fmt.Errorf("not a valid whom repository: pods/ directory not found")
	}
	if _, err := os.Stat("config.toml"); err != nil {
		return fmt.Errorf("not a valid whom repository: config.toml file not found")
	}
	if _, err := os.Stat("pods.toml"); err != nil {
		return fmt.Errorf("not a valid whom repository: pods.toml file not found")
	}
	return nil
}
