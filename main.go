package main

import (
	"fmt"

	"github.com/hugermuger/gator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Test: %v", config)
	return
}
