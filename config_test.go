package config_test

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/skysparq/config-go"
)

type TestConfig struct {
	Name   string
	Option int
}

func TestExportConfig(t *testing.T) {
	cfg := TestConfig{
		Name:   "Hello",
		Option: 2,
	}
	result := config.Export(cfg)
	expected := base64.URLEncoding.EncodeToString([]byte(`{"Name":"Hello","Option":2}`))
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
	t.Log(expected)
}

func TestLoadConfig(t *testing.T) {
	err := os.Setenv(`CONFIG`, `eyJOYW1lIjoiSGVsbG8iLCJPcHRpb24iOjJ9`)
	if err != nil {
		t.Fatal(err)
	}
	cfg, err := config.Load[TestConfig]()
	if err != nil {
		t.Fatal(err)
	}
	expected := TestConfig{
		Name:   "Hello",
		Option: 2,
	}
	if cfg != expected {
		t.Errorf("got %+v, want %+v", cfg, expected)
	}
}

func TestLoadConfigEnvNotSet(t *testing.T) {
	err := os.Setenv("CONFIG", "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = config.Load[TestConfig]()
	if err == nil {
		t.Fatal(`expected an error`)
	}
	t.Log(err)
}

func TestLoadConfigInvalidBase64(t *testing.T) {
	err := os.Setenv("CONFIG", `{"Name":"Hello","Option":2}`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = config.Load[TestConfig]()
	if err == nil {
		t.Fatal(`expected an error`)
	}
	t.Log(err)
}

func TestLoadConfigInvalidJson(t *testing.T) {
	err := os.Setenv("CONFIG", `aGVsbG8gd29ybGQ=`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = config.Load[TestConfig]()
	if err == nil {
		t.Fatal(`expected an error`)
	}
	t.Log(err)
}

func TestLoadConfigFromFile(t *testing.T) {
	cfg, err := config.LoadFromPath[TestConfig](`./test.json`)
	if err != nil {
		t.Fatal(err)
	}
	expected := TestConfig{
		Name:   "Hello",
		Option: 2,
	}
	if cfg != expected {
		t.Errorf("got %+v, want %+v", cfg, expected)
	}
}
