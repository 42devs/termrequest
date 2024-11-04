package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

func getInitialConfig(fileconfig string) Configuration {
	hash := generateHash()

	configuration := Configuration{
		Hash: hash,
	}

	return configuration

}

func getConfigFile() {}
