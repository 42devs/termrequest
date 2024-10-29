package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Method string

const (
	Get     Method = "GET"
	Post    Method = "POST"
	Put     Method = "PUT"
	Patch   Method = "PATCH"
	Delete  Method = "DELETE"
	Options Method = "OPTIONS"
	HEAD    Method = "HEAD"
)

type Header struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type URLSearchParam struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type Request struct {
	Url             string           `yaml:"url"`
	Method          string           `yaml:"method"`
	Headers         []Header         `yaml:"headers,omitempty"`
	Path            string           `yaml:"path,omitempty"`
	URLSearchParams []URLSearchParam `yaml:"urlSearchParams,omitempty"`
	Body            string           `yaml:"body,omitempty"`
}

type EnvironmentVariable struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type Folders struct {
	Name                 string                `yaml:"name"`
	Description          string                `yaml:"description,omitempty"`
	Requests             []Request             `yaml:"requests,omitempty"`
	EnvironmentVariables []EnvironmentVariable `yaml:"env,omitempty"`
}

type Workspace struct {
	Name                 string                `yaml:"name"`
	EnvironmentVariables []EnvironmentVariable `yaml:"env"`
}

type Configuration struct {
	Hash       string      `yaml:"hash"`
	Workspaces []Workspace `yaml:"workspaces"`
}

var configuration Configuration

func getWorkspaces() Workspace {
	if configuration = nil {
	 // TODO: validar aca

	}
}

func createConfigFile() Configuration {

}

func readConfigFile() Configuration {

}

func getConfiguration() Configuration {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	configPath := filepath.Join(homeDir, ".config", "termrequest")

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(configPath)
		if err != nil {
			panic(err)
		}
	}

	configFile := filepath.Join(configPath, "settings.yml")

	if _, err := is, Stat(configFile); os.IsNotExist(err) {
		newFile, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}

		defer newFIle.close()

	}
}
