package strutil

import (
	"errors"
	"math"
	"testing"

	errors2 "utils/pkg/internal/errors"
	"utils/pkg/internal/types"
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
		{"NotReal", Algorithm(100), ""},
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
		{"Fake", "Hello World", "Hello World", Algorithm(99), 0.0, errors.New("Illegal argument for algorithm method")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := *New(tt.input).
				WithComparisonManager().
				Similarity(tt.compare, tt.algo).
				GetComparisonManager().
				GetSimilarityResultsMap()[tt.algo][tt.compare]
			s1, s2 := res.GetStrings()
			score, err := res.GetScore()
			scoreComp := math.Abs(float64(score) - float64(tt.score))
			if err != nil && tt.algo != Algorithm(99) {
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
			if scoreComp > types.Float64EqualityThreshold {
				t.Errorf("SimilarityResultGetters() = %f, want %f - %f", score, tt.score,
					math.Abs(float64(score-tt.score)))
			}
			if !errors2.CompareErrors(res.GetError(), tt.err) {
				t.Errorf("SimilarityResultGetters() = %v, want %v", res, tt.err)
			}
		})
	}
}

var (
	fullSimilarityResults1 = New("Hello, World!").
				WithComparisonManager().
				Similarity("Hello World", Levenshtein).
				Similarity("Hello World", DamerauLevenshtein).
				Similarity("Hello World", JaroWinkler).
				Similarity("Hello World", Hamming).
				Similarity("Hello World", Jaccard).
				Similarity("Hello World", Jaro).
				Similarity("Hello World", JaroWinkler).
				Similarity("Hello World", Cosine).
				Similarity("Hello World", SorensenDice).
				Similarity("Hello World", Lcs).
				Similarity("Hello World", QGram).
				GetComparisonManager().
				GetSimilarityResultsMap()

	fullSimilarityResults2 = New("Hi, World!").
				WithComparisonManager().
				Similarity("Hello World", Levenshtein).
				Similarity("Hello World", DamerauLevenshtein).
				Similarity("Hello World", JaroWinkler).
				Similarity("Hello World", Jaccard).
				Similarity("Hello World", Jaro).
				Similarity("Hello World", JaroWinkler).
				Similarity("Hello World", Cosine).
				Similarity("Hello World", SorensenDice).
				Similarity("Hello World", Lcs).
				Similarity("Hello World", QGram).
				Similarity("Hello World", Hamming).
				GetComparisonManager().
				GetSimilarityResultsMap()
)

func TestSimilarityIsMatch(t *testing.T) {
	tests := []struct {
		name    string
		simRes1 SimilarityResultsMap
		algo1   Algorithm
		simRes2 SimilarityResultsMap
		algo2   Algorithm
		want    bool
	}{
		{"IsMatchSame", fullSimilarityResults1, Levenshtein, fullSimilarityResults1, Levenshtein, true},
		{"IsMatchDifferent", fullSimilarityResults1, Levenshtein, fullSimilarityResults2, Levenshtein, false},
		{"IsMatchDifferentAlgo", fullSimilarityResults1, Levenshtein, fullSimilarityResults1, DamerauLevenshtein, false},
		{"IsMatchDifferentAlgoDiff", fullSimilarityResults1, Levenshtein, fullSimilarityResults2, Lcs, false},
	}
	key := "Hello World"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry1 := tt.simRes1[tt.algo1][key]
			entry2 := tt.simRes2[tt.algo2][key]
			if entry1.IsMatch(entry2) != tt.want {
				t.Errorf("SimilarityIsMatch() = %v, want %v", entry1.IsMatch(entry2), tt.want)
			}
		})
	}
}

var (
	levShort      = "Levenshtein: 0.846154\n"
	levLong       = "Comparison: Levenshtein\nFirst String: Hello, World!\nSecond String: Hello World\nScore: 0.846154\n"
	hamErrorShort = "Hamming GetError: Undefined for strings of unequal length\n"
	hamLongError  = "GetError during processing Hamming\n" +
		"First String: Hello, World!\n" +
		"Second String: Hello World\n" +
		"GetError: Undefined for strings of unequal length\n"
	prntInput   = "Hello, World!"
	prntCompare = "Hello World"
)

func TestSimilarityResultFormatOutput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		compare  string
		algo     Algorithm
		verbose  bool
		expected string
	}{
		{"LevenshteinShort", prntInput, prntCompare, Levenshtein, false, levShort},
		{"LevenshteinLong", prntInput, prntCompare, Levenshtein, true, levLong},
		{"HammingErrorShort", prntInput, prntCompare, Hamming, false, hamErrorShort},
		{"HammingLongError", prntInput, prntCompare, Hamming, true, hamLongError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := *New(tt.input).
				WithComparisonManager().
				Similarity(tt.compare, tt.algo).
				GetComparisonManager().
				GetSimilarityResultsMap()[tt.algo][tt.compare]
			if formatSimilarityResultOutput(&res, tt.verbose) != tt.expected {
				t.Errorf("SimilarityResultFormatOutput() = %s, want %s",
					formatSimilarityResultOutput(&res, tt.verbose), tt.expected)
			}
			res.Print(tt.verbose)
		})
	}
}
