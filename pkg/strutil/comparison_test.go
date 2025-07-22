package strutil

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

// TODO update test cases to address result objects

func TestLevenshteinDistance(t *testing.T) {
	test := []struct {
		name     string
		input1   string
		input2   string
		expected int
	}{
		{"LevenshteinDistance1", "hello", "hello", 0},
		{"LevenshteinDistance2", "hello", "helloo", 1},
		{"LevenshteinDistance3", "ABCDEFG", "abcdefg", 7},
		{"LevenshteinDistance4", "ABCDEFG", "ABCDEFGH", 1},
		{"LevenshteinDistance5", "hello", "world", 4},
		{"LevenshteinDistance6", "My name is John", "My name is Jane", 3},
		{"LevenshteinDistance7", "tpyo", "typo", 2},
		{"LevenshteinDistance8", "teal", "tale", 2},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := levenshteinDistance(tt.input1, tt.input2)
			result := LevenshteinDistance(tt.input1, tt.input2)
			builderResult := New(tt.input1).LevenshteinDistance(tt.input2)
			if helperResult != tt.expected ||
				result != tt.expected ||
				*builderResult.ComparisonData().GetLevenshteinDist() != tt.expected {
				t.Errorf("LevenshteinDistance - expected %d - got %d / %d / %d",
					tt.expected,
					helperResult,
					result,
					*builderResult.comparisonData.LevenshteinDist)
			}
		})
	}
}

func TestDamerauLevenshteinDistance(t *testing.T) {
	test := []struct {
		name     string
		input1   string
		input2   string
		expected int
	}{
		{"DamerauLevenshteinDistanceDistance1", "hello", "hello", 0},
		{"DamerauLevenshteinDistanceDistance2", "hello", "helloo", 1},
		{"DamerauLevenshteinDistance3", "ABCDEFG", "abcdefg", 7},
		{"DamerauLevenshteinDistance4", "ABCDEFG", "ABCDEFGH", 1},
		{"DamerauLevenshteinDistance5", "hello", "world", 4},
		{"DamerauLevenshteinDistance6", "My name is John", "My name is Jane", 3},
		{"DamerauLevenshteinDistance7", "tpyo", "typo", 1},
		{"DamerauLevenshteinDistance8", "teal", "tale", 2},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := damerauLevenshteinDistance(tt.input1, tt.input2)
			result := DamerauLevenshteinDistance(tt.input1, tt.input2)
			builderResult := New(tt.input1).DamerauLevenshteinDistance(tt.input2)
			if helperResult != tt.expected ||
				result != tt.expected ||
				*builderResult.comparisonData.GetDamerauLevDist() != tt.expected {
				t.Errorf("Damarau-LevenshteinDistance - expected %d - got %d / %d / %d",
					tt.expected,
					helperResult,
					result,
					*builderResult.comparisonData.GetDamerauLevDist())
			}
		})
	}
}

func TestOSADamerauLevenshteinDistance(t *testing.T) {
	test := []struct {
		name     string
		input1   string
		input2   string
		expected int
	}{
		{"OSADamerauLevenshteinDistanceDistance1", "hello", "hello", 0},
		{"OSADamerauLevenshteinDistanceDistance2", "hello", "helloo", 1},
		{"OSADamerauLevenshteinDistance3", "ABCDEFG", "abcdefg", 7},
		{"OSADamerauLevenshteinDistance4", "ABCDEFG", "ABCDEFGH", 1},
		{"OSADamerauLevenshteinDistance5", "hello", "world", 4},
		{"OSADamerauLevenshteinDistance6", "My name is John", "My name is Jane", 3},
		{"OSADamerauLevenshteinDistance7", "tpyo", "typo", 1},
		{"OSADamerauLevenshteinDistance8", "teal", "tale", 2},
		{"OSADamerauLevenshteinDistance9", "For All Mankind", "For All of Maknidn", 5},
		{"OSADamerauLevenshteinDistance10", "ABCDEFG", "badcfeg", 7},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := osaDamerauLevenshteinDistance(tt.input1, tt.input2)
			result := OSADamerauLevenshteinDistance(tt.input1, tt.input2)
			builderResult := New(tt.input1).OSADamerauLevenshteinDistance(tt.input2)
			if helperResult != tt.expected ||
				result != tt.expected ||
				*builderResult.comparisonData.GetOSADamerauLevDist() != tt.expected {
				t.Errorf("OSALevenshteinDistance - expected %d - got %d / %d / %d",
					tt.expected,
					helperResult,
					result,
					*builderResult.comparisonData.GetOSADamerauLevDist())
			}
		})
	}
}

func TestLCS(t *testing.T) {
	test := []struct {
		name     string
		input1   string
		input2   string
		expected int
	}{
		{"LCS1", "hello", "hello", 5},
		{"LCS2", "hello", "hello world", 5},
		{"LCS3", "ABCDEFG", "abcdefg", 0},
		{"LCS4", "ABCDEFG", "ABCDEFGH", 7},
		{"LCS5", "hello", "world", 1},
		{"LCS6", "My name is John", "My name is Jane", 13},
		{"LCS7", "tpyo", "typo", 3},
		{"LCS8", "teal", "tale", 3},
		{"LCS9", "For All Mankind", "For All of Maknidn", 13},
		{"LCS10", "ABCDEFG", "badcfeg", 0},
		{"LCS11",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			56,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := lcs(tt.input1, tt.input2)
			result := LCS(tt.input1, tt.input2)
			builderResult := New(tt.input1).LCS(tt.input2)
			if helperResult != tt.expected || result != tt.expected || *builderResult.comparisonData.GetLCS() != tt.expected {
				t.Errorf("LCS - expected %d - got %d / %d / %d",
					tt.expected, helperResult, result, *builderResult.comparisonData.GetLCS())
			}
		})
	}
}

