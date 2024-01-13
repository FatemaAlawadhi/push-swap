package main

import (
	"fmt"
	"os"
	"push-swap/Organizer"
	"strconv"
	"strings"
)

func main() {
	arg := os.Args
	if len(arg) < 2 {
		return
	} else {
		var Values []int
		var Valuesstr []string

		if len(arg)>2 {
			Valuesstr =arg[1:]
		}else {
			Valuesstr = strings.Split(arg[1], " ")
		}
		for i := 0; i < len(Valuesstr); i++ {
			Num, err := strconv.Atoi(Valuesstr[i])
			if err != nil {
				fmt.Println("Error: Invalid input:", Valuesstr[i], "is not a valid number.")
				return
			}
			for i := 0; i < len(Values); i++ {
				if Num == Values[i] {
					fmt.Println("Error: Do not repeat the same number")
					return
				}
			}
			Values = append(Values, Num)

		}

		Arranged:=Check(Values)
		if Arranged == true {
			return
		}

		ArrangedValues, UsedInstructions, InstructionCount := Organizer.Organizer(Values)

		fmt.Println("Values in ascending order: ", ArrangedValues)
		fmt.Println("Instructions used: ", UsedInstructions)
		fmt.Println("Total Number of Instructiona: ", InstructionCount)

	}
}

func Check(Values []int) bool{
	for i := 0; i < len(Values)-1; i++ {
		if Values[i] > Values[i+1] {
			return false
		}
	}
	return true
}