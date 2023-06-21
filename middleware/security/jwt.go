// Create use and consume JWT tokens
package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"os"
	"path"

	"github.com/degreane/ezekod.com/middleware/ezelogger"
	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Security ConfigElements `yaml:"security"`
}
type ConfigElements struct {
	Type   string `yaml:"type"`
	Secret string `yaml:"secret"`
	EncKey string `yaml:"encKey"`
}

type JWTClaims struct {
	ID    string
	Group string
	jwt.RegisteredClaims
}

var (
	cfg Config
)

func readConfig() {
	// read COnfiguration from file in current Directory
	// it is a yaml file so we need to parse it
	lhe := "Middleware->Security->jwt->readConfig : <Error> "
	lhs := "Middleware->Security->jwt->readConfig : <Success> "

	rootdir, err := os.Getwd()
	if err != nil {
		ezelogger.Ezelogger.Fatalf("%s % +v", lhe, err)
	}
	cfgFile := path.Join(rootdir, "middleware", "security", "config.yaml")
	if _, ok := os.Stat(cfgFile); ok != nil {
		ezelogger.Ezelogger.Fatalf("%s % +v", lhe, ok)
	}
	cfgContent, err := os.ReadFile(cfgFile)
	if err != nil {
		ezelogger.Ezelogger.Fatalf("%s % +v", lhe, err)
	}
	err = yaml.Unmarshal(cfgContent, &cfg)
	if err != nil {
		ezelogger.Ezelogger.Fatalf("%s % +v", lhe, err)
	}
	ezelogger.Ezelogger.Printf("%+s %+v", lhs, cfg.Security.Secret)
}

func encrypt(plaintext string) string {

	lhe := "MiddleWare->Security->JWT->encrypt : <Error> "
	// lhs := "MiddleWare->Security->JWT->encrypt : <Success> "
	var secretKey [16]byte = md5.Sum([]byte(cfg.Security.EncKey))

	aes, err := aes.NewCipher([]byte(secretKey[:]))
	if err != nil {
		ezelogger.Ezelogger.Printf("%s (newCipher) %+v", lhe, err)
	} else {
		gcm, err := cipher.NewGCM(aes)
		if err != nil {

			ezelogger.Ezelogger.Printf("%s (newGCM) %+v", lhe, err)
		} else {
			// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
			// A nonce should always be randomly generated for every encryption.
			nonce := make([]byte, gcm.NonceSize())
			_, err = rand.Read(nonce)
			if err != nil {
				ezelogger.Ezelogger.Printf("%s (readNonce) %+v", lhe, err)
			} else {
				// ciphertext here is actually nonce+ciphertext
				// So that when we decrypt, just knowing the nonce size
				// is enough to separate it from the ciphertext.
				ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
				return base64.StdEncoding.EncodeToString(ciphertext)
				// return string(ciphertext)
			}
		}
	}
	return ""
}
func decrypt(b64ciphertext string) string {
	lhe := "MiddleWare->Security->JWT->decrypt : <Error> "
	ciphertext, err := base64.StdEncoding.DecodeString(b64ciphertext)
	if err != nil {
		ezelogger.Ezelogger.Printf("%s (b64) %+v", lhe, err)
	} else {
		var secretKey [16]byte = md5.Sum([]byte(cfg.Security.EncKey))
		aes, err := aes.NewCipher([]byte(secretKey[:]))
		if err != nil {
			ezelogger.Ezelogger.Printf("%s (newCipher) %+v", lhe, err)
		} else {
			gcm, err := cipher.NewGCM(aes)
			if err != nil {
				ezelogger.Ezelogger.Printf("%s (newGCM) %+v", lhe, err)
			} else {
				// Since we know the ciphertext is actually nonce+ciphertext
				// And len(nonce) == NonceSize(). We can separate the two.
				nonceSize := gcm.NonceSize()
				nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
				plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
				if err != nil {
					ezelogger.Ezelogger.Printf("%s (gcm.Open) %+v", lhe, err)
				} else {
					return string(plaintext)
				}
			}
		}
	}
	return ""
}

func newJWT(cl JWTClaims) string {

	lhe := "Middleware->Security->JWT->New : <Err> "
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = cl.ID
	claims["Group"] = cl.Group
	t, err := token.SignedString([]byte(cfg.Security.Secret))
	if err != nil {
		ezelogger.Ezelogger.Printf("%s %+v", lhe, err)
	}
	return encrypt(t)
	// ezelogger.Ezelogger.Printf("%s", t)
}

func parseJWT(token string) (JWTClaims, error) {
	claims := &JWTClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.Security.Secret), nil
	})
	if err != nil {
		ezelogger.Ezelogger.Fatalf("No Claime %s", err)
	}
	if !tkn.Valid {
		ezelogger.Ezelogger.Fatalf("Token Not Valid %s", tkn.Header)
	}
	return *claims, nil
}
