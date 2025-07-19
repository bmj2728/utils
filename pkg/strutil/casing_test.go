package strutil

import "testing"

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
