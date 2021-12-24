package hello

import "fmt"

func main() {
	fmt.Println(Hello("", ""))
}

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const french = "French"
const russianHelloPrefix = "Привет, "
const russian = "Russian"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

//lowercase - private
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
