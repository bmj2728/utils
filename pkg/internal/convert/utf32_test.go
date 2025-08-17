package convert

import (
	"testing"

	utf33 "golang.org/x/text/encoding/unicode/utf32"
)

func TestToUTF32(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"ToUTF32-1", "Café"},
		{"ToUTF32-2", "Hello World!"},
		{"ToUTF32-3", "1234567890"},
		{"ToUTF32-4", "abc123"},
		{"ToUTF32-5", "Go Bananas!!"},
		{"ToUTF32Empty", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utf32, err := ToUTF32(tt.input)
			if err != nil {
				t.Errorf("ToUTF32() error = %v", err)
			}
			expected, err := utf33.UTF32(utf33.BigEndian, utf33.IgnoreBOM).NewEncoder().String(tt.input)
			if err != nil {
				t.Errorf("ToUTF32() error = %v", err)
			}
			if utf32 != expected {
				t.Errorf("ToUTF32() = %v, want %v", utf32, expected)
			}
		})
	}
}
