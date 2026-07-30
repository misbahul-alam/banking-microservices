package password

import "golang.org/x/crypto/bcrypt"

type Service interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}

type service struct {
	cost int
}

func New(cost int) Service {
	return &service{
		cost: cost,
	}
}

func (s *service) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		s.cost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s *service) Compare(
	hash,
	password string,
) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
