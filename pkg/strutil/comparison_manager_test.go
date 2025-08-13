package strutil

import "testing"

var (
	testMan1 = New("Hello, World!").
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
			Shingle(2).
			Shingle(1).
			ShingleSlice(2).
			LCSBacktrack("Hello World").
			LCSBacktrackAll("Hello World").
			LCSDiff("Hello World").
			Similarity("Hello World", Levenshtein).
			Similarity("Hello World", DamerauLevenshtein).
			Similarity("Hello World", OSADamerauLevenshtein).
			Similarity("Hello World", Jaro).
			Similarity("Hello World", JaroWinkler).
			Similarity("Hello World", Jaccard).
			Similarity("Hello World", Cosine).
			Similarity("Hello World", SorensenDice).
			Similarity("Hello World", Lcs).
			Similarity("Hello World", QGram).
			Similarity("Hello World", Hamming).
			GetComparisonManager()

	testMan2 = New("Hello, World!").
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
			Shingle(2).
			Shingle(1).
			ShingleSlice(2).
			LCSBacktrack("Hello World").
			LCSBacktrackAll("Hello World").
			LCSDiff("Hello World").
			Similarity("Hello World", Levenshtein).
			Similarity("Hello World", DamerauLevenshtein).
			Similarity("Hello World", OSADamerauLevenshtein).
			Similarity("Hello World", Jaro).
			Similarity("Hello World", JaroWinkler).
			Similarity("Hello World", QGram).
			Similarity("Hello World", Jaccard).
			Similarity("Hello World", Cosine).
			Similarity("Hello World", SorensenDice).
			Similarity("Hello World", Lcs).
			Similarity("Hello World", Hamming).
			GetComparisonManager()

	testMan3 = New("Hello, World!").
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
			Shingle(2).
			Shingle(1).
			ShingleSlice(2).
			LCSBacktrack("Hello World").
			LCSBacktrackAll("Hello World").
			LCSDiff("Hello World").
			Similarity("Hello World", Levenshtein).
			Similarity("Hello World", DamerauLevenshtein).
			Similarity("Hello World", OSADamerauLevenshtein).
			Similarity("Hello World", Jaro).
			Similarity("Hello World", JaroWinkler).
			Similarity("Hello World", QGram).
			Similarity("Hello World", Jaccard).
			Similarity("Hello World", Cosine).
			Similarity("Hello World", SorensenDice).
			Similarity("Hello World", Lcs).
			Similarity("Hello World", Hamming).
			GetComparisonManager()

	testMan4 = New("Hello, World!").
			WithComparisonManager().
			LevenshteinDistance("Hello World").
			DamerauLevenshteinDistance("Hello World").
			Shingle(2).
			LCSDiff("Hello World").
			Similarity("Hello World", Levenshtein).
			Similarity("Hello World", Cosine).
			GetComparisonManager()

	testMan5 = New("Hello, World!").
			WithComparisonManager().
			GetComparisonManager()
)

func TestNewComparisonManager(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"CM1"},
		{"CM2"},
		{"CM3"},
		{"CM4"},
		{"CM5"},
		{"CM6"},
		{"CM7"},
		{"CM8"},
		{"CM9"},
		{"CM10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := NewComparisonManager()
			if cm.ComparisonResults == nil ||
				cm.SimilarityResults == nil ||
				cm.ShingleResults == nil ||
				cm.LCSResults == nil {
				t.Errorf("NewComparisonManager() = %v, want %v", cm, nil)
			}
		})
	}
}

func TestComparisonManager_AddComparisonResult(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"AddComparisonResult1"},
		{"AddComparisonResult2"},
		{"AddComparisonResult3"},
	}
	sb := NewLoremSentence().WithComparisonManager()
	i := 0
	for i < 100 {
		comp := randomAlphaNumericString(10)
		sb = sb.LevenshteinDistance(comp)
		i++
	}
	for range tests {
		if len(sb.GetComparisonManager().ComparisonResults.GetByType(LevDist)) != 100 {
			t.Errorf("GetComparisonManager.AddComparisonResult() = %v, want %v",
				len(sb.GetComparisonManager().ComparisonResults), 100)
		}
	}
}

