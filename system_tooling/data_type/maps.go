package main

import "fmt"

func main() {
	birthdays := map[string]string{
		"keith": "02/11/2016",
		"abc":   "02/12/2017",
		"qwe":   "02/1/2018",
	}
	fmt.Println(birthdays)

	ages := map[string]int{}
	ages["Keith"] = 28
	ages["Kevin"] = 12
	ages["Kayla"] = 25
	fmt.Println(ages)

}
