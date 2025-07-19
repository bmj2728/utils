package strutil

import "fmt"

func ExampleNew() {
	sb := New("Hello World!")
	fmt.Printf("%q", sb.String())
	// Output: "Hello World!"
}

func ExampleStringBuilder_Print() {
	New("Hello World!").Print()
	// Output: "Hello World!"
}

func ExampleSlugify() {
	original := "The Life and Strange Surprising Adventures of Robinson Crusoe, Of York, Mariner"
	slug := Slugify(original, 50)
	fmt.Println(slug)
	// Output: the-life-and-strange-surprising-adventures-of-robi
}

func ExampleStringBuilder_Slugify() {
	New("The Life and Strange Surprising Adventures of Robinson Crusoe, Of York, Mariner").
		Slugify(50).
		Print()
	// Output: the-life-and-strange-surprising-adventures-of-robi
}
