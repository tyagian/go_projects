package main

// maps example, key and values like dictionary {key: value}

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	// output
	// List of all languages: map[JS:JavaScript PY:Python RB:Ruby]
	// if you notice, they are not comma separated

	fmt.Println("JS shorts for:", languages["JS"])

	// delete
	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// loop through maps
	for key, value := range languages {
		fmt.Printf("FOR key %v , value is %v \n", key, value)
	}

	// another way
	// if we dont care what key is, use this approach
	for _, value := range languages {
		fmt.Printf("FOR key v , value is %v \n", value)
	}
}
