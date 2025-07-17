package strutil

import "slices"

// EdLibData represents a structure for various edit distance and similarity measurement metrics.
type EdLibData struct {
	// LCS family
	LCS             *int      `json:"lcs,omitempty"`
	LCSBacktrack    *string   `json:"lcs_backtrack,omitempty"`
	LCSBacktrackAll *[]string `json:"lcs_backtrack_all,omitempty"`
	LCSDiff         *[]string `json:"lcs_diff,omitempty"`
	LCSEditDistance *int      `json:"lcs_edit_distance,omitempty"`

	// Levenshtein family
	LevenshteinDist   *int `json:"levenshtein_dist,omitempty"`
	DamerauLevDist    *int `json:"damerau_lev_dist,omitempty"`
	OSADamerauLevDist *int `json:"osa_damerau_lev_dist,omitempty"`

	// Other algorithms
	HammingDist      *int     `json:"hamming_dist,omitempty"`
	JaroSimilarity   *float64 `json:"jaro_similarity,omitempty"`
	JaroWinklerSim   *float64 `json:"jaro_winkler_similarity,omitempty"`
	JaccardSim       *float64 `json:"jaccard_similarity,omitempty"`
	CosineSimilarity *float64 `json:"cosine_similarity,omitempty"`
	SorensenDiceCo   *float64 `json:"sorensen_dice_similarity,omitempty"`

	// QGram
	QGramDist       *int `json:"qgram_dist,omitempty"`
	QGramDistCustom *int `json:"qgram_dist_custom,omitempty"`
	QGramSim        *int `json:"qgram_sim,omitempty"`

	// Shingle
	Shingle      *map[string]int `json:"shingle,omitempty"`
	ShingleSlice *[]string       `json:"shingle_slice,omitempty"`
}

// NewEdLibData initializes and returns a pointer to a new instance of EdLibData with default values.
func NewEdLibData() *EdLibData {
	return &EdLibData{}
}

// SetLCS sets the LCS (Longest Common Subsequence) value in the EdLibData structure.
func (e *EdLibData) SetLCS(lcs *int) {
	e.LCS = lcs
}

// GetLCS retrieves the pointer to the LCS (Longest Common Subsequence) value stored in EdLibData.
func (e *EdLibData) GetLCS() *int {
	return e.LCS
}

// SetLCSBacktrack sets the LCS backtrack field with the given string pointer value.
func (e *EdLibData) SetLCSBacktrack(lcsBacktrack *string) {
	e.LCSBacktrack = lcsBacktrack
}

func (e *EdLibData) GetLCSBacktrack() *string {
	return e.LCSBacktrack
}

// SetLCSBacktrackAll sets the LCSBacktrackAll field with the provided pointer to a slice of strings.
func (e *EdLibData) SetLCSBacktrackAll(lcsBacktrackAll *[]string) {
	e.LCSBacktrackAll = lcsBacktrackAll
}

// GetLCSBacktrackAll retrieves all possible backtracking results for the longest common subsequence (LCS).
func (e *EdLibData) GetLCSBacktrackAll() *[]string {
	return e.LCSBacktrackAll
}

// SetLCSDiff updates the LCSDiff property with the provided pointer to a slice of strings.
func (e *EdLibData) SetLCSDiff(lcsDiff *[]string) {
	e.LCSDiff = lcsDiff
}

// GetLCSDiff retrieves the LCS difference as a slice of strings from the EdLibData structure.
func (e *EdLibData) GetLCSDiff() *[]string {
	return e.LCSDiff
}

// SetLCSEditDistance sets the LCS edit distance value in the EdLibData struct.
func (e *EdLibData) SetLCSEditDistance(lcsEditDistance *int) {
	e.LCSEditDistance = lcsEditDistance
}

// GetLCSEditDistance retrieves the LCS edit distance value from an EdLibData instance.
// Returns a pointer to an integer representing the LCS edit distance, or nil if not set.
func (e *EdLibData) GetLCSEditDistance() *int {
	return e.LCSEditDistance
}

// SetLevenshteinDist sets the Levenshtein distance value in the EdLibData structure.
func (e *EdLibData) SetLevenshteinDist(levenshteinDist *int) {
	e.LevenshteinDist = levenshteinDist
}

// GetLevenshteinDist retrieves the Levenshtein distance value associated with the EdLibData instance.
func (e *EdLibData) GetLevenshteinDist() *int {
	return e.LevenshteinDist
}

// SetDamerauLevDist sets the Damerau-Levenshtein distance value for the EdLibData structure.
func (e *EdLibData) SetDamerauLevDist(damerauLevDist *int) {
	e.DamerauLevDist = damerauLevDist
}

// GetDamerauLevDist returns the Damerau-Levenshtein distance value, or nil if not set.
func (e *EdLibData) GetDamerauLevDist() *int {
	return e.DamerauLevDist
}

// SetOSADamerauLevDist sets the optimal string alignment Damerau-Levenshtein distance for the EdLibData instance.
func (e *EdLibData) SetOSADamerauLevDist(osaDamerauLevDist *int) {
	e.OSADamerauLevDist = osaDamerauLevDist
}

// GetOSADamerauLevDist retrieves the optimal string alignment
// Damerau-Levenshtein distance stored in the EdLibData struct.
func (e *EdLibData) GetOSADamerauLevDist() *int {
	return e.OSADamerauLevDist
}

