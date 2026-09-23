package main

import (
	"testing"
)

func TestCodec(t *testing.T) {
	tests := []struct {
		name    string
		plain   []byte
		encoded string
	}{
		{"multi-word string with spaces", []byte("a simple test"), "YSBzaW1wbGUgdGVzdA=="},
		{"empty input", []byte(""), ""},
		{"single byte, double padding", []byte("a"), "YQ=="},
		{"two bytes, single padding", []byte("ab"), "YWI="},
		{"1 byte, no leftover", []byte("f"), "Zg=="},
		{"2 bytes, no leftover", []byte("fo"), "Zm8="},
		{"3 bytes, no padding", []byte("foo"), "Zm9v"},
		{"4 bytes", []byte("foob"), "Zm9vYg=="},
		{"5 bytes", []byte("fooba"), "Zm9vYmE="},
		{"6 bytes", []byte("foobar"), "Zm9vYmFy"},
		{"longer ascii string", []byte("hello world"), "aGVsbG8gd29ybGQ="},
		{"null byte, double padding", []byte{0x00}, "AA=="},
		{"two null bytes, single padding", []byte{0x00, 0x00}, "AAA="},
		{"three null bytes, no padding", []byte{0x00, 0x00, 0x00}, "AAAA"},
		{"null byte sequence", []byte{0x00, 0x01, 0x02}, "AAEC"},
		{"mixed binary with high bytes", []byte{0x00, 0xFF, 0x80, 0x7F}, "AP+Afw=="},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Base64Encode(tt.plain)
			if got != tt.encoded {
				t.Errorf("Encode(%q) = %q, want %q", tt.plain, got, tt.encoded)
			}

			got2 := Base64Decode([]byte(tt.encoded))
			if got2 != string(tt.plain) {
				t.Errorf("Decode(%q) = %q, want %q", tt.encoded, got2, tt.plain)
			}
		})
	}
}
