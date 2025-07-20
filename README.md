
<div style="text-align: center;">
  <h1>Utils - The Missing Go Utilities Package</h1>
  <img src="assets/utils-gopher_transparent.png" alt="utils Gopher Logo" width="500"/>
</div>


<!-- Badges -->
[![CodeQL](https://github.com/bmj2728/utils/workflows/CodeQL/badge.svg)](https://github.com/bmj2728/utils/actions?query=workflow%3ACodeQL)
[![Linter/Formatter](https://github.com/bmj2728/utils/actions/workflows/ci.yml/badge.svg)](https://github.com/bmj2728/utils/actions)
[![Testing](https://github.com/bmj2728/utils/actions/workflows/test.yml/badge.svg)](https://github.com/bmj2728/utils/actions)
[![codecov](https://codecov.io/gh/bmj2728/utils/branch/main/graph/badge.svg)](https://codecov.io/gh/bmj2728/utils)
[![dev-codecov](https://codecov.io/gh/bmj2728/utils/branch/feature%2Fstrutil/graph/badge.svg?token=3gw1kjEFCr)](https://codecov.io/gh/bmj2728/utils)
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
├── pkg/
│   ├── strutil/                # String manipulation & validation
│   │   ├── builder.go          # Core StringBuilder implementation
│   │   ├── casing.go           # Main case conversion functions
│   │   ├── casing_builder.go   # Fluent API for case operations
│   │   ├── casing_helpers.go   # Internal case conversion implementations
│   │   ├── comparison.go       # Main comparison functions
│   │   ├── comparison_builder.go # Fluent API for comparison operations
│   │   ├── comparison_data.go  # ComparisonData struct definition
│   │   ├── comparison_helpers.go # Internal comparison implementations
│   │   ├── similarity_result.go # SimilarityResult struct definition
│   │   ├── generation.go       # Main string generation functions
│   │   ├── generation_builder.go # Fluent API for generation operations
│   │   ├── lorem.go            # Main lorem ipsum functions
│   │   ├── lorem_builder.go    # Fluent API for lorem ipsum operations
│   │   ├── sanitization.go     # Main sanitization functions
│   │   ├── sanitization_builder.go # Fluent API for sanitization operations
│   │   ├── transform.go        # Main string transformation functions
│   │   ├── transform_builder.go # Fluent API for transformation operations
│   │   ├── validation.go       # Main validation functions
│   │   ├── validation_builder.go # Fluent API for validation operations
│   │   └── *_test.go           # Test files for each functionality area
│   └── version/                # Version information utilities
├── httputil/        # HTTP client patterns (planned)
├── fileutl/         # File operations (planned)
├── jsonutil/        # JSON utilities (planned)
├── cryptoutil/      # Common crypto patterns (planned)
├── configutil/      # Configuration management (planned)
├── sliceutil/       # Generic slice operations (planned)
├── validationutil/  # Input validation (planned)
├── errorutil/       # Error handling (planned)
└── testutil/        # Test helpers (planned)
```

This structure maintains a clear separation between:
- Main public API functions (e.g., `comparison.go`)
- Internal implementation functions (e.g., `comparison_helpers.go`)
- Fluent builder pattern methods (e.g., `comparison_builder.go`)
- Type definitions (e.g., `similarity_result.go`, `comparison_data.go`)

## Recent Updates

The project has recently added several significant features and improvements:

- **Full Comparison Suite**: Implemented a comprehensive suite of string comparison functions including:
  - Longest Common Subsequence (LCS) with backtracking and diff capabilities
  - Hamming Distance
  - Jaro and Jaro-Winkler Distance
  - Jaccard Similarity
  - Cosine Similarity
  - Q-gram Distance and Similarity
  - Sorensen-Dice Coefficient
  - Shingle generation for text fingerprinting
- **Similarity Result Handling**: Added a `SimilarityResult` struct to standardize and enhance comparison results
- **StringBuilder History**: Enhanced the string builder to maintain a history of similarity results
- **Project Reorganization**: Restructured the codebase into more manageable files with clear separation between:
  - Internal implementation functions
  - Main public API functions
  - Fluent builder pattern methods
- **Test Organization**: Separated tests into multiple files by functionality area for better maintainability

## Roadmap

Upcoming features planned for future releases:

1. **Enhanced Sanitization Functions**: Additional sanitization for various data types and contexts
2. **Documentation Generation**: GitHub Actions for automatic documentation and pages generation
3. **Performance Optimizations**: Further optimizations for string operations on large datasets


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

// String comparison with similarity tracking
result, comparisonData, err := strutil.New(input).
    LevenshteinDistance(otherString).
    JaroSimilarity(otherString).
    Result()
```

#### Comparison Data Structures

> **Note**: The comparison functionality is evolving and likely to change before v0.1.0.

The string builder maintains two separate structures for tracking comparison results:

##### ComparisonData Struct

The `ComparisonData` struct stores various metrics from different string comparison algorithms:

```
type ComparisonData struct {
    // LCS family
    LCS             *int      `json:"lcs,omitempty"`
    LCSBacktrack    *string   `json:"lcs_backtrack,omitempty"`
    // Levenshtein family
    LevenshteinDist   *int `json:"levenshtein_dist,omitempty"`
    DamerauLevDist    *int `json:"damerau_lev_dist,omitempty"`
    // Other algorithms
    HammingDist      *int     `json:"hamming_dist,omitempty"`
    JaroSimilarity   *float32 `json:"jaro_similarity,omitempty"`
    // ... and many more metrics
}
```

**Purpose**: Stores raw metrics from comparison operations.

**Limitations**:
- Doesn't store information about the comparison text itself
- Can only hold comparison to a single text at a time
- Only updated by specific comparison methods (LevenshteinDistance, JaroSimilarity, etc.)

**Usage**:
```
// Access comparison data after operations
sb := strutil.New("hello").
    LevenshteinDistance("hallo").
    JaroSimilarity("hallo")

data := sb.ComparisonData()
levDist := data.GetLevenshteinDist()  // Get Levenshtein distance
jaroSim := data.GetJaroSimilarity()   // Get Jaro similarity score
```

##### SimilarityResult Struct

The `SimilarityResult` struct provides a standardized way to handle normalized string comparison results:

```
type SimilarityResult struct {
    Algorithm  string    // The algorithm used for comparison
    Str1       string    // The first string being compared
    Str2       string    // The second string being compared
    Similarity *float32  // The similarity score (nil if error occurred)
    Err        error     // Any error that occurred during comparison
}
```

**Purpose**: Provides normalized results using one of 11 algorithms.

**Benefits**:
- Self-contained with both the comparison result and the strings being compared
- Maintains context about which algorithm was used
- Multiple SimilarityResults can be stored in the `similarities` slice
- Standardized format regardless of the algorithm used

**Usage**:
```
// Using the Similarity function directly
result := strutil.Similarity("hello", "hallo", "levenshtein")
fmt.Println(result.String())  // Prints formatted comparison result

// Using the StringBuilder API to add to similarities slice
sb := strutil.New("hello").Similarity("hallo", "levenshtein")
// Additional comparisons can be added
sb.Similarity("hallo", "jaro")
```

## Implemented Functions (as of 7/16/2025)

The following functions have been fully implemented and are ready for use:

### String Utilities (`strutil`)

#### UUID and Random String Generation
- `GenerateUUID()` - Generates a random UUID
- `GenerateUUIDV7()` - Generates a UUID v7 (recommended for new applications)
- `RandomAlphaString(length)` - Generates a random alphanumeric string
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
- `Capitalize(s)` - Capitalizes the first letter of a string
- `Uncapitalize(s)` - Converts the first letter of a string to lowercase
- `ToTitleCase(s)` - Converts a string to title case
- `SplitCamelCase(s)` - Splits a camelCase string into space-separated words
- `SplitPascalCase(s)` - Splits a PascalCase string into space-separated words
- `ToSnakeCase(s, scream)` - Converts a string to snake_case or SCREAMING_SNAKE_CASE
- `ToSnakeCaseWithIgnore(s, scream, ignore)` - Converts a string to snake_case with custom ignore characters
- `ToKebabCase(s, scream)` - Converts a string to kebab-case or SCREAMING-KEBAB-CASE
- `ToCamelCase(s)` - Converts a string to camelCase
- `ToPascalCase(s)` - Converts a string to PascalCase
- `ToDelimited(s, delim, ignore, scream)` - Converts a string to a custom delimited format
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

#### String Comparison
- `LevenshteinDistance(s1, s2)` - Calculates the Levenshtein distance between two strings
- `DamerauLevenshteinDistance(s1, s2)` - Calculates the Damerau-Levenshtein distance between two strings
- `OSADamerauLevenshteinDistance(s1, s2)` - Calculates the Optimal String Alignment variant of Damerau-Levenshtein distance
- `LCS(s1, s2)` - Calculates the length of the longest common subsequence
- `LCSBacktrack(s1, s2)` - Calculates the longest common subsequence and returns the result
- `LCSBacktrackAll(s1, s2)` - Computes all longest common subsequences
- `LCSDiff(s1, s2)` - Computes the differences between strings using LCS
- `LCSEditDistance(s1, s2)` - Computes the edit distance using LCS
- `HammingDistance(s1, s2)` - Computes the Hamming distance between two strings
- `JaroSimilarity(s1, s2)` - Calculates the Jaro similarity between two strings
- `JaroWinklerSimilarity(s1, s2)` - Computes the Jaro-Winkler similarity between two strings
- `JaccardSimilarity(s1, s2, splitLength)` - Computes the Jaccard similarity coefficient
- `CosineSimilarity(s1, s2, splitLength)` - Computes the cosine similarity between two strings
- `SorensenDiceCoefficient(s1, s2, splitLength)` - Computes the Sørensen–Dice coefficient
- `QgramDistance(s1, s2, q)` - Calculates the q-gram distance between two strings
- `QgramDistanceCustomNgram(nmap1, nmap2)` - Computes the q-gram distance between n-gram frequency maps
- `QgramSimilarity(s1, s2, q)` - Calculates the q-gram similarity between two strings
- `Shingle(s, k)` - Generates k-shingles from the input string
- `ShingleSlice(s, k)` - Generates k-length shingles as a string slice
- `Similarity(s1, s2, algorithm)` - Computes the similarity score using a specified algorithm
- `CompareStringSlices(a, b, nulls)` - Compares two slices of strings for equality

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

### Test Organization

Tests have been organized into multiple files by functionality area for better maintainability:

- `casing_test.go` - Tests for case conversion functions
- `comparison_test.go` - Tests for string comparison functions
- `generation_test.go` - Tests for string generation functions
- `lorem_test.go` - Tests for lorem ipsum generation
- `sanitization_test.go` - Tests for HTML and text sanitization
- `transform_test.go` - Tests for string transformation functions
- `validation_test.go` - Tests for string validation functions
- `examples_test.go` - Example usage patterns for documentation

This organization makes it easier to locate and maintain tests for specific functionality areas.

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
go test ./pkg/strutil
```

To run tests for a specific functionality area (using test filtering):

```
go test ./pkg/strutil -run "TestComparison.*"
```

To run tests with benchmarks:

```
go test -bench=. ./pkg/strutil
```

To run tests with coverage reporting:

```
go test -cover ./pkg/strutil
go test -coverprofile=coverage.out ./pkg/strutil
go tool cover -html=coverage.out  # Opens coverage report in browser
```

### Continuous Integration

This project uses GitHub Actions for continuous integration. The workflow automatically runs:

1. **Linting & Formatting** - Using golangci-lint, go vet, and go fmt to ensure code quality and consistent formatting
2. **Testing** - Running all tests across multiple Go versions (1.22, 1.23, 1.24) with race detection
3. **Security Scanning** - Using gosec, govulncheck, and dependency review to identify security issues
4. **Building** - Building for multiple platforms (Linux, macOS, Windows) and architectures (amd64, arm64)

The CI/CD workflows run on all pull requests and pushes to the main branch, ensuring code quality, security, and functionality are maintained. Weekly security scans are also scheduled to catch newly discovered vulnerabilities.

#### CI/CD Workflow Details (as of 7/20/2025)

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
- [go-diacritics](github.com/Regis24GmbH/go-diacritics) - A lightweight function for normalizing diacritics.
- [strcase](https://github.com/iancoleman/strcase) - A comprehensive library for converting strings between different case formats.
- [camelcase](https://github.com/fatih/camelcase) - A library for splitting camelCase or PascalCase words into components.
- [go-edlib](https://github.com/hbollon/go-edlib) - A string comparison and edit distance library that provides various algorithms for measuring string similarity.

## License

This project is licensed under the terms of the LICENSE file included in the repository.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
