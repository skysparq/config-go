package config

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func Load[T any]() (T, error) {
	var config T
	configBytes := []byte(os.Getenv("CONFIG"))
	if len(configBytes) == 0 {
		return config, errors.New("error loading config: CONFIG environment variable not set")
	}

	decoded, err := base64.URLEncoding.DecodeString(string(configBytes))
	if err != nil {
		return config, fmt.Errorf("error loading config: config is not Base64 encoded: %v", err)
	}

	err = json.Unmarshal(decoded, &config)
	if err != nil {
		return config, fmt.Errorf("error loading config: config is not valid JSON: %v", err)
	}

	return config, nil
}

func LoadFromPath[T any](configPath string) (T, error) {
	var config T
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("error loading config: could not read config file: %v", err)
	}

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return config, fmt.Errorf("error loading config: config is not valid JSON: %v", err)
	}

	return config, nil
}

func Export[T any](config T) string {
	configBytes, _ := json.Marshal(config)
	return base64.URLEncoding.EncodeToString(configBytes)
}
