package strutil

import (
	"strings"
	"testing"
)

func TestLoremWord(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremWord1"},
		{"LoremWord2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremWord()
			result := LoremWord()
			builderResult := NewLoremWord().String()
			if helperResult == "" || builderResult == "" || result == "" {
				t.Errorf("LoremWord - %q, %q, %q", result, helperResult, builderResult)
			}

		})
	}
}

func TestLoremWords(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{"LoremWords1", 1},
		{"LoremWords30", 30},
		{"LoremWords100", 100},
		{"LoremWordsZero", 0},
		{"LoremWordsNegative", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremWords(tt.count)
			result := LoremWords(tt.count)
			builderResult := NewLoremWords(tt.count).String()

			if tt.count < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremWords - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.count > 0 && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremWords - %q, %q, %q", result, helperResult, builderResult)
			}

			helperResultCount := len(strings.Split(helperResult, " ")) - 1
			helperArr := strings.Split(helperResult, " ")
			resultCount := len(strings.Split(result, " ")) - 1
			resultArr := strings.Split(result, " ")
			builderResultCount := len(strings.Split(builderResult, " ")) - 1
			builderArr := strings.Split(builderResult, " ")
			if tt.count > 0 && (helperResultCount != tt.count || resultCount != tt.count || builderResultCount != tt.count) {
				t.Errorf("LoremWords(%d, %d, %d) - expected %d - words: %q / %q / %q",
					resultCount, helperResultCount, builderResultCount, tt.count, resultArr, helperArr, builderArr)
			}
		})
	}
}

func TestLoremSentence(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremSentence1"},
		{"LoremSentence2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremSentence()
			result := LoremSentence()
			builderResult := NewLoremSentence().String()
			if helperResult == "" || builderResult == "" || result == "" {
				t.Errorf("LoremSentence - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if helperResultCount != 8 || builderResultCount != 8 || resultCount != 8 {
				t.Errorf("LoremSentence - expected 8 - words: %d / %d / %d", resultCount, helperResultCount, builderResultCount)
			}

			if helperResult[0] < 'A' ||
				helperResult[0] > 'Z' ||
				result[0] < 'A' || result[0] > 'Z' ||
				builderResult[0] < 'A' ||
				builderResult[0] > 'Z' {
				t.Errorf("LoremSentence - first character not uppercase: %q / %q / %q",
					helperResult[0], result[0], builderResult[0])
			}

			if helperResult[len(helperResult)-1] != '.' ||
				result[len(result)-1] != '.' ||
				builderResult[len(builderResult)-1] != '.' {
				t.Errorf("LoremSentence - last character not period: %q / %q / %q",
					helperResult[len(helperResult)-1], result[len(result)-1], builderResult[len(builderResult)-1])
			}
		})
	}
}

func TestLoremSentenceCustom(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"LoremSentenceCustom1", 1},
		{"LoremSentenceCustom11", 11},
		{"LoremSentenceCustom25", 25},
		{"LoremSentenceCustomZero", 0},
		{"LoremSentenceCustomNeg", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremSentenceCustom(tt.length)
			result := LoremSentenceCustom(tt.length)
			builderResult := NewLoremSentenceCustom(tt.length).String()
			if tt.length > 0 && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremSentenceCustom - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.length < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremSentenceCustom - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if tt.length > 0 && (helperResultCount != tt.length || builderResultCount != tt.length || resultCount != tt.length) {
				t.Errorf("LoremSentenceCustom - expected %d - words: %d / %d / %d",
					tt.length, resultCount, helperResultCount, builderResultCount)
			}

			if tt.length > 0 &&
				(helperResult[0] < 'A' ||
					helperResult[0] > 'Z' ||
					result[0] < 'A' ||
					result[0] > 'Z' ||
					builderResult[0] < 'A' ||
					builderResult[0] > 'Z') {
				t.Errorf("LoremSentenceCustom - first character not uppercase: %q / %q / %q",
					helperResult[0], result[0], builderResult[0])
			}

			if tt.length > 0 &&
				(helperResult[len(helperResult)-1] != '.' ||
					result[len(result)-1] != '.' ||
					builderResult[len(builderResult)-1] != '.') {
				t.Errorf("LoremSentenceCustom - last character not period: %q / %q / %q",
					helperResult[len(helperResult)-1], result[len(result)-1], builderResult[len(builderResult)-1])
			}
		})
	}
}

