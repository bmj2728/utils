package strutil

import (
	"maps"
	"testing"

	"utils/pkg/internal/comparison"
	"utils/pkg/internal/errors"
)

var (
	shinMapPrt = New("ababababababababa").
			WithComparisonManager().
			Shingle(2).
			ShingleSlice(2).
			GetComparisonManager().
			GetShingleResultsMap()

	shinMapPrntShort    = "Shingle Map (ababababababababa/2):\n2 shingles found\n"
	shinMapPrntLongOpt1 = "Results for Shingle Map\nWord: ababababababababa\n" +
		"N-Gram Length: 2\nShingles Map:\nab: 8\nba: 8\n"
	shinMapPrntLongOpt2 = "Results for Shingle Map\nWord: ababababababababa\n" +
		"N-Gram Length: 2\nShingles Map:\nba: 8\nab: 8\n"
	shinSlicePrntShort    = "Shingle Slice (ababababababababa/2):\n2 shingles found\n"
	shinSlicePrntLongOpt1 = "Results for Shingle Slice\nWord: ababababababababa\n" +
		"N-Gram Length: 2\nShingles Slice:\nab\nba\n"
	shinSlicePrntLongOpt2 = "Results for Shingle Slice\nWord: ababababababababa\n" +
		"N-Gram Length: 2\nShingles Slice:\nba\nab\n"

	shinMap1 = New("ooooo").
			WithComparisonManager().
			Shingle(2).
			ShingleSlice(2).
			GetComparisonManager().
			GetShingleResultsMap()

	shinMap2 = New("ooooo").
			WithComparisonManager().
			Shingle(2).
			ShingleSlice(2).
			GetComparisonManager().
			GetShingleResultsMap()

	shinMap3 = New("aaaa").
			WithComparisonManager().
			Shingle(1).
			Shingle(2).
			ShingleSlice(1).
			GetComparisonManager().
			GetShingleResultsMap()

	shinMap4 = New("abab").
			WithComparisonManager().
			ShingleSlice(1).
			GetComparisonManager().
			GetShingleResultsMap()

	shinMap5 = New("none").
			WithComparisonManager().
			GetComparisonManager().
			GetShingleResultsMap()
)

func TestShingleMapResultGetters(t *testing.T) {
	tests := []struct {
		name        string
		resMap      ShingleResultsMap
		shinResType ShingleResultType
		input       string
		nGramLength int
		result      map[string]int
		err         error
	}{
		{"ShingleMap1", shinMap1, ShinglesMap, "ooooo", 2, map[string]int{"oo": 4}, nil},
		{"ShingleMap2", shinMap2, ShinglesMap, "ooooo", 2, map[string]int{"oo": 4}, nil},
		{"ShingleMap3", shinMap3, ShinglesMap, "aaaa", 1, map[string]int{"a": 4}, nil},
		{"ShingleMap4", shinMap3, ShinglesMap, "aaaa", 2, map[string]int{"aa": 3}, nil},
		{"ShingleMap5", shinMap5, ShinglesMap, "none", 1, nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resMap[tt.shinResType][tt.nGramLength] == nil && tt.input != "none" {
				t.Errorf("ShingleMapResultGetters() = %v, want %v", tt.resMap[tt.shinResType][tt.nGramLength], nil)
			} else if tt.resMap[tt.shinResType][tt.nGramLength] != nil {
				casted, ok := (*tt.resMap[tt.shinResType][tt.nGramLength]).(*ShingleMapResult)
				if (!ok && tt.input != "none") || (ok && tt.input == "none") {
					t.Errorf("ShingleMapResultGetters() = %v, want %v", ok, true)
				}
				if casted.GetType() != tt.shinResType {
					t.Errorf("ShingleMapResultGetters() = %v, want %v", casted.GetType(), tt.shinResType)
				}
				if casted.GetTypeName() != tt.shinResType.String() {
					t.Errorf("ShingleMapResultGetters() = %v, want %v", casted.GetTypeName(), tt.shinResType.String())
				}
				if casted.GetInput() != tt.input {
					t.Errorf("ShingleMapResultGetters() = %v, want %v", casted.GetInput(), tt.input)
				}
				if casted.GetNgramLength() != tt.nGramLength {
					t.Errorf("ShingleMapResultGetters() = %v, want %v", casted.GetNgramLength(), tt.nGramLength)
				}
				if !errors.CompareErrors(casted.GetError(), tt.err) {
					t.Errorf("ShingleMapResultGetters() = %v, want %v", casted.GetError(), tt.err)
				}
				if !maps.Equal(casted.GetShinglesMap(), tt.result) {
					t.Errorf("ShingleMapResultGetters() = %v, want %v", casted.GetShinglesMap(), tt.result)
				}
			}
		})
	}
}

