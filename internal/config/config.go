package config

import (
    "os"
    "gopkg.in/yaml.v3"
)

type Configuration struct {
    Server  ServerConfig    `yaml:"server"`
}
type ServerConfig struct {
    Port    string      `yaml:"port"`
    Version string      `yaml:"version"`
    Admin   AdminConfig `yaml:"admin"`
}
type AdminConfig struct {
    Username    string  `yaml:"username"`
    Password    string  `yaml:"password"`
}

var Config Configuration

func Init() {
    yamlFile, err := os.ReadFile("config.yaml")
    if err != nil {
        panic(err)
    }
    err = yaml.Unmarshal(yamlFile, &Config)
    if err != nil {
        panic(err)
    }
}