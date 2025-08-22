package strutil

import (
	"testing"

	"github.com/bmj2728/utils/pkg/internal/comparison"
)

func TestFuzzyMatch(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		compList []string
		algo     Algorithm
		expected string
	}{
		{"FuzzyMatch1",
			"Katharine",
			[]string{"Catherine", "Kathrine", "Katherine", "Kathryn", "Ellen"},
			JaroWinkler,
			"Kathrine"},
		{"FuzzyMatch2",
			"Katherine",
			[]string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"},
			JaroWinkler,
			"Katharine"},
		{"FuzzyMatch3",
			"Katherine",
			[]string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"},
			Cosine,
			"Catherine"},
		{"FuzzyMatch4",
			"Katherine",
			[]string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"},
			Cosine,
			"Catherine"},
		{"FuzzyMatch5",
			"Katherine",
			[]string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"},
			Hamming,
			"Catherine"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult, err := fuzzyMatch(tt.input, tt.compList, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if helperResult != tt.expected {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, helperResult)
			}
			result, err := FuzzyMatch(tt.input, tt.compList, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if result != tt.expected {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, result)
			}
			builderResult, err := New(tt.input).FuzzyMatch(tt.compList, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if builderResult != tt.expected {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, builderResult)
			}
		})
	}
}

func TestFuzzyMatchWithThreshold(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		compList  []string
		threshold float32
		algo      Algorithm
		expected  string
	}{
		{"FuzzyMatch1",
			"Katharine",
			[]string{"Catherine", "Kathrine", "Katherine", "Kathryn", "Ellen"},
			0.9,
			JaroWinkler,
			"Kathrine"},
		{"FuzzyMatch2",
			"Katherine",
			[]string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"},
			0.8,
			JaroWinkler,
			"Katharine"},
		{"FuzzyMatch3",
			"Katherine",
			[]string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"},
			0.9,
			Cosine,
			""},
		{"FuzzyMatch4",
			"Katherine",
			[]string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"},
			0.8,
			Cosine,
			"Catherine"},
		{"FuzzyMatch5",
			"Katherine",
			[]string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"},
			0.8,
			Hamming,
			"Catherine"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult, err := fuzzyMatchWithThreshold(tt.input, tt.compList, tt.threshold, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if helperResult != tt.expected {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, helperResult)
			}
			result, err := FuzzyMatchWithThreshold(tt.input, tt.compList, tt.threshold, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if result != tt.expected {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, result)
			}
			builderResult, err := New(tt.input).FuzzyMatchWithThreshold(tt.compList, tt.threshold, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if builderResult != tt.expected {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, builderResult)
			}
		})
	}
}

var (
	input            = "Katherine"
	fuzzSet          = []string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"}
	lcsExpected      = []string{"Kathrine", "Catherine", "Katharine", "Kathryn", ""}
	jwExpected       = []string{"Katharine", "Kathrine", "Catherine", "Kathryn", "Ellen"}
	cosExpected      = []string{"Catherine", "Kathryn", "Katharine", "Kathrine", ""}
	qgramExpected    = []string{"Catherine", "Kathrine", "Katharine", "Kathryn", ""}
	jaccardExpected  = []string{"Catherine", "Kathrine", "Katharine", "Kathryn", ""}
	levExpected      = []string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"}
	sorensenExpected = []string{"Catherine", "Kathrine", "Katharine", "Kathryn", ""}
	damerauExpected  = []string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"}
	osadExpected     = []string{"Catherine", "Kathrine", "Katharine", "Kathryn", "Ellen"}
	hammingExpected  = []string{"Catherine", "Katharine", "", "", ""}
)

