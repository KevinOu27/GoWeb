package main

import (
	"encoding/json"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

func main() {
	urls := []string{
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

	// Initialize the collector
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// Apply Random User-Agent
	extensions.RandomUserAgent(c)

	// Create the output file
	file, err := os.Create("results.jl")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Handle the scraping process
	c.OnHTML("div.mw-parser-output", func(e *colly.HTMLElement) {
		// Extract text
		text := e.Text
		// Create a map to encode to JSON
		data := map[string]string{
			"url":  e.Request.URL.String(),
			"text": text,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			return
		}
		// Write JSON data to the file with newline
		file.Write(jsonData)
		file.WriteString("\n")
	})

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	// Iterate over all URLs
	for _, url := range urls {
		c.Visit(url)
	}
}
