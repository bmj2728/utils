package strutil

import (
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
