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

#### Functional API

The functional API provides standalone functions for string operations:

```
// Generate a UUID
strutil.NewUUID()

// Validate an email
strutil.IsEmail(email)

// Sanitize user input
strutil.SanitizeUsername(rawName)
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

## License

This project is licensed under the terms of the LICENSE file included in the repository.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
