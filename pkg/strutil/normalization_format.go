package strutil

import (
	"golang.org/x/text/unicode/norm"
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
