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

## Installation

```bash
go get github.com/bmj2728/utils
```

## Key Features

### String Utilities (`strutil`)

- **Comprehensive String Manipulation**: Transform, validate, and sanitize strings with a rich set of functions
- **Comparison Manager**: Central repository for managing and tracking string comparisons across different algorithms
- **Transformation History**: Optional tracking of string transformations through a history slice
- **Sanitization Functions**: Safely clean and sanitize strings for various contexts (HTML, filenames, etc.)
- **Multiple API Styles**: Choose between functional API or fluent builder pattern

### Version Utilities (`version`)

- **Version Information**: Get current version and build information
- **Semantic Version Validation**: Validate semantic versioning format

## Example Use Cases

### String Manipulation with Builder Pattern

```golang
// Chain multiple operations with the fluent builder API
package main

import "github.com/bmj2728/utils/pkg/strutil"

func example() {
    userInput := "<div>Some user input with   extra   spaces</div>"
    
    result, err := strutil.New(userInput).
        CleanWhitespace().
        Truncate(100, "...").
        SanitizeHTML().
        Result()

    if err != nil {
        // Handle error
    }
    
    // Use the result
    _ = result
}
```

### String Validation

```golang
package main

import "github.com/bmj2728/utils/pkg/strutil"

func validateExample() {
    // Validate email addresses
    email := "user@example.com"
    if strutil.IsEmail(email) {
        // Valid email
    }

    // Chain validations with the builder API
    input := "test@example.com"
    valid, err := strutil.New(input).
        RequireNotEmpty().
        RequireEmail().
        Result()
        
    // Use valid and err
    _ = valid
    _ = err
}
```

### String Comparison with History Tracking

```golang
package main

import (
	"fmt"

	"github.com/bmj2728/utils/pkg/strutil"
)

func comparisonExample() {
	// Compare strings and track results
	input := "hello world"
	otherString := "hello there"

	compMan := strutil.New(input).
		WithComparisonManager().
		LevenshteinDistance(otherString).
		JaroSimilarity(otherString).
		GetComparisonManager()

	// Access comparison results
	levenResult := compMan.GetComparisonResult(strutil.Levenshtein, otherString)
	jaroResult := compMan.GetSimilarityResult(strutil.Jaro, otherString)
	
	
}
```

### HTML Sanitization

```golang
package main

import "github.com/bmj2728/utils/pkg/strutil"

func sanitizationExample() {
    userInput := "<script>alert('xss')</script><p>Hello world</p>"
    
    // Sanitize HTML with default settings
    cleanHTML := strutil.SanitizeHTML(userInput)
    
    // Sanitize HTML with custom allowed tags
    allowedTags := []string{"p", "br", "strong", "em"}
    customCleanHTML := strutil.SanitizeHTMLCustom(userInput, allowedTags)
    
    // Use the results
    _ = cleanHTML
    _ = customCleanHTML
}
```

### Lorem Ipsum Generation

```golang
package main

import "github.com/bmj2728/utils/pkg/strutil"

func loremExample() {
    // Generate random text for testing
    sentence := strutil.LoremSentence()
    paragraph := strutil.LoremParagraph()
    email := strutil.LoremEmail()
    
    // Use the generated text
    _ = sentence
    _ = paragraph
    _ = email
}
```

## Documentation

For complete documentation of all available functions and their usage, please refer to the [Go Reference Documentation](https://pkg.go.dev/github.com/bmj2728/utils).

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