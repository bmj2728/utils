package strutil

import (
	"testing"

	"github.com/bmj2728/utils/pkg/internal/comparison"

	punkt "github.com/neurosnap/sentences"
)

func TestGetPunktTokenizer(t *testing.T) {
	tokenizer := getPunktTokenizer()
	if tokenizer == nil {
		t.Error("Punkt tokenizer is nil")
	}
}

func TestPunktSentenceData(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		want   int
		result []*punkt.Sentence
	}{
		{"SentenceData1", "Hello world! How are you today?", 2,
			[]*punkt.Sentence{
				{0, 12, "Hello world!"},
				{12, 31, " How are you today?"},
			},
		},
		{"SentenceData2", "Hello, Dr. Smith, I presume? How are you today?", 2,
			[]*punkt.Sentence{
				{0, 28, "Hello, Dr. Smith, I presume?"},
				{28, 47, " How are you today?"},
			},
		},
		{"SentenceData3", "Hello world! How are you today? Goodbye World. See you tomorrow.", 4,
			[]*punkt.Sentence{
				{0, 12, "Hello world!"},
				{12, 31, " How are you today?"},
				{31, 46, " Goodbye World."},
				{46, 63, " See you tomorrow."},
			},
		},
		{"SentenceData4", "", 0,
			[]*punkt.Sentence{},
		},
		{"SentenceData5", "Hi... Dr.? Sorry about the mess.", 2,
			[]*punkt.Sentence{
				{0, 10, "Hi... Dr.?"},
				{10, 32, " Sorry about the mess."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := punktSentenceData(tt.input)
			result := PunktSentenceData(tt.input)
			builderResult := New(tt.input).PunktSentenceData()
			if len(helperResult) != tt.want || len(result) != tt.want || len(builderResult) != tt.want {
				t.Errorf("PunktSentenceData - expected %d - got %d / %d / %d",
					tt.want,
					len(helperResult),
					len(result),
					len(builderResult))
			}
			hf := false
			f := false
			bf := false
			for _, v := range tt.result {
				for _, v2 := range helperResult {
					if v.Start == v2.Start && v.End == v2.End && v.Text == v2.Text {
						hf = true
						break
					}
				}
				for _, v2 := range result {
					if v.Start == v2.Start && v.End == v2.End && v.Text == v2.Text {
						f = true
						break
					}
				}
				for _, v2 := range builderResult {
					if v.Start == v2.Start && v.End == v2.End && v.Text == v2.Text {
						bf = true
						break
					}
				}
			}
			if len(tt.result) > 0 && (!hf || !f || !bf) {
				t.Errorf("PunktSentenceData - expected %v - got %v / %v / %v",
					tt.result,
					helperResult,
					result,
					builderResult)
			}
		})
	}
}

func TestPunktSentences(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		want   int
		result []string
	}{
		{"Sentences1",
			"Hello world! How are you today?",
			2,
			[]string{"Hello world!", "How are you today?"},
		},
		{"Sentences2",
			"Hello, Dr. Smith, I presume? How are you today?",
			2,
			[]string{"Hello, Dr. Smith, I presume?", "How are you today?"},
		},
		{"Sentences3",
			"Hello world! How are you today? Goodbye World. See you tomorrow.",
			4,
			[]string{"Hello world!", "How are you today?", "Goodbye World.", "See you tomorrow."},
		},
		{"Sentences4",
			"",
			0,
			[]string{},
		},
		{"Sentences5",
			"Hi... Dr.? Sorry about the mess.",
			2,
			[]string{"Hi... Dr.?", "Sorry about the mess."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := punktSentences(tt.input)
			hc := punktSentenceCount(tt.input)
			result := PunktSentences(tt.input)
			rc := PunktSentenceCount(tt.input)
			builderResult := New(tt.input).PunktSentences()
			bc := New(tt.input).PunktSentenceCount()
			if hc != tt.want || rc != tt.want || bc != tt.want {
				t.Errorf("PunktSentences - expected %d - got %d / %d / %d",
					tt.want,
					hc,
					rc,
					bc)
			}
			if len(tt.result) > 0 && (!comparison.CompareSlices(helperResult, tt.result, true) ||
				!comparison.CompareSlices(result, tt.result, true) ||
				!comparison.CompareSlices(builderResult, tt.result, true)) {
				t.Errorf("PunktSentences - expected %v - got %v / %v / %v",
					tt.result,
					helperResult,
					result,
					builderResult)
			}
		})
	}
}

