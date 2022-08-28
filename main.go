package main

import (
	"fmt"
	"strconv"
)

func getWeight() float32 {
	for {
		fmt.Printf("What's your weight? (In Kilograms)\n")

		var input string
    _, err := fmt.Scanln(&input)

    if err != nil {
			fmt.Printf("Invalid weight, try again!\n")

			continue
    }

		inputHeight, err := strconv.Atoi(input)

		if err != nil {
			fmt.Printf("Invalid weight, try again!\n")

			continue
		}

		if inputHeight < 1 || inputHeight > 300 {
			fmt.Printf("Invalid height, try again!\n")

			continue
		}

		return float32(inputHeight)
	}
}

func getHeight() float32 {
	for {
		fmt.Printf("What's your height? (In centimeters, between 50 and 300)\n")

		var input string
    _, err := fmt.Scanln(&input)

    if err != nil {
			fmt.Printf("Invalid height, try again!\n")

			continue
    }

		inputHeight, err := strconv.Atoi(input)

		if err != nil {
			fmt.Printf("Invalid height, try again!\n")

			continue
		}

		if inputHeight < 50 || inputHeight > 300 {
			fmt.Printf("Invalid height, try again!\n")

			continue
		}

		return float32(inputHeight) / 100
	}
}

func getNutritionalStatus(bmi float32) string {
  if bmi < 18.5 {
    return "Underweight"
  }

	if bmi < 25 {
    return "Normal weight"
  }

	if bmi < 30 {
    return "Pre-obesity"
  }

	if bmi < 35 {
    return "Obesity class I"
  }

	if bmi < 40 {
    return "Obesity class II"
  }

	return "Obesity class III"
}

func main() {
	fmt.Printf("Check your BMI!\n")

	weight := getWeight()
	height := getHeight()

	bmi := weight / (height * height)

	nutritionalStatus := getNutritionalStatus(bmi)

	fmt.Printf("Your BMI is: %.2f\n", bmi)
	fmt.Printf("Your nutritional status is: %s\n", nutritionalStatus)
}
