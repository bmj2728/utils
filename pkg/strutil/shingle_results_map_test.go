package strutil

import "testing"

var (
	testShinMap1 = New("aaaaa").
			WithComparisonManager().
			Shingle(5).
			Shingle(4).
			Shingle(3).
			Shingle(2).
			Shingle(1).
			ShingleSlice(5).
			ShingleSlice(4).
			ShingleSlice(3).
			ShingleSlice(2).
			ShingleSlice(1).
			GetComparisonManager().
			GetShingleResultsMap()

	testShinMap2 = New("aaaaa").
			WithComparisonManager().
			Shingle(5).
			Shingle(4).
			Shingle(3).
			Shingle(2).
			Shingle(1).
			ShingleSlice(5).
			ShingleSlice(4).
			ShingleSlice(3).
			ShingleSlice(2).
			ShingleSlice(1).
			GetComparisonManager().
			GetShingleResultsMap()

	testShinMap3 = New("ooooo").
			WithComparisonManager().
			Shingle(5).
			Shingle(4).
			Shingle(3).
			Shingle(2).
			Shingle(1).
			ShingleSlice(5).
			ShingleSlice(4).
			ShingleSlice(3).
			ShingleSlice(2).
			ShingleSlice(1).
			GetComparisonManager().
			GetShingleResultsMap()

	testShinMap4 = New("xxxxx").
			WithComparisonManager().
			Shingle(1).
			ShingleSlice(1).
			GetComparisonManager().
			GetShingleResultsMap()

	testShinMap5 = New("aaaaa").
			WithComparisonManager().
			GetComparisonManager().
			GetShingleResultsMap()

	shinMapPrtLongOpt1 = "Shingle Results for Shingle Map\n" +
		"Results for Shingle Map\n" +
		"Word: xxxxx\n" +
		"N-Gram Length: 1\n" +
		"Shingles Map:\n" +
		"x: 5\n\n" +
		"Shingle Results for Shingle Slice\n" +
		"Results for Shingle Slice\n" +
		"Word: xxxxx\n" +
		"N-Gram Length: 1\n" +
		"Shingles Slice:\n" +
		"x\n\n"

	shinMapPrtLongOpt2 = "Shingle Results for Shingle Slice\n" +
		"Results for Shingle Slice\n" +
		"Word: xxxxx\n" +
		"N-Gram Length: 1\n" +
		"Shingles Slice:\n" +
		"x\n\n" +
		"Shingle Results for Shingle Map\n" +
		"Results for Shingle Map\n" +
		"Word: xxxxx\n" +
		"N-Gram Length: 1\n" +
		"Shingles Map:\n" +
		"x: 5\n\n"

	shinMapPrtShortOpt1 = "Shingle Results for Shingle Map\n" +
		"Shingle Map (xxxxx/1):\n" +
		"1 shingles found\n\n" +
		"Shingle Results for Shingle Slice\n" +
		"Shingle Slice (xxxxx/1):\n" +
		"1 shingles found\n\n"

	shinMapPrtShortOpt2 = "Shingle Results for Shingle Slice\n" +
		"Shingle Slice (xxxxx/1):\n" +
		"1 shingles found\n\n" +
		"Shingle Results for Shingle Map\n" +
		"Shingle Map (xxxxx/1):\n" +
		"1 shingles found\n\n"
)

func TestShingleResultsMapTypeCount(t *testing.T) {
	tests := []struct {
		name     string
		resMap   ShingleResultsMap
		expected int
	}{
		{"ShingleMap1", testShinMap1, 2},
		{"ShingleMap2", testShinMap2, 2},
		{"ShingleMap3", testShinMap3, 2},
		{"ShingleMap4", testShinMap4, 2},
		{"ShingleMap5", testShinMap5, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resMap.TypeCount() != tt.expected {
				t.Errorf("ShingleResultsMapTypeCount() = %v, want %v", tt.resMap.TypeCount(), tt.expected)
			}
		})
	}
}

func TestShingleResultsMapEntryCount(t *testing.T) {
	tests := []struct {
		name     string
		resMap   ShingleResultsMap
		expected int
	}{
		{"ShingleMap1", testShinMap1, 10},
		{"ShingleMap2", testShinMap2, 10},
		{"ShingleMap3", testShinMap3, 10},
		{"ShingleMap4", testShinMap4, 2},
		{"ShingleMap5", testShinMap5, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resMap.EntryCount() != tt.expected {
				t.Errorf("ShingleResultsMapEntryCount() = %v, want %v", tt.resMap.EntryCount(), tt.expected)
			}
		})
	}
}

