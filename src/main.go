package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/0219angry/CPU-Sim/command"
	"github.com/0219angry/CPU-Sim/cpuboard"
)

/* ==============================================================================
 *    CPU Board States
 * ============================================================================*/
var cpuboards [2]cpuboard.Cpub

/* ==============================================================================
 *    Initialization of CPU Board
 * ============================================================================*/
func initCpub() int {
	cpuboards[0].Ibuf = &(cpuboards[1].Obuf)
	cpuboards[1].Ibuf = &(cpuboards[0].Obuf)
	return 0
}

/* ==============================================================================
 *    Main Routine
 * ============================================================================*/
func main() {
	var (
		cpub_id int
		cpub    *cpuboard.Cpub
	)
	/*
	 *	Create newscanner
	 */
	scanner := bufio.NewScanner(os.Stdin)

	/*
	 *	Initalize the CPU board state
	 */
	cpub_id = initCpub()
	cpub = &(cpuboards[cpub_id])

	/*
	 *	Interpret commands
	 */

	for {
		/*
		 *	Prompt
		 */
		fmt.Fprintf(os.Stderr, "CPU%d,PC=0x%02x> ", cpub_id, cpub.Pc)

		/*
		 *	Read a command line input
		 */
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Scanner error: %q\n", err)
			os.Exit(0)
		}
		input := strings.Split(scanner.Text(), " ")
		if len(input) <= 0 {
			continue
		}

		/*
		 *	Parse a command
		 */
		if len(input[0]) != 1 {
			command.UnknownCommand()
			continue
		}

		switch input[0] {
		case "i":
			cpuboard.Step(cpub)
		case "c":
			if len(input) == 1 {
				command.Continue(cpub, "")
			} else {
				command.Continue(cpub, input[1])
			}

		case "d":
			command.DisplayRegs(cpub)

		case "s":
			if len(input) == 3 {
				command.SetReg(cpub, input[1], input[2])
			} else {
				command.InvalidInputCount(input[0], len(input)-1)
			}
		case "m":
			switch len(input) {
			case 1:
				command.DisplayMem(cpub, "")
			case 2:
				command.DisplayMem(cpub, input[1])
			default:
				command.InvalidInputCount(input[0], len(input)-1)
			}
		case "w":

		case "t":

		case "h", "?":
			command.Help()

		case "q":

		}

	}
}
