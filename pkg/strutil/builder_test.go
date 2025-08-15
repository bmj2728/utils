package strutil

import (
	"math/rand"
	"testing"

	"github.com/bmj2728/utils/pkg/internal/errors"
)

func TestBuilderGetOriginal(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Test 1", "Hello World"},
		{"Test 2", "Hi"},
		{"Test 3", "www.example.com"},
		{"Test 4", "asdfghjkl;1234567890!@#$%^&*("},
		{"Test 5", " "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := New(tt.input). //Hello World 0
						Append("append 1", "-").
						Append("append 2", "-").
						Append("append 3", "-").
						Append("append 4", "-")

			if rb.GetOriginalValue() != tt.input {
				t.Errorf("GetOriginal failed, expected Hello World, got %s", rb.GetOriginalValue())
			}
		})
	}

	rb := New("Hello World"). //Hello World 0
					Append("append 1", "-").
					Append("append 2", "-").
					Append("append 3", "-").
					Append("append 4", "-")

	if rb.GetOriginalValue() != "Hello World" {
		t.Errorf("GetOriginal failed, expected Hello World, got %s", rb.GetOriginalValue())
	}
}

func TestStringBuilder_RevertToOriginal(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					WithHistory(25).
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello World 4
					TrimCharsRight(" World"). //Hello 5
					Append("John", " ")

	if rb.RevertToOriginal().String() != "Hello World" {
		t.Errorf("RevertToOriginal failed, expected Hello World, got %s", rb.RevertToOriginal().String())
	}
}

