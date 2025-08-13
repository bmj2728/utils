<div style="text-align: center;">
  <h1>Utils - The Missing Go Utilities Package</h1>
  <img src="assets/utils-gopher_transparent.png" alt="utils Gopher Logo" width="500"/>
</div>


<!-- Quality & Testing -->
[![CodeQL](https://github.com/bmj2728/utils/workflows/CodeQL/badge.svg)](https://github.com/bmj2728/utils/actions?query=workflow%3ACodeQL)
[![Testing](https://github.com/bmj2728/utils/actions/workflows/test.yml/badge.svg)](https://github.com/bmj2728/utils/actions)
[![Linter/Formatter](https://github.com/bmj2728/utils/actions/workflows/ci.yml/badge.svg)](https://github.com/bmj2728/utils/actions)
[![codecov](https://codecov.io/gh/bmj2728/utils/branch/main/graph/badge.svg)](https://codecov.io/gh/bmj2728/utils)

<!-- Code Quality & Docs -->
[![Go Report Card](https://goreportcard.com/badge/github.com/bmj2728/utils)](https://goreportcard.com/report/github.com/bmj2728/utils)
[![Go Reference](https://pkg.go.dev/badge/github.com/bmj2728/utils.svg)](https://pkg.go.dev/github.com/bmj2728/utils)

<!-- Project Info -->
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/bmj2728/utils)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/bmj2728/utils?include_prereleases)](https://github.com/bmj2728/utils/releases)

> **Status**: Ready for v0.1.1 🚀 - Stable API with comprehensive testing

A comprehensive collection of utility packages for Go, designed to fill the gaps in the standard library and provide a consistent, well-tested set of tools for common programming tasks.

## Overview

This project aims to provide a set of utility packages that follow these principles:

- **Simple**: Easy to understand and use
- **Consistent**: Predictable APIs across all packages
- **Well-tested**: High test coverage and robust error handling
- **Performant**: Optimized for speed and memory usage
- **Modular**: Use only what you need

## Installation 📦

### Full Library
```bash
go get github.com/bmj2728/utils
```

### Specific Components
```bash
# String utilities only
go get github.com/bmj2728/utils/pkg/strutil
```
```bash
# Version utilities only  
go get github.com/bmj2728/utils/pkg/version
```

### Import Examples
```go
// Import the full string utilities package
import "github.com/bmj2728/utils/pkg/strutil"
```
```go
// Import specific utilities
import (
    "github.com/bmj2728/utils/pkg/strutil"
    "github.com/bmj2728/utils/pkg/version"
)
```

## Key Features ✨

### 🔤 String Utilities (`strutil`)
The heart of the library - comprehensive string manipulation with dual APIs!

- **🎯 Dual API Design**: Choose your style - functional for simplicity, builder for chaining
- **🧹 Sanitization & Cleaning**: HTML sanitization, whitespace normalization, character filtering
- **📏 String Comparison**: Levenshtein, Jaro-Winkler, LCS, and more algorithms
- **📊 Comparison Manager**: Track and organize multiple comparison results *(optional)*
- **🕐 History Tracking**: Revert transformations with full history *(optional)*
- **🔧 Text Transformation**: Case conversion, slug generation, truncation, padding
- **✅ Validation**: Email, URL, numeric, and custom pattern validation
- **🎲 Generation**: Lorem ipsum text, emails, and placeholder content

### 📊 Version Utilities (`version`)
Build-time version management made simple!

- **📋 Version Information**: Embedded build-time version and build details
- **🔍 Semantic Validation**: Validate and parse semantic versioning format

## Usage Guide 🚀

### 🎭 Choose Your API Style

**Functional API** - Direct function calls for simple operations:
```go
import "github.com/bmj2728/utils/pkg/strutil"

// Simple operations
cleaned := strutil.CleanWhitespace("  hello   world  ")  // "hello world"
slug := strutil.Slugify("Hello World!", 20)              // "hello-world"
isValid := strutil.IsEmail("user@example.com")           // true
```

**Builder API** - Fluent chaining for complex operations:
```go
import "github.com/bmj2728/utils/pkg/strutil"

// Chain multiple operations
result, err := strutil.New("  <div>Hello World!</div>  ").
    CleanWhitespace().
    SanitizeHTML().
    ToLower().
    Slugify(50).
    Result()
// Result: "hello-world"
```

### 🧹 String Cleaning & Sanitization

```go
// Remove dangerous HTML but keep safe tags
userInput := "<script>alert('xss')</script><p>Safe content</p>"
clean := strutil.SanitizeHTML(userInput)  // "<p>Safe content</p>"

// Clean whitespace and normalize
messy := "  \t  hello    world  \n  "
tidy := strutil.CleanWhitespace(messy)  // "hello world"

// Remove non-printable characters
withControl := "hello\x00\x01world"
printable := strutil.RemoveNonPrintable(withControl)  // "helloworld"
```

### 🔧 Text Transformation

```go
// Case conversions
text := "hello_world"
camel := strutil.ToCamelCase(text)     // "helloWorld"
pascal := strutil.ToPascalCase(text)   // "HelloWorld"
kebab := strutil.ToKebabCase(text)     // "hello-world"

// String manipulation
original := "Hello World"
prepended := strutil.Prepend(original, "*********")      // "*********Hello World"
truncated := strutil.Truncate(original, 5, "...")  // "Hello..."
```

### ✅ Validation & Generation

```go
// Validation
strutil.IsEmail("test@example.com")     // true
strutil.IsURL("https://example.com")    // true
strutil.IsNumeric("12345")              // true

// Lorem ipsum generation
sentence := strutil.LoremSentence()     // "Lorem ipsum dolor sit amet."
email := strutil.LoremEmail()           // "lorem@ipsum.com"
paragraph := strutil.LoremParagraph()   // Full paragraph of lorem text
```

### 📊 Advanced: Comparison Manager *(Optional)*

Track multiple string comparison results in one place:

```go
// Create a builder with comparison manager
manager := strutil.New("hello world").
    WithComparisonManager().
    LevenshteinDistance("hello there").      // Distance: 6
    JaroSimilarity("hello there").           // Similarity: 0.79
    LCSLength("hello there").                // LCS: 8
    GetComparisonManager()

// Access individual results
distance := manager.GetComparisonResult(strutil.Levenshtein, "hello there").Distance
similarity := manager.GetSimilarityResult(strutil.Jaro, "hello there").Score
lcsLen := manager.GetLCSResult("hello there").Length

// Get all results for analysis
allComparisons := manager.GetComparisonResultsMap()
allSimilarities := manager.GetSimilarityResultsMap()
```

### 🕐 Advanced: History Tracking *(Optional)*

Track transformations and revert when needed:

```go
// Enable history tracking
result, err := strutil.New("  Hello WORLD!  ").
    WithHistory(10).                    // Track up to 10 transformations
    CleanWhitespace().                  // "Hello WORLD!"
    ToLower().                          // "hello world!"
    ToTitle().                          // "Hello World!"
    Slugify(20).                        // "hello-world"
    Result()

// Access transformation history
history := strutil.New("  Hello WORLD!  ").
    WithHistory(10).
    CleanWhitespace().
    ToLower().
    GetHistory()

fmt.Println(history.GetAll())  // ["  Hello WORLD!  ", "Hello WORLD!", "hello world!"]

// Revert to previous states
reverted, err := strutil.New("  Hello WORLD!  ").
    WithHistory(10).
    CleanWhitespace().
    ToLower().
    RevertToPrevious().  // Back to "Hello WORLD!"
    Result()
```

## Documentation 📚

For complete documentation of all available functions and their usage, please refer to the [Go Reference Documentation](https://pkg.go.dev/github.com/bmj2728/utils) (available shortly after v0.1.0 release).

## Roadmap 🗺️

Exciting utilities coming in future releases:

- 📁 **fileoputils** - Safe file operations leveraging Go 1.24+ features
- 🗄️ **dbutils** - Database utilities and connection management  
- 🍕 **sliceutils** - Advanced slice manipulation and algorithms
- 🔢 **floatutils** - Floating-point utilities and math helpers
- 🐚 **shellutils** - Shell command execution and process management
- 🌐 **netutils** - Network utilities and HTTP helpers

## Acknowledgements 🙏

This project stands on the shoulders of giants! We leverage these excellent open-source libraries:

- [**go-edlib**](https://github.com/hbollon/go-edlib) - String comparison and edit distance algorithms for measuring similarity
- [**bluemonday**](https://github.com/microcosm-cc/bluemonday) - HTML sanitizer for safe HTML cleaning
- [**go-sanitize**](https://github.com/mrz1836/go-sanitize) - Powerful string cleaning and sanitization functions
- [**strcase**](https://github.com/iancoleman/strcase) - Converting strings between different case formats
- [**camelcase**](https://github.com/fatih/camelcase) - Splitting camelCase/PascalCase words into components
- [**lorelai**](https://github.com/UltiRequiem/lorelai) - Versatile lorem ipsum generator for placeholder content
- [**go-diacritics**](https://github.com/Regis24GmbH/go-diacritics) - Lightweight diacritics normalization
- [**stripansi**](https://github.com/acarl005/stripansi) - ANSI escape sequence removal for clean output
- [**google/uuid**](https://github.com/google/uuid) - Robust UUID implementation

## License

This project is licensed under the terms of the LICENSE file included in the repository.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.