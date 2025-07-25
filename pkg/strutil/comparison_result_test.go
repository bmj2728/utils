package strutil

import (
	"testing"
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

func TestComparisonResultPrint(t *testing.T) {
	tests := []struct {
		name    string
		resType ComparisonResultType
	}{
		{"PrintLevDist", LevDist},
		{"PrintDamLevDist", DamLevDist},
		{"PrintOSADamLevDist", OSADamLevDist},
		{"PrintLCS", LCSLength},
		{"PrintLCSDist", LCSDist},
	}
	correct := "With great power there must also come great responsibility."
	common := "With great power comes great responsibility."

	ls := New(correct).
		WithComparisonManager().
		LevenshteinDistance(common).
		DamerauLevenshteinDistance(common).
		OSADamerauLevenshteinDistance(common).
		LCS(common).
		LCSEditDistance(common).
		QgramDistance(common, 3).
		JaroSimilarity(common).
		JaroWinklerSimilarity(common).
		JaccardSimilarity(common, 3).
		CosineSimilarity(common, 3).
		QgramSimilarity(common, 3).
		SorensenDiceCoefficient(common, 3).
		HammingDistance(correct).
		ComparisonManager().
		ComparisonResultsMap()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls.GetByType(tt.resType)[0].Print(true)
			ls.GetByType(tt.resType)[0].Print(false)
		})
	}
}
