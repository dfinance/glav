package glav

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestModulePath(t *testing.T) {
	access := ModuleAccessVector(codeCoreAddress, "Block")
	expected, _ := hex.DecodeString("009934e51f2eb6b334a0111bc088588497e2a304aac277526593fd68f816c20417")
	if !reflect.DeepEqual(access, expected) {
		t.Errorf("Expected %X path. Actual: %X", expected, access)
	}
}
