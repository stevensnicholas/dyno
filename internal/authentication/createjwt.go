package authentication

import (
	"time"
	"io/ioutil"
	"dyno/internal/logger"
	"github.com/golang-jwt/jwt"
)

type JWT struct {
	privateKey []byte
	publicKey  []byte
}
 
func NewJWT(privateKey []byte, publicKey []byte) JWT {
	return JWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}


func CreateToken(ttl time.Duration, gittoken string) (string, error) {

	prvKey, err := ioutil.ReadFile("cert/id_rsa")
	if err != nil {
		logger.Error(err.Error())
			return "", err
	}
	pubKey, err := ioutil.ReadFile("cert/id_rsa")
	if err != nil {
		logger.Error(err.Error())
			return "", err
	}

	jwtToken := NewJWT(prvKey, pubKey)

	key, err := jwt.ParseRSAPrivateKeyFromPEM(jwtToken.privateKey)

	if err != nil {
		logger.Error(err.Error())
			return "", err
	}

	claims := make(jwt.MapClaims)
	claims["token"] = gittoken

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		logger.Error(err.Error())
			return "", err
	}

	return token, nil
}
