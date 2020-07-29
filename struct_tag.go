package glav

import (
	"bytes"

	"golang.org/x/crypto/sha3"
)

const (
	ResourceTag = 0x1
)

type StructTag struct {
	address    [20]byte
	moduleName string
	structName string
	typeParams []TypeParam
}

func (tag StructTag) AccessVector() []byte {
	h := sha3.New256()
	access := make([]byte, 33)
	access[0] = ResourceTag
	h.Write(salt[:])
	var buffer bytes.Buffer
	tag.marshal(&buffer)
	h.Write(buffer.Bytes())

	h.Sum(access[:1])

	return access
}

func (tag StructTag) marshal(buffer *bytes.Buffer) {
	buffer.Write(tag.address[0:])
	writeStringToBuffer(buffer, tag.moduleName)
	writeStringToBuffer(buffer, tag.structName)
	writeUintToBuffer(buffer, uint64(len(tag.typeParams)))

	for _, param := range tag.typeParams {
		param.marshal(buffer)
	}
}

func NewStructTag(address [20]byte, moduleName string, structName string, typeParams []TypeParam) StructTag {
	return StructTag{
		address,
		moduleName,
		structName,
		typeParams,
	}
}
