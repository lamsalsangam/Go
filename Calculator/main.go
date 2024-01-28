package main

import "fmt"

type Calculator interface {
	Add() float64
	Subtract() float64
	Multiply() float64
	Divide() (float64, error)
}

type CalculatorImp1 struct {
	Operand1 float64
	Operand2 float64
}

func (c CalculatorImp1) Add() float64 {
	return c.Operand1 + c.Operand2
}

func (c CalculatorImp1) Subtract() float64 {
	return c.Operand1 - c.Operand2
}

func (c CalculatorImp1) Multiply() float64 {
	return c.Operand1 * c.Operand2
}

func (c CalculatorImp1) Divide() (float64, error) {
	if c.Operand2 == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return c.Operand1 / c.Operand2, nil
}

func main() {
	calculator := CalculatorImp1{Operand1: 10, Operand2: 2}

	// Perform calculations
	additionResult := calculator.Add()
	fmt.Printf("Addition Result: %.2f\n", additionResult)

	subtractionResult := calculator.Subtract()
	fmt.Printf("Subtraction Result: %.2f\n", subtractionResult)

	multiplicationResult := calculator.Multiply()
	fmt.Printf("Multiplication Result: %.2f\n", multiplicationResult)

	divisionResult, err := calculator.Divide()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Division Result: %.2f\n", divisionResult)
	}

}
