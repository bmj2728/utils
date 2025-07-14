# Utils - The Missing Go Utilities Package

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
strutil.NewUUID()

// Validate an email
strutil.IsEmail(email)

// Sanitize user input
strutil.SanitizeUsername(rawName)

// Keep only alphabetic characters
strutil.KeepAlpha(input, false)

// Keep only alphanumeric characters
strutil.KeepAlphaNumeric(input, true)
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

1. **Linting** - Using golangci-lint to ensure code quality
2. **Testing** - Running all tests to verify functionality

The CI workflow runs on all pull requests and pushes to the main branch, ensuring code quality and functionality are maintained.

## Acknowledgements

This project leverages several excellent open-source libraries:

- [go-sanitize](https://github.com/mrz1836/go-sanitize) - A powerful Go sanitization package that provides robust string cleaning and sanitization functions. The `strutil` package uses go-sanitize for its `KeepAlpha` and `KeepAlphaNumeric` functions.
- [google/uuid](https://github.com/google/uuid) - A robust UUID implementation for Go.

## License

This project is licensed under the terms of the LICENSE file included in the repository.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
