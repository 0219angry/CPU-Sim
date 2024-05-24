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
 *    Command: Display a Help Menu
 * ============================================================================*/
func help() {
	fmt.Fprint(os.Stderr, "   i\t\t--- execute an instruction (one step execution)\n")
	fmt.Fprint(os.Stderr, "   c [addr]\t--- continue(start) execution [to address(hex)]\n")
	fmt.Fprint(os.Stderr, "   d\t\t--- display the contents of registers\n")
	fmt.Fprint(os.Stderr, "   s reg data\t--- set data(hex) to the register\n\t\t\treg: pc,acc,ix,cf,vf,nf,zf,ibuf,if,obuf,of\n")
	fmt.Fprint(os.Stderr, "   m [addr]\t--- dump memory or display data at memory address(hex)\n")
	fmt.Fprint(os.Stderr, "   w addr data\t--- write data(hex) at memory address(hex)\n")
	fmt.Fprint(os.Stderr, "   r file\t--- load a program into the main memory from the file\n")
	fmt.Fprint(os.Stderr, "   t\t\t--- toggle current computer(context)\n")
	fmt.Fprint(os.Stderr, "   h\t\t--- help (this menu)\n")
	fmt.Fprint(os.Stderr, "   ?\t\t--- help (this menu)\n")
	fmt.Fprint(os.Stderr, "   q\t\t--- quit\n")
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
		}

	}
}
