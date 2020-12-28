package main

import (
	"fmt"
	"strconv"

	"github.com/webview/webview"
)

func main() {
	go Quark_server()

	app := webview.New(true)
	defer app.Destroy()

	app.SetSize(300, 365, webview.HintNone)
	app.Navigate("http://localhost:8080/index.html")
	app.Bind("goEquals", func(expr string) {
		app.Eval(fmt.Sprintf(`dSet("%f")`, parse_add(expr)))
	})
	app.Bind("goLog", fmt.Println)
	app.Run()
}

func parse_add(expr string) float64 {
	for i, c := range expr {
		if c == '+' {
			return parse_mul(expr[0:i]) + parse_add(expr[i+1:len(expr)])
		}
		if c == '-' {
			return parse_mul(expr[0:i]) - parse_add(expr[i+1:len(expr)])
		}
	}

	return parse_mul(expr)
}

func parse_mul(expr string) float64 {
	for i, c := range expr {
		if c == '*' {
			return parse_atom(expr[0:i]) * parse_mul(expr[i+1:len(expr)])
		}
		if c == '/' {
			return parse_atom(expr[0:i]) / parse_mul(expr[i+1:len(expr)])
		}
	}

	return parse_atom(expr)
}

func parse_atom(expr string) float64 {
	ret, _ := strconv.ParseFloat(expr, 64)
	return ret
}
