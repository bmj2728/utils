package strutil

import (
	"testing"

	"github.com/bmj2728/utils/pkg/internal/comparison"
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
