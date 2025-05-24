package main

import (
	"reflect"
	"testing"
)

func Test_replaceMathSymbolWithLanguageOperator(t *testing.T) {
	var actual string
	actual = replaceMathSymbolWithLanguageOperator("x")
	expected := "*"
	if actual != expected {
		t.Errorf("actual = %s, expected = %s", actual, expected)
	}

	actual = replaceMathSymbolWithLanguageOperator("x mod ÷ ^ ")
	expected = "*%/^"
	if actual != expected {
		t.Errorf("actual = %s, expected = %s", actual, expected)
	}

	actual = replaceMathSymbolWithLanguageOperator("x23 mod 3÷ 44^ ")
	expected = "*23%3/44^"
	if actual != expected {
		t.Errorf("actual = %s, expected = %s", actual, expected)
	}

}

func Test_generateExpressList(t *testing.T) {
	actual := generateExpressionList("12")
	expected := []string{"12"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	actual = generateExpressionList("12+7")
	expected = []string{"12", "+", "7"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	actual = generateExpressionList("122 + 34 * 5 / 9 ")
	expected = []string{"122", "+", "34", "*", "5", "/", "9"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	actual = generateExpressionList("122+34!*5^2/9")
	expected = []string{"122", "+", "34", "!", "*", "5", "^", "2", "/", "9"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
	actual = generateExpressionList("!%^*()-+/*√")
	expected = []string{"!", "%", "^", "*", "(", ")", "-", "+", "/", "*", "√"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_evaluateMultiplications(t *testing.T) {
	sample := []string{"12"}
	actual := evaluateMultiplications(sample)
	expected := []string{"12"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"2", "*", "2"}
	actual = evaluateMultiplications(sample)
	expected = []string{"4.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"2", "*", "2", "*", "2"}
	actual = evaluateMultiplications(sample)
	expected = []string{"8.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"2", "*", "2", "*", "2", "+", "2"}
	actual = evaluateMultiplications(sample)
	expected = []string{"8.0", "+", "2"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"2", "+", "2", "*", "2", "*", "2"}
	actual = evaluateMultiplications(sample)
	expected = []string{"2", "+", "8.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_evaluateDivisions(t *testing.T) {
	sample := []string{"12"}
	actual := evaluateDivisions(sample)
	expected := []string{"12"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	//catch error test (division by 0
	//sample = []string{"2", "/", "0"}
	//actual = evaluateDivisions(sample)
	//expected = []string{"4.0"}
	//if !reflect.DeepEqual(actual, expected) {
	//	t.Errorf("actual = %v, expected = %v", actual, expected)
	//}

	sample = []string{"2", "/", "2"}
	actual = evaluateDivisions(sample)
	expected = []string{"1.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"2", "/", "2", "/", "2"}
	actual = evaluateDivisions(sample)
	expected = []string{"0.5"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"2", "/", "2", "/", "2", "+", "2"}
	actual = evaluateDivisions(sample)
	expected = []string{"0.5", "+", "2"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"2", "+", "2", "/", "2", "/", "2"}
	actual = evaluateDivisions(sample)
	expected = []string{"2", "+", "0.5"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_evaluateRemainders(t *testing.T) {
	sample := []string{"12"}
	actual := evaluateRemainders(sample)
	expected := []string{"12"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"10", "%", "2"}
	actual = evaluateRemainders(sample)
	expected = []string{"0.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"10", "%", "3"}
	actual = evaluateRemainders(sample)
	expected = []string{"1.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"100", "%", "13", "%", "2"}
	actual = evaluateRemainders(sample)
	expected = []string{"1.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"100", "%", "13", "%", "2", "+", "9"}
	actual = evaluateRemainders(sample)
	expected = []string{"1.0", "+", "9"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"9", "+", "100", "%", "13", "%", "2"}
	actual = evaluateRemainders(sample)
	expected = []string{"9", "+", "1.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_evaluateAdditions(t *testing.T) {
	sample := []string{"12"}
	actual := evaluateAdditions(sample)
	expected := []string{"12"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"10", "+", "2"}
	actual = evaluateAdditions(sample)
	expected = []string{"12.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"10", "+", "3"}
	actual = evaluateAdditions(sample)
	expected = []string{"13.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"100", "+", "13", "+", "2"}
	actual = evaluateAdditions(sample)
	expected = []string{"115.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"100", "+", "13", "+", "2", "/", "9"}
	actual = evaluateAdditions(sample)
	expected = []string{"115.0", "/", "9"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"9", "*", "100", "+", "13", "+", "2"}
	actual = evaluateAdditions(sample)
	expected = []string{"9", "*", "115.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_evaluateSubstractions(t *testing.T) {
	sample := []string{"12"}
	actual := evaluateSubstractions(sample)
	expected := []string{"12"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"10", "-", "2"}
	actual = evaluateSubstractions(sample)
	expected = []string{"8.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"10", "-", "3"}
	actual = evaluateSubstractions(sample)
	expected = []string{"7.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"100", "-", "13", "-", "2"}
	actual = evaluateSubstractions(sample)
	expected = []string{"85.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"100", "-", "13", "-", "2", "+", "9"}
	actual = evaluateSubstractions(sample)
	expected = []string{"85.0", "+", "9"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"9", "+", "100", "-", "13", "-", "2"}
	actual = evaluateSubstractions(sample)
	expected = []string{"9", "+", "85.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_evaluatePolarity(t *testing.T) {
	sample := "a+-b--c++d"
	actual := evaluatePolarity(sample)
	expected := "a-b+c+d"
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = "--x++y"
	actual = evaluatePolarity(sample)
	expected = "+x+y"
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_evaluateFactorials(t *testing.T) {
	sample := []string{"12"}
	actual := evaluateFactorials(sample)
	expected := []string{"12"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"2", "!", "-", "2"}
	actual = evaluateFactorials(sample)
	expected = []string{"2.0", "-", "2"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"3", "!"}
	actual = evaluateFactorials(sample)
	expected = []string{"6.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"3", "!", "+", "4", "!", "/", "5", "!"}
	actual = evaluateFactorials(sample)
	expected = []string{"6.0", "+", "24.0", "/", "120.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_evaluateExponents(t *testing.T) {
	sample := []string{"12"}
	actual := evaluateExponents(sample)
	expected := []string{"12"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"12", "^", "2"}
	actual = evaluateExponents(sample)
	expected = []string{"144.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"3", "^", "2", "+", "34"}
	actual = evaluateExponents(sample)
	expected = []string{"9.0", "+", "34"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_evaluateSquareRoots(t *testing.T) {
	sample := []string{"12"}
	actual := evaluateSquareRoots(sample)
	expected := []string{"12"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"√", "9"}
	actual = evaluateSquareRoots(sample)
	expected = []string{"3.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"√", "36", "+", "36"}
	actual = evaluateSquareRoots(sample)
	expected = []string{"6.0", "+", "36"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"3", "-", "√", "36"}
	actual = evaluateSquareRoots(sample)
	expected = []string{"3", "-", "6.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_evaluateParenthesis(t *testing.T) {
	sample := []string{"12"}
	actual := evaluateParenthesis(sample)
	expected := []string{"12"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"(", "12", ")"}
	actual = evaluateParenthesis(sample)
	expected = []string{"12.0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = []string{"12", "+", "(", "3", "+", "5", "+", "1", ")", "+", "19"}
	actual = evaluateParenthesis(sample)
	expected = []string{"12", "+", "9.0", "+", "19"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

}

func Test_evaluateCalculate(t *testing.T) {
	sample := "12"
	actual := calculate(sample)
	expected := 12.0
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = "2+2"
	actual = calculate(sample)
	expected = 4.0
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = "2x3000"
	actual = calculate(sample)
	expected = 6000.0
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = "122 + 34 * 5 / 9"
	actual = calculate(sample)
	expected = 140.88888888888889
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = "(7+3^2)x34÷4"
	actual = calculate(sample)
	expected = 136.0
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = "-7*3"
	actual = calculate(sample)
	expected = -21.0
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	sample = "+7*3"
	actual = calculate(sample)
	expected = 21.0
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_handlesDivisionByZero(t *testing.T) {
	sample := "1/0"
	actual := calculate(sample)
	expected := 0.0
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

func Test_stressTest(t *testing.T) {
	//1. Mixed Operations with Nested Parentheses
	sample := "√(100) + (3! * 2^5) / (10 % 3) - (-5 * 2)"
	actual := calculate(sample)
	expected := 212.0
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	////2. High-Precision Floating Point
	//sample = "(0.1 + 0.2) * 5^3 - √(2.25) / 0.5%3"
	//actual = calculate(sample)
	//expected = 34.5
	//if actual != expected {
	//	t.Errorf("actual = %v, expected = %v", actual, expected)
	//}

	////3. Factorial & Exponent Stress
	//sample = "10! / (5^3 * (2%7)) + √(1000000) - 3!!"
	//actual = calculate(sample)
	//expected = 15_512.2
	//if actual != expected {
	//	t.Errorf("actual = %v, expected = %v", actual, expected)
	//}

	//4. Deeply Nested Logic
	sample = "((((5 + 3) * 2)^2 % 10) / √(4)) - (2^(3! - 4))"
	actual = calculate(sample)
	expected = -1.0
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	////5. Consecutive Operators
	//sample = "5--+-+√(9) + 3%2 * 10/--2"
	//actual = calculate(sample)
	//expected = 13.0
	//if actual != expected {
	//	t.Errorf("actual = %v, expected = %v", actual, expected)
	//}

	////6. Large-Number Computation
	//sample = "999999^2 % 12345 + √(987654321) * 20! / 1000"
	//actual = calculate(sample)
	//expected = 2.43e+18
	//fmt.Println(actual)
	//if actual != expected {
	//	t.Errorf("actual = %v, expected = %v", actual, expected)
	//}

	//7. Mixed Precedence Chaos
	sample = "3 + 4 * 2 / (1 - 5)^2 + 10%3 + √(4^2!)"
	actual = calculate(sample)
	expected = 8.5
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	////8. Minimal Whitespace Challenge
	//sample = "√4+3!-2^3*5%2/1"
	//actual = calculate(sample)
	//expected = 0.0
	//if actual != expected {
	//	t.Errorf("actual = %v, expected = %v", actual, expected)
	//}

	//9. Redundant Parentheses
	sample = "((((5))) + (√((36))) / ((2%(3))) - ((2^(3))))"
	actual = calculate(sample)
	expected = 0.0
	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	////10. All Operators in One
	//sample = "√(5! + 3^4) * (10%3) - (2^(6/2)) + (-5 + 3!)"
	//actual = calculate(sample)
	//expected = 7.177
	//if actual != expected {
	//	t.Errorf("actual = %v, expected = %v", actual, expected)
	//}

}