func TestLCSBacktrack(t *testing.T) {
	tests := []struct {
		name     string
		input1   string
		input2   string
		expected string
	}{
		{"LCSBacktrack1", "hello", "hello", "hello"},
		{"LCSBacktrack2", "hello", "hello world", "hello"},
		{"LCSBacktrack3", "ABCDEFG", "abcdefg", ""},
		{"LCSBacktrack4", "ABCDEFG", "ABCDEFGH", "ABCDEFG"},
		{"LCSBacktrack5", "rust", "golang", ""},
		{"LCSBacktrack6", "My name is John", "My name is Jane", "My name is Jn"},
		{"LCSBacktrack7", "tpyo", "typo", "tpo"},
		{"LCSBacktrack8", "teal", "tale", "tal"},
		{"LCSBacktrack9", "For All Mankind", "For All of Maknidn", "For All Manin"},
		{"LCSBacktrack10",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			"Some text dolor sit amet, consectetur an end",
			"oe  dolor sit amet, consectetur an e",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult, err := lcsBacktrack(tt.input1, tt.input2)
			if err != nil {
				t.Errorf("Error: %s", err)
			}
			result, err := LCSBacktrack(tt.input1, tt.input2)
			if err != nil {
				t.Errorf("Error: %s", err)
			}
			builderResult := New(tt.input1).LCSBacktrack(tt.input2)
			fmt.Printf("LCSBacktrack - expected %s - got %s / %s / %s",
				tt.expected, helperResult, result, *builderResult.comparisonData.GetLCSBacktrack())
			if helperResult != tt.expected ||
				result != tt.expected ||
				*builderResult.comparisonData.GetLCSBacktrack() != tt.expected {
				t.Errorf("LCSBacktrack - expected %s - got %s / %s / %s",
					tt.expected, helperResult, result, *builderResult.comparisonData.GetLCSBacktrack())
			}
		})
	}
}

func TestLCSBacktrackAll(t *testing.T) {
	tests := []struct {
		name     string
		input1   string
		input2   string
		expected []string
	}{
		{"LCSBacktrackAll1", "hello", "hello", []string{"hello"}},
		{"LCSBacktrackAll2", "hello", "hello world", []string{"hello"}},
		{"LCSBacktrackAll3", "ABCDEFG", "abcdefg", []string{""}},
		{"LCSBacktrackAll4", "ABCDEFG", "ABCDEFGH", []string{"ABCDEFG"}},
		{"LCSBacktrackAll5", "rust", "golang", []string{""}},
		{"LCSBacktrackAll6", "My name is John", "My name is Jane", []string{"My name is Jn"}},
		{"LCSBacktrackAll7", "tpyo", "typo", []string{"tpo", "tyo"}},
		{"LCSBacktrackAll8", "teal", "tale", []string{"tal"}},
		{"LCSBacktrackAll9",
			"For All Mankind",
			"For All of Maknidn",
			[]string{"For All Makin", "For All Manid", "For All Makid", "For All Maknd", "For All Manin"}},
		{"LCSBacktrackAll10",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			"Some text dolor sit amet, consectetur an end",
			[]string{"oe  dolor sit amet, consectetur an e", "om  dolor sit amet, consectetur an e"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult, err := lcsBacktrackAll(tt.input1, tt.input2)
			if err != nil {
				t.Errorf("Error: %s", err)
			}
			result, err := LCSBacktrackAll(tt.input1, tt.input2)
			if err != nil {
				t.Errorf("Error: %s", err)
			}
			builderResult := New(tt.input1).LCSBacktrackAll(tt.input2)
			if tt.expected != nil && (!CompareStringSlices(tt.expected, helperResult, false) ||
				!CompareStringSlices(tt.expected, result, false) ||
				!CompareStringSlices(tt.expected, *builderResult.comparisonData.GetLCSBacktrackAll(), false)) {
				t.Errorf("LCSBacktrackAllA - expected %s - got %s / %s / %s",
					tt.expected, helperResult, result, *builderResult.comparisonData.GetLCSBacktrackAll())
			}
			if tt.expected == nil && (helperResult != nil ||
				result != nil ||
				*builderResult.comparisonData.GetLCSBacktrackAll() != nil) {
				t.Errorf("LCSBacktrackAllB - expected %d - got %d / %d / %d",
					len(tt.expected), len(helperResult), len(result), len(*builderResult.comparisonData.GetLCSBacktrackAll()))
			}
		})
	}
}

func TestCompareStringSlices(t *testing.T) {
	tests := []struct {
		name     string
		input1   []string
		input2   []string
		nulls    bool
		expected bool
	}{
		{"CompareStringSlices1", []string{"hello", "world"}, []string{"hello", "world"}, true, true},
		{"CompareStringSlices2", []string{"hello", "world"}, []string{"hello", "world"}, false, true},
		{"CompareStringSlices3", []string{"hello", "world"}, []string{"world", "hello"}, true, true},
		{"CompareStringSlices4", []string{"hello", "world"}, []string{"world", "hello"}, false, true},
		{"CompareStringSlices5", []string{"hello", "world"}, []string{"hello", "world", "hello"}, true, false},
		{"CompareStringSlices6", []string{"hello", "world"}, []string{"hello", "world", "hello"}, false, false},
		{"CompareStringSlices7", nil, nil, false, false},
		{"CompareStringSlices8", nil, nil, true, true},
		{"CompareStringSlices9", []string{"hello", "world"}, nil, false, false},
		{"CompareStringSlices10", nil, []string{"hello", "world"}, true, false},
		{"CompareStringSlices11", nil, []string{"hello", "world"}, false, false},
		{"CompareStringSlices12", []string{"hello", "world"}, nil, true, false},
		{"CompareStringSlices13",
			[]string{"hello", "world", "hello", "world"},
			[]string{"hello", "world", "goodnight", "moon"},
			false,
			false},
		{"CompareStringSlices14", []string{"hello", "world"}, []string{"hello", "hello"}, true, false},
		{"CompareStringSlices15",
			[]string{"to", "be", "or", "not", "to", "be"},
			[]string{"be", "or", "to", "not", "be", "to"},
			false,
			true},
		{"CompareStringSlices16",
			[]string{"to", "be", "or", "not", "to", "be"},
			[]string{"be", "or", "to", "not", "not", "not"},
			false,
			false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := compareStringSlices(tt.input1, tt.input2, tt.nulls)
			result := CompareStringSlices(tt.input1, tt.input2, tt.nulls)
			if helperResult != tt.expected || result != tt.expected {
				t.Errorf("CompareStringSlices - expected %t - got %t / %t",
					tt.expected,
					helperResult,
					result)
			}
		})
	}
}

