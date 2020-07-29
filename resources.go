package glav

import "strings"

const (
	DfiModule = "DFI"
	DfiStruct = "T"

	CoinsModule = "Coins"
	PriceStruct = "Price"

	BlockModule = "Block"
	BlockStruct = "BlockMetadata"

	TimeModule = "Time"
	TimeStruct = "CurrentTimestamp"
)

func OracleAccessVector(first string, second string) []byte {
	params := [2]TypeParam{
		currencyType(first),
		currencyType(second),
	}
	tag := NewStructTag(codeCoreAddress, CoinsModule, PriceStruct, params[0:])

	return tag.AccessVector()
}

func BlockMetadataVector() []byte {
	empty := [0]TypeParam{}
	tag := NewStructTag(codeCoreAddress, BlockModule, BlockStruct, empty[0:])

	return tag.AccessVector()
}

func TimeMetadataVector() []byte {
	empty := [0]TypeParam{}
	tag := NewStructTag(codeCoreAddress, TimeModule, TimeStruct, empty[0:])

	return tag.AccessVector()
}

func currencyType(curr string) TypeParam {
	empty := [0]TypeParam{}

	curr = strings.ToUpper(curr)
	if curr == "DFI" {
		return NewStructTypeParam(NewStructTag(codeCoreAddress, DfiModule, DfiStruct, empty[0:]))
	}

	return NewStructTypeParam(NewStructTag(codeCoreAddress, CoinsModule, curr, empty[0:]))
}
