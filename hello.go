package main

// func main() {
// 	fmt.Println(Hello("", english))
// }

type language int32

const (
	english language = 1
	spanish language = 2
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, idiom language) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(idiom) + name
}

func greetingPrefix(idiom language) (prefix string) {
	switch idiom {
	case spanish:
		prefix = spanishHelloPrefix
	case english:
		prefix = englishHelloPrefix
	}
	return
}
