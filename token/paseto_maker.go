package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetorMake struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func (maker *PasetorMake) CreateToken(username string, id int32, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, id, duration)
	if err != nil {
		return "", nil, err
	}
	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

func (maker *PasetorMake) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly #{chacha20poly1305.KeySize} characters")
	}

	maker := &PasetorMake{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}
