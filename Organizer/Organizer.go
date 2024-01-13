package Organizer

import (
	"push-swap/Functions"
)


func Organizer(SliceA []int) ([]int,[]string,int) {
	var SliceB []int
	if len(SliceA) <= 3 {
		if len(SliceA) == 2 {
			if SliceA[0] > SliceA[1] {
				SliceA = Functions.Sa(SliceA)
			}
		} else if len(SliceA) == 3 {
			SliceA, _ = ArrangeMap(SliceA, SliceB, 0,0)
		}
	} else if len(SliceA) > 3 {
		//Step 1: Shift two values to B
		SliceA, SliceB = Functions.Pb(SliceA, SliceB)
		SliceA, SliceB = Functions.Pb(SliceA, SliceB)

		for len(SliceA) > 3 {
			FewerSteps, RaMoves, RbMoves, RrMoves, RraMoves, RrbMoves, RrrMoves := CheapestNumber(0, SliceA, SliceB)
			var Steps int
			for i := 0; i < len(SliceA); i++ {
				Steps, _, _, _, _, _, _ = CheapestNumber(i, SliceA, SliceB)
				if Steps < FewerSteps {
					FewerSteps, RaMoves, RbMoves, RrMoves, RraMoves, RrbMoves, RrrMoves = CheapestNumber(i, SliceA, SliceB)
				}
			}

			SliceA = RepeatFunction1Array(RaMoves, Functions.Ra, SliceA)
			SliceB = RepeatFunction1Array(RbMoves, Functions.Rb, SliceB)
			SliceA,SliceB = RepeatFunction2Array(RrMoves, Functions.Rr, SliceA,SliceB)
			SliceA = RepeatFunction1Array(RraMoves, Functions.Rra, SliceA)
			SliceB = RepeatFunction1Array(RrbMoves, Functions.Rrb, SliceB)
			SliceA,SliceB = RepeatFunction2Array(RrrMoves, Functions.Rrr, SliceA,SliceB)

			SliceA,SliceB = Functions.Pb(SliceA,SliceB)

		}

		//To Organize B
		LenofB := len(SliceB)
		MaximumB := SliceB[0]
		MaximumBIndex := 0
		for i := 0; i < LenofB; i++ {
			if SliceB[i] > MaximumB {
				MaximumB = SliceB[i]
				MaximumBIndex = i
			}
		}

		RbforB := 0
		RrbforB := 0
		if MaximumBIndex <= LenofB/2 {
			RbforB = MaximumBIndex
		} else if MaximumBIndex > LenofB/2 {
			RrbforB = LenofB - MaximumBIndex
		}
		SliceA,SliceB = ArrangeMap(SliceA,SliceB, RbforB, RrbforB)

		//
		for len(SliceB) != 0 {
			Ra, Rra := BacktoA(SliceA,SliceB)
			SliceA = RepeatFunction1Array(Ra, Functions.Ra, SliceA)
			SliceA = RepeatFunction1Array(Rra, Functions.Rra, SliceA)
			SliceA,SliceB = Functions.Pa(SliceA,SliceB)
		}

		for SliceA[0] > SliceA[len(SliceA)-1] {
			SliceA = Functions.Rra(SliceA)
		}
	}

	InstructionCount, UsedInstructions := Functions.Instructions()

	return SliceA, UsedInstructions, InstructionCount
}

func RepeatFunction1Array(num int, Operation func([]int) []int, Slice []int) []int {
	for i := 0; i < num; i++ {
		Slice = Operation(Slice)
	}

	return Slice
}

func RepeatFunction2Array(num int, Operation func([]int,[]int)([]int,[]int), SliceA []int, SliceB []int) ([]int,[]int){
	for i := 0; i < num; i++ {
		SliceA,SliceB = Operation(SliceA,SliceB)
	}

	return SliceA,SliceB
}

func Checker(SliceA []int,SliceB []int) bool {
	LenofA:= len(SliceA)
	LenofB:= len(SliceB)

	for i := 0; i < LenofA-1; i++ {
		if SliceA[0] > SliceA[i+1] {
			return false
		}
	}
	
	for i := 0; i < LenofB-1; i++ {
		if SliceB[i] < SliceB[i+1] {
			return false
		}
	}

	return true
}

