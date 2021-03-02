package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Mister Pressly"
	text := fmt.Sprintf(`
    <!DOCTYPEhtml>
    <html lang="en">
    <head>
    <meta charset="UTF-8">
    <title>Hello There!</title>
    </head>
    <body>
    <h1>` + name + `</h1>
    </body>
    </html>`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(text))
}
