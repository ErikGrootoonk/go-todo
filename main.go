package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("### Welcome to our to our Todo list App ###")
	// var shortGolang = "Watch the Go crash course"
	// var fullGolang = "Watch Nana's full course"
	// var reward = "Bananabread"

	// var taskItems = []string{shortGolang, fullGolang, reward}
	// printTasks(taskItems)
	// taskItems = addTask(taskItems, "Go to the gym")
	// printTasks(taskItems)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, welcome to our Todo list App!")
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		taskItems := []string{"Watch the Go crash course", "Watch Nana's full course", "Bananabread"}
		for i, task := range taskItems {
			fmt.Fprintf(w, "%d. %s\n", i+1, task)
		}
	})

	http.ListenAndServe(":8080", nil)

}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")

	for i, task := range taskItems {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func addTask(taskItems []string, task string) []string {
	return append(taskItems, task)
}
