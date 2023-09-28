package main

import "fmt"

func main() {
	input := []float64{1, 2, 3, 2.5}
	weights1 := []float64{0.2, 0.8, -0.5, 1}
	weights2 := []float64{0.5, -0.91, 0.26, -0.5}
	weights3 := []float64{-0.26, -0.27, 0.17, 0.87}

	bias1 := 2.0
	bias2 := 3.0
	bias3 := 0.5

	var result []float64
	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Weights1: %v\n", weights1)
	fmt.Printf("Weights2: %v\n", weights2)
	fmt.Printf("Weights3: %v\n", weights3)
	fmt.Printf("Bias1: %f\n", bias1)
	fmt.Printf("Bias2: %f\n", bias2)
	fmt.Printf("Bias3: %f\n", bias3)

	result = append(result, bias1)
	for i := range input {
		result[0] += input[i] * weights1[i]
	}

	result = append(result, bias2)
	for i := range input {
		result[1] += input[i] * weights2[i]
	}

	result = append(result, bias3)
	for i := range input {
		result[2] += input[i] * weights3[i]
	}

	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Ok!\n")
}
