package api

import (
	"io"
	"net/http"
	"time"

	"github.com/alpgozbasi/textsentimentor/internal/sentiment"
	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	doneChan := make(chan bool)

	go func() {
		// simulate small work
		time.Sleep(100 * time.Millisecond)
		doneChan <- true
	}()

	<-doneChan

	c.String(http.StatusOK, "ok!")
}

func AnalyzeHandler(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return
	}
	defer c.Request.Body.Close()

	text := string(bodyBytes)

	resultChan := make(chan string)

	go func(txt string, ch chan<- string) {
		if loadErr := sentiment.LoadModel(); loadErr != nil {
			ch <- "error loading model"
			return
		}

		// analyze text
		res, analyzeErr := sentiment.AnalyzeSentiment(txt)
		if analyzeErr != nil {
			ch <- "analysis error"
			return
		}

		ch <- "sentiment result: " + res
	}(text, resultChan)

	result := <-resultChan
	c.String(http.StatusOK, result)
}
