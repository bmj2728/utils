package strutil

import "testing"

var (
	testMap1 = New("Hello, World!").
			WithComparisonManager().
			LevenshteinDistance("Hello World").
			LevenshteinDistance("Hi").
			DamerauLevenshteinDistance("Hello World").
			OSADamerauLevenshteinDistance("Hello World").
			LCS("Hello World").
			LCSEditDistance("Hello World").
			JaroSimilarity("Hello World").
			JaroWinklerSimilarity("Hello World").
			JaccardSimilarity("Hello World", 3).
			CosineSimilarity("Hello World", 3).
			SorensenDiceCoefficient("Hello World", 3).
			GetComparisonManager().
			GetComparisonResultsMap()

	testMap2 = New("Hello, World!").
			WithComparisonManager().
			LevenshteinDistance("Hello World").
			LevenshteinDistance("Hi").
			DamerauLevenshteinDistance("Hello World").
			OSADamerauLevenshteinDistance("Hello World").
			LCS("Hello World").
			LCSEditDistance("Hello World").
			JaroSimilarity("Hello World").
			JaroWinklerSimilarity("Hello World").
			JaccardSimilarity("Hello World", 3).
			CosineSimilarity("Hello World", 3).
			SorensenDiceCoefficient("Hello World", 3).
			GetComparisonManager().
			GetComparisonResultsMap()

	testMap3 = New("Hello, World!").
			WithComparisonManager().
			LevenshteinDistance("Hello World").
			LevenshteinDistance("Yo").
			DamerauLevenshteinDistance("Hello World").
			OSADamerauLevenshteinDistance("Hello World").
			LCS("Hello World").
			LCSEditDistance("Hello World").
			JaroSimilarity("Hello World").
			JaroWinklerSimilarity("Hello World").
			JaccardSimilarity("Hello World", 3).
			CosineSimilarity("Hello World", 3).
			SorensenDiceCoefficient("Hello World", 3).
			GetComparisonManager().
			GetComparisonResultsMap()

	testMap4 = New("Hello, World!").
			WithComparisonManager().
			LevenshteinDistance("Hello World").
			DamerauLevenshteinDistance("Hello World").
			GetComparisonManager().
			GetComparisonResultsMap()

	testMap5 = New("Hello, World!").
			WithComparisonManager().
			GetComparisonManager().
			GetComparisonResultsMap()
)

func TestComparisonResultsMapTypeCount(t *testing.T) {
	tests := []struct {
		name     string
		compMap  ComparisonResultsMap
		expected int
	}{
		{"TypeCount1", testMap1, 10},
		{"TypeCount2", testMap2, 10},
		{"TypeCount3", testMap3, 10},
		{"TypeCount4", testMap4, 2},
		{"TypeCount5", testMap5, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.compMap.TypeCount() != test.expected {
				t.Errorf("Expected type count to be %d, got %d", test.expected, test.compMap.TypeCount())
			}
		})
	}
}

func TestComparisonResultsMapEntryCount(t *testing.T) {
	tests := []struct {
		name     string
		compMap  ComparisonResultsMap
		expected int
	}{
		{"EntryCount1", testMap1, 11},
		{"EntryCount2", testMap2, 11},
		{"EntryCount3", testMap3, 11},
		{"EntryCount4", testMap4, 2},
		{"EntryCount5", testMap5, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.compMap.EntryCount() != test.expected {
				t.Errorf("Expected entry count to be %d, got %d", test.expected, test.compMap.EntryCount())
			}
		})
	}
}

func TestComparisonResultsMapIsMatch(t *testing.T) {
	tests := []struct {
		name     string
		compMap1 ComparisonResultsMap
		compMap2 ComparisonResultsMap
		expected bool
	}{
		{"IsMatch1", testMap1, testMap2, true},
		{"IsMatch2", testMap1, testMap3, false},
		{"IsMatch3", testMap1, testMap4, false},
		{"IsMatch4", testMap1, testMap5, false},
		{"IsMatch5", testMap1, testMap1, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.compMap1.IsMatch(test.compMap2) != test.expected {
				t.Errorf("Expected match to be %t, got %t", test.expected, test.compMap1.IsMatch(test.compMap2))
			}
		})
	}
}

func TestComparisonResultsMapGetCopy(t *testing.T) {
	tests := []struct {
		name     string
		compMap  ComparisonResultsMap
		expected bool
	}{
		{"GetCopy1", testMap1, true},
		{"GetCopy2", testMap2, true},
		{"GetCopy3", testMap3, true},
		{"GetCopy4", testMap4, true},
		{"GetCopy5", testMap5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dup := tt.compMap.GetCopy()
			if dup.TypeCount() != tt.compMap.TypeCount() {
				t.Errorf("Expected type count to be %d, got %d", tt.compMap.TypeCount(), dup.TypeCount())
			}
			if dup.EntryCount() != tt.compMap.EntryCount() {
				t.Errorf("Expected entry count to be %d, got %d", tt.compMap.EntryCount(), dup.EntryCount())
			}
			if !dup.IsMatch(tt.compMap) {
				t.Errorf("Expected match to be %t, got %t", true, dup.IsMatch(tt.compMap))
			}
		})
	}
}

