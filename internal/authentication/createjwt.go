package authentication

import (
	"crypto/x509"
	"encoding/pem"
	"time"

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

func CreateJWT(ttl time.Duration, gittoken string) (string, error) {
	priPEM, err := getSSMParameterValue("prikey")
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}


	block, _ := pem.Decode([]byte(priPEM))

	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	claims := make(jwt.MapClaims)
	claims["token"] = gittoken

	jwt, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(pri)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return jwt, nil
}