func TestCompManGetComparisonResultsMap(t *testing.T) {
	tests := []struct {
		name     string
		cm       *ComparisonManager
		typeCnt  int
		entryCnt int
	}{
		{"CompManGetComparisonResultsMap1", testMan1, 10, 11},
		{"CompManGetComparisonResultsMap2", testMan2, 10, 11},
		{"CompManGetComparisonResultsMap3", testMan3, 10, 11},
		{"CompManGetComparisonResultsMap4", testMan4, 2, 2},
		{"CompManGetComparisonResultsMap5", testMan5, 0, 0},
		{"CompManGetComparisonResultsMap5", &ComparisonManager{}, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.cm.GetComparisonResultsMap().TypeCount() != tt.typeCnt {
				t.Errorf("Expected type count to be %d, got %d",
					tt.typeCnt, tt.cm.GetComparisonResultsMap().TypeCount())
			}
			if tt.cm.GetComparisonResultsMap().EntryCount() != tt.entryCnt {
				t.Errorf("Expected entry count to be %d, got %d",
					tt.entryCnt, tt.cm.GetComparisonResultsMap().EntryCount())
			}
		})
	}
}

func TestCompManCopyComparisonResultsMap(t *testing.T) {
	tests := []struct {
		name     string
		cm       *ComparisonManager
		typeCnt  int
		entryCnt int
	}{
		{"CompManGetComparisonResultsMap1", testMan1, 10, 11},
		{"CompManGetComparisonResultsMap2", testMan2, 10, 11},
		{"CompManGetComparisonResultsMap3", testMan3, 10, 11},
		{"CompManGetComparisonResultsMap4", testMan4, 2, 2},
		{"CompManGetComparisonResultsMap5", testMan5, 0, 0},
		{"CompManGetComparisonResultsMap5", &ComparisonManager{}, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.cm.CopyComparisonResultsMap().TypeCount() != tt.typeCnt {
				t.Errorf("Expected type count to be %d, got %d",
					tt.typeCnt, tt.cm.CopyComparisonResultsMap().TypeCount())
			}
			if tt.cm.CopyComparisonResultsMap().EntryCount() != tt.entryCnt {
				t.Errorf("Expected entry count to be %d, got %d",
					tt.entryCnt, tt.cm.CopyComparisonResultsMap().EntryCount())
			}
			if !tt.cm.CopyComparisonResultsMap().IsMatch(tt.cm.GetComparisonResultsMap()) {
				t.Errorf("Expected match to be %t, got %t",
					true, tt.cm.CopyComparisonResultsMap().IsMatch(tt.cm.GetComparisonResultsMap()))
			}
		})
	}
}

func TestComparisonManager_GetComparisonResult(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		compType ComparisonResultType
		compStr  string
		expected bool
	}{
		{"MapGet1", testMan1, LevDist, "Hello World", true},
		{"MapGet2", testMan1, LevDist, "Hi", true},
		{"MapGet3", testMan1, LevDist, "Yo", false},
		{"MapGet4", testMan5, LevDist, "Hello World", false},
		{"MapGet5", testMan2, JaroWinklerSim, "Hello World", true},
		{"MapGet6", testMan2, LCSLength, "Hello World", true},
		{"MapGet7", testMan3, LevDist, "Yo", true},
		{"MapGet8", testMan4, DamLevDist, "Hello World", true},
		{"MapGet9", testMan5, JaroSim, "Hello World", false},
		{"MapGet10", testMan5, QGramDist, "Hello World", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (tt.compMan.GetComparisonResult(tt.compType, tt.compStr) != nil) != tt.expected {
				t.Errorf("Expected match to be %t, got %t",
					tt.expected, tt.compMan.GetComparisonResult(tt.compType, tt.compStr) != nil)
			}
		})
	}
}

func TestComparisonManagerGetComparisonResultsByType(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		compType ComparisonResultType
		expected int
	}{
		{"MapGetByCompStr1", testMan1, LevDist, 2},
		{"MapGetByCompStr2", testMan2, LevDist, 2},
		{"MapGetByCompStr3", testMan3, LevDist, 2},
		{"MapGetByCompStr4", testMan4, LevDist, 1},
		{"MapGetByCompStr5", testMan5, LevDist, 0},
		{"MapGetByCompStr6", testMan1, HammingDist, 0},
		{"MapGetByCompStr7", testMan2, CosineSim, 1},
		{"MapGetByCompStr8", testMan3, LCSLength, 1},
		{"MapGetByCompStr9", testMan4, DamLevDist, 1},
		{"MapGetByCompStr10", testMan5, SorensenDiceCo, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.compMan.GetComparisonResultsByType(tt.compType)
			if len(res) != tt.expected {
				t.Errorf("Expected match to be %d, got %d", tt.expected, len(res))
			}
		})
	}
}

