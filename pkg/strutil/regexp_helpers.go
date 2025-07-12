package strutil

import "regexp"

var (

	// LabelRegex defines a regular expression for validating domain labels, ensuring they meet DNS hostname requirements.
	LabelRegex = `[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?`

	// TLDRegex is a regular expression that matches valid top-level domain (TLD) strings
	// consisting of 2 or more alphabetic characters.
	TLDRegex = `[a-zA-Z]{2,}`

	// DomainRegex is a compiled regular expression validating fully qualified domain names (FQDN) based on DNS rules.
	DomainRegex = regexp.MustCompile(`^(` + LabelRegex + `\.)+` + TLDRegex + `$`)

	// CamelCaseRegex is a regular expression used to identify transitions between lowercase
	// and uppercase characters or consecutive capitals in camelCase or PascalCase strings.
	CamelCaseRegex = regexp.MustCompile(`([a-z0-9])([A-Z])|([A-Z])([A-Z][a-z])`)

	// WhiteSpaceRegex is a compiled regular expression that matches one or more whitespace characters.
	WhiteSpaceRegex = regexp.MustCompile(`\s+`)
)
