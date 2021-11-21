package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	ConfigFileName         = ".pongbox"
	DefaultTimeoutDuration = 3 * time.Second
)

type Config struct {
	General  GeneralConfig            `json:"general"`
	Commands map[string]CommandConfig `json:"commands"`
}

type GeneralConfig struct {
	Timeout Duration `json:"timeout"`
}

type CommandConfig struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

func loadConfig() (Config, error) {
	var config Config

	home, err := os.UserHomeDir()

	if err != nil {
		return config, err
	}

	data, err := ioutil.ReadFile(filepath.Join(home, ConfigFileName))

	if err != nil {
		return config, err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}

func withTimeout(ctx context.Context, config Config) (context.Context, context.CancelFunc) {
	if config.General.Timeout.Duration != 0 {
		return context.WithTimeout(ctx, config.General.Timeout.Duration)
	}

	return context.WithTimeout(ctx, DefaultTimeoutDuration)
}
