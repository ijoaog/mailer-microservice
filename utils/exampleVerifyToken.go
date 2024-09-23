package utils

import (
	"fmt"
	"os"
)

func ValidateToken(tokenString string) bool {
	var staticToken = os.Getenv("TOKEN")
	fmt.Println("Token estático:", staticToken)

	return tokenString == staticToken
}
