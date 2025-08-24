package strutil

import (
	"maps"
	"testing"

	"github.com/bmj2728/utils/pkg/internal/comparison"
)

func TestSubstrings(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{"Substrings", "hello world!", []string{"hello", "world!"}},
		{"Substrings2", "hello world", []string{"hello", "world"}},
		{"Substrings3", "hello", []string{"hello"}},
		{"Substrings4", "", []string{}},
		{"Substrings5",
			"With great power there must also come great responsibility.",
			[]string{"With", "great", "power", "there", "must", "also", "come", "great", "responsibility."}},
		{"Substrings6", "abc 123 @username", []string{"abc", "123", "@username"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := substrings(tt.input)
			result := Substrings(tt.input)
			builderResult := New(tt.input).Substrings()
			if len(tt.want) > 0 && (!comparison.CompareSlices(helperResult, tt.want, false) ||
				!comparison.CompareSlices(result, tt.want, false) ||
				!comparison.CompareSlices(builderResult, tt.want, false)) {
				t.Errorf("Substrings - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult)
			}
			if len(tt.want) == 0 && (len(helperResult) > 0 || len(result) > 0 || len(builderResult) > 0) {
				t.Errorf("Substrings - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult,
				)
			}
		})
	}
}

func TestWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{"Words1", "hello world!", []string{"hello", "world"}},
		{"Words2", "hello world", []string{"hello", "world"}},
		{"Words3", "hello", []string{"hello"}},
		{"Words4", "", []string{}},
		{"Words5",
			"With great power there must also come great responsibility.",
			[]string{"With", "great", "power", "there", "must", "also", "come", "great", "responsibility"}},
		{"Words6", "abc 123 @username", []string{"abc", "123", "username"}},
		{"Words7", " [{'Hello', 'World'}] ", []string{"Hello", "World"}},
		{"Words8", "Hello !!! World", []string{"Hello", "World"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := words(tt.input)
			result := Words(tt.input)
			builderResult := New(tt.input).Words()
			if len(tt.want) > 0 && (!comparison.CompareSlices(helperResult, tt.want, false) ||
				!comparison.CompareSlices(result, tt.want, false) ||
				!comparison.CompareSlices(builderResult, tt.want, false)) {
				t.Errorf("Words - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					builderResult,
					result,
				)
			}
			if len(tt.want) == 0 && (len(helperResult) > 0 || len(result) > 0 || len(builderResult) > 0) {
				t.Errorf("Words - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult)
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		int
	}{
		{"Words1", "hello world!", 2},
		{"Words2", "hello world", 2},
		{"Words3", "hello", 1},
		{"Words4", "", 0},
		{"Words5",
			"With great power there must also come great responsibility.",
			9},
		{"Words6", "abc 123 @username", 3},
		{"Words7", " [{'Hello', 'World'}] ", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := wordCount(tt.input)
			result := WordCount(tt.input)
			builderResult := New(tt.input).WordCount()
			if helperResult != tt.int || result != tt.int || builderResult != tt.int {
				t.Errorf("WordCount - expected %d - got %d / %d / %d",
					tt.int,
					helperResult,
					result,
					builderResult)
			}
		})
	}
}

func TestWordMap(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{"WordMap1", "hello world!", map[string]int{"hello": 1, "world": 1}},
		{"WordMap2", "hello world", map[string]int{"hello": 1, "world": 1}},
		{"WordMap3", "hello", map[string]int{"hello": 1}},
		{"WordMap4", "", map[string]int{}},
		{"WordMap5",
			"With great power there must also come great responsibility.",
			map[string]int{"With": 1, "great": 2, "power": 1, "there": 1, "must": 1, "come": 1, "also": 1, "responsibility": 1}},
		{"WordMap6", "abc 123 @username", map[string]int{"abc": 1, "123": 1, "username": 1}},
		{"WordMap7", " [{'Hello', 'World'}] ", map[string]int{"Hello": 1, "World": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := wordMap(tt.input)
			result := WordMap(tt.input)
			builderResult := New(tt.input).WordMap()
			if !maps.Equal(helperResult, tt.expected) ||
				!maps.Equal(result, tt.expected) ||
				!maps.Equal(builderResult, tt.expected) {
				t.Errorf("WordMap - expected %v - got %v / %v / %v",
					tt.expected,
					helperResult,
					result,
					builderResult)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		sep         string
		preserveSep bool
		want        []string
	}{
		{"Split1", "hello world!", " ", false, []string{"hello", "world!"}},
		{"Split2", "hello world!", " ", true, []string{"hello ", "world!"}},
		{"Split3", "a, b, c, d, e", ", ", false, []string{"a", "b", "c", "d", "e"}},
		{"Split4",
			"a, b, c, d, e",
			", ",
			true,
			[]string{"a, ", "b, ", "c, ", "d, ", "e"}},
		{"Split5", "1 & 2 & 3 & 4", " & ", false, []string{"1", "2", "3", "4"}},
		{"Split6", "1 & 2 & 3 & 4", " & ", true, []string{"1 & ", "2 & ", "3 & ", "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := split(tt.input, tt.sep, tt.preserveSep)
			result := Split(tt.input, tt.sep, tt.preserveSep)
			builderResult := New(tt.input).Split(tt.sep, tt.preserveSep)
			if len(tt.want) > 0 && (!comparison.CompareSlices(helperResult, tt.want, false) ||
				!comparison.CompareSlices(result, tt.want, false) ||
				!comparison.CompareSlices(builderResult, tt.want, false)) {
				t.Errorf("Split - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult)
			}
			if len(tt.want) == 0 && (len(helperResult) > 0 || len(result) > 0 || len(builderResult) > 0) {
				t.Errorf("Split - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult)
			}
		})
	}
}

func TestSplitAt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		index int
		want  []string
	}{
		{"SplitAt1", "hello world!", 5, []string{"hello", "world!"}},
		{"SplitAt2", "hello world!", 0, []string{"", "hello world!"}},
		{"SplitAt3", "hello world!", 1, []string{"h", "ello world!"}},
		{"SplitAt4", "hello world!", 2, []string{"he", "llo world!"}},
		{"SplitAt5", "hello world!", -1, []string{}},
		{"SplitAt6", "hello world!", 100, []string{}},
		{"SplitAt7", "hello world!", 6, []string{"hello", "world!"}},
		{"SplitAt8", "", 0, []string{"", ""}},
		{"SplitAt9", "hello world!", 11, []string{"hello world", "!"}},
		{"SplitAt10", "\n\n\n\n\nHi!", 4, []string{"", "Hi!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := splitAt(tt.input, tt.index)
			result := SplitAt(tt.input, tt.index)
			builderResult := New(tt.input).SplitAt(tt.index)
			if len(tt.want) > 0 && (!comparison.CompareSlices(helperResult, tt.want, false) ||
				!comparison.CompareSlices(result, tt.want, false) ||
				!comparison.CompareSlices(builderResult, tt.want, false)) {
				t.Errorf("SplitAt - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult,
				)
			}
		})
	}
}

func TestSplitN(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		sep         string
		quantity    int
		preserveSep bool
		want        []string
	}{
		{"Split1", "hello world!", " ", 2, false, []string{"hello", "world!"}},
		{"Split2", "hello world!", " ", 2, true, []string{"hello ", "world!"}},
		{"Split3", "a, b, c, d, e", ", ", 2, false, []string{"a", "b, c, d, e"}},
		{"Split4",
			"a, b, c, d, e",
			", ",
			2,
			true,
			[]string{"a, ", "b, c, d, e"}},
		{"Split5", "1 & 2 & 3 & 4", " & ", 2,
			false, []string{"1", "2 & 3 & 4"}},
		{"Split6", "1 & 2 & 3 & 4", " & ", 2,
			true, []string{"1 & ", "2 & 3 & 4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := splitN(tt.input, tt.sep, tt.quantity, tt.preserveSep)
			result := SplitN(tt.input, tt.sep, tt.quantity, tt.preserveSep)
			builderResult := New(tt.input).SplitN(tt.sep, tt.quantity, tt.preserveSep)
			if len(tt.want) > 0 && (!comparison.CompareSlices(helperResult, tt.want, false) ||
				!comparison.CompareSlices(result, tt.want, false) ||
				!comparison.CompareSlices(builderResult, tt.want, false)) {
				t.Errorf("SplitN - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult)
			}
			if len(tt.want) == 0 && (len(helperResult) > 0 || len(result) > 0 || len(builderResult) > 0) {
				t.Errorf("SplitN - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult)
			}
		})
	}
}

func TestLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{"Lines1", "hello world!", []string{"hello world!"}},
		{"Lines2", "hello\nworld", []string{"hello", "world"}},
		{"Lines3", "hello\n", []string{"hello"}},
		{"Lines4", "\n\n\n", []string{"", "", ""}},
		{"Lines5", "This is\nmulti-line\ntext entry", []string{"This is", "multi-line", "text entry"}},
		{"Lines6", "\nHello\nWorld\n", []string{"", "Hello", "World"}},
		{"Lines7", "", []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := lines(tt.input)
			result := Lines(tt.input)
			builderResult := New(tt.input).Lines()
			if len(tt.want) > 0 && (!comparison.CompareSlices(helperResult, tt.want, false) ||
				!comparison.CompareSlices(result, tt.want, false) ||
				!comparison.CompareSlices(builderResult, tt.want, false)) {
				t.Errorf("Lines - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult,
				)
			}
			if len(tt.want) == 0 && (len(helperResult) > 0 || len(result) > 0 || len(builderResult) > 0) {
				t.Errorf("Lines - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult,
				)
			}
		})
	}
}

func TestLineCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Lines1", "hello world!", 1},
		{"Lines2", "hello\nworld", 2},
		{"Lines3", "hello\n", 1},
		{"Lines4", "\n\n\n", 3},
		{"Lines5", "This is\nmulti-line\ntext entry", 3},
		{"Lines6", "\nHello\nWorld\n", 3},
		{"Lines7", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := lineCount(tt.input)
			result := LineCount(tt.input)
			builderResult := New(tt.input).LineCount()
			if helperResult != tt.want || result != tt.want || builderResult != tt.want {
				t.Errorf("LineCount - expected %d - got %d / %d / %d",
					tt.want,
					helperResult,
					result,
					builderResult,
				)
			}
		})
	}
}

