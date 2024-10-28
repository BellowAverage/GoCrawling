package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

var urls = []string{
	"https://en.wikipedia.org/wiki/Robotics",
	"https://en.wikipedia.org/wiki/Robot",
	"https://en.wikipedia.org/wiki/Reinforcement_learning",
	"https://en.wikipedia.org/wiki/Robot_Operating_System",
	"https://en.wikipedia.org/wiki/Intelligent_agent",
	"https://en.wikipedia.org/wiki/Software_agent",
	"https://en.wikipedia.org/wiki/Robotic_process_automation",
	"https://en.wikipedia.org/wiki/Chatbot",
	"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
	"https://en.wikipedia.org/wiki/Android_(robot)",
}

type Article struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

func main() {

	start := time.Now()

	var wg sync.WaitGroup
	articles := make(chan Article, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			scrape(u, articles)
		}(url)
	}

	go func() {
		wg.Wait()
		close(articles)
	}()

	writeToFile(articles)

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func scrape(url string, articles chan<- Article) {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnHTML("div#mw-content-text", func(e *colly.HTMLElement) {
		// Extract text and clean it
		text := e.DOM.Find("p, li").Text()
		text = strings.TrimSpace(text)
		articles <- Article{URL: url, Text: text}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.Visit(url)
}

func writeToFile(articles <-chan Article) {
	file, err := os.Create("articles_go.jsonl")
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}
	defer file.Close()

	for article := range articles {
		jsonData, err := json.Marshal(article)
		if err != nil {
			log.Printf("Could not marshal data: %v", err)
			continue
		}
		file.WriteString(string(jsonData) + "\n")
	}
}
