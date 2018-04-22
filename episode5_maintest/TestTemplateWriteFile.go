package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	log.Println("compose letter..")

	// Define a template.
	const letter = `
	Dear {{.Name}},
	{{if .Attended}}
	It was a pleasure to see you at the wedding.
	{{- else}}
	It is a shame you couldn't make it to the wedding.
	{{- end}}
	{{with .Gift -}}
	Thank you for the lovely {{.}}.
	{{end}}
	Best wishes,
	Josie
	`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

	log.Println("start writing to file .\\letter.txt")

	content := []byte("hello world!")
	err := ioutil.WriteFile(".\\test_0000.txt", content, 0000)
	//err2 := ioutil.WriteFile("E:\\test_0664.txt", content, 0000)

	if err != nil {
		fmt.Println(err)
	}
	// if err2 != nil {
	// 	fmt.Println(err)
	// }

	log.Println("complete writing file")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
