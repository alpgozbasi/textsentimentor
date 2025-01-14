package sentiment

import (
	"sync"

	"github.com/cdipaolo/sentiment"
)

var (
	model sentiment.Models
	once  sync.Once
)

func LoadModel() error {
	var err error
	once.Do(func() {
		model, err = sentiment.Restore()
	})
	return err
}

func AnalyzeSentiment(text string) (string, error) {
	analysis := model.SentimentAnalysis(text, sentiment.English)
	switch analysis.Score {
	case 0:
		return "negative", nil
	case 1:
		return "positive", nil
	default:
		return "neutral", nil
	}

}
