package glav

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestOraclePath(t *testing.T) {
	access := OracleAccessVector("XFI", "BTC")
	expected, _ := hex.DecodeString("018c2f213d25358a39f9370a494dbe4bd80f84734137a01ec8f468c3b2ef16360a")
	if !reflect.DeepEqual(access, expected) {
		t.Errorf("Expected %X path. Actual: %X", expected, access)
	}

	access = OracleAccessVector("BTC", "ETH")
	expected, _ = hex.DecodeString("01a7183ec0c4d32fd9a2705e1e6844035c5238598bf45167742e9db3735af96cc1")
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

func TestBalanceVector(t *testing.T) {
	t.Run("eth denom", func(t *testing.T) {
		access := BalanceVector("eth")
		expected, _ := hex.DecodeString("0138f4f2895881c804de0e57ced1d44f02e976f9c6561c889f7b7eef8e660d2c9a")
		if !reflect.DeepEqual(access, expected) {
			t.Errorf("Expected %X path. Actual: %X", expected, access)
		}
	})

	t.Run("xfi denom", func(t *testing.T) {
		access := BalanceVector("xfi")
		expected, _ := hex.DecodeString("01226844e85ad6e3867f4ff1a4300e71ed6057538631a5a5330512772b7104b585")
		if !reflect.DeepEqual(access, expected) {
			t.Errorf("Expected %X path. Actual: %X", expected, access)
		}
	})
}

func TestCurrencyInfoVector(t *testing.T) {
	t.Run("eth denom", func(t *testing.T) {
		access := CurrencyInfoVector("eth")
		expected, _ := hex.DecodeString("012a00668b5325f832c28a24eb83dffa8295170c80345fbfbf99a5263f962c76f4")
		if !reflect.DeepEqual(access, expected) {
			t.Errorf("Expected %X path. Actual: %X", expected, access)
		}
	})

	t.Run("xfi denom", func(t *testing.T) {
		access := CurrencyInfoVector("xfi")
		expected, _ := hex.DecodeString("01b9ed21c23abf8c7a53fb868a36e106d45394c30127fb722f8dd2d45aae719585")
		if !reflect.DeepEqual(access, expected) {
			t.Errorf("Expected %X path. Actual: %X", expected, access)
		}
	})
}
