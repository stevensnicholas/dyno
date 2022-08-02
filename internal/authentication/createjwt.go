package authentication

import (
	"time"
	"encoding/pem"
	"crypto/x509"
	
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
	priPEM := "-----BEGIN PRIVATE KEY-----\n"+
				getSSMParameterValue("testkey")+
				"\n-----END PRIVATE KEY-----"

	block, _ := pem.Decode([]byte(priPEM))
    pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	claims := make(jwt.MapClaims)
	claims["token"] = gittoken

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(pri)

	if err != nil {
		logger.Error(err.Error())
			return "", err
	}

	return token, nil
}
