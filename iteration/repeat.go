package iteration

// Repeat concatenates the given character string N times and returns the final string.
//
// This function takes a string `character` and an integer `cycles`. It will
// repeat the `character` string `cycles` number of times.
//
// Example:
//
//     result := Repeat("a", 5)
//     fmt.Println(result) // Output: "aaaaa"
//
// If `cycles` is 0, it will return an empty string.
func Repeat(character string, cycles int) string {
	var repeated string

	for i := 0; i < cycles; i++ {
		repeated += character
	}
	return repeated
}
