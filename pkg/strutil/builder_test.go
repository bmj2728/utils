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
			_ = tt.input.LevenshteinDistance(compWord)
			if tt.input.String() == "" != tt.expected {
				t.Errorf("LevenshteinDistance failed, expected %t, got %t", tt.expected,
					tt.input.String() == "")
			}
			_ = tt.input.LevenshteinDistance(compWord)
			if tt.input.String() == "" != tt.expected {
				t.Errorf("Damarau-LevenshteinDistance failed, expected %t, got %t", tt.expected,
					tt.input.String() == "")
			}
			_ = tt.input.OSADamerauLevenshteinDistance(compWord)
			if tt.input.String() == "" != tt.expected {
				t.Errorf("OSA Damarau-LevenshteinDistance failed, expected %t, got %t", tt.expected,
					tt.input.String() == "")
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
			if tt.expected == true && tt.input.IsNumeric(true) != false {
				t.Errorf("IsNumeric failed, expected %t, got %t", false,
					tt.input.IsNumeric(true))
			}
			if tt.expected == true && tt.input.IsNumeric(false) != false {
				t.Errorf("IsNumeric failed, expected %t, got %t", false,
					tt.input.IsNumeric(false))
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
			if tt.input.AddLeftPadding(10).String() == "" != tt.expected {
				t.Errorf("AddLeftPadding failed, expected %t, got %t", tt.expected,
					tt.input.AddLeftPadding(10).String() == "")
			}
			if tt.input.AddRightPadding(10).String() == "" != tt.expected {
				t.Errorf("AddRightPadding failed, expected %t, got %t", tt.expected,
					tt.input.AddRightPadding(10).String() == "")
			}
			if tt.input.AddPadding(10).String() == "" != tt.expected {
				t.Errorf("AddPadding failed, expected %t, got %t", tt.expected,
					tt.input.AddPadding(10).String() == "")
			}
			if tt.input.LeftPadToLength(10).String() == "" != tt.expected {
				t.Errorf("LeftPadToLength failed, expected %t, got %t", tt.expected,
					tt.input.LeftPadToLength(10).String() == "")
			}
			if tt.input.RightPadToLength(10).String() == "" != tt.expected {
				t.Errorf("RightPadToLength failed, expected %t, got %t", tt.expected,
					tt.input.RightPadToLength(10).String() == "")
			}
			if tt.input.PadToLength(10, true).String() == "" != tt.expected {
				t.Errorf("PadToLength failed, expected %t, got %t", tt.expected,
					tt.input.PadToLength(10, true).String() == "")
			}
			if tt.expected == true && tt.input.Index("x", true) != -1 {
				t.Errorf("Index failed expected %t, got %t",
					tt.expected, tt.input.Index("x", true) != -1)
			}
			if tt.expected == true && tt.input.LastIndex("x", false) != -1 {
				t.Errorf("Index failed expected %t, got %t",
					tt.expected, tt.input.Index("x", false) != -1)
			}
			if tt.expected == true && tt.input.IndexAll("x", true) != nil {
				t.Errorf("IndexAll failed expected %t, got %t",
					tt.expected, tt.input.IndexAll("x", true) != nil)
			}
			if tt.expected == true && tt.input.IndexRune('x') != -1 {
				t.Errorf("IndexRune failed expected %t, got %t",
					tt.expected, tt.input.IndexRune('x') != -1)
			}
			if tt.expected == true && tt.input.LastIndexRune('x') != -1 {
				t.Errorf("LastIndexRune failed expected %t, got %t",
					tt.expected, tt.input.LastIndexRune('x') != -1)
			}
			if tt.expected == true && tt.input.IndexAllRune('x') != nil {
				t.Errorf("IndexAllRune failed expected %t, got %t",
					tt.expected, tt.input.IndexAllRune('x') != nil)
			}
			if tt.expected == true && tt.input.UTF32NormalizedIndex("x") != -1 {
				t.Errorf("UTF32NormalizedIndex failed expected %t, got %t",
					tt.expected, tt.input.UTF32NormalizedIndex("x") != -1)
			}
			if tt.expected == true && tt.input.UTF32NormalizedLastIndex("x") != -1 {
				t.Errorf("UTF32NormalizedLastIndex failed expected %t, got %t",
					tt.expected, tt.input.UTF32NormalizedLastIndex("x") != -1)
			}
			if tt.expected == true && tt.input.UTF32NormalizedIndexAll("x") != nil {
				t.Errorf("UTF32NormalizedIndexAll failed expected %t, got %t",
					tt.expected, tt.input.UTF32NormalizedIndexAll("x") != nil)
			}
			if tt.expected == true && tt.input.UTF32NormalizedLastIndex("x") != -1 {
				t.Errorf("UTF32NormalizedLastIndex failed expected %t, got %t",
					tt.expected, tt.input.UTF32NormalizedLastIndex("x") != -1)
			}
			if tt.expected == true && tt.input.UTF32NormalizedIndexAll("x") != nil {
				t.Errorf("UTF32NormalizedIndexAll failed expected %t, got %t",
					tt.expected, tt.input.UTF32NormalizedIndexAll("x") != nil)
			}
			if tt.expected == true && tt.input.FirstAlphaNumericIndex() != -1 {
				t.Errorf("FirstAlphaNumericIndex failed expected %t, got %t",
					tt.expected, tt.input.FirstAlphaNumericIndex() != -1)
			}
			if tt.expected == true && tt.input.LastAlphaNumericIndex() != -1 {
				t.Errorf("LastAlphaNumericIndex failed expected %t, got %t",
					tt.expected, tt.input.LastAlphaNumericIndex() != 1)
			}
			fm, err := tt.input.FuzzyMatch([]string{"x"}, Cosine)
			if tt.expected == true && (fm != "" || err == nil) {
				t.Errorf("FuzzyMatch failed expected %s got %s", "", fm)
			}
			fmt, err := tt.input.FuzzyMatchWithThreshold([]string{"x"}, 0.5, Cosine)
			if tt.expected == true && (fmt != "" || err == nil) {
				t.Errorf("FuzzyMatchWithThreshold failed expected %s got %s", "", fm)
			}
			fms, err := tt.input.FuzzyMatchSet([]string{"x"}, 1, Cosine)
			if tt.expected == true && (fms != nil || err == nil) {
				t.Errorf("FuzzyMatchSet failed expected %s got %s", "", fms)
			}
			fmst, err := tt.input.FuzzyMatchSetWithThreshold([]string{"x"}, 1, 0.5, Cosine)
			if tt.expected == true && (fmst != nil || err == nil) {
				t.Errorf("FuzzyMatchSetWithThreshold failed expected %s got %s", "", fmst)
			}
			if tt.input.Substrings() == nil != tt.expected {
				t.Errorf("Substrings failed, expected %t, got %t", tt.expected,
					tt.input.Substrings() == nil)
			}
			if tt.input.Words() == nil != tt.expected {
				t.Errorf("Words failed, expected %t, got %t", tt.expected,
					tt.input.Words() == nil)
			}
			if tt.input.WordCount() == -1 != tt.expected {
				t.Errorf("WordCount failed, expected %t, got %t", tt.expected,
					tt.input.WordCount() == -1)
			}
			if tt.input.WordMap() == nil != tt.expected {
				t.Errorf("WordMap failed, expected %t, got %t", tt.expected,
					tt.input.WordMap() == nil)
			}
			if tt.input.Split(", ", false) == nil != tt.expected {
				t.Errorf("Split failed, expected %t, got %t", tt.expected,
					tt.input.Split(", ", false) == nil)
			}
			if tt.input.SplitN(", ", 10, false) == nil != tt.expected {
				t.Errorf("SplitN failed, expected %t, got %t", tt.expected,
					tt.input.SplitN(", ", 10, false) == nil)
			}
			if tt.input.Lines() == nil != tt.expected {
				t.Errorf("Lines failed, expected %t, got %t", tt.expected,
					tt.input.Lines() == nil)
			}
			if tt.input.LineCount() == -1 != tt.expected {
				t.Errorf("LineCount failed, expected %t, got %t", tt.expected,
					tt.input.LineCount() == -1)
			}
			if tt.input.Sentences("en") == nil != tt.expected {
				t.Errorf("Sentences failed, expected %t, got %t", tt.expected,
					tt.input.Sentences("en") == nil)
			}
			if tt.input.SentenceCount("en") == -1 != tt.expected {
				t.Errorf("SentenceCount failed, expected %t, got %t", tt.expected,
					tt.input.SentenceCount("en") == -1)
			}
			if tt.input.SplitAt(0) == nil != tt.expected {
				t.Errorf("SplitAt failed, expected %t, got %t", tt.expected,
					tt.input.SplitAt(0) == nil)
			}
			if tt.input.PunktSentenceData() == nil != tt.expected {
				t.Errorf("PunktSentenceData failed, expected %t, got %t", tt.expected,
					tt.input.PunktSentenceData() == nil)
			}
			if tt.input.PunktSentences() == nil != tt.expected {
				t.Errorf("PunktSentences failed, expected %t, got %t", tt.expected,
					tt.input.PunktSentences() == nil)
			}
			if tt.input.PunktSentenceCount() == -1 != tt.expected {
				t.Errorf("PunktSentenceCount failed, expected %t, got %t", tt.expected,
					tt.input.PunktSentenceCount() == -1)
			}
			if tt.input.PunktSentencesCleaned() == nil != tt.expected {
				t.Errorf("PunktSentencesCleaned failed, expected %t, got %t", tt.expected,
					tt.input.PunktSentencesCleaned() == nil)
			}
			if tt.input.PunktSentenceCountCleaned() == -1 != tt.expected {
				t.Errorf("PunktSentenceCountCleaned failed, expected %t, got %t", tt.expected,
					tt.input.PunktSentenceCountCleaned() == -1)
			}
			if tt.input.PunktAnnotatedTokensData() == nil != tt.expected {
				t.Errorf("PunktAnnotatedTokensData failed, expected %t, got %t", tt.expected,
					tt.input.PunktAnnotatedTokensData() == nil)
			}
			if tt.input.PunktAnnotatedTokens() == nil != tt.expected {
				t.Errorf("PunktAnnotatedTokens failed, expected %t, got %t", tt.expected,
					tt.input.PunktAnnotatedTokens() == nil)
			}
			if tt.input.PunktAnnotatedTokenCount() == -1 != tt.expected {
				t.Errorf("PunktAnnotatedTokenCount failed, expected %t, got %t", tt.expected,
					tt.input.PunktAnnotatedTokenCount() == -1)
			}
		})
	}
}

