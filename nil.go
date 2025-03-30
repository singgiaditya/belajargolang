package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}

	return map[string]string{
		"name": name,
	}
}

func main() {
	var data = newMap("singgi")
	fmt.Println(data["name"])
}
