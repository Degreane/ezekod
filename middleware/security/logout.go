package security

import (
	"strings"

	"github.com/degreane/ezekod.com/middleware/ezelogger"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	// To Actually logout we need to be holding an Authorization header
	lhs := "Middleware->Security->Logout : <Success>"
	lhe := "Middleware->Security->Logout : <Error>"
	// Get Headers
	headers := c.GetReqHeaders()
	if authHeader, ok := headers["Authorization"]; !ok {
		// no Authorization header supplied then we short circuit and return false to Locals['LoggedIn']
		c.Locals("LoggedIn", false)
	} else {
		if len(strings.Fields(authHeader)) != 2 {
			// Check for length of the fields found in the header it should be 2 [Bearer, jwtToken]
			ezelogger.Ezelogger.Printf("%s MalFormed Header '%s'", lhe, authHeader)
			c.Locals("LoggedIn", false)
		} else {
			authHeader = strings.Fields(authHeader)[1]
			ezelogger.Ezelogger.Printf("%s %s", lhs, authHeader)
			decryptedAuthHeader := decrypt(authHeader)
			if decryptedAuthHeader != "" {
				claims, _ := parseJWT(decryptedAuthHeader)
				ezelogger.Ezelogger.Printf("%+v", claims)
			}
		}

	}

	return c.Next()
}
