package password

import (
	"golang.org/x/crypto/bcrypt"
)

type bcryptPassword struct {
	salt      string
	delimiter string
}

func NewBcryptPassword(salt string) Password {
	return &bcryptPassword{salt, "#|#"}
}

func (s *bcryptPassword) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s.concatPasswordSalt(password)), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *bcryptPassword) Check(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(s.concatPasswordSalt(password)))
	return err == nil
}

func (s *bcryptPassword) concatPasswordSalt(password string) string {
	return password + s.delimiter + s.salt
}
