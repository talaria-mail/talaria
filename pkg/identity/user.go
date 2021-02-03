package identity

type User struct {
	Login        string
	Name         string
	EmailAliases []string

	PasswordHash []byte // Argon2 hash of password
	Salt         []byte // 32 byte salf for key derivation of kek encryption key

	// Content key for encryption of user data
	ContentKey struct {
		Private []byte // Encrypted Curve25519 private key
		Public  []byte // Public key
	}
}