func TestComparisonManagerFilterComparisonResultsByType(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		compType ComparisonResultType
		expected int
	}{
		{"MapGetByCompStr1", testMan1, LevDist, 2},
		{"MapGetByCompStr2", testMan2, LevDist, 2},
		{"MapGetByCompStr3", testMan3, LevDist, 2},
		{"MapGetByCompStr4", testMan4, LevDist, 1},
		{"MapGetByCompStr5", testMan5, LevDist, 0},
		{"MapGetByCompStr6", testMan1, HammingDist, 0},
		{"MapGetByCompStr7", testMan2, CosineSim, 1},
		{"MapGetByCompStr8", testMan3, LCSLength, 1},
		{"MapGetByCompStr9", testMan4, DamLevDist, 1},
		{"MapGetByCompStr10", testMan5, SorensenDiceCo, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.compMan.FilterComparisonResultsByType(tt.compType)
			if res.EntryCount() != tt.expected {
				t.Errorf("Expected match to be %d, got %d", tt.expected, res.EntryCount())
			}
		})
	}
}

func TestComparisonManager_GetComparisonResultsByString(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		compStr  string
		expected int
	}{
		{"MapGetByCompStr1", testMan1, "Hello World", 10},
		{"MapGetByCompStr2", testMan2, "Hello World", 10},
		{"MapGetByCompStr3", testMan3, "Hello World", 10},
		{"MapGetByCompStr4", testMan4, "Hello World", 2},
		{"MapGetByCompStr5", testMan5, "Hello World", 0},
		{"MapGetByCompStr6", testMan1, "Hi", 1},
		{"MapGetByCompStr7", testMan2, "Hi", 1},
		{"MapGetByCompStr8", testMan3, "Hi", 0},
		{"MapGetByCompStr9", testMan4, "Hi", 0},
		{"MapGetByCompStr10", testMan5, "Hi", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.compMan.GetComparisonResultsByString(tt.compStr)
			if len(res) != tt.expected {
				t.Errorf("Expected match to be %d, got %d", tt.expected, len(res))
			}
		})
	}
}

func TestComparisonManagerFilterComparisonResultsByComparisonString(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		compStr  string
		expected int
	}{
		{"MapGetByCompStr1", testMan1, "Hello World", 10},
		{"MapGetByCompStr2", testMan2, "Hello World", 10},
		{"MapGetByCompStr3", testMan3, "Hello World", 10},
		{"MapGetByCompStr4", testMan4, "Hello World", 2},
		{"MapGetByCompStr5", testMan5, "Hello World", 0},
		{"MapGetByCompStr6", testMan1, "Hi", 1},
		{"MapGetByCompStr7", testMan2, "Hi", 1},
		{"MapGetByCompStr8", testMan3, "Hi", 0},
		{"MapGetByCompStr9", testMan4, "Hi", 0},
		{"MapGetByCompStr10", testMan5, "Hi", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.compMan.FilterComparisonResultsByComparisonString(tt.compStr)
			if res.EntryCount() != tt.expected {
				t.Errorf("Expected match to be %d, got %d", tt.expected, res.EntryCount())
			}
		})
	}
}

func TestComparisonManagerGetSimilarityResultsMap(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		typeCnt  int
		entryCnt int
	}{
		{"MapGet1", testMan1, 11, 11},
		{"MapGet2", testMan2, 11, 11},
		{"MapGet3", testMan3, 11, 11},
		{"MapGet4", testMan4, 2, 2},
		{"MapGet5", testMan5, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.compMan.GetSimilarityResultsMap().TypeCount() != tt.typeCnt {
				t.Errorf("Expected type count to be %d, got %d",
					tt.typeCnt, tt.compMan.GetSimilarityResultsMap().TypeCount())
			}
			if tt.compMan.GetSimilarityResultsMap().EntryCount() != tt.entryCnt {
				t.Errorf("Expected entry count to be %d, got %d",
					tt.entryCnt, tt.compMan.GetSimilarityResultsMap().EntryCount())
			}
		})
	}
}

