package strutil

// ComparisonData represents a structure for various edit distance and similarity measurement metrics.
type ComparisonData struct {
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
	JaroSimilarity   *float32 `json:"jaro_similarity,omitempty"`
	JaroWinklerSim   *float32 `json:"jaro_winkler_similarity,omitempty"`
	JaccardSim       *float32 `json:"jaccard_similarity,omitempty"`
	CosineSimilarity *float32 `json:"cosine_similarity,omitempty"`
	SorensenDiceCo   *float32 `json:"sorensen_dice_similarity,omitempty"`

	// QGram
	QGramDist       *int     `json:"qgram_dist,omitempty"`
	QGramDistCustom *int     `json:"qgram_dist_custom,omitempty"`
	QGramSim        *float32 `json:"qgram_sim,omitempty"`

	// Shingle
	Shingle      *map[string]int `json:"shingle,omitempty"`
	ShingleSlice *[]string       `json:"shingle_slice,omitempty"`
}

// NewComparisonData initializes and returns a pointer to a new instance of ComparisonData with default values.
func NewComparisonData() *ComparisonData {
	return &ComparisonData{}
}

// SetLCS sets the LCS (Longest Common Subsequence) value in the ComparisonData structure.
func (cd *ComparisonData) SetLCS(lcs *int) {
	cd.LCS = lcs
}

// GetLCS retrieves the pointer to the LCS (Longest Common Subsequence) value stored in ComparisonData.
func (cd *ComparisonData) GetLCS() *int {
	return cd.LCS
}

// SetLCSBacktrack sets the LCS backtrack field with the given string pointer value.
func (cd *ComparisonData) SetLCSBacktrack(lcsBacktrack *string) {
	cd.LCSBacktrack = lcsBacktrack
}

// GetLCSBacktrack retrieves the LCS backtrack string from the ComparisonData structure.
func (cd *ComparisonData) GetLCSBacktrack() *string {
	return cd.LCSBacktrack
}

// SetLCSBacktrackAll sets the LCSBacktrackAll field with the provided pointer to a slice of strings.
func (cd *ComparisonData) SetLCSBacktrackAll(lcsBacktrackAll *[]string) {
	cd.LCSBacktrackAll = lcsBacktrackAll
}

// GetLCSBacktrackAll retrieves all possible backtracking results for the longest common subsequence (LCS).
func (cd *ComparisonData) GetLCSBacktrackAll() *[]string {
	return cd.LCSBacktrackAll
}

// SetLCSDiff updates the LCSDiff property with the provided pointer to a slice of strings.
func (cd *ComparisonData) SetLCSDiff(lcsDiff *[]string) {
	cd.LCSDiff = lcsDiff
}

// GetLCSDiff retrieves the LCS difference as a slice of strings from the ComparisonData structure.
func (cd *ComparisonData) GetLCSDiff() *[]string {
	return cd.LCSDiff
}

// SetLCSEditDistance sets the LCS edit distance value in the ComparisonData struct.
func (cd *ComparisonData) SetLCSEditDistance(lcsEditDistance *int) {
	cd.LCSEditDistance = lcsEditDistance
}

// GetLCSEditDistance retrieves the LCS edit distance value from an ComparisonData instance.
// Returns a pointer to an integer representing the LCS edit distance, or nil if not set.
func (cd *ComparisonData) GetLCSEditDistance() *int {
	return cd.LCSEditDistance
}

// SetLevenshteinDist sets the Levenshtein distance value in the ComparisonData structure.
func (cd *ComparisonData) SetLevenshteinDist(levenshteinDist *int) {
	cd.LevenshteinDist = levenshteinDist
}

// GetLevenshteinDist retrieves the Levenshtein distance value associated with the ComparisonData instance.
func (cd *ComparisonData) GetLevenshteinDist() *int {
	return cd.LevenshteinDist
}

// SetDamerauLevDist sets the Damerau-Levenshtein distance value for the ComparisonData structure.
func (cd *ComparisonData) SetDamerauLevDist(damerauLevDist *int) {
	cd.DamerauLevDist = damerauLevDist
}

// GetDamerauLevDist returns the Damerau-Levenshtein distance value, or nil if not set.
func (cd *ComparisonData) GetDamerauLevDist() *int {
	return cd.DamerauLevDist
}

// SetOSADamerauLevDist sets the optimal string alignment Damerau-Levenshtein distance for the ComparisonData instance.
func (cd *ComparisonData) SetOSADamerauLevDist(osaDamerauLevDist *int) {
	cd.OSADamerauLevDist = osaDamerauLevDist
}

// GetOSADamerauLevDist retrieves the optimal string alignment
// Damerau-Levenshtein distance stored in the ComparisonData struct.
func (cd *ComparisonData) GetOSADamerauLevDist() *int {
	return cd.OSADamerauLevDist
}

