package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tomasji/test/internal/colors"
)

type args struct {
	color   string
	dataDir string
}

func parseArgs() args {
	var cfg args

	flag.StringVar(&cfg.color, "color", "blue", "color to count")
	flag.StringVar(&cfg.dataDir, "dir", "./data", "data directory")
	flag.Parse()
	return cfg
}

func main() {
	cfg := parseArgs()

	colorCount, err := colors.Count(cfg.dataDir, cfg.color)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of occurrences of '%s': %d\n", cfg.color, colorCount)
}
