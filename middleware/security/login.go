package security

import (
	"github.com/degreane/ezekod.com/middleware/ezelogger"
	"github.com/degreane/ezekod.com/model/users"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// POSTED Data should be in the body
func Login(c *fiber.Ctx) error {
	// Key Concept here is to take the body content of c.Body()
	// unmarshal it from json to bson
	// connect with users model and find if document match
	lhe := "MiddleWare->Security->Login : <Error> "
	lhs := "MiddleWare->Security->Login : <Success> "
	var userBody interface{}
	readConfig()
	err := bson.UnmarshalExtJSON(c.Body(), true, &userBody)
	if err != nil {
		ezelogger.Ezelogger.Printf("%s % +v", lhe, err)
	} else {
		ezelogger.Ezelogger.Printf("%s % +v", lhs, err)
	}

	user, ok := users.Find(userBody)
	if ok {
		ezelogger.Ezelogger.Printf("%s % +v", lhs, user)
		newJWTToken := newJWT(JWTClaims{
			ID:    user.ID.String(),
			Group: user.Group,
		})
		c.Locals("auth", newJWTToken)
		c.Set("X-AuthToken", newJWTToken)
	} else {
		ezelogger.Ezelogger.Printf("%s : No User Found ", lhe)
	}

	return c.Next()
}
