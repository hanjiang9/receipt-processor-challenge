package main

import (
	"fmt"
	"math"
	"strings"
	"github.com/google/uuid"
	"strconv"
)

func generateID() string {
	return uuid.New().String()
}

func calculatePoints(receipt Receipt) int {
	points := 0

	// One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}
	fmt.Printf("Retailer name points: %d\n", points)

	// 50 points if the total is a round dollar amount with no cents.
	if isRoundDollar(receipt.Total) {
		points += 50
		fmt.Println("Added 50 points for round dollar total")
	}

	// 25 points if the total is a multiple of 0.25.
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
		fmt.Println("Added 25 points for total being a multiple of 0.25")
	}

	// 5 points for every two items on the receipt.
	itemPairs := len(receipt.Items) / 2
	points += itemPairs * 5
	fmt.Printf("Added %d points for %d item pairs\n", itemPairs*5, itemPairs)

	// Points for item description length.
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			itemPoints := int(math.Ceil(price * 0.2))
			points += itemPoints
			fmt.Printf("Added %d points for item '%s'\n", itemPoints, item.ShortDescription)
		}
	}

	// 6 points if the day in the purchase date is odd.
	if isOddDay(receipt.PurchaseDate) {
		points += 6
		fmt.Println("Added 6 points for odd purchase day")
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if isBetweenTwoAndFour(receipt.PurchaseTime) {
		points += 10
		fmt.Println("Added 10 points for purchase time between 2:00pm and 4:00pm")
	}

	fmt.Printf("Total points: %d\n", points)
	return points
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func isRoundDollar(total string) bool {
	value, err := strconv.ParseFloat(total, 64)
	return err == nil && value == float64(int(value))
}

func isMultipleOfQuarter(total string) bool {
	value, err := strconv.ParseFloat(total, 64)
	return err == nil && math.Mod(value, 0.25) == 0
}

func totalGreaterThanTen(total string) bool {
	value, err := strconv.ParseFloat(total, 64)
	return err == nil && value > 10.00
}

func isOddDay(date string) bool {
	day := strings.Split(date, "-")[2]
	dayInt, err := strconv.Atoi(day)
	return err == nil && dayInt%2 != 0
}

func isBetweenTwoAndFour(timeStr string) bool {
	hour := strings.Split(timeStr, ":")[0]
	hourInt, err := strconv.Atoi(hour)
	return err == nil && hourInt >= 14 && hourInt < 16
}