func TestBuilderRequiresEmail(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		passed bool
	}{
		{"ReqEmail1", New("guy@email.com"), true},
		{"ReqEmail2", New("guy"), false},
		{"ReqEmail3", New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireEmail().String() == "" {
				t.Errorf("RequireEmail failed, expected %t, got %t", tt.passed,
					tt.input.RequireEmail().String() == "")
			}
			if !tt.passed && tt.input.RequireEmail().String() != "" {
				t.Errorf("RequireEmail failed, expected %t, got %t", tt.passed,
					tt.input.RequireEmail().String() == "")
			}
		})
	}
}

func TestBuilderRequiresDomain(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		passed bool
	}{
		{"ReqDom1", New("sub.my.home"), true},
		{"ReqDom2", New("9081---12"), false},
		{"ReqDom3", New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.passed && tt.input.RequireDomain().String() == "" {
				t.Errorf("RequireDomain failed, expected %t, got %t", tt.passed,
					tt.input.RequireDomain().String() == "")
			}
			if !tt.passed && tt.input.RequireDomain().String() != "" {
				t.Errorf("RequireDomain failed, expected %t, got %t", tt.passed,
					tt.input.RequireDomain().String() == "")
			}
		})
	}
}

func TestBuilderRequiresURL(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		passed bool
	}{
		{"ReqURL1", New("https://www.google.com"), true},
		{"ReqURL2", New("hello world"), false},
		{"ReqURL3", New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireURL().String() == "" {
				t.Errorf("RequireURL failed, expected %t, got %t", tt.passed,
					tt.input.RequireURL().String() == "")
			}
			if !tt.passed && tt.input.RequireURL().String() != "" {
				t.Errorf("RequireURL failed, expected %t, got %t", tt.passed,
					tt.input.RequireURL().String() == "")
			}
		})
	}
}