func TestShingleResultsMapIsMatch(t *testing.T) {
	tests := []struct {
		name     string
		resMap1  ShingleResultsMap
		resMap2  ShingleResultsMap
		expected bool
	}{
		{"ShingleMap1", testShinMap1, testShinMap2, true},
		{"ShingleMap2", testShinMap1, testShinMap3, false},
		{"ShingleMap3", testShinMap1, testShinMap4, false},
		{"ShingleMap4", testShinMap1, testShinMap5, false},
		{"ShingleMap5", testShinMap1, testShinMap1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resMap1.IsMatch(tt.resMap2) != tt.expected {
				t.Errorf("ShingleResultsMapIsMatch() = %v, want %v", tt.resMap1.IsMatch(tt.resMap2), tt.expected)
			}
		})
	}
}

func TestShingleResultsMapGetCopy(t *testing.T) {
	tests := []struct {
		name     string
		resMap   ShingleResultsMap
		expected bool
	}{
		{"ShingleMap1", testShinMap1, true},
		{"ShingleMap2", testShinMap2, true},
		{"ShingleMap3", testShinMap3, true},
		{"ShingleMap4", testShinMap4, true},
		{"ShingleMap5", testShinMap5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dup := tt.resMap.GetCopy()
			if dup.TypeCount() != tt.resMap.TypeCount() {
				t.Errorf("ShingleResultsMapGetCopy() = %v, want %v", dup.TypeCount(), tt.resMap.TypeCount())
			}
			if dup.EntryCount() != tt.resMap.EntryCount() {
				t.Errorf("ShingleResultsMapGetCopy() = %v, want %v", dup.EntryCount(), tt.resMap.EntryCount())
			}
			if !tt.resMap.IsMatch(dup) {
				t.Errorf("ShingleResultsMapGetCopy() = %v, want %v", dup.IsMatch(tt.resMap), true)
			}
		})
	}
}

func TestShingleResultsMapGet(t *testing.T) {
	tests := []struct {
		name     string
		resMap   ShingleResultsMap
		resType  ShingleResultType
		nGramLen int
		expected bool
	}{
		{"ShingleMap1", testShinMap1, ShinglesMap, 5, true},
		{"ShingleMap2", testShinMap1, ShinglesMap, 4, true},
		{"ShingleMap3", testShinMap1, ShinglesMap, 3, true},
		{"ShingleMap4", testShinMap1, ShinglesMap, 2, true},
		{"ShingleMap5", testShinMap1, ShinglesMap, 1, true},
		{"ShingleMap6", testShinMap1, ShinglesMap, 0, false},
		{"ShingleMap7", testShinMap1, ShinglesSlice, 5, true},
		{"ShingleMap8", testShinMap1, ShinglesSlice, 4, true},
		{"ShingleMap9", testShinMap1, ShinglesSlice, 3, true},
		{"ShingleMap10", testShinMap1, ShinglesSlice, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resMap.Get(tt.resType, tt.nGramLen) == nil && tt.expected {
				t.Errorf("ShingleResultsMapGet() = %v, want %v", tt.resMap.Get(tt.resType, tt.nGramLen), nil)
			} else if tt.resMap.Get(tt.resType, tt.nGramLen) != nil && !tt.expected {
				t.Errorf("ShingleResultsMapGet() = %v, want %v", tt.resMap.Get(tt.resType, tt.nGramLen), nil)
			}
		})
	}
}

func TestShingleResultsMapGetByType(t *testing.T) {
	tests := []struct {
		name     string
		resMap   ShingleResultsMap
		resType  ShingleResultType
		expected int
	}{
		{"ShingleMap1", testShinMap1, ShinglesMap, 5},
		{"ShingleMap2", testShinMap1, ShinglesSlice, 5},
		{"ShingleMap3", testShinMap2, ShinglesMap, 5},
		{"ShingleMap4", testShinMap2, ShinglesSlice, 5},
		{"ShingleMap5", testShinMap3, ShinglesMap, 5},
		{"ShingleMap6", testShinMap3, ShinglesSlice, 5},
		{"ShingleMap7", testShinMap4, ShinglesMap, 1},
		{"ShingleMap8", testShinMap4, ShinglesSlice, 1},
		{"ShingleMap9", testShinMap5, ShinglesMap, 0},
		{"ShingleMap10", testShinMap5, ShinglesSlice, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.resMap.GetByType(tt.resType)) != tt.expected {
				t.Errorf("ShingleResultsMapGetByType() = %v, want %v", tt.resMap.GetByType(tt.resType), tt.expected)
			}
		})
	}
}

