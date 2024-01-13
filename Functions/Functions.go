package Functions

//import "fmt"

var (
	InstructionCount int
	UsedInstructions []string
)

func Instructions()(int, []string) {
	return InstructionCount, UsedInstructions
}

func Pa(SliceA []int, SliceB []int) ([]int, []int) {
	if len(SliceB) > 0 {
		TempNum := SliceB[0]
		SliceA = append([]int{TempNum}, SliceA...)
		SliceB = SliceB[1:]
	}
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "pa")

	return SliceA, SliceB
}

func Pb(SliceA []int, SliceB []int) ([]int, []int) {
	if len(SliceA) > 0 {
		TempNum := SliceA[0]
		SliceB = append([]int{TempNum}, SliceB...)
		SliceA = SliceA[1:]
	}
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "pb")
	return SliceA, SliceB
}


func Sa(SliceA []int) []int {
	if len(SliceA) >= 2 {
		TempNum := SliceA[0]
		SliceA[0] = SliceA[1]
		SliceA[1] = TempNum
	}
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "sa")
	return SliceA
}

func Sb(SliceB []int) []int {
	if len(SliceB) >= 2 {
		TempNum := SliceB[0]
		SliceB[0] = SliceB[1]
		SliceB[1] = TempNum
	}
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "sb")
	return SliceB
}


func Ss(SliceA []int, SliceB []int) ([]int, []int) {
	SliceA = Sa(SliceA)
	SliceB= Sb(SliceB)
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "ss")
	return SliceA, SliceB
}

func Ra(SliceA []int) []int {
	if len(SliceA) >= 2 {
		firstElement := SliceA[0]
		copy(SliceA, SliceA[1:])
		SliceA[len(SliceA)-1] = firstElement
	}
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "ra")
	return SliceA
}



func Rb(SliceB []int) []int {
	if len(SliceB) >= 2 {
		firstElement := SliceB[0]
		copy(SliceB, SliceB[1:])
		SliceB[len(SliceB)-1] = firstElement
	}
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "rb")
	return SliceB
}

func Rr(SliceA []int, SliceB []int) ([]int, []int) {
	SliceA = Ra(SliceA)
	SliceB= Rb(SliceB)
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "rr")
	return SliceA, SliceB
}

func Rra(SliceA []int) []int {
	if len(SliceA) >= 2 {
		lastIndex := len(SliceA) - 1
		lastElement := SliceA[lastIndex]
		copy(SliceA[1:], SliceA[:lastIndex])
		SliceA[0] = lastElement
	}
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "rra")
	return SliceA
}

func Rrb(SliceB []int) []int {
	if len(SliceB) >= 2 {
		lastIndex := len(SliceB) - 1
		lastElement := SliceB[lastIndex]
		copy(SliceB[1:], SliceB[:lastIndex])
		SliceB[0] = lastElement
	}
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "rrb")
	return SliceB
}


func Rrr(SliceA []int, SliceB []int) ([]int, []int) {
	SliceA = Rra(SliceA)
	SliceB= Rrb(SliceB)
	InstructionCount += 1
	UsedInstructions = append(UsedInstructions, "rrr")
	return SliceA, SliceB
}