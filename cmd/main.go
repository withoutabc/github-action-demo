package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"log"
)

func main() {
	saying := Cat()
	fmt.Println(saying)
	node, err := snowflake.NewNode(0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(node)
}

func Cat() string {
	return "Miao~~~~~~~"
}
