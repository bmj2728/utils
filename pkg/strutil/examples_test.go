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
	slug := Slugify("The Life and Strange Surprising Adventures of Robinson Crusoe, Of York, Mariner", 50)
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
	s := AppendString("v0.1.0", "removeNonAlpha", "-")
	fmt.Printf("%s", s)
	// Output: v0.1.0-removeNonAlpha
}

func ExampleStringBuilder_Append() {
	New("v0.1.0").Append("removeNonAlpha", "-").Print()
	// Output: v0.1.0-removeNonAlpha
}

func ExamplePrependString() {
	s := PrependString("ENV_VAR", "APP", "_")
	fmt.Printf("%s", s)
	// Output: APP_ENV_VAR
}

func ExampleStringBuilder_Prepend() {
	New("ENV_VAR").Prepend("APP", "_").Print()
	// Output: APP_ENV_VAR
}
