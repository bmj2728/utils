package strutil

import (
	"testing"
	"unicode"
)

func TestGetRuneSet(t *testing.T) {
	tests := []struct {
		name   string
		rsName string
		rs     RuneSet
	}{
		{"GetAlpha", "Letter", Letter},
		{"GetDigit", "Digit", Digit},
		{"GetWhiteSpaceChars", "WhiteSpaceChars", WhiteSpace},
		{"GetPrintable", "Printable", Printable},
		{"GetBad", "Fake", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := GetRuneSet(tt.rsName)
			if set != tt.rs {
				t.Errorf("Expected rune set to be %v, got %v", tt.rs, set)
			}
		})
	}
}

func TestRegisterRuneSet(t *testing.T) {
	tests := []struct {
		name   string
		rsName string
		fn     func(rune) bool
	}{
		{"RegisterAlphaDigit",
			"AlphaDigit",
			func(r rune) bool {
				if unicode.IsLetter(r) || unicode.IsDigit(r) {
					return true
				}
				return false
			}},
		{"RegisterAlphaDigitPunct",
			"AlphaDigitPunct",
			func(r rune) bool {
				if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsPunct(r) {
					return true
				}
				return false
			}},
		{"RegisterAlphaDigit",
			"DigitPunct",
			func(r rune) bool {
				if unicode.IsPunct(r) || unicode.IsDigit(r) {
					return true
				}
				return false
			}},
		{"NotPrintable",
			"NoPrint",
			func(r rune) bool {
				return !unicode.IsPrint(r)
			}},
		{"RegisterAlphaPunct",
			"AlphaPunct",
			func(r rune) bool {
				if unicode.IsLetter(r) || unicode.IsPunct(r) {
					return true
				}
				return false
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := RegisterRuneSet(tt.rsName, tt.fn)
			byName := GetRuneSet(tt.rsName)
			if set != byName {
				t.Errorf("Expected rune set to be %v, got %v", byName, set)
			}
		})
	}
}

func TestCheckRunesBySetName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		rsName   string
		expected bool
	}{
		{"Alpha", "hello123", "Letter", true},
		{"Punctuation", "hello123", "Punctuation", false},
		{"Digit", "hello123", "Digit", true},
		{"Printable", "hello123", "Printable", true},
		{"Upper", "hello123", "Upper", false},
		{"Graphic", "hello123", "Graphic", true},
		{"Bad", "hello123", "Fake", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if CheckRunesBySetName(tt.input, tt.rsName) != tt.expected {
				t.Errorf("Expected match to be %t, got %t",
					tt.expected, CheckRunesBySetName(tt.input, tt.rsName))
			}
		})
	}
}

func TestCheckRunes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		rs       RuneSet
		expected bool
	}{
		{"Alpha", "hello123", Letter, true},
		{"Punctuation", "hello123", Punctuation, false},
		{"Digit", "hello123", Digit, true},
		{"Printable", "hello123", Printable, true},
		{"Upper", "hello123", Upper, false},
		{"Graphic", "hello123", Graphic, true},
		{"Bad", "hello123", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if CheckRunes(tt.input, tt.rs) != tt.expected {
				t.Errorf("Expected match to be %t, got %t",
					tt.expected, CheckRunes(tt.input, tt.rs))
			}
		})
	}
}
