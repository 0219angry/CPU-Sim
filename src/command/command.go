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
	return fmt.Sprintf("0x%02X (%3d, %3d)", value, int8(value), value)
}

/* ==============================================================================
 *    Command: Set CPU Board Register
 * ============================================================================*/
func SetReg(cpub *cpuboard.Cpub, target string, strv string) error {
	uiv, e := parseUword(strv)
	if e != nil {
		return e
	}
	switch target {
	case "pc", "Pc", "PC":
		cpub.Pc = uiv
	case "acc", "Acc", "ACC":
		cpub.Acc = uiv
	case "ix", "Ix", "IX":
		cpub.Ix = uiv
	case "IBUF", "Ibuf", "ibuf":
		cpub.Ibuf.Buf = uiv
		cpub.Ibuf.Flag = 1
	case "OBUF", "Obuf", "obuf":
		cpub.Obuf.Buf = uiv
		cpub.Obuf.Flag = 1
	case "if":
		cpub.Ibuf.Buf = uiv
	case "of":
		cpub.Obuf.Buf = uiv
	default:
		invalidTargetRegName()
	}
	return e
}

func invalidTargetRegName() {
	fmt.Fprintf(os.Stderr, "Unknown register name.\n")
}

/* ==============================================================================
 *    Command: Display CPU Board Memory
 * ============================================================================*/
func DisplayMem(cpub *cpuboard.Cpub, straddr string) error {
	if straddr == "" {
		displayMemAll(cpub)
	} else {
		uiaddr, e := parseAddr(straddr)
		if e != nil {
			return e
		}
		displayMemParts(cpub, int(uiaddr)/16)
	}
	return nil
}

/* ==============================================================================
 *    Command: Display CPU Board Memory (32bytes)
 * ============================================================================*/
func displayMemParts(cpub *cpuboard.Cpub, line int) {
	var area string
	if line < 16 {
		area = "TEXT"
	} else {
		area = "DATA"
	}
	displayColumnName(area)
	for i := 0; i < 2; i++ {
		displayMemLine(cpub, line+i)
	}
}

/* ==============================================================================
 *    Command: Display CPU Board Memory (ALL)
 * ============================================================================*/
func displayMemAll(cpub *cpuboard.Cpub) {
	displayColumnName("TEXT")
	for i := 0; i < 32; i++ {
		if i == 16 {
			displayColumnName("DATA")
		}
		displayMemLine(cpub, i)
	}
}

func displayMemLine(cpub *cpuboard.Cpub, line int) {
	fmt.Fprintf(os.Stderr, "%03X |  ", line*16)
	for i := 0; i < 16; i++ {
		fmt.Fprintf(os.Stderr, "%02X ", cpub.Mem[line*16+i])
	}
	fmt.Fprintf(os.Stderr, "\n")
}

func displayColumnName(area string) {
	fmt.Fprintf(os.Stderr, "%s    0  1  2  3  4  5  6  7  8  9  A  B  C  D  E  F\n", area)
}

/* ==============================================================================
 *    Command: Switch Current CPU Board
 * ============================================================================*/
func SwitchCPU(cpuboards *[2]cpuboard.Cpub, cpubid int) (*cpuboard.Cpub, int) {
	fmt.Fprintf(os.Stderr, "\x1b[31m[DEBUG]\x1b[39m cpu%d acc = %02X\n", cpubid, cpuboards[cpubid].Acc)
	cpubid ^= 1
	return &(cpuboards[cpubid]), cpubid
}

/* ==============================================================================
 *    Command: Quit CPU Board Simulator
 * ============================================================================*/
func Quit() {
	os.Exit(0)
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

/* ==============================================================================
 *    Support Function: String Addr to Unsigned Int (For 9 bits value)
 * ============================================================================*/
func parseAddr(str string) (cpuboard.Addr, error) {
	p, e := strconv.ParseUint(str, 16, 9)
	if enum, ok := e.(*strconv.NumError); ok {
		switch enum.Err {
		case strconv.ErrRange:
			fmt.Fprintf(os.Stderr, "Input address is out of range.")
		case strconv.ErrSyntax:
		}
	}
	return cpuboard.Addr(p), e
}

/* ==============================================================================
 *    Support Function: String Uword to Unsigned Int (For 8 bits value)
 * ============================================================================*/
func parseUword(str string) (cpuboard.Uword, error) {
	p, e := strconv.ParseUint(str, 16, 8)
	if enum, ok := e.(*strconv.NumError); ok {
		switch enum.Err {
		case strconv.ErrRange:
		case strconv.ErrSyntax:
		}
	}
	return cpuboard.Uword(p), e
}
