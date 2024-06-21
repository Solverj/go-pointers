package main

import "fmt"

type YoutubePerson struct {
	name      string
	telephone string
}

func main() {
	youtubePerson := YoutubePerson{
		name:      "",
		telephone: "",
	}
	fmt.Printf("initial youtubePerson: %v\n ", youtubePerson)

	modifyYoutubePerson(youtubePerson)
	fmt.Printf("after modifyYoutubePerson: %v\n ", youtubePerson)

	actuallyModifyYoutubePerson(&youtubePerson)
	fmt.Printf("after actuallyModifyYoutubePerson: %v\n ", youtubePerson)
}

func modifyYoutubePerson(youtubePerson YoutubePerson) {
	youtubePerson.name = "John Doe"
	youtubePerson.telephone = "1234567890"
}

func actuallyModifyYoutubePerson(youtubePerson *YoutubePerson) {
	youtubePerson.name = "John Doe"
	youtubePerson.telephone = "1234567890"
}
