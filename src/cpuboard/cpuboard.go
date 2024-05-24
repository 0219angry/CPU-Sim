package cpuboard

/* ==============================================================================
 *    Architectural Data Types
 * ============================================================================*/
type Sword int8
type Uword uint8
type Addr uint16
type Bit bool

/* ==============================================================================
 *    CPU Board Resources
 * ============================================================================*/
const MEMORY_SIZE = 256 * 2
const IMEMORY_SIZE = 256

type IOBuf struct {
	Flag Bit
	Buf  Uword
}

type Cpub struct {
	Pc             Uword
	Acc            Uword
	Ix             Uword
	Cf, Vf, Nf, Zf Bit
	Ibuf           *IOBuf
	Obuf           IOBuf
	Ir             Uword
	Mem            [MEMORY_SIZE]Uword
}

/* ==============================================================================
 *    Top Function if an Instruction Simulation
 * ============================================================================*/
const RUN_HALT = 0
const RUN_STEP = 1

/* ==============================================================================
 *    Simulation of Single Instruction
 * ============================================================================*/
func Step(*Cpub) int {
	/* Add any code for simulation of single instruction */

	return RUN_HALT
}