func TestComparisonResultsMap_Get(t *testing.T) {
	tests := []struct {
		name     string
		compMap  ComparisonResultsMap
		compType ComparisonResultType
		compStr  string
		expected bool
	}{
		{"MapGet1", testMap1, LevDist, "Hello World", true},
		{"MapGet2", testMap1, LevDist, "Hi", true},
		{"MapGet3", testMap1, LevDist, "Yo", false},
		{"MapGet4", testMap5, LevDist, "Hello World", false},
		{"MapGet5", testMap2, JaroWinklerSim, "Hello World", true},
		{"MapGet6", testMap2, LCSLength, "Hello World", true},
		{"MapGet7", testMap3, LevDist, "Yo", true},
		{"MapGet8", testMap4, DamLevDist, "Hello World", true},
		{"MapGet9", testMap5, JaroSim, "Hello World", false},
		{"MapGet10", testMap5, QGramDist, "Hello World", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (tt.compMap.Get(tt.compType, tt.compStr) != nil) != tt.expected {
				t.Errorf("Expected match to be %t, got %t",
					tt.expected, tt.compMap.Get(tt.compType, tt.compStr) != nil)
			}
		})
	}
}

func TestComparisonResultsMapGetByType(t *testing.T) {
	tests := []struct {
		name     string
		compMap  ComparisonResultsMap
		compType ComparisonResultType
		expected int
	}{
		{"MapGetByCompStr1", testMap1, LevDist, 2},
		{"MapGetByCompStr2", testMap2, LevDist, 2},
		{"MapGetByCompStr3", testMap3, LevDist, 2},
		{"MapGetByCompStr4", testMap4, LevDist, 1},
		{"MapGetByCompStr5", testMap5, LevDist, 0},
		{"MapGetByCompStr6", testMap1, HammingDist, 0},
		{"MapGetByCompStr7", testMap2, CosineSim, 1},
		{"MapGetByCompStr8", testMap3, LCSLength, 1},
		{"MapGetByCompStr9", testMap4, DamLevDist, 1},
		{"MapGetByCompStr10", testMap5, SorensenDiceCo, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.compMap.GetByType(tt.compType)
			if len(res) != tt.expected {
				t.Errorf("Expected match to be %d, got %d", tt.expected, len(res))
			}
		})
	}
}

func TestComparisonResultsMapFilterByType(t *testing.T) {
	tests := []struct {
		name     string
		compMap  ComparisonResultsMap
		compType ComparisonResultType
		expected int
	}{
		{"MapFilterByType1", testMap1, LevDist, 2},
		{"MapFilterByType2", testMap2, LevDist, 2},
		{"MapFilterByType3", testMap3, LevDist, 2},
		{"MapFilterByType4", testMap4, LevDist, 1},
		{"MapFilterByType5", testMap5, LevDist, 0},
		{"MapFilterByType6", testMap1, HammingDist, 0},
		{"MapFilterByType7", testMap2, CosineSim, 1},
		{"MapFilterByType8", testMap3, LCSLength, 1},
		{"MapFilterByType9", testMap4, DamLevDist, 1},
		{"MapFilterByType10", testMap5, SorensenDiceCo, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.compMap.FilterByType(tt.compType)
			if res.EntryCount() != tt.expected {
				t.Errorf("Expected match to be %d, got %d", tt.expected, res.EntryCount())
			}
		})
	}
}

func TestComparisonResultsMapGetByCompStr(t *testing.T) {
	tests := []struct {
		name     string
		compMap  ComparisonResultsMap
		compStr  string
		expected int
	}{
		{"MapGetByCompStr1", testMap1, "Hello World", 10},
		{"MapGetByCompStr2", testMap2, "Hello World", 10},
		{"MapGetByCompStr3", testMap3, "Hello World", 10},
		{"MapGetByCompStr4", testMap4, "Hello World", 2},
		{"MapGetByCompStr5", testMap5, "Hello World", 0},
		{"MapGetByCompStr6", testMap1, "Hi", 1},
		{"MapGetByCompStr7", testMap2, "Hi", 1},
		{"MapGetByCompStr8", testMap3, "Hi", 0},
		{"MapGetByCompStr9", testMap4, "Hi", 0},
		{"MapGetByCompStr10", testMap5, "Hi", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.compMap.GetByComparisonString(tt.compStr)
			if len(res) != tt.expected {
				t.Errorf("Expected match to be %d, got %d", tt.expected, len(res))
			}
		})
	}
}

