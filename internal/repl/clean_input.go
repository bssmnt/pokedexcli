package repl

import ("strings")

func CleanInput(text string) []string {
	cleanText := strings.TrimSpace(text)
	lowerCase := strings.ToLower(cleanText)
	splitText := strings.Fields(lowerCase)
	return splitText
}