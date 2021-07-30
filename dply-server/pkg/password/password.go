package password

type Password interface {
	Hash(password string) (string, error)
	Check(password, hashedPassword string) bool
}