func TestComparisonManager_CopySimilarityResultsMap(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		typeCnt  int
		entryCnt int
		isMatch  bool
	}{
		{"MapGet1", testMan1, 11, 11, true},
		{"MapGet2", testMan2, 11, 11, true},
		{"MapGet3", testMan3, 11, 11, true},
		{"MapGet4", testMan4, 2, 2, true},
		{"MapGet5", testMan5, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dup := tt.compMan.CopySimilarityResultsMap()
			if dup.TypeCount() != tt.typeCnt {
				t.Errorf("Expected type count to be %d, got %d",
					tt.typeCnt, dup.TypeCount())
			}
			if dup.EntryCount() != tt.entryCnt {
				t.Errorf("Expected entry count to be %d, got %d",
					tt.entryCnt, dup.EntryCount())
			}
			if !dup.IsMatch(tt.compMan.GetSimilarityResultsMap()) {
				t.Errorf("Expected match to be %t, got %t",
					true, dup.IsMatch(tt.compMan.GetSimilarityResultsMap()))
			}
		})
	}
}

func TestComparisonManager_GetSimilarityResult(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		algo     Algorithm
		str      string
		expected bool
	}{
		{"GetTest1", testMan1, Levenshtein, "Hello World", true},
		{"GetTest2", testMan1, DamerauLevenshtein, "Hello World", true},
		{"GetTest3", testMan1, OSADamerauLevenshtein, "Hello World", true},
		{"GetTest4", testMan1, Levenshtein, "Hello World", true},
		{"GetTest5", testMan1, DamerauLevenshtein, "Hello World", true},
		{"GetTest6", testMan5, OSADamerauLevenshtein, "Hello World", false},
		{"GetTest7", testMan5, Levenshtein, "Hello World", false},
		{"GetTest8", testMan5, DamerauLevenshtein, "Hello World", false},
		{"GetTest9", testMan5, OSADamerauLevenshtein, "Hello World", false},
		{"GetTest10", testMan5, Levenshtein, "Hello World", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (tt.compMan.GetSimilarityResult(tt.algo, tt.str) != nil) != tt.expected {
				t.Errorf("Expected match to be %t, got %t",
					tt.expected, tt.compMan.GetSimilarityResult(tt.algo, tt.str) != nil)
			}
		})
	}
}

func TestComparisonManagerGetSimilarityResultsByType(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		algo     Algorithm
		expected int
	}{
		{"GetTest1", testMan1, Levenshtein, 1},
		{"GetTest2", testMan1, DamerauLevenshtein, 1},
		{"GetTest3", testMan1, Cosine, 1},
		{"GetTest4", testMan1, SorensenDice, 1},
		{"GetTest5", testMan1, Hamming, 1},
		{"GetTest6", testMan5, OSADamerauLevenshtein, 0},
		{"GetTest7", testMan5, Levenshtein, 0},
		{"GetTest8", testMan5, DamerauLevenshtein, 0},
		{"GetTest9", testMan5, OSADamerauLevenshtein, 0},
		{"GetTest10", testMan5, Levenshtein, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.compMan.GetSimilarityResultsByType(tt.algo)
			if len(res) != tt.expected {
				t.Errorf("Expected match to be %d, got %d", tt.expected, len(res))
			}
		})
	}
}

func TestComparisonManagerFilterSimilarityResultsByType(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		algo     Algorithm
		expected int
	}{
		{"GetTest1", testMan1, Levenshtein, 1},
		{"GetTest2", testMan1, DamerauLevenshtein, 1},
		{"GetTest3", testMan1, Cosine, 1},
		{"GetTest4", testMan1, SorensenDice, 1},
		{"GetTest5", testMan1, Hamming, 1},
		{"GetTest6", testMan5, OSADamerauLevenshtein, 0},
		{"GetTest7", testMan5, Levenshtein, 0},
		{"GetTest8", testMan5, DamerauLevenshtein, 0},
		{"GetTest9", testMan5, OSADamerauLevenshtein, 0},
		{"GetTest10", testMan5, Levenshtein, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.compMan.FilterSimilarityResultsByType(tt.algo)
			if res.EntryCount() != tt.expected {
				t.Errorf("Expected match to be %d, got %d", tt.expected, res.EntryCount())
			}
		})
	}
}

