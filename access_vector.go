package libra_access_vector

import (
	"bytes"
	"golang.org/x/crypto/sha3"
)

const PREFIX = "LIBRA::StructTag"
const ResourceTag = 0x1

type StructTag struct {
	address    [20]byte
	moduleName string
	structName string
	typeParams []TypeParam
}

func Tag(address [20]byte, moduleName string, structName string, typeParams []TypeParam) StructTag {
	return StructTag{
		address,
		moduleName,
		structName,
		typeParams,
	}
}

func (tag StructTag) AccessVector() []byte {
	h := sha3.New256()
	access := make([]byte, 33)
	access[0] = ResourceTag
	salt := salt()
	h.Write(salt[0:])
	var buffer bytes.Buffer
	tag.marshal(&buffer)
	h.Write(buffer.Bytes())

	h.Sum(access[:1])
	return access
}

func (tag StructTag) marshal(buffer *bytes.Buffer) {
	buffer.Write(tag.address[0:])
	writeStr(buffer, tag.moduleName)
	writeStr(buffer, tag.structName)
	writeVarUint(buffer, uint64(len(tag.typeParams)))

	for _, param := range tag.typeParams {
		param.marshal(buffer)
	}
}

func writeStr(buffer *bytes.Buffer, val string) {
	writeVarUint(buffer, uint64(len(val)))
	buffer.Write([]byte(val))
}

func salt() [32]byte {
	return sha3.Sum256([]byte(PREFIX))
}

type TypeParam struct {
	pType     byte
	typeParam *TypeParam
	tag       *StructTag
}

func (p TypeParam) marshal(buffer *bytes.Buffer) {
	buffer.WriteByte(p.pType)
	if p.typeParam != nil {
		p.typeParam.marshal(buffer)
	}

	if p.tag != nil {
		p.tag.marshal(buffer)
	}
}

const (
	BoolType    = 0x0
	U8Type      = 0x1
	U64Type     = 0x2
	U128Type    = 0x3
	AddressType = 0x4
	SignerType  = 0x5
	VectorType  = 0x6
	StructType  = 0x7
)

func Bool() TypeParam {
	return TypeParam{BoolType, nil, nil}
}

func U8() TypeParam {
	return TypeParam{U8Type, nil, nil}
}

func U64() TypeParam {
	return TypeParam{U64Type, nil, nil}
}

func U128() TypeParam {
	return TypeParam{U128Type, nil, nil}
}

func Address() TypeParam {
	return TypeParam{AddressType, nil, nil}
}

func Signer() TypeParam {
	return TypeParam{SignerType, nil, nil}
}

func Vector(param TypeParam) TypeParam {
	return TypeParam{VectorType, &param, nil}
}

func Struct(tag StructTag) TypeParam {
	return TypeParam{StructType, nil, &tag}
}

func CoreCodeAddress() [20]byte {
	return [20]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
}

func writeVarUint(w *bytes.Buffer, v uint64) {
	var buf []byte
	for {
		c := uint8(v & 0x7f)
		v >>= 7
		if v != 0 {
			c |= 0x80
		}
		buf = append(buf, c)
		if c&0x80 == 0 {
			break
		}
	}
	w.Write(buf)
}
