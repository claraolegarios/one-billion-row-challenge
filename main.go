package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	
	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan(){
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]
		
		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurements, ok := data[location]
		if !ok {
			measurements = Measurement{
				TempMin: temp,
				TempMax: temp,
				Sum: temp,
				Count: 1,
			}
		} else {
			measurements.TempMin = min(measurements.TempMin, temp)
			measurements.TempMax = min(measurements.TempMax, temp)
			measurements.Sum += temp
			measurements.Count++
		}
		data[location] = measurements
	}

		for name, measurements := range data {
			fmt.Printf("%s: %#+v\n", name, measurements)
		}
	}
