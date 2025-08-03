package strutil

import "testing"

func TestStringBuilder_RevertToOriginal(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					WithHistory().
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello World 4
					TrimCharsRight(" World"). //Hello 5
					Append("John", " ")

	if rb.RevertToOriginal().String() != "Hello World" {
		t.Errorf("RevertToOriginal failed, expected Hello World, got %s", rb.RevertToOriginal().String())
	}
}

func TestBuilderRevertToPrevious(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					WithHistory().
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello-World 4
					TrimCharsRight(" World"). //Hello-5
					Append("John", " ")       //Hello- John

	if rb.String() != "Hello- John" {
		t.Errorf("RevertToPrevious failed, expected Hello- John, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "Hello-" {
		t.Errorf("RevertToPrevious failed, expected Hello-, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "Hello-World" {
		t.Errorf("RevertToPrevious failed, expected Hello-World, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "hello-world" {
		t.Errorf("RevertToPrevious failed, expected hello-world, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "HELLO-WORLD" {
		t.Errorf("RevertToPrevious failed, expected HELLO-WORLD, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "helloWorld" {
		t.Errorf("RevertToPrevious failed, expected helloWorld, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "Hello World" {
		t.Errorf("RevertToPrevious failed, expected Hello World, got %s", rb.String())
	}
	if rb.RevertToPrevious().String() != "" || !compareErrors(rb.Error(), ErrInvalidHistoryIndex) {
		t.Errorf("RevertToPrevious failed, expected empty string, got %s", rb.String())
	}
}

func TestBuilderRevertToIndex(t *testing.T) {
	tests := []struct {
		name     string
		index    int
		expected string
	}{
		{"Hist1", 1, "helloWorld"},
		{"Hist2", 2, "HELLO-WORLD"},
		{"Hist3", 3, "hello-world"},
		{"Hist4", 4, "Hello-World"},
		{"Hist5", 5, "Hello-"},
		{"Hist6", 6, "Hello- John"},
		{"Hist7", 7, ""},
		{"Hist8", 0, "Hello World"},
		{"Hist9", -1, ""},
		{"Hist10", -2, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := New("Hello World"). //Hello World 0
							WithHistory().
							ToCamelCase().            //helloWorld 1
							ToKebabCase(true).        //HELLO-WORLD 2
							ToKebabCase(false).       //hello-world 3
							ToTitleCase().            //Hello-World 4
							TrimCharsRight(" World"). //Hello-5
							Append("John", " ")
			if rb.RevertToIndex(tt.index).String() != tt.expected {
				t.Errorf("RevertToIndex(%d) failed, expected %s, got %s",
					tt.index, tt.expected, rb.RevertToIndex(tt.index).String())
			}
		})
	}
}

func TestRevertByFunc(t *testing.T) {
	tests := []struct {
		name     string
		offset   int
		expected string
	}{
		{"Hist1", 6, "helloWorld"},
		{"Hist2", 5, "HELLO-WORLD"},
		{"Hist3", 4, "hello-world"},
		{"Hist4", 3, "Hello-World"},
		{"Hist5", 2, "Hello-"},
		{"Hist6", 1, "Hello- John"},
		{"Hist7", 11, ""},
		{"Hist8", 0, ""},
		{"Hist9", -1, ""},
		{"Hist10", -2, ""},
		{"Hist11", -6, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := New("Hello World"). //Hello World 0
							WithHistory().
							ToCamelCase().            //helloWorld 1
							ToKebabCase(true).        //HELLO-WORLD 2
							ToKebabCase(false).       //hello-world 3
							ToTitleCase().            //Hello-World 4
							TrimCharsRight(" World"). //Hello-5
							Append("John", " ")       // 6

			revStr := rb.RevertWithFunction(func(history *StringHistory) int {
				return history.Len() - tt.offset
			}).String()
			if revStr != tt.expected {
				t.Errorf("RevertWithFunction failed, expected %s, got %s", tt.expected, revStr)
			}
		})
	}
}

func TestRevertNoHistoryOrig(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello-World 4
					TrimCharsRight(" World"). //Hello-5
					Append("John", " ")       // 6

	str := rb.RevertToOriginal().String()
	err := rb.Error()
	if str != "Hello- John" || !compareErrors(err, ErrHistoryNotInitialized) {
		t.Errorf("RevertToPrevious failed, expected empty string, got %s", str)
	}
}

func TestRevertNoHistoryPrev(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello-World 4
					TrimCharsRight(" World"). //Hello-5
					Append("John", " ")       // 6

	str := rb.RevertToPrevious().String()
	err := rb.Error()
	if str != "Hello- John" || !compareErrors(err, ErrHistoryNotInitialized) {
		t.Errorf("RevertToPrevious failed, expected empty string, got %s", str)
	}
}

func TestRevertNoHistoryIndex(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello-World 4
					TrimCharsRight(" World"). //Hello-5
					Append("John", " ")       // 6

	str := rb.RevertToIndex(1).String()
	err := rb.Error()
	if str != "Hello- John" || !compareErrors(err, ErrHistoryNotInitialized) {
		t.Errorf("RevertToPrevious failed, expected empty string, got %s", str)
	}
}

func TestRevertNoHistoryFunc(t *testing.T) {
	rb := New("Hello World"). //Hello World 0
					ToCamelCase().            //helloWorld 1
					ToKebabCase(true).        //HELLO-WORLD 2
					ToKebabCase(false).       //hello-world 3
					ToTitleCase().            //Hello-World 4
					TrimCharsRight(" World"). //Hello-5
					Append("John", " ")       // 6

	str := rb.RevertWithFunction(func(history *StringHistory) int {
		return len(*rb.GetHistory()) - 2
	}).String()
	err := rb.Error()
	if str != "Hello- John" || !compareErrors(err, ErrHistoryNotInitialized) {
		t.Errorf("RevertToPrevious failed, expected empty string, got %s", str)
	}
}

func TestMissingOrig(t *testing.T) {
	s := New("Hello World")
	s.history = &StringHistory{}
	s.RevertToOriginal()
	str := s.String()
	err := s.Error()
	if str != "" || !compareErrors(err, ErrHistoryIsEmpty) {
		t.Errorf("RevertToPrevious failed, expected %q, got %q", "", str)
		t.Errorf("RevertToPrevious failed, expected %s, got %s", ErrHistoryIsEmpty, err)
	}
}
