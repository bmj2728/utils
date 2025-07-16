# Utils - The Missing Go Utilities Package


<!-- Badges -->
[![CodeQL](https://github.com/bmj2728/utils/workflows/CodeQL/badge.svg)](https://github.com/bmj2728/utils/actions?query=workflow%3ACodeQL)
[![Linter/Formatter](https://github.com/bmj2728/utils/actions/workflows/ci.yml/badge.svg)](https://github.com/bmj2728/utils/actions)
[![Testing](https://github.com/bmj2728/utils/actions/workflows/test.yml/badge.svg)](https://github.com/bmj2728/utils/actions)
[![codecov](https://codecov.io/gh/bmj2728/utils/branch/main/graph/badge.svg)](https://codecov.io/gh/bmj2728/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/bmj2728/utils)](https://goreportcard.com/report/github.com/bmj2728/utils)
[![Go Reference](https://pkg.go.dev/badge/github.com/bmj2728/utils.svg)](https://pkg.go.dev/github.com/bmj2728/utils)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/bmj2728/utils)](https://golang.org/)

> **Status**: Pre-release (approaching v0.1.0) - API may change

A comprehensive collection of utility packages for Go, designed to fill the gaps in the standard library and provide a consistent, well-tested set of tools for common programming tasks.

## Overview

This project aims to provide a set of utility packages that follow these principles:

- **Simple**: Easy to understand and use
- **Consistent**: Predictable APIs across all packages
- **Well-tested**: High test coverage and robust error handling
- **Performant**: Optimized for speed and memory usage
- **Modular**: Use only what you need

## Package Structure

```
utils/
├── strutil/         # String manipulation & validation (currently in development)
├── httputil/        # HTTP client patterns
├── fileutl/         # File operations
├── jsonutil/        # JSON utilities
├── cryptoutil/      # Common crypto patterns
├── configutil/      # Configuration management
├── sliceutil/       # Generic slice operations
├── validationutil/  # Input validation
├── errorutil/       # Error handling
└── testutil/        # Test helpers
```

## Recent Updates

The project has recently added several new features:

- **HTML Sanitization**: Added functions for sanitizing HTML content with various security levels
- **Lorem Ipsum Generation**: Added comprehensive lorem ipsum generation for text, emails, URLs, and more


## Current Implementation

### String Utilities (`strutil`)

The `strutil` package provides comprehensive string manipulation, validation, and sanitization functions. It offers both a functional API and a fluent builder pattern.

#### Error Constants

The package provides standardized error constants for validation failures:

```
// Error constants for validation
ErrInvalidEmail                   = "invalid email address"
ErrInvalidURL                     = "invalid URL"
ErrInvalidUUID                    = "invalid UUID"
ErrInvalidLengthRange             = "invalid length range"
ErrInvalidLength                  = "invalid length"
ErrInvalidEmpty                   = "empty string"
ErrInvalidEmptyAfterNormalization = "empty string after whitespace normalization"
ErrInvalidNotAlphaNumeric         = "string contains non-alphanumeric characters"
```

These constants are used throughout the package for consistent error messaging and can be checked when using the builder API's validation methods.

#### Functional API

The functional API provides standalone functions for string operations:

```
// Generate a UUID
strutil.GenerateUUID()

// Validate an email
strutil.IsEmail(email)

// Create a URL-friendly slug
strutil.Slugify(rawText, 50)

// Keep only alphabetic characters
strutil.KeepAlpha(input, false)

// Keep only alphanumeric characters
strutil.KeepAlphaNumeric(input, true)

// HTML sanitization
strutil.StripHTML(input)
strutil.SanitizeHTML(input)

// Lorem ipsum generation
strutil.LoremWord()
strutil.LoremSentence()
strutil.LoremParagraph()
strutil.LoremEmail()
strutil.LoremURL()
```

#### StringBuilder API

The string builder API allows for chaining multiple operations:

```
// Chain multiple operations
strutil.New(input).
    CleanWhitespace().
    Truncate(100, "...").
    SanitizeHTML(allowedTags).
    String()

// Validation with error handling
strutil.New(input).
    RequireNotEmpty().
    RequireEmail().
    Result()
```

## Implemented Functions (as of 7/16/2025)

The following functions have been fully implemented and are ready for use:

### String Utilities (`strutil`)

#### UUID and Random String Generation
- `GenerateUUID()` - Generates a random UUID
- `GenerateUUIDV7()` - Generates a UUID v7 (recommended for new applications)
- `RandomString(length)` - Generates a random alphanumeric string
- `RandomHex(length)` - Generates a random hexadecimal string
- `RandomUrlSafe(length)` - Generates a random URL-safe string

#### Lorem Ipsum Generation
- `LoremWord()` - Generates a random lorem ipsum word
- `LoremWords(count)` - Generates multiple lorem ipsum words
- `LoremSentence()` - Generates a lorem ipsum sentence
- `LoremSentenceCustom(count)` - Generates a custom-length lorem ipsum sentence
- `LoremSentences(count)` - Generates multiple lorem ipsum sentences
- `LoremSentencesCustom(count, length)` - Generates multiple custom-length lorem ipsum sentences
- `LoremSentencesVariable(count, min, max)` - Generates variable-length lorem ipsum sentences
- `LoremParagraph()` - Generates a lorem ipsum paragraph
- `LoremParagraphs(count)` - Generates multiple lorem ipsum paragraphs
- `LoremDomain()` - Generates a random domain name
- `LoremURL()` - Generates a random URL
- `LoremEmail()` - Generates a random email address

#### String Validation
- `IsEmail(s)` - Validates email addresses
- `IsURL(s)` - Validates URLs
- `IsDomain(domain)` - Validates domain names
- `IsUUID(s)` - Validates UUIDs
- `IsValidLength(s, min, max)` - Validates string length
- `IsEmpty(s)` - Checks if a string is empty
- `IsEmptyNormalized(s)` - Checks if a string is empty after normalization
- `IsAlphaNumericString(s)` - Checks if a string contains only alphanumeric characters
- `IsAlphaString(s)` - Checks if a string contains only alphabetic characters

#### String Transformation
- `ToUpper(s)` - Converts a string to uppercase
- `ToLower(s)` - Converts a string to lowercase
- `ReplaceWhitespace(s, replacement)` - Replaces whitespace with a specified string
- `ReplaceSpaces(s, replacement)` - Replaces spaces with a specified string
- `Trim(s)` - Trims whitespace from both ends of a string
- `TrimLeft(s)` - Trims whitespace from the left side of a string
- `TrimRight(s)` - Trims whitespace from the right side of a string
- `AlphaReplace(s, replacement)` - Replaces non-alphabetic characters
- `AlphaNumericReplace(s, replacement)` - Replaces non-alphanumeric characters
- `Slugify(s, length)` - Creates a URL-friendly slug
- `Truncate(s, length, suffix)` - Truncates a string to a specified length
- `KeepAlpha(s, ws)` - Keeps only alphabetic characters
- `KeepAlphaNumeric(s, ws)` - Keeps only alphanumeric characters
- `CleanWhitespace(s)` - Removes all whitespace
- `NormalizeWhitespace(s)` - Normalizes whitespace
- `CollapseWhitespace(s)` - Collapses consecutive whitespace characters
- `NormalizeDiacritics(s)` - Replaces diacritics with their ASCII equivalent

#### HTML and Sanitization
- `StripHTML(s)` - Removes all HTML tags
- `SanitizeHTML(s)` - Sanitizes HTML using UGC policy
- `SanitizeHTMLCustom(s, allowedTags)` - Sanitizes HTML with custom allowed tags
- `EscapeHTML(s)` - Escapes HTML special characters

### Version Utilities (`version`)

- `GetVersion()` - Returns the current version string
- `GetBuildInfo()` - Returns complete build information
- `IsValidSemVer(version)` - Validates semantic versioning format

The `BuildInfo` struct also provides methods:
- `String()` - Returns a human-readable version string
- `IsDevelopment()` - Checks if this is a development build
- `GetShortCommit()` - Returns the short commit hash

## Installation

```
go get github.com/bmj2728/utils
```

## Usage

Import the specific package you need:

```
import "github.com/bmj2728/utils/strutil"
```

Then use the functions or builders as needed:

```
// Using functional API
strutil.Slugify("Hello World!")

// Using string builder API
strutil.New(userInput).
    CleanWhitespace().
    RequireNotEmpty().
    Result()
```

## Testing

This project follows Go's standard testing practices. Each package includes comprehensive tests to ensure functionality, edge cases, and regression prevention.

### Running Tests Locally

To run all tests in the project:

```
go test ./...
```

For verbose output:

```
go test -v ./...
```

To run tests for a specific package:

```
go test ./strutil
```

### Continuous Integration

This project uses GitHub Actions for continuous integration. The workflow automatically runs:

1. **Linting & Formatting** - Using golangci-lint, go vet, and go fmt to ensure code quality and consistent formatting
2. **Testing** - Running all tests across multiple Go versions (1.22, 1.23, 1.24) with race detection
3. **Security Scanning** - Using gosec, govulncheck, and dependency review to identify security issues
4. **Building** - Building for multiple platforms (Linux, macOS, Windows) and architectures (amd64, arm64)

The CI/CD workflows run on all pull requests and pushes to the main branch, ensuring code quality, security, and functionality are maintained. Weekly security scans are also scheduled to catch newly discovered vulnerabilities.

#### CI/CD Workflow Details (as of 7/16/2025)

- **CI Workflow**: Handles linting, formatting, quick tests, and import verification
- **Test Workflow**: Runs comprehensive tests with race detection across multiple Go versions
- **Security Workflow**: Performs security scanning, vulnerability checking, and dependency review
- **Build Workflow**: Builds binaries for multiple platforms and architectures

All workflows use concurrency groups to avoid redundant runs and optimize CI/CD resources.

## Acknowledgements

This project leverages several excellent open-source libraries:

- [go-sanitize](https://github.com/mrz1836/go-sanitize) - A powerful Go sanitization package that provides robust string cleaning and sanitization functions.
- [google/uuid](https://github.com/google/uuid) - A robust UUID implementation for Go.
- [lorelai](https://github.com/UltiRequiem/lorelai) - A versatile lorem ipsum generator used for creating placeholder text, emails, URLs, and more.
- [bluemonday](https://github.com/microcosm-cc/bluemonday) - A HTML sanitizer implementation used for the HTML sanitization functions.
- [go-diacritics](github.com/Regis24GmbH/go-diacritics) - A lightweight function for normalizing diacritics

## License

This project is licensed under the terms of the LICENSE file included in the repository.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
