package glav

import "golang.org/x/crypto/sha3"

const (
	structTag   = "LIBRA::StructTag"
	moduleIdTag = "LIBRA::ModuleId"
)

var (
	codeCoreAddress = [20]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	structTagSalt   = sha3.Sum256([]byte(structTag))
	moduleIdSalt    = sha3.Sum256([]byte(moduleIdTag))
)
