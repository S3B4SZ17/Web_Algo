package app

import (
	"os"

	"gopkg.in/yaml.v2"
)

// NewConfig returns a new decoded Config struct
func ReadYaml(configPath string) (*Config, error) {
    // Create config structure
    config := &Config{}

    // Open config file
    file, err := os.Open(configPath)
    if err != nil {
      return nil, err
    }
    defer file.Close()

    // Init new YAML decode
    d := yaml.NewDecoder(file)

    // Start YAML decoding from file
    if err := d.Decode(&config); err != nil {
        return nil, err
    }

    return config, nil
}

type Config struct {
  Gin_mode string `yaml:"gin_mode"`
  HttpPort string `yaml:"http_port"`
  Cors struct {
    Enabled bool `yaml:"enabled"`
    List_hosts []string `yaml:"list_hosts"`
  } `yaml:"cors"`
}