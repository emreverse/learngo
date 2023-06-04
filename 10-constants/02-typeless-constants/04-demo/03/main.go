package main

import (
	"fmt"
)

type mathExpressionType = string

const (
	addExpression   = mathExpressionType("add")
	subExpression   = mathExpressionType("subtract")
	multiExpression = mathExpressionType("multiply")
)

func main() {
	value := mathExpression(subExpression)
	fmt.Println(value(2, 3))
	//fmt.Println(mathExpression(subExpression)(2, 3))
}

func mathExpression(expression mathExpressionType) func(float64, float64) float64 {
	switch expression {
	case addExpression:
		return func(x, y float64) float64 {
			return x + y
		}
	case subExpression:
		return func(a, b float64) float64 {
			return a - b
		}
	case multiExpression:
		return func(a, b float64) float64 {
			return a * b
		}
	default:
		fmt.Println("invalid math expression")
		return nil
	}
}
