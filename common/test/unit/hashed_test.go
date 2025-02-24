package unit_test

import (
	"rupamic-arch/common"
	"testing"
)

func Test(t *testing.T) {
	enc := common.Encrypt("MyJit")
	if enc == "" {
		t.Errorf("Test failed: %v", enc)
	}
	matcked, err := common.Decrypt(enc, "MyJit")
	if err != nil {
		t.Errorf("Test failed: got %v, want %v", err, nil)
	}
	if !matcked {
		t.Errorf("Test failed: ")
	}
}
func TestUnyuiu(t *testing.T) {
	t.Skip("mera marji")
	enc := common.Encrypt("MyJit")
	if enc == "" {
		t.Errorf("Test failed: %v", enc)
	}
	matcked, err := common.Decrypt(enc, "MyJit")
	if err != nil {
		t.Errorf("Test failed: got %v, want %v", err, nil)
	}
	if !matcked {
		t.Errorf("Test failed: ")
	}
}
