package glav

import (
	"bytes"
)

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

func NewBoolTypeParam() TypeParam {
	return TypeParam{BoolType, nil, nil}
}

func NewU8TypeParam() TypeParam {
	return TypeParam{U8Type, nil, nil}
}

func NewU64TypeParam() TypeParam {
	return TypeParam{U64Type, nil, nil}
}

func NewU128TypeParam() TypeParam {
	return TypeParam{U128Type, nil, nil}
}

func NewAddressTypeParam() TypeParam {
	return TypeParam{AddressType, nil, nil}
}

func NewSignerTypeParam() TypeParam {
	return TypeParam{SignerType, nil, nil}
}

func NewVectorTypeParam(param TypeParam) TypeParam {
	return TypeParam{VectorType, &param, nil}
}

func NewStructTypeParam(tag StructTag) TypeParam {
	return TypeParam{StructType, nil, &tag}
}