func TestBuilderRequiresUUID(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		passed bool
	}{
		{"ReqUUID1", New("0198a5ea-423d-7ad2-916d-d91b01dad7c7"), true},
		{"ReqUUID1", New("hello world"), false},
		{"ReqUUID3", New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireUUID().String() == "" {
				t.Errorf("RequireUUID failed, expected %t, got %t", tt.passed,
					tt.input.RequireUUID().String() == "")
			}
			if !tt.passed && tt.input.RequireUUID().String() != "" {
				t.Errorf("RequireUUID failed, expected %t, got %t", tt.passed,
					tt.input.RequireUUID().String() == "")
			}
		})
	}
}

func TestBuilderRequiresLength(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		min    int
		max    int
		passed bool
	}{
		{"ReqLength1", New("hello world"), 1, 100, true},
		{"ReqLength2", New("hello world"), 25, 255, false},
		{"ReqLength3", New("guy@email").setError(errors.ErrUnknownError, true), 1, 100, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireLength(tt.min, tt.max).String() == "" {
				t.Errorf("RequireLength failed, expected %t, got %t", tt.passed,
					tt.input.RequireLength(tt.min, tt.max).String() == "")
			}
			if !tt.passed && tt.input.RequireLength(tt.min, tt.max).String() != "" {
				t.Errorf("RequireLength failed, expected %t, got %t", tt.passed,
					tt.input.RequireLength(tt.min, tt.max).String() == "")
			}
		})
	}
}

