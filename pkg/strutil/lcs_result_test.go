package strutil

import (
	"testing"

	"utils/pkg/internal/comparison"
	"utils/pkg/internal/errors"
)

var (
	lcsAll1 = New("With great power there must also come great responsibility.").
		WithComparisonManager().
		LCSBacktrack("With great power comes great responsibility.").
		LCSBacktrackAll("With great power comes great responsibility.").
		LCSDiff("With great power comes great responsibility.").
		GetComparisonManager().
		GetLCSResultsMap()
	bt1  = lcsAll1[LCSBacktrackWord]["With great power comes great responsibility."]
	bta1 = lcsAll1[LCSBacktrackWordAll]["With great power comes great responsibility."]
	ld1  = lcsAll1[LCSDiffSlice]["With great power comes great responsibility."]

	lcsAll2 = New("With great power there must also come great responsibility.").
		WithComparisonManager().
		LCSBacktrack("With great power comes great responsibility.").
		LCSBacktrackAll("With great power comes great responsibility.").
		LCSDiff("With great power comes great responsibility.").
		GetComparisonManager().
		GetLCSResultsMap()
	bt2  = lcsAll2[LCSBacktrackWord]["With great power comes great responsibility."]
	bta2 = lcsAll2[LCSBacktrackWordAll]["With great power comes great responsibility."]
	ld2  = lcsAll2[LCSDiffSlice]["With great power comes great responsibility."]

	lcsAll3 = New("Hello, World!").
		WithComparisonManager().
		LCSBacktrack("Hello World").
		LCSBacktrackAll("Hello World").
		LCSDiff("Hello World").
		GetComparisonManager().
		GetLCSResultsMap()
	bt3 = lcsAll3[LCSBacktrackWord]["Hello World"]

	ld3 = lcsAll3[LCSDiffSlice]["Hello World"]

	btp1Sh   = "LCS Backtrack: With great power come great responsibility.\n"
	btp1Verb = "LCS Backtrack:\n" +
		"First: With great power there must also come great responsibility.\n" +
		"Second: With great power comes great responsibility.\n" +
		"Word: With great power come great responsibility.\n"
	btap1Sh = "LCS Backtrack All(1):\n" +
		"First: With great power come great responsibility.\n"
	btap1Verb = "LCS Backtrack All:\n" +
		"First: With great power there must also come great responsibility.\n" +
		"Second: With great power comes great responsibility.\n" +
		"Words:\n" +
		"With great power come great responsibility.\n"
	ldp1Sh = "LCS Diff:\n" +
		" W i t h   g r e a t   p o w e r   t h e r e   m u s t   a l s o   " +
		"c o m e s   g r e a t   r e s p o n s i b i l i t y ." +
		"\n                             - - - - -     - - - - - - - - - - -           " +
		"+                                            \n"
	ldp1Verb = "LCS Diff (With great power there must also come great responsibility." +
		"/With great power comes great responsibility.):\n" +
		" W i t h   g r e a t   p o w e r   t h e r e   m u s t   a l s o   " +
		"c o m e s   g r e a t   r e s p o n s i b i l i t y .\n" +
		"                             - - - - -     - - - - - - - - - - -           " +
		"+                                            \n"
)

func TestLCSResultTypeString(t *testing.T) {
	tests := []struct {
		name          string
		lcsResultType LCSResultType
		want          string
	}{
		{"BacktackWord", LCSBacktrackWord, "LCS Backtrack"},
		{"BacktackWordAll", LCSBacktrackWordAll, "LCS Backtrack All"},
		{"BacktackDiff", LCSDiffSlice, "LCS Diff"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lcsResultType.String(); got != tt.want {
				t.Errorf("LCSResultTypeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCSResultGetters(t *testing.T) {
	tests := []struct {
		name          string
		lcsResultType LCSResultType
		str1          string
		str2          string
		result        []string
		err           error
	}{
		{"BacktackWord",
			LCSBacktrackWord,
			"Hello, World!",
			"Hello World",
			[]string{"Hello World"},
			nil},
		{"BacktackWordAll",
			LCSBacktrackWordAll,
			"Hello, World!",
			"Hello World",
			[]string{"Hello World"},
			nil},
		{"BacktackDiff",
			LCSDiffSlice,
			"Hello, World!",
			"Hello World",
			[]string{" H e l l o ,   W o r l d !", "           -             -"},
			nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := *New(tt.str1).
				WithComparisonManager().
				LCSBacktrack(tt.str2).
				LCSBacktrackAll(tt.str2).
				LCSDiff(tt.str2).
				GetComparisonManager().
				GetLCSResultsMap()[tt.lcsResultType][tt.str2]
			s1, s2 := res.GetStrings()
			if res.GetType() != tt.lcsResultType ||
				res.GetTypeName() != tt.lcsResultType.String() ||
				res.GetString1() != tt.str1 ||
				res.GetString2() != tt.str2 ||
				s1 != tt.str1 ||
				s2 != tt.str2 ||
				!errors.CompareErrors(res.GetError(), tt.err) ||
				!comparison.CompareStringSlices(res.GetResult(), tt.result, false) {
				t.Errorf("LCSResultGetters() = %v, want %v", res.GetResult(), tt.result)
			}
		})
	}
}

func TestLCSResultIsMatch(t *testing.T) {
	tests := []struct {
		name string
		res1 *LCSResult
		res2 *LCSResult
		want bool
	}{
		{"LCSIsMatch1", bt1, bt2, true},
		{"LCSIsMatch2", bt1, bta2, false},
		{"LCSIsMatch3", bt1, ld2, false},
		{"LCSIsMatch4", bta1, bt2, false},
		{"LCSIsMatch5", bta1, bta2, true},
		{"LCSIsMatch6", bta1, ld2, false},
		{"LCSIsMatch7", bt1, bt3, false},
		{"LCSIsMatch8", bt1, bt1, true},
		{"LCSIsMatch9", ld3, ld2, false},
		{"LCSIsMatch10", ld1, ld2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.res1.IsMatch(tt.res2); got != tt.want {
				t.Errorf("LCSResultIsMatch() = %t, want %t", got, tt.want)
			}
		})
	}
}

func TestLCSResultFormatOutput(t *testing.T) {
	tests := []struct {
		name      string
		lcsResult *LCSResult
		verbose   bool
		expected  string
	}{
		{"BacktackWord", bt1, true, btp1Verb},
		{"BacktackWordAll", bt1, false, btp1Sh},
		{"BacktackDiff", bta1, true, btap1Verb},
		{"BacktackDiffAll", bta1, false, btap1Sh},
		{"LCSDiff", ld1, true, ldp1Verb},
		{"LCSDiffAll", ld1, false, ldp1Sh},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if formatLCSResultOutput(tt.lcsResult, tt.verbose) != tt.expected {
				t.Errorf("LCSResultFormatOutput() = %s, want %s",
					formatLCSResultOutput(tt.lcsResult, tt.verbose), tt.expected)
			}
			tt.lcsResult.Print(tt.verbose)
		})
	}
}
