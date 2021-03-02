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

			//write response
			responder(conn, method, url)
		}
		if ln == "" {
			break
		}

		i++
	}
}

func responder(conn net.Conn, method string, url string) {
	if method == "GET" {
		switch url {
		case "/":
			index(conn)
		case "/profile":
			profile(conn)
		case "/about":
			about(conn)
		default:
			not_found(conn)
		}
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPEhtml>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>HOME</title>
  </head>
  <body>
  <h1>Home</h1>
	<strong>Everywhere is good but home is the best!</strong>
	<ul>
	<li><a href="/">go home<a></li>
	<li><a href="/profile">go profile<a></li>
	<li><a href="/about">go about<a></li>
	</ul>
  </body>
  </html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func profile(conn net.Conn) {
	body := `<!DOCTYPEhtml>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>PROFILE</title>
  </head>
  <body>
  <h1>PROFILE</h1>
	<p> This guy and his story .... </p>
	<a href="/">go home<a>
	<a href="/profile">go profile<a>
	<a href="/about">go about<a>
  </body>
  </html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPEhtml>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>About</title>
  </head>
  <body>
  <h1>About</h1>
	<p> Of course this is page  </p>
	<a href="/">go home<a>
	<a href="/profile">go profile<a>
	<a href="/about">go about<a>
  </body>
  </html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func not_found(conn net.Conn) {
	body := `<!DOCTYPEhtml>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>PROFILE</title>
  </head>
  <body>
  <h1>404 Page Not Found...</h1>
  </body>
  </html>`

	fmt.Fprint(conn, "HTTP/1.1 404 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
