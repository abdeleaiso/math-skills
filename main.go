package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run program.go data.txt")
		return
	}

	// read data file
	filename := os.Args[1]
	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// fields automatically handles spaces, tabs and newlines.
	// split will fail if there is empty element in the end.
	fileStr := strings.Fields(string(fileData))

	if len(fileStr) == 0 {
		fmt.Println("Error: empty file")
		return
	}

	// transform to numbers.
	Data := make([]float64, 0)
	for _, val := range fileStr {
		val, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		valf := float64(val)
		Data = append(Data, valf)
	}

	// calculate average
	var sum float64
	for _, v := range Data {
		sum += v
	}
	l := float64(len(Data))
	average := sum / l
	fmt.Printf("Average: %.0f\n", average)

	// calcul Median
	var median float64
	slices.Sort(Data)
	n := len(Data)
	if n%2 == 0 {
		median = (Data[n/2-1] + Data[(n-1)/2]) / 2
	} else {
		median = Data[n/2]
	}
	fmt.Printf("Median: %.0f\n", median)

	// calcul Variance
	// var = 1/sum * Loop(xi - E(x))^2
	// Var := E(X^2) - [E(X)]^2  -> E(X^2) = 1/len(Data)*(Data[i]^2)
	mean := math.Pow(average, 2)
	var seg float64
	for i := 0; i < len(Data); i++ {
		seg += math.Pow(Data[i], 2)
	}
	SecondRawMoment := seg / float64(len(Data))
	variance := SecondRawMoment - mean

	fmt.Printf("variance: %.0f\n", variance)

	// calculate the standard Deviation
	StandardDeviation := math.Sqrt(variance)
	fmt.Printf("Standard Deviation: %.0f\n", StandardDeviation)
}
