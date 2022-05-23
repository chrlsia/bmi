package main

import (
	"fmt"

	"github.com/chrlsia/bmi/info"
)


func main() {
	//Output information
	info.PrintWelcome()

	weight, height:= getUserMetrics()
	

	//Calculate the bmi (weight/(height * height))
	bmi := weight / (height * height)
	//Output the calculated bmi
	fmt.Printf("Your BMI : %.2f\n", bmi)
}

