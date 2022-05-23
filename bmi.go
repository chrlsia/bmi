package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const mainTitle = "BMI calculator"
const separator = "----------------------"
const weightPrompt = "Ente your weight (kg): "
const heightPrompt = "Ente your height (m): "

func main() {
	//Output information
	fmt.Println(mainTitle)
	fmt.Println(separator)

	//Prompt for user input (weight + height)
	fmt.Print(weightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(heightPrompt)
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

// const : pairnoyn timew sthn diarkeia toy 
// compile time
// varaibles: pairnoyn timew sthn diarkeia toy run time