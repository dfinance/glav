package libra_access_vector

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestOraclePath(t *testing.T) {
	access := OracleAccessVector("USD", "BTC")
	expected, _ := hex.DecodeString("0148b6fd935bea9299b852c31f00b4bcf6e919dc6712f4bf36d5b3e9e4d4fce0d2")
	if !reflect.DeepEqual(access, expected) {
		t.Errorf("Expected %X path. Actual: %X", expected, access)
	}
}

func TestBlockMetadataVector(t *testing.T) {
	access := BlockMetadataVector()
	expected, _ := hex.DecodeString("01ada6f79e8eddfdf986687174de1000df3c5fa45e9965ece812fed33332ec543a")
	if !reflect.DeepEqual(access, expected) {
		t.Errorf("Expected %X path. Actual: %X", expected, access)
	}
}

func TestTimeMetadataVector(t *testing.T) {
	access := TimeMetadataVector()
	expected, _ := hex.DecodeString("01bb4980dfba817aaa64cb4b3a75ee1e1d8352111dead64c5c4f05fb7b4c85bb3e")
	if !reflect.DeepEqual(access, expected) {
		t.Errorf("Expected %X path. Actual: %X", expected, access)
	}
}
