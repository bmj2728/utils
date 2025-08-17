package strutil

// Index finds the first occurrence of the given substring within the StringBuilder's value and returns its byte index.
// If normalize is true, the StringBuilder's value and substring are normalized using Unicode NFC before searching.
// Returns -1 if the substring is not found or if processing is halted by shouldContinueProcessing.
func (sb *StringBuilder) Index(substr string, normalize bool) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return index(sb.value, substr, normalize)
}

// IndexAll returns a slice of starting indices of all non-overlapping occurrences
// of substr in the StringBuilder's value.
// If normalize is true, both the StringBuilder's value and substr will be normalized to NFC before searching.
// Returns nil if processing should not continue.
func (sb *StringBuilder) IndexAll(substr string, normalize bool) []int {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return indexAll(sb.value, substr, normalize)
}

// IndexRune returns the index of the first occurrence of the specified rune in the StringBuilder, or -1 if not found.
func (sb *StringBuilder) IndexRune(r rune) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return indexRune(sb.value, r)
}

// IndexAllRune returns a slice of all indices where the specified rune appears in the StringBuilder's value.
func (sb *StringBuilder) IndexAllRune(r rune) []int {
	if !sb.shouldContinueProcessing() {
		return nil
	}
	return indexAllRune(sb.value, r)
}

// LastIndex finds the last occurrence of substr in the StringBuilder's value.
// If normalize is true, the inputs are Unicode normalized before searching.
// Returns -1 if processing should not continue or substr is not found.
func (sb *StringBuilder) LastIndex(substr string, normalize bool) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return lastIndex(sb.value, substr, normalize)
}

// LastIndexRune returns the index of the last occurrence of the specified
// rune in the StringBuilder or -1 if not found.
func (sb *StringBuilder) LastIndexRune(r rune) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return lastIndexRune(sb.value, r)
}

// UTF32NormalizedIndex returns the normalized UTF-32 index of the given substring in the StringBuilder's value.
// Returns -1 if the substring is not found or processing should not continue.
func (sb *StringBuilder) UTF32NormalizedIndex(substr string) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return utf32NormalizedIndex(sb.value, substr)
}

// UTF32NormalizedLastIndex finds the last index of a UTF-32 normalized substring within the StringBuilder content.
// It normalizes both the StringBuilder's value and the substring using NFC Unicode normalization before searching.
// Returns -1 if normalization or UTF-32 conversion fails, or if the substring is not found.
func (sb *StringBuilder) UTF32NormalizedLastIndex(substr string) int {
	if !sb.shouldContinueProcessing() {
		return -1
	}
	return utf32NormalizedLastIndex(sb.value, substr)
}
