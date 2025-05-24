package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func replaceMathSymbolWithLanguageOperator(expression string) string {
	expression = strings.Replace(expression, " ", "", -1)
	expression = strings.Replace(expression, "x", "*", -1)
	expression = strings.Replace(expression, "÷", "/", -1)
	expression = strings.Replace(expression, "mod", "%", -1)
	return expression
}

func evaluatePolarity(expression string) string {
	runes := []rune(expression)
	for i := 0; i < len(runes); i++ {
		if i+1 >= len(runes) {
			continue
		}

		current := runes[i]
		next := runes[i+1]

		// Check for consecutive operators
		if _isPolarityOperator(current) && _isPolarityOperator(next) {
			var replacement rune

			// Determine replacement based on operator combination
			if current == next {
				replacement = '+'
			} else {
				replacement = '-'
			}

			// Rebuild expression with replacement and recurse
			newExpr := string(runes[:i]) + string(replacement) + string(runes[i+2:])
			return evaluatePolarity(newExpr)
		}
	}
	return expression
}

func _isPolarityOperator(character rune) bool {
	return character == '+' || character == '-'
}

func generateExpressionList(expression string) []string {
	expression = replaceMathSymbolWithLanguageOperator(expression)

	var newList []string
	temp := ""
	for _, char := range expression {
		switch char {
		case '%', '*', '/', '+', '-', '!', '^', '(', ')', '√':
			newList = append(newList, temp, string(char))
			temp = ""
		default:
			temp += string(char)
		}
	}
	newList = append(newList, temp)

	var newList2 []string
	for _, s := range newList {
		if s != "" {
			newList2 = append(newList2, s)
		}
	}
	newList2 = _join1stAnd2ndElementIf1stElementIsPolarAnd2ndIsNotAnOperator(newList2)
	return newList2
}

func _join1stAnd2ndElementIf1stElementIsPolarAnd2ndIsNotAnOperator(newList []string) []string {
	if len(newList) >= 2 {
		firstOperators := map[string]bool{"+": true, "-": true}
		forbiddenOperators := map[string]bool{
			"%": true, "*": true, "/": true, "+": true, "-": true,
			"!": true, "^": true, "(": true, ")": true, "√": true,
		}

		if firstOperators[newList[0]] && !forbiddenOperators[newList[1]] {
			newList[1] = newList[0] + newList[1]
			return newList[1:]
		}
	}
	return newList
}

func __resolvedExpression(expressionList []string, index int, operation string) []string {
	var newElement float64

	switch operation {
	case "!":
		return _resolveFactorial(expressionList, index)
	case "√":
		return _resolveSquareRoot(expressionList, index)
	case ")":
		return _resolveParenthesis(expressionList, index)
	}

	a, _ := strconv.ParseFloat(expressionList[index-1], 64)
	b, _ := strconv.ParseFloat(expressionList[index+1], 64)

	switch operation {
	case "*":
		newElement = a * b
	case "/":
		newElement = a / b
	case "%":
		newElement = math.Mod(a, b)
	case "+":
		newElement = a + b
	case "-":
		newElement = a - b
	case "^":
		newElement = math.Pow(a, b)
	}

	newElementStr := fmt.Sprintf("%g", newElement)
	if !strings.Contains(newElementStr, ".") {
		newElementStr += ".0" // Ensure at least one decimal digit
	}
	newList := append([]string{}, expressionList[:index-1]...)
	newList = append(newList, newElementStr)
	newList = append(newList, expressionList[index+2:]...)
	return newList
}

func _resolveFactorial(expressionList []string, index int) []string {
	a, _ := strconv.Atoi(expressionList[index-1])
	var newElement = float64(_calculateFactorials(a))
	newElementStr := fmt.Sprintf("%g", newElement)
	if !strings.Contains(newElementStr, ".") {
		newElementStr += ".0" // Ensure at least one decimal digit
	}
	newList := append([]string{}, expressionList[:index-1]...)
	newList = append(newList, newElementStr)
	newList = append(newList, expressionList[index+1:]...)
	return newList
}

