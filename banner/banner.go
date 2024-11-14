package custom

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func InitBanner() {
	var banner string
	colors := []string{
		"\033[31m",       // Red
		"\033[38;5;214m", // Orange
		"\033[33m",       // Yellow
		"\033[32m",       // Green
		"\033[34m",       // Blue
		"\033[38;5;57m",  // Indigo
		"\033[35m",       // Violet
	}
	data, err := ioutil.ReadFile("./banner/ascii-art.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	for i, line := range strings.Split(string(data), ",") {
		color := colors[i%len(colors)]
		banner += fmt.Sprintf("%s%s\033[0m", color, line)
	}

	fmt.Println(banner)
}
