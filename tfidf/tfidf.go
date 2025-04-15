package tfidf

import (
	"math"
	"sort"

	"github.com/phishing-email-classifier/preprocess"
)

type TFIDF struct {
	Vocabulary  map[string]int // Word to index
	IDF         []float64      // Inverse document frequency
	MaxFeatures int            // Limit vocabulary size
	TotalDocs   int
	DocFreq     map[string]int // Document frequency per term
}

func NewTFIDF(maxFeatures int) *TFIDF {
	return &TFIDF{
		Vocabulary:  make(map[string]int),
		MaxFeatures: maxFeatures,
		DocFreq:     make(map[string]int),
	}
}

func (t *TFIDF) Fit(emails []string) {
	t.TotalDocs = len(emails)

	// Build document frequency
	for _, email := range emails {
		tokens := preprocess.Preprocess(email)
		unique := make(map[string]bool)
		for _, token := range tokens {
			if !unique[token] {
				t.DocFreq[token]++
				unique[token] = true
			}
		}
	}

	// Select top MaxFeatures terms by document frequency
	type termFreq struct {
		term string
		freq int
	}
	var terms []termFreq
	for term, freq := range t.DocFreq {
		terms = append(terms, termFreq{term, freq})
	}
	sort.Slice(terms, func(i, j int) bool {
		return terms[i].freq > terms[j].freq
	})

	// Build vocabulary
	numFeatures := t.MaxFeatures
	if len(terms) < numFeatures {
		numFeatures = len(terms)
	}
	for i := 0; i < numFeatures; i++ {
		t.Vocabulary[terms[i].term] = i
	}

	// Compute IDF
	t.IDF = make([]float64, numFeatures)
	for term, idx := range t.Vocabulary {
		t.IDF[idx] = math.Log(float64(t.TotalDocs) / float64(t.DocFreq[term]+1))
	}
}

func (t *TFIDF) Transform(email string) []float64 {
	tokens := preprocess.Preprocess(email)
	termFreq := make(map[int]int)

	// Compute term frequency
	for _, token := range tokens {
		if idx, exists := t.Vocabulary[token]; exists {
			termFreq[idx]++
		}
	}

	// Compute TF-IDF
	features := make([]float64, len(t.Vocabulary))
	totalWords := len(tokens)
	if totalWords == 0 {
		return features
	}

	for idx, freq := range termFreq {
		tf := float64(freq) / float64(totalWords)
		features[idx] = tf * t.IDF[idx]
	}

	return features
}

func (t *TFIDF) FitTransform(emails []string) [][]float64 {
	t.Fit(emails)
	features := make([][]float64, len(emails))
	for i, email := range emails {
		features[i] = t.Transform(email)
	}
	return features
}
