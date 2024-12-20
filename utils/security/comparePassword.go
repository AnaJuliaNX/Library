package security

import "golang.org/x/crypto/bcrypt"

type Password struct{}

func NewPass() *Password {
	return &Password{}
}

func (p *Password) Generate(raw string) (hash string, err error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashByte), nil
}

func (p *Password) ComaprePassword(p1, p2 string) error {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2))
}
