package strutil

import "testing"

var (
	lcsMap1 = New("With great power there must also come great responsibility.").
		WithComparisonManager().
		LCSBacktrack("With great power comes great responsibility.").
		LCSBacktrackAll("With great power comes great responsibility.").
		LCSDiff("With great power comes great responsibility.").
		GetComparisonManager().
		GetLCSResultsMap()

	lcsMap2 = New("With great power there must also come great responsibility.").
		WithComparisonManager().
		LCSBacktrack("With great power comes great responsibility.").
		LCSBacktrackAll("With great power comes great responsibility.").
		LCSDiff("With great power comes great responsibility.").
		GetComparisonManager().
		GetLCSResultsMap()

	lcsMap3 = New("Hello, World!").
		WithComparisonManager().
		LCSBacktrack("Hello World").
		LCSBacktrack("Hello Whirled").
		LCSBacktrackAll("Hello World").
		LCSDiff("Hello World").
		GetComparisonManager().
		GetLCSResultsMap()

	lcsMap4 = New("With great power there must also come great responsibility.").
		WithComparisonManager().
		LCSBacktrack("With great power comes great responsibility.").
		LCSDiff("With great power comes great responsibility.").
		GetComparisonManager().
		GetLCSResultsMap()

	lcsMap5 = New("With great power there must also come great responsibility.").
		WithComparisonManager().
		GetComparisonManager().
		GetLCSResultsMap()

	lcsMapPrintLongOpt1 = "***LCS Results for LCS Backtrack***\n\n" +
		"LCS Backtrack:\n" +
		"First: With great power there must also come great responsibility.\n" +
		"Second: With great power comes great responsibility.\n" +
		"Word: With great power come great responsibility.\n\n" +
		"***LCS Results for LCS Diff***\n\n" +
		"LCS Diff (With great power there must also come great responsibility./With great power comes great" +
		" responsibility.):\n W i t h   g r e a t   p o w e r   t h e r e   m u s t   a l s o   c o m e s   g r e a t   " +
		"r e s p o n s i b i l i t y .\n                             - - - - -     - - - - - - - - - - -           +" +
		"                                            \n\n"
	lcsMapPrintLongOpt2 = "***LCS Results for LCS Diff***\n\n" +
		"LCS Diff (With great power there must also come great responsibility./With great power" +
		" comes great responsibility.):\n W i t h   g r e a t   p o w e r   t h e r e   m u s t   a l s o   " +
		"c o m e s   g r e a t   r e s p o n s i b i l i t y .\n" +
		"                             - - - - -     - - - - - - - - - - -           +" +
		"                                            \n\n" +
		"***LCS Results for LCS Backtrack***\n\nLCS Backtrack:\n" +
		"First: With great power there must also come great responsibility.\n" +
		"Second: With great power comes great responsibility.\n" +
		"Word: With great power come great responsibility.\n\n"
	lcsMapPrintShortOpt1 = "***LCS Results for LCS Backtrack***\n\n" +
		"LCS Backtrack: With great power come great responsibility.\n\n" +
		"***LCS Results for LCS Diff***\n\n" +
		"LCS Diff:\n" +
		" W i t h   g r e a t   p o w e r   t h e r e   m u s t   a l s o   c o m e s   g r e a t   " +
		"r e s p o n s i b i l i t y .\n                             - - - - -     - - - - - - - - - - -           +" +
		"                                            \n\n"
	lcsMapPrintShortOpt2 = "***LCS Results for LCS Diff***\n\n" +
		"LCS Diff:\n" +
		" W i t h   g r e a t   p o w e r   t h e r e   m u s t   a l s o   c o m e s   g r e a t   " +
		"r e s p o n s i b i l i t y .\n                             - - - - -     - - - - - - - - - - -           +" +
		"                                            \n\n" +
		"***LCS Results for LCS Backtrack***\n\n" +
		"LCS Backtrack: With great power come great responsibility.\n\n"
)

