package strutil

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestCleanWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"NoWhitespace", "hello world", "helloworld"},
		{"Whitespace", "hello world ", "helloworld"},
		{"Tabs", "hello\tworld", "helloworld"},
		{"Newlines", "hello\nworld", "helloworld"},
		{"CarriageReturns", "hello\rworld", "helloworld"},
		{"MixedWhitespace", "hello\tworld\n\r", "helloworld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := cleanWhitespace(tt.input)
			result := CleanWhitespace(tt.input)
			builderResult := New(tt.input).CleanWhitespace().String()
			builderError := New(tt.input).CleanWhitespace().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("cleanWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCleanWhitespaceWithIgnore(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		ignore   string
		expected string
	}{
		{"NoWhitespace", "hello world", "", "helloworld"},
		{"WhitespacePreserved", "hello world ", " ", "hello world "},
		{"Tabs", "hello	world", "\n\r", "helloworld"},
		{"TabsPreserved", "hello	world", "\t", "hello	world"},
		{"Newlines", "hello\nworld", "\t ", "helloworld"},
		{"NewlinesPreserved", "hello\nworld", "\n ", `hello
world`},
		{"CarriageReturns", "hello\rworld", " ", "helloworld"},
		{"CarriageReturnsPreserved", "hello\rworld", "\r", "hello\rworld"},
		{"MixedWhitespace", "\rhello   world\f", "\t\n", "helloworld"},
		{"MixedWhitespacePreserved", "hello	 world\n", " \n", "hello world\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := cleanWhitespaceWithIgnore(tt.input, tt.ignore)
			result := CleanWhitespaceWithIgnore(tt.input, tt.ignore)
			builderResult := New(tt.input).CleanWhitespaceWithIgnore(tt.ignore)
			builderString := builderResult.String()
			builderError := builderResult.Error()
			if result != tt.expected || helperResult != tt.expected || builderString != tt.expected || builderError != nil {
				t.Errorf("cleanWhitespaceWithIgnore(%q) = %q / %q / %q; want %q",
					tt.input, helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestNormalizeWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"NoWhitespace", "hello world", "hello world"},
		{"ExtraWhitespace", "hello     world", "hello world"},
		{"Tabs", "hello\t\tworld", "hello world"},
		{"Newlines", "hello\n\nworld", "hello world"},
		{"CarriageReturns", "hello\r\rworld", "hello world"},
		{"MixedWhitespace", "hello  \t\tworld\n\n\r\r", "hello world"},
		{"Empty", "", ""},
		{"LeadingWhitespace", "   hello world", "hello world"},
		{"TrailingWhitespace", "hello world   ", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := normalizeWhitespace(tt.input)
			result := NormalizeWhitespace(tt.input)
			builderResult := New(tt.input).NormalizeWhitespace().String()
			builderError := New(tt.input).NormalizeWhitespace().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("normalizeWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCollapseWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"NoWhitespace", "hello world", "hello world"},
		{"ExtraWhitespace", "hello     world", "hello world"},
		{"Tabs", "hello\t\tworld", "hello world"},
		{"Newlines", "hello\n\nworld", "hello world"},
		{"CarriageReturns", "hello\r\rworld", "hello world"},
		{"MixedWhitespace", "hello  \t\tworld\n\n\r\r", "hello world "},
		{"Empty", "", ""},
		{"LeadingWhitespace", "   hello world", " hello world"},
		{"LeadingAndTrailingWhitespace", "   hello world   ", " hello world "},
		{"TrailingWhitespace", "hello world   ", "hello world "},
		{"LeadingTrailingWhitespaceMixed", " hello   world\n\r\t   ", " hello world "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := collapseWhitespace(tt.input)
			result := CollapseWhitespace(tt.input)
			builderResult := New(tt.input).CollapseWhitespace().String()
			builderError := New(tt.input).CollapseWhitespace().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("collapseWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"ToUpperNormal", "hello world", "HELLO WORLD"},
		{"ToUpperEmptyString", "", ""},
		{"ToUpperSingleChar", "a", "A"},
		{"ToUpperSingleCharUpper", "A", "A"},
		{"ToUpperWithSpecials", "hello world!@#$%^&*()_+", "HELLO WORLD!@#$%^&*()_+"},
		{"ToUpperWithWhitespace", "hello world", "HELLO WORLD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := toUpper(tt.input)
			result := ToUpper(tt.input)
			builderResult := New(tt.input).ToUpper().String()
			builderError := New(tt.input).ToUpper().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("ToUpper(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"ToLowerNormal", "Hello World", "hello world"},
		{"ToLowerEmptyString", "", ""},
		{"ToLowerSingleChar", "A", "a"},
		{"ToLowerSingleCharLower", "a", "a"},
		{"ToLowerWithSpecials", "HELLO WORLD!@#$%^&*()_+", "hello world!@#$%^&*()_+"},
		{"ToLowerWithWhitespace", " HELLO WORLD ", " hello world "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := toLower(tt.input)
			result := ToLower(tt.input)
			builderResult := New(tt.input).ToLower().String()
			builderError := New(tt.input).ToLower().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("ToLower(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestKeepAlphaNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		ws       bool
		expected string
	}{
		{"KeepAlphaNumericRemoveSpace", "hello world 123", false, "helloworld123"},
		{"KeepAlphaNumericPreserveSpace", "hello world 123", true, "hello world 123"},
		{"KeepAlphaNumericNoSpaceSpecialsValid", "hello! world! 123!@#$%^&*()", false, "helloworld123"},
		{"KeepAlphaNumericPreserveSpaceSpecialsValid", "hello! world! 123!@#$%^&*()", true, "hello world 123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := alphaNumeric(tt.input, tt.ws)
			result := KeepAlphaNumeric(tt.input, tt.ws)
			builderResult := New(tt.input).KeepAlphaNumeric(tt.ws).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("KeepAlphaNumeric(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestKeepAlpha(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		ws       bool
		expected string
	}{
		{"KeepAlphaRemoveSpace", "hello world", false, "helloworld"},
		{"KeepAlphaPreserveSpace", "hello world", true, "hello world"},
		{"KeepAlphaRemoveSpaceDigits", "hello world 123", false, "helloworld"},
		{"KeepAlphaPreserveSpaceDigits", "hello world 123", true, "hello world "},
		{"KeepAlphaNoSpaceSpecialsValid", "hello! world! 123!@#$%^&*()", false, "helloworld"},
		{"KeepAlphaPreserveSpaceSpecialsValid", "hello! world! 123!@#$%^&*()", true, "hello world "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := alpha(tt.input, tt.ws)
			result := KeepAlpha(tt.input, tt.ws)
			builderResult := New(tt.input).KeepAlpha(tt.ws).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("KeepAlpha(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReplaceWhitespace(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		replacement string
		expected    string
	}{
		{"ReplaceWhiteSpaceDash", "a b\nc\rd\fe\vf\tg", "-", "a-b-c-d-e-f-g"},
		{"ReplaceWhiteSpaceSpace", "a b\nc\rd\fe\vf\tg", " ", "a b c d e f g"},
		{"ReplaceWhiteSpaceSlash", "a b\nc\rd\fe\vf\tg", "/", "a/b/c/d/e/f/g"},
		{"ReplaceWhiteSpaceText", "a b\nc\rd\fe\vf\tg", ", and ", "a, and b, and c, and d, and e, and f, and g"},
		{"ReplaceWhiteSpaceEmpty", "", " ", ""},
		{"ReplaceWhiteSpaceEmptyReplacement", "", " ", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := replaceWhitespace(tt.input, tt.replacement)
			result := ReplaceWhitespace(tt.input, tt.replacement)
			builderResult := New(tt.input).ReplaceWhitespace(tt.replacement).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ReplaceWhitespace - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestReplaceSpaces(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		replacement string
		expected    string
	}{
		{"ReplaceSpacesDash", "a b\nc\rd\fe\vf\tg", "-", "a-b\nc\rd\fe\vf\tg"},
		{"ReplaceSpacesSpace", "a b\nc\rd\fe\vf\tg", " ", "a b\nc\rd\fe\vf\tg"},
		{"ReplaceSpacesSlash", "a b\nc\rd\fe\vf\tg", "/", "a/b\nc\rd\fe\vf\tg"},
		{"ReplaceSpacesText", "a b\nc\rd\fe\vf\tg", ", and ", "a, and b\nc\rd\fe\vf\tg"},
		{"ReplaceSpacesEmpty", "", " ", ""},
		{"ReplaceSpacesEmptyReplacement", "", " ", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := replaceSpaces(tt.input, tt.replacement)
			result := ReplaceSpaces(tt.input, tt.replacement)
			builderResult := New(tt.input).ReplaceSpaces(tt.replacement).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ReplaceSpaces - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		suffix   string
		expected string
	}{
		{"TruncateNormalNoSuffix", "abcdefghijklmnopqrstuvwxyz", 10, "", "abcdefghij"},
		{"TruncateNormalEllipsis", "abcdefghijklmnopqrstuvwxyz", 10, "...", "abcdefghij..."},
		{"TruncateNoChange", "abcde", 25, "", "abcde"},
		{"TruncateNoChangeNoEllipsis", "abcde", 25, "...", "abcde"},
		{"TruncateEmpty", "", 25, "", ""},
		{"TruncateEmptySuffix", "", 25, "...", ""},
		{"TruncateNegLen", "abcde", -1, "", ""},
		{"TruncateNegLenSuffix", "abcde", -1, "...", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := truncate(tt.input, tt.length, tt.suffix)
			result := Truncate(tt.input, tt.length, tt.suffix)
			builderResult := New(tt.input).Truncate(tt.length, tt.suffix).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("Truncate - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestStripHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"StripHTML1",
			`<a onblur="alert(secret)" href="https://www.google.com">Google</a>`,
			`Google`},
		// New test cases:
		{"StripScriptTag",
			"<p>text</p><script>alert('XSS');</script>",
			"text"},
		{"StripJSInHref",
			`<a href="javascript:alert('XSS')">Click Me</a>`,
			`Click Me`},
		{"StripImgOnError",
			`<img src="image.png" onerror="alert('XSS');" />`,
			``},
		{"StripAllowSafeHTML",
			"<p>This is <b>bold</b> and <em>emphasized</em> text.</p>",
			"This is bold and emphasized text."},
		{"HandlePlainText",
			"This is a string with no HTML.",
			"This is a string with no HTML."},
		{"HandleEmptyInput",
			"",
			""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := stripHTML(tt.input)
			result := StripHTML(tt.input)
			builderResult := New(tt.input).StripHTML().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("StripHTML - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestSanitizeHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"StripHTML1",
			`<a onblur="alert(secret)" href="https://www.google.com">Google</a>`,
			`<a href="https://www.google.com" rel="nofollow">Google</a>`},
		// New test cases:
		{"StripScriptTag",
			"<p>text</p><script>alert('XSS');</script>",
			"<p>text</p>"},
		{"StripJSInHref",
			`<a href="javascript:alert('XSS')">Click Me</a>`,
			`Click Me`},
		{"StripImgOnError",
			`<img src="image.png" onerror="alert('XSS');" />`,
			`<img src="image.png"/>`},
		{"AllowSafeHTML",
			"<p>This is <b>bold</b> and <em>emphasized</em> text.</p>",
			"<p>This is <b>bold</b> and <em>emphasized</em> text.</p>"},
		{"HandlePlainText",
			"This is a string with no HTML.",
			"This is a string with no HTML."},
		{"HandleEmptyInput",
			"",
			""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := sanitizeHTML(tt.input)
			result := SanitizeHTML(tt.input)
			builderResult := New(tt.input).SanitizeHTML().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("SanitizeHTML - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestSanitizeHTMLCustom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		allowed  []string
		expected string
	}{
		{"StripHTML1",
			`<a onblur="alert(secret)" href="https://www.google.com">Google</a>`,
			[]string{"body", "p", "b", "em", "h1", "img"},
			`Google`},
		// New test cases:
		{"StripScriptTag",
			"<div><p>text</p><script>alert('XSS');</script><b>bold</b></div>",
			[]string{"div", "p", "b"},
			"<div><p>text</p><b>bold</b></div>"},
		{"StripJSInHref",
			`<html><body><div><p><h1><a href="javascript:alert('XSS')">Click Me</a></h1></p></div></body></html>`,
			[]string{"body", "div", "p"},
			`<body><div><p>Click Me</p></div></body>`},
		{"StripImgOnError",
			`<div><body><img src="image.png" onerror="alert('XSS');" /></body></div>`,
			[]string{"body", "div"},
			`<div><body></body></div>`},
		{"AllowSafeHTML",
			"<p>This is <b>bold</b> and <em>emphasized</em> text.</p>",
			[]string{"p", "b", "em"},
			"<p>This is <b>bold</b> and <em>emphasized</em> text.</p>"},
		{"HandlePlainText",
			"This is a string with no HTML.",
			[]string{"p", "b", "em"},
			"This is a string with no HTML."},
		{"HandleEmptyInput",
			"",
			[]string{"p", "b", "em"},
			""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := sanitizeHTMLCustom(tt.input, tt.allowed)
			result := SanitizeHTMLCustom(tt.input, tt.allowed)
			builderResult := New(tt.input).SanitizeHTMLCustom(tt.allowed).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("SanitizeHTMLCustom - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"TrimSpace", " abc ", "abc"},
		{"TrimSpaceEmpty", " ", ""},
		{"TrimSpaceEmptyInput", "", ""},
		{"TrimSpaceLeading", " abc", "abc"},
		{"TrimSpaceTrailing", "abc ", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := trim(tt.input)
			result := Trim(tt.input)
			builderResult := New(tt.input).Trim().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("Trim - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestTrimLeft(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"TrimLeftSpace", " abc ", "abc "},
		{"TrimLeftSpaceEmpty", " ", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := trimLeft(tt.input)
			result := TrimLeft(tt.input)
			builderResult := New(tt.input).TrimLeft().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("TrimLeft - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestTrimRight(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"TrimRightSpace", " abc ", " abc"},
		{"TrimRightSpaceEmpty", " ", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := trimRight(tt.input)
			result := TrimRight(tt.input)
			builderResult := New(tt.input).TrimRight().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("TrimRight - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestAlphaReplace(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		replacement string
		expected    string
	}{
		{"AlphaReplace", "1234567890", "x", "xxxxxxxxxx"},
		{"AlphaReplaceEmpty", "Hello World!!!", "-", "Hello-World---"},
		{"AlphaReplaceEmpty", "", "x", ""},
		{"AlphaReplaceEmptyReplacement", "   ", "x", "xxx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := alphaReplace(tt.input, tt.replacement)
			result := AlphaReplace(tt.input, tt.replacement)
			builderResult := New(tt.input).AlphaReplace(tt.replacement).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("AlphaReplace - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestAlphaNumericReplace(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		replacement string
		expected    string
	}{
		{"AlphaNumericReplace1", "abcd1234", "x", "abcd1234"},
		{"AlphaNumericReplace2", "a b c d e f", "-", "a-b-c-d-e-f"},
		{"AlphaNumericReplace3", "a b c d e f", "", "abcdef"},
		{"AlphaNumericBlank", "", "x", ""},
		{"AlphaNumericBlankReplacement", "a b c d e f", "", "abcdef"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := alphaNumericReplace(tt.input, tt.replacement)
			result := AlphaNumericReplace(tt.input, tt.replacement)
			builderResult := New(tt.input).AlphaNumericReplace(tt.replacement).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("AlphaNumericReplace - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		expected string
	}{
		{"Slugify",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ",
			45,
			"lorem-ipsum-dolor-sit-amet-consectetur-adipis"},
		{"SlugifyEmpty", "Sed porttitor orci lacinia ipsum efficitur eleifend.", 21, "sed-porttitor-orci-la"},
		{"SlugifyEmptyReplacement", "Donec risus mauris, facilisis eu egestas sed, convallis a ligula. " +
			"Morbi pharetra placerat dapibus. Praesent vitae nisl viverra, malesuada felis non, tincidunt arcu. " +
			"Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Integer " +
			"nisi justo, ultricies vel enim auctor, viverra tincidunt turpis.",
			125,
			"donec-risus-mauris-facilisis-eu-egestas-sed-convallis-a-ligula-morbi-phar" +
				"etra-placerat-dapibus-praesent-vitae-nisl-viverra-ma"},
		{"SlugifySpecialLeadTrail", "!Hello World!", 21, "hello-world"},
		{"SlugifyMultipleSpaces", "Hello   World   Test", 16, "hello-world-test"},
		{"SlugifySpecialChars", "Hello! @#$%^&*() World", 11, "hello-world"},
		{"SlugifyNonAscii", "Café & Résumé", 12, "cafe-resume"},
		{"SlugifyEmpty", "", 10, ""},
		{"SlugifyZeroLength", "Hello World", 0, ""},
		{"SlugifyNegativeLength", "Hello World", -1, ""},
		{"SlugifyMultipleHyphens", "hello---world", 11, "hello-world"},
		{"SlugifyLeadingHyphens", "---hello-world", 11, "hello-world"},
		{"SlugifyTrailingHyphens", "hello-world---", 11, "hello-world"},
		{"SlugifyCamelCase", "camelCase", 100, "camel-case"},
		{"SlugifyPascalCase", "PascalCase", 100, "pascal-case"},
		{"SlugifySnakeCase", "snake_case", 100, "snake-case"},
		{"SlugifyBlob", "somelongstringwecannotsplit", 100, "somelongstringwecannotsplit"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := slugify(tt.input, tt.length)
			result := Slugify(tt.input, tt.length)
			builderResult := New(tt.input).Slugify(tt.length).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("Slugify - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestEscapeHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"EscapeHTML",
			"<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. </p>",
			"&lt;p&gt;Lorem ipsum dolor sit amet, consectetur adipiscing elit. &lt;/p&gt;"},
		{"EscapeHTMLEmpty", "", ""},
		{"EscapeHTMLSpecialChars", "Hello! @#$%^&*() World", "Hello! @#$%^&amp;*() World"},
		{"EscapeHTMLAll", `"\&<>`, "&quot;\\&amp;&lt;&gt;"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := escapeHTML(tt.input)
			result := EscapeHTML(tt.input)
			builderResult := New(tt.input).EscapeHTML().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("EscapeHTML - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestNormalizeDiacritics(t *testing.T) {
	test := []struct {
		name     string
		input    string
		expected string
	}{
		{"NormDiacriticAcute", "café", "cafe"},
		{"NormDiacriticCedilla", "façade", "facade"},
		{"NormDiacriticUmlaut", "naïve", "naive"},
		{"NormDiacriticTilde", "jalapeño", "jalapeno"},
		{"NormDiacriticCircumGraveAcute", "élevàtor ôperàtor", "elevator operator"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := normalizeDiacritics(tt.input)
			result := NormalizeDiacritics(tt.input)
			builderResult := New(tt.input).NormalizeDiacritics().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("NormalizeDiacritics - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestTrimChars(t *testing.T) {
	test := []struct {
		name     string
		input    string
		chars    string
		expected string
	}{
		{"TrimChars", "Hello World!!!", "!", "Hello World"},
		{"TrimCharsEmpty", "Hello World!!!", "", "Hello World!!!"},
		{"TrimCharsEmptyInput", "", "", ""},
		{"TrimCharsLeading", "Hello World!!!", "H", "ello World!!!"},
		{"TrimCharsTrailing", "Hello World!!!", "!", "Hello World"},
		{"TrimCharsPrefix", "x-Hello World!!!", "x-", "Hello World!!!"},
		{"TrimCharsSuffix", "Hello World-alpha", "-alpha", "Hello World"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := trimChars(tt.input, tt.chars)
			result := TrimChars(tt.input, tt.chars)
			builderResult := New(tt.input).TrimChars(tt.chars).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("TrimChars - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestTrimCharsLeft(t *testing.T) {
	test := []struct {
		name     string
		input    string
		chars    string
		expected string
	}{
		{"TrimCharsL", "Hello World!!!", "!", "Hello World!!!"},
		{"TrimCharsLEmpty", "Hello World!!!", "", "Hello World!!!"},
		{"TrimCharsLEmptyInput", "", "", ""},
		{"TrimCharsLLeading", "Hello World!!!", "H", "ello World!!!"},
		{"TrimCharsLTrailing", "Hello World!!!", "!", "Hello World!!!"},
		{"TrimCharsLPrefix", "x-Hello World!!!", "x-", "Hello World!!!"},
		{"TrimCharsLSuffix", "Hello World-alpha", "-alpha", "Hello World-alpha"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := trimCharsLeft(tt.input, tt.chars)
			result := TrimCharsLeft(tt.input, tt.chars)
			builderResult := New(tt.input).TrimCharsLeft(tt.chars).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("TrimCharsLeft - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestTrimCharsRight(t *testing.T) {
	test := []struct {
		name     string
		input    string
		chars    string
		expected string
	}{
		{"TrimCharsR", "Hello World!!!", "!", "Hello World"},
		{"TrimCharsREmpty", "Hello World!!!", "", "Hello World!!!"},
		{"TrimCharsREmptyInput", "", "", ""},
		{"TrimCharsRLeading", "Hello World!!!", "H", "Hello World!!!"},
		{"TrimCharsRTrailing", "Hello World!!!", "!", "Hello World"},
		{"TrimCharsRPrefix", "x-Hello World!!!", "x-", "x-Hello World!!!"},
		{"TrimCharsRSuffix", "Hello World-alpha", "-alpha", "Hello World"},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := trimCharsRight(tt.input, tt.chars)
			result := TrimCharsRight(tt.input, tt.chars)
			builderResult := New(tt.input).TrimCharsRight(tt.chars).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("TrimCharsRight - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	test := []struct {
		name     string
		input    string
		scream   bool
		expected string
	}{
		{"ToSnakeCase", "Hello World!!!", false, "hello_world!!!"},
		{"ToSnakeCaseEmpty", "  ", false, ""},
		{"ToSnakeCaseNil", "", false, ""},
		{"ToSnakeCaseDiacritic", "Golang Café", false, "golang_cafe"},
		{"ToSnakeCaseKebab", "hello-world", false, "hello_world"},
		{"ToSnakeCaseCamel", "helloWorld", false, "hello_world"},
		{"ToSnakeCasePascal", "HelloWorld", false, "hello_world"},
		{"ToSnakeCaseSnake", "hello_world", false, "hello_world"},
		{"ToSnakeCase", "Hello World!!!", true, "HELLO_WORLD!!!"},
		{"ToSnakeCaseEmpty", "  ", true, ""},
		{"ToSnakeCaseNil", "", true, ""},
		{"ToSnakeCaseDiacritic", "Golang Café", true, "GOLANG_CAFE"},
		{"ToSnakeCaseKebab", "hello-world", true, "HELLO_WORLD"},
		{"ToSnakeCaseCamel", "helloWorld", true, "HELLO_WORLD"},
		{"ToSnakeCasePascal", "HelloWorld", true, "HELLO_WORLD"},
		{"ToSnakeCaseSnake", "hello_world", true, "HELLO_WORLD"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := toSnakeCase(tt.input, tt.scream)
			result := ToSnakeCase(tt.input, tt.scream)
			builderResult := New(tt.input).ToSnakeCase(tt.scream).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ToSnakeCase - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestSplitCamelCase(t *testing.T) {
	test := []struct {
		name     string
		fn       func(string) string
		input    string
		expected string
	}{
		{"SplitCamelCaseCamel", SplitCamelCase, "helloWorld", "hello World"},
		{"SplitCamelCasePascal", SplitCamelCase, "HelloWorld", "Hello World"},
		{"SplitCamelCaseNil", SplitCamelCase, "", ""},
		{"SplitCamelCaseAliasCamel", SplitPascalCase, "helloWorld", "hello World"},
		{"SplitCamelCaseAliasPascal", SplitPascalCase, "HelloWorld", "Hello World"},
		{"SplitCamelCaseAliasNil", SplitPascalCase, "", ""},
		{"SplitCamelCaseLong", SplitCamelCase, "helloWorldHowAreYouToday", "hello World How Are You Today"},
		{"SplitCamelCaseLongAlias", SplitPascalCase, "HelloWorldHowAreYouToday", "Hello World How Are You Today"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := tt.fn(tt.input)
			result := tt.fn(tt.input)
			if helperResult != tt.expected || result != tt.expected {
				t.Errorf("SplitCamelCase - expected %q - got %q / %q", tt.expected, helperResult, result)
			}
		})
	}
}

func TestSplitCamelCaseBuilder(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"SplitCamelCaseCamel", "helloWorld", "hello World"},
		{"SplitCamelCasePascal", "HelloWorld", "Hello World"},
		{"SplitCamelCaseNil", "", ""},
		{"SplitCamelCaseAliasNil", "", ""},
		{"SplitCamelCaseLong", "helloWorldHowAreYouToday", "hello World How Are You Today"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb1 := New(tt.input).SplitCamelCase().String()
			sb2 := New(tt.input).SplitPascalCase().String()
			if sb1 != tt.expected || sb2 != tt.expected {
				t.Errorf("SplitCamelCase - expected %q - got %q / %q", tt.expected, sb1, sb2)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	test := []struct {
		name     string
		input    string
		scream   bool
		expected string
	}{
		{"ToKebabCase", "Hello World!!!", false, "hello-world!!!"},
		{"ToKebabCaseEmpty", "  ", false, ""},
		{"ToKebabCaseNil", "", false, ""},
		{"ToKebabCaseDiacritic", "Golang Café", false, "golang-cafe"},
		{"ToKebabCaseKebab", "hello-world", false, "hello-world"},
		{"ToKebabCaseCamel", "helloWorld", false, "hello-world"},
		{"ToKebabCasePascal", "HelloWorld", false, "hello-world"},
		{"ToKebabCaseSnake", "hello_world", false, "hello-world"},
		{"ToKebabCase", "Hello World!!!", true, "HELLO-WORLD!!!"},
		{"ToKebabCaseEmpty", "  ", true, ""},
		{"ToKebabCaseNil", "", true, ""},
		{"ToKebabCaseDiacritic", "Golang Café", true, "GOLANG-CAFE"},
		{"ToKebabCaseKebab", "hello-world", true, "HELLO-WORLD"},
		{"ToKebabCaseCamel", "helloWorld", true, "HELLO-WORLD"},
		{"ToKebabCasePascal", "HelloWorld", true, "HELLO-WORLD"},
		{"ToKebabCaseSnake", "hello_world", true, "HELLO-WORLD"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := ToKebabCase(tt.input, tt.scream)
			result := ToKebabCase(tt.input, tt.scream)
			builderResult := New(tt.input).ToKebabCase(tt.scream).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ToKebabCase - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestToTitleCase(t *testing.T) {
	test := []struct {
		name     string
		input    string
		expected string
	}{
		{"ToTitleCase", "Hello World!!!", "Hello World!!!"},
		{"ToTitleCaseEmpty", "  ", ""},
		{"ToTitleCaseNil", "", ""},
		{"ToTitleLong",
			"dr. strangelove or how I learned to stop worrying and love the bomb",
			"Dr. Strangelove Or How I Learned To Stop Worrying And Love The Bomb"},
		{"ToTitleCaseCamel", "helloWorld", "Hello World"},
		{"ToTitleCasePascal", "HelloWorld", "Hello World"},
		{"ToTitleCaseSnake", "hello_world", "Hello_world"},
		{"ToTitleCaseKebab", "hello-world", "Hello-World"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := ToTitleCase(tt.input)
			result := ToTitleCase(tt.input)
			builderResult := New(tt.input).ToTitleCase().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ToTitleCase - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	test := []struct {
		name     string
		input    string
		expected string
	}{
		{"Capitalize", "hello world", "Hello world"},
		{"CapitalizeEmpty", "  ", "  "},
		{"CapitalizeNil", "", ""},
		{"CapitalizeCamel", "helloWorld", "HelloWorld"},
		{"CapitalizePascal", "HelloWorld", "HelloWorld"},
		{"CapitalizeSnake", "hello_world", "Hello_world"},
		{"CapitalizeKebab", "hello-world", "Hello-world"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := capitalize(tt.input)
			result := Capitalize(tt.input)
			builderResult := New(tt.input).Capitalize().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("Capitalize - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestUncapitalize(t *testing.T) {
	test := []struct {
		name     string
		input    string
		expected string
	}{
		{"Uncapitalize", "Hello world", "hello world"},
		{"UncapitalizeEmpty", "  ", "  "},
		{"UncapitalizeNil", "", ""},
		{"UncapitalizeCamel", "HelloWorld", "helloWorld"},
		{"UncapitalizePascal", "HelloWorld", "helloWorld"},
		{"UncapitalizeSnake", "Hello_world", "hello_world"},
		{"UncapitalizeKebab", "Hello-world", "hello-world"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := uncapitalize(tt.input)
			result := Uncapitalize(tt.input)
			builderResult := New(tt.input).Uncapitalize().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("Uncapitalize - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestToSnakeCaseWithIgnored(t *testing.T) {
	test := []struct {
		name     string
		input    string
		scream   bool
		ignore   string
		expected string
	}{
		{"ToSnakeCase", "Hello World!!!", false, "", "hello_world!!!"},
		{"ToSnakeCaseEmpty", "  ", false, "", ""},
		{"ToSnakeCaseNil", "", false, "", ""},
		{"ToSnakeCaseDiacritic", "Golang Café.com", false, ".", "golang_cafe.com"},
		{"ToSnakeCaseKebab", "hello-world+2", false, "+", "hello_world+2"},
		{"ToSnakeCaseCamel", "helloWorld", false, "", "hello_world"},
		{"ToSnakeCasePascal", "HelloWorld", false, ".", "hello_world"},
		{"ToSnakeCaseSnake", "hello_world", false, "", "hello_world"},
		{"ToSnakeCase", "Hello World!!!", true, "", "HELLO_WORLD!!!"},
		{"ToSnakeCaseEmpty", "  ", true, "", ""},
		{"ToSnakeCaseNil", "", true, "", ""},
		{"ToSnakeCaseDiacritic", "Golang Café.com", true, ".", "GOLANG_CAFE.COM"},
		{"ToSnakeCaseKebab", "hello-world", true, "-", "HELLO-WORLD"},
		{"ToSnakeCaseCamel", "helloWorld", true, "", "HELLO_WORLD"},
		{"ToSnakeCasePascal", "HelloWorld", true, "", "HELLO_WORLD"},
		{"ToSnakeCaseSnake", "hello_world", true, "", "HELLO_WORLD"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := toSnakeCaseWithIgnore(tt.input, tt.scream, tt.ignore)
			result := ToSnakeCaseWithIgnore(tt.input, tt.scream, tt.ignore)
			builderResult := New(tt.input).ToSnakeCaseWithIgnore(tt.scream, tt.ignore).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ToSnakeCase - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"ToCamelCase", "helloWorld", "helloWorld"},
		{"ToCamelCaseEmpty", "  ", ""},
		{"ToCamelCaseNil", "", ""},
		{"ToCamelCaseKebab", "hello-world", "helloWorld"},
		{"ToCamelCaseSnake", "hello_world", "helloWorld"},
		{"ToCamelCasePascal", "HelloWorld", "helloWorld"},
		{"ToCamelCaseMessy", "I'm-a_messy.String", "imAMessyString"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := ToCamelCase(tt.input)
			result := ToCamelCase(tt.input)
			builderResult := New(tt.input).ToCamelCase().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ToCamelCase - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"ToPascalCase", "helloWorld", "HelloWorld"},
		{"ToPascalCaseEmpty", "  ", ""},
		{"ToPascalCaseNil", "", ""},
		{"ToPascalCaseKebab", "hello-world", "HelloWorld"},
		{"ToPascalCaseSnake", "hello_world", "HelloWorld"},
		{"ToPascalCasePascal", "HelloWorld", "HelloWorld"},
		{"ToPascalCaseMessy", "I'm-a_messy.String", "ImAMessyString"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := ToPascalCase(tt.input)
			result := ToPascalCase(tt.input)
			builderResult := New(tt.input).ToPascalCase().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ToPascalCase - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestToDelimited(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		delim    uint8
		scream   bool
		expected string
	}{
		{"ToDelimCase", "Hello World!!!", '_', false, "hello_world!!!"},
		{"ToDelimCaseEmpty", "  ", ' ', false, ""},
		{"ToDelimCaseNil", "", ' ', false, ""},
		{"ToDelimCaseDiacritic", "Golang Café.com", '.', false, "golang.cafe.com"},
		{"ToDelimCaseKebab", "hello-world+2", '+', false, "hello+world+2"},
		{"ToDelimCaseCamel", "helloWorld", '-', false, "hello-world"},
		{"ToDelimCasePascal", "HelloWorld", '|', false, "hello|world"},
		{"ToDelimCaseSnake", "hello_world", ' ', false, "hello world"},
		{"ToDelimCase", "Hello World!!!", '_', true, "HELLO_WORLD!!!"},
		{"ToDelimCaseEmpty", "  ", ' ', true, ""},
		{"ToDelimCaseNil", "", ' ', true, ""},
		{"ToDelimCaseDiacritic", "Golang Café.com", '.', true, "GOLANG.CAFE.COM"},
		{"ToDelimCaseKebab", "hello-world", '-', true, "HELLO-WORLD"},
		{"ToDelimCaseCamel", "helloWorld", '|', true, "HELLO|WORLD"},
		{"ToDelimCasePascal", "HelloWorld", '*', true, "HELLO*WORLD"},
		{"ToDelimCaseSnake", "hello_world", '&', true, "HELLO&WORLD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := toDelimited(tt.input, tt.delim, "", tt.scream)
			result := ToDelimited(tt.input, tt.delim, "", tt.scream)
			builderResult := New(tt.input).ToDelimited(tt.delim, "", tt.scream).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ToPascalCase - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

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
				*builderResult.Comparison().GetLevenshteinDist() != tt.expected {
				t.Errorf("LevenshteinDistance - expected %d - got %d / %d / %d",
					tt.expected,
					helperResult,
					result,
					*builderResult.comparison.LevenshteinDist)
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
				*builderResult.comparison.GetDamerauLevDist() != tt.expected {
				t.Errorf("Damarau-LevenshteinDistance - expected %d - got %d / %d / %d",
					tt.expected,
					helperResult,
					result,
					*builderResult.comparison.GetDamerauLevDist())
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
				*builderResult.comparison.GetOSADamerauLevDist() != tt.expected {
				t.Errorf("OSALevenshteinDistance - expected %d - got %d / %d / %d",
					tt.expected,
					helperResult,
					result,
					*builderResult.comparison.GetOSADamerauLevDist())
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
			if helperResult != tt.expected || result != tt.expected || *builderResult.comparison.GetLCS() != tt.expected {
				t.Errorf("LCS - expected %d - got %d / %d / %d",
					tt.expected, helperResult, result, *builderResult.comparison.GetLCS())
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
				tt.expected, helperResult, result, *builderResult.comparison.GetLCSBacktrack())
			if helperResult != tt.expected ||
				result != tt.expected ||
				*builderResult.comparison.GetLCSBacktrack() != tt.expected {
				t.Errorf("LCSBacktrack - expected %s - got %s / %s / %s",
					tt.expected, helperResult, result, *builderResult.comparison.GetLCSBacktrack())
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
				!CompareStringSlices(tt.expected, *builderResult.comparison.GetLCSBacktrackAll(), false)) {
				t.Errorf("LCSBacktrackAllA - expected %s - got %s / %s / %s",
					tt.expected, helperResult, result, *builderResult.comparison.GetLCSBacktrackAll())
			}
			if tt.expected == nil && (helperResult != nil ||
				result != nil ||
				*builderResult.comparison.GetLCSBacktrackAll() != nil) {
				t.Errorf("LCSBacktrackAllB - expected %d - got %d / %d / %d",
					len(tt.expected), len(helperResult), len(result), len(*builderResult.comparison.GetLCSBacktrackAll()))
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
					*builderResult.comparison.GetLCSDiff() == nil) {
				t.Errorf("LCSDiff - expected %s - got %s / %s / %s",
					tt.expected,
					helperResult,
					result,
					*builderResult.comparison.GetLCSDiff())
			}
			if tt.expected == nil &&
				(helperResult != nil ||
					result != nil ||
					*builderResult.comparison.GetLCSDiff() != nil) {
				t.Errorf("LCSDiff - expected %s - got %s / %s / %s",
					tt.expected,
					helperResult,
					result,
					*builderResult.comparison.GetLCSDiff())
			}
			if tt.expected != nil && (!CompareStringSlices(tt.expected, helperResult, false) ||
				!CompareStringSlices(tt.expected, result, false) ||
				!CompareStringSlices(tt.expected, *builderResult.comparison.GetLCSDiff(), false)) {
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
				*builderResult.comparison.GetLCSEditDistance() != tt.expected {
				t.Errorf("LCSEditDistance - expected %d - got %d / %d /%d",
					tt.expected,
					helperResul,
					result,
					*builderResult.comparison.GetLCSEditDistance(),
				)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		suffix    string
		separator string
		expected  string
	}{
		{"Append1", "hello", "world", "", "helloworld"},
		{"Append2", "hello", "world", " ", "hello world"},
		{"Append3", "hello", "world", "-", "hello-world"},
		{"Append4", "hello", "world", "_", "hello_world"},
		{"Append5", "hello", "world", ".", "hello.world"},
		{"Append6", "hello world", "alpha", "-", "hello world-alpha"},
		{"Append7", "hello", "", "!", "hello"},
		{"Append8", "hello", "Mr. Bond", ", ", "hello, Mr. Bond"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := appendString(tt.s, tt.suffix, tt.separator)
			result := AppendString(tt.s, tt.suffix, tt.separator)
			builderResult := New(tt.s).Append(tt.suffix, tt.separator).String()
			if helperResult != tt.expected ||
				result != tt.expected ||
				builderResult != tt.expected {
				t.Errorf("AppendString - expected %s - got %s / %s / %s",
					tt.expected,
					helperResult,
					result,
					builderResult,
				)
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		prefix    string
		separator string
		expected  string
	}{
		{"Prepend1", "hello", "world", "", "worldhello"},
		{"Prepend2", "hello", "world", " ", "world hello"},
		{"Prepend3", "hello", "world", "-", "world-hello"},
		{"Prepend4", "hello", "world", "_", "world_hello"},
		{"Prepend5", "hello", "world", ".", "world.hello"},
		{"Prepend6", "hello world", "alpha", "-", "alpha-hello world"},
		{"Prepend7", "hello", "", "!", "hello"},
		{"Prepend8", "hello", "Mr. Bond", ", ", "Mr. Bond, hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := prependString(tt.s, tt.prefix, tt.separator)
			result := PrependString(tt.s, tt.prefix, tt.separator)
			builderResult := New(tt.s).Prepend(tt.prefix, tt.separator).String()
			if helperResult != tt.expected ||
				result != tt.expected ||
				builderResult != tt.expected {
				t.Errorf("PrependString - expected %s - got %s / %s / %s",
					tt.expected,
					helperResult,
					result,
					builderResult,
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
			builderHD := builderResult.Comparison().GetHammingDist()
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
				math.Abs(float64(tt.expected)-float64(*builderResult.Comparison().GetJaroSimilarity())) > float64EqualityThreshold {
				t.Errorf("JaroSimilarity - expected %f - got %f / %f / %f",
					tt.expected,
					helperResult,
					result,
					*builderResult.Comparison().GetJaroSimilarity(),
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
				math.Abs(float64(tt.expected)-float64(*builderResult.Comparison().GetJaroWinklerSim())) > float64EqualityThreshold {
				t.Errorf("JaroWinklerSimilarity - expected %f - got %f / %f / %f",
					tt.expected,
					helperResult,
					result,
					*builderResult.Comparison().GetJaroWinklerSim(),
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
				math.Abs(float64(*tt.expected)-float64(*builderResult.Comparison().GetJaccardSim())) > float64EqualityThreshold) {
				t.Errorf("JaccardSimilarity - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.Comparison().GetJaccardSim(),
				)
			}
			if tt.expected == nil &&
				helperResult != nil &&
				result != nil &&
				builderResult.Comparison().GetJaccardSim() != nil {
				t.Errorf("JaccardSimilarity - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.Comparison().GetJaccardSim(),
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
				math.Abs(float64(*tt.expected)-float64(*builderResult.Comparison().GetCosineSimilarity())) >
					float64EqualityThreshold) {
				t.Errorf("CosineSimilarity - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.Comparison().GetCosineSimilarity(),
				)
			}
			if tt.expected == nil &&
				helperResult != nil &&
				result != nil &&
				builderResult.Comparison().GetCosineSimilarity() != nil {
				t.Errorf("CosineSimilarity - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.Comparison().GetCosineSimilarity(),
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
		{"SorensenDiceCoefficient11", "abc", "", 1, &val10},
		{"SorensenDiceCoefficient12", "", "xyz", 1, &val10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := SorensenDiceCoefficient(tt.input1, tt.input2, tt.splitLen)
			result := SorensenDiceCoefficient(tt.input1, tt.input2, tt.splitLen)
			builderResult := New(tt.input1).SorensenDiceCoefficient(tt.input2, tt.splitLen)
			if tt.expected != nil && (math.Abs(float64(*tt.expected)-float64(*helperResult)) > float64EqualityThreshold ||
				math.Abs(float64(*tt.expected)-float64(*result)) > float64EqualityThreshold ||
				math.Abs(float64(*tt.expected)-float64(*builderResult.Comparison().GetSorensenDiceCo())) >
					float64EqualityThreshold) {
				t.Errorf("SorensenDiceCoefficient - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.Comparison().GetSorensenDiceCo(),
				)
			}
			if tt.expected == nil &&
				helperResult != nil &&
				result != nil &&
				builderResult.Comparison().GetSorensenDiceCo() != nil {
				t.Errorf("SorensenDiceCoefficient - expected %f - got %f / %f / %f",
					*tt.expected,
					*helperResult,
					*result,
					*builderResult.Comparison().GetSorensenDiceCo(),
				)
			}
		})
	}
}