func TestComparisonManagerGetSimilarityResultsByString(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		str      string
		expected int
	}{
		{"GetTest1", testMan1, "Hello World", 11},
		{"GetTest2", testMan1, "Hi", 0},
		{"GetTest3", testMan1, "Yo", 0},
		{"GetTest4", testMan4, "Hello World", 2},
		{"GetTest5", testMan3, "Hi", 0},
		{"GetTest6", testMan5, "Yo", 0},
		{"GetTest7", testMan5, "Yo", 0},
		{"GetTest8", testMan5, "Yo", 0},
		{"GetTest9", testMan5, "Yo", 0},
		{"GetTest10", testMan5, "Yo", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.compMan.GetSimilarityResultsByComparisonString(tt.str)) != tt.expected {
				t.Errorf("Expected match to be %d, got %d",
					tt.expected, len(tt.compMan.GetSimilarityResultsByComparisonString(tt.str)))
			}
		})
	}

}

func TestComparisonManagerFilterSimilarityResultsByComparisonString(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		str      string
		expected int
	}{
		{"GetTest1", testMan1, "Hello World", 11},
		{"GetTest2", testMan1, "Hi", 0},
		{"GetTest3", testMan1, "Yo", 0},
		{"GetTest4", testMan4, "Hello World", 2},
		{"GetTest5", testMan3, "Hi", 0},
		{"GetTest6", testMan5, "Yo", 0},
		{"GetTest7", testMan5, "Yo", 0},
		{"GetTest8", testMan5, "Yo", 0},
		{"GetTest9", testMan5, "Yo", 0},
		{"GetTest10", testMan5, "Yo", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.compMan.FilterSimilarityResultsByComparisonString(tt.str).EntryCount() != tt.expected {
				t.Errorf("Expected match to be %d, got %d",
					tt.expected, tt.compMan.FilterSimilarityResultsByComparisonString(tt.str).EntryCount())
			}
		})
	}
}

func TestComparisonManager_GetShingleResultsMap(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		typeCnt  int
		entryCnt int
	}{
		{"MapGet1", testMan1, 2, 3},
		{"MapGet2", testMan2, 2, 3},
		{"MapGet3", testMan3, 2, 3},
		{"MapGet4", testMan4, 1, 1},
		{"MapGet5", testMan5, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.compMan.GetShingleResultsMap().TypeCount() != tt.typeCnt {
				t.Errorf("Expected type count to be %d, got %d",
					tt.typeCnt, tt.compMan.GetShingleResultsMap().TypeCount())
			}
			if tt.compMan.GetShingleResultsMap().EntryCount() != tt.entryCnt {
				t.Errorf("Expected entry count to be %d, got %d",
					tt.entryCnt, tt.compMan.GetShingleResultsMap().EntryCount())
			}
		})
	}
}

func TestComparisonManager_CopyShingleResultsMap(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		typeCnt  int
		entryCnt int
		isMatch  bool
	}{
		{"MapGet1", testMan1, 2, 3, true},
		{"MapGet2", testMan2, 2, 3, true},
		{"MapGet3", testMan3, 2, 3, true},
		{"MapGet4", testMan4, 1, 1, true},
		{"MapGet5", testMan5, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dup := tt.compMan.CopyShingleResultsMap()
			if dup.TypeCount() != tt.typeCnt {
				t.Errorf("Expected type count to be %d, got %d",
					tt.typeCnt, dup.TypeCount())
			}
		})
	}
}

func TestComparisonManager_GetShingleResult(t *testing.T) {
	tests := []struct {
		name        string
		compMan     *ComparisonManager
		shingleType ShingleResultType
		ngramLength int
		expected    bool
	}{
		{"GetShingle1", testMan1, ShinglesMap, 2, true},
		{"GetShingle2", testMan1, ShinglesMap, 1, true},
		{"GetShingle3", testMan1, ShinglesMap, 4, false},
		{"GetShingle4", testMan1, ShinglesSlice, 2, true},
		{"GetShingle5", testMan5, ShinglesMap, 6, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (tt.compMan.GetShingleResult(tt.shingleType, tt.ngramLength) != nil) != tt.expected {
				t.Errorf("Expected match to be %t, got %t",
					tt.expected, tt.compMan.GetShingleResult(tt.shingleType, tt.ngramLength) != nil)
			}
		})
	}
}

