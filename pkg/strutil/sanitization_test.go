package strutil

import "testing"

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
			helperResult := removeWhitespace(tt.input)
			result := RemoveWhitespace(tt.input)
			builderResult := New(tt.input).RemoveWhitespace().String()
			builderError := New(tt.input).RemoveWhitespace().Error()
			if result != tt.expected || helperResult != tt.expected || builderResult != tt.expected || builderError != nil {
				t.Errorf("removeWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
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
			helperResult := removeWhitespaceWithIgnore(tt.input, tt.ignore)
			result := RemoveWhitespaceWithIgnore(tt.input, tt.ignore)
			builderResult := New(tt.input).RemoveWhitespaceWithIgnore(tt.ignore)
			builderString := builderResult.String()
			builderError := builderResult.Error()
			if result != tt.expected || helperResult != tt.expected || builderString != tt.expected || builderError != nil {
				t.Errorf("removeWhitespaceWithIgnore(%q) = %q / %q / %q; want %q",
					tt.input, helperResult, result, builderResult, tt.expected)
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
			helperResult := removeNonAlphaNumeric(tt.input, tt.ws)
			result := RemoveNonAlphaNumeric(tt.input, tt.ws)
			builderResult := New(tt.input).RemoveNonAlphaNumeric(tt.ws).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("RemoveNonAlphaNumeric(%q) = %q; want %q", tt.input, result, tt.expected)
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
			helperResult := removeNonAlpha(tt.input, tt.ws)
			result := RemoveNonAlpha(tt.input, tt.ws)
			builderResult := New(tt.input).RemoveNonAlpha(tt.ws).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("RemoveNonAlpha(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRemoveHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		preserve bool
		expected string
	}{
		{"StripHTML1",
			`<a onblur="alert(secret)" href="https://www.google.com">Google</a>`,
			false, `Google`},
		// New test cases:
		{"StripScriptTag",
			"<p>text</p><script>alert('XSS');</script>",
			false, "text"},
		{"StripJSInHref",
			`<a href="javascript:alert('XSS')">Click Me</a>`,
			false, `Click Me`},
		{"StripImgOnError",
			`<img src="image.png" onerror="alert('XSS');" />`,
			false, ``},
		{"StripSafeHTML",
			"<p>This is <b>bold</b> and <em>emphasized</em> text.</p>",
			false, "This is bold and emphasized text."},
		{"StripSafeHTML2",
			"<p>This is <b>bold</b> and <em>emphasized</em> text.</p>",
			true, "This is  bold  and  emphasized  text."},
		{"HandlePlainText",
			"This is a string with no HTML.",
			false, "This is a string with no HTML."},
		{"HandleEmptyInput",
			"",
			false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := removeHTML(tt.input, tt.preserve)
			result := RemoveHTML(tt.input, tt.preserve)
			builderResult := New(tt.input).RemoveHTML(tt.preserve).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("RemoveHTML - expected %q - got %q / %q / %q", tt.expected, helperResult, result, builderResult)
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
			"<p>text</p><script>alert('XSS');</script><p>text 2</p>",
			"<p>text</p><p>text 2</p>"},
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
		{"EscapeHTMLAll", `"\&<>`, "&#34;\\&amp;&lt;&gt;"},
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

func TestRemoveNonPrintable(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"NonPrintable1", "hello \rworld", "hello  world"},
		{"NonPrintable2", "hello \f\f\f\f\f\fworld", "hello       world"},
		{"NonPrintable3", "hello \r\f\vworld", "hello    world"},
		{"NonPrintable4", "hello world\v\n", "hello world  "},
		{"NonPrintable5", "hello world", "hello world"},
		{"NonPrintable6", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := removeNonPrintable(tt.input)
			result := RemoveNonPrintable(tt.input)
			builderResult := New(tt.input).RemoveNonPrintable().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("RemoveNonPrintable - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestRemoveANSIEscapeCodes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"ANSI1",
			"hello \x1b[31mworld\x1b[0m", "hello world"},
		{"ANSI2",
			"hello \x1b[31m\x1b[32mworld\x1b[0m", "hello world"},
		{"ANSI3",
			"hello \x1b[31m\x1b[32m\x1b[33mworld\x1b[0m", "hello world"},
		{"ANSI4",
			"hello \x1b[31m\x1b[32m\x1b[33m\x1b[34mworld\x1b[0m", "hello world"},
		{"ANSI5",
			"hello \x1b[31m\x1b[32m\x1b[33m\x1b[34m\x1b[35mworld\x1b[0m", "hello world"},
		{"ANSI6",
			"hello \x1b[31m\x1b[32m\x1b[33m\x1b[34m\x1b[35m\x1b[36mworld\x1b[0m", "hello world"},
		{"ANSIBlank", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := removeANSIEscapeCodes(tt.input)
			result := RemoveANSIEscapeCodes(tt.input)
			builderResult := New(tt.input).RemoveANSIEscapeCodes().String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("RemoveANSIEscapeCodes - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
			}
		})
	}
}