func TestLCSMapTypeCount(t *testing.T) {
	tests := []struct {
		name   string
		lcsMap LCSResultsMap
		want   int
	}{
		{"LCSMapTypeCount1", lcsMap1, 3},
		{"LCSMapTypeCount2", lcsMap2, 3},
		{"LCSMapTypeCount3", lcsMap3, 3},
		{"LCSMapTypeCount4", lcsMap4, 2},
		{"LCSMapTypeCount5", lcsMap5, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lcsMap.TypeCount(); got != tt.want {
				t.Errorf("LCSMapTypeCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCSMapEntryCount(t *testing.T) {
	tests := []struct {
		name   string
		lcsMap LCSResultsMap
		want   int
	}{
		{"LCSMapEntryCount1", lcsMap1, 3},
		{"LCSMapEntryCount2", lcsMap2, 3},
		{"LCSMapEntryCount3", lcsMap3, 4},
		{"LCSMapEntryCount4", lcsMap4, 2},
		{"LCSMapEntryCount5", lcsMap5, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lcsMap.EntryCount(); got != tt.want {
				t.Errorf("LCSMapEntryCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCSMapIsMatch(t *testing.T) {
	tests := []struct {
		name    string
		lcsMap1 LCSResultsMap
		lcsMap2 LCSResultsMap
		want    bool
	}{
		{"LCSMapIsMatch1", lcsMap1, lcsMap1, true},
		{"LCSMapIsMatch2", lcsMap1, lcsMap2, true},
		{"LCSMapIsMatch3", lcsMap1, lcsMap3, false},
		{"LCSMapIsMatch4", lcsMap1, lcsMap4, false},
		{"LCSMapIsMatch5", lcsMap1, lcsMap5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lcsMap1.IsMatch(tt.lcsMap2); got != tt.want {
				t.Errorf("LCSMapIsMatch() = %t, want %t", got, tt.want)
			}
		})
	}
}

func TestLCSMapGetCopy(t *testing.T) {
	tests := []struct {
		name     string
		lcsMap   LCSResultsMap
		expected bool
	}{
		{"LCSMapGetCopy1", lcsMap1, true},
		{"LCSMapGetCopy2", lcsMap2, true},
		{"LCSMapGetCopy3", lcsMap3, true},
		{"LCSMapGetCopy4", lcsMap4, true},
		{"LCSMapGetCopy5", lcsMap5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dup := tt.lcsMap.GetCopy()
			if dup.TypeCount() != tt.lcsMap.TypeCount() {
				t.Errorf("LCSMapGetCopy() = %v, want %v", dup.TypeCount(), tt.lcsMap.TypeCount())
			}
			if dup.EntryCount() != tt.lcsMap.EntryCount() {
				t.Errorf("LCSMapGetCopy() = %v, want %v", dup.EntryCount(), tt.lcsMap.EntryCount())
			}
			if tt.expected && !tt.lcsMap.IsMatch(dup) {
				t.Errorf("LCSMapGetCopy() = %v, want %v", dup.IsMatch(tt.lcsMap), true)
			}
		})
	}
}

func TestLCSMapGet(t *testing.T) {
	tests := []struct {
		name     string
		lcsMap   LCSResultsMap
		lcsType  LCSResultType
		compStr  string
		expected bool
	}{
		{"LCSMapGet1", lcsMap1, LCSBacktrackWord,
			"With great power comes great responsibility.", true},
		{"LCSMapGet2", lcsMap2, LCSBacktrackWord,
			"With great power comes great responsibility.", true},
		{"LCSMapGet3", lcsMap1, LCSBacktrackWord,
			"Hi", false},
		{"LCSMapGet4", lcsMap3, LCSBacktrackWord,
			"With great power comes great responsibility.", false},
		{"LCSMapGet5", lcsMap3, LCSBacktrackWord,
			"Hello World", true},
		{"LCSMapGet6", lcsMap3, LCSBacktrackWord,
			"Hello Whirled", true},
		{"LCSMapGet7", lcsMap4, LCSBacktrackWord,
			"With great power comes great responsibility.", true},
		{"LCSMapGet8", lcsMap4, LCSDiffSlice,
			"With great power comes great responsibility.", true},
		{"LCSMapGet9", lcsMap5, LCSBacktrackWord,
			"With great power comes great responsibility.", false},
		{"LCSMapGet10", lcsMap5, LCSBacktrackWordAll,
			"With great power comes great responsibility.", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (tt.lcsMap.Get(tt.lcsType, tt.compStr) != nil) != tt.expected {
				t.Errorf("LCSMapGet() = %v, want %v", tt.lcsMap.Get(tt.lcsType, tt.compStr) != nil, tt.expected)
			}
		})
	}
}

func TestLCSMapGetByType(t *testing.T) {
	tests := []struct {
		name     string
		lcsMap   LCSResultsMap
		lcsType  LCSResultType
		expected int
	}{
		{"LCSMapGetByType1", lcsMap1, LCSBacktrackWord, 1},
		{"LCSMapGetByType2", lcsMap2, LCSBacktrackWord, 1},
		{"LCSMapGetByType3", lcsMap3, LCSBacktrackWord, 2},
		{"LCSMapGetByType4", lcsMap4, LCSBacktrackWord, 1},
		{"LCSMapGetByType5", lcsMap5, LCSBacktrackWord, 0},
		{"LCSMapGetByType6", lcsMap1, LCSBacktrackWordAll, 1},
		{"LCSMapGetByType7", lcsMap2, LCSDiffSlice, 1},
		{"LCSMapGetByType8", lcsMap3, LCSBacktrackWordAll, 1},
		{"LCSMapGetByType9", lcsMap4, LCSBacktrackWordAll, 0},
		{"LCSMapGetByType10", lcsMap5, LCSDiffSlice, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.lcsMap.GetByType(tt.lcsType)
			if len(result) != tt.expected {
				t.Errorf("LCSMapGetByType() = %v, want %v", len(result), tt.expected)
			}
		})
	}
}

func TestLCSMapFilterByType(t *testing.T) {
	tests := []struct {
		name     string
		lcsMap   LCSResultsMap
		lcsType  LCSResultType
		expected int
	}{
		{"LCSMapFilterByType1", lcsMap1, LCSBacktrackWord, 1},
		{"LCSMapFilterByType2", lcsMap2, LCSBacktrackWord, 1},
		{"LCSMapFilterByType3", lcsMap3, LCSBacktrackWord, 2},
		{"LCSMapFilterByType4", lcsMap4, LCSBacktrackWord, 1},
		{"LCSMapFilterByType5", lcsMap5, LCSBacktrackWord, 0},
		{"LCSMapFilterByType6", lcsMap1, LCSBacktrackWordAll, 1},
		{"LCSMapFilterByType7", lcsMap2, LCSDiffSlice, 1},
		{"LCSMapFilterByType8", lcsMap3, LCSBacktrackWordAll, 1},
		{"LCSMapFilterByType9", lcsMap4, LCSBacktrackWordAll, 0},
		{"LCSMapFilterByType10", lcsMap5, LCSDiffSlice, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.lcsMap.FilterByType(tt.lcsType)
			if result.EntryCount() != tt.expected {
				t.Errorf("LCSMapFilterByType() = %v, want %v", len(result), tt.expected)
			}
		})
	}
}

func TestLCSMapGetByCompStr(t *testing.T) {
	tests := []struct {
		name     string
		lcsMap   LCSResultsMap
		compStr  string
		expected int
	}{
		{"LCSMapGetByCompStr1", lcsMap1, "With great power comes great responsibility.", 3},
		{"LCSMapGetByCompStr2", lcsMap2, "With great power comes great responsibility.", 3},
		{"LCSMapGetByCompStr3", lcsMap3, "With great power comes great responsibility.", 0},
		{"LCSMapGetByCompStr4", lcsMap4, "With great power comes great responsibility.", 2},
		{"LCSMapGetByCompStr5", lcsMap5, "With great power comes great responsibility.", 0},
		{"LCSMapGetByCompStr6", lcsMap1, "Hi", 0},
		{"LCSMapGetByCompStr7", lcsMap3, "With great power comes great responsibility.", 0},
		{"LCSMapGetByCompStr8", lcsMap3, "Hello World", 3},
		{"LCSMapGetByCompStr9", lcsMap3, "Hello Whirled", 1},
		{"LCSMapGetByCompStr10", lcsMap4, "With great power comes great responsibility.", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.lcsMap.GetByComparisonString(tt.compStr)
			if len(res) != tt.expected {
				t.Errorf("LCSMapGetByCompStr() = %v, want %v", len(res), tt.expected)
			}
		})
	}
}

func TestLCSMapFilterByCompStr(t *testing.T) {
	tests := []struct {
		name     string
		lcsMap   LCSResultsMap
		compStr  string
		expected int
	}{
		{"LCSMapFilterByCompStr1", lcsMap1, "With great power comes great responsibility.", 3},
		{"LCSMapFilterByCompStr2", lcsMap2, "With great power comes great responsibility.", 3},
		{"LCSMapFilterByCompStr3", lcsMap3, "With great power comes great responsibility.", 0},
		{"LCSMapFilterByCompStr4", lcsMap4, "With great power comes great responsibility.", 2},
		{"LCSMapFilterByCompStr5", lcsMap5, "With great power comes great responsibility.", 0},
		{"LCSMapFilterByCompStr6", lcsMap1, "Hi", 0},
		{"LCSMapFilterByCompStr7", lcsMap3, "With great power comes great responsibility.", 0},
		{"LCSMapFilterByCompStr8", lcsMap3, "Hello World", 3},
		{"LCSMapFilterByCompStr9", lcsMap3, "Hello Whirled", 1},
		{"LCSMapFilterByCompStr10", lcsMap4, "With great power comes great responsibility.", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.lcsMap.FilterByComparisonString(tt.compStr)
			if res.EntryCount() != tt.expected {
				t.Errorf("LCSMapFilterByCompStr() = %v, want %v", res.EntryCount(), tt.expected)
			}
		})
	}
}

func TestLCSMapPrint(t *testing.T) {
	tests := []struct {
		name     string
		lcsMap   LCSResultsMap
		expected bool
	}{
		{"LCSMapPrint1", lcsMap4, true},
		{"LCSMapPrint2", lcsMap2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var isShortMatch bool
			var isLongMatch bool
			resShort := formatLCSResultsMapOutput(tt.lcsMap, false)
			resLong := formatLCSResultsMapOutput(tt.lcsMap, true)
			if resShort == lcsMapPrintShortOpt1 || resShort == lcsMapPrintShortOpt2 {
				isShortMatch = true
			}
			if resLong == lcsMapPrintLongOpt1 || resLong == lcsMapPrintLongOpt2 {
				isLongMatch = true
			}
			if isShortMatch != tt.expected {
				t.Errorf("LCSMapPrint() = %v, want %v", isShortMatch, tt.expected)
			}
			if isLongMatch != tt.expected {
				t.Errorf("LCSMapPrint() = %v, want %v", isLongMatch, tt.expected)
			}
			tt.lcsMap.Print(true)
			tt.lcsMap.Print(false)
		})
	}
}
