package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
	<html>
		<head>
			<title>Hello Bro</title>
		</head>
		<body>
			Hello World!, Today is a beautiful day...
		</body>
	</html>`,
	)
}
func main() {
	http.HandleFunc("/", hello)
	// http.Handle("/assets/",
	// 	http.StripPrefix(
	// 		"/assets/",
	// 		http.FileServer(http.Dir("assets")),
	// 	))
	http.ListenAndServe(":9000", nil)
}
