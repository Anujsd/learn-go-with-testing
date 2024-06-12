package hello_world

import "fmt"

const englishprefix = "Hello, "
const spanishprefix = "Hola, "
const frenchprefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := getGreetingPrefix(lang)
	return prefix + name
}

func getGreetingPrefix(lang string) (prefix string) {
	switch lang {
	case "Spanish":
		prefix = spanishprefix
	case "French":
		prefix = frenchprefix
	default:
		prefix = englishprefix
	}
	return
}

func main() {
	fmt.Println(Hello("Anuj", "French"))
}