func TestFuzzyMatchSet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		compList []string
		algo     Algorithm
		expected []string
	}{
		{"FuzzMatchSet1", input, fuzzSet, JaroWinkler, jwExpected},
		{"FuzzMatchSet2", input, fuzzSet, Cosine, cosExpected},
		{"FuzzMatchSet3", input, fuzzSet, Levenshtein, levExpected},
		{"FuzzMatchSet4", input, fuzzSet, SorensenDice, sorensenExpected},
		{"FuzzMatchSet5", input, fuzzSet, DamerauLevenshtein, damerauExpected},
		{"FuzzMatchSet6", input, fuzzSet, OSADamerauLevenshtein, osadExpected},
		{"FuzzMatchSet7", input, fuzzSet, QGram, qgramExpected},
		{"FuzzMatchSet8", input, fuzzSet, Jaccard, jaccardExpected},
		{"FuzzMatchSet9", input, fuzzSet, Lcs, lcsExpected},
		{"FuzzMatchSet10", input, fuzzSet, Hamming, hammingExpected},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult, err := fuzzyMatchSet(tt.input, tt.compList, 5, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if !comparison.CompareSlices(helperResult, tt.expected, false) {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, helperResult)
			}
			result, err := FuzzyMatchSet(tt.input, tt.compList, 5, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if !comparison.CompareSlices(result, tt.expected, false) {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, result)
			}
			builderResult, err := New(tt.input).FuzzyMatchSet(tt.compList, 5, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if !comparison.CompareSlices(builderResult, tt.expected, false) {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, builderResult)
			}
		})
	}
}

var (
	lcsExpectedThresh      = []string{"Kathrine", "", "", "", ""}
	jwExpectedThresh       = []string{"Katharine", "Kathrine", "Catherine", "Kathryn", ""}
	cosExpectedThresh      = []string{"Catherine", "Kathrine", "", "", ""}
	qgramExpectedThresh    = []string{"Catherine", "Kathrine", "", "", ""}
	jaccardExpectedThresh  = []string{"", "", "", "", ""}
	levExpectedThresh      = []string{"Catherine", "Kathrine", "Katharine", "", ""}
	sorensenExpectedThresh = []string{"Catherine", "Kathrine", "", "", ""}
	damerauExpectedThresh  = []string{"Catherine", "Kathrine", "Katharine", "", ""}
	osadExpectedThresh     = []string{"Catherine", "Kathrine", "Katharine", "", ""}
	hammingExpectedThresh  = []string{"Catherine", "Katharine", "", "", ""}
)

func TestFuzzyMatchSetWithThreshold(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		compList  []string
		threshold float32
		algo      Algorithm
		expected  []string
	}{
		{"FuzzMatchSet1", input, fuzzSet, 0.8, JaroWinkler, jwExpectedThresh},
		{"FuzzMatchSet2", input, fuzzSet, 0.8, Cosine, cosExpectedThresh},
		{"FuzzMatchSet3", input, fuzzSet, 0.8, Levenshtein, levExpectedThresh},
		{"FuzzMatchSet4", input, fuzzSet, 0.8, SorensenDice, sorensenExpectedThresh},
		{"FuzzMatchSet5", input, fuzzSet, 0.8, DamerauLevenshtein, damerauExpectedThresh},
		{"FuzzMatchSet6", input, fuzzSet, 0.8, OSADamerauLevenshtein, osadExpectedThresh},
		{"FuzzMatchSet7", input, fuzzSet, 0.8, QGram, qgramExpectedThresh},
		{"FuzzMatchSet8", input, fuzzSet, 0.8, Jaccard, jaccardExpectedThresh},
		{"FuzzMatchSet9", input, fuzzSet, 0.8, Lcs, lcsExpectedThresh},
		{"FuzzMatchSet10", input, fuzzSet, 0.8, Hamming, hammingExpectedThresh},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult, err := fuzzyMatchSetWithThreshold(tt.input, tt.compList, 5, tt.threshold, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if !comparison.CompareSlices(helperResult, tt.expected, false) {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, helperResult)
			}
			result, err := FuzzyMatchSetWithThreshold(tt.input, tt.compList, 5, tt.threshold, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if !comparison.CompareSlices(result, tt.expected, false) {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, result)
			}
			builderResult, err := New(tt.input).FuzzyMatchSetWithThreshold(tt.compList, 5, tt.threshold, tt.algo)
			if err != nil {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, err)
			}
			if !comparison.CompareSlices(builderResult, tt.expected, false) {
				t.Errorf("FuzzyMatch - expected %q - got %q", tt.expected, builderResult)
			}
		})
	}
}
