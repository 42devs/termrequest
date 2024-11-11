package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

type Method string

const (
	Get     Method = "GET"
	Post    Method = "POST"
	Put     Method = "PUT"
	Patch   Method = "PATCH"
	Delete  Method = "DELETE"
	Options Method = "OPTIONS"
	Head    Method = "HEAD"
)

type Header struct {
	Key    string `yaml:"key"`
	Value  string `yaml:"value"`
	Enable bool   `default:"true" yaml:"enable"`
}

type URLSearchParam struct {
	Key    string `yaml:"key"`
	Value  string `yaml:"value"`
	Enable bool   `default:"true" yaml:"enable"`
}

type Request struct {
	Url             string           `yaml:"url"`
	Method          string           `yaml:"method"`
	Headers         []Header         `yaml:"headers,omitempty"`
	Path            string           `yaml:"path,omitempty"`
	URLSearchParams []URLSearchParam `yaml:"urlSearchParams,omitempty"`
	Body            string           `yaml:"body,omitempty"`
	Enable          bool             `default:"true" yaml:"enable"`
}

type EnvironmentVariable struct {
	Key    string `yaml:"key"`
	Value  string `yaml:"value"`
	Enable bool   `default:"true" yaml:"enable"`
}

type Folders struct {
	Name                 string                `yaml:"name"`
	Description          string                `yaml:"description,omitempty"`
	Requests             []Request             `yaml:"requests,omitempty"`
	EnvironmentVariables []EnvironmentVariable `yaml:"env,omitempty"`
	Enable               bool                  `default:"true" yaml:"enable"`
}

type Workspace struct {
	Name                 string                `yaml:"name"`
	EnvironmentVariables []EnvironmentVariable `yaml:"env"`
	Enable               bool                  `default:"true" yaml:"enable"`
}

type Configuration struct {
	Hash                        string                `yaml:"hash"`
	InitialDate                 time.Time             `yaml:"initialDate"`
	Version                     string                `yaml:"version"`
	GlobalEnvironmentsVariables []EnvironmentVariable `yaml:"env,omitempty"`
	Workspaces                  []Workspace           `yaml:"workspaces"`
}

var configuration Configuration

func generateHash() string {
	now := time.Now()
	ts := now.Unix()
	tsString := strconv.FormatInt(ts, 10)

	hash := sha256.New()

	hash.Write([]byte(tsString))

	sum := hash.Sum(nil)

	hashString := fmt.Sprintf("%x", sum)
	return hashString

}

func getConfigPath() string {
	homedir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	configpath := filepath.Join(homedir, ".config", "termrequest")

	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		err := os.MkdirAll(configpath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	return configpath
}

func getInitialConfig() Configuration {
	hash := generateHash()

	configuration := Configuration{
		Hash:        hash,
		InitialDate: time.Now(),
	}

	return configuration

}

func getConfigFile() Configuration {
	configFile := filepath.Join(getConfigPath(), "settings.yml")
	var configuration Configuration
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		configuration = getInitialConfig()
		yamlFile, err := yaml.Marshal(&configuration)
		if err != nil {
			panic(err)
		}

		_, err = io.WriteString(f, string(yamlFile))
		if err != nil {
			panic(err)
		}
	} else {
		yamlFile, err := os.ReadFile(configFile)
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(yamlFile, &configuration)
		if err != nil {
			panic(err)
		}
	}

	return configuration
}
