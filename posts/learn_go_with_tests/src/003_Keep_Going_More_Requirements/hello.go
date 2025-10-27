package main

import "fmt"

var translations = map[string]string{
	"en": "Hello",
	"es": "Hola",
	"fr": "Bonjour",
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix, ok := translations[language]
	if !ok {
		prefix = translations["en"]
	}

	return fmt.Sprintf("%s, %s!", prefix, name)
}
