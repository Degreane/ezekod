// In package server/routes we should:
// Identify reading routes.yaml file
package server

import (
	"fmt"
	"os"
	"path"

	"github.com/degreane/ezekod.com/middleware"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type APath struct {
	Path        string   `yaml:"path"`
	Method      string   `yaml:"method"`
	MiddleWares []string `yaml:"middlewares"`
}
type Routes struct {
	Paths []APath `yaml:"routes"`
}

func (server *Server) applyRoutes(pth []APath) {
	for pIdx := 0; pIdx < len(pth); pIdx++ {
		thePath := pth[pIdx]
		var elements []func(c *fiber.Ctx) error
		for _, element := range thePath.MiddleWares {
			elements = append(elements, middleware.MiddleWares[element])
		}
		server.App.Add(thePath.Method, thePath.Path, elements...)

	}
}

func (server *Server) readRoutesConfiguration(fileName string) {
	defaultPathConstruct := Routes{
		Paths: []APath{
			{
				Path:        "/",
				Method:      "GET",
				MiddleWares: []string{"defaultGet"},
			},
			{
				Path:        "/",
				Method:      "POST",
				MiddleWares: []string{"defaultPost"},
			},
		},
	}
	var theRoutes Routes
	if fileName == "" {
		theRoutes = defaultPathConstruct
	} else {
		fileContent, err := os.ReadFile(fileName)
		if err != nil {
			theRoutes = defaultPathConstruct
		} else {
			err = yaml.Unmarshal(fileContent, &theRoutes)
			if err != nil {
				fmt.Printf("Err: unmarshalling %+v \n", err)
			} else {
				// e := reflect.ValueOf(&theRoutes).Elem()
				// fmt.Printf("%s \n", e.Type().Field(0).Name)
				// loop over the Paths and get the path to work on
				server.applyRoutes(theRoutes.Paths)
			}
		}
	}

}

// SetRoutes() Called
//
// then (server *Server)readRoutesConfiguration(routes.yaml)
//
// then (server *Server)applyRoutes([]APath).
func (server *Server) SetRoutes() {
	configFile := path.Join(".", "server", "routes.yaml")
	_, err := os.Stat(configFile)
	if err != nil {
		server.readRoutesConfiguration("")
	} else {
		server.readRoutesConfiguration(configFile)
	}
}
