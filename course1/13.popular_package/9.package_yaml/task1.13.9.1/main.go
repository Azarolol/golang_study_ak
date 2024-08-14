package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func getYAML(input []Config) (string, error) {
	output, err := yaml.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func main() {
	input := []Config{
		{
			Server{"8080"},
			Db{"localhost", "5432", "admin", "password123"},
		},
	}
	output, err := getYAML(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
