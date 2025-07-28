package strutil

import "testing"

func TestAlgorithmString(t *testing.T) {
	tests := []struct {
		name string
		algo Algorithm
		want string
	}{
		{"Levenshtein", Levenshtein, "Levenshtein"},
		{"DamerauLevenshtein", DamerauLevenshtein, "DamerauLevenshtein"},
		{"JaroWinkler", JaroWinkler, "JaroWinkler"},
		{"Hamming", Hamming, "Hamming"},
		{"Jaccard", Jaccard, "Jaccard"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.algo.String(); got != tt.want {
				t.Errorf("AlgorithmString() = %v, want %v", got, tt.want)
			}
		})
	}
}
