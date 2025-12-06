package whom

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pelletier/go-toml/v2"
)

var (
	baseCommandFile = fmt.Appendf([]byte{}, "podman run \\\n	--rm alpine:latest \\\n	sh -c 'echo Hello, World!'\n")
	baseComposeFile = fmt.Appendf([]byte{}, "services:\n  app:\n    image: alpine:latest\n    command: [\"sh\", \"-c\", \"echo Hello, World!\"]\n")
)

type PodType string

const (
	CommandPod PodType = "command"
	ComposePod PodType = "compose"
)

type Pod struct {
	Type      PodType   `toml:"type"`
	Name      string    `toml:"name"`
	Timestamp time.Time `toml:"timestamp"`
}

type PodConfig struct {
	Pods map[string]*Pod `toml:"pods"`
}

func ReadPodConfig() (*PodConfig, error) {
	file, err := os.ReadFile("pods.toml")
	if err != nil {
		return nil, fmt.Errorf("failed to read pods.toml: %w", err)
	}

	data := new(PodConfig)
	data.Pods = make(map[string]*Pod)
	if err := toml.Unmarshal(file, data); err != nil {
		return nil, fmt.Errorf("failed to parse pods.toml: %w", err)
	}

	return data, nil
}

func RegisterPod(pc *PodConfig, p *Pod) error {
	if p == nil {
		return fmt.Errorf("pod is nil")
	}
	if pc == nil {
		return fmt.Errorf("pod config is nil")
	}

	if _, exists := pc.Pods[p.Name]; exists {
		return fmt.Errorf("pod with name %s already exists", p.Name)
	}

	p.Timestamp = time.Now()
	pc.Pods[p.Name] = p

	podsData, err := toml.Marshal(pc)
	if err != nil {
		return fmt.Errorf("failed to marshal pod config: %w", err)
	}

	output := "pods.toml"
	err = os.WriteFile(output, podsData, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write pod config to %s: %w", output, err)
	}

	return nil
}

func CreatePodDirectory(podName string) error {
	if strings.TrimSpace(podName) == "" {
		return fmt.Errorf("pod name cannot be empty")
	}
	if err := os.MkdirAll("pods/"+podName, 0o755); err != nil {
		return fmt.Errorf("failed to create pod directory: %w", err)
	}
	return nil
}

func CreatePod(name string, typ PodType) error {
	conf, err := ReadPodConfig()
	if err != nil {
		return err
	}

	if _, exists := conf.Pods[name]; exists {
		return fmt.Errorf("pod with name '%s' already exists", name)
	}

	if err = CreatePodDirectory(name); err != nil {
		return err
	}

	switch typ {
	case CommandPod:
		err = os.WriteFile(fmt.Sprintf("pods/%s/command", name), baseCommandFile, 0o644)
		if err != nil {
			return fmt.Errorf("failed to create command pod script: %w", err)
		}
	case ComposePod:
		err = os.WriteFile(fmt.Sprintf("pods/%s/compose.yml", name), baseComposeFile, 0o644)
		if err != nil {
			return fmt.Errorf("failed to create compose pod file: %w", err)
		}
	}

	newPod := &Pod{Name: name, Type: typ}
	if err = RegisterPod(conf, newPod); err != nil {
		return err
	}

	return nil
}
