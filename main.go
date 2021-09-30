package main

import "fmt"

const (
	StepUp    = "up"
	StepRight = "right"
	StepDown  = "down"
)

func main() {
	var resultX []int
	var resultY []int
	mapStepsWithCoordinate := make(map[string]string)

	fmt.Println("List of treasure location probability coordinates")

	// find probable location for treasure on every step probability
	// minimum step is set to 1, you can set it to 0 if you want
	// loop step up
	minimumStep := 1
	for u := minimumStep; u <= 3; u++ {
		x, y := 0, 0

		// validate step, if hit obstacle
		if isValid(x, y, u, StepUp) {
			for r := minimumStep; r <= 5; r++ {
				if isValid(x, u, r, StepRight) {
					for d := minimumStep; d <= 3; d++ {
						if d <= u {
							if isValid(r, u, d, StepDown) {
								x = r
								y = u - d // get delta value for y coordinates
								resultX = append(resultX, x)
								resultY = append(resultY, y)

								// fmt.Printf("%d-%d-%d --> %d,%d\n", u, r, d, x, y)
								fmt.Printf("(%d,%d)\n", x, y)

								mapStepsWithCoordinate[fmt.Sprintf("%d,%d,%d", u, r, d)] = fmt.Sprintf("%d,%d", x, y)
							}

						}

					}
				}
			}
		}
	}

	var userStepUp, userStepRight, userStepDown int

	fmt.Println("\nEnter your steps: ")
	fmt.Println("Up: ")
	fmt.Scanln(&userStepUp)
	if !isValid(0, 0, userStepUp, StepUp) {
		fmt.Print("You hit an obstacle!")
		printMapGraph(resultX, resultY)

		return
	}

	fmt.Println("Right: ")
	fmt.Scanln(&userStepRight)
	if !isValid(0, userStepUp, userStepRight, StepRight) {
		fmt.Print("You hit an obstacle!")
		printMapGraph(resultX, resultY)

		return
	}

	fmt.Println("Down: ")
	fmt.Scanln(&userStepDown)
	if !isValid(userStepRight, userStepUp, userStepDown, StepDown) {
		fmt.Print("You hit an obstacle!")
		printMapGraph(resultX, resultY)

		return
	}

	// todo check if user hit the probable treasure location
	userStep := fmt.Sprintf("%d,%d,%d", userStepUp, userStepRight, userStepDown)
	if _, ok := mapStepsWithCoordinate[userStep]; !ok {
		fmt.Println("\nUnfortunately you miss the treasure location!")
	} else {
		fmt.Println("\nCongratulation, you find the treasure!")
	}

	printMapGraph(resultX, resultY)

	return
}

func isBlocker(x, y int) bool {
	blockX := []int{1, 3, 5, 1, 2, 3}
	blockY := []int{0, 1, 1, 2, 2, 2}
	valid := false
	for i := 0; i < 6; i++ {
		if blockX[i] == x && blockY[i] == y {
			return true
		}
	}
	return valid
}

func isResult(blockX, blockY []int, x, y int) bool {
	valid := false
	for i := 0; i < len(blockX); i++ {
		if blockX[i] == x && blockY[i] == y {
			return true
		}
	}
	return valid

}

func isValid(x, y, moveCount int, move string) bool {
	// validate step, cant more than 3 step in y axis because of map size
	if (move == StepUp || move == StepDown) && moveCount > 3 {
		return false
	}

	// validate step, cant more than 5 in x axis step because of map size
	if move == StepRight && moveCount > 5 {
		return false
	}

	// define obstacle
	blockX := []int{1, 3, 5, 1, 2, 3}
	blockY := []int{0, 1, 1, 2, 2, 2}

	valid := true
	if move == StepUp {
		for m := 1; m <= moveCount; m++ {
			MoveY := y + m
			for i := 0; i < 6; i++ {
				if blockX[i] == x && blockY[i] == MoveY {
					return false
				}
			}

		}
	} else if move == StepDown {
		for m := 1; m <= moveCount; m++ {
			MoveY := y - m
			for i := 0; i < 6; i++ {
				if blockX[i] == x && blockY[i] == MoveY {
					return false
				}
			}

		}

	} else {
		for m := 1; m <= moveCount; m++ {
			MoveX := x + m
			for i := 0; i < 6; i++ {
				if blockX[i] == MoveX && blockY[i] == y {
					return false
				}
			}

		}
	}

	return valid

}

func printMapGraph(resultX, resultY []int) {
	fmt.Println("\nMap graph with probable treasure locations")
	for i := 5; i >= 0; i-- {

		for j := 0; j <= 7; j++ {
			if i == 5 || i == 0 || j == 0 || j == 7 {
				fmt.Print("#")
			} else {
				if isBlocker(j-1, i-1) {
					fmt.Print("#")
				} else if isResult(resultX, resultY, j-1, i-1) {
					fmt.Print("$")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}
