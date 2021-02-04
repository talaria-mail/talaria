package encryption

import "golang.org/x/crypto/argon2"

func HashPassword(passwd, salt []byte) []byte {
	var (
		time    uint32 = 3
		memory  uint32 = 512 * 1024
		threads uint8  = 4
		len     uint32 = 32
	)

	return argon2.IDKey(passwd, salt, time, memory, threads, len)
}
