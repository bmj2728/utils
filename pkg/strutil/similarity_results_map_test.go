package strutil

import "testing"

var (
	simMap1 = New("Hello, World!").
		WithComparisonManager().
		Similarity("Hello World", Levenshtein).
		Similarity("Hi", Levenshtein).
		Similarity("Hello World", DamerauLevenshtein).
		Similarity("Hello World", OSADamerauLevenshtein).
		Similarity("Hello World", Lcs).
		Similarity("Hello World", Jaro).
		Similarity("Hello World", JaroWinkler).
		Similarity("Hello World", Jaccard).
		Similarity("Hello World", Cosine).
		Similarity("Hello World", SorensenDice).
		Similarity("Hello World", QGram).
		Similarity("Hello World", Hamming).
		GetComparisonManager().
		GetSimilarityResultsMap()

	simMap2 = New("Hello, World!").
		WithComparisonManager().
		Similarity("Hello World", Levenshtein).
		Similarity("Hi", Levenshtein).
		Similarity("Hello World", DamerauLevenshtein).
		Similarity("Hello World", OSADamerauLevenshtein).
		Similarity("Hello World", Lcs).
		Similarity("Hello World", Jaro).
		Similarity("Hello World", JaroWinkler).
		Similarity("Hello World", Jaccard).
		Similarity("Hello World", Cosine).
		Similarity("Hello World", SorensenDice).
		Similarity("Hello World", QGram).
		Similarity("Hello World", Hamming).
		GetComparisonManager().
		GetSimilarityResultsMap()

	simMap3 = New("Hello, World!").
		WithComparisonManager().
		Similarity("Hello World", Levenshtein).
		Similarity("Yo", Levenshtein).
		Similarity("Hello World", DamerauLevenshtein).
		Similarity("Hello World", OSADamerauLevenshtein).
		Similarity("Hello World", Lcs).
		Similarity("Hello World", Jaro).
		Similarity("Hello World", JaroWinkler).
		Similarity("Hello World", Jaccard).
		Similarity("Hello World", Cosine).
		Similarity("Hello World", SorensenDice).
		Similarity("Hello World", QGram).
		Similarity("Hello World", Hamming).
		GetComparisonManager().
		GetSimilarityResultsMap()

	simMap4 = New("Hello, World!").
		WithComparisonManager().
		Similarity("Hello World", Levenshtein).
		Similarity("Hello World", DamerauLevenshtein).
		GetComparisonManager().
		GetSimilarityResultsMap()

	simMap5 = New("Hello, World!").
		WithComparisonManager().
		GetComparisonManager().
		GetSimilarityResultsMap()
)

func TestSimilarityResultsMapTypeCount(t *testing.T) {
	tests := []struct {
		name   string
		simMap SimilarityResultsMap
		want   int
	}{
		{"SimMap1", simMap1, 11},
		{"SimMap2", simMap2, 11},
		{"SimMap3", simMap3, 11},
		{"SimMap4", simMap4, 2},
		{"SimMap5", simMap5, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.simMap.TypeCount() != tt.want {
				t.Errorf("SimilarityResultsMapTypeCount() = %v, want %v", tt.simMap.TypeCount(), tt.want)
			}
		})
	}
}

func TestSimilarityResultsMapEntryCount(t *testing.T) {
	tests := []struct {
		name   string
		simMap SimilarityResultsMap
		want   int
	}{
		{"SimMap1", simMap1, 12},
		{"SimMap2", simMap2, 12},
		{"SimMap3", simMap3, 12},
		{"SimMap4", simMap4, 2},
		{"SimMap5", simMap5, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.simMap.EntryCount() != tt.want {
				t.Errorf("SimilarityResultsMapEntryCount() = %v, want %v", tt.simMap.EntryCount(), tt.want)
			}
		})
	}
}

func TestSimilarityResultsMapIsMatch(t *testing.T) {
	tests := []struct {
		name    string
		simMap1 SimilarityResultsMap
		simMap2 SimilarityResultsMap
		want    bool
	}{
		{"SimMapIsMatch1", simMap1, simMap2, true},
		{"SimMapIsMatch2", simMap1, simMap3, false},
		{"SimMapIsMatch3", simMap1, simMap4, false},
		{"SimMapIsMatch4", simMap1, simMap5, false},
		{"SimMapIsMatch5", simMap1, simMap1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.simMap1.IsMatch(tt.simMap2) != tt.want {
				t.Errorf("SimilarityResultsMapIsMatch() = %v, want %v", tt.simMap1.IsMatch(tt.simMap2), tt.want)
			}
		})
	}
}

