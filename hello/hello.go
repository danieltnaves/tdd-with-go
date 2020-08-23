package main

import "fmt"

const hello = "Hello, "
const helloPtBr = "Olá, "
const helloDe = "Hallo, "
const de = "de"
const ptBr = "pt-BR"

func Hello(name, language string) string {
	helloByLanguage := getHelloByLanguage(language)
	if name == "" {
		return helloByLanguage + "World"
	} else {
		return helloByLanguage + name
	}
}

func getHelloByLanguage(language string) (prefix string) {
	switch language {
	case ptBr:
		prefix = helloPtBr
	case de:
		prefix = helloDe
	default:
		prefix = hello
	}
	return
}

func main() {
	fmt.Println(Hello("Hello", ""))
}
