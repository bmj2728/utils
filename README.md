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

> **Status**: v0.3.1 🚀—Stable API with comprehensive testing

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
- **🧭 Indexing**: Build and query multiple index types (byte, rune, normalized UTF-32) *(optional)*
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
strutil.IsNumeric("12345", true)        // true

// Lorem ipsum generation
sentence := strutil.LoremSentence()     // "Dapibus dictum sollicitudin congue dignissim hendrerit massa commodo."
email := strutil.LoremEmail()           // "lorem@ipsum.amet"
paragraph := strutil.LoremParagraph()   // Full paragraph of lorem text
```

### 🧭 Advanced: Indexing (Optional)

Build and query indices across multiple coordinate systems:

```go
import "github.com/bmj2728/utils/pkg/strutil"

// Create indices and add positions for a string and substring
idxMap := strutil.NewIndexMap()
idxMap.Add(strutil.NewStringIndex("Hello, 世界", "世", strutil.Byte, 7))
idxMap.Add(strutil.NewStringIndex("Hello, 世界", "世", strutil.Rune, 7))
idxMap.Add(strutil.NewStringIndex("Hello, 世界", "世", strutil.UTF32Norm, 7))

// Retrieve by type
runeEntries := idxMap.GetByIndexType(strutil.Rune)

// Filter by comparison string
matches := idxMap.GetByComparisonString("Hello, 世界")

// Work with the Indices helper
indices := strutil.NewIndices()
indices.IdxMap.Add(strutil.NewStringIndex("abc", "b", strutil.Byte, 1))
indices.Reset() // clears and prepares for re-indexing after transformations
```

### 🕐 Advanced: History Tracking *(Optional)*

Track transformations and revert when needed:

```go
// Enable history tracking
result, err := strutil.New("  Hello WORLD!  ").
    WithHistory(10).                    // Track history, rotating after the tenth change
    Trim().                             // "Hello WORLD!"
    ToLower().                          // "hello world!"
    ToTitle().                          // "Hello World!"
    Slugify(20).                        // "hello-world"
    Result()

// Access transformation history
history := strutil.New("  Hello WORLD!  ").
    WithHistory(10).
    Trim().
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

For complete documentation of all available functions and their usage, please refer to the [Go Reference Documentation](https://pkg.go.dev/github.com/bmj2728/utils).

## Roadmap 🗺️

Exciting utilities coming in future releases:

- 📁 **fileoputils** - Safe file operations leveraging Go 1.24+ features
- 🗄️ **dbutils** - Database utilities and connection management  
- 🍕 **sliceutils** - Advanced slice manipulation and algorithms
- 🔢 **floatutils** - Floating-point utilities and math helpers
- 🐚 **shellutils** - Shell command execution and process management
- 🌐 **netutils** - Network utilities and HTTP helpers

## Acknowledgements 🙏

This project stands on the shoulders of giants—many thanks to the maintainers of these libraries we use:

- [**bluemonday**](https://github.com/microcosm-cc/bluemonday) - HTML sanitization
- [**go-sanitize**](https://github.com/mrz1836/go-sanitize) - String cleaning and sanitization helpers
- [**strcase**](https://github.com/iancoleman/strcase) - Case conversions
- [**camelcase**](https://github.com/fatih/camelcase) - Word splitting for camel/Pascal case
- [**lorelai**](https://github.com/UltiRequiem/lorelai) - Lorem ipsum generation
- [**go-diacritics**](https://github.com/Regis24GmbH/go-diacritics) - Diacritics normalization
- [**stripansi**](https://github.com/acarl005/stripansi) - ANSI escape removal
- [**google/uuid**](https://github.com/google/uuid) - UUID utilities

## License

This project is licensed under the terms of the LICENSE file included in the repository.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.