package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"liteWeightCryptoConverter/internal/converter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s <amount> <from_currency> <to_currency>", os.Args[0])
	}

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("Invalid amount: %s", os.Args[1])
	}

	fromCurrency := os.Args[2]
	toCurrency := os.Args[3]

	rate, err := converter.GetConversionRate(fromCurrency, toCurrency)
	if err != nil {
		log.Fatalf("Error fetching conversion rate: %v", err)
	}

	result := amount * rate
	fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
}
