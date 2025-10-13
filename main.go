package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Measurement struct {
	Name 		string
	TempMin float64
	TempMax float64
	Sum 		float64
	Count 	int64
}

func main() {
	measurements, err := os.Open("measurement.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan(){
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		temp := rawData[semicolon+1:]
		fmt.Println(location, temp)
		return //testando a primeira saída
	}
}