func TestRequireNotEmpty(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		passed bool
	}{
		{"NotEmpty1", New("hello world"), true},
		{"NotEmpty2", New(""), false},
		{"NotEmpty3", New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireNotEmpty().String() == "" {
				t.Errorf("RequireNotEmpty failed, expected %t, got %t", tt.passed,
					tt.input.RequireNotEmpty().String() == "")
			}
			if !tt.passed && tt.input.RequireNotEmpty().String() != "" {
				t.Errorf("RequireNotEmpty failed, expected %t, got %t", tt.passed,
					tt.input.RequireNotEmpty().String() == "")
			}
		})
	}
}

func TestRequireNotEmptyNormalized(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		passed bool
	}{
		{"NotEmptyNorm1", New("hello world"), true},
		{"NotEmptyNorm2", New("   "), false},
		{"NotEmptyNorm3", New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireNotEmptyNormalized().String() == "" {
				t.Errorf("RequireNotEmptyNormalized failed, expected %t, got %t", tt.passed,
					tt.input.RequireNotEmptyNormalized().String() == "")
			}
			if !tt.passed && tt.input.RequireNotEmptyNormalized().String() != "" {
				t.Errorf("RequireNotEmptyNormalized failed, expected %t, got %t", tt.passed,
					tt.input.RequireNotEmptyNormalized().String() == "")
			}
		})
	}
}

func TestRequireAlpha(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		passed bool
	}{
		{"ReqAlpha1", New("helloworld"), true},
		{"ReqAlpha2", New("hello world!"), false},
		{"ReqAlpha3", New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireAlpha().String() == "" {
				t.Errorf("RequireAlpha failed, expected %t, got %t", tt.passed,
					tt.input.RequireAlpha().String() == "")
			}
			if !tt.passed && tt.input.RequireAlpha().String() != "" {
				t.Errorf("RequireAlpha failed, expected %t, got %t", tt.passed,
					tt.input.RequireAlpha().String() == "")
			}
		})
	}
}

