package glav

import "golang.org/x/crypto/sha3"

const (
	saltPrefix = "LIBRA::StructTag"
)

var (
	codeCoreAddress = [20]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	salt            = sha3.Sum256([]byte(saltPrefix))
)