func TestComparisonManagerGetShingleResultsByType(t *testing.T) {
	tests := []struct {
		name        string
		compMan     *ComparisonManager
		shingleType ShingleResultType
		expected    int
	}{
		{"ShingleTest1", testMan1, ShinglesMap, 2},
		{"ShingleTest2", testMan1, ShinglesSlice, 1},
		{"ShingleTest3", testMan2, ShinglesMap, 2},
		{"ShingleTest4", testMan2, ShinglesSlice, 1},
		{"ShingleTest5", testMan4, ShinglesMap, 1},
		{"ShingleTest6", testMan5, ShinglesSlice, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.compMan.GetShingleResultsByType(tt.shingleType)) != tt.expected {
				t.Errorf("Expected match to be %d, got %d",
					tt.expected, len(tt.compMan.GetShingleResultsByType(tt.shingleType)))
			}
		})
	}
}

func TestComparisonManagerFilterShingleResultsByType(t *testing.T) {
	tests := []struct {
		name        string
		compMan     *ComparisonManager
		shingleType ShingleResultType
		expected    int
	}{
		{"ShingleTest1", testMan1, ShinglesMap, 2},
		{"ShingleTest2", testMan1, ShinglesSlice, 1},
		{"ShingleTest3", testMan2, ShinglesMap, 2},
		{"ShingleTest4", testMan2, ShinglesSlice, 1},
		{"ShingleTest5", testMan4, ShinglesMap, 1},
		{"ShingleTest6", testMan5, ShinglesSlice, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.compMan.FilterShingleResultsByType(tt.shingleType).EntryCount() != tt.expected {
				t.Errorf("Expected match to be %d, got %d",
					tt.expected, tt.compMan.FilterShingleResultsByType(tt.shingleType).EntryCount())
			}
		})
	}
}

func TestComparisonManager_GetShingleResultsByNGramLength(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		ngram    int
		expected int
	}{
		{"ShingleTest1", testMan1, 2, 2},
		{"ShingleTest2", testMan1, 1, 1},
		{"ShingleTest3", testMan2, 2, 2},
		{"ShingleTest4", testMan2, 1, 1},
		{"ShingleTest5", testMan4, 2, 1},
		{"ShingleTest6", testMan5, 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.compMan.GetShingleResultsByNGramLength(tt.ngram)) != tt.expected {
				t.Errorf("Expected match to be %d, got %d",
					tt.expected, len(tt.compMan.GetShingleResultsByNGramLength(tt.ngram)))
			}
		})
	}
}

func TestComparisonManager_FilterShingleResultsByNGramLength(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		ngram    int
		expected int
	}{
		{"ShingleTest1", testMan1, 2, 2},
		{"ShingleTest2", testMan1, 1, 1},
		{"ShingleTest3", testMan2, 2, 2},
		{"ShingleTest4", testMan2, 1, 1},
		{"ShingleTest5", testMan4, 2, 1},
		{"ShingleTest6", testMan5, 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.compMan.FilterShingleResultsByNGramLength(tt.ngram).EntryCount() != tt.expected {
				t.Errorf("Expected match to be %d, got %d",
					tt.expected, tt.compMan.FilterShingleResultsByNGramLength(tt.ngram).EntryCount())
			}
		})
	}
}

func TestComparisonManager_GetLCSResultsMap(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		typeCnt  int
		entryCnt int
	}{
		{"MapGet1", testMan1, 3, 3},
		{"MapGet2", testMan2, 3, 3},
		{"MapGet3", testMan3, 3, 3},
		{"MapGet4", testMan4, 1, 1},
		{"MapGet5", testMan5, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.compMan.GetLCSResultsMap().TypeCount() != tt.typeCnt {
				t.Errorf("Expected type count to be %d, got %d",
					tt.typeCnt, tt.compMan.GetLCSResultsMap().TypeCount())
			}
			if tt.compMan.GetLCSResultsMap().EntryCount() != tt.entryCnt {
				t.Errorf("Expected entry count to be %d, got %d",
					tt.entryCnt, tt.compMan.GetLCSResultsMap().EntryCount())
			}
		})
	}
}

