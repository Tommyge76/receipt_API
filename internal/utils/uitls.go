package utils

import (
	"github.com/google/uuid"
	"process_receipts/internal/request_models"
	"math"
	"strconv"
	"unicode"
	"strings"
	"time"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func CalculatePoints(receipt request_models.Receipt) int {
	points := 0

	// Calculate points for each alphanumeric character in the retailer name
	points += CalculateAlphaNumericCharPoints(receipt.Retailer)

	// Convert receipt.Total from string to float64
	totalStr := receipt.Total
	total, err := strconv.ParseFloat(totalStr, 64)
	if err != nil {
		return points
	}

	// Add 50 points if the total is a whole dollar amount
	if total == math.Floor(total) {
		points += 50
	}

	// Add 25 points if the total is a multiple of 0.25
	if total != 0 && math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Add 5 points for every two items on the receipt
	points += int(math.Floor(float64(len(receipt.Items)) / 2)) * 5

	// If description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer
	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDescription)
		if len(description) % 3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}

	// Add 6 points if the day is odd
	day := strings.Split(receipt.PurchaseDate, "-")[2]
	dayInt, err := strconv.Atoi(day)
	if err == nil && dayInt % 2 != 0 {
		points += 6
	}

	// Add 10 points if the time is after 2:00pm and before 4:00pm
	time := strings.Split(receipt.PurchaseTime, ":")
	hour, _ := strconv.Atoi(time[0])
	minute, _ := strconv.Atoi(time[1])
	if hour >= 14 && hour < 16 && minute > 0 && minute < 60 {
		points += 10
	}
	return points
}

func CalculateAlphaNumericCharPoints(str string) int {
	points := 0
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}
	return points
}

func ValidateReceipt(receipt request_models.Receipt) bool {
	if receipt.Retailer == "" ||
		receipt.PurchaseDate == "" ||
		receipt.Total == "" ||
		receipt.PurchaseTime == "" ||
		len(receipt.Items) == 0 {
		return false
	}

	// Validate PurchaseDate is in format YYYY-MM-DD
	_, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return false
	}

	// Validate purchase time is in format HH:MM
	_, err = time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return false
	}

	// Validate total is a floating point number
	_, err = strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return false
	}

	// validate items have a description and price
	for _, item := range receipt.Items {
		if item.ShortDescription == "" || item.Price == "" {
			return false
		}
		_, err = strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return false
		}
	}

	return true
}