func TestCompareStringBuilderSlices(t *testing.T) {
	tests := []struct {
		name     string
		input1   []StringBuilder
		input2   []StringBuilder
		nulls    bool
		expected bool
	}{
		{"CompareStringBuilderSlices1",
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			true,
			true},
		{"CompareStringBuilderSlices2",
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			false,
			true},
		{"CompareStringBuilderSlices3",
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			[]StringBuilder{{value: "world"}, {value: "hello"}},
			true,
			true},
		{"CompareStringBuilderSlices4",
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			[]StringBuilder{{value: "world"}, {value: "hello"}},
			false,
			true},
		{"CompareStringBuilderSlices5",
			nil,
			nil,
			true,
			true},
		{"CompareStringBuilderSlices6",
			nil,
			nil,
			false,
			false},
		{"CompareStringBuilderSlices7",
			nil,
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			true,
			false},
		{"CompareStringBuilderSlices8",
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			nil,
			false,
			false},
		{"CompareStringBuilderSlices9",
			[]StringBuilder{{value: "hello"}, {value: "world"}, {value: "hello"}},
			[]StringBuilder{{value: "hello"}, {value: "world"}},
			true,
			false},
		{"CompareStringBuilderSlices10",
			[]StringBuilder{{value: "I'm a sentence"}, {value: "so am I"}},
			[]StringBuilder{{value: "so am I"}, {value: "I'm a sentence"}},
			true,
			true},
		{"CompareStringBuilderSlices11",
			[]StringBuilder{{value: "some"}, {value: "super"}, {value: "good"}, {value: "strings"}},
			[]StringBuilder{{value: "some"}, {value: "other"}, {value: "awesome"}, {value: "strings"}},
			true,
			false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := compareStringBuilderSlices(tt.input1, tt.input2, tt.nulls)
			result := CompareStringBuilderSlices(tt.input1, tt.input2, tt.nulls)
			if helperResult != tt.expected || result != tt.expected {
				t.Errorf("CompareStringBuilderSlices - expected %t - got %t / %t",
					tt.expected,
					helperResult,
					result,
				)
			}
		})
	}
}

func TestLCSDiff(t *testing.T) {
	tests := []struct {
		name     string
		input1   string
		input2   string
		expected []string
	}{
		{"LCSDiff1",
			"hello",
			"help",
			[]string{" h e l p l o", "       + - -"}},
		{"LCSDiff2",
			"",
			"",
			nil},
		{"LCSDiff3",
			"It was the best of times",
			"It was the worst of times",
			[]string{" I t   w a s   t h e   w o r b e s t   o f   t i m e s",
				"                       + + + - -                      "},
		},
		{"LCSDiff4",
			"wEdNeSdAy",
			"WeDnEsDaY",
			[]string{" W e D n w E s D a Y d N e S d A y", " + + + + -   + + + + - - - - - - -"}},
		{"LCSDiff5",
			"Georgia",
			"Jorja",
			[]string{" J G e o r j g i a", " + - -     + - -  "}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult, err := lcsDiff(tt.input1, tt.input2)
			if err != nil && !errors.Is(err, ErrLCSDiffFailure) {
				t.Errorf("Error: %s", err)
			}
			result, err := LCSDiff(tt.input1, tt.input2)
			if err != nil && !errors.Is(err, ErrLCSDiffFailure) {
				t.Errorf("Error: %s", err)
			}
			builderResult := New(tt.input1).LCSDiff(tt.input2)
			if builderResult.Error() != nil && !errors.Is(builderResult.Error(), ErrLCSDiffFailure) {
				t.Errorf("Error: %s", builderResult.Error())
			}
			if tt.expected != nil &&
				(helperResult == nil ||
					result == nil ||
					*builderResult.comparisonData.GetLCSDiff() == nil) {
				t.Errorf("LCSDiff - expected %s - got %s / %s / %s",
					tt.expected,
					helperResult,
					result,
					*builderResult.comparisonData.GetLCSDiff())
			}
			if tt.expected == nil &&
				(helperResult != nil ||
					result != nil ||
					*builderResult.comparisonData.GetLCSDiff() != nil) {
				t.Errorf("LCSDiff - expected %s - got %s / %s / %s",
					tt.expected,
					helperResult,
					result,
					*builderResult.comparisonData.GetLCSDiff())
			}
			if tt.expected != nil && (!CompareStringSlices(tt.expected, helperResult, false) ||
				!CompareStringSlices(tt.expected, result, false) ||
				!CompareStringSlices(tt.expected, *builderResult.comparisonData.GetLCSDiff(), false)) {
				t.Errorf("LCSDiff - expected %s - got %s / %s",
					tt.expected,
					helperResult,
					result,
				)
			}
		})
	}
}

