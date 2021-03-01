package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type result struct {
	Score   int    `json:"score"`
	Message string `json:"message"`
}

// PrintResult prints score to stdout
func PrintResult() {
	b, _ := json.Marshal(result{Score: totalScore})
	fmt.Println(string(b))
}

// PrintFatal は JSON を出しつつエラーログを出してかつ終了もします
func PrintFatal(message string) {
	b, _ := json.Marshal(result{Score: 0, Message: message})
	log.Println(message)
	fmt.Println(string(b))
	os.Exit(1)
}
