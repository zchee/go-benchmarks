// go:build amd64 && !gccgo
// +build amd64,!gccgo

#include "textflag.h"

DATA decodeLowerBound<>+0x00(SB)/8, $0x01012b3040506170
DATA decodeLowerBound<>+0x08(SB)/8, $0x0101010101010101
GLOBL decodeLowerBound<>(SB),RODATA,$16

DATA decodeUpperBound<>+0x00(SB)/8, $0x00002b394f5a6f7a
GLOBL decodeUpperBound<>(SB),RODATA,$16

DATA decodeShifts<>+0x00(SB)/8, $0x00003e34000f1a29
GLOBL decodeShifts<>(SB),RODATA,$16

DATA decodeNibble<>+0x00(SB)/8, $0x0f0f0f0f0f0f0f0f
DATA decodeNibble<>+0x08(SB)/8, $0x0f0f0f0f0f0f0f0f
GLOBL decodeNibble<>(SB),RODATA,$16

DATA decode2f<>+0x00(SB)/8, $0x2f2f2f2f2f2f2f2f
DATA decode2f<>+0x08(SB)/8, $0x2f2f2f2f2f2f2f2f
GLOBL decode2f<>(SB),RODATA,$16

DATA decode2fOffset<>+0x00(SB)/8, $0xfdfdfdfdfdfdfdfd
DATA decode2fOffset<>+0x08(SB)/8, $0xfdfdfdfdfdfdfdfd
GLOBL decode2fOffset<>(SB),RODATA,$16

DATA decodeMerge<>+0x00(SB)/4, $0x40014001
DATA decodeMerge<>+0x04(SB)/4, $0x40014001
DATA decodeMerge<>+0x08(SB)/4, $0x40014001
DATA decodeMerge<>+0x0c(SB)/4, $0x40014001
DATA decodeMerge<>+0x10(SB)/4, $0x10000001
DATA decodeMerge<>+0x14(SB)/4, $0x10000001
DATA decodeMerge<>+0x18(SB)/4, $0x10000001
DATA decodeMerge<>+0x1c(SB)/4, $0x10000001
GLOBL decodeMerge<>(SB),RODATA,$32

TEXT ·decode(SB),NOSPLIT,$0
	MOVQ dst+0(FP), DI
	MOVQ src+8(FP), SI
	MOVQ len+16(FP), BX
	MOVL padding+24(FP), DX
	MOVB url+28(FP), AX
	MOVQ SI, R14
	MOVQ DI, R15
	MOVW $65535, DX
	CMPQ BX, $24
	JB loop
	MOVQ $decodeMerge<>(SB), R15
	MOVQ decodeLowerBound<>(SB), X13
	MOVQ decodeUpperBound<>(SB), X14
	MOVQ decodeShifts<>(SB), X15
bigloop_avx:
	MOVOU (SI), X1
	VPSRLD $4, X1, X2
	PAND decodeNibble<>(SB), X2
	VPSHUFB X2, X13, X3
	VPSHUFB X2, X14, X0
	VPCMPGTB X1, X3, X4
	// BYTE $0xc5; BYTE $0xe1; BYTE $0x64; BYTE $0xe1
	VPCMPGTB X0, X1, X5
	// BYTE $0xc5; BYTE $0xf1; BYTE $0x64; BYTE $0xe8
	VPCMPEQB decode2f<>(SB), X1, X6
	POR X5, X4
	PANDN X4, X6
	PMOVMSKB X6, AX
	VPSHUFB X2, X15, X7
	PSUBB X0, X1
	PADDB X7, X1
	PAND decode2fOffset<>(SB), X6
	PADDB X7, X1
	PMADDUBSW (R15), X1
	// BYTE $0x66; BYTE $0x41; BYTE $0x0f; BYTE $0x38; BYTE $0x04; BYTE $0x0f
	PMADDWL 16(R15), X1
	MOVOU X1, (DI)
	SUBQ $16, BX
	JZ ret
	ADDQ $16, SI
	ADDQ $12, DI
loop:
	XORQ AX, AX
	JMP invalid
ret:
	SUBQ R15, DI
	MOVQ DI, n+32(FP)
	MOVB $1, ok+40(FP)
	RET
invalid:
	BSFW AX, AX
	SUBQ R14, SI
	ADDQ SI, AX
	MOVQ AX, n+32(FP)
	MOVB $0, ok+40(FP)
	RET
