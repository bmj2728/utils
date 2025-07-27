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
				cm.LCSData == nil {
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
