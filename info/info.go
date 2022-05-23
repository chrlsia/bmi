package info

import "fmt"

const mainTitle = "BMI calculator"
const separator = "----------------------"
const WeightPrompt = "Enter your weight (kg): "
const HeightPrompt = "Enter your height (m): "

func PrintWelcome(){
	fmt.Println(mainTitle)
	fmt.Println(separator)
}