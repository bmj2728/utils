package strutil

import "testing"

func TestIndexType_String(t *testing.T) {
	tests := []struct {
		name string
		i    IndexType
		want string
	}{
		{"IndexType1", Byte, "Byte"},
		{"IndexType2", ByteNorm, "ByteNorm"},
		{"IndexType3", Rune, "Rune"},
		{"IndexType4", UTF32Norm, "UTF32Norm"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("IndexType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStringIndex(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		substr string
		i      IndexType
		want   int
	}{
		{"Index1", normalizeUnicode("Café Gopher", NFD), "Gopher", Byte, 7},
		{"Index2", normalizeUnicode("Café Gopher", NFD), "Gopher", ByteNorm, 6},
		{"Index3", normalizeUnicode("Café Gopher", NFD), "Gopher", UTF32Norm, 5},
		{"Index5", "Café Gopher", "Gopher", Byte, 6},
		{"Index6", "Café Gopher", "Gopher", ByteNorm, 6},
		{"Index7", "Café Gopher", "Gopher", UTF32Norm, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ix int
			switch tt.i {
			case Byte:
				ix = Index(tt.s, tt.substr, false)
			case ByteNorm:
				ix = Index(tt.s, tt.substr, true)
			case UTF32Norm:
				ix = UTF32NormalizedIndex(tt.s, tt.substr)
			default:
				ix = -1
			}
			idx := NewStringIndex(tt.s, tt.substr, tt.i, ix)
			if idx.GetString() != tt.s {
				t.Errorf("NewStringIndex.GetString() = %s, want %s",
					idx.GetString(), tt.s)
			}
			if idx.GetIndex() != tt.want {
				t.Errorf("NewStringIndex.GetIndex() = %d, want %d",
					idx.GetIndex(), tt.want)
			}
			if idx.GetRune() != 0 {
				t.Errorf("NewStringIndex.GetIndex() = %d, want %d",
					idx.GetIndex(), ix)
			}
			if idx.GetIndexType() != tt.i {
				t.Errorf("NewStringIndex.GetIndexType() = %s, want %s",
					idx.GetIndexType(), tt.i)
			}
			if idx.GetIndexType().String() != tt.i.String() {
				t.Errorf("NewStringIndex.GetIndexType().String() = %s, want %s",
					idx.GetIndexType().String(), tt.i.String())
			}
			if idx.GetSubstr() != tt.substr {
				t.Errorf("NewStringIndex.GetIndex() = %d, want %d",
					idx.GetIndex(), tt.want)
			}
		})
	}
}

func TestNewStringRuneIndex(t *testing.T) {
	tests := []struct {
		name string
		s    string
		r    rune
		want int
	}{
		{"Index1", normalizeUnicode("Café Gopher", NFD), 'é', -1},
		{"Index2", "Café Gopher", 'é', 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := NewStringRuneIndex(tt.s, tt.r, IndexRune(tt.s, tt.r))
			if idx.GetString() != tt.s {
				t.Errorf("NewStringIndex.GetString() = %s, want %s",
					idx.GetString(), tt.s)
			}
			if idx.GetSubstr() != string(tt.r) {
				t.Errorf("NewStringIndex.GetSubstr() = %s, want %s",
					idx.GetSubstr(), string(tt.r))
			}
			if idx.GetRune() != tt.r {
				t.Errorf("NewStringIndex.GetRune() = %v, want %v",
					idx.GetRune(), tt.r)
			}
			if idx.GetIndex() != tt.want {
				t.Errorf("NewStringIndex.GetIndex() = %d, want %d",
					idx.GetIndex(), tt.want)
			}
			if idx.GetIndexType() != Rune {
				t.Errorf("NewStringIndex.GetIndexType() = %s, want %s",
					idx.GetIndexType(), Rune)
			}
			if idx.GetIndexType().String() != Rune.String() {
				t.Errorf("NewStringIndex.GetIndexType().String() = %s, want %s",
					idx.GetIndexType().String(), Rune.String())
			}
			notRune := NewStringIndex(tt.s, "hi", Byte, 0)
			if notRune.GetRune() != 0 {
				t.Errorf("NewStringIndex.GetRune() = %v, want %v",
					notRune.GetRune(), 0)
			}
		})
	}
}

var (
	idx1 = NewStringIndex("Hello World", "World", Byte, 6)
	idx2 = NewStringIndex("Hello World", "World", ByteNorm, 6)
	idx3 = NewStringIndex("Hello World", "World", UTF32Norm, 6)
	idx4 = NewStringIndex("Hello World", "World", Byte, 6)
	idx5 = NewStringIndex("Hi World", "World", ByteNorm, 6)
	idx6 = NewStringIndex("Hello World", "World", UTF32Norm, 6)
	idx7 = NewStringRuneIndex("Hello World", 'o', 5)
	idx8 = NewStringRuneIndex("Hello World", 'o', 5)
	idx9 *StringIndex
)

func TestIndexIsMatch(t *testing.T) {
	tests := []struct {
		name      string
		idx1      *StringIndex
		idx2      *StringIndex
		typeMatch bool
		expected  bool
	}{
		{"IndexIsMatch1", idx1, idx2, false, true},
		{"IndexIsMatch2", idx1, idx2, true, false},
		{"IndexIsMatch3", idx1, idx4, true, true},
		{"IndexIsMatch4", idx1, idx5, true, false},
		{"IndexIsMatch5", idx1, idx7, true, false},
		{"IndexIsMatch6", idx1, idx8, false, false},
		{"IndexIsMatch7", idx3, idx6, false, true},
		{"IndexIsMatch8", idx7, idx8, true, true},
		{"IndexIsMatch9", idx1, idx9, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.idx1.IsMatch(tt.idx2, tt.typeMatch) != tt.expected {
				t.Errorf("IndexIsMatch.IsMatch() = %v, want %v",
					tt.idx1.IsMatch(tt.idx2, tt.typeMatch), tt.expected)
			}
		})
	}
}

var (
	shortOut = "Index Type: Byte, Index: 6\n"
	longOut  = "String: Hello World\nSubstring: World\nIndexType: Byte\nIndex: 6\n"
)

func TestStringIndexFormatOutput(t *testing.T) {
	input := "Hello World"
	substr := "World"
	idx := NewStringIndex(input, substr, Byte, Index(input, substr, true))
	shS := formatStringIndexOutput(idx, false)
	shL := formatStringIndexOutput(idx, true)
	var bad *StringIndex
	badFormat := formatStringIndexOutput(bad, true)
	if shS != shortOut {
		t.Errorf("StringIndexFormatOutput.Short() = %s, want %s",
			shS, shortOut)
	}
	if shL != longOut {
		t.Errorf("StringIndexFormatOutput.Long() = %s, want %s",
			shL, longOut)
	}
	if badFormat != "" {
		t.Errorf("StringIndexFormatOutput.Bad() = %s, want %s",
			badFormat, "")
	}
	idx.Print(true)
	idx.Print(false)
}
