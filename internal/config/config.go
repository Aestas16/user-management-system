package config

import (
    "os"
    "gopkg.in/yaml.v3"
)

type Configuration struct {
    Server  ServerConfig    `yaml:"server"`
    SQL     SQLConfig       `yaml:"postgresql"`
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
type SQLConfig struct {
    User    string  `yaml:"user"`
    Password    string  `yaml:"password"`
    Port    string  `yaml:"port"`
    DBName    string  `yaml:"dbname"`
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