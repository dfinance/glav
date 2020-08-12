package glav

import "strings"

const (
	AccountModule  = "Account"
	BlockModule    = "Block"
	CoinsModule    = "Coins"
	XfiModule      = "XFI"
	DfinanceModule = "Dfinance"
	TimeModule     = "Time"

	BalanceStruct = "Balance"
	BlockStruct   = "BlockMetadata"
	XfiStruct     = "T"
	PriceStruct   = "Price"
	TimeStruct    = "CurrentTimestamp"
	InfoStruct    = "Info"
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

func BalanceVector(curr string) []byte {
	tag := NewStructTag(codeCoreAddress, AccountModule, BalanceStruct, []TypeParam{currencyType(curr)})

	return tag.AccessVector()
}

func CurrencyInfoVector(curr string) []byte {
	tag := NewStructTag(codeCoreAddress, DfinanceModule, InfoStruct, []TypeParam{currencyType(curr)})

	return tag.AccessVector()
}

func currencyType(curr string) TypeParam {
	empty := [0]TypeParam{}

	curr = strings.ToUpper(curr)
	if curr == "XFI" {
		return NewStructTypeParam(NewStructTag(codeCoreAddress, XfiModule, XfiStruct, empty[0:]))
	}

	return NewStructTypeParam(NewStructTag(codeCoreAddress, CoinsModule, curr, empty[0:]))
}
