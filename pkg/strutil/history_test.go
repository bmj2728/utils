package strutil

import (
	"testing"

	"utils/pkg/internal/errors"
)

func TestNewStringHistory(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Hist1", "Hello World"},
		{"Hist2", "Hello"},
		{"Hist3", "Hi"},
		{"Hist4", "World"},
		{"Hist5", "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hist := New(tt.want).WithHistory(10).GetHistory()
			if (*hist).transforms[0] != tt.want {
				t.Errorf("NewStringHistory() = %v, want %v", (*hist).transforms[0], tt.want)
			}
		})
	}
}

func TestHistoryLength(t *testing.T) {
	tests := []struct {
		name  string
		loops int
		want  int
	}{
		{"Hist1", 1, 2},
		{"Hist2", 2, 3},
		{"Hist3", 3, 4},
		{"Hist4", 4, 5},
		{"Hist5", 5, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := New("Hello World").WithHistory(10)
			i := tt.loops
			for i > 0 {
				str.Append("loop", "-")
				i--
			}
			if str.GetHistory().Len() != tt.want {
				t.Errorf("HistoryLength() = %v, want %v", str.GetHistory().Len(), tt.want)
			}
		})
	}
}

func TestHistoryLengthNil(t *testing.T) {
	tests := []struct {
		name  string
		loops int
		want  int
	}{
		{"Hist1", 1, 0},
		{"Hist2", 2, 0},
		{"Hist3", 3, 0},
		{"Hist4", 4, 0},
		{"Hist5", 5, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := New("Hello World")
			i := tt.loops
			for i > 0 {
				str.Append("loop", "-")
				i--
			}
			if str.GetHistory().Len() != tt.want {
				t.Errorf("HistoryLength() = %v, want %v", str.GetHistory().Len(), tt.want)
			}
		})
	}
}

func TestFormatHistory(t *testing.T) {
	long := "\nHistory: \n1: Hello, World!\n2: hello-world\n3: HELLO-WORLD\n4: hello-world\n"
	short := "\nHistory: \nHello, World!, hello-world, HELLO-WORLD, hello-world\n"

	x := New("Hello, World!").
		WithHistory(10).
		Slugify(25).
		ToUpper().
		ToLower().
		GetHistory()
	longOut := formatHistoryOutput(*x, true)
	shortOut := formatHistoryOutput(*x, false)
	if longOut != long {
		t.Errorf("FormatHistory() = %s, want %s", longOut, long)
	}
	if shortOut != short {
		t.Errorf("FormatHistory() = %s, want %s", shortOut, short)
	}
	x.Print(true)
	x.Print(false)
}

func TestHistoryGetNil(t *testing.T) {
	empty := &StringHistory{}
	prev, err2 := empty.GetPreviousValue()
	if !errors.CompareErrors(err2, errors.ErrInvalidHistoryIndex) {
		t.Errorf("HistoryGetPrevious() = %s, want %s", err2.Error(), errors.ErrInvalidHistoryIndex.Error())
	}
	if prev != "" {
		t.Errorf("HistoryGetPrevious() = %s, want %s", prev, "")
	}
	ind, err3 := empty.GetByIndex(1)
	if !errors.CompareErrors(err3, errors.ErrInvalidHistoryIndex) {
		t.Errorf("HistoryGetByIndex(1) = %s, want %s", err3.Error(), errors.ErrInvalidHistoryIndex.Error())
	}
	if ind != "" {
		t.Errorf("HistoryGetByIndex(1) = %s, want %s", ind, "")
	}
	ind2, err4 := empty.GetByIndex(-1)
	if !errors.CompareErrors(err4, errors.ErrInvalidHistoryIndex) {
		t.Errorf("HistoryGetByIndex(-1) = %s, want %s", err4.Error(), errors.ErrInvalidHistoryIndex.Error())
	}
	if ind2 != "" {
		t.Errorf("HistoryGetByIndex(-1) = %s, want %s", ind, "")
	}
}

func TestHistoryGetPrevious(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected1 string
	}{
		{"Hist1", "Hello World", "hello-world"},
		{"Hist2", "With great power there must also come great responsibility",
			"with-great-power-there-mu"},
		{"Hist3", "Hi", "hi"},
		{"Hist4", "This is the way!", "this-is-the-way"},
		{"Hist5", "Hello, World!", "hello-world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := New(tt.input).
				WithHistory(10).
				Slugify(25).
				ToLower().
				ToUpper().
				GetHistory()
			prev, err := x.GetPreviousValue()
			if err != nil {
				t.Errorf("HistoryGetOriginal() = %v, want %v", err, nil)
			}
			if prev != tt.expected1 {
				t.Errorf("HistoryGetOriginal() = %v, want %v", prev, tt.expected1)
			}
		})
	}
}

func TestHistoryGetByIndex(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		index     int
		expected1 string
	}{
		{"Hist1", "Hello World", 1, "hello-world"},
		{"Hist2", "Hello World", 3, "hello-world"},
		{"Hist3", "Hello World", 0, "Hello World"},
		{"Hist4", "Hi", 2, "HI"},
		{"Hist5", "This is the way!", 1, "this-is-the-way"},
		{"Hist6", "Hello, World!", 3, "hello-world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := New(tt.input).
				WithHistory(10).
				Slugify(25).
				ToUpper().
				ToLower().
				GetHistory()
			ind, err := x.GetByIndex(tt.index)
			if err != nil {
				t.Errorf("HistoryGetOriginal() = %v, want %v", err, nil)
			}
			if ind != tt.expected1 {
				t.Errorf("HistoryGetOriginal() = %v, want %v", ind, tt.expected1)
			}
		})
	}
}
