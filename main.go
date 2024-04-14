package main

import (
	"gen-stage/psc"
	"strings"
)

func main() {
	// Create the channels
	data := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry", "Fig", "Grape", "Honeydew", "Icaco", "Jackfruit", "Kiwi", "Lemon", "Mango", "Nectarine", "Orange", "Papaya", "Quince", "Raspberry", "Strawberry", "Tangerine", "Ugli fruit", "Vanilla bean", "Watermelon", "Xigua", "Yuzu", "Zucchini", "Apricot", "Blackberry", "Cantaloupe", "Dragon fruit"}
	out := make(chan string)
	processed := make(chan string)
	ready := make(chan struct{}, 1) // Buffered channel
	batchSize := 2                  // Number of items Consumer can process at a time

	// Simple example processing functions
	ate := func(item string) string {
		return "ate: " + item
	}

	title := func(item string) string {
		return strings.Title(item)
	}

	// Start the producer and consumer
	go psc.Producer[string](data, out, ready, batchSize)
	firstOut := psc.ProducerConsumer(out, title, ready, batchSize)
	go psc.Consumer(firstOut, processed, ate, ready, batchSize)

	// Initially signal the producer to start sending items
	ready <- struct{}{}

	// Print the processed items
	for item := range processed {
		println("Processed item:", item)
	}
}