func TestSimilarityResultsMapGetCopy(t *testing.T) {
	tests := []struct {
		name     string
		simMap   SimilarityResultsMap
		expected bool
	}{
		{"SimMapGetCopy1", simMap1, true},
		{"SimMapGetCopy2", simMap2, true},
		{"SimMapGetCopy3", simMap3, true},
		{"SimMapGetCopy4", simMap4, true},
		{"SimMapGetCopy5", simMap5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dup := tt.simMap.GetCopy()
			if dup.TypeCount() != tt.simMap.TypeCount() {
				t.Errorf("SimilarityResultsMapGetCopy() = %v, want %v", dup.TypeCount(), tt.simMap.TypeCount())
			}
			if dup.EntryCount() != tt.simMap.EntryCount() {
				t.Errorf("SimilarityResultsMapGetCopy() = %v, want %v", dup.EntryCount(), tt.simMap.EntryCount())
			}
			if tt.expected && !tt.simMap.IsMatch(dup) {
				t.Errorf("SimilarityResultsMapGetCopy() = %v, want %v", dup.IsMatch(tt.simMap), true)
			}
		})
	}
}

func TestSimilarityResultsMapGet(t *testing.T) {
	tests := []struct {
		name     string
		simMap   SimilarityResultsMap
		algo     Algorithm
		comStr   string
		expected bool
	}{
		{"SimMapGet1", simMap1, Levenshtein, "Hello World", true},
		{"SimMapGet2", simMap1, Levenshtein, "Hi", true},
		{"SimMapGet3", simMap1, Levenshtein, "Yo", false},
		{"SimMapGet4", simMap5, Levenshtein, "Hello World", false},
		{"SimMapGet5", simMap2, JaroWinkler, "Hello World", true},
		{"SimMapGet6", simMap2, Lcs, "Hello World", true},
		{"SimMapGet7", simMap3, Levenshtein, "Yo", true},
		{"SimMapGet8", simMap4, DamerauLevenshtein, "Hello World", true},
		{"SimMapGet9", simMap5, Jaro, "Hello World", false},
		{"SimMapGet10", simMap5, QGram, "Hello World", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (tt.simMap.Get(tt.algo, tt.comStr) != nil) != tt.expected {
				t.Errorf("SimilarityResultsMapGet() = %v, want %v",
					tt.simMap.Get(tt.algo, tt.comStr), tt.expected)
			}
		})
	}
}

func TestSimilarityResultsMapGetByType(t *testing.T) {
	tests := []struct {
		name     string
		simMap   SimilarityResultsMap
		algo     Algorithm
		expected int
	}{
		{"SimMapGetByTypes1", simMap1, Levenshtein, 2},
		{"SimMapGetByTypes2", simMap2, Levenshtein, 2},
		{"SimMapGetByTypes3", simMap3, Levenshtein, 2},
		{"SimMapGetByTypes4", simMap4, Levenshtein, 1},
		{"SimMapGetByTypes5", simMap5, Levenshtein, 0},
		{"SimMapGetByTypes6", simMap1, Hamming, 1},
		{"SimMapGetByTypes7", simMap2, Cosine, 1},
		{"SimMapGetByTypes8", simMap3, Lcs, 1},
		{"SimMapGetByTypes9", simMap4, DamerauLevenshtein, 1},
		{"SimMapGetByTypes10", simMap5, SorensenDice, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.simMap.GetByType(tt.algo)
			if len(res) != tt.expected {
				t.Errorf("SimilarityResultsMapGetByTypes() = %v, want %v", len(res), tt.expected)
			}
		})
	}
}

func TestSimilarityResultsMapFilterByType(t *testing.T) {
	tests := []struct {
		name     string
		simMap   SimilarityResultsMap
		algo     Algorithm
		expected int
	}{
		{"SimMapGetByTypes1", simMap1, Levenshtein, 2},
		{"SimMapGetByTypes2", simMap2, Levenshtein, 2},
		{"SimMapGetByTypes3", simMap3, Levenshtein, 2},
		{"SimMapGetByTypes4", simMap4, Levenshtein, 1},
		{"SimMapGetByTypes5", simMap5, Levenshtein, 0},
		{"SimMapGetByTypes6", simMap1, Hamming, 1},
		{"SimMapGetByTypes7", simMap2, Cosine, 1},
		{"SimMapGetByTypes8", simMap3, Lcs, 1},
		{"SimMapGetByTypes9", simMap4, DamerauLevenshtein, 1},
		{"SimMapGetByTypes10", simMap5, SorensenDice, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.simMap.FilterByType(tt.algo)
			if res.EntryCount() != tt.expected {
				t.Errorf("SimilarityResultsMapFilterByType() = %v, want %v", len(res), tt.expected)
			}
		})
	}
}

