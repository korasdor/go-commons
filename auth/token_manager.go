package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenManager interface {
	NewJWT(userId string, ttl time.Duration) (string, error)
	ParseJWT(token string) (string, error)
}

type Manager struct {
	signingKey []byte
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{
		signingKey: []byte(signingKey),
	}, nil
}

func (m *Manager) NewJWT(userId string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(ttl).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(m.signingKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (m *Manager) ParseJWT(tokenString string) (string, error) {
	var userId string

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return m.signingKey, nil
	})
	if err != nil {
		return userId, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return userId, errors.New("error get user claims from token")
	}

	userId, ok = claims["sub"].(string)
	if !ok {
		return userId, errors.New("error while geting userId from jwt claims")
	}

	return userId, nil
}
