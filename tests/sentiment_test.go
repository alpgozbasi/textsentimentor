package tests

import (
	"testing"

	"github.com/alpgozbasi/textsentimentor/internal/sentiment"
)

func TestAnalyzeSentiment(t *testing.T) {
	err := sentiment.LoadModel()
	if err != nil {
		t.Fatalf("failed to load model: %v", err)
	}

	// positive test
	pos, _ := sentiment.AnalyzeSentiment("i love Ataturk")
	if pos != "positive" {
		t.Errorf("expected positive, got %s", pos)
	}

	// negative test
	neg, _ := sentiment.AnalyzeSentiment("i hate bugs")
	if neg != "negative" {
		t.Errorf("expected negative, got %s", neg)
	}

	// neutral test
	neu, _ := sentiment.AnalyzeSentiment("it is a day")
	if neu != "neutral" {
		t.Errorf("expected neutral, got %s", neu)
	}
}
