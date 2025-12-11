package main

import (
	lib "aoc/2025"
	"fmt"
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

func findFewestPresses(state []bool, machine Machine) int {
	type State struct {
		lights  []bool
		presses int
	}

	queue := []State{{lights: state, presses: 0}}
	visited := make(map[string]bool)

	stateKey := func(s []bool) string {
		return fmt.Sprint(s)
	}

	visited[stateKey(state)] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if slices.Equal(current.lights, machine.LightDiagram) {
			return current.presses
		}

		for _, button := range machine.ButtonWiringSchema {
			newState := make([]bool, len(current.lights))
			copy(newState, current.lights)

			// Flip the lights
			for _, i := range button {
				newState[i] = !newState[i]
			}

			key := stateKey(newState)
			if !visited[key] {
				visited[key] = true
				queue = append(queue, State{lights: newState, presses: current.presses + 1})
			}
		}
	}

	return -1 // If no solution found
}

func solvePartOne(data string) int {
	machines := parse(data)

	sum := 0
	for _, m := range machines {
		sum += findFewestPresses(make([]bool, len(m.LightDiagram)), m)
	}

	return sum
}

func copyMachine(machine Machine) Machine {
	initialMachine := Machine{
		LightDiagram:        append([]bool(nil), machine.LightDiagram...),
		ButtonWiringSchema:  make([][]int, len(machine.ButtonWiringSchema)),
		JoltageRequirements: append([]int(nil), machine.JoltageRequirements...),
	}
	for i := range machine.ButtonWiringSchema {
		initialMachine.ButtonWiringSchema[i] = append([]int(nil), machine.ButtonWiringSchema[i]...)
	}
	return initialMachine
}

func findFewestPresses2(machine Machine) int {
	type State struct {
		machine Machine
		presses int
	}

	initialMachine := copyMachine(machine)
	queue := []State{{machine: initialMachine, presses: 0}}
	visited := make(map[string]bool)

	stateKey := func(s []int) string {
		return fmt.Sprint(s)
	}

	visited[stateKey(initialMachine.JoltageRequirements)] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		solved := true
		for _, j := range current.machine.JoltageRequirements {
			if j != 0 {
				solved = false
			}
		}

		if solved {
			return current.presses
		}

		for i := len(current.machine.ButtonWiringSchema) - 1; i >= 0; i-- {
			button := current.machine.ButtonWiringSchema[i]
			shouldRemove := false
			for _, idx := range button {
				val := current.machine.JoltageRequirements[idx]
				if val-1 < 0 {
					shouldRemove = true
					break
				}
			}
			if shouldRemove {
				current.machine.ButtonWiringSchema = append(current.machine.ButtonWiringSchema[:i], current.machine.ButtonWiringSchema[i+1:]...)
			}
		}

		for i := len(current.machine.ButtonWiringSchema) - 1; i >= 0; i-- {
			newMachine := copyMachine(current.machine)

			for _, i := range newMachine.ButtonWiringSchema[i] {
				newMachine.JoltageRequirements[i]--
			}

			key := stateKey(newMachine.JoltageRequirements)
			if !visited[key] {
				visited[key] = true
				queue = append(queue, State{machine: newMachine, presses: current.presses + 1})
			}
		}
	}

	return -1 // If no solution found
}

func solvePartTwo(data string) int {
	machines := parse(data)

	sum := 0

	for _, m := range machines {
		sum += findFewestPresses2(m)
	}

	// m := machines[189]
	// sum += findFewestPresses2(m)

	return sum
}