func TestLoremSentences(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		expected int
	}{
		{"LoremSentences1", 1, 8},
		{"LoremSentences3", 3, 24},
		{"LoremSentences10", 10, 80},
		{"LoremSentencesZero", 0, 0},
		{"LoremSentencesNegative", -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremSentences(tt.count)
			result := LoremSentences(tt.count)
			builderResult := NewLoremSentences(tt.count).String()
			if tt.count < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremSentencesNotNil - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.count > 0 && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremSentencesIsNil - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if tt.count > 0 && (helperResultCount != tt.expected ||
				resultCount != tt.expected ||
				builderResultCount != tt.expected) {
				t.Errorf("LoremSentencesLen1 - expected %d - sentences: %d / %d / %d",
					tt.expected, helperResultCount, resultCount, builderResultCount)
			}
			if tt.count < 1 && (helperResultCount != tt.expected+1 ||
				resultCount != tt.expected+1 ||
				builderResultCount != tt.expected+1) {
				t.Errorf("LoremSentenceslen2 - expected %d - sentences: %d / %d / %d",
					tt.expected, helperResultCount, resultCount, builderResultCount)
			}
		})
	}
}

func TestLoremSentencesCustom(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		length   int
		expected int
	}{
		{"LoremSentencesCustom1", 1, 1, 1},
		{"LoremSentencesCustom11", 3, 7, 21},
		{"LoremSentencesCustom25", 5, 12, 60},
		{"LoremSentencesCustomZero", 0, 0, 0},
		{"LoremSentencesCustomNeg", -1, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremSentencesCustom(tt.count, tt.length)
			result := LoremSentencesCustom(tt.count, tt.length)
			builderResult := NewLoremSentencesCustom(tt.count, tt.length).String()
			if tt.count < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremSentencesCustomNotNil - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.count > 0 && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremSentencesCustomIsNil - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if tt.count > 0 && (helperResultCount != tt.expected ||
				resultCount != tt.expected ||
				builderResultCount != tt.expected) {
				t.Errorf("LoremSentencesCustomLen1 - expected %d - sentences: %d / %d / %d",
					tt.expected, helperResultCount, resultCount, builderResultCount)
			}
			if tt.count < 1 && (helperResultCount != tt.expected+1 ||
				resultCount != tt.expected+1 ||
				builderResultCount != tt.expected+1) {
				t.Errorf("LoremSentencesCustomLen2 - expected %d - sentences: %d / %d / %d",
					tt.expected, helperResultCount, resultCount, builderResultCount)
			}
		})
	}
}

func TestLoremSentencesVariable(t *testing.T) {
	tests := []struct {
		name  string
		count int
		min   int
		max   int
	}{
		{"LoremSentencesVariable1", 1, 1, 10},
		{"LoremSentencesVariable3", 3, 1, 10},
		{"LoremSentencesVariable10", 5, 6, 14},
		{"LoremSentencesVariableZero", 0, 0, 0},
		{"LoremSentencesVariableNeg", -1, 0, 0},
		{"LoremSentencesVariableInvalid", 3, 10, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremSentencesVariable(tt.count, tt.min, tt.max)
			result := LoremSentencesVariable(tt.count, tt.min, tt.max)
			builderResult := NewLoremSentencesVariable(tt.count, tt.min, tt.max).String()
			if tt.count < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremSentencesVariableNotNil - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.count > 0 && tt.min <= tt.max && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremSentencesVariableIsNil - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			minRange := tt.min * tt.count
			maxRange := tt.max * tt.count
			if tt.count > 0 && tt.min <= tt.max &&
				(helperResultCount < minRange ||
					helperResultCount > maxRange ||
					resultCount < minRange ||
					resultCount > maxRange ||
					builderResultCount < minRange ||
					builderResultCount > maxRange) {
				t.Errorf("LoremSentencesVariableLen1 - expected %d - %d - %d - %d - sentences: %d / %d / %d / %d",
					tt.count, tt.min, tt.max, minRange, helperResultCount, resultCount, builderResultCount, maxRange)
			}
			if tt.max < tt.min && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremSentencesVariableLen2 - expected %d - %d - %d - %d - sentences: %d / %d / %d / %d",
					tt.count, tt.min, tt.max, minRange, helperResultCount, resultCount, builderResultCount, maxRange)
			}
		})
	}
}

