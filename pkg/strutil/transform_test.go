package strutil

import (
	"errors"
	"testing"

	"utils/pkg/internal"
)

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
		{"TrimCharsSuffix", "Hello World-alphaRemove", "-alphaRemove", "Hello World"},
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
		{"TrimCharsLSuffix", "Hello World-alphaRemove", "-alphaRemove", "Hello World-alphaRemove"},
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
		{"TrimCharsRSuffix", "Hello World-alphaRemove", "-alphaRemove", "Hello World"},
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
		{"Append6", "hello world", "alphaRemove", "-", "hello world-alphaRemove"},
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
		{"Prepend6", "hello world", "alphaRemove", "-", "alphaRemove-hello world"},
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

func TestNormalizeUnicode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		format   NormalizationFormat
		expected string
	}{
		// NFD: Decomposes é (U+00E9) into e (U+0065) + ́ (U+0301)
		{"NFD_AccentedE", "café", NFD, "cafe\u0301"},
		{"NFD_AccentedE_Resume", "résumé", NFD, "re\u0301sume\u0301"},

		// NFC: Composes e + ́ back into é
		{"NFC_CombiningAccent", "cafe\u0301", NFC, "café"},
		{"NFC_MultipleCombining", "e\u0301\u0302", NFC, "é̂"}, // e + acute + circumflex

		// NFKD: Compatibility decomposition (converts special chars)
		{"NFKD_FractionHalf", "½", NFKD, "1⁄2"},
		{"NFKD_SuperscriptTwo", "x²", NFKD, "x2"},
		{"NFKD_Ligature", "ﬁle", NFKD, "file"},
		{"NFKD_CircledOne", "①", NFKD, "1"},

		// NFKC: Compatibility composition
		{"NFKC_FractionHalf", "½", NFKC, "1⁄2"},
		{"NFKC_WithAccents", "ﬁlé", NFKC, "filé"},

		// Edge cases
		{"NFD_Empty", "", NFD, ""},
		{"NFC_AlreadyComposed", "café", NFC, "café"},               // No change needed
		{"NFD_AlreadyDecomposed", "cafe\u0301", NFD, "cafe\u0301"}, // No change needed

		// Mixed content
		{"NFD_Mixed", "naïve café", NFD, "nai\u0308ve cafe\u0301"},
		{"NFKD_Mixed", "½ café", NFKD, "1⁄2 cafe\u0301"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := normalizeUnicode(tt.input, tt.format)
			result := NormalizeUnicode(tt.input, tt.format)
			builderResult := New(tt.input).NormalizeUnicode(tt.format).String()

			if helperResult != tt.expected {
				t.Errorf("normalizeUnicode() = %q, expected %q", helperResult, tt.expected)
				t.Errorf("  Input bytes: % X", []byte(tt.input))
				t.Errorf("  Result bytes: % X", []byte(helperResult))
				t.Errorf("  Expected bytes: % X", []byte(tt.expected))
			}
			if result != tt.expected {
				t.Errorf("NormalizeUnicode() = %q, expected %q", result, tt.expected)
			}
			if builderResult != tt.expected {
				t.Errorf("Builder.NormalizeUnicode() = %q, expected %q", builderResult, tt.expected)
			}
		})
	}
}

func TestBuilderIf(t *testing.T) {
	tests := []struct {
		name      string
		input     *StringBuilder
		condition bool
		transform func(b string) string
		expected  string
	}{
		{"TransformAnon", New("Hello World"),
			true, func(b string) string { return b + "!" }, "Hello World!"},
		{"TransformLower", New("Hello World"), true, ToLower, "hello world"},
		{"TransformUpper", New("Hello World"), true, ToUpper, "HELLO WORLD"},
		{"TransformTrim", New("   Hello World   "), true, Trim, "Hello World"},
		{"TransformCamel", New("Hello World"), true, ToCamelCase, "helloWorld"},
		{"TransformAnon", New("Hello World"),
			false, func(b string) string { return b + "!" }, "Hello World"},
		{"TransformLower", New("Hello World"), false, ToLower, "Hello World"},
		{"TransformUpper", New("Hello World"), false, ToUpper, "Hello World"},
		{"TransformTrim", New("   Hello World   "),
			false, Trim, "   Hello World   "},
		{"TransformCamel", New("Hello World"), false, ToCamelCase, "Hello World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := tt.input.If(tt.condition, tt.transform).String()
			if out != tt.expected {
				t.Errorf("If() = %q, expected %q", out, tt.expected)
			}
		})
	}
}

func TestBuilderTransform(t *testing.T) {
	tests := []struct {
		name      string
		input     *StringBuilder
		transform func(b string) string
		expected  string
	}{
		{"TransformAnon", New("Hello World"), func(b string) string { return b + "!" }, "Hello World!"},
		{"TransformLower", New("Hello World"), ToLower, "hello world"},
		{"TransformUpper", New("Hello World"), ToUpper, "HELLO WORLD"},
		{"TransformTrim", New("   Hello World   "), Trim, "Hello World"},
		{"TransformCamel", New("Hello World"), ToCamelCase, "helloWorld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := tt.input.Transform(tt.transform).String()
			if out != tt.expected {
				t.Errorf("Transform() = %q, expected %q", out, tt.expected)
			}
		})
	}
}

func TestBuilderSetterValue(t *testing.T) {
	tests := []struct {
		name      string
		input     *StringBuilder
		setString string
		expected  string
	}{
		{"Set", New("Hello World"), "Hello John", "Hello John"},
		{"SetEmpty", New("Hello World"), "", ""},
		{"SetUnempty", New(""), "Hello World", "Hello World"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input.setValue(tt.setString).String() != tt.expected {
				t.Errorf("Set() = %q, expected %q", tt.input.String(), tt.expected)
			}
		})
	}
}

func TestBuilderSetterError(t *testing.T) {
	tests := []struct {
		name        string
		input       *StringBuilder
		setError    error
		expectedErr error
		isFatal     bool
		expected    string
	}{
		{"Set", New("Hello World"), internal.ErrUnknownError, internal.ErrUnknownError, false, "Hello World"},
		{"Set", New("Hello World"), internal.ErrNoSplitLengthSet, internal.ErrNoSplitLengthSet, false, "Hello World"},
		{"Set", New("Hello World"), internal.ErrInvalidEmpty, internal.ErrInvalidEmpty, false, "Hello World"},
		{"Set", New("Hello World"), internal.ErrNoSplitLengthSet, internal.ErrNoSplitLengthSet, true, ""},
		{"Set", New("Hello World"), internal.ErrInvalidEmpty, internal.ErrInvalidEmpty, true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !errors.Is(tt.input.setError(tt.setError, tt.isFatal).Error(), tt.expectedErr) ||
				tt.input.String() != tt.expected {
				t.Errorf("Set() = %q, expected %q", tt.input.String(), tt.expectedErr)
			}

		})
	}
}