func TestLCSEditDistance(t *testing.T) {
	tests := []struct {
		name     string
		input1   string
		input2   string
		expected int
	}{
		{"LCSEditDistance1",
			"hello",
			"help",
			3,
		},
		{"LCSEditDistance2",
			"",
			"",
			0,
		},
		{"LCSEditDistance3",
			"It was the best of times",
			"It was the worst of times",
			5,
		},
		{"LCSEditDistance4",
			"wEdNeSdAy",
			"WeDnEsDaY",
			16,
		},
		{"LCSEditDistance5",
			"Georgia",
			"Jorja",
			6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResul := lcsEditDistance(tt.input1, tt.input2)
			result := LCSEditDistance(tt.input1, tt.input2)
			builderResult := New(tt.input1).LCSEditDistance(tt.input2)
			if helperResul != tt.expected ||
				result != tt.expected ||
				*builderResult.comparisonData.GetLCSEditDistance() != tt.expected {
				t.Errorf("LCSEditDistance - expected %d - got %d / %d /%d",
					tt.expected,
					helperResul,
					result,
					*builderResult.comparisonData.GetLCSEditDistance(),
				)
			}
		})
	}
}

func TestHammingDistance(t *testing.T) {

	testVal2 := 3
	testVal3 := 3
	testVal4 := 2
	testVal5 := 5
	testVal6 := 0

	tests := []struct {
		name     string
		input1   string
		input2   string
		expected *int
	}{
		{"HammingDistance1", "hello", "help", nil},
		{"HammingDistance2", "karolin", "kathrin", &testVal2},
		{"HammingDistance3", "karolin", "kerstin", &testVal3},
		{"HammingDistance4", "10111", "10010", &testVal4},
		{"HammingDistance5", "00000", "11111", &testVal5},
		{"HammingDistance6", "11111", "11111", &testVal6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult, err := hammingDistance(tt.input1, tt.input2)
			if err != nil && helperResult != nil {
				t.Errorf("Error: %s", err)
			}
			result, err := HammingDistance(tt.input1, tt.input2)
			if err != nil && result != nil {
				t.Errorf("Error: %s", err)
			}
			builderResult := New(tt.input1).HammingDistance(tt.input2)
			builderErr := builderResult.Error()
			builderHD := builderResult.ComparisonData().GetHammingDist()
			if builderErr != nil && builderHD != nil {
				t.Errorf("Error: %s", builderErr)
			}
			if tt.expected != nil &&
				(*helperResult != *tt.expected ||
					*result != *tt.expected ||
					*builderHD != *tt.expected) {
				t.Errorf("HammingDistance - expected %d - got %d / %d / %d",
					*tt.expected,
					*helperResult,
					*result,
					*builderHD,
				)
			}
		})
	}
}

func TestJaroSimilarity(t *testing.T) {

	const float64EqualityThreshold = 1e-6

	tests := []struct {
		name     string
		input1   string
		input2   string
		expected float32
	}{
		{"JaroSimilarity1", "hello", "help", 0.783333},
		{"JaroSimilarity5", "00000", "11111", 0.0},
		{"JaroSimilarity6", "11111", "11111", 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := jaroSimilarity(tt.input1, tt.input2)
			result := JaroSimilarity(tt.input1, tt.input2)
			builderResult := New(tt.input1).JaroSimilarity(tt.input2)
			if math.Abs(float64(tt.expected)-float64(helperResult)) > float64EqualityThreshold ||
				math.Abs(float64(tt.expected)-float64(result)) > float64EqualityThreshold ||
				math.Abs(float64(tt.expected)-float64(*builderResult.ComparisonData().GetJaroSimilarity())) >
					float64EqualityThreshold {
				t.Errorf("JaroSimilarity - expected %f - got %f / %f / %f",
					tt.expected,
					helperResult,
					result,
					*builderResult.ComparisonData().GetJaroSimilarity(),
				)
			}
		})
	}
}

const float64EqualityThreshold = 1e-6

func TestJaroWinklerSimilarity(t *testing.T) {

	tests := []struct {
		name     string
		input1   string
		input2   string
		expected float32
	}{
		{"JaroWinklerSimilarity1", "hello", "help", 0.848333},
		{"JaroWinklerSimilarity5", "00000", "11111", 0.0},
		{"JaroWinklerSimilarity6", "11111", "11111", 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := jaroWinklerSimilarity(tt.input1, tt.input2)
			result := JaroWinklerSimilarity(tt.input1, tt.input2)
			builderResult := New(tt.input1).JaroWinklerSimilarity(tt.input2)
			if math.Abs(float64(tt.expected)-float64(helperResult)) > float64EqualityThreshold ||
				math.Abs(float64(tt.expected)-float64(result)) > float64EqualityThreshold ||
				math.Abs(float64(tt.expected)-float64(*builderResult.ComparisonData().GetJaroWinklerSim())) >
					float64EqualityThreshold {
				t.Errorf("JaroWinklerSimilarity - expected %f - got %f / %f / %f",
					tt.expected,
					helperResult,
					result,
					*builderResult.ComparisonData().GetJaroWinklerSim(),
				)
			}
		})
	}
}

func TestJaccardSimilarity(t *testing.T) {

	var val1 float32 = 1.0
	var val2 float32 = 0.25
	var val3 float32 = 0.5
	var val4 float32 = 0.068966
	var val5 float32 = 0.125
	var val6 float32 = 0.666667
	var val7 float32 = 0.75
	var val8 float32 = 1.0
	var val10 float32 = 0.0

	tests := []struct {
		name        string
		input1      string
		input2      string
		splitLength int
		expected    *float32
	}{
		{"JaccardSimilarity1", "hello", "hello", 0, &val1},
		{"JaccardSimilarity1", "hello", "help", 3, &val2},
		{"JaccardSimilarity1", "abcd", "abc", 3, &val3},
		{"JaccardSimilarity1", "this is a sentence", "this guy sent me home", 5, &val4},
		{"JaccardSimilarity1", "this is a sentence", "this guy sent me home", 0, &val5},
		{"JaccardSimilarity1", "abcd", "abc", 2, &val6},
		{"JaccardSimilarity1", "abcd", "abc", 1, &val7},
		{"JaccardSimilarity1", "abcd", "abcd", 4, &val8},
		{"JaccardSimilarity1", "abcd", "abc", -1, nil},
		{"JaccardSimilarity1", "abc", "xyz", 1, &val10},
		{"JaccardSimilarity1", "abc", "", 1, &val10},
		{"JaccardSimilarity1", "", "xyz", 1, &val10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := jaccardSimilarity(tt.input1, tt.input2, tt.splitLength)
			result := JaccardSimilarity(tt.input1, tt.input2, tt.splitLength)
			builderResult := New(tt.input1).JaccardSimilarity(tt.input2, tt.splitLength)
			if tt.expected != nil && (math.Abs(float64(*tt.expected)-float64(*helperResult)) > float64EqualityThreshold ||
				math.Abs(float64(*tt.expected)-float64(*result)) > float64EqualityThreshold ||
				math.Abs(float64(*tt.expected)-float64(*builderResult.ComparisonData().GetJaccardSim())) >
					float64EqualityThreshold) {
				t.Errorf("JaccardSimilarity - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.ComparisonData().GetJaccardSim(),
				)
			}
			if tt.expected == nil &&
				helperResult != nil &&
				result != nil &&
				builderResult.ComparisonData().GetJaccardSim() != nil {
				t.Errorf("JaccardSimilarity - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.ComparisonData().GetJaccardSim(),
				)
			}
		})
	}
}

