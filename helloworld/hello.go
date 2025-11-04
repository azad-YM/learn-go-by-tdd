package helloworld

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrifix = "Hello, "
	spanishHelloPrifix = "Hola, "
	frechHelloPrifix   = "Salut, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrifix
	case french:
		prefix = frechHelloPrifix
	default:
		prefix = englishHelloPrifix
	}

	return
}