// SetHammingDist updates the HammingDist field with the provided pointer to an integer.
func (e *EdLibData) SetHammingDist(hammingDist *int) {
	e.HammingDist = hammingDist
}

// GetHammingDist retrieves the Hamming distance value from the EdLibData instance.
func (e *EdLibData) GetHammingDist() *int {
	return e.HammingDist
}

// SetJaroSimilarity sets the Jaro similarity value for the EdLibData instance. It accepts a pointer to a float64 value.
func (e *EdLibData) SetJaroSimilarity(jaroSimilarity *float64) {
	e.JaroSimilarity = jaroSimilarity
}

// GetJaroSimilarity retrieves the Jaro similarity value, a measure of similarity between two strings, if available.
func (e *EdLibData) GetJaroSimilarity() *float64 {
	return e.JaroSimilarity
}

// SetJaroWinklerSim sets the Jaro-Winkler similarity score for the EdLibData instance.
func (e *EdLibData) SetJaroWinklerSim(jaroWinklerSim *float64) {
	e.JaroWinklerSim = jaroWinklerSim
}

// GetJaroWinklerSim retrieves the Jaro-Winkler similarity value as a pointer to float64 from the EdLibData structure.
func (e *EdLibData) GetJaroWinklerSim() *float64 {
	return e.JaroWinklerSim
}

// SetJaccardSim sets the Jaccard similarity value in the EdLibData structure.
func (e *EdLibData) SetJaccardSim(jaccardSim *float64) {
	e.JaccardSim = jaccardSim
}

// GetJaccardSim retrieves a pointer to the Jaccard similarity value from the EdLibData structure.
func (e *EdLibData) GetJaccardSim() *float64 {
	return e.JaccardSim
}

// SetCosineSimilarity sets the CosineSimilarity field to the provided float64 pointer value.
func (e *EdLibData) SetCosineSimilarity(cosineSimilarity *float64) {
	e.CosineSimilarity = cosineSimilarity
}

// GetCosineSimilarity retrieves the cosine similarity value from the EdLibData structure.
// Returns a pointer to a float64 representing the similarity or nil if not set.
func (e *EdLibData) GetCosineSimilarity() *float64 {
	return e.CosineSimilarity
}

// SetSorensenDiceCo sets the Sorensen-Dice similarity coefficient in the EdLibData structure.
func (e *EdLibData) SetSorensenDiceCo(sorensenDiceCo *float64) {
	e.SorensenDiceCo = sorensenDiceCo
}

// GetSorensenDiceCo retrieves the Sorensen-Dice coefficient as a pointer to a float64.
// This value indicates similarity between two sets.
func (e *EdLibData) GetSorensenDiceCo() *float64 {
	return e.SorensenDiceCo
}

// SetQGramDist sets the QGramDist value for the EdLibData instance.
func (e *EdLibData) SetQGramDist(qGramDist *int) {
	e.QGramDist = qGramDist
}

// GetQGramDist retrieves the QGramDist field, representing the Q-gram distance in the EdLibData structure.
func (e *EdLibData) GetQGramDist() *int {
	return e.QGramDist
}

// SetQGramDistCustom sets the custom Q-gram distance value for the EdLibData instance.
func (e *EdLibData) SetQGramDistCustom(qGramDistCustom *int) {
	e.QGramDistCustom = qGramDistCustom
}

// GetQGramDistCustom retrieves the custom Q-Gram distance associated with the EdLibData instance.
func (e *EdLibData) GetQGramDistCustom() *int {
	return e.QGramDistCustom
}

// SetQGramSim sets the Q-gram similarity value for the EdLibData instance.
func (e *EdLibData) SetQGramSim(qGramSim *int) {
	e.QGramSim = qGramSim
}

// GetQGramSim retrieves the QGram similarity score from the EdLibData structure.
func (e *EdLibData) GetQGramSim() *int {
	return e.QGramSim
}

// SetShingle sets the Shingle field of the EdLibData struct with the provided map pointer.
func (e *EdLibData) SetShingle(shingle *map[string]int) {
	e.Shingle = shingle
}

// GetShingle returns the shingle map representation associated with the EdLibData instance.
func (e *EdLibData) GetShingle() *map[string]int {
	return e.Shingle
}

// SetShingleSlice sets the ShingleSlice field of the EdLibData struct to the provided slice of strings.
func (e *EdLibData) SetShingleSlice(shingleSlice *[]string) {
	e.ShingleSlice = shingleSlice
}

// GetShingleSlice returns a pointer to the ShingleSlice field, which represents a slice of shingles.
func (e *EdLibData) GetShingleSlice() *[]string {
	return e.ShingleSlice
}

func CompareStringSlices(s1, s2 []string, nulls bool) bool {
	// true if both nil and null = true
	if nulls && s1 == nil && s2 == nil {
		return true
	}
	// any other nil scenario is false
	if s1 == nil || s2 == nil {
		return false
	}
	// confirm length
	if len(s1) != len(s2) {
		return false
	}
	// confirm everything in one is in the other
	for _, s := range s1 {
		if slices.Contains(s2, s) {
			continue
		}
		return false
	}
	return true
}
