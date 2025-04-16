package main

import (
	"fmt"
	"log"
	"net/http"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="ru">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Demo App</title>
	<style>
		body {
			background-color: #D8BFD8;
			display: flex;
			justify-content: center;
			align-items: center;
			height: 100vh;
			margin: 0;
			color: white;
			text-align: center;
		}
		.button {
			border: 2px solid purple;
			padding: 10px 20px;
			background-color: transparent;
			color: white;
			cursor: pointer;
			font-size: 16px;
		}
	</style>
</head>
<body>
	<div>
		<h1>Hello, 世界</h1>
		<button class="button">Нажми на меня</button>
	</div>
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, htmlTemplate)
}


func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
func newFeature() {
    fmt.Println("Это новая функциональность!")
}