package gonnarain

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Interval time.Duration `yaml:"interval"`
}

var DefaultConfig = Config{
	Interval: time.Second,
}

func (c *Config) Load(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	configFlag := flags.String("config", "", "Path to config file")

	if err := flags.Parse(args[1:]); err != nil {
		return fmt.Errorf("args parse failed: %w", err)
	}

	if *configFlag != "" {
		cfg, err := fromYaml(*configFlag)

		if err != nil {
			return fmt.Errorf("config load failed: %w", err)
		}

		*c = *cfg
	} else {
		*c = DefaultConfig
	}

	return nil
}

func fromYaml(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)

	if err != nil {
		return nil, fmt.Errorf("file read failed: %w", err)
	}

	cfg := Config{}

	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, fmt.Errorf("file parse failed: %w", err)
	}

	return &cfg, nil
}