func TestLoremParagraph(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremParagraph1"},
		{"LoremParagraph2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremParagraph()
			result := LoremParagraph()
			builderResult := NewLoremParagraph().String()
			if helperResult == "" || builderResult == "" || result == "" {
				t.Errorf("LoremParagraph - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if helperResultCount != 45 || builderResultCount != 45 || resultCount != 45 {
				t.Errorf("LoremParagraph - expected 45 - words: %d / %d / %d", resultCount, helperResultCount, builderResultCount)
			}
		})
	}
}

func TestLoremParagraphs(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		expected int
	}{
		{"LoremParagraphs1", 1, 45},
		{"LoremParagraphs3", 3, 133},
		{"LoremParagraphs10", 10, 441},
		{"LoremParagraphs1000", 1000, 44001},
		{"LoremParagraphsZero", 0, 0},
		{"LoremParagraphsNegative", -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremParagraphs(tt.count)
			result := LoremParagraphs(tt.count)
			builderResult := NewLoremParagraphs(tt.count).String()
			if tt.count < 1 && (helperResult != "" || builderResult != "" || result != "") {
				t.Errorf("LoremParagraphsNotNil - %q, %q, %q", result, helperResult, builderResult)
			}
			if tt.count > 0 && (helperResult == "" || builderResult == "" || result == "") {
				t.Errorf("LoremParagraphsIsNil - %q, %q, %q", result, helperResult, builderResult)
			}
			helperResultCount := len(strings.Split(helperResult, " "))
			resultCount := len(strings.Split(result, " "))
			builderResultCount := len(strings.Split(builderResult, " "))
			if tt.count > 0 &&
				(helperResultCount != tt.expected ||
					resultCount != tt.expected ||
					builderResultCount != tt.expected) {
				t.Errorf("LoremParagraphsLen1 - expected %d - sentences: %d / %d / %d",
					tt.expected, helperResultCount, resultCount, builderResultCount)
			}
			if tt.count < 1 &&
				(helperResultCount != tt.expected+1 ||
					resultCount != tt.expected+1 ||
					builderResultCount != tt.expected+1) {
				t.Errorf("LoremParagraphsLen2 - expected %d - sentences: %d / %d / %d",
					tt.expected, helperResultCount, resultCount, builderResultCount)
			}
		})
	}
}

func TestLoremDomain(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremDomain1"},
		{"LoremDomain2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremDomain()
			result := LoremDomain()
			builderResult := NewLoremDomain().String()
			if helperResult == "" || builderResult == "" || result == "" {
				t.Errorf("LoremDomain - %q, %q, %q", result, helperResult, builderResult)
			}
			if !isValidDomain(result) || !isValidDomain(builderResult) || !isValidDomain(helperResult) {
				t.Errorf("LoremDomain - invalid domain: %q / %q / %q", result, builderResult, helperResult)
			}
		})
	}
}

func TestLoremURL(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremURL1"},
		{"LoremURL2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremURL()
			result := LoremURL()
			builderResult := NewLoremURL().String()
			if result == "" || builderResult == "" || helperResult == "" {
				t.Errorf("LoremURL - %q, %q", result, builderResult)
			}
			if !isValidURL(result) || !isValidURL(builderResult) || !isValidURL(helperResult) {
				t.Errorf("LoremURL - invalid url: %q / %q", result, builderResult)
			}

		})
	}
}

func TestLoremEmail(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"LoremEmail1"},
		{"LoremEmail2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := loremEmail()
			result := LoremEmail()
			builderResult := NewLoremEmail().String()
			if result == "" || builderResult == "" || helperResult == "" {
				t.Errorf("LoremEmail - %q, %q", result, builderResult)
			}
			if !isValidEmail(result) || !isValidEmail(builderResult) || !isValidEmail(helperResult) {
				t.Errorf("LoremEmail - invalid email: %q / %q", result, builderResult)
			}
		})
	}
}
