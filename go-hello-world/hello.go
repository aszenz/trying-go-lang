package main

import "fmt"

const prefix = "Hello, "

func getGreetingPrefix(lang string) string {
	englishHello := "Hello, "

	dutchLangCode := "nl_NL"
	dutchHello := "Hallo, "

	frenchLangCode := "fr_FR"
	frenchHello := "Bonjour, "

	prefix := ""
	switch lang {
	case dutchLangCode:
		prefix = dutchHello
	case frenchLangCode:
		prefix = frenchHello
	default:
		prefix = englishHello
	}
	return prefix
}

// Hello returns Hello, <name>
func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return getGreetingPrefix(lang) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
