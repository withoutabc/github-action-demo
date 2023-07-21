package main

import "fmt"

func main() {
	saying := Cat()
	fmt.Println(saying)
}

func smallCat() {

}

func Cat() string {
	return "Miao~~~~~~~"
}
