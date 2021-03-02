package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}

		go handle(conn)

	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//read the request with header
	request(conn)

	//write response
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// Position 0 is always request type and URL
			method := strings.Fields(ln)[0]
      url := strings.Fields(ln)[1]
			fmt.Println("### METHOD: ", method)
      fmt.Println("### URI: ", url)
		}
		if ln == "" {
			break
		}

		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPEhtml>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>BAR</title>
  </head>
  <body>
  <h1>BAR</h1>
  </body>
  </html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
