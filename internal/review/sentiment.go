package review

import (
	"strings"
)

// Very simple sentiment engine for the assignment.
// In Prod wee should use a transformer model or OpenAI API.
func SentimentScore(text string) float64 {
	text = strings.ToLower(text)

	positive := []string{"good", "great", "amazing", "love", "fantastic", "tasty"}
	negative := []string{"bad", "terrible", "awful", "hate", "disgusting"}

	score := 0.0

	for _, w := range positive {
		if strings.Contains(text, w) {
			score += 1
		}
	}

	for _, w := range negative {
		if strings.Contains(text, w) {
			score -= 1
		}
	}

	// normalize to 0–1
	return (score + 1) / 2
}