func CheapestNumber(APosition int, SliceA []int, SliceB []int) (int, int, int, int, int, int, int) {
	MinimumB := SliceB[0]
	MaximumB := SliceB[0]
	MaximumBIndex := 0
	LenofB := len(SliceB)
	AValue := SliceA[APosition]
	var LessThanAIndex int
	LTADifference := 10000
	for i := 0; i < LenofB; i++ {
		if SliceB[i] > MaximumB {
			MaximumB = SliceB[i]
			MaximumBIndex = i
		}
		if SliceB[i] < MinimumB {
			MinimumB = SliceB[i]
		}

		DifferenceTemp := AValue - SliceB[i]
		if DifferenceTemp > 0 {
			if DifferenceTemp < LTADifference {
				LessThanAIndex = i
				LTADifference = DifferenceTemp
			}
		}

	}

	TotalSteps := 0
	Ra := 0
	Rb := 0
	Rra := 0
	Rrb := 0
	Rr := 0
	Rrr := 0

	LenofA := len(SliceA)
	if APosition <= LenofA/2 {
		Ra = APosition
	} else if APosition > LenofA/2 {
		Rra = LenofA - APosition
	}

	if AValue < MinimumB {
		LessThanAIndex = MaximumBIndex
	}
	if LessThanAIndex <= LenofB/2 {
		Rb = LessThanAIndex
	} else if LessThanAIndex > LenofB/2 {
		Rrb = LenofB - LessThanAIndex
	}

	if Ra == Rb {
		Rr = Ra
		Ra = 0
		Rb = 0
	} else if Ra > Rb {
		Ra = Ra - Rb
		Rr = Rb
		Rb = 0
	} else if Ra < Rb {
		Rb = Rb - Ra
		Rr = Ra
		Ra = 0
	}

	if Rra == Rrb {
		Rrr = Rra
		Rra = 0
		Rrb = 0
	} else if Rra > Rrb {
		Rra = Rra - Rrb
		Rrr = Rrb
		Rrb = 0
	} else if Rra < Rrb {
		Rrb = Rrb - Rra
		Rrr = Rra
		Rra = 0
	}

	TotalSteps = Ra + Rb + Rr + Rra + Rrb + Rrr

	return TotalSteps, Ra, Rb, Rr, Rra, Rrb, Rrr
}

func ArrangeMap(SliceA []int,SliceB []int, RbforB int, RrbforB int) ([]int,[]int) {
	checker := false
	LenofA := len(SliceA)
	for checker == false {
		if SliceA[0] > SliceA[LenofA-1] && RrbforB > 0 {
			SliceA,SliceB = Functions.Rrr(SliceA,SliceB)
			RrbforB -= 1
		} else if SliceA[0]  > SliceA[LenofA-1] && RbforB > 0 {
			SliceA,SliceB = Functions.Rr(SliceA,SliceB)
			RbforB -= 1
		} else if SliceA[0] > SliceA[LenofA-1] {
			SliceA = Functions.Ra(SliceA)
		} else if SliceA[0] > SliceA[1] {
			SliceA = Functions.Sa(SliceA)
		} else if RbforB > 0 {
			SliceB = Functions.Rb(SliceB)
			RbforB -= 1
		} else if RrbforB > 0 {
			SliceB = Functions.Rrb(SliceB)
			RrbforB -= 1
		} else {
			SliceA = Functions.Rra(SliceA)
		}

		checker = Checker(SliceA,SliceB)
	}
	return SliceA,SliceB
}

func BacktoA(SliceA []int, SliceB []int) (int, int) {
	MinimumA := SliceA[0]
	MaximumA := SliceA [0]
	MinimumAIndex := 0
	BValue := SliceB[0]
	var MoreThanBIndex int
	MTBDifference := 10000
	LenofA := len(SliceA)
	for i := 0; i < LenofA; i++ {
		if SliceA[i] > MaximumA {
			MaximumA = SliceA[i]
		}
		if SliceA[i] < MinimumA {
			MinimumA = SliceA[i]
			MinimumAIndex = i
		}
		DifferenceTemp := SliceA[i] - BValue
		if DifferenceTemp > 0 {
			if DifferenceTemp < MTBDifference {
				MoreThanBIndex = i
				MTBDifference = DifferenceTemp
			}
		}

	}
	Ra := 0
	Rra := 0
	if BValue > MaximumA {
		MoreThanBIndex = MinimumAIndex
	}

	if MoreThanBIndex <= LenofA/2 {
		Ra = MoreThanBIndex
	} else if MoreThanBIndex > LenofA/2 {
		Rra = LenofA - MoreThanBIndex
	}

	return Ra, Rra
}
