package validator

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// validator represents data about a card number validator.
type validator struct {
    Number  string  `json:"number"`
    Result bool  `json:"result"`
}

// GetValidationByCardNumber check whose Number value matches the rules
// parameter sent by the client, then returns true or false.
func GetValidationByCardNumber(c *gin.Context) {
    		number := c.Param("number")

				fmt.Printf("Type of a is %T\n", number)
        fmt.Println("Received number:", number)

				var result bool;

				if len(number) == 13 || len(number) == 16  {
					result = validation(number)
				} else if len(number) == 15 {
					result = amexValidation(number)
				} else {
					result = false
				}

        var response = validator{
            Number: number,
            Result: result,
        }

        c.IndentedJSON(http.StatusOK, response)
}

func amexValidation (number string) bool {
	prefix := number[:2]
	if prefix == "34" || prefix == "37" {
		return true
	}
	return false
}

func validation (number string) bool {
	if len(number) == 16 {
		prefix := number[:2]
		if prefix == "51" || prefix == "52" || prefix == "53" || prefix == "54" || prefix == "55" {
			return true
		}
	}
	prefix := number[:1]
	if prefix == "4" {
		return true
	}
	return false
}

func splitInt(n int) []int {
	digits := []int{}

	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}

	return digits
}

func stringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Error during conversion: %v\n", err)
		return -1
	}

	return num
}
