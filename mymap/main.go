package main

import "fmt"

func main() {

	fmt.Println("Welcome to the maps class")

	LanguageList := make(map[string]string)

	LanguageList["Go"] = "Golang"
	LanguageList["JS"] = "Javascript"
	LanguageList["C#"] = "C Sharp"

	fmt.Println("Language list: ", LanguageList)
	fmt.Println("Language list length: ", len(LanguageList))
	fmt.Println("Language list of C#: ", LanguageList["C#"])

	delete(LanguageList, "C#")
	fmt.Println("Language list: ", LanguageList)

	for key, value := range LanguageList {
		fmt.Println("Key:", key, "Value:", value)
	}

}
