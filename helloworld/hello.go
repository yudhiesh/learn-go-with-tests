package 

import "fmt"

const (
	englishHelloWorldPrefix = "Hello, "
	spanishHelloWorldPrefix = "Hola, "
	frenchHelloWorldPrefix  = "Bonjour, "
	spanish                 = "Spanish"
	french                  = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	prefix := englishHelloWorldPrefix
	switch language {
	case spanish:
		prefix = spanishHelloWorldPrefix
	case french:
		prefix = frenchHelloWorldPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