func TestCosineSimilarity(t *testing.T) {

	var val1 float32 = 1.0
	var val2 float32 = 0.408248
	var val3 float32 = 0.707107
	var val4 float32 = 0.129641
	var val5 float32 = 0.223607
	var val6 float32 = 0.816497
	var val7 float32 = 0.866025
	var val8 float32 = 1.0
	var val10 float32 = 0.0

	tests := []struct {
		name     string
		input1   string
		input2   string
		splitLen int
		expected *float32
	}{
		{"CosineSimilarity1", "hello", "hello", 0, &val1},
		{"CosineSimilarity2", "hello", "help", 3, &val2},
		{"CosineSimilarity3", "abcd", "abc", 3, &val3},
		{"CosineSimilarity4", "this is a sentence", "this guy sent me home", 5, &val4},
		{"CosineSimilarity5", "this is a sentence", "this guy sent me home", 0, &val5},
		{"CosineSimilarity6", "abcd", "abc", 2, &val6},
		{"CosineSimilarity7", "abcd", "abc", 1, &val7},
		{"CosineSimilarity8", "abcd", "abcd", 4, &val8},
		{"CosineSimilarity9", "abcd", "abc", -1, nil},
		{"CosineSimilarity10", "abc", "xyz", 1, &val10},
		{"CosineSimilarity11", "abc", "", 1, &val10},
		{"CosineSimilarity12", "", "xyz", 1, &val10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := cosineSimilarity(tt.input1, tt.input2, tt.splitLen)
			result := CosineSimilarity(tt.input1, tt.input2, tt.splitLen)
			builderResult := New(tt.input1).CosineSimilarity(tt.input2, tt.splitLen)
			if tt.expected != nil && (math.Abs(float64(*tt.expected)-float64(*helperResult)) > float64EqualityThreshold ||
				math.Abs(float64(*tt.expected)-float64(*result)) > float64EqualityThreshold ||
				math.Abs(float64(*tt.expected)-float64(*builderResult.ComparisonData().GetCosineSimilarity())) >
					float64EqualityThreshold) {
				t.Errorf("CosineSimilarity - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.ComparisonData().GetCosineSimilarity(),
				)
			}
			if tt.expected == nil &&
				helperResult != nil &&
				result != nil &&
				builderResult.ComparisonData().GetCosineSimilarity() != nil {
				t.Errorf("CosineSimilarity - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.ComparisonData().GetCosineSimilarity(),
				)
			}
		})
	}
}

func TestSorensenDiceCoefficient(t *testing.T) {

	var val1 float32 = 1.0
	var val2 float32 = 0.40
	var val3 float32 = 0.666667
	var val4 float32 = 0.129032
	var val5 float32 = 0.223607
	var val6 float32 = 0.80
	var val7 float32 = 0.857142
	var val8 float32 = 1.0
	var val10 float32 = 0.0

	tests := []struct {
		name     string
		input1   string
		input2   string
		splitLen int
		expected *float32
	}{
		{"SorensenDiceCoefficient1", "hello", "hello", 0, &val1},
		{"SorensenDiceCoefficient2", "hello", "help", 3, &val2},
		{"SorensenDiceCoefficient3", "abcd", "abc", 3, &val3},
		{"SorensenDiceCoefficient4", "this is a sentence", "this guy sent me home", 5, &val4},
		{"SorensenDiceCoefficient5", "this is a sentence", "this guy sent me home", 0, &val5},
		{"SorensenDiceCoefficient6", "abcd", "abc", 2, &val6},
		{"SorensenDiceCoefficient7", "abcd", "abc", 1, &val7},
		{"SorensenDiceCoefficient8", "abcd", "abcd", 4, &val8},
		{"SorensenDiceCoefficient9", "abcd", "abc", -1, nil},
		{"SorensenDiceCoefficient10", "abc", "xyz", 1, &val10},
		{"SorensenDiceCoefficient11", "abc", "", 4, &val10},
		{"SorensenDiceCoefficient12", "", "xyz", 1, &val10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := SorensenDiceCoefficient(tt.input1, tt.input2, tt.splitLen)
			result := SorensenDiceCoefficient(tt.input1, tt.input2, tt.splitLen)
			builderResult := New(tt.input1).SorensenDiceCoefficient(tt.input2, tt.splitLen)
			if tt.expected != nil && (math.Abs(float64(*tt.expected)-float64(*helperResult)) > float64EqualityThreshold ||
				math.Abs(float64(*tt.expected)-float64(*result)) > float64EqualityThreshold ||
				math.Abs(float64(*tt.expected)-float64(*builderResult.ComparisonData().GetSorensenDiceCo())) >
					float64EqualityThreshold) {
				t.Errorf("SorensenDiceCoefficient - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.ComparisonData().GetSorensenDiceCo(),
				)
			}
			if tt.expected == nil &&
				helperResult != nil &&
				result != nil &&
				builderResult.ComparisonData().GetSorensenDiceCo() != nil {
				t.Errorf("SorensenDiceCoefficient - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.ComparisonData().GetSorensenDiceCo(),
				)
			}
		})
	}
}

