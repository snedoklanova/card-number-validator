package validator

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

// validator represents data about a card number validator.
type validator struct {
	Number string `json:"number"`
	Result bool   `json:"result"`
}

// GetValidationByCardNumber check whose Number value matches the rules
// parameter sent by the client, then returns true or false.
func GetValidationByCardNumber(c *gin.Context) {
	numString := c.Param("number")

	var result bool

	if len(numString) == 13 || len(numString) == 16 {
		result = validation(numString)
	} else if len(numString) == 15 {
		result = amexValidation(numString)
	} else {
		result = false
	}

	var response = validator{
		Number: numString,
		Result: result,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func amexValidation(numString string) bool {
	prefix := numString[:2]
	if prefix == "34" || prefix == "37" {
		result := validate(numString)

		return result
	}
	return false
}

func validation(numString string) bool {
	if len(numString) == 16 {
		result := validate(numString)

		return result
	}
	return false
}

func numToDigits(numString string) []int {
	number := stringToInt(numString)
	digits := []int{}

	for number > 0 {
		digits = append([]int{number % 10}, digits...)
		number /= 10
	}

	return digits
}

func stringToInt(numString string) int {
	number, err := strconv.Atoi(numString)
	if err != nil {
		// ... handle error
		panic(err)
	}

	return number
}

func validate(numString string) bool {
	digits := numToDigits(numString)
	slices.Reverse(digits)
	sum := 0

	fmt.Println("digits:", digits)

	for index, value := range digits {
		if index%2 != 0 {
			digits[index] = value * 2
		}

		if digits[index] > 9 {
			digits[index] = digits[index] - 9
		}

		sum += digits[index]
	}

	fmt.Println("digits:", digits)
	fmt.Println("sum:", sum)

	if sum%10 == 0 {
		return true
	}

	return false
}
