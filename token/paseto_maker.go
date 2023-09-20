package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	claims, err := NewPayload(username, duration)
	if err != nil {
		return "", claims, err
	}
	token, err := maker.paseto.Encrypt(maker.symmetricKey, claims, nil)
	return token, claims, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	claims := &Payload{}
	err := maker.paseto.Decrypt(token, maker.symmetricKey, claims, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = claims.Valid()
	if err != nil {
		return nil, err
	}
	return claims, nil
}