func TestQGramDistance(t *testing.T) {

	var val1 = 0
	var val2 = 3
	var val3 = 6
	var val4 = 1
	var val5 = 21
	var val6 = 23
	var val7 = 6
	var val8 = 2
	var val9 = 1
	var val10 = 6
	var val11 = 6
	var val12 = 0

	tests := []struct {
		name     string
		input1   string
		input2   string
		q        int
		expected *int
	}{
		{"QGramDist1", "Hello", "Hello", 1, &val1},
		{"QGramDist2", "Hello", "Help", 3, &val2},
		{"QGramDist3", "Hello", "Helping", 2, &val3},
		{"QGramDist4", "abcd", "abc", 4, &val4},
		{"QGramDist5", "this is a sentence", "this guy sent me home", 10, &val5},
		{"QGramDist6", "this is a sentence", "this guy sent me home", 3, &val6},
		{"QGramDist7", "abc", "xyz", 1, &val7},
		{"QGramDist8", "abc", "", 2, &val8},
		{"QGramDist9", "", "xyz", 3, &val9},
		{"QGramDist10", "HELLO", "hello", 3, &val10},
		{"QGramDist11", "HELLO", "hello", -5, nil},
		{"QGramDist11", "HELLO", "hellohellohello", 10, &val11},
		{"QGramDist11", "HELLOHELLOHELLO", "hello", 20, &val12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := qgramDistance(tt.input1, tt.input2, tt.q)
			result := QgramDistance(tt.input1, tt.input2, tt.q)
			builderResult := New(tt.input1).QgramDistance(tt.input2, tt.q).comparisonData.GetQGramDist()
			if tt.expected != nil && (*helperResult != *tt.expected ||
				*result != *tt.expected ||
				*builderResult != *tt.expected) {
				t.Errorf("QgramDistance - expected %d - got %d / %d / %d",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult,
				)
			}
			if tt.expected == nil && (helperResult != nil ||
				result != nil ||
				builderResult != nil) {
				t.Errorf("QgramDistance - expected %d - got %d / %d / %d",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult,
				)
			}
		})
	}
}

func TestQgramDistanceCustomNgram(t *testing.T) {

	testMap1 := map[string]int{
		"he": 1,
		"el": 1,
		"ll": 1,
		"lo": 1,
	}
	testMap2 := map[string]int{
		"he": 1,
		"el": 1,
		"lp": 1,
	}
	testMap3 := map[string]int{
		"this": 5,
		"help": 2,
		"gets": 1,
	}
	testMap4 := map[string]int{
		"this": 1,
		"gets": 1,
		"home": 8,
	}

	tests := []struct {
		name     string
		input1   map[string]int
		input2   map[string]int
		expected int
	}{
		{"QgramDist1", testMap1, testMap1, 0},
		{"QgramDist2", testMap1, testMap2, 3},
		{"QgramDist3", testMap1, testMap3, 12},
		{"QgramDist4", testMap1, testMap4, 14},
		{"QgramDist5", testMap2, testMap3, 11},
		{"QgramDist6", testMap4, testMap2, 13},
		{"QgramDist7", testMap3, testMap4, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := qgramDistanceCustomNgram(tt.input1, tt.input2)
			result := QgramDistanceCustomNgram(tt.input1, tt.input2)
			if helperResult != tt.expected || result != tt.expected {
				t.Errorf("QgramDistanceCustomNgram - expected %d - got %d / %d",
					tt.expected,
					helperResult,
					result,
				)
			}
		})
	}
}

func TestQgramDistanceCustomNgramBuilder(t *testing.T) {
	tests := []struct {
		name     string
		input1   *StringBuilder
		input2   map[string]int
		expected int
	}{
		{"QgramDist1", New("hello"), map[string]int{"he": 1, "el": 1, "ll": 1, "lo": 1}, 0},
		{"QgramDist2", New("hello"), map[string]int{"he": 1, "el": 1, "lp": 1}, 3},
		{"QgramDist3", New("this is a sentence"), map[string]int{"this": 5, "help": 2, "gets": 1}, 21},
		{"QgramDist4", New("this is a sentence"), map[string]int{"this": 1, "gets": 1, "home": 8}, 23},
		{"QgramDist1", nil, map[string]int{"he": 1, "el": 1, "ll": 1, "lo": 1}, 0},
		{"QgramDist2", nil, map[string]int{"he": 1, "el": 1, "lp": 1}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input1 == nil {
				tt.input1 = New("hello").WithComparisonManager().Shingle(2)
			}
			builderResult := tt.input1.QgramDistanceCustomNgram(tt.input2)
			if *builderResult.ComparisonData().GetQGramDistCustom() != tt.expected {
				t.Errorf("QgramDistanceCustomNgramBuilder - expected %d - got %d",
					tt.expected,
					*builderResult.ComparisonData().GetQGramDistCustom(),
				)
			}
		})
	}
}

