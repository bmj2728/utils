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
	s := AppendString("v0.1.0", "alphaRemove", "-")
	fmt.Printf("%s", s)
	// Output: v0.1.0-alphaRemove
}

func ExampleStringBuilder_Append() {
	New("v0.1.0").Append("alphaRemove", "-").Print()
	// Output: v0.1.0-alphaRemove
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

func ExampleCleanWhitespace() {
	s := `Hello 
		  World
		  !`
	clean := CleanWhitespace(s)
	fmt.Printf("%s", clean)
	// Output: HelloWorld!
}

func ExampleStringBuilder_CleanWhitespace() {
	s := `Hello 
		  World
		  !`
	New(s).CleanWhitespace().Print()
	// Output: HelloWorld!
}

func ExampleComparisonResultsMap_Print() {
	New("Hello, World!").
		WithComparisonManager().
		LevenshteinDistance("Hello World").
		DamerauLevenshteinDistance("Hello World").
		ComparisonManager().
		GetComparisonResultsMap().
		Print(true)
	// Output: ***Comparison Results for Levenshtein Distance***
	//Comparison: Levenshtein Distance
	//First String: Hello, World!
	//Second String: Hello World
	//Score: 2
	//***Comparison Results for Damerau-Levenshtein Distance***
	//Comparison: Damerau-Levenshtein Distance
	//First String: Hello, World!
	//Second String: Hello World
	//Score: 2
}