func TestRequireAlphaNumeric(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		passed bool
	}{
		{"ReqAlphaNum1", New("helloworld123"), true},
		{"ReqAlphaNum2", New("hello world!"), false},
		{"ReqAlphaNum3", New("guy@email").setError(errors.ErrUnknownError, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireAlphaNumeric().String() == "" {
				t.Errorf("RequireAlphaNumeric failed, expected %t, got %t", tt.passed,
					tt.input.RequireAlphaNumeric().String() == "")
			}
			if !tt.passed && tt.input.RequireAlphaNumeric().String() != "" {
				t.Errorf("RequireAlphaNumeric failed, expected %t, got %t", tt.passed,
					tt.input.RequireAlphaNumeric().String() == "")
			}
		})
	}
}

func TestRequireNumeric(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		strict bool
		passed bool
	}{
		{"ReqNum1", New("123"), true, true},
		{"ReqNum2", New("hello world!"), true, false},
		{"ReqNum3", New("guy@email").setError(errors.ErrUnknownError, true), true, false},
		{"ReqNum4", New("123"), false, true},
		{"ReqNum5", New("hello world!"), false, false},
		{"ReqNum6", New("guy@email").setError(errors.ErrUnknownError, true), false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireNumeric(tt.strict).String() == "" {
				t.Errorf("RequireNumeric failed, expected %t, got %t", tt.passed,
					tt.input.RequireNumeric(tt.strict).String() == "")
			}
			if !tt.passed && tt.input.RequireNumeric(tt.strict).String() != "" {
				t.Errorf("RequireNumeric failed, expected %t, got %t", tt.passed,
					tt.input.RequireNumeric(tt.strict).String() == "")
			}
		})
	}
}

func TestRequireNormalizedUnicode(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		form   NormalizationFormat
		passed bool
	}{
		{"ReqNormUn1", New("hello world"), NFC, true},
		{"ReqNormUn2", New("w̤𝓲𝔱𝙝 𝔣🇦m̤𝗂𝚕ⓨ ⒜ｎ𝐝 f́r̤ï𝘦⒩𝚍s̤❗❕"), NFD, false},
		{"ReqNormUn3", New("guy@email").setError(errors.ErrUnknownError, true), NFC, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireNormalizedUnicode(tt.form).String() == "" {
				t.Errorf("RequireNormalizedUnicode failed, expected %t, got %t", tt.passed,
					tt.input.RequireNormalizedUnicode(tt.form).String() == "")
			}
			res2 := tt.input.RequireNormalizedUnicode(tt.form).String()
			if !tt.passed && res2 != "" {
				t.Errorf("RequireNormalizedUnicode failed, expected %t, got %t", tt.passed,
					tt.input.RequireNormalizedUnicode(tt.form).String() == "")
			}
		})
	}
}

func TestRequireContains(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		comp   string
		passed bool
	}{
		{"ReqCont1", New("hello world"), "world", true},
		{"ReqCont2", New("hello world"), "hi", false},
		{"ReqCont3", New("hello world"), "HELLO", false},
		{"ReqCont4", New("hello world"), "", false},
		{"ReqCont5", New(""), "hello", false},
		{"ReqCont6", New(""), "", false},
		{"ReqCont7", New("guy@email").setError(errors.ErrUnknownError, true), "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireContains(tt.comp).String() == "" {
				t.Errorf("RequireContains failed, expected %t, got %t", tt.passed,
					tt.input.RequireContains(tt.comp).String() == "")
			}
			if !tt.passed && tt.input.RequireContains(tt.comp).String() != "" {
				t.Errorf("RequireContains failed, expected %t, got %t", tt.passed,
					tt.input.RequireContains(tt.comp).String() == "")
			}
		})
	}
}

func TestRequireContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		comp   string
		passed bool
	}{
		{"ReqContIC1", New("hello world"), "world", true},
		{"ReqContIC2", New("hello world"), "hi", false},
		{"ReqContIC3", New("hello world"), "HELLO", true},
		{"ReqContIC4", New("hello world"), "", false},
		{"ReqContIC5", New(""), "hello", false},
		{"ReqContIC6", New(""), "", false},
		{"ReqContIC7", New("guy@email").setError(errors.ErrUnknownError, true), "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireContainsIgnoreCase(tt.comp).String() == "" {
				t.Errorf("RequireContainsIgnoreCase failed, expected %t, got %t", tt.passed,
					tt.input.RequireContains(tt.comp).String() == "")
			}
			if !tt.passed && tt.input.RequireContainsIgnoreCase(tt.comp).String() != "" {
				t.Errorf("RequireContainsIgnoreCase failed, expected %t, got %t", tt.passed,
					tt.input.RequireContainsIgnoreCase(tt.comp).String() == "")
			}
		})
	}
}