func TestPunktSentencesCleaned(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		want   int
		result []string
	}{
		{"Sentences1",
			"Hello world! How are you today?",
			2,
			[]string{"Hello world!", "How are you today?"},
		},
		{"Sentences2",
			"Hello, Dr. Smith, I presume? How are you today?",
			2,
			[]string{"Hello, Dr. Smith, I presume?", "How are you today?"},
		},
		{"Sentences3",
			"Hello world! How are you today? Goodbye World. See you tomorrow.",
			4,
			[]string{"Hello world!", "How are you today?", "Goodbye World.", "See you tomorrow."},
		},
		{"Sentences4",
			"",
			0,
			[]string{},
		},
		{"Sentences5",
			"Hi... Dr.? Sorry about the mess.",
			3,
			[]string{"Hi...", "Dr.?", "Sorry about the mess."},
		},
		{"Sentences6",
			"Hi Dr.! Sorry about the mess.",
			2,
			[]string{"Hi Dr.!", "Sorry about the mess."},
		},
		{"Sentences7",
			"It's 8 p.m.? How'd it get so late?",
			2,
			[]string{"It's 8 p.m.?", "How'd it get so late?"},
		},
		{"Sentences8",
			"...the plot thickens. Hi... Dr.? Sorry about the mess.",
			4,
			[]string{"...the plot thickens.", "Hi...", "Dr.?", "Sorry about the mess."},
		},
		{"Sentences9",
			"Yes!",
			1,
			[]string{"Yes!"},
		},
		{"Sentences10",
			"$1,999,999,99.99 is a lot of money. A larger amount is $2,000,000,00.00.",
			2,
			[]string{"$1,999,999,99.99 is a lot of money.", "A larger amount is $2,000,000,00.00."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := punktSentencesCleaned(tt.input)
			hc := punktSentenceCountCleaned(tt.input)
			result := PunktSentencesCleaned(tt.input)
			rc := PunktSentenceCountCleaned(tt.input)
			builderResult := New(tt.input).PunktSentencesCleaned()
			bc := New(tt.input).PunktSentenceCountCleaned()
			if hc != tt.want || rc != tt.want || bc != tt.want {
				t.Errorf("PunktSentencesCleaned - expected %d - got %d / %d / %d",
					tt.want,
					hc,
					rc,
					bc)
			}
			if len(tt.result) > 0 && (!comparison.CompareSlices(helperResult, tt.result, true) ||
				!comparison.CompareSlices(result, tt.result, true) ||
				!comparison.CompareSlices(builderResult, tt.result, true)) {
				t.Errorf("PunktSentencesCleaned - expected %v - got %v / %v / %v",
					tt.result,
					helperResult,
					result,
					builderResult)
			}
		})
	}
}

func TestPunktAnnotatedTokens(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		want   int
		result []string
	}{
		{"AnnotatedTokens1", "Hello world! How are you today?", 6,
			[]string{"Hello", "world!", "How", "are", "you", "today?"},
		},
		{"AnnotatedTokens2", "Hello, Dr. Smith, I presume?", 5,
			[]string{"Hello,", "Dr.", "Smith,", "I", "presume?"},
		},
		{"AnnotatedTokens3", "Hello world!", 2,
			[]string{"Hello", "world!"},
		},
		{"AnnotatedTokens4", "", 0,
			[]string{},
		},
		{"AnnotatedTokens5", "Hi... Dr.? Sorry about the mess.", 6,
			[]string{"Hi...", "Dr.?", "Sorry", "about", "the", "mess."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := punktAnnotatedTokens(tt.input)
			hc := punktAnnotatedTokenCount(tt.input)
			result := PunktAnnotatedTokens(tt.input)
			rc := PunktAnnotatedTokenCount(tt.input)
			builderResult := New(tt.input).PunktAnnotatedTokens()
			bc := New(tt.input).PunktAnnotatedTokenCount()
			if hc != tt.want || rc != tt.want || bc != tt.want {
				t.Errorf("PunktAnnotatedTokens - expected %d - got %d / %d / %d",
					tt.want,
					hc,
					rc,
					bc)
			}
			if len(tt.result) > 0 && (!comparison.CompareSlices(helperResult, tt.result, true) ||
				!comparison.CompareSlices(result, tt.result, true) ||
				!comparison.CompareSlices(builderResult, tt.result, true)) {
				t.Errorf("PunktAnnotatedTokens - expected %v - got %v / %v / %v",
					tt.result,
					helperResult,
					result,
					builderResult)
			}
		})
	}
}

func TestPunktAnnotatedTokensData(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		want   int
		result []string
	}{
		{"AnnotatedData1",
			"Hello World!",
			2,
			[]string{`<Token Tok: "Hello", SentBreak: false, Abbr: false, Position: 5>`,
				`<Token Tok: "World!", SentBreak: true, Abbr: false, Position: 12>`},
		},
		{"AnnotatedData2",
			"Hello Dr. Smith.",
			3,
			[]string{`<Token Tok: "Hello", SentBreak: false, Abbr: false, Position: 5>`,
				`<Token Tok: "Dr.", SentBreak: false, Abbr: true, Position: 9>`,
				`<Token Tok: "Smith.", SentBreak: true, Abbr: false, Position: 16>`},
		},
		{"AnnotatedData3",
			"Hi!!!",
			1,
			[]string{`<Token Tok: "Hi!!!", SentBreak: true, Abbr: false, Position: 5>`},
		},
		{"AnnotatedData4",
			"",
			0,
			[]string{},
		},
		{"AnnotatedData5",
			"11:45 p.m.",
			2,
			[]string{`<Token Tok: "11:45", SentBreak: false, Abbr: false, Position: 5>`,
				`<Token Tok: "p.m.", SentBreak: false, Abbr: true, Position: 10>`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperResult := punktAnnotatedTokensData(tt.input)
			var hrSl []string
			result := PunktAnnotatedTokensData(tt.input)
			var rSl []string
			builderResult := New(tt.input).PunktAnnotatedTokensData()
			var brSl []string
			for _, v := range helperResult {
				hrSl = append(hrSl, v.String())
			}
			for _, v := range result {
				rSl = append(rSl, v.String())
			}
			for _, v := range builderResult {
				brSl = append(brSl, v.String())
			}
			if len(tt.result) > 0 && (!comparison.CompareSlices(hrSl, tt.result, true) ||
				!comparison.CompareSlices(rSl, tt.result, true) ||
				!comparison.CompareSlices(brSl, tt.result, true)) {
				t.Errorf("PunktAnnotatedTokensData - expected %v - got %v / %v / %v",
					tt.result,
					hrSl,
					rSl,
					brSl)
			}
		})
	}
}
