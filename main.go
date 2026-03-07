package main

import "fmt"

func main() {
	shortGolang := "Watch the Go crash course"
	fullGolang := "Watch Nana's full course"
	reward := "Bananabread"

	var taskItems = []string{shortGolang, fullGolang, reward}
	fmt.Println(taskItems)
}