func TestRequireContainsAny(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		comp   []string
		passed bool
	}{
		{"ReqContAny1", New("hello world"), []string{"world", "hi"}, true},
		{"ReqContAny2", New("HELLO WORLD"), []string{"hi", "HELLO"}, true},
		{"ReqContAny3", New("hello world"), []string{"HELLO", "hi"}, false},
		{"ReqContAny4", New("hello world"), []string{"", "hi"}, false},
		{"ReqContAny5", New("hello world"), []string{}, false},
		{"ReqContAny6", New(""), []string{"hello", "hi"}, false},
		{"ReqContAny7", New(""), []string{}, false},
		{"ReqContAny8", New("guy@email").setError(errors.ErrUnknownError, true), []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireContainsAny(tt.comp).String() == "" {
				t.Errorf("RequireContainsAny failed, expected %t, got %t", tt.passed,
					tt.input.RequireContainsAny(tt.comp).String() == "")
			}
			if !tt.passed && tt.input.RequireContainsAny(tt.comp).String() != "" {
				t.Errorf("RequireContainsAny failed, expected %t, got %t", tt.passed,
					tt.input.RequireContainsAny(tt.comp).String() != "")
			}
		})
	}
}

func TestRequireContainsAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		comp   []string
		passed bool
	}{
		{"ReqContAnyIC1", New("hello world"), []string{"world", "hi"}, true},
		{"ReqContAnyIC2", New("HELLO WORLD"), []string{"hi", "HELLO"}, true},
		{"ReqContAnyIC3", New("hello world"), []string{"HELLO", "hi"}, true},
		{"ReqContAnyIC4", New("hello world"), []string{"", "hi"}, false},
		{"ReqContAnyIC5", New("hello world"), []string{}, false},
		{"ReqContAnyIC6", New(""), []string{"hello", "hi"}, false},
		{"ReqContAnyIC7", New(""), []string{}, false},
		{"ReqContAnyIC8", New("guy@email").setError(errors.ErrUnknownError, true),
			[]string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireContainsAnyIgnoreCase(tt.comp).String() == "" {
				t.Errorf("RequireContainsAnyIgnoreCase failed, expected %t, got %t", tt.passed,
					tt.input.RequireContainsAnyIgnoreCase(tt.comp).String() == "")
			}
			if !tt.passed && tt.input.RequireContainsAnyIgnoreCase(tt.comp).String() != "" {
				t.Errorf("RequireContainsAnyIgnoreCase failed, expected %t, got %t", tt.passed,
					tt.input.RequireContainsAnyIgnoreCase(tt.comp).String() != "")
			}
		})
	}
}

func TestRequireContainsAll(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		comp   []string
		passed bool
	}{
		{"ReqContAll1", New("hello world"), []string{"world", "hello"}, true},
		{"ReqContAll2", New("HELLO WORLD"), []string{"hi", "HELLO"}, false},
		{"ReqContAll3", New("hello world"), []string{"HELLO", "world"}, false},
		{"ReqContAll4", New("hello world"), []string{"", "hi"}, false},
		{"ReqContAll5", New("hello world"), []string{}, false},
		{"ReqContAll6", New(""), []string{"hello", "hi"}, false},
		{"ReqContAll7", New(""), []string{}, false},
		{"ReqContAll8", New("guy@email").setError(errors.ErrUnknownError, true), []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireContainsAll(tt.comp).String() == "" {
				t.Errorf("RequireContainsAll failed, expected %t, got %t", tt.passed,
					tt.input.RequireContainsAll(tt.comp).String() == "")
			}
			if !tt.passed && tt.input.RequireContainsAll(tt.comp).String() != "" {
				t.Errorf("RequireContainsAll failed, expected %t, got %t", tt.passed,
					tt.input.RequireContainsAll(tt.comp).String() != "")
			}
		})
	}
}