func TestShingleSliceResultGetters(t *testing.T) {
	tests := []struct {
		name        string
		resMap      ShingleResultsMap
		shinResType ShingleResultType
		input       string
		nGramLength int
		result      []string
		err         error
	}{
		{"ShingleMap1", shinMap1,
			ShinglesSlice, "ooooo", 2, []string{"oo"}, nil},
		{"ShingleMap2", shinMap2,
			ShinglesSlice, "ooooo", 2, []string{"oo"}, nil},
		{"ShingleMap3", shinMap3,
			ShinglesSlice, "aaaa", 1, []string{"a"}, nil},
		{"ShingleMap4", shinMap4,
			ShinglesSlice, "abab", 1, []string{"a", "b"}, nil},
		{"ShingleMap5", shinMap5,
			ShinglesSlice, "none", 1, nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resMap[tt.shinResType][tt.nGramLength] == nil && tt.input != "none" {
				t.Errorf("ShingleSliceResultGetters() = %v, want %v", tt.resMap[tt.shinResType][tt.nGramLength], nil)
			} else if tt.resMap[tt.shinResType][tt.nGramLength] != nil {
				casted, ok := (*tt.resMap[tt.shinResType][tt.nGramLength]).(*ShingleSliceResult)
				if (!ok && tt.input != "none") || (ok && tt.input == "none") {
					t.Errorf("ShingleSliceResultGetters() = %v, want %v", ok, true)
				}
				if casted.GetType() != tt.shinResType {
					t.Errorf("ShingleSliceResultGetters() = %v, want %v", casted.GetType(), tt.shinResType)
				}
				if casted.GetTypeName() != tt.shinResType.String() {
					t.Errorf("ShingleSliceResultGetters() = %v, want %v", casted.GetTypeName(), tt.shinResType.String())
				}
				if casted.GetInput() != tt.input {
					t.Errorf("ShingleSliceResultGetters() = %v, want %v", casted.GetInput(), tt.input)
				}
				if casted.GetNgramLength() != tt.nGramLength {
					t.Errorf("ShingleSliceResultGetters() = %v, want %v", casted.GetNgramLength(), tt.nGramLength)
				}
				if !errors.CompareErrors(casted.GetError(), tt.err) {
					t.Errorf("ShingleSliceResultGetters() = %v, want %v", casted.GetError(), tt.err)
				}
				if !comparison.CompareStringSlices(casted.GetShinglesSlice(), tt.result, false) {
					t.Errorf("ShingleSliceResultGetters() = %v, want %v", casted.GetShinglesSlice(), tt.result)
				}
			}
		})
	}
}

