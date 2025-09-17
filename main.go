package main

import (
	"fmt"
	"log"
	"github.com/smwalke83/gator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", c)
	err = c.SetUser("scott")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}
	c, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}
	fmt.Printf("Read config again: %+v\n", c)
}