func TestStringBuilderRequireContainsAllIgnoreCase(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		comp   []string
		passed bool
	}{
		{"ReqContAllIC1", New("hello world"), []string{"world", "hello"}, true},
		{"ReqContAllIC2", New("HELLO WORLD"), []string{"hi", "HELLO"}, false},
		{"ReqContAllIC3", New("hello world"), []string{"HELLO", "world"}, true},
		{"ReqContAllIC4", New("hello world"), []string{"", "hi"}, false},
		{"ReqContAllIC5", New("hello world"), []string{}, false},
		{"ReqContAllIC6", New(""), []string{"hello", "hi"}, false},
		{"ReqContAllIC7", New(""), []string{}, false},
		{"ReqContAllIC8", New("guy@email").setError(errors.ErrUnknownError, true), []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireContainsAllIgnoreCase(tt.comp).String() == "" {
				t.Errorf("RequireContainsAllI failed, expected %t, got %t", tt.passed,
					tt.input.RequireContainsAll(tt.comp).String() == "")
			}
			if !tt.passed && tt.input.RequireContainsAllIgnoreCase(tt.comp).String() != "" {
				t.Errorf("RequireContainsAllIgnoreCase failed, expected %t, got %t", tt.passed,
					tt.input.RequireContainsAllIgnoreCase(tt.comp).String() != "")
			}
		})
	}
}

func TestRequireHasPrefix(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		comp   string
		passed bool
	}{
		{"ReqPrefix1", New("pre-release"), "pre", true},
		{"ReqPrefix2", New("APP_ENV_VAR"), "APP", true},
		{"ReqPrefix3", New("hello world"), "", false},
		{"ReqPrefix4", New(""), "hello", false},
		{"ReqPrefix5", New(""), "", false},
		{"ReqPrefix6", New("guy@email").setError(errors.ErrUnknownError, true), "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireHasPrefix(tt.comp).String() == "" {
				t.Errorf("RequireHasPrefix failed, expected %t, got %t", tt.passed,
					tt.input.RequireHasPrefix(tt.comp).String() == "")
			}
			if !tt.passed && tt.input.RequireHasPrefix(tt.comp).String() != "" {
				t.Errorf("RequireHasPrefix failed, expected %t, got %t", tt.passed,
					tt.input.RequireHasPrefix(tt.comp).String() == "")
			}
		})
	}
}

func TestRequireHasSuffix(t *testing.T) {
	tests := []struct {
		name   string
		input  *StringBuilder
		comp   string
		passed bool
	}{
		{"ReqSuffix1", New("Hello World!"), "!", true},
		{"ReqSuffix2", New("Hello World"), "!", false},
		{"ReqSuffix3", New("hello world"), "", false},
		{"ReqSuffix4", New(""), "hello", false},
		{"ReqSuffix5", New(""), "", false},
		{"ReqSuffix6", New("guy@email").setError(errors.ErrUnknownError, true), "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.passed && tt.input.RequireHasSuffix(tt.comp).String() == "" {
				t.Errorf("RequireHasSuffix failed, expected %t, got %t", tt.passed,
					tt.input.RequireHasSuffix(tt.comp).String() == "")
			}
			if !tt.passed && tt.input.RequireHasSuffix(tt.comp).String() != "" {
				t.Errorf("RequireHasSuffix failed, expected %t, got %t", tt.passed,
					tt.input.RequireHasSuffix(tt.comp).String() == "")
			}
		})
	}
}

func TestWithIndices(t *testing.T) {
	tests := []struct {
		name  string
		input *StringBuilder
		isNil bool
	}{
		{"TestWithIndicesWithout", New("hello world"), true},
		{"TestWithIndicesWithLazy", New("hello world").WithIndices(true), true},
		{"TestWithIndicesWithNotLazy", New("hello world").WithIndices(false), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input.indices == nil != tt.isNil {
				t.Errorf("WithIndices failed, expected %t, got %t", tt.isNil,
					tt.input.indices == nil)
			}
		})
	}
}
