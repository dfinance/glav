package libra_access_vector

import "strings"

const DfiModule = "DFI"
const DfiStruct = "T"

const CoinsModule = "Coins"
const PriceStruct = "Price"

const BlockModule = "Block"
const BlockStruct = "BlockMetadata"

const TimeModule = "Time"
const TimeStruct = "CurrentTimestamp"

func OracleAccessVector(first string, second string) []byte {
	params := [2]TypeParam{
		currencyType(first),
		currencyType(second),
	}
	tag := Tag(CoreCodeAddress(), CoinsModule, PriceStruct, params[0:])
	return tag.AccessVector()
}

func BlockMetadataVector() []byte {
	empty := [0]TypeParam{}
	tag := Tag(CoreCodeAddress(), BlockModule, BlockStruct, empty[0:])
	return tag.AccessVector()
}

func TimeMetadataVector() []byte {
	empty := [0]TypeParam{}
	tag := Tag(CoreCodeAddress(), TimeModule, TimeStruct, empty[0:])
	return tag.AccessVector()
}

func currencyType(curr string) TypeParam {
	empty := [0]TypeParam{}

	curr = strings.ToUpper(curr)
	if curr == "DFI" {
		return Struct(Tag(CoreCodeAddress(), DfiModule, DfiStruct, empty[0:]))
	} else {
		return Struct(Tag(CoreCodeAddress(), CoinsModule, curr, empty[0:]))
	}
}
