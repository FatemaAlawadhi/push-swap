package main

import ("fmt"
        "os"
	    "bufio"
	    "push-swap/Functions"
	    "strings"
	    "strconv")

func main() {
	arg := os.Args
	if len(arg) < 2 {
		return
	} else {
		var Values []int
		var ValuesSlice []string

		if len(arg)>2 {
			ValuesSlice=arg[1:]
		}else {
			ValuesSlice = strings.Split(arg[1], " ")
		}

		for i := 0; i < len(ValuesSlice); i++ {
			Num, err := strconv.Atoi(ValuesSlice[i])
			if err != nil {
				fmt.Println("Error: Invalid input:", ValuesSlice[i], "is not a valid number.")
				return
			}
			for j := 0; j < len(Values); j++ {
				if Num == Values[j] {
					fmt.Println("Error: Do not repeat the same number")
					return
				}
			}
			Values = append(Values, Num)
		}

		var Instructions []string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			instruction := scanner.Text()
			Instructions = append(Instructions, instruction)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error occurred while reading input:", err)
			return
		}
		
		
		var SliceA []int
		if strings.HasPrefix(Instructions[1], "Instructions used:") {
			instructionsString := Instructions[1]
			instructionsString = strings.TrimPrefix(instructionsString, "Instructions used:")
			instructionsString = strings.TrimSpace(instructionsString)
			instructionsString = strings.TrimPrefix(instructionsString, "[")
			instructionsString = strings.TrimSuffix(instructionsString, "]")
			Instructions = strings.Split(instructionsString, " ")
		}

		SliceA = Values
		var SliceB []int
		for i := 0; i < len(Instructions); i++ {
			SliceA, SliceB = ApplyInstruction(SliceA, SliceB, Instructions[i])
		}

		
		Arranged := Check(SliceA)
		if Arranged == true {
			fmt.Println("OK")
			
		} else {
			fmt.Println("KO")
			
		}
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

func ApplyInstruction(SliceA []int,SliceB []int, Instruction string)([]int,[]int) {
	if Instruction == "pa" {
		SliceA, SliceB = Functions.Pa(SliceA, SliceB)
	}else if Instruction == "pb" {
		SliceA, SliceB = Functions.Pb(SliceA, SliceB)
	}else if Instruction == "ss" {
		SliceA, SliceB= Functions.Ss(SliceA, SliceB)
	}else if Instruction == "sa" {
		SliceA = Functions.Sa(SliceA)
	}else if Instruction == "sb" {
		SliceB = Functions.Sb(SliceB)
	}else if Instruction == "rr" {
		SliceA, SliceB = Functions.Rr(SliceA, SliceB)
	}else if Instruction == "ra" {
		SliceA = Functions.Ra(SliceA)
	}else if Instruction == "rb" {
		SliceB = Functions.Rb(SliceB)
	}else if Instruction == "rrr" {
		SliceA, SliceB = Functions.Rrr(SliceA, SliceB)
	}else if Instruction == "rra" {
		SliceA = Functions.Rra(SliceA)
	}else if Instruction == "rrb" {
		SliceB = Functions.Rrb(SliceB)
	}
	return SliceA,SliceB
}