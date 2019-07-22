package bigbinary

import (
	"testing"
)

func TestConvertWithBigEndian_ShouldConvertSuccessfully(t *testing.T) {
	got := ConvertWithBigEndian([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb, 0xfa, 0xf9, 0xf8, 0xf7, 0xf6, 0xf5, 0xf4, 0xf3, 0xf2, 0xf1, 0xf0})
	if got.String() != "340277133820332220657323652036036850160" {
		t.Errorf("unexpected error has come: %s", got.String())
	}
}

func TestConvertWithLittleEndian_ShouldConvertSuccessfully(t *testing.T) {
	got := ConvertWithLittleEndian([]byte{0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff})
	if got.String() != "340277133820332220657323652036036850160" {
		t.Errorf("unexpected error has come: %s", got.String())
	}
}
