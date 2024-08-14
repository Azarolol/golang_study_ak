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

func getConfigFromYAML(data []byte) (Config, error) {
	var output Config
	err := yaml.Unmarshal(data, &output)
	if err != nil {
		return Config{}, err
	}
	return output, nil
}

func main() {
	input := `server: 
 port: "8080"
db:
 host: localhost
 port: "5432"
 user: admin
 password: password123`
	output, err := getConfigFromYAML([]byte(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
