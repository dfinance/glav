package libra_access_vector

const CurrencyModule = "Currency"
const PriceStruct = "Price"

const BlockModule = "Block"
const BlockStruct = "BlockMetadata"

const TimeModule = "Time"
const TimeStruct = "CurrentTimestamp"

func OracleAccessVector(first string, second string) []byte {
	empty := [0]TypeParam{}

	params := [2]TypeParam{
		Struct(Tag(CoreCodeAddress(), CurrencyModule, first, empty[0:])),
		Struct(Tag(CoreCodeAddress(), CurrencyModule, second, empty[0:])),
	}
	tag := Tag(CoreCodeAddress(), CurrencyModule, PriceStruct, params[0:])
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
