package hello_world

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func Bye(name string) string {
	return "Bye, have a great day!"
}

// func main() {
// 	fmt.Println(Hello("Chris"))
// }
