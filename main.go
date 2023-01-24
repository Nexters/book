/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/nexters/book/cmd"
	"github.com/nexters/book/docs"
)

var Port string

func main() {
	docs.SwaggerInfo.Title = "Book API 문서"
	docs.SwaggerInfo.Description = "독서기록 작성 서비스 API 문서"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	cmd.Execute()
}