func TestSentences(t *testing.T) {
	tests := []struct {
		name  string
		input string
		lang  string
		want  []string
	}{
		{"Sentences1", "hello world!", "en",
			[]string{"hello world!"}},
		{"Sentences2", "hello world. How are you?", "en",
			[]string{"hello world.", "How are you?"}},
		{"Sentences3", "hello world. How are you? Hello world. How are you?", "en",
			[]string{"hello world.", "How are you?", "Hello world.", "How are you?"}},
		{"Sentences4", "", "en",
			[]string{}},
		{"Sentences5", "This is a sentence. This is another sentence.", "xx",
			[]string{"This is a sentence.", "This is another sentence."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := sentences(tt.input, tt.lang)
			result := Sentences(tt.input, tt.lang)
			builderResult := New(tt.input).Sentences(tt.lang)
			if len(tt.want) > 0 && (!comparison.CompareSlices(helperResult, tt.want, false) ||
				!comparison.CompareSlices(result, tt.want, false) ||
				!comparison.CompareSlices(builderResult, tt.want, false)) {
				t.Errorf("Sentences - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult,
				)
			}
			if len(tt.want) == 0 && (len(helperResult) > 0 || len(result) > 0 || len(builderResult) > 0) {
				t.Errorf("Sentences - expected %v - got %v / %v / %v",
					tt.want,
					helperResult,
					result,
					builderResult,
				)
			}
		})
	}
}

func TestSentenceCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		lang  string
		want  int
	}{
		{"Sentences1", "hello world!", "en",
			1},
		{"Sentences2", "hello world. How are you?", "en",
			2},
		{"Sentences3", "hello world. How are you? Hello world. How are you?", "en",
			4},
		{"Sentences4", "", "en",
			0},
		{"Sentences5", "This is a sentence. This is another sentence.", "xx",
			2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := sentenceCount(tt.input, tt.lang)
			result := SentenceCount(tt.input, tt.lang)
			builderResult := New(tt.input).SentenceCount(tt.lang)
			if helperResult != tt.want || result != tt.want || builderResult != tt.want {
				t.Errorf("SentenceCount - expected %d - got %d / %d / %d",
					tt.want,
					helperResult,
					result,
					builderResult,
				)
			}
		})
	}
}
