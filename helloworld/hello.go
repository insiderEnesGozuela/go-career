package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Halo, "
const Spanish = "Spanish"
const frenchHelloPrefix = "Bonjour, "
const French = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == Spanish {
		return spanishHelloPrefix + name
	}

	if language == French {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func greetingPrefix(prefix string, name string) string {
	if prefix == Spanish {
		return spanishHelloPrefix + name
	} else if prefix == French {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
}
