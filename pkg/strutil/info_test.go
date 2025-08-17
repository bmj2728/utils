package strutil

import (
	"testing"

	"github.com/bmj2728/utils/pkg/internal/comparison"
	"github.com/bmj2728/utils/pkg/internal/convert"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		substr    string
		normalize bool
		expected  int
	}{
		{"Index1", "Hello World", "World", true, 6},
		{"Index2", "Hello World", "Hello", true, 0},
		{"Index3", "Hello World", "Hello World", false, 0},
		{"Index4", normalizeUnicode("Café Gopher Gopher Gopher Gopher Gopher Gopher Café", NFD),
			"Go", false, 7},
		{"Index5", "Café Gopher Gopher Gopher Gopher Gopher Gopher Café",
			"Go", true, 6},
		{"Index6", "", "Hello World", false, -1},
		{"Index7", "Hello World", "", false, -1},
		{"Index8", "", "", true, -1},
		{"Index9", "Hello World", "nope", true, -1},
		{"Index10", "Hello World", "nope", false, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := index(tt.input, tt.substr, tt.normalize)
			result := Index(tt.input, tt.substr, tt.normalize)
			builderResult := New(tt.input).Index(tt.substr, tt.normalize)
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("Expected index to be %d / %d / %d, got %d",
					helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestIndexAll(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		substr    string
		normalize bool
		expected  []int
	}{
		{"IndexAll1", "Hello World, Hello", "Hello", true, []int{0, 13}},
		{"IndexAll2", "Hello World", "World", true, []int{6}},
		{"IndexAll3", "Hello World", "Hello World", false, []int{0}},
		{"IndexAll4", normalizeUnicode("Café Gopher Gopher Gopher Gopher Gopher Gopher Café", NFD),
			"Go", false, []int{7, 14, 21, 28, 35, 42}},
		{"IndexAll5", "Café Gopher Gopher Gopher Gopher Gopher Gopher Café",
			"Go", true, []int{6, 13, 20, 27, 34, 41}},
		{"IndexAll6", "", "Hello World", false, nil},
		{"IndexAll7", "Hello World", "", false, nil},
		{"IndexAll8", "", "", true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := indexAll(tt.input, tt.substr, tt.normalize)
			result := IndexAll(tt.input, tt.substr, tt.normalize)
			builderResult := New(tt.input).IndexAll(tt.substr, tt.normalize)
			if !comparison.CompareSlices(helperResult, tt.expected, true) ||
				!comparison.CompareSlices(result, tt.expected, true) ||
				!comparison.CompareSlices(builderResult, tt.expected, true) {
				t.Errorf("Expected index to be %v. Got %v / %v / %v",
					tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestLastIndex(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		substr    string
		normalize bool
		expected  int
	}{
		{"LastIndex1", "Hello World", "World", true, 6},
		{"LastIndex2", "Hello World", "Hello", true, 0},
		{"LastIndex3", "Hello World", "Hello World", false, 0},
		{"LastIndex4", "Café Gopher Gopher Gopher Gopher Gopher Gopher Café", "Go",
			true, 41},
		{"LastIndex5", "Café Gopher Gopher Gopher Gopher Gopher Gopher Café", "Go",
			false, 41},
		{"LastIndex6", NormalizeUnicode("Café Gopher Gopher Gopher Gopher Gopher Gopher Café", NFD),
			"Go", true, 41},
		{"LastIndex7", NormalizeUnicode("Café Gopher Gopher Gopher Gopher Gopher Gopher Café", NFD),
			"Go", false, 42},
		{"LastIndex8", "", "Hello World", false, -1},
		{"LastIndex9", "Hello World", "", false, -1},
		{"LastIndex10", "", "", true, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := lastIndex(tt.input, tt.substr, tt.normalize)
			result := LastIndex(tt.input, tt.substr, tt.normalize)
			builderResult := New(tt.input).LastIndex(tt.substr, tt.normalize)
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("Got: %d / %d / %d - Expected: %d", helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestIndexRune(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		r         rune
		wantIndex int
	}{
		{"IndexRune1", "Hello World", 'o', 4},
		{"IndexRune2", "Hello World", ' ', 5},
		{"IndexRune3", "Hello World", 'x', -1},
		{"IndexRune4", NormalizeUnicode("Café Gopher", NFD), 'é', -1},
		{"IndexRune5", "Café Gopher", 'é', 3},
		{"IndexRune6", "", 'H', -1},
		{"IndexRune7", "Hello World", 0, -1},
		{"IndexRune8", "", 0, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := indexRune(tt.input, tt.r)
			result := IndexRune(tt.input, tt.r)
			builderResult := New(tt.input).IndexRune(tt.r)
			if helperResult != tt.wantIndex || result != tt.wantIndex || builderResult != tt.wantIndex {
				t.Errorf("Got: %d / %d / %d | Wanted: %d",
					helperResult, result, builderResult, tt.wantIndex)
			}
		})
	}
}

func TestIndexRuneAll(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		r         rune
		wantIndex []int
	}{
		{"IndexRuneAll1", "Hello World", 'o', []int{4, 7}},
		{"IndexRuneAllRev", "Hello World", 'o', []int{7, 4}},
		{"IndexRuneAll2", "Hello World", ' ', []int{5}},
		{"IndexRuneAll3", "Hello World", 'x', nil},
		{"IndexRuneAll4", "Café Gopher Gopher Gopher Gopher Gopher Gopher Café",
			'é', []int{3, 51}},
		{"IndexRuneAll5", NormalizeUnicode("Café Gopher Gopher Gopher Gopher Gopher Gopher Café", NFD),
			'é', nil},
		{"IndexRuneAll6", "", 'x', nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := indexAllRune(tt.input, tt.r)
			result := IndexAllRune(tt.input, tt.r)
			builderResult := New(tt.input).IndexAllRune(tt.r)

			if !comparison.CompareSlices(helperResult, tt.wantIndex, true) ||
				!comparison.CompareSlices(result, tt.wantIndex, true) ||
				!comparison.CompareSlices(builderResult, tt.wantIndex, true) {
				t.Errorf("Got: %v / %v / %v | Wanted: %v",
					helperResult, result, builderResult, tt.wantIndex)
			}
		})
	}
}

func TestLastIndexRune(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		r         rune
		wantIndex int
	}{
		{"LastIndexRune1", "Hello World", 'o', 7},
		{"LastIndexRune2", "Hello World", ' ', 5},
		{"LastIndexRune3", "Hello World", 'x', -1},
		{"LastIndexRune4", NormalizeUnicode("Café Gopher", NFD), 'é', -1},
		{"LastIndexRune5", "Café Gopher Gopher Gopher Gopher Gopher Gopher Café", 'é', 51},
		{"LastIndexRune6", "", 'H', -1},
		{"LastIndexRune7", "Hello World", 0, -1},
		{"LastIndexRune8", "", 0, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := lastIndexRune(tt.input, tt.r)
			result := LastIndexRune(tt.input, tt.r)
			builderResult := New(tt.input).LastIndexRune(tt.r)
			if helperResult != tt.wantIndex || result != tt.wantIndex || builderResult != tt.wantIndex {
				t.Errorf("Got: %d / %d / %d | Wanted: %d",
					helperResult, result, builderResult, tt.wantIndex)
			}
		})
	}
}

var (
	invalidUTF8 = string([]byte{0xff, 0xfe})
)

func TestUTF32NormalizedIndex(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		substr    string
		wantIndex int
	}{
		{"UTF32NormalizedIndex1",
			"Café Gopher Gopher Gopher Gopher Gopher Gopher Café", "Go", 5},
		{"UTF32NormalizedIndex2", "Hello World", "Hello", 0},
		{"UTF32NormalizedIndex3", "Hello World", "Hello World", 0},
		{"UTF32NormalizedIndex4",
			"Café Gopher Gopher Gopher Gopher Gopher Gopher Café", "Café", 0},
		{"UTF32NormalizedIndex5",
			normalizeUnicode("Café Gopher Gopher Gopher Gopher Gopher Gopher Café", NFD),
			"Café",
			0},
		{"UTF32NormalizedIndex6",
			"",
			"Café",
			-1},
		{"UTF32NormalizedIndex8",
			"Hello World",
			"",
			-1},
		{"UTF32NormalizedIndex8",
			"",
			"",
			-1},
		{"UTF32InvalidUTF81", invalidUTF8, "x", -1},
		{"UTF32InvalidUTF82", "hello world", invalidUTF8, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := utf32NormalizedIndex(tt.input, tt.substr)
			result := UTF32NormalizedIndex(tt.input, tt.substr)
			builderResult := New(tt.input).UTF32NormalizedIndex(tt.substr)
			if helperResult != tt.wantIndex || result != tt.wantIndex || builderResult != tt.wantIndex {
				t.Errorf("Got: %d / %d / %d | Wanted: %d",
					helperResult, result, builderResult, tt.wantIndex)
			}
		})
	}
}

func TestUTF32NormalizedLastIndex(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		substr    string
		wantIndex int
	}{
		{"UTF32NormalizedIndex1",
			"Café Gopher Gopher Gopher Gopher Gopher Gopher Café", "Go", 40},
		{"UTF32NormalizedIndex2", "Hello World", "Hello", 0},
		{"UTF32NormalizedIndex3", "Hello World", "Hello World", 0},
		{"UTF32NormalizedIndex4",
			"Café Gopher Gopher Gopher Gopher Gopher Gopher Café", "Café", 47},
		{"UTF32NormalizedIndex5",
			normalizeUnicode("Café Gopher Gopher Gopher Gopher Gopher Gopher Café", NFD),
			"Café",
			47},
		{"UTF32NormalizedIndex6",
			"",
			"Café",
			-1},
		{"UTF32NormalizedIndex8",
			"Hello World",
			"",
			-1},
		{"UTF32NormalizedIndex8",
			"",
			"",
			-1},
		{"UTF32InvalidUTF81", invalidUTF8, "x", -1},
		{"UTF32InvalidUTF82", "hello world", invalidUTF8, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := utf32NormalizedLastIndex(tt.input, tt.substr)
			result := UTF32NormalizedLastIndex(tt.input, tt.substr)
			builderResult := New(tt.input).UTF32NormalizedLastIndex(tt.substr)
			if helperResult != tt.wantIndex || result != tt.wantIndex || builderResult != tt.wantIndex {
				t.Errorf("Got: %d / %d / %d | Wanted: %d",
					helperResult, result, builderResult, tt.wantIndex)
			}
		})
	}
}

