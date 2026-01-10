package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/learn-go/note"
)

func main() {

	userNote, err := note.New(getNoteData())

	if err != nil {
		fmt.Print(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return
	}

	fmt.Println("file saved successfully")

}

func getNoteData() (string, string) {
	title := getUserInput("Note tile :")
	content := getUserInput("Note Content :")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
