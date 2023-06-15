package server

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type ServerAddress struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}
type ServerConfig struct {
	Address      ServerAddress
	Prefork      bool   `yaml:"prefork"`
	ServerHeader string `yaml:"serverheader"`
	AppName      string `yaml:"appname"`
}

type Server struct {
	App  *fiber.App
	Port string
	Host string
}

var App *Server = &Server{}

// func packageName(v interface{}) string {
// 	if v == nil {
// 		return ""
// 	}

//		val := reflect.ValueOf(v)
//		if val.Kind() == reflect.Ptr {
//			return val.Elem().Type().PkgPath()
//		}
//		return val.Type().PkgPath()
//	}
func readConfig() ServerConfig {
	configFile := path.Join(".", "server", "config.yaml")
	_, err := os.Stat(configFile)
	defaultConfig := ServerConfig{
		Address: ServerAddress{
			Port: "3000",
			Host: "0.0.0.0",
		},
		Prefork:      true,
		ServerHeader: "EzeKod",
		AppName:      "ExeCod",
	}
	if err != nil {
		return defaultConfig
	} else {
		content, err := os.ReadFile(configFile)
		if err != nil {
			return defaultConfig
		} else {

			config := make(map[string]ServerConfig)
			err = yaml.Unmarshal(content, &config)
			if err != nil {
				return defaultConfig
			} else {
				if _, ok := config["server"]; ok {
					return config["server"]
				} else {
					return defaultConfig
				}

			}
		}
	}
}

func Init() *Server {
	sAddr := readConfig()
	app := fiber.New(
		fiber.Config{
			Prefork:      sAddr.Prefork,
			ServerHeader: sAddr.ServerHeader,
			AppName:      sAddr.AppName,
		},
	)
	port := sAddr.Address.Port
	host := sAddr.Address.Host
	server := Server{
		App:  app,
		Port: port,
		Host: host,
	}
	return &server
}

func (server *Server) StartServer() {

	log.Fatal(server.App.Listen(fmt.Sprintf("%s:%s", server.Host, server.Port)))
}