func TestBuilderRevertToPrevious(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					WithHistory(10).
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello-World 4
					TrimCharsRight(" World"). //Hello-5
					Append("John", " ")       //Hello- John

	if rb.String() != "Hello- John" {
		t.Errorf("RevertToPrevious failed, expected Hello- John, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "Hello-" {
		t.Errorf("RevertToPrevious failed, expected Hello-, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "Hello-World" {
		t.Errorf("RevertToPrevious failed, expected Hello-World, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "hello-world" {
		t.Errorf("RevertToPrevious failed, expected hello-world, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "HELLO-WORLD" {
		t.Errorf("RevertToPrevious failed, expected HELLO-WORLD, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "helloWorld" {
		t.Errorf("RevertToPrevious failed, expected helloWorld, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "Hello World" {
		t.Errorf("RevertToPrevious failed, expected Hello World, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "" || !errors.CompareErrors(rb.Error(), errors.ErrInvalidHistoryIndex) {
		t.Errorf("RevertToPrevious failed, expected empty string, got %s", rb.String())
	}
}

func TestBuilderRevertToIndex(t *testing.T) {
	tests := []struct {
		name     string
		index    int
		expected string
	}{
		{"Hist1", 1, "helloWorld"},
		{"Hist2", 2, "HELLO-WORLD"},
		{"Hist3", 3, "hello-world"},
		{"Hist4", 4, "Hello-World"},
		{"Hist5", 5, "Hello-"},
		{"Hist6", 6, "Hello- John"},
		{"Hist7", 7, ""},
		{"Hist8", 0, "Hello World"},
		{"Hist9", -1, ""},
		{"Hist10", -2, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := New("Hello World"). //Hello World 0
							WithHistory(10).
							ToCamelCase().            //helloWorld 1
							ToKebabCase(true).        //HELLO-WORLD 2
							ToKebabCase(false).       //hello-world 3
							ToTitleCase().            //Hello-World 4
							TrimCharsRight(" World"). //Hello-5
							Append("John", " ")
			if rb.RevertToIndex(tt.index).String() != tt.expected {
				t.Errorf("RevertToIndex(%d) failed, expected %s, got %s",
					tt.index, tt.expected, rb.RevertToIndex(tt.index).String())
			}
		})
	}
}

func TestRevertByFunc(t *testing.T) {
	tests := []struct {
		name     string
		offset   int
		expected string
	}{
		{"Hist1", 6, "helloWorld"},
		{"Hist2", 5, "HELLO-WORLD"},
		{"Hist3", 4, "hello-world"},
		{"Hist4", 3, "Hello-World"},
		{"Hist5", 2, "Hello-"},
		{"Hist6", 1, "Hello- John"},
		{"Hist7", 11, ""},
		{"Hist8", 0, ""},
		{"Hist9", -1, ""},
		{"Hist10", -2, ""},
		{"Hist11", -6, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := New("Hello World"). //Hello World 0
							WithHistory(10).
							ToCamelCase().            //helloWorld 1
							ToKebabCase(true).        //HELLO-WORLD 2
							ToKebabCase(false).       //hello-world 3
							ToTitleCase().            //Hello-World 4
							TrimCharsRight(" World"). //Hello-5
							Append("John", " ")       // 6

			revStr := rb.RevertWithFunction(func(history *StringHistory) int {
				return history.Len() - tt.offset
			}).String()
			if revStr != tt.expected {
				t.Errorf("RevertWithFunction failed, expected %s, got %s", tt.expected, revStr)
			}
		})
	}
}

func TestRevertNoHistory(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello-World 4
					TrimCharsRight(" World"). //Hello-5
					Append("John", " ")       // 6
	str := rb.RevertToPrevious().String()
	err := rb.Error()
	if str != "Hello- John" || !errors.CompareErrors(err, errors.ErrHistoryNotInitialized) {
		t.Errorf("RevertToPrevious failed, expected empty string, got %s", str)
	}
}

func TestRevertNoHistoryIndex(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello-World 4
					TrimCharsRight(" World"). //Hello-5
					Append("John", " ")       // 6

	str := rb.RevertToIndex(1).String()
	err := rb.Error()
	if str != "Hello- John" || !errors.CompareErrors(err, errors.ErrHistoryNotInitialized) {
		t.Errorf("RevertToPrevious failed, expected empty string, got %s", str)
	}
}

func TestRevertNoHistoryFunc(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello-World 4
					TrimCharsRight(" World"). //Hello-5
					Append("John", " ")       // 6

	str := rb.RevertWithFunction(func(history *StringHistory) int {
		return (*rb.GetHistory()).Len() - 2
	}).String()
	err := rb.Error()
	if str != "Hello- John" || !errors.CompareErrors(err, errors.ErrHistoryNotInitialized) {
		t.Errorf("RevertToPrevious failed, expected empty string, got %s", str)
	}
}

var (
	inWord       = "hi"
	compWord     = "Hi"
	ngramLength  = 2
	allWithFatal = New(inWord).
			WithComparisonManager().
			WithHistory(10).
			setError(errors.ErrUnknownError, true)
	allWithNonFatal = New(inWord).
			WithComparisonManager().
			WithHistory(10).
			setError(errors.ErrUnknownError, false)
)

func TestBuilderError(t *testing.T) {
	tests := []struct {
		name     string
		input    *StringBuilder
		expected bool
	}{
		{"Error1", allWithFatal, true},
		{"Error2", allWithNonFatal, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input.ToLower().String() == "" != tt.expected {
				t.Errorf("ToLower failed, expected %t, got %t", tt.expected,
					tt.input.ToLower().String() == "")
			}
			if tt.input.ToUpper().String() == "" != tt.expected {
				t.Errorf("ToUpper failed, expected %t, got %t", tt.expected,
					tt.input.ToUpper().String() == "")
			}
			if tt.input.Capitalize().String() == "" != tt.expected {
				t.Errorf("Capitalize failed, expected %t, got %t", tt.expected,
					tt.input.Capitalize().String() == "")
			}
			if tt.input.Uncapitalize().String() == "" != tt.expected {
				t.Errorf("Uncapitalize failed, expected %t, got %t", tt.expected,
					tt.input.Uncapitalize().String() == "")
			}
			if tt.input.ToTitleCase().String() == "" != tt.expected {
				t.Errorf("ToTitleCase failed, expected %t, got %t", tt.expected,
					tt.input.ToTitleCase().String() == "")
			}
			if tt.input.SplitPascalCase().String() == "" != tt.expected {
				t.Errorf("SplitPascalCase failed, expected %t, got %t", tt.expected,
					tt.input.SplitPascalCase().String() == "")
			}
			if tt.input.SplitCamelCase().String() == "" != tt.expected {
				t.Errorf("SplitCamelCase failed, expected %t, got %t", tt.expected,
					tt.input.SplitCamelCase().String() == "")
			}
			if tt.input.ToSnakeCase(false).String() == "" != tt.expected {
				t.Errorf("ToSnakeCase failed, expected %t, got %t", tt.expected,
					tt.input.ToSnakeCase(false).String() == "")
			}
			if tt.input.ToSnakeCaseWithIgnore(false, "_").String() == "" != tt.expected {
				t.Errorf("ToSnakeCaseWithIgnore failed, expected %t, got %t", tt.expected,
					tt.input.ToSnakeCaseWithIgnore(false, "-").String() == "")
			}
			if tt.input.ToKebabCase(false).String() == "" != tt.expected {
				t.Errorf("ToKebabCase failed, expected %t, got %t", tt.expected,
					tt.input.ToKebabCase(false).String() == "")
			}
			if tt.input.ToCamelCase().String() == "" != tt.expected {
				t.Errorf("ToCamelCase failed, expected %t, got %t", tt.expected,
					tt.input.ToCamelCase().String() == "")
			}
			if tt.input.ToPascalCase().String() == "" != tt.expected {
				t.Errorf("ToPascalCase failed, expected %t, got %t", tt.expected,
					tt.input.ToPascalCase().String() == "")
			}
			if tt.input.ToDelimited('-', "", false).String() == "" != tt.expected {
				t.Errorf("ToDelimited failed, expected %t, got %t", tt.expected,
					tt.input.ToDelimited('-', "", false).String() == "")
			}
			if tt.input.Append("appended", "-").String() == "" != tt.expected {
				t.Errorf("Append failed, expected %t, got %t", tt.expected,
					tt.input.Append("appended", "-").String() == "")
			}
			if tt.input.Prepend("prepended", "-").String() == "" != tt.expected {
				t.Errorf("Prepend failed, expected %t, got %t", tt.expected,
					tt.input.Prepend("prepended", "-").String() == "")
			}
			if tt.input.Trim().String() == "" != tt.expected {
				t.Errorf("Trim failed, expected %t, got %t", tt.expected,
					tt.input.Trim().String() == "")
			}
			if tt.input.TrimLeft().String() == "" != tt.expected {
				t.Errorf("TrimLeft failed, expected %t, got %t", tt.expected,
					tt.input.TrimLeft().String() == "")
			}
			if tt.input.TrimRight().String() == "" != tt.expected {
				t.Errorf("TrimRight failed, expected %t, got %t", tt.expected,
					tt.input.TrimRight().String() == "")
			}
			if tt.input.TrimChars("abc").String() == "" != tt.expected {
				t.Errorf("TrimChars failed, expected %t, got %t", tt.expected,
					tt.input.TrimChars("abc").String() == "")
			}
			if tt.input.TrimCharsLeft("abc").String() == "" != tt.expected {
				t.Errorf("TrimCharsLeft failed, expected %t, got %t", tt.expected,
					tt.input.TrimCharsLeft("abc").String() == "")
			}
			if tt.input.TrimCharsRight("abc").String() == "" != tt.expected {
				t.Errorf("TrimCharsRight failed, expected %t, got %t", tt.expected,
					tt.input.TrimCharsRight("abc").String() == "")
			}
			if tt.input.NormalizeDiacritics().String() == "" != tt.expected {
				t.Errorf("NormalizeDiacritics failed, expected %t, got %t", tt.expected,
					tt.input.NormalizeDiacritics().String() == "")
			}
			wtf := tt.input.Slugify(25).String()
			if wtf == "" != tt.expected {
				t.Errorf("Slugify failed, expected %t, got %t", tt.expected,
					tt.input.Slugify(25).String() == "")
			}
			if tt.input.Truncate(25, "...").String() == "" != tt.expected {
				t.Errorf("Truncate failed, expected %t, got %t", tt.expected,
					tt.input.Truncate(25, "...").String() == "")
			}

			if tt.input.If(rand.Int() > 1, func(s string) string {
				return s + "(If)"
			}).String() == "" != tt.expected {
				t.Errorf("If failed, expected %t, got %t", tt.expected,
					tt.input.If(rand.Int() > 1,
						func(s string) string {
							return s + "(IF)"
						}).String() == "" != tt.expected)
			}

			if tt.input.Transform(func(s string) string {
				return s + "(Transform)"
			}).String() == "" != tt.expected {
				t.Errorf("Transform failed, expected %t, got %t", tt.expected,
					tt.input.Transform(func(s string) string {
						return s + "(Transform)"
					}).String() == "" != tt.expected != tt.expected)
			}

			if tt.input.NormalizeWhitespace(' ').String() == "" != tt.expected {
				t.Errorf("NormalizeWhitespace failed, expected %t, got %t", tt.expected,
					tt.input.NormalizeWhitespace(' ').String() == "")
			}
			if tt.input.NormalizeWhitespaceWithIgnore(' ', "\n\t").String() == "" != tt.expected {
				t.Errorf("NormalizeWhitespace failed, expected %t, got %t", tt.expected,
					tt.input.NormalizeWhitespaceWithIgnore(' ', "\n\t").String() == "")
			}
			if tt.input.CollapseWhitespace().String() == "" != tt.expected {
				t.Errorf("CollapseWhitespace failed, expected %t, got %t", tt.expected,
					tt.input.CollapseWhitespace().String() == "")
			}
			if tt.input.CollapseWhitespaceWithIgnore("\n\t").String() == "" != tt.expected {
				t.Errorf("CollapseWhitespace failed, expected %t, got %t", tt.expected,
					tt.input.CollapseWhitespaceWithIgnore("\n\t").String() == "")
			}
			if tt.input.ReplaceWhitespace("x").String() == "" != tt.expected {
				t.Errorf("ReplaceWhitespace failed, expected %t, got %t", tt.expected,
					tt.input.ReplaceWhitespace("x").String() == "")
			}
			if tt.input.ReplaceWhitespaceWithIgnore("x", "\n\t").String() == "" != tt.expected {
				t.Errorf("ReplaceWhitespace failed, expected %t, got %t", tt.expected,
					tt.input.ReplaceWhitespaceWithIgnore("x", "\n\t").String() == "")
			}
			if tt.input.ReplaceSpaces("x").String() == "" != tt.expected {
				t.Errorf("ReplaceSpaces failed, expected %t, got %t", tt.expected,
					tt.input.ReplaceSpaces("x").String() == "")
			}
			if tt.input.ReplaceNonAlpha("x").String() == "" != tt.expected {
				t.Errorf("ReplaceNonAlpha failed, expected %t, got %t", tt.expected,
					tt.input.ReplaceNonAlpha("x").String() == "")
			}
			if tt.input.ReplaceNonAlphaWithIgnore("x", "\n\t").String() == "" != tt.expected {
				t.Errorf("ReplaceNonAlpha failed, expected %t, got %t", tt.expected,
					tt.input.ReplaceNonAlphaWithIgnore("x", "\n\t").String() == "")
			}
			if tt.input.ReplaceNonAlphaNumeric("x").String() == "" != tt.expected {
				t.Errorf("ReplaceNonAlphaNumeric failed, expected %t, got %t", tt.expected,
					tt.input.ReplaceNonAlphaNumeric("x").String() == "")
			}
			if tt.input.ReplaceNonAlphaNumericWithIgnore("x",
				"\n\t").String() == "" != tt.expected {
				t.Errorf("ReplaceNonAlphaNumeric failed, expected %t, got %t", tt.expected,
					tt.input.ReplaceNonAlphaNumericWithIgnore("x", "\n\t").String() == "")
			}
			if tt.input.NormalizeUnicode(NFC).String() == "" != tt.expected {
				t.Errorf("NormalizeUnicode failed, expected %t, got %t", tt.expected,
					tt.input.NormalizeUnicode(NFC).String() == "")
			}
			if tt.input.RemoveWhitespace().String() == "" != tt.expected {
				t.Errorf("RemoveWhitespace failed, expected %t, got %t", tt.expected,
					tt.input.RemoveWhitespace().String() == "")
			}
			if tt.input.RemoveWhitespaceWithIgnore("abc").String() == "" != tt.expected {
				t.Errorf("RemoveWhitespaceWithIgnore failed, expected %t, got %t", tt.expected,
					tt.input.RemoveWhitespaceWithIgnore("abc").String() == "")
			}
			if tt.input.RemoveNonAlpha(true).String() == "" != tt.expected {
				t.Errorf("RemoveNonAlpha failed, expected %t, got %t", tt.expected,
					tt.input.RemoveNonAlpha(true).String() == "")
			}
			if tt.input.RemoveNonAlphaNumeric(true).String() == "" != tt.expected {
				t.Errorf("RemoveNonAlphaNumeric failed, expected %t, got %t", tt.expected,
					tt.input.RemoveNonAlphaNumeric(true).String() == "")
			}
			if tt.input.RemoveHTML(true).String() == "" != tt.expected {
				t.Errorf("RemoveHTML failed, expected %t, got %t", tt.expected,
					tt.input.RemoveHTML(true).String() == "")
			}
			if tt.input.EscapeHTML().String() == "" != tt.expected {
				t.Errorf("EscapeHTML failed, expected %t, got %t", tt.expected,
					tt.input.EscapeHTML().String() == "")
			}
			if tt.input.SanitizeHTML().String() == "" != tt.expected {
				t.Errorf("SanitizeHTML failed, expected %t, got %t", tt.expected,
					tt.input.SanitizeHTML().String() == "")
			}
			if tt.input.LevenshteinDistance(compWord).String() == "" != tt.expected {
				t.Errorf("LevenshteinDistance failed, expected %t, got %t", tt.expected,
					tt.input.LevenshteinDistance(compWord).String() == "")
			}
			if tt.input.DamerauLevenshteinDistance(compWord).String() == "" != tt.expected {
				t.Errorf("Damarau-LevenshteinDistance failed, expected %t, got %t", tt.expected,
					tt.input.DamerauLevenshteinDistance(compWord).String() == "")
			}
			if tt.input.OSADamerauLevenshteinDistance(compWord).String() == "" != tt.expected {
				t.Errorf("OSA Damarau-LevenshteinDistance failed, expected %t, got %t", tt.expected,
					tt.input.OSADamerauLevenshteinDistance(compWord).String() == "")
			}
			if tt.input.LCS(compWord).String() == "" != tt.expected {
				t.Errorf("LCS failed, expected %t, got %t", tt.expected,
					tt.input.LCS(compWord).String() == "")
			}
			if tt.input.LCSEditDistance(compWord).String() == "" != tt.expected {
				t.Errorf("LCSEditDistance failed, expected %t, got %t", tt.expected,
					tt.input.LCSEditDistance(compWord).String() == "")
			}
			if tt.input.LCSBacktrack(compWord).String() == "" != tt.expected {
				t.Errorf("LCS Backtrack failed, expected %t, got %t", tt.expected,
					tt.input.LCSBacktrack(compWord).String() == "")
			}
			if tt.input.LCSBacktrackAll(compWord).String() == "" != tt.expected {
				t.Errorf("LCSBacktrackAll failed, expected %t, got %t", tt.expected,
					tt.input.LCSBacktrackAll(compWord).String() == "")
			}
			if tt.input.LCSDiff(compWord).String() == "" != tt.expected {
				t.Errorf("LCS Diff failed, expected %t, got %t", tt.expected,
					tt.input.LCSDiff(compWord).String() == "")
			}
			if tt.input.HammingDistance(compWord).String() == "" != tt.expected {
				t.Errorf("HammingDistance failed, expected %t, got %t", tt.expected,
					tt.input.HammingDistance(compWord).String() == "")
			}
			if tt.input.JaroSimilarity(compWord).String() == "" != tt.expected {
				t.Errorf("JaroSimilarity failed, expected %t, got %t", tt.expected,
					tt.input.JaroSimilarity(compWord).String() == "")
			}
			if tt.input.JaroWinklerSimilarity(compWord).String() == "" != tt.expected {
				t.Errorf("JaroWinklerSimilarity failed, expected %t, got %t", tt.expected,
					tt.input.JaroWinklerSimilarity(compWord).String() == "")
			}
			if tt.input.JaccardSimilarity(compWord, ngramLength).String() == "" != tt.expected {
				t.Errorf("JaccardSimilarity failed, expected %t, got %t", tt.expected,
					tt.input.JaccardSimilarity(compWord, ngramLength).String() == "")
			}
			if tt.input.CosineSimilarity(compWord, ngramLength).String() == "" != tt.expected {
				t.Errorf("CosineSimilarity failed, expected %t, got %t", tt.expected,
					tt.input.CosineSimilarity(compWord, ngramLength).String() == "")
			}
			if tt.input.SorensenDiceCoefficient(compWord, ngramLength).String() == "" != tt.expected {
				t.Errorf("SorensenDiceCoefficient failed, expected %t, got %t", tt.expected,
					tt.input.SorensenDiceCoefficient(compWord, ngramLength).String() == "")
			}
			if tt.input.QgramDistance(compWord, ngramLength).String() == "" != tt.expected {
				t.Errorf("QgramDistance failed, expected %t, got %t", tt.expected,
					tt.input.QgramDistance(compWord, ngramLength).String() == "")
			}
			if tt.input.QgramSimilarity(compWord, ngramLength).String() == "" != tt.expected {
				t.Errorf("QgramSimilarity failed, expected %t, got %t", tt.expected,
					tt.input.QgramSimilarity(compWord, ngramLength).String() == "")
			}
			if tt.input.QgramDistanceCustomNgram(map[string]int{"hi": 1},
				"custom").String() == "" != tt.expected {
				t.Errorf("QgramDistanceCustom failed, expected %t, got %t", tt.expected,
					tt.input.QgramDistanceCustomNgram(map[string]int{"hi": 1},
						"custom").String() == "")
			}
			if tt.input.Shingle(ngramLength).String() == "" != tt.expected {
				t.Errorf("Shingle failed, expected %t, got %t", tt.expected,
					tt.input.Shingle(ngramLength).String() == "")
			}
			if tt.input.ShingleSlice(ngramLength).String() == "" != tt.expected {
				t.Errorf("ShingleSlice failed, expected %t, got %t", tt.expected,
					tt.input.ShingleSlice(ngramLength).String() == "")
			}
			if tt.input.Similarity(compWord, Levenshtein).String() == "" != tt.expected {
				t.Errorf("Similarity failed, expected %t, got %t", tt.expected,
					tt.input.Similarity(compWord, Levenshtein).String() == "")
			}
			if tt.input.RemoveNonPrintable().String() == "" != tt.expected {
				t.Errorf("RemoveNonPrintable failed, expected %t, got %t", tt.expected,
					tt.input.RemoveNonPrintable().String() == "")
			}
			if tt.input.RemoveANSIEscapeCodes().String() == "" != tt.expected {
				t.Errorf("RemoveANSIEscapeCodes failed, expected %t, got %t", tt.expected,
					tt.input.RemoveANSIEscapeCodes().String() == "")
			}
			if tt.input.RemovePrefix(".").String() == "" != tt.expected {
				t.Errorf("RemovePrefix failed, expected %t, got %t", tt.expected,
					tt.input.RemovePrefix(".").String() == "")
			}
			if tt.input.RemoveSuffix(".exe").String() == "" != tt.expected {
				t.Errorf("RemoveSuffix failed, expected %t, got %t", tt.expected,
					tt.input.RemoveSuffix(".exe").String() == "")
			}
			prefRes, prefOk := tt.input.RemovePrefixWithResult(".")
			if prefRes.String() == "" != tt.expected || prefOk {
				t.Errorf("RemovePrefixWithResult failed, expected %t, got %t", tt.expected,
					prefRes.String() == "")
			}
			sufRes, sufOk := tt.input.RemoveSuffixWithResult(".exe")
			if sufRes.String() == "" != tt.expected || sufOk {
				t.Errorf("RemoveSuffixWithResult failed, expected %t, got %t", tt.expected,
					sufRes.String() == "")
			}
			if tt.expected == true && tt.input.IsEmail() != false {
				t.Errorf("IsEmail failed, expected %t, got %t", tt.expected,
					tt.input.IsEmail())
			}
			if tt.expected == true && tt.input.IsDomain() != false {
				t.Errorf("IsDomain failed, expected %t, got %t", tt.expected,
					tt.input.IsDomain())
			}
			if tt.expected == true && tt.input.IsURL() != false {
				t.Errorf("IsURL failed, expected %t, got %t", tt.expected,
					tt.input.IsURL())
			}
			if tt.expected == true && tt.input.IsUUID() != false {
				t.Errorf("IsUUID failed, expected %t, got %t", tt.expected,
					tt.input.IsUUID())
			}
			if tt.expected == true && tt.input.IsLengthInRange(1, 100) != false {
				t.Errorf("IsLengthInRange failed, expected %t, got %t", tt.expected,
					tt.input.IsLengthInRange(1, 100))
			}
			if tt.expected == true && tt.input.IsEmpty() != tt.expected {
				t.Errorf("IsEmpty failed, expected %t, got %t", tt.expected,
					tt.input.IsEmpty())
			}
			if tt.expected == true && tt.input.IsEmptyNormalized() != tt.expected {
				t.Errorf("IsEmptyNormalized failed, expected %t, got %t", tt.expected,
					tt.input.IsEmptyNormalized())
			}
			if tt.expected == true && tt.input.IsAlphaNumeric() != false {
				t.Errorf("IsAlphaNumeric failed, expected %t, got %t", false,
					tt.input.IsAlphaNumeric())
			}
			if tt.expected == true && tt.input.IsAlpha() != false {
				t.Errorf("IsAlpha failed, expected %t, got %t", false,
					tt.input.IsAlpha())
			}
			if tt.expected == true && tt.input.IsNormalizedUnicode(NFC) != false {
				t.Errorf("IsNormalizedUnicode failed, expected %t, got %t", false,
					tt.input.IsNormalizedUnicode(NFC))
			}
			if tt.expected == true && tt.input.Contains(compWord) != false {
				t.Errorf("Contains failed, expected %t, got %t", false,
					tt.input.Contains(compWord))
			}
			if tt.expected == true && tt.input.ContainsIgnoreCase(compWord) != false {
				t.Errorf("ContainsIgnoreCase failed, expected %t, got %t", false,
					tt.input.ContainsIgnoreCase(compWord))
			}
			if tt.expected == true && tt.input.ContainsAny([]string{compWord}) != false {
				t.Errorf("ContainsAny failed, expected %t, got %t", false,
					tt.input.ContainsAny([]string{compWord}))
			}
			if tt.expected == true && tt.input.ContainsAnyIgnoreCase([]string{compWord}) != false {
				t.Errorf("ContainsAnyIgnoreCase failed, expected %t, got %t", false,
					tt.input.ContainsAnyIgnoreCase([]string{compWord}))
			}
			if tt.expected == true && tt.input.ContainsAll([]string{compWord}) != false {
				t.Errorf("ContainsAll failed, expected %t, got %t", false,
					tt.input.ContainsAll([]string{compWord}))
			}
			if tt.expected == true && tt.input.ContainsAllIgnoreCase([]string{compWord}) != false {
				t.Errorf("ContainsAllIgnoreCase failed, expected %t, got %t", false,
					tt.input.ContainsAllIgnoreCase([]string{compWord}))
			}
			if tt.expected == true && tt.input.HasPrefix(compWord) != false {
				t.Errorf("HasPrefix failed, expected %t, got %t", false,
					tt.input.HasPrefix(compWord))
			}
			if tt.expected == true && tt.input.HasSuffix(compWord) != false {
				t.Errorf("HasSuffix failed, expected %t, got %t", false,
					tt.input.HasSuffix(compWord))
			}
		})
	}
}

func TestBuilderRequiresEmail(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		passed bool
	}{
		{New("guy@email.com"), true},
		{New("guy"), false},
		{New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireEmail().String() == "" {
			t.Errorf("RequireEmail failed, expected %t, got %t", tt.passed,
				tt.input.RequireEmail().String() == "")
		}
		if !tt.passed && tt.input.RequireEmail().String() != "" {
			t.Errorf("RequireEmail failed, expected %t, got %t", tt.passed,
				tt.input.RequireEmail().String() == "")
		}
	}
}

func TestBuilderRequiresDomain(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		passed bool
	}{
		{New("sub.my.home"), true},
		{New("9081---12"), false},
		{New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireDomain().String() == "" {
			t.Errorf("RequireDomain failed, expected %t, got %t", tt.passed,
				tt.input.RequireDomain().String() == "")
		}
		if !tt.passed && tt.input.RequireDomain().String() != "" {
			t.Errorf("RequireDomain failed, expected %t, got %t", tt.passed,
				tt.input.RequireDomain().String() == "")
		}
	}
}

func TestBuilderRequiresURL(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		passed bool
	}{
		{New("https://www.google.com"), true},
		{New("hello world"), false},
		{New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireURL().String() == "" {
			t.Errorf("RequireURL failed, expected %t, got %t", tt.passed,
				tt.input.RequireURL().String() == "")
		}
		if !tt.passed && tt.input.RequireURL().String() != "" {
			t.Errorf("RequireURL failed, expected %t, got %t", tt.passed,
				tt.input.RequireURL().String() == "")
		}
	}
}

func TestBuilderRequiresUUID(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		passed bool
	}{
		{New("0198a5ea-423d-7ad2-916d-d91b01dad7c7"), true},
		{New("hello world"), false},
		{New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireUUID().String() == "" {
			t.Errorf("RequireUUID failed, expected %t, got %t", tt.passed,
				tt.input.RequireUUID().String() == "")
		}
		if !tt.passed && tt.input.RequireUUID().String() != "" {
			t.Errorf("RequireUUID failed, expected %t, got %t", tt.passed,
				tt.input.RequireUUID().String() == "")
		}
	}
}

func TestBuilderRequiresLength(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		min    int
		max    int
		passed bool
	}{
		{New("hello world"), 1, 100, true},
		{New("hello world"), 25, 255, false},
		{New("guy@email").setError(errors.ErrUnknownError, true), 1, 100, false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireLength(tt.min, tt.max).String() == "" {
			t.Errorf("RequireLength failed, expected %t, got %t", tt.passed,
				tt.input.RequireLength(tt.min, tt.max).String() == "")
		}
		if !tt.passed && tt.input.RequireLength(tt.min, tt.max).String() != "" {
			t.Errorf("RequireLength failed, expected %t, got %t", tt.passed,
				tt.input.RequireLength(tt.min, tt.max).String() == "")
		}
	}
}

func TestRequireNotEmpty(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		passed bool
	}{
		{New("hello world"), true},
		{New(""), false},
		{New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireNotEmpty().String() == "" {
			t.Errorf("RequireNotEmpty failed, expected %t, got %t", tt.passed,
				tt.input.RequireNotEmpty().String() == "")
		}
		if !tt.passed && tt.input.RequireNotEmpty().String() != "" {
			t.Errorf("RequireNotEmpty failed, expected %t, got %t", tt.passed,
				tt.input.RequireNotEmpty().String() == "")
		}
	}
}

func TestRequireNotEmptyNormalized(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		passed bool
	}{
		{New("hello world"), true},
		{New("   "), false},
		{New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireNotEmptyNormalized().String() == "" {
			t.Errorf("RequireNotEmptyNormalized failed, expected %t, got %t", tt.passed,
				tt.input.RequireNotEmptyNormalized().String() == "")
		}
		if !tt.passed && tt.input.RequireNotEmptyNormalized().String() != "" {
			t.Errorf("RequireNotEmptyNormalized failed, expected %t, got %t", tt.passed,
				tt.input.RequireNotEmptyNormalized().String() == "")
		}
	}
}

func TestRequireAlpha(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		passed bool
	}{
		{New("helloworld"), true},
		{New("hello world!"), false},
		{New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireAlpha().String() == "" {
			t.Errorf("RequireAlpha failed, expected %t, got %t", tt.passed,
				tt.input.RequireAlpha().String() == "")
		}
		if !tt.passed && tt.input.RequireAlpha().String() != "" {
			t.Errorf("RequireAlpha failed, expected %t, got %t", tt.passed,
				tt.input.RequireAlpha().String() == "")
		}
	}

}

func TestRequireAlphaNumeric(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		passed bool
	}{
		{New("helloworld123"), true},
		{New("hello world!"), false},
		{New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireAlphaNumeric().String() == "" {
			t.Errorf("RequireAlphaNumeric failed, expected %t, got %t", tt.passed,
				tt.input.RequireAlphaNumeric().String() == "")
		}
		if !tt.passed && tt.input.RequireAlphaNumeric().String() != "" {
			t.Errorf("RequireAlphaNumeric failed, expected %t, got %t", tt.passed,
				tt.input.RequireAlphaNumeric().String() == "")
		}
	}
}

func TestRequireNormalizedUnicode(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		form   NormalizationFormat
		passed bool
	}{
		{New("hello world"), NFC, true},
		{New("w̤𝓲𝔱𝙝 𝔣🇦m̤𝗂𝚕ⓨ ⒜ｎ𝐝 f́r̤ï𝘦⒩𝚍s̤❗❕"), NFD, false},
		{New("guy@email").setError(errors.ErrUnknownError, true), NFC, false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireNormalizedUnicode(tt.form).String() == "" {
			t.Errorf("RequireNormalizedUnicode failed, expected %t, got %t", tt.passed,
				tt.input.RequireNormalizedUnicode(tt.form).String() == "")
		}
		res2 := tt.input.RequireNormalizedUnicode(tt.form).String()
		if !tt.passed && res2 != "" {
			t.Errorf("RequireNormalizedUnicode failed, expected %t, got %t", tt.passed,
				tt.input.RequireNormalizedUnicode(tt.form).String() == "")
		}
	}
}

func TestRequireContains(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		comp   string
		passed bool
	}{
		{New("hello world"), "world", true},
		{New("hello world"), "hi", false},
		{New("hello world"), "HELLO", false},
		{New("hello world"), "", false},
		{New(""), "hello", false},
		{New(""), "", false},
		{New("guy@email").setError(errors.ErrUnknownError, true), "", false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireContains(tt.comp).String() == "" {
			t.Errorf("RequireContains failed, expected %t, got %t", tt.passed,
				tt.input.RequireContains(tt.comp).String() == "")
		}
		if !tt.passed && tt.input.RequireContains(tt.comp).String() != "" {
			t.Errorf("RequireContains failed, expected %t, got %t", tt.passed,
				tt.input.RequireContains(tt.comp).String() == "")
		}
	}
}

func TestRequireContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		comp   string
		passed bool
	}{
		{New("hello world"), "world", true},
		{New("hello world"), "hi", false},
		{New("hello world"), "HELLO", true},
		{New("hello world"), "", false},
		{New(""), "hello", false},
		{New(""), "", false},
		{New("guy@email").setError(errors.ErrUnknownError, true), "", false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireContainsIgnoreCase(tt.comp).String() == "" {
			t.Errorf("RequireContainsIgnoreCase failed, expected %t, got %t", tt.passed,
				tt.input.RequireContains(tt.comp).String() == "")
		}
		if !tt.passed && tt.input.RequireContainsIgnoreCase(tt.comp).String() != "" {
			t.Errorf("RequireContainsIgnoreCase failed, expected %t, got %t", tt.passed,
				tt.input.RequireContainsIgnoreCase(tt.comp).String() == "")
		}
	}
}

func TestRequireContainsAny(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		comp   []string
		passed bool
	}{
		{New("hello world"), []string{"world", "hi"}, true},
		{New("HELLO WORLD"), []string{"hi", "HELLO"}, true},
		{New("hello world"), []string{"HELLO", "hi"}, false},
		{New("hello world"), []string{"", "hi"}, false},
		{New("hello world"), []string{}, false},
		{New(""), []string{"hello", "hi"}, false},
		{New(""), []string{}, false},
		{New("guy@email").setError(errors.ErrUnknownError, true), []string{}, false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireContainsAny(tt.comp).String() == "" {
			t.Errorf("RequireContainsAny failed, expected %t, got %t", tt.passed,
				tt.input.RequireContainsAny(tt.comp).String() == "")
		}
		if !tt.passed && tt.input.RequireContainsAny(tt.comp).String() != "" {
			t.Errorf("RequireContainsAny failed, expected %t, got %t", tt.passed,
				tt.input.RequireContainsAny(tt.comp).String() != "")
		}
	}
}

func TestRequireContainsAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		comp   []string
		passed bool
	}{
		{New("hello world"), []string{"world", "hi"}, true},
		{New("HELLO WORLD"), []string{"hi", "HELLO"}, true},
		{New("hello world"), []string{"HELLO", "hi"}, true},
		{New("hello world"), []string{"", "hi"}, false},
		{New("hello world"), []string{}, false},
		{New(""), []string{"hello", "hi"}, false},
		{New(""), []string{}, false},
		{New("guy@email").setError(errors.ErrUnknownError, true), []string{}, false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireContainsAnyIgnoreCase(tt.comp).String() == "" {
			t.Errorf("RequireContainsAnyIgnoreCase failed, expected %t, got %t", tt.passed,
				tt.input.RequireContainsAnyIgnoreCase(tt.comp).String() == "")
		}
		if !tt.passed && tt.input.RequireContainsAnyIgnoreCase(tt.comp).String() != "" {
			t.Errorf("RequireContainsAnyIgnoreCase failed, expected %t, got %t", tt.passed,
				tt.input.RequireContainsAnyIgnoreCase(tt.comp).String() != "")
		}
	}
}

func TestRequireContainsAll(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		comp   []string
		passed bool
	}{
		{New("hello world"), []string{"world", "hello"}, true},
		{New("HELLO WORLD"), []string{"hi", "HELLO"}, false},
		{New("hello world"), []string{"HELLO", "world"}, false},
		{New("hello world"), []string{"", "hi"}, false},
		{New("hello world"), []string{}, false},
		{New(""), []string{"hello", "hi"}, false},
		{New(""), []string{}, false},
		{New("guy@email").setError(errors.ErrUnknownError, true), []string{}, false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireContainsAll(tt.comp).String() == "" {
			t.Errorf("RequireContainsAll failed, expected %t, got %t", tt.passed,
				tt.input.RequireContainsAll(tt.comp).String() == "")
		}
		if !tt.passed && tt.input.RequireContainsAll(tt.comp).String() != "" {
			t.Errorf("RequireContainsAll failed, expected %t, got %t", tt.passed,
				tt.input.RequireContainsAll(tt.comp).String() != "")
		}
	}
}

func TestStringBuilderRequireContainsAllIgnoreCase(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		comp   []string
		passed bool
	}{
		{New("hello world"), []string{"world", "hello"}, true},
		{New("HELLO WORLD"), []string{"hi", "HELLO"}, false},
		{New("hello world"), []string{"HELLO", "world"}, true},
		{New("hello world"), []string{"", "hi"}, false},
		{New("hello world"), []string{}, false},
		{New(""), []string{"hello", "hi"}, false},
		{New(""), []string{}, false},
		{New("guy@email").setError(errors.ErrUnknownError, true), []string{}, false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireContainsAllIgnoreCase(tt.comp).String() == "" {
			t.Errorf("RequireContainsAllI failed, expected %t, got %t", tt.passed,
				tt.input.RequireContainsAll(tt.comp).String() == "")
		}
		if !tt.passed && tt.input.RequireContainsAllIgnoreCase(tt.comp).String() != "" {
			t.Errorf("RequireContainsAllIgnoreCase failed, expected %t, got %t", tt.passed,
				tt.input.RequireContainsAllIgnoreCase(tt.comp).String() != "")
		}
	}
}

func TestRequireHasPrefix(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		comp   string
		passed bool
	}{
		{New("pre-release"), "pre", true},
		{New("APP_ENV_VAR"), "APP", true},
		{New("hello world"), "", false},
		{New(""), "hello", false},
		{New(""), "", false},
		{New("guy@email").setError(errors.ErrUnknownError, true), "", false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireHasPrefix(tt.comp).String() == "" {
			t.Errorf("RequireHasPrefix failed, expected %t, got %t", tt.passed,
				tt.input.RequireHasPrefix(tt.comp).String() == "")
		}
		if !tt.passed && tt.input.RequireHasPrefix(tt.comp).String() != "" {
			t.Errorf("RequireHasPrefix failed, expected %t, got %t", tt.passed,
				tt.input.RequireHasPrefix(tt.comp).String() == "")
		}
	}
}

func TestRequireHasSuffix(t *testing.T) {
	tests := []struct {
		input  *StringBuilder
		comp   string
		passed bool
	}{
		{New("Hello World!"), "!", true},
		{New("Hello World"), "!", false},
		{New("hello world"), "", false},
		{New(""), "hello", false},
		{New(""), "", false},
		{New("guy@email").setError(errors.ErrUnknownError, true), "", false},
	}
	for _, tt := range tests {
		if tt.passed && tt.input.RequireHasSuffix(tt.comp).String() == "" {
			t.Errorf("RequireHasSuffix failed, expected %t, got %t", tt.passed,
				tt.input.RequireHasSuffix(tt.comp).String() == "")
		}
		if !tt.passed && tt.input.RequireHasSuffix(tt.comp).String() != "" {
			t.Errorf("RequireHasSuffix failed, expected %t, got %t", tt.passed,
				tt.input.RequireHasSuffix(tt.comp).String() == "")
		}
	}
}