func TestQgramSimilarity(t *testing.T) {

	var val1 float32 = 0.588235
	var val2 float32 = 0.00
	var val3 float32 = 0.00
	var val4 float32 = 0.782609
	var val5 float32 = 0.0
	var val6 float32 = 0.0
	var val7 float32 = 0.631579
	var val8 float32 = 0.0
	var val9 float32 = 0.0
	var val10 float32 = 0.0

	tests := []struct {
		name     string
		input1   string
		input2   string
		q        int
		expected *float32
	}{
		{"QgramSim1", "Hello World", "ello world", 3, &val1},
		{"QgramSim2", "Hello", "help", 3, &val2},
		{"QgramSim3", "Hello", "World", 3, &val3},
		{"QgramSim4", "It was the best of times.",
			"It was the worst of times", 3, &val4},
		{"QgramSim5", "rock", "paper", 3, &val5},
		{"QgramSim6", "orea", "books", 3, &val6},
		{"QgramSim7", "Hello World", "Hello world!", 3, &val7},
		{"QgramSim8", "HELLO WORLD", "hello world", 3, &val8},
		{"QgramSim9", "Hello World", "", 3, &val9},
		{"QgramSim10", "", "Hello world", 3, &val10},
		{"QgramSim11", "Hello World", "ello world", -99, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := qgramSimilarity(tt.input1, tt.input2, tt.q)
			result := QgramSimilarity(tt.input1, tt.input2, tt.q)
			builderResult := New(tt.input1).QgramSimilarity(tt.input2, tt.q)
			if tt.expected != nil && (math.Abs(float64(*tt.expected)-float64(*helperResult)) > float64EqualityThreshold ||
				math.Abs(float64(*tt.expected)-float64(*result)) > float64EqualityThreshold ||
				math.Abs(float64(*tt.expected)-float64(*builderResult.ComparisonData().GetQGramSim())) >
					float64EqualityThreshold) {
				t.Errorf("QgramSimilarity - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.ComparisonData().GetQGramSim(),
				)
			}
		})
	}
}

func TestShingle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		k        int
		expected map[string]int
	}{
		{"Shingle1", "Hello World", 1, map[string]int{" ": 1, "H": 1, "e": 1, "l": 3, "o": 2, "W": 1, "r": 1, "d": 1}},
		{"Shingle2",
			"Hello World",
			2,
			map[string]int{"He": 1, "el": 1, "ll": 1, "lo": 1, "o ": 1, " W": 1, "Wo": 1, "or": 1, "rl": 1, "ld": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := shingle(tt.input, tt.k)
			result := Shingle(tt.input, tt.k)
			for k, v := range tt.expected {
				if (*helperResult.shingles)[k] == 0 || (*result.shingles)[k] == 0 {
					t.Errorf("Shingle 0 - expected %d - got %d / %d",
						v,
						(*helperResult.shingles)[k],
						(*result.shingles)[k],
					)
				}
				if (*helperResult.shingles)[k] != v ||
					(*result.shingles)[k] != v {
					t.Errorf("Shingle val - expected %d - got %d / %d",
						v,
						(*helperResult.shingles)[k],
						(*result.shingles)[k],
					)
				}
				builderResult := New(tt.input).
					WithComparisonManager().
					Shingle(tt.k).
					ComparisonManager().
					ShingleData[ShinglesMap][tt.k]
				if brMap, ok := (*builderResult).(*ShingleMapResult); ok {
					if (*brMap.shingles)[k] == 0 || (*brMap.shingles)[k] != v {
						t.Errorf("Shingle b0 - expected %d - got %d",
							v,
							(*brMap.shingles)[k],
						)
					}
					if (*brMap.shingles)[k] != v {
						t.Errorf("Shingle bval - expected %d - got %d",
							v,
							(*brMap.shingles)[k],
						)
					}
				}
			}
		})
	}
}

func TestShingleSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		k        int
		expected []string
	}{
		{"ShingleSlice1", "Hello World", 1, []string{" ", "H", "e", "l", "o", "W", "r", "d"}},
		{"ShingleSlice2",
			"Hello World",
			2,
			[]string{"He", "el", "ll", "lo", "o ", " W", "Wo", "or", "rl", "ld"},
		},
		{"ShingleSlice3",
			"Hello World",
			3,
			[]string{"Hel", "ell", "llo", "lo ", "o W", " Wo", "Wor", "orl", "rld"}},
		{"ShingleSlice4",
			"Hello World", 4,
			[]string{"Hell", "ello", "llo ", "lo W", "o Wo", " Wor", "Worl", "orld"}},
		{"ShingleSlice5",
			"Hello World",
			5,
			[]string{"Hello", "ello ", "llo W", "lo Wo", "o Wor", " Worl", "World"}},
		{"ShingleSlice6", "Hello World", 6, []string{"Hello ", "ello W", "llo Wo", "lo Wor", "o Worl", " World"}},
		{"ShingleSlice7", "Hello World", 7, []string{"Hello W", "ello Wo", "llo Wor", "lo Worl", "o World"}},
		{"ShingleSlice8", "Hello World", 8, []string{"Hello Wo", "ello Wor", "llo Worl", "lo World"}},
		{"ShingleSlice9", "Hello World", 9, []string{"Hello Wor", "ello Worl", "llo World"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := shingleSlice(tt.input, tt.k)
			result := ShingleSlice(tt.input, tt.k)
			if !CompareStringSlices(*helperResult.shingles, tt.expected, false) ||
				!CompareStringSlices(*result.shingles, tt.expected, false) {
				t.Errorf("ShingleSlice - expected %v - got %v",
					tt.expected,
					*helperResult,
				)
			}
			builderResult := New(tt.input).ShingleSlice(tt.k).ComparisonManager().ShingleData[ShinglesSlice][tt.k]
			if brSlice, ok := (*builderResult).(*ShingleSliceResult); ok {
				if !CompareStringSlices(*brSlice.shingles, tt.expected, false) {
					t.Errorf("ShingleSlice - expected %v - got %v",
						tt.expected,
						*brSlice,
					)
				}
			}
		})
	}
}

