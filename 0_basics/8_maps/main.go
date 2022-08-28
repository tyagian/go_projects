package main

import "fmt"

func main() {

	fmt.Println("Maps tutorial")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["KT"] = "Kotlin"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS short for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("After deleting:", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	/*
		// if we don't want to use key
		for _, value := range languages {
			fmt.Printf("For key v, value is %v\n", value)
		}

	*/
}
