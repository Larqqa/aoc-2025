package main

import (
	lib "aoc/2025"
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := lib.ReadFile("inputs/dayten.txt")

	p1 := solvePartOne(data)
	fmt.Println("Part One:", p1)

	p2 := solvePartTwo(data)
	fmt.Println("Part Two:", p2)
}

type Machine struct {
	LightDiagram        []bool
	ButtonWiringSchema  [][]int
	JoltageRequirements []int
}

func parse(data string) []Machine {
	parsed := strings.Split(strings.ReplaceAll(data, "\r", ""), "\n")

	machines := make([]Machine, len(parsed))
	for i, line := range parsed {
		first := strings.Split(line, "]")

		lightStr := strings.Replace(first[0], "[", "", 1)
		lights := make([]bool, len(lightStr))

		for i, char := range lightStr {
			if char == '#' {
				lights[i] = true
			} else {
				lights[i] = false
			}
		}

		second := strings.Split(first[1], "{")

		joltageChars := strings.Split(strings.Replace(second[1], "}", "", 1), ",")
		joltage := make([]int, len(joltageChars))
		for i, char := range joltageChars {
			val, _ := strconv.Atoi(strings.TrimSpace(char))
			joltage[i] = val
		}

		buttonStrs := strings.Split(strings.TrimSpace(second[0]), " ")
		buttons := make([][]int, len(buttonStrs))

		for i, char := range buttonStrs {
			vals := strings.SplitSeq(strings.ReplaceAll(strings.ReplaceAll(char, "(", ""), ")", ""), ",")
			for v := range vals {
				val, _ := strconv.Atoi(v)
				buttons[i] = append(buttons[i], val)
			}
		}
		machines[i] = Machine{
			LightDiagram:        lights,
			ButtonWiringSchema:  buttons,
			JoltageRequirements: joltage,
		}
	}

	return machines
}

// Find all button combos of unique buttons that yield the target state
func findAllButtonCombos(target []bool, machine Machine) [][][]int {
	allPaths := [][][]int{}

	stateKey := func(state []bool, usedButtons map[int]bool) string {
		return fmt.Sprintf("%v-%v", state, usedButtons)
	}

	seenStates := make(map[string]bool)
	var dfs func(state []bool, path [][]int, usedButtons map[int]bool)
	dfs = func(state []bool, path [][]int, usedButtons map[int]bool) {
		key := stateKey(state, usedButtons)
		if seenStates[key] {
			return
		}

		seenStates[key] = true

		if slices.Equal(state, target) {
			pathCopy := make([][]int, len(path))
			copy(pathCopy, path)
			allPaths = append(allPaths, pathCopy)
			return
		}

		for idx, button := range machine.ButtonWiringSchema {
			if usedButtons[idx] {
				continue
			}

			newState := make([]bool, len(state))
			copy(newState, state)

			for _, i := range button {
				newState[i] = !newState[i]
			}

			newUsed := make(map[int]bool)
			maps.Copy(newUsed, usedButtons)
			newUsed[idx] = true

			dfs(newState, append(path, button), newUsed)
		}
	}

	dfs(make([]bool, len(target)), [][]int{}, make(map[int]bool))

	return allPaths
}

func solvePartOne(data string) int {
	machines := parse(data)

	sum := 0
	for _, m := range machines {
		allButtons := findAllButtonCombos(m.LightDiagram, m)
		slices.SortFunc(allButtons, func(a, b [][]int) int {
			return len(a) - len(b)
		})
		sum += len(allButtons[0])
	}

	return sum
}

func convertJoltageRegToParity(state []int) []bool {
	parity := make([]bool, len(state))
	for i, val := range state {
		parity[i] = val%2 == 1
	}
	return parity
}

func solveMachine(machine Machine) int {
	cache := make(map[string]int)

	var solve func(target []int) int
	solve = func(target []int) int {
		allZero := 0
		for _, val := range target {
			if val < 0 {
				return math.MaxInt32
			}
			allZero += val
		}
		if allZero == 0 {
			return 0
		}

		stateKey := fmt.Sprint(target)
		if val, exists := cache[stateKey]; exists {
			return val
		}

		parity := convertJoltageRegToParity(target)
		buttons := findAllButtonCombos(parity, machine)

		if len(buttons) == 0 {
			cache[stateKey] = math.MaxInt32
			return math.MaxInt32
		}

		minCost := math.MaxInt32
		for _, btnCombo := range buttons {
			newT := make([]int, len(target))
			copy(newT, target)

			for _, button := range btnCombo {
				for _, i := range button {
					newT[i]--
				}
			}

			valid := true
			for _, val := range newT {
				if val < 0 || val%2 != 0 {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			for i := range newT {
				newT[i] /= 2
			}

			recursiveCost := solve(newT)
			if recursiveCost >= math.MaxInt32 {
				continue
			}

			minCost = min(minCost, len(btnCombo)+2*recursiveCost)
		}

		cache[stateKey] = minCost
		return minCost
	}

	return solve(machine.JoltageRequirements)
}

func solvePartTwo(data string) int {
	machines := parse(data)

	sum := 0
	for i, m := range machines {
		fmt.Println("Machine", i+1)
		sum += solveMachine(m)
	}

	return sum
}
