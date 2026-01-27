package main

import (
	"fmt"

	"github.com/hugermuger/gator/internal/config"
)

func main() {
	jsonconfig, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = jsonconfig.SetUser("hugermuger")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	jsonconfig, err = config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Test: %v", jsonconfig)
}
