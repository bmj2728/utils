package strutil

import (
	"testing"

	"github.com/bmj2728/utils/pkg/internal/comparison"
)

func TestCompareStringSlices(t *testing.T) {
	tests := []struct {
		name     string
		input1   []string
		input2   []string
		nulls    bool
		expected bool
	}{
		{"CompareStringSlices1", []string{"hello", "world"}, []string{"hello", "world"}, true, true},
		{"CompareStringSlices2", []string{"hello", "world"}, []string{"hello", "world"}, false, true},
		{"CompareStringSlices3", []string{"hello", "world"}, []string{"world", "hello"}, true, true},
		{"CompareStringSlices4", []string{"hello", "world"}, []string{"world", "hello"}, false, true},
		{"CompareStringSlices5", []string{"hello", "world"}, []string{"hello", "world", "hello"}, true, false},
		{"CompareStringSlices6", []string{"hello", "world"}, []string{"hello", "world", "hello"}, false, false},
		{"CompareStringSlices7", nil, nil, false, false},
		{"CompareStringSlices8", nil, nil, true, true},
		{"CompareStringSlices9", []string{"hello", "world"}, nil, false, false},
		{"CompareStringSlices10", nil, []string{"hello", "world"}, true, false},
		{"CompareStringSlices11", nil, []string{"hello", "world"}, false, false},
		{"CompareStringSlices12", []string{"hello", "world"}, nil, true, false},
		{"CompareStringSlices13",
			[]string{"hello", "world", "hello", "world"},
			[]string{"hello", "world", "goodnight", "moon"},
			false,
			false},
		{"CompareStringSlices14", []string{"hello", "world"}, []string{"hello", "hello"}, true, false},
		{"CompareStringSlices15",
			[]string{"to", "be", "or", "not", "to", "be"},
			[]string{"be", "or", "to", "not", "be", "to"},
			false,
			true},
		{"CompareStringSlices16",
			[]string{"to", "be", "or", "not", "to", "be"},
			[]string{"be", "or", "to", "not", "not", "not"},
			false,
			false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := comparison.CompareSlices(tt.input1, tt.input2, tt.nulls)
			result := CompareStringSlices(tt.input1, tt.input2, tt.nulls)
			if helperResult != tt.expected || result != tt.expected {
				t.Errorf("CompareStringSlices - expected %t - got %t / %t",
					tt.expected,
					helperResult,
					result)
			}
		})
	}
}

func TestCompareStringBuilderSlices(t *testing.T) {
	tests := []struct {
		name     string
		input1   []StringBuilder
		input2   []StringBuilder
		nulls    bool
		expected bool
	}{
		{"CompareStringBuilderSlices1",
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			true,
			true},
		{"CompareStringBuilderSlices2",
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			false,
			true},
		{"CompareStringBuilderSlices3",
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			[]StringBuilder{{value: "world"}, {value: "hello"}},
			true,
			true},
		{"CompareStringBuilderSlices4",
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			[]StringBuilder{{value: "world"}, {value: "hello"}},
			false,
			true},
		{"CompareStringBuilderSlices5",
			nil,
			nil,
			true,
			true},
		{"CompareStringBuilderSlices6",
			nil,
			nil,
			false,
			false},
		{"CompareStringBuilderSlices7",
			nil,
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			true,
			false},
		{"CompareStringBuilderSlices8",
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			nil,
			false,
			false},
		{"CompareStringBuilderSlices9",
			[]StringBuilder{{value: "hello"}, {value: "world"}, {value: "hello"}},
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			true,
			false},
		{"CompareStringBuilderSlices10",
			[]StringBuilder{{value: "I'm a sentence"}, {value: "so am I"}},
			[]StringBuilder{{value: "so am I"}, {value: "I'm a sentence"}},
			true,
			true},
		{"CompareStringBuilderSlices11",
			[]StringBuilder{{value: "some"}, {value: "super"}, {value: "good"}, {value: "strings"}},
			[]StringBuilder{{value: "some"}, {value: "other"}, {value: "awesome"}, {value: "strings"}},
			true,
			false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := compareStringBuilderSlices(tt.input1, tt.input2, tt.nulls)
			result := CompareStringBuilderSlices(tt.input1, tt.input2, tt.nulls)
			if helperResult != tt.expected || result != tt.expected {
				t.Errorf("CompareStringBuilderSlices - expected %t - got %t / %t",
					tt.expected,
					helperResult,
					result,
				)
			}
		})
	}
}