func TestSimilarityResultsMapGetByComparisonString(t *testing.T) {
	tests := []struct {
		name     string
		simMap   SimilarityResultsMap
		comStr   string
		expected int
	}{
		{"SimMapGetByTypes1", simMap1, "Hello World", 11},
		{"SimMapGetByTypes2", simMap2, "Hello World", 11},
		{"SimMapGetByTypes3", simMap3, "Hello World", 11},
		{"SimMapGetByTypes4", simMap4, "Hello World", 2},
		{"SimMapGetByTypes5", simMap5, "Hello World", 0},
		{"SimMapGetByTypes6", simMap1, "Hi", 1},
		{"SimMapGetByTypes7", simMap2, "Hi", 1},
		{"SimMapGetByTypes8", simMap3, "Hi", 0},
		{"SimMapGetByTypes9", simMap4, "Hi", 0},
		{"SimMapGetByTypes10", simMap5, "Hi", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.simMap.GetByComparisonString(tt.comStr)
			if len(res) != tt.expected {
				t.Errorf("SimilarityResultsMap_GetByComparisonString() = %v, want %v", len(res), tt.expected)
			}
		})
	}
}

func TestSimilarityResultsMapFilterByComparisonString(t *testing.T) {
	tests := []struct {
		name     string
		simMap   SimilarityResultsMap
		comStr   string
		expected int
	}{
		{"SimMapGetByTypes1", simMap1, "Hello World", 11},
		{"SimMapGetByTypes2", simMap2, "Hello World", 11},
		{"SimMapGetByTypes3", simMap3, "Hello World", 11},
		{"SimMapGetByTypes4", simMap4, "Hello World", 2},
		{"SimMapGetByTypes5", simMap5, "Hello World", 0},
		{"SimMapGetByTypes6", simMap1, "Hi", 1},
		{"SimMapGetByTypes7", simMap2, "Hi", 1},
		{"SimMapGetByTypes8", simMap3, "Hi", 0},
		{"SimMapGetByTypes9", simMap4, "Hi", 0},
		{"SimMapGetByTypes10", simMap5, "Hi", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.simMap.FilterByComparisonString(tt.comStr)
			if res.EntryCount() != tt.expected {
				t.Errorf("SimilarityResultsMap_FilterByComparisonString() = %v, want %v",
					res.EntryCount(), tt.expected)
			}
		})
	}
}

var (
	testSimMap1 = New("Hello, World!").
			WithComparisonManager().
			Similarity("Hello World", Levenshtein).
			Similarity("Hello World", JaroWinkler).
			GetComparisonManager().
			GetSimilarityResultsMap()

	testSimMap2 = New("Hello, World!").
			WithComparisonManager().
			Similarity("Hello World", DamerauLevenshtein).
			Similarity("Hello World", Cosine).
			GetComparisonManager().
			GetSimilarityResultsMap()

	short1 = "***Similarity Results for Levenshtein***\n\n" +
		"Levenshtein: 0.846154\n\n***Similarity Results for Jaro-Winkler***\n\nJaro-Winkler: 0.969231\n\n"
	short2 = "***Similarity Results for Jaro-Winkler***\n\n" +
		"Jaro-Winkler: 0.969231\n\n***Similarity Results for Levenshtein***\n\nLevenshtein: 0.846154\n\n"
	long1 = "***Similarity Results for Levenshtein***\n\n" +
		"Comparison: Levenshtein\nFirst String: Hello, World!\nSecond String: Hello World\nScore: 0.846154\n\n" +
		"***Similarity Results for Jaro-Winkler***\n\n" +
		"Comparison: Jaro-Winkler\nFirst String: Hello, World!\nSecond String: Hello World\nScore: 0.969231\n\n"
	long2 = "***Similarity Results for Jaro-Winkler***\n\n" +
		"Comparison: Jaro-Winkler\nFirst String: Hello, World!\nSecond String: Hello World\nScore: 0.969231\n\n" +
		"***Similarity Results for Levenshtein***\n\n" +
		"Comparison: Levenshtein\nFirst String: Hello, World!\nSecond String: Hello World\nScore: 0.846154\n\n"
)

func TestSimilarityResultsMapPrint(t *testing.T) {
	tests := []struct {
		name     string
		simMap   SimilarityResultsMap
		expected bool
	}{
		{"SimMapPrint1", testSimMap1, true},
		{"SimMapPrint2", testSimMap2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var isShortMatch bool
			var isLongMatch bool
			resShort := formatSimilarityMapOutput(tt.simMap, false)
			resLong := formatSimilarityMapOutput(tt.simMap, true)
			if resShort == short1 || resShort == short2 {
				isShortMatch = true
			}
			if resLong == long1 || resLong == long2 {
				isLongMatch = true
			}
			if isShortMatch != tt.expected {
				t.Errorf("SimilarityResultsMapPrint() = %v, want %v", isShortMatch, tt.expected)
			}
			if isLongMatch != tt.expected {
				t.Errorf("SimilarityResultsMapPrint() = %v, want %v", isLongMatch, tt.expected)
			}
			tt.simMap.Print(true)
			tt.simMap.Print(false)
		})
	}
}
