package glav

import (
	"bytes"
	"golang.org/x/crypto/sha3"
)

func ModuleAccessVector(address [20]byte, name string) []byte {
	h := sha3.New256()
	access := make([]byte, 33)
	access[0] = ModuleTag
	h.Write(moduleIdSalt[:])

	var buffer bytes.Buffer
	buffer.Write(address[:])
	writeStringToBuffer(&buffer, name)

	h.Write(buffer.Bytes())
	h.Sum(access[:1])

	return access
}
