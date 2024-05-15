package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
	"time"
)

var (
	ErrInvalidClaims = errors.New("invalid claims")
)

func GenerateTokenByHS256(secret string, payload map[string]interface{}, expiry int64) (string, error) {
	claims := jwt.MapClaims{
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expiry))),
	}
	for k, v := range payload {
		claims[k] = v
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return t.SignedString([]byte(secret))
}

func parseToken[T any](token *jwt.Token) (*T, error) {
	if !token.Valid {
		return nil, ErrInvalidClaims
	}

	var payload T
	err := mapstructure.Decode(token.Claims, &payload)
	if err != nil {
		return nil, errors.Join(err, ErrInvalidClaims)
	}

	return &payload, nil
}

func VerifyByHS256[T any](secret string, myToken string) (*T, error) {
	res, err := jwt.Parse(myToken, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return parseToken[T](res)
}

type RsaGetter func(source string) (*rsa.PublicKey, bool)

func VerifyByRsaKeyFromToken[T any](myToken string, sourceKey string, keyGetter RsaGetter) (*T, error) {
	res, err := jwt.Parse(myToken, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			return nil, ErrInvalidClaims
		}

		source, ok := claims[sourceKey].(string)
		if !ok {
			return nil, fmt.Errorf("invalid source %s", claims[sourceKey])
		}

		publicKey, ok := keyGetter(source)
		if !ok {
			return nil, fmt.Errorf("source %s not support", source)
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	return parseToken[T](res)
}