var (
	utf32In, _  = convert.ToUTF32("Hello World")
	utf32Sub, _ = convert.ToUTF32("Hello")
)

func TestUTF32NormalizedIndexAll(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		substr    string
		wantIndex []int
	}{
		{"UTF32All1", "", "", nil},
		{"UTF32All2", "Hello World", "Hello", []int{0}},
		{"UTF32All3", "Hello World", "Hello World", []int{0}},
		{"UTF32All4",
			"Café Gopher Gopher Gopher Gopher Gopher Gopher Café",
			"Go", []int{5, 12, 19, 26, 33, 40}},
		{"UTF32All5", "Hello World", "John", nil},
		{"UTF32All6", "Hello World", "\r\n\t", nil},
		{"UTF32All7",
			"Café Gopher Gopher Gopher Gopher Gopher Gopher Café", "Café", []int{0, 47}},
		{"UTF32All8",
			normalizeUnicode("Café Gopher Gopher Gopher Gopher Gopher Gopher Café", NFD),
			"Café",
			[]int{0, 47}},
		{"UTF32All9",
			"",
			"Café",
			nil},
		{"UTF32All10",
			"Hello World",
			"",
			nil},
		{"UTF32All11",
			"",
			"",
			nil},
		{"UTF32All12", utf32In, "Hello", nil},
		{"UTF32All13", "Hello World", utf32Sub, nil},
		{"UTF32InvalidUTF81", invalidUTF8, "x", nil},
		{"UTF32InvalidUTF82", "hello world", invalidUTF8, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := utf32NormalizedIndexAll(tt.input, tt.substr)
			result := UTF32NormalizedIndexAll(tt.input, tt.substr)
			builderResult := New(tt.input).UTF32NormalizedIndexAll(tt.substr)
			if !comparison.CompareSlices(helperResult, tt.wantIndex, true) ||
				!comparison.CompareSlices(result, tt.wantIndex, true) ||
				!comparison.CompareSlices(builderResult, tt.wantIndex, true) {
				t.Errorf("Got: %v / %v / %v | Wanted: %v",
					helperResult, result, builderResult, tt.wantIndex)
			}
		})
	}
}