func TestShingleResultPrint(t *testing.T) {
	tests := []struct {
		name     string
		resMap   ShingleResultsMap
		expected bool
	}{
		{"ShingleMap1", shinMapPrt, true},
		{"ShingleMap2", shinMap1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resM := tt.resMap[ShinglesMap][2]
			castedM, ok := (*resM).(*ShingleMapResult)
			if !ok {
				t.Errorf("Panic casting to map")
			}
			resS := tt.resMap[ShinglesSlice][2]
			castedS, ok := (*resS).(*ShingleSliceResult)
			if !ok {
				t.Errorf("Panic casting to slice")
			}
			mS := formatShingleResultOutput(castedM, false)
			sS := formatShingleResultOutput(castedS, false)
			mL := formatShingleResultOutput(castedM, true)
			sL := formatShingleResultOutput(castedS, true)
			isShortMatch := true
			isLongMatch := true
			if mS != shinMapPrntShort || sS != shinSlicePrntShort {
				isShortMatch = false
			}
			if (mL != shinMapPrntLongOpt1 && mL != shinMapPrntLongOpt2) ||
				(sL != shinSlicePrntLongOpt1 && sL != shinSlicePrntLongOpt2) {
				isLongMatch = false
			}
			if isShortMatch != tt.expected {
				t.Errorf("ShingleResultPrintMap() = %v, want %v", mS, shinMapPrntShort)
			}
			if isLongMatch != tt.expected {
				t.Errorf("ShingleResultPrintMap() = %v, want %v", mL, shinMapPrntLongOpt1)
			}
			if isShortMatch != tt.expected {
				t.Errorf("ShingleResultPrintMap() = %v, want %v", sS, shinSlicePrntShort)
			}
			castedS.Print(true)
			castedM.Print(true)
			castedS.Print(false)
			castedM.Print(false)
		})
	}
}

func TestShingleResultIsMatchMap(t *testing.T) {
	tests := []struct {
		name     string
		resMap1  ShingleResultsMap
		rt1      ShingleResultType
		n1       int
		resMap2  ShingleResultsMap
		rt2      ShingleResultType
		n2       int
		expected bool
	}{
		{"ShingleMap1", shinMap1, ShinglesMap, 2, shinMap2, ShinglesMap, 2, true},
		{"ShingleMap2", shinMap1, ShinglesMap, 2, shinMap3, ShinglesMap, 2, false},
		{"ShingleMap3", shinMap1, ShinglesMap, 2, shinMap3, ShinglesMap, 1, false},
		{"ShingleMap4", shinMap1, ShinglesMap, 2, shinMap3, ShinglesMap, 2, false},
		{"ShingleMap5", shinMap1, ShinglesMap, 2, shinMap1, ShinglesMap, 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr1, ok := (*tt.resMap1[tt.rt1][tt.n1]).(*ShingleMapResult)
			if !ok {
				t.Errorf("Panic casting to map")
			}
			sr2, ok := (*tt.resMap2[tt.rt2][tt.n2]).(*ShingleMapResult)
			if !ok {
				t.Errorf("Panic casting to map")
			}
			if sr1.IsMatch(sr2) != tt.expected {
				t.Errorf("ShingleResultIsMatchMap() = %v, want %v", sr1.IsMatch(sr2), tt.expected)
			}
		})
	}
}

func TestShingleResultIsMatchSlice(t *testing.T) {
	tests := []struct {
		name     string
		resMap1  ShingleResultsMap
		rt1      ShingleResultType
		n1       int
		resMap2  ShingleResultsMap
		rt2      ShingleResultType
		n2       int
		expected bool
	}{
		{"ShingleSlice1", shinMap1, ShinglesSlice, 2, shinMap2, ShinglesSlice, 2, true},
		{"ShingleSlice2", shinMap1, ShinglesSlice, 2, shinMap3, ShinglesSlice, 1, false},
		{"ShingleSlice3", shinMap1, ShinglesSlice, 2, shinMap4, ShinglesSlice, 1, false},
		{"ShingleSlice4", shinMap1, ShinglesSlice, 2, shinMap3, ShinglesSlice, 1, false},
		{"ShingleSlice5", shinMap1, ShinglesSlice, 2, shinMap1, ShinglesSlice, 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr1, ok := (*tt.resMap1[tt.rt1][tt.n1]).(*ShingleSliceResult)
			if !ok {
				t.Errorf("Panic casting to map")
			}
			sr2, ok := (*tt.resMap2[tt.rt2][tt.n2]).(*ShingleSliceResult)
			if !ok {
				t.Errorf("Panic casting to map")
			}
			if sr1.IsMatch(sr2) != tt.expected {
				t.Errorf("ShingleResultIsMatchSlice() = %v, want %v", sr1.IsMatch(sr2), tt.expected)
			}
		})
	}
}
