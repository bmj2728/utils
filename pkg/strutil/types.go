package strutil

import (
	"regexp"

	"golang.org/x/text/unicode/norm"
)

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
)

// NormalizationFormat represents a wrapper around norm.Form for specifying Unicode normalization forms.
type NormalizationFormat norm.Form

// NFC represents the Canonical Composition normalization format (NFC).
// NFD represents the Canonical Decomposition normalization format (NFD).
// NFKC represents the Compatibility Composition normalization format (NFKC).
// NFKD represents the Compatibility Decomposition normalization format (NFKD).
const (
	NFC  = NormalizationFormat(norm.NFC)
	NFD  = NormalizationFormat(norm.NFD)
	NFKC = NormalizationFormat(norm.NFKC)
	NFKD = NormalizationFormat(norm.NFKD)
)