func TestSimilarity(t *testing.T) {

	var val1 float32 = 1.0
	var val2 float32 = 0.2
	var val3 float32 = 1.0
	var val4 float32 = 0.2
	var val5 float32 = 1.0
	var val6 float32 = 0.2
	var val7 float32 = 1.0
	var val8 float32 = -.60
	var val9 float32 = 1.0
	var val10 float32 = 0.2
	var val11 float32 = 1.0
	var val12 float32 = 0.466667
	var val13 float32 = 1.0
	var val14 float32 = 0.466667
	var val15 float32 = 1.0
	var val16 float32 = 0.0
	var val17 float32 = 1.0
	var val18 float32 = 0.0
	var val19 float32 = 1.0
	var val20 float32 = 0.0
	var val21 float32 = 1.0
	var val22 float32 = 0.0

	tests := []struct {
		name      string
		input1    string
		input2    string
		algorithm Algorithm
		expected  *SimilarityResult
	}{
		{"Similarity1",
			"Hello World",
			"Hello World",
			Levenshtein,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  Levenshtein,
				similarity: &val1,
				err:        nil,
			}},
		{"Similarity2",
			"Hello",
			"World",
			Levenshtein,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  Levenshtein,
				similarity: &val2,
				err:        nil,
			}},
		{"Similarity3",
			"Hello World",
			"Hello World",
			DamerauLevenshtein,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  DamerauLevenshtein,
				similarity: &val3,
				err:        nil,
			}},
		{"Similarity4",
			"Hello",
			"World",
			DamerauLevenshtein,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  DamerauLevenshtein,
				similarity: &val4,
				err:        nil,
			}},
		{"Similarity5",
			"Hello World",
			"Hello World",
			OSADamerauLevenshtein,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  OSADamerauLevenshtein,
				similarity: &val5,
				err:        nil,
			}},
		{"Similarity6",
			"Hello",
			"World",
			OSADamerauLevenshtein,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  OSADamerauLevenshtein,
				similarity: &val6,
				err:        nil,
			}},
		{"Similarity7",
			"Hello World",
			"Hello World",
			Lcs,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  Lcs,
				similarity: &val7,
				err:        nil,
			}},
		{"Similarity8",
			"Hello",
			"World",
			Lcs,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  Lcs,
				similarity: &val8,
				err:        nil,
			}},
		{"Similarity9",
			"Hello World",
			"Hello World",
			Hamming,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  Hamming,
				similarity: &val9,
				err:        nil,
			}},
		{"Similarity10",
			"Hello",
			"World",
			Hamming,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  Hamming,
				similarity: &val10,
				err:        nil,
			}},
		{"Similarity11",
			"Hello World",
			"Hello World",
			Jaro,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  Jaro,
				similarity: &val11,
				err:        nil,
			}},
		{"Similarity12",
			"Hello",
			"World",
			Jaro,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  Jaro,
				similarity: &val12,
				err:        nil,
			}},
		{"Similarity13",
			"Hello World",
			"Hello World",
			JaroWinkler,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  JaroWinkler,
				similarity: &val13,
				err:        nil,
			}},
		{"Similarity14",
			"Hello",
			"World",
			JaroWinkler,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  JaroWinkler,
				similarity: &val14,
				err:        nil,
			}},
		{"Similarity15",
			"Hello World",
			"Hello World",
			Cosine,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  Cosine,
				similarity: &val15,
				err:        nil,
			}},
		{"Similarity16",
			"Hello",
			"World",
			Cosine,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  Cosine,
				similarity: &val16,
				err:        nil,
			}},
		{"Similarity17",
			"Hello World",
			"Hello World",
			Jaccard,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  Jaccard,
				similarity: &val17,
				err:        nil,
			}},
		{"Similarity18",
			"Hello",
			"World",
			Jaccard,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  Jaccard,
				similarity: &val18,
				err:        nil,
			}},
		{"Similarity19",
			"Hello World",
			"Hello World",
			SorensenDice,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  SorensenDice,
				similarity: &val19,
				err:        nil,
			}},
		{"Similarity20",
			"Hello",
			"World",
			SorensenDice,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  SorensenDice,
				similarity: &val20,
				err:        nil,
			}},
		{"Similarity21",
			"Hello World",
			"Hello World",
			QGram,
			&SimilarityResult{
				string1:    "Hello World",
				string2:    "Hello World",
				algorithm:  QGram,
				similarity: &val21,
				err:        nil,
			}},
		{"Similarity22",
			"Hello",
			"World",
			QGram,
			&SimilarityResult{
				string1:    "Hello",
				string2:    "World",
				algorithm:  QGram,
				similarity: &val22,
				err:        nil,
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := similarity(tt.input1, tt.input2, tt.algorithm)
			result := Similarity(tt.input1, tt.input2, tt.algorithm)
			brSim := New(tt.input1).
				WithComparisonManager().
				Similarity(tt.input2, tt.algorithm).
				comparisonManager.
				SimilarityResults[tt.algorithm][tt.input2]
			if helperResult.algorithm != tt.expected.algorithm ||
				helperResult.string1 != tt.expected.string1 ||
				helperResult.string2 != tt.expected.string2 ||
				math.Abs(float64(*tt.expected.GetSimilarity())-float64(*helperResult.GetSimilarity())) > 1e-6 ||
				!errors.Is(helperResult.err, tt.expected.err) {
				t.Errorf("SimilarityA - expected %f - got %f",
					*tt.expected.GetSimilarity(),
					*helperResult.GetSimilarity(),
				)
			}
			if result.algorithm != tt.expected.algorithm ||
				result.string1 != tt.expected.string1 ||
				result.string2 != tt.expected.string2 ||
				math.Abs(float64(*tt.expected.GetSimilarity())-float64(*result.GetSimilarity())) > 1e-6 ||
				!errors.Is(result.err, tt.expected.err) {
				t.Errorf("SimilarityB - expected %f - got %f\n",
					*tt.expected.GetSimilarity(),
					*helperResult.GetSimilarity(),
				)
			}
			if brSim.algorithm != tt.expected.algorithm ||
				brSim.string1 != tt.expected.string1 ||
				brSim.string2 != tt.expected.string2 ||
				math.Abs(float64(*tt.expected.GetSimilarity())-float64(*brSim.GetSimilarity())) > 1e-6 ||
				!errors.Is(brSim.err, tt.expected.err) {
				t.Errorf("SimilarityC - expected %f - got %f",
					*tt.expected.GetSimilarity(),
					*helperResult.GetSimilarity(),
				)
			}
		})
	}
}
