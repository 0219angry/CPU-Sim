package cpuboard

/* ==============================================================================
 *    Architectural Data Types
 * ============================================================================*/
type Sword int8
type Uword uint8
type Addr uint16
type Bit uint8

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
func Step(cpub *Cpub) int {
	/* Add any code for simulation of single instruction */
	var (
		mar Uword
		ir  Uword
	)

	/*
	 *	Phase 0 : Instruction Fetch
	 */
	mar = cpub.Pc
	cpub.Pc++

	/*
	 *	Phase 1 : Instruction Fetch
	 */
	ir = cpub.Mem[mar]

	/*
	 *	Phase 2 : Instruction decode
	 */
	switch (ir & 0x0f) >> 4 {
	case 0x0: /* NOP or HLT or JAL or JR */

	case 0x1: /* OUT or IN */
	case 0x2: /* RCF or SCF */
	case 0x3: /* Bbc */
	case 0x4: /* Ssm or Rsm */
	case 0x5: /* empty */
	case 0x6: /* LD */
	case 0x7: /* ST */
	case 0x8: /* SBC */
	case 0x9: /* ADC */
	case 0xa: /* SUB */
	case 0xb: /* ADD */
	case 0xc: /* EOR */
	case 0xd: /* OR */
	case 0xe: /* AND */
	case 0xf: /* CMP */

	}

	return RUN_HALT
}