func _calculateFactorials(a int) int {
	if a == 0 {
		return 1
	}
	return a * _calculateFactorials(a-1)
}

func _resolveSquareRoot(expressionList []string, index int) []string {
	b, _ := strconv.ParseFloat(expressionList[index+1], 64)
	newElement := math.Sqrt(b)
	newElementStr := fmt.Sprintf("%g", newElement)
	if !strings.Contains(newElementStr, ".") {
		newElementStr += ".0" // Ensure at least one decimal digit
	}
	newList := append([]string{}, expressionList[:index]...)
	newList = append(newList, newElementStr)
	newList = append(newList, expressionList[index+2:]...)
	return newList
}

func _resolveParenthesis(expressionList []string, closeIdx int) []string {
	openIdx := _getIndexOfLastOpeningParenthesisBeforeIndex(expressionList, closeIdx)
	var newExpression strings.Builder
	for _, elem := range expressionList[openIdx+1 : closeIdx] {
		newExpression.WriteString(elem)
	}

	var result float64 = calculate(newExpression.String())
	//add zeros

	newElementStr := fmt.Sprintf("%g", result)
	if !strings.Contains(newElementStr, ".") {
		newElementStr += ".0" // Ensure at least one decimal digit
	}

	return append(
		append(expressionList[:openIdx],
			newElementStr),
		expressionList[closeIdx+1:]...,
	)

}

func _getIndexOfLastOpeningParenthesisBeforeIndex(expressionList []string, closeParenthesisIndex int) int {
	for i := closeParenthesisIndex - 1; i >= 0; i-- {
		if expressionList[i] == "(" {
			return i
		}
	}
	return -1
}

func _evaluateOperation(expressionList []string, operator string) []string {
	for index, element := range expressionList {
		if element == operator && len(expressionList) > 1 {
			expressionList = __resolvedExpression(expressionList, index, operator)
			return _evaluateOperation(expressionList, operator)
		}
	}
	return expressionList
}

func evaluateMultiplications(expressionList []string) []string {
	return _evaluateOperation(expressionList, "*")
}

func evaluateDivisions(expressionList []string) []string {
	return _evaluateOperation(expressionList, "/")
}

func evaluateRemainders(expressionList []string) []string {
	return _evaluateOperation(expressionList, "%")
}

func evaluateAdditions(expressionList []string) []string {
	return _evaluateOperation(expressionList, "+")
}

func evaluateSubstractions(expressionList []string) []string {
	return _evaluateOperation(expressionList, "-")
}

func evaluateExponents(expressionList []string) []string {
	return _evaluateOperation(expressionList, "^")
}

func evaluateFactorials(expressionList []string) []string {
	return _evaluateOperation(expressionList, "!")
}

func evaluateSquareRoots(expressionList []string) []string {
	return _evaluateOperation(expressionList, "√")
}

func evaluateParenthesis(expressionList []string) []string {
	return _evaluateOperation(expressionList, ")")
}

func calculate(expression string) float64 {
	var expressionList []string
	expression = replaceMathSymbolWithLanguageOperator(expression)
	expression = evaluatePolarity(expression)
	expressionList = generateExpressionList(expression)
	expressionList = evaluateParenthesis(expressionList)
	expressionList = evaluateFactorials(expressionList)
	expressionList = evaluateSquareRoots(expressionList)
	expressionList = evaluateExponents(expressionList)
	expressionList = evaluateMultiplications(expressionList)
	expressionList = evaluateDivisions(expressionList)
	expressionList = evaluateRemainders(expressionList)
	expressionList = evaluateAdditions(expressionList)
	expressionList = evaluateSubstractions(expressionList)
	result, _ := strconv.ParseFloat(expressionList[0], 64)
	return result
}
