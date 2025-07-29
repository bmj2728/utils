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
			ComparisonManager()

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
			ComparisonManager()

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
			ComparisonManager()

	testMan4 = New("Hello, World!").
			WithComparisonManager().
			LevenshteinDistance("Hello World").
			DamerauLevenshteinDistance("Hello World").
			ComparisonManager()

	testMan5 = New("Hello, World!").
			WithComparisonManager().
			ComparisonManager()
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
				cm.ShingleData == nil ||
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
		if len(sb.ComparisonManager().ComparisonResults.GetByType(LevDist)) != 100 {
			t.Errorf("ComparisonManager.AddComparisonResult() = %v, want %v", len(sb.ComparisonManager().ComparisonResults), 100)
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
