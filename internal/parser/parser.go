package parser

type Parser interface {
	// Parse always accepts a string to parse using the rules
	// provided and returns the parsed string.
	Parse(string) error
}