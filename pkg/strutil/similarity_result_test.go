package strutil

import (
	"math"
	"testing"
)

func TestAlgorithmString(t *testing.T) {
	tests := []struct {
		name string
		algo Algorithm
		want string
	}{
		{"Levenshtein", Levenshtein, "Levenshtein"},
		{"DamerauLevenshtein", DamerauLevenshtein, "Damerau-Levenshtein"},
		{"JaroWinkler", JaroWinkler, "Jaro-Winkler"},
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

func TestSimilarityResultGetters(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		compare string
		algo    Algorithm
		score   float32
		err     error
	}{
		{"Levenshtein", "Hello World", "Hello World", Levenshtein, 1.0, nil},
		{"DamerauLevenshtein", "Hello World", "Hello World", DamerauLevenshtein, 1.0, nil},
		{"JaroWinkler", "Hello World", "Hello World", JaroWinkler, 1.0, nil},
		{"Hamming", "Hello World", "Hello World", Hamming, 1.0, nil},
		{"Jaccard", "Hello World", "Hello World", Jaccard, 1.0, nil},
		{"Levenshtein", "Hello World", "Hello", Levenshtein, 0.454545, nil},
		{"DamerauLevenshtein", "Hello World", "Hello", DamerauLevenshtein, 0.454545, nil},
		{"JaroWinkler", "Hello World", "Hello", JaroWinkler, 0.890909, nil},
		{"Hamming", "Hello World", "Hello World", Hamming, 1.0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := *New(tt.input).
				WithComparisonManager().
				Similarity(tt.compare, tt.algo).
				ComparisonManager().
				GetSimilarityResults()[tt.algo][tt.compare]
			s1, s2 := res.GetStrings()
			score, err := res.GetScore()
			scoreComp := math.Abs(float64(score) - float64(tt.score))
			if err != nil {
				t.Errorf("SimilarityResultGetters() = %v, want %v", err, nil)
			}
			if res.GetAlgorithm() != tt.algo {
				t.Errorf("SimilarityResultGetters() = %v, want %v", res.GetAlgorithm(), tt.algo)
			}
			if res.GetAlgorithmName() != tt.algo.String() {
				t.Errorf("SimilarityResultGetters() = %s, want %s", res.GetAlgorithmName(), tt.algo.String())
			}
			if res.GetString1() != tt.input {
				t.Errorf("SimilarityResultGetters() = %s, want %s", res.GetString1(), tt.input)
			}
			if res.GetString2() != tt.compare {
				t.Errorf("SimilarityResultGetters() = %s, want %s", res.GetString2(), tt.compare)
			}
			if s1 != tt.input || s2 != tt.compare {
				t.Errorf("SimilarityResultGetters() = %s, %s, want %s", s1, s2, tt.input)
			}
			if scoreComp > float64EqualityThreshold {
				t.Errorf("SimilarityResultGetters() = %f, want %f - %f", score, tt.score,
					math.Abs(float64(score-tt.score)))
			}
			if !compareErrors(res.GetError(), tt.err) {
				t.Errorf("SimilarityResultGetters() = %v, want %v", res, tt.err)
			}
		})
	}
}
