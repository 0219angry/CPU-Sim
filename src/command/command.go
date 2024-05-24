package command

import (
	"fmt"
	"os"
	"strconv"

	"github.com/0219angry/CPU-Sim/cpuboard"
)

func Continue(cpub *cpuboard.Cpub, straddr string) {
	MAX_EXEC_COUNT := 500
	var (
		addr   int
		breakp cpuboard.Addr
		count  int
		temp   int
	)

	/*
	 *	Check and set a break-point address
	 */

	if straddr == "" {
		breakp = 0xffff
	} else {
		p, _ := strconv.ParseUint(straddr, 16, 16)
		breakp = cpuboard.Addr(p)

	}
}

func UnknownCommand() {
	fmt.Fprintf(os.Stderr, "Unknown command. Type 'h' for help.\n")
}
