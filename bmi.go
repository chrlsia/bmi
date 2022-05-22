package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	//Output information
	fmt.Println("BMI calculator")
	fmt.Println("----------------------")

	//Prompt for user input (weight + height)
	fmt.Print("Ente your weight (kg): ")
	weightInput,_:=reader.ReadString('\n')

	fmt.Print("Ente your height (m): ")
	heightInput,_:=reader.ReadString('\n')

	//Save that user input in variables
	weightInput = strings.Replace(weightInput,"\n","",-1)
	heightInput = strings.Replace(heightInput,"\n","",-1)

	weight,_:= strconv.ParseFloat(weightInput,64)
	height,_:= strconv.ParseFloat(heightInput,64)

	//Calculate the bmi (weight/(height * height))
	bmi:= weight/(height * height)
	//Output the calculated bmi 
	fmt.Printf("Your BMI : %.2f\n",bmi)
}