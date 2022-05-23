package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chrlsia/bmi/info"
)


func main() {
	//Output information
	info.PrintWelcome()

	//Prompt for user input (weight + height)
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	//Save that user input in variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	//Calculate the bmi (weight/(height * height))
	bmi := weight / (height * height)
	//Output the calculated bmi
	fmt.Printf("Your BMI : %.2f\n", bmi)
}