func TestComparisonResultsMapFilterByCompStr(t *testing.T) {
	tests := []struct {
		name     string
		compMap  ComparisonResultsMap
		compStr  string
		expected int
	}{
		{"MapGetByCompStr1", testMap1, "Hello World", 10},
		{"MapGetByCompStr2", testMap2, "Hello World", 10},
		{"MapGetByCompStr3", testMap3, "Hello World", 10},
		{"MapGetByCompStr4", testMap4, "Hello World", 2},
		{"MapGetByCompStr5", testMap5, "Hello World", 0},
		{"MapGetByCompStr6", testMap1, "Hi", 1},
		{"MapGetByCompStr7", testMap2, "Hi", 1},
		{"MapGetByCompStr8", testMap3, "Hi", 0},
		{"MapGetByCompStr9", testMap4, "Hi", 0},
		{"MapGetByCompStr10", testMap5, "Hi", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.compMap.FilterByComparisonString(tt.compStr)
			if res.EntryCount() != tt.expected {
				t.Errorf("Expected match to be %d, got %d", tt.expected, res.EntryCount())
			}
		})
	}
}

var (
	testFormatOpt1 = "***Comparison Results for Levenshtein Distance***\n\n" +
		"Levenshtein Distance (\"Hello, World!\"/\"Hello World\"): 2\n\n" +
		"***Comparison Results for Sorensen-Dice Coefficient***\n\n" +
		"Sorensen-Dice Coefficient (\"Hello, World!\"/\"Hello World\"): 0.7\n\n"

	testFormatOpt2 = "***Comparison Results for Sorensen-Dice Coefficient***\n\n" +
		"Sorensen-Dice Coefficient (\"Hello, World!\"/\"Hello World\"): 0.7\n\n" +
		"***Comparison Results for Levenshtein Distance***\n\n" +
		"Levenshtein Distance (\"Hello, World!\"/\"Hello World\"): 2\n\n"

	testFormatLongOpt1 = "***Comparison Results for Levenshtein Distance***\n\n" +
		"Comparison: Levenshtein Distance\nFirst String: Hello, World!\n" +
		"Second String: Hello World\n" +
		"Score: 2\n\n" +
		"***Comparison Results for Sorensen-Dice Coefficient***\n\n" +
		"Comparison: Sorensen-Dice Coefficient\n" +
		"First String: Hello, World!\n" +
		"Second String: Hello World\n" +
		"Score: 0.7\n\n"

	testFormatLongOpt2 = "***Comparison Results for Sorensen-Dice Coefficient***\n\n" +
		"Comparison: Sorensen-Dice Coefficient\n" +
		"First String: Hello, World!\n" +
		"Second String: Hello World\n" +
		"Score: 0.7\n\n" +
		"***Comparison Results for Levenshtein Distance***\n\n" +
		"Comparison: Levenshtein Distance\n" +
		"First String: Hello, World!\n" +
		"Second String: Hello World\n" +
		"Score: 2\n\n"

	testFormatMap1 = New("Hello, World!").
			WithComparisonManager().
			LevenshteinDistance("Hello World").
			SorensenDiceCoefficient("Hello World", 3).
			GetComparisonManager().
			GetComparisonResultsMap()

	testFormatMap2 = New("Hello, World!").
			WithComparisonManager().
			DamerauLevenshteinDistance("Hello World").
			CosineSimilarity("Hello World", 3).
			GetComparisonManager().
			GetComparisonResultsMap()
)

func TestComparisonResultsMapFormat(t *testing.T) {
	tests := []struct {
		name     string
		compMap  ComparisonResultsMap
		expected bool
	}{
		{"Format1", testFormatMap1, true},
		{"Format2", testFormatMap2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var isShortMatch bool
			var isLongMatch bool
			resShort := formatComparisonMapOutput(tt.compMap, false)
			resLong := formatComparisonMapOutput(tt.compMap, true)
			if resShort == testFormatOpt1 ||
				resShort == testFormatOpt2 {
				isShortMatch = true
			} else {
				isShortMatch = false
			}
			if resLong == testFormatLongOpt1 ||
				resLong == testFormatLongOpt2 {
				isLongMatch = true
			} else {
				isLongMatch = false
			}
			if isShortMatch != tt.expected {
				t.Errorf("Got:\n%s\nExpected:\n%s\nor\n%s", resShort, testFormatOpt1, testFormatOpt2)
			}
			if isLongMatch != tt.expected {
				t.Errorf("Got:\n%s\nExpected:\n%s\nor\n%s", resLong, testFormatLongOpt1, testFormatLongOpt2)
			}
			tt.compMap.Print(true)
			tt.compMap.Print(false)
		})
	}
}
