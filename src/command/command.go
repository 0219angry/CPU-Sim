package command

import (
	"fmt"
	"os"
	"strconv"

	"github.com/0219angry/CPU-Sim/cpuboard"
)

func Continue(cpub *cpuboard.Cpub, straddr string) {
	// MAX_EXEC_COUNT := 500
	// var (
	// 	addr   int
	// 	breakp cpuboard.Addr
	// 	count  int
	// 	temp   int
	// )

	// /*
	//  *	Check and set a break-point address
	//  */

	// if straddr == "" {
	// 	breakp = 0xffff
	// } else {
	// 	p, _ := strconv.ParseUint(straddr, 16, 16)
	// 	breakp = cpuboard.Addr(p)

	// }
}

/* ==============================================================================
 *    Command: Display a Help Menu
 * ============================================================================*/
func Help() {
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
 *    Command: Display CPU Board Registers
 * ============================================================================*/
func DisplayRegs(cpub *cpuboard.Cpub) {
	/*
	 *	Output format:
	 *		acc  =   0xXX (ddd, -ddd)   ix   =   0xXX (ddd, -ddd)
	 *		cf=b   vf=b   nf=b   zf=b
	 *		ibuf = b:0xXX (ddd, -ddd)   obuf = b:0xXX (ddd, -ddd)
	 */
	fmt.Fprintf(os.Stderr, "\t acc  =   %s      ", formatRegVal(cpub.Acc))
	fmt.Fprintf(os.Stderr, "\t ix   =   %s\n", formatRegVal(cpub.Ix))
	fmt.Fprintf(os.Stderr, "\t cf=%d   vf=%d   nf=%d   zf=%d\n", cpub.Cf, cpub.Vf, cpub.Nf, cpub.Zf)
	fmt.Fprintf(os.Stderr, "\t ibuf = %d:%s   ", cpub.Ibuf.Flag, formatRegVal(cpub.Ibuf.Buf))
	fmt.Fprintf(os.Stderr, "\t obuf = %d:%s\n", cpub.Obuf.Flag, formatRegVal(cpub.Obuf.Buf))
}

func formatRegVal(value cpuboard.Uword) string {
	/*
	 *	return string : "Hex value(signed value, unsigned value)"
	 */
	return fmt.Sprintf("0x%02d (%3d, %3d)", value, int8(value), value)
}

/* ==============================================================================
 *    Command: Set CPU Board Register
 * ============================================================================*/
func SetReg(cpub *cpuboard.Cpub, target string, value string) {
	switch target {
	case "acc", "Acc", "ACC":
		p, e := strconv.ParseUint(value, 16, 8)
		if e != nil {

		}
	case "ix", "Ix", "IX":

	default:
		invalidTargetRegName()
	}
}

func invalidTargetRegName() {

}

/* ==============================================================================
 *    Error: Input unknown command message
 * ============================================================================*/
func UnknownCommand() {
	fmt.Fprintf(os.Stderr, "Unknown command. Type 'h' for help.\n")
}

/* ==============================================================================
 *    Error: Invalid number of arguments message
 * ============================================================================*/
func InvalidInputCount(command string, actual int) {
	var expect string /* Expected Number of Arguments */
	switch command {
	case "i", "h", "q", "t", "d", "?":
		expect = "0"
	case "r":
		expect = "1"
	case "c", "m":
		expect = "0 or 1"
	case "s", "w":
		expect = "2"
	}
	fmt.Fprintf(os.Stderr, "Invalid number of arguments. ")
	fmt.Fprintf(os.Stderr, "Expected %s, but actual %d. Type 'h' for help.\n", expect, actual)
}