// SetHammingDist updates the HammingDist field with the provided pointer to an integer.
func (cd *ComparisonData) SetHammingDist(hammingDist *int) {
	cd.HammingDist = hammingDist
}

// GetHammingDist retrieves the Hamming distance value from the ComparisonData instance.
func (cd *ComparisonData) GetHammingDist() *int {
	return cd.HammingDist
}

// SetJaroSimilarity sets the Jaro similarity value for the ComparisonData instance.
// It accepts a pointer to a float64 value.
func (cd *ComparisonData) SetJaroSimilarity(jaroSimilarity *float32) {
	cd.JaroSimilarity = jaroSimilarity
}

// GetJaroSimilarity retrieves the Jaro similarity value, a measure of similarity between two strings, if available.
func (cd *ComparisonData) GetJaroSimilarity() *float32 {
	return cd.JaroSimilarity
}

// SetJaroWinklerSim sets the Jaro-Winkler similarity score for the ComparisonData instance.
func (cd *ComparisonData) SetJaroWinklerSim(jaroWinklerSim *float32) {
	cd.JaroWinklerSim = jaroWinklerSim
}

// GetJaroWinklerSim retrieves the Jaro-Winkler similarity value as a
// pointer to float64 from the ComparisonData structure.
func (cd *ComparisonData) GetJaroWinklerSim() *float32 {
	return cd.JaroWinklerSim
}

// SetJaccardSim sets the Jaccard similarity value in the ComparisonData structure.
func (cd *ComparisonData) SetJaccardSim(jaccardSim *float32) {
	cd.JaccardSim = jaccardSim
}

// GetJaccardSim retrieves a pointer to the Jaccard similarity value from the ComparisonData structure.
func (cd *ComparisonData) GetJaccardSim() *float32 {
	return cd.JaccardSim
}

// SetCosineSimilarity sets the CosineSimilarity field to the provided float64 pointer value.
func (cd *ComparisonData) SetCosineSimilarity(cosineSimilarity *float32) {
	cd.CosineSimilarity = cosineSimilarity
}

// GetCosineSimilarity retrieves the cosine similarity value from the ComparisonData structure.
// Returns a pointer to a float64 representing the similarity or nil if not set.
func (cd *ComparisonData) GetCosineSimilarity() *float32 {
	return cd.CosineSimilarity
}

// SetSorensenDiceCo sets the Sorensen-Dice similarity coefficient in the ComparisonData structure.
func (cd *ComparisonData) SetSorensenDiceCo(sorensenDiceCo *float32) {
	cd.SorensenDiceCo = sorensenDiceCo
}

// GetSorensenDiceCo retrieves the Sorensen-Dice coefficient as a pointer to a float64.
// This value indicates similarity between two sets.
func (cd *ComparisonData) GetSorensenDiceCo() *float32 {
	return cd.SorensenDiceCo
}

// SetQGramDist sets the QGramDist value for the ComparisonData instance.
func (cd *ComparisonData) SetQGramDist(qGramDist *int) {
	cd.QGramDist = qGramDist
}

// GetQGramDist retrieves the QGramDist field, representing the Q-gram distance in the ComparisonData structure.
func (cd *ComparisonData) GetQGramDist() *int {
	return cd.QGramDist
}

// SetQGramDistCustom sets the custom Q-gram distance value for the ComparisonData instance.
func (cd *ComparisonData) SetQGramDistCustom(qGramDistCustom *int) {
	cd.QGramDistCustom = qGramDistCustom
}

// GetQGramDistCustom retrieves the custom Q-Gram distance associated with the ComparisonData instance.
func (cd *ComparisonData) GetQGramDistCustom() *int {
	return cd.QGramDistCustom
}

// SetQGramSim sets the Q-gram similarity value for the ComparisonData instance.
func (cd *ComparisonData) SetQGramSim(qGramSim *float32) {
	cd.QGramSim = qGramSim
}

// GetQGramSim retrieves the QGram similarity score from the ComparisonData structure.
func (cd *ComparisonData) GetQGramSim() *float32 {
	return cd.QGramSim
}

// SetShingle sets the Shingle field of the ComparisonData struct with the provided map pointer.
func (cd *ComparisonData) SetShingle(shingle *map[string]int) {
	cd.Shingle = shingle
}

// GetShingle returns the shingle map representation associated with the ComparisonData instance.
func (cd *ComparisonData) GetShingle() *map[string]int {
	return cd.Shingle
}

// SetShingleSlice sets the ShingleSlice field of the ComparisonData struct to the provided slice of strings.
func (cd *ComparisonData) SetShingleSlice(shingleSlice *[]string) {
	cd.ShingleSlice = shingleSlice
}

// GetShingleSlice returns a pointer to the ShingleSlice field, which represents a slice of shingles.
func (cd *ComparisonData) GetShingleSlice() *[]string {
	return cd.ShingleSlice
}