func TestComparisonManager_CopyLCSResultsMap(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		typeCnt  int
		entryCnt int
		isMatch  bool
	}{
		{"MapGet1", testMan1, 3, 3, true},
		{"MapGet2", testMan2, 3, 3, true},
		{"MapGet3", testMan3, 3, 3, true},
		{"MapGet4", testMan4, 1, 1, true},
		{"MapGet5", testMan5, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dup := tt.compMan.CopyLCSResultsMap()
			if dup.TypeCount() != tt.typeCnt {
				t.Errorf("Expected type count to be %d, got %d",
					tt.typeCnt, dup.TypeCount())
			}
		})
	}
}

func TestComparisonManager_GetLCSResult(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		lcsType  LCSResultType
		compStr  string
		expected bool
	}{
		{"LCSGet1", testMan1, LCSBacktrackWord, "Hello World", true},
		{"LCSGet2", testMan1, LCSBacktrackWordAll, "Hello World", true},
		{"LCSGet3", testMan1, LCSDiffSlice, "Hello World", true},
		{"LCSGet4", testMan4, LCSDiffSlice, "Hello World", true},
		{"LCSGet5", testMan5, LCSBacktrackWord, "Yo", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (tt.compMan.GetLCSResult(tt.lcsType, tt.compStr) != nil) != tt.expected {
				t.Errorf("Expected match to be %t, got %t",
					tt.expected, tt.compMan.GetLCSResult(tt.lcsType, tt.compStr) != nil)
			}
		})
	}
}

func TestComparisonManagerGetLCSResultsByType(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		lcsType  LCSResultType
		expected int
	}{
		{"LCSGet1", testMan1, LCSBacktrackWord, 1},
		{"LCSGet2", testMan1, LCSBacktrackWordAll, 1},
		{"LCSGet3", testMan1, LCSDiffSlice, 1},
		{"LCSGet4", testMan4, LCSDiffSlice, 1},
		{"LCSGet5", testMan5, LCSBacktrackWord, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.compMan.GetLCSResultsByType(tt.lcsType)) != tt.expected {
				t.Errorf("Expected match to be %d, got %d",
					tt.expected, len(tt.compMan.GetLCSResultsByType(tt.lcsType)))
			}
		})
	}
}

func TestComparisonManagerFilterLCSResultsByType(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		lcsType  LCSResultType
		expected int
	}{
		{"LCSGet1", testMan1, LCSBacktrackWord, 1},
		{"LCSGet2", testMan1, LCSBacktrackWordAll, 1},
		{"LCSGet3", testMan1, LCSDiffSlice, 1},
		{"LCSGet4", testMan4, LCSDiffSlice, 1},
		{"LCSGet5", testMan5, LCSBacktrackWord, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.compMan.FilterLCSResultsByType(tt.lcsType).EntryCount() != tt.expected {
				t.Errorf("Expected match to be %d, got %d",
					tt.expected, tt.compMan.FilterLCSResultsByType(tt.lcsType).EntryCount())
			}
		})
	}
}

func TestComparisonManager_GetLCSResultsByComparisonString(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		str      string
		expected int
	}{
		{"LCSGet1", testMan1, "Hello World", 3},
		{"LCSGet2", testMan1, "Hi", 0},
		{"LCSGet3", testMan1, "Yo", 0},
		{"LCSGet4", testMan4, "Hello World", 1},
		{"LCSGet5", testMan5, "Hi", 0},
		{"LCSGet6", testMan5, "Yo", 0},
		{"LCSGet7", testMan2, "Hello World", 3},
		{"LCSGet8", testMan2, "Hello World", 3},
		{"LCSGet9", testMan3, "Hello World", 3},
		{"LCSGet10", testMan3, "Hello World", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.compMan.GetLCSResultsByComparisonString(tt.str)) != tt.expected {
				t.Errorf("Expected match to be %d, got %d",
					tt.expected, len(tt.compMan.GetLCSResultsByComparisonString(tt.str)))
			}
		})
	}
}

