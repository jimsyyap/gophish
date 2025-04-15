package preprocess

import (
	"github.com/kljensen/snowball"
	"regexp"
	"strings"
)

// Stopwords (simplified list; expand as needed)
var stopwords = map[string]bool{
	"a": true, "an": true, "and": true, "are": true, "as": true, "at": true,
	"be": true, "by": true, "for": true, "from": true, "has": true, "he": true,
	"in": true, "is": true, "it": true, "its": true, "of": true, "on": true,
	"that": true, "the": true, "to": true, "was": true, "were": true, "with": true,
}

var (
	urlRegex    = regexp.MustCompile(`http[s]?://\S+`)
	emailRegex  = regexp.MustCompile(`\S+@\S+`)
	nonAlphaNum = regexp.MustCompile(`[^a-zA-Z\s]`)
	headerRegex = regexp.MustCompile(`(?s)^.*?\n\n`)
)

func Preprocess(email string) []string {
	// remove headers (everything before first double newline)
	body := headerRegex.ReplaceAllString(email, "")

	//remove urls, email addresses, and non-alphanumeric characters
	body = urlRegex.ReplaceAllString(body, "")
	body = emailRegex.ReplaceAllString(body, "")
	body = nonAlphaNum.ReplaceAllString(body, "")

	// convert to lowercase and tokensize
	words := strings.Fields(strings.ToLower(body))

	// stem and remove stopwords
	var tokens []string
	for _, word := range words {
		if !stopwords[word] {
			stemmed, err := snowball.Stem(word, "english", true)
			if err == nil && stemmed != "" {
				tokens = append(tokens, stemmed)
			}
		}
	}
	return tokens
}
