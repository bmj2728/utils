package strutil

import (
	"testing"

	"github.com/bmj2728/utils/pkg/internal/errors"
)

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		suffix   string
		expected string
	}{
		{"TruncateNormalNoSuffix", "abcdefghijklmnopqrstuvwxyz",
			10, "", "abcdefghij"},
		{"TruncateNormalEllipsis", "abcdefghijklmnopqrstuvwxyz",
			10, "...", "abcdefghij..."},
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
				t.Errorf("Truncate - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
				t.Errorf("Trim - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
				t.Errorf("TrimLeft - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
				t.Errorf("TrimRight - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
		{"SlugifyEmpty",
			"Sed porttitor orci lacinia ipsum efficitur eleifend.",
			21,
			"sed-porttitor-orci-la"},
		{"SlugifyEmptyReplacement",
			"Donec risus mauris, facilisis eu egestas sed, convallis a ligula. " +
				"Morbi pharetra placerat dapibus. Praesent vitae nisl viverra, " +
				"malesuada felis non, tincidunt arcu. " +
				"Pellentesque habitant morbi tristique senectus et netus et malesuada " +
				"fames ac turpis egestas. Integer " +
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
				t.Errorf("Slugify - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
				t.Errorf("NormalizeDiacritics - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
		{"TrimCharsSuffix", "Hello World-removeNonAlpha", "-removeNonAlpha", "Hello World"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := trimChars(tt.input, tt.chars)
			result := TrimChars(tt.input, tt.chars)
			builderResult := New(tt.input).TrimChars(tt.chars).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("TrimChars - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
		{"TrimCharsLSuffix",
			"Hello World-removeNonAlpha",
			"-removeNonAlpha",
			"Hello World-removeNonAlpha"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := trimCharsLeft(tt.input, tt.chars)
			result := TrimCharsLeft(tt.input, tt.chars)
			builderResult := New(tt.input).TrimCharsLeft(tt.chars).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("TrimCharsLeft - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
		{"TrimCharsRSuffix", "Hello World-removeNonAlpha", "-removeNonAlpha", "Hello World"},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := trimCharsRight(tt.input, tt.chars)
			result := TrimCharsRight(tt.input, tt.chars)
			builderResult := New(tt.input).TrimCharsRight(tt.chars).String()
			if helperResult != tt.expected ||
				result != tt.expected ||
				builderResult != tt.expected {
				t.Errorf("TrimCharsRight - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
		{"Append6", "hello world", "removeNonAlpha", "-", "hello world-removeNonAlpha"},
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
		{"Prepend6", "hello world", "removeNonAlpha", "-", "removeNonAlpha-hello world"},
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
		ws       rune
		expected string
	}{
		{"NoWhitespace", "hello world", ' ', "hello world"},
		{"ExtraWhitespace", "hello     world", ' ', "hello     world"},
		{"Tabs", "hello\t\tworld", ' ', "hello  world"},
		{"Newlines", "hello\n\nworld", ' ', "hello  world"},
		{"CarriageReturns", "hello\r\rworld", ' ', "hello  world"},
		{"MixedWhitespace", "hello  \t\tworld\n\n\r\r", ' ', "hello    world"},
		{"Empty", "", ' ', ""},
		{"LeadingWhitespace", "   hello world", ' ', "hello world"},
		{"TrailingWhitespace", "hello world   ", ' ', "hello world"},
		{"NormalizeToTab",
			"hello world	goodnight\nmoon",
			'\t',
			"hello	world	goodnight	moon"},
		{"NormalizeToNewline",
			"hello world	goodnight\nmoon",
			'\n',
			"hello\nworld\ngoodnight\nmoon"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := normalizeWhitespace(tt.input, tt.ws)
			result := NormalizeWhitespace(tt.input, tt.ws)
			builderResult := New(tt.input).NormalizeWhitespace(tt.ws).String()
			builderError := New(tt.input).NormalizeWhitespace(tt.ws).Error()
			if result != tt.expected ||
				helperResult != tt.expected ||
				builderResult != tt.expected ||
				builderError != nil {
				t.Errorf("normalizeWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
			}
			bad := New(tt.input).NormalizeWhitespace('p')
			if bad.String() != tt.input {
				t.Errorf("normalizeWhitespace(%q) = %q; want %q", tt.input, bad.String(), tt.input)
			}
		})
	}
}

func TestNormalizeWhitespaceWithIgnore(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		ws       rune
		ignore   string
		expected string
	}{
		{"NoWhitespace", "hello world", ' ', "\t\n", "hello world"},
		{"ExtraWhitespace", "hello     world", ' ', "\t\n", "hello     world"},
		{"Tabs", "hello\t\tworld", ' ', "\t\n", "hello\t\tworld"},
		{"Newlines", "hello\n\nworld", ' ', "\t\n", "hello\n\nworld"},
		{"CarriageReturns", "hello\r\rworld", ' ', "\t\n", "hello  world"},
		{"MixedWhitespace",
			"hello  \t\tworld\n\n\r\r",
			' ',
			"\t\n",
			"hello  \t\tworld"},
		{"Empty", "", ' ', "\n\t", ""},
		{"LeadingWhitespace", "   hello world", ' ', "\n\t", "hello world"},
		{"TrailingWhitespace", "hello world   ", ' ', "\n\t", "hello world"},
		{"NormalizeToTab",
			"hello world	goodnight\nmoon",
			'\t',
			"",
			"hello	world	goodnight	moon"},
		{"NormalizeToNewline",
			"hello world	goodnight\nmoon",
			'\n',
			"",
			"hello\nworld\ngoodnight\nmoon"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := normalizeWhitespaceWithIgnore(tt.input, tt.ws, tt.ignore)
			result := NormalizeWhitespaceWithIgnore(tt.input, tt.ws, tt.ignore)
			builderResult := New(tt.input).NormalizeWhitespaceWithIgnore(tt.ws, tt.ignore).String()
			builderError := New(tt.input).NormalizeWhitespaceWithIgnore(tt.ws, tt.ignore).Error()
			if result != tt.expected ||
				helperResult != tt.expected ||
				builderResult != tt.expected ||
				builderError != nil {
				t.Errorf("normalizeWhitespaceWithIgnore(%q) = %q; want %q", tt.input, result, tt.expected)
			}
			bad := New(tt.input).NormalizeWhitespaceWithIgnore('p', "")
			if bad.String() != tt.input {
				t.Errorf("normalizeWhitespaceWithIgnore(%q) = %q; want %q", tt.input, bad.String(), tt.input)
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
		{"Tabs", "hello\t\tworld", "hello	world"},
		{"Newlines", "hello\n\nworld", "hello\nworld"},
		{"CarriageReturns", "hello\r\rworld", "hello\rworld"},
		{"MixedWhitespace", "hello  \t\tworld\n\n\r\r", "hello \tworld\n\r"},
		{"Empty", "", ""},
		{"LeadingWhitespace", "   hello world", " hello world"},
		{"LeadingAndTrailingWhitespace", "   hello world   ", " hello world "},
		{"TrailingWhitespace", "hello world   ", "hello world "},
		{"LeadingTrailingWhitespaceMixed", " hello   world\n\r\t   ", " hello world\n\r\t "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := collapseWhitespace(tt.input)
			result := CollapseWhitespace(tt.input)
			builderResult := New(tt.input).CollapseWhitespace().String()
			builderError := New(tt.input).CollapseWhitespace().Error()
			if result != tt.expected ||
				helperResult != tt.expected ||
				builderResult != tt.expected ||
				builderError != nil {
				t.Errorf("collapseWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCollapseWhitespaceWithIgnore(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		ignore   string
		expected string
	}{
		{"NoWhitespace", "hello world", "\n\t", "hello world"},
		{"ExtraWhitespace", "hello     world", "\n\t", "hello world"},
		{"Tabs", "hello\t\tworld", "\n\t", "hello		world"},
		{"Newlines", "hello\n\nworld", "\n\t", "hello\n\nworld"},
		{"CarriageReturns", "hello\r\rworld", "\n\t", "hello\rworld"},
		{"MixedWhitespace", "hello  \t\tworld\n\n\r\r", "\n\t", "hello \t\tworld\n\n\r"},
		{"Empty", "", "\n\t", ""},
		{"LeadingWhitespace", "   hello world", "\n\t", " hello world"},
		{"LeadingAndTrailingWhitespace", "   hello world   ", "\n\t", " hello world "},
		{"TrailingWhitespace", "hello world   ", "\n\t", "hello world "},
		{"LeadingTrailingWhitespaceMixed",
			" hello   world\n\r\t   ",
			"\n\t",
			" hello world\n\r\t "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := collapseWhitespaceWithIgnore(tt.input, tt.ignore)
			result := CollapseWhitespaceWithIgnore(tt.input, tt.ignore)
			builderResult := New(tt.input).CollapseWhitespaceWithIgnore(tt.ignore).String()
			builderError := New(tt.input).CollapseWhitespaceWithIgnore(tt.ignore).Error()
			if result != tt.expected ||
				helperResult != tt.expected ||
				builderResult != tt.expected ||
				builderError != nil {
				t.Errorf("collapseWhitespaceWithIgnore(%q) = %q; want %q", tt.input, result, tt.expected)
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
		{"ReplaceWhiteSpaceText", "a b\nc\rd\fe\vf\tg",
			", and ", "a, and b, and c, and d, and e, and f, and g"},
		{"ReplaceWhiteSpaceEmpty", "", " ", ""},
		{"ReplaceWhiteSpaceEmptyReplacement", "", " ", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := replaceWhitespace(tt.input, tt.replacement)
			result := ReplaceWhitespace(tt.input, tt.replacement)
			builderResult := New(tt.input).ReplaceWhitespace(tt.replacement).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ReplaceWhitespace - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestReplaceWhitespaceWithIgnore(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		replacement string
		ignore      string
		expected    string
	}{
		{"ReplaceWhiteSpaceDash", "a b\nc\rd\fe\vf\tg",
			"-", "\n\t", "a-b\nc-d-e-f\tg"},
		{"ReplaceWhiteSpaceSpace", "a b\nc\rd\fe\vf\tg",
			" ", "\n\t", "a b\nc d e f\tg"},
		{"ReplaceWhiteSpaceSlash", "a b\nc\rd\fe\vf\tg",
			"/", "\n\t", "a/b\nc/d/e/f\tg"},
		{"ReplaceWhiteSpaceText", "a b\nc\rd\fe\vf\tg",
			", and ", " ", "a b, and c, and d, and e, and f, and g"},
		{"ReplaceWhiteSpaceEmpty", "",
			" ", "\n\t", ""},
		{"ReplaceWhiteSpaceEmptyReplacement", "",
			" ", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := replaceWhitespaceWithIgnore(tt.input, tt.replacement, tt.ignore)
			result := ReplaceWhitespaceWithIgnore(tt.input, tt.replacement, tt.ignore)
			builderResult := New(tt.input).ReplaceWhitespaceWithIgnore(tt.replacement, tt.ignore).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ReplaceWhitespace - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
				t.Errorf("ReplaceSpaces - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestReplaceNonAlpha(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		replacement string
		expected    string
	}{
		{"ReplaceNonAlpha", "1234567890", "x", "xxxxxxxxxx"},
		{"AlphaReplaceEmpty", "Hello World!!!", "-", "Hello-World---"},
		{"AlphaReplaceEmpty", "", "x", ""},
		{"AlphaReplaceEmptyReplacement", "   ", "x", "xxx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := replaceNonAlpha(tt.input, tt.replacement)
			result := ReplaceNonAlpha(tt.input, tt.replacement)
			builderResult := New(tt.input).ReplaceNonAlpha(tt.replacement).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ReplaceNonAlpha - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestReplaceNonAlphaWithIgnore(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		replacement string
		ignore      string
		expected    string
	}{
		{"ReplaceNonAlpha", "1234567890", "x", "3", "xx3xxxxxxx"},
		{"AlphaReplaceEmpty", "Hello World!!!", "-", "!", "Hello-World!!!"},
		{"AlphaReplaceEmpty", "", "x", "hello", ""},
		{"AlphaReplaceEmptyReplacement", "   ", "", "\n", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := replaceNonAlphaWithIgnore(tt.input, tt.replacement, tt.ignore)
			result := ReplaceNonAlphaWithIgnore(tt.input, tt.replacement, tt.ignore)
			builderResult := New(tt.input).ReplaceNonAlphaWithIgnore(tt.replacement, tt.ignore).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ReplaceNonAlpha - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestReplaceNonAlphaNumeric(t *testing.T) {
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
			helperResult := replaceNonAlphaNumeric(tt.input, tt.replacement)
			result := ReplaceNonAlphaNumeric(tt.input, tt.replacement)
			builderResult := New(tt.input).ReplaceNonAlphaNumeric(tt.replacement).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ReplaceNonAlphaNumeric - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
			}
		})
	}
}

func TestReplaceNonAlphaNumericWithIgnore(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		replacement string
		ignore      string
		expected    string
	}{
		{"AlphaNumericReplace1", "abcd1234", "x", "", "abcd1234"},
		{"AlphaNumericReplace2", "a b c/d e f", "-", "/", "a-b-c/d-e-f"},
		{"AlphaNumericReplace3", "a b c+++d e f", "", "+++", "abc+++def"},
		{"AlphaNumericBlank", "", "x", "()+-", ""},
		{"AlphaNumericBlankReplacement", "a b c d e f", "", "", "abcdef"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := replaceNonAlphaNumericWithIgnore(tt.input, tt.replacement, tt.ignore)
			result := ReplaceNonAlphaNumericWithIgnore(tt.input, tt.replacement, tt.ignore)
			builderResult := New(tt.input).ReplaceNonAlphaNumericWithIgnore(tt.replacement, tt.ignore).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("ReplaceNonAlphaNumeric - expected %q - got %q / %q / %q",
					tt.expected, helperResult, result, builderResult)
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
		{"SetError1",
			New("Hello World"),
			errors.ErrUnknownError,
			errors.ErrUnknownError,
			false,
			"Hello World"},
		{"SetError2",
			New("Hello World"),
			errors.ErrNoSplitLengthSet,
			errors.ErrNoSplitLengthSet,
			false,
			"Hello World"},
		{"SetError3",
			New("Hello World"),
			errors.ErrInvalidEmpty,
			errors.ErrInvalidEmpty,
			false,
			"Hello World"},
		{"SetError4",
			New("Hello World"),
			errors.ErrNoSplitLengthSet,
			errors.ErrNoSplitLengthSet,
			true,
			""},
		{"SetError5",
			New("Hello World"),
			errors.ErrInvalidEmpty,
			errors.ErrInvalidEmpty,
			true,
			""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !errors.CompareErrors(tt.input.setError(tt.setError, tt.isFatal).Error(), tt.expectedErr) ||
				tt.input.String() != tt.expected {
				t.Errorf("setError(%v, %t): expected %q/%t", tt.setError, tt.isFatal, tt.expected, tt.isFatal)
				t.Errorf("  Input: %q", tt.input.String())
				t.Errorf("  Expected: %q", tt.expected)
				t.Errorf("  Actual: %q", tt.input.String())
			}
		})
	}
}

func TestRemovePrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		prefix   string
		expected string
	}{
		{"RemovePrefix", "Pre-Release", "Pre-", "Release"},
		{"RemovePrefixNotThere", "Release", "Pre-", "Release"},
		{"RemovePrefixEmpty", "", "Pre-", ""},
		{"RemovePrefixEmptyPrefix", "Pre-", "", "Pre-"},
		{"RemovePrefixEmptyInput", "", "Pre-", ""},
		{"RemovePrefixEmptyBoth", "", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := removePrefix(tt.input, tt.prefix)
			result := RemovePrefix(tt.input, tt.prefix)
			builderResult := New(tt.input).RemovePrefix(tt.prefix).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("RemovePrefix() = %s / %s / %s, expected %s",
					helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestRemoveSuffix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		suffix   string
		expected string
	}{
		{"RemoveSuffix", "v0.1.0-alpha", "-alpha", "v0.1.0"},
		{"RemoveSuffixNotThere", "Release", "-alpha", "Release"},
		{"RemoveSuffixEmptySuffix", "-alpha", "", "-alpha"},
		{"RemoveSuffixEmptyInput", "", "-alpha", ""},
		{"RemoveSuffixEmptyBoth", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := removeSuffix(tt.input, tt.suffix)
			result := RemoveSuffix(tt.input, tt.suffix)
			builderResult := New(tt.input).RemoveSuffix(tt.suffix).String()
			if helperResult != tt.expected || result != tt.expected || builderResult != tt.expected {
				t.Errorf("RemoveSuffix() = %s / %s / %s, expected %s",
					helperResult, result, builderResult, tt.expected)
			}
		})
	}
}

func TestRemovePrefixWithResult(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		prefix string
		out    string
		result bool
	}{
		{"RemovePrefix", "Pre-Release", "Pre-", "Release", true},
		{"RemovePrefixNotThere", "Release", "Pre-", "Release", false},
		{"RemovePrefixEmpty", "", "Pre-", "", false},
		{"RemovePrefixEmptyPrefix", "Pre-", "", "Pre-", false},
		{"RemovePrefixEmptyInput", "", "Pre-", "", false},
		{"RemovePrefixEmptyBoth", "", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperOut, helperResult := removePrefixWithResult(tt.input, tt.prefix)
			out, result := RemovePrefixWithResult(tt.input, tt.prefix)
			builderOut, builderResult := New(tt.input).RemovePrefixWithResult(tt.prefix)

			if helperOut != tt.out || out != tt.out || builderOut.String() != tt.out {
				t.Errorf("RemovePrefixWithResult() = %s / %s / %s / expected %s",
					helperOut, out, builderOut.String(), tt.out)
			}
			if helperResult != tt.result || result != tt.result || builderResult != tt.result {
				t.Errorf("RemovePrefixWithResult() = %t / %t / %t / expected %t",
					helperResult, result, builderResult, tt.result)
			}
		})
	}
}

func TestRemoveSuffixWithResult(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		suffix string
		out    string
		result bool
	}{
		{"RemoveSuffix", "v0.1.0-alpha", "-alpha", "v0.1.0", true},
		{"RemoveSuffixNotThere", "Release", "-alpha", "Release", false},
		{"RemoveSuffixEmptySuffix", "-alpha", "", "-alpha", false},
		{"RemoveSuffixEmptyInput", "", "-alpha", "", false},
		{"RemoveSuffixEmptyBoth", "", "", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperOut, helperResult := removeSuffixWithResult(tt.input, tt.suffix)
			out, result := RemoveSuffixWithResult(tt.input, tt.suffix)
			builderOut, builderResult := New(tt.input).RemoveSuffixWithResult(tt.suffix)
			if helperOut != tt.out || out != tt.out || builderOut.String() != tt.out {
				t.Errorf("RemoveSuffixWithResult() = %s / %s / %s / expected %s",
					helperOut, out, builderOut.String(), tt.out)
			}
			if helperResult != tt.result || result != tt.result || builderResult != tt.result {
				t.Errorf("RemoveSuffixWithResult() = %t / %t / %t / expected %t",
					helperResult, result, builderResult, tt.result)
			}
		})
	}
}