func TestComparisonManager_FilterLCSResultsByComparisonString(t *testing.T) {
	tests := []struct {
		name     string
		compMan  *ComparisonManager
		str      string
		expected int
	}{
		{"LCSGet1", testMan1, "Hello World", 3},
		{"LCSGet2", testMan1, "Hi", 0},
		{"LCSGet3", testMan1, "Yo", 0},
		{"LCSGet4", testMan4, "Hello World", 1},
		{"LCSGet5", testMan5, "Hi", 0},
		{"LCSGet6", testMan5, "Yo", 0},
		{"LCSGet7", testMan2, "Hello World", 3},
		{"LCSGet8", testMan2, "Hello World", 3},
		{"LCSGet9", testMan3, "Hello World", 3},
		{"LCSGet10", testMan3, "Hello World", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.compMan.FilterLCSResultsByComparisonString(tt.str).EntryCount() != tt.expected {
				t.Errorf("Expected match to be %d, got %d",
					tt.expected, tt.compMan.FilterLCSResultsByComparisonString(tt.str).EntryCount())
			}
		})
	}
}

func TestEmptyTestGet(t *testing.T) {
	testManager := ComparisonManager{}
	//map gets
	if testManager.GetShingleResultsMap() != nil {
		t.Errorf("Expected map to be nil")
	}
	if testManager.GetLCSResultsMap() != nil {
		t.Errorf("Expected map to be nil")
	}
	if testManager.GetComparisonResultsMap() != nil {
		t.Errorf("Expected map to be nil")
	}
	if testManager.GetSimilarityResultsMap() != nil {
		t.Errorf("Expected map to be nil")
	}
	//map copies
	if testManager.CopyShingleResultsMap() != nil {
		t.Errorf("Expected map to be nil")
	}
	if testManager.CopyLCSResultsMap() != nil {
		t.Errorf("Expected map to be nil")
	}
	if testManager.CopyComparisonResultsMap() != nil {
		t.Errorf("Expected map to be nil")
	}
	if testManager.CopySimilarityResultsMap() != nil {
		t.Errorf("Expected map to be nil")
	}
	//simple gets
	if testManager.GetShingleResult(ShinglesMap, 4) != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.GetLCSResult(LCSBacktrackWord, "Hi") != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.GetComparisonResult(LevDist, "Hi") != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.GetSimilarityResult(Cosine, "Hi") != nil {
		t.Errorf("Expected result to be nil")
	}
	//get by type
	if testManager.GetShingleResultsByType(ShinglesMap) != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.GetLCSResultsByType(LCSBacktrackWord) != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.GetComparisonResultsByType(LevDist) != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.GetSimilarityResultsByType(Levenshtein) != nil {
		t.Errorf("Expected result to be nil")
	}
	//filter by type
	if testManager.FilterShingleResultsByType(ShinglesMap) != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.FilterLCSResultsByType(LCSBacktrackWord) != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.FilterComparisonResultsByType(LevDist) != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.FilterSimilarityResultsByType(Levenshtein) != nil {
		t.Errorf("Expected result to be nil")
	}
	// get by other key
	if testManager.GetShingleResultsByNGramLength(2) != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.GetComparisonResultsByString("Hi") != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.GetSimilarityResultsByComparisonString("Hi") != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.GetLCSResultsByComparisonString("Hi") != nil {
		t.Errorf("Expected result to be nil")
	}
	//filter by other key
	if testManager.FilterShingleResultsByNGramLength(2) != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.FilterComparisonResultsByComparisonString("Hi") != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.FilterSimilarityResultsByComparisonString("Hi") != nil {
		t.Errorf("Expected result to be nil")
	}
	if testManager.FilterLCSResultsByComparisonString("Hi") != nil {
		t.Errorf("Expected result to be nil")
	}
}

func TestEmptyAdds(t *testing.T) {
	testManager := ComparisonManager{}
	l := LCSBacktrack("Hi", "Hi")
	s := Shingle("Hi", 1)
	c := LevenshteinDistance("Hi", "Hi")
	si := Similarity("Hi", "Hi", Cosine)
	testManager.AddLCSResult(*l)
	testManager.AddShingleResult(s)
	testManager.AddComparisonResult(c)
	testManager.AddSimilarityResult(*si)

	if testManager.GetLCSResult(LCSBacktrackWord, "Hi") == nil {
		t.Errorf("Expected result to not be nil")
	}
	if testManager.GetShingleResult(ShinglesMap, 1) == nil {
		t.Errorf("Expected result to not be nil")
	}
	if testManager.GetComparisonResult(LevDist, "Hi") == nil {
		t.Errorf("Expected result to not be nil")
	}
	if testManager.GetSimilarityResult(Cosine, "Hi") == nil {
		t.Errorf("Expected result to not be nil")
	}
}