func TestShingleResultsMapFilterByType(t *testing.T) {
	tests := []struct {
		name     string
		resMap   ShingleResultsMap
		resType  ShingleResultType
		expected int
	}{
		{"ShingleMap1", testShinMap1, ShinglesMap, 5},
		{"ShingleMap2", testShinMap1, ShinglesSlice, 5},
		{"ShingleMap3", testShinMap2, ShinglesMap, 5},
		{"ShingleMap4", testShinMap2, ShinglesSlice, 5},
		{"ShingleMap5", testShinMap3, ShinglesMap, 5},
		{"ShingleMap6", testShinMap3, ShinglesSlice, 5},
		{"ShingleMap7", testShinMap4, ShinglesMap, 1},
		{"ShingleMap8", testShinMap4, ShinglesSlice, 1},
		{"ShingleMap9", testShinMap5, ShinglesMap, 0},
		{"ShingleMap10", testShinMap5, ShinglesSlice, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resMap.FilterByType(tt.resType).EntryCount() != tt.expected {
				t.Errorf("ShingleResultsMapFilterByType() = %v, want %v", tt.resMap.FilterByType(tt.resType), tt.expected)
			}
		})
	}
}

func TestShingleResultsMapGetByNGramLength(t *testing.T) {
	tests := []struct {
		name        string
		resMap      ShingleResultsMap
		ngramLength int
		expected    int
	}{
		{"ShingleMap1", testShinMap1, 5, 2},
		{"ShingleMap2", testShinMap1, 4, 2},
		{"ShingleMap3", testShinMap1, 3, 2},
		{"ShingleMap4", testShinMap1, 2, 2},
		{"ShingleMap5", testShinMap1, 1, 2},
		{"ShingleMap6", testShinMap1, 0, 0},
		{"ShingleMap7", testShinMap1, 6, 0},
		{"ShingleMap8", testShinMap2, 5, 2},
		{"ShingleMap9", testShinMap2, 4, 2},
		{"ShingleMap10", testShinMap2, 3, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.resMap.GetByNGramLength(tt.ngramLength)) != tt.expected {
				t.Errorf("ShingleResultsMap_GetByNGramLength() = %d, want %d",
					len(tt.resMap.GetByNGramLength(tt.ngramLength)), tt.expected)
			}
		})
	}
}

func TestShingleResultsMapFilterByNGramLength(t *testing.T) {
	tests := []struct {
		name        string
		resMap      ShingleResultsMap
		ngramLength int
		expected    int
	}{
		{"ShingleMap1", testShinMap1, 5, 2},
		{"ShingleMap2", testShinMap1, 4, 2},
		{"ShingleMap3", testShinMap1, 3, 2},
		{"ShingleMap4", testShinMap1, 2, 2},
		{"ShingleMap5", testShinMap1, 1, 2},
		{"ShingleMap6", testShinMap1, 0, 0},
		{"ShingleMap7", testShinMap1, 6, 0},
		{"ShingleMap8", testShinMap2, 5, 2},
		{"ShingleMap9", testShinMap2, 4, 2},
		{"ShingleMap10", testShinMap2, 3, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resMap.FilterByNGramLength(tt.ngramLength).EntryCount() != tt.expected {
				t.Errorf("ShingleResultsMap_GetByNGramLength() = %d, want %d",
					tt.resMap.FilterByNGramLength(tt.ngramLength).EntryCount(), tt.expected)
			}
		})
	}
}

func TestShingleResultsMapPrint(t *testing.T) {
	tests := []struct {
		name     string
		resMap   ShingleResultsMap
		expected bool
	}{
		{"ShingleMap1", testShinMap4, true},
		{"ShingleMap2", testShinMap1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isShortMatch := true
			isLongMatch := true
			short := formatShingleResultsMapOutput(tt.resMap, false)
			long := formatShingleResultsMapOutput(tt.resMap, true)
			if short != shinMapPrtShortOpt1 && short != shinMapPrtShortOpt2 {
				isShortMatch = false
			}
			if isShortMatch != tt.expected {
				t.Errorf("ShingleResultsMapPrint() = %s, want %s/%s",
					short, shinMapPrtShortOpt1, shinMapPrtShortOpt2)
			}
			if long != shinMapPrtLongOpt1 && long != shinMapPrtLongOpt2 {
				isLongMatch = false
			}
			if isLongMatch != tt.expected {
				t.Errorf("ShingleResultsMapPrint() = %s, want %s/%s",
					long, shinMapPrtLongOpt1, shinMapPrtLongOpt2)
			}
			tt.resMap.Print(true)
			tt.resMap.Print(false)
		})
	}
}
