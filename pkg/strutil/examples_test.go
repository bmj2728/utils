package strutil

import "fmt"

func ExampleNew() {
	sb := New("Hello World!")
	fmt.Printf("%q", sb.String())
	// Output: "Hello World!"
}

func ExampleStringBuilder_Print() {
	New("Hello World!").Print()
	// Output: Hello World!
}

func ExampleSlugify() {
	original := "The Life and Strange Surprising Adventures of Robinson Crusoe, Of York, Mariner"
	slug := Slugify(original, 50)
	fmt.Printf("%s", slug)
	// Output: the-life-and-strange-surprising-adventures-of-robi
}

func ExampleStringBuilder_Slugify() {
	New("The Life and Strange Surprising Adventures of Robinson Crusoe, Of York, Mariner").
		Slugify(50).
		Print()
	// Output: the-life-and-strange-surprising-adventures-of-robi
}

func ExampleAppendString() {
	s := "v0.1.0"
	suffix := "alpha"
	separator := "-"
	s = AppendString(s, suffix, separator)
	fmt.Printf("%s", s)
	// Output: v0.1.0-alpha
}

func ExampleStringBuilder_Append() {
	New("v0.1.0").Append("alpha", "-").Print()
	// Output: v0.1.0-alpha
}
