// go:build amd64 && !gccgo
// +build amd64,!gccgo

#include "go_asm.h"
#include "textflag.h"

DATA encodeShuf<>+0x00(SB)/4, $0xff000102
DATA encodeShuf<>+0x04(SB)/4, $0xff030405
DATA encodeShuf<>+0x08(SB)/4, $0xff060708
DATA encodeShuf<>+0x0c(SB)/4, $0xff090a0b
GLOBL encodeShuf<>(SB), RODATA, $16

DATA encodeShufOut<>+0x00(SB)/4, $0x00010203
DATA encodeShufOut<>+0x04(SB)/4, $0x04050607
DATA encodeShufOut<>+0x08(SB)/4, $0x08090a0b
DATA encodeShufOut<>+0x0c(SB)/4, $0x0c0d0e0f
GLOBL encodeShufOut<>(SB), RODATA, $16

DATA encodeAnd<>+0x00(SB)/8, $0x00000fff00000fff
DATA encodeAnd<>+0x08(SB)/8, $0x00000fff00000fff
DATA encodeAnd<>+0x10(SB)/8, $0x0fff00000fff0000
DATA encodeAnd<>+0x18(SB)/8, $0x0fff00000fff0000
DATA encodeAnd<>+0x20(SB)/8, $0x003f003f003f003f
DATA encodeAnd<>+0x28(SB)/8, $0x003f003f003f003f
DATA encodeAnd<>+0x30(SB)/8, $0x3f003f003f003f00
DATA encodeAnd<>+0x38(SB)/8, $0x3f003f003f003f00
GLOBL encodeAnd<>(SB), RODATA, $64

DATA encodeCompare<>+0x00(SB)/8, $0x1919191919191919
DATA encodeCompare<>+0x08(SB)/8, $0x1919191919191919
DATA encodeCompare<>+0x10(SB)/8, $0x3333333333333333
DATA encodeCompare<>+0x18(SB)/8, $0x3333333333333333
DATA encodeCompare<>+0x20(SB)/8, $0x3e3e3e3e3e3e3e3e
DATA encodeCompare<>+0x28(SB)/8, $0x3e3e3e3e3e3e3e3e
DATA encodeCompare<>+0x30(SB)/8, $0x3f3f3f3f3f3f3f3f
DATA encodeCompare<>+0x38(SB)/8, $0x3f3f3f3f3f3f3f3f
GLOBL encodeCompare<>(SB), RODATA, $64

DATA encodeBase<>+0x00(SB)/8, $0x4141414141414141
DATA encodeBase<>+0x08(SB)/8, $0x4141414141414141
DATA encodeBase<>+0x10(SB)/8, $0x4747474747474747
DATA encodeBase<>+0x18(SB)/8, $0x4747474747474747
DATA encodeBase<>+0x20(SB)/8, $0xfcfcfcfcfcfcfcfc
DATA encodeBase<>+0x28(SB)/8, $0xfcfcfcfcfcfcfcfc
DATA encodeBase<>+0x30(SB)/8, $0x2b2b2b2b2b2b2b2b
DATA encodeBase<>+0x38(SB)/8, $0x2b2b2b2b2b2b2b2b
DATA encodeBase<>+0x40(SB)/8, $0x2f2f2f2f2f2f2f2f
DATA encodeBase<>+0x48(SB)/8, $0x2f2f2f2f2f2f2f2f
DATA encodeBase<>+0x50(SB)/8, $0x2d2d2d2d2d2d2d2d
DATA encodeBase<>+0x58(SB)/8, $0x2d2d2d2d2d2d2d2d
DATA encodeBase<>+0x60(SB)/8, $0x5f5f5f5f5f5f5f5f
DATA encodeBase<>+0x68(SB)/8, $0x5f5f5f5f5f5f5f5f
GLOBL encodeBase<>(SB), RODATA, $112

DATA encodeLookup<>+0x00(SB)/8, $"ABCDEFGH"
DATA encodeLookup<>+0x08(SB)/8, $"IJKLMNOP"
DATA encodeLookup<>+0x10(SB)/8, $"QRSTUVWX"
DATA encodeLookup<>+0x18(SB)/8, $"YZabcdef"
DATA encodeLookup<>+0x20(SB)/8, $"ghijklmn"
DATA encodeLookup<>+0x28(SB)/8, $"opqrstuv"
DATA encodeLookup<>+0x30(SB)/8, $"wxyz0123"
DATA encodeLookup<>+0x38(SB)/8, $"456789+/"
DATA encodeLookup<>+0x40(SB)/8, $"ABCDEFGH"
DATA encodeLookup<>+0x48(SB)/8, $"IJKLMNOP"
DATA encodeLookup<>+0x50(SB)/8, $"QRSTUVWX"
DATA encodeLookup<>+0x58(SB)/8, $"YZabcdef"
DATA encodeLookup<>+0x60(SB)/8, $"ghijklmn"
DATA encodeLookup<>+0x68(SB)/8, $"opqrstuv"
DATA encodeLookup<>+0x70(SB)/8, $"wxyz0123"
DATA encodeLookup<>+0x78(SB)/8, $"456789-_"
GLOBL encodeLookup<>(SB), RODATA, $128

// func encode(dst *byte, src *byte, len uint64, padding int32, url bool)
TEXT ·encode(SB), NOSPLIT|NOFRAME, $0-0
	MOVQ  dst+0(FP), DI
	MOVQ  src+8(FP), SI
	MOVQ  len+16(FP), BX
	MOVL  padding+24(FP), AX
	XORQ  R14, R14
	MOVB  url+28(FP), R14
	SHLQ  $3, R14
	MOVQ  $encodeLookup<>(SB), DX
	LEAQ  (DX)(R14*8), DX
	CMPQ  BX, $3
	JB    tail
	CMPQ  BX, $16
	JB    loop_preheader
	SHRQ  $1, R14
	MOVOU encodeBase<>(SB), X11
	MOVOU encodeBase<>+0x10(SB), X12
	MOVOU encodeBase<>+0x20(SB), X13
	MOVQ  $encodeBase<>(SB), R13
	MOVOU 48(R13)(R14*8), X14
	MOVOU 64(R13)(R14*8), X15
	MOVQ  $encodeCompare<>(SB), R14
	CMPB  golang·org∕x∕sys∕cpu·X86+const_offsetX86HasAVX(SB), $1
	JNE   bigloop_sse

bigloop_avx:
	MOVOU  (SI), X1
	PSHUFB encodeShuf<>(SB), X1
	VPAND  encodeAnd<>(SB), X1, X0
	PSLLL  $4, X1
	PAND   encodeAnd<>+0x10(SB), X1
	POR    X0, X1
	VPAND  encodeAnd<>+0x20(SB), X1, X0
	PSLLL  $2, X1
	PAND   encodeAnd<>+0x30(SB), X1
	POR    X0, X1
	PSHUFB encodeShufOut<>(SB), X1

	// VPCMPGTB (R14), X1, X2
	BYTE $0xc4; BYTE $0xc1; BYTE $0x71; BYTE $0x64; BYTE $0x16

	// VPCMPGTB 16(R14), X1, X3
	BYTE     $0xc4; BYTE $0xc1; BYTE $0x71; BYTE $0x64; BYTE $0x5e; BYTE $0x10
	VPCMPEQB 32(R14), X1, X4
	VPCMPEQB 48(R14), X1, X5

	// VPBLENDVB X2, X12, X11, X2
	BYTE $0xc4; BYTE $0xc3; BYTE $0x21; BYTE $0x4c; BYTE $0xd4; BYTE $0x20

	// VPBLENDVB X3, X13, X2, X2
	BYTE  $0xc4; BYTE $0xc3; BYTE $0x69; BYTE $0x4c; BYTE $0xd5; BYTE $0x30
	PADDB X2, X1

	// VPBLENDVB X4, X14, X1, X1
	BYTE $0xc4; BYTE $0xc3; BYTE $0x71; BYTE $0x4c; BYTE $0xce; BYTE $0x40

	// VPBLENDVB X5, X15, X1, X1
	BYTE  $0xc4; BYTE $0xc3; BYTE $0x71; BYTE $0x4c; BYTE $0xcf; BYTE $0x50
	MOVOU X1, (DI)
	SUBQ  $12, BX
	JZ    ret
	ADDQ  $12, SI
	ADDQ  $16, DI
	CMPQ  BX, $16
	JAE   bigloop_avx
	CMPQ  BX, $3
	JB    tail

loop_preheader:
	XORQ R9, R9
	XORQ R10, R10
	XORQ R11, R11

loop:
	MOVB 2(SI), R9
	MOVB 1(SI), R10
	MOVB (SI), R11
	MOVQ R9, R12
	ANDB $63, R12
	MOVQ R10, R13
	ANDB $15, R13
	SHLB $2, R13
	SHRB $6, R9
	ORB  R9, R13
	MOVQ R11, R14
	ANDB $3, R14
	SHLB $4, R14
	SHRB $4, R10
	ORB  R10, R14
	SHRB $2, R11
	MOVB (DX)(R12*1), R12
	MOVB (DX)(R13*1), R13
	MOVB (DX)(R14*1), R14
	MOVB (DX)(R11*1), R15
	MOVB R12, 3(DI)
	MOVB R13, 2(DI)
	MOVB R14, 1(DI)
	MOVB R15, (DI)
	SUBQ $3, BX
	JZ   ret
	ADDQ $3, SI
	ADDQ $4, DI
	CMPQ BX, $3
	JAE  loop

tail:
	XORQ R11, R11
	MOVB (SI), R11
	MOVQ R11, R14
	ANDB $3, R14
	SHLB $4, R14
	CMPQ BX, $2
	JB   tail_1
	XORQ R10, R10
	MOVB 1(SI), R10
	MOVB R10, R9
	SHRB $4, R9
	ORB  R9, R14
	ANDB $15, R10
	SHLB $2, R10
	MOVB (DX)(R10*1), R13
	MOVB R13, 2(DI)

tail_1:
	SHRB $2, R11
	MOVB (DX)(R14*1), R14
	MOVB (DX)(R11*1), R15
	MOVB R14, 1(DI)
	MOVB R15, (DI)
	CMPL AX, $-1
	JE   ret
	MOVB AX, 1(DI)(BX*1)
	CMPQ BX, $2
	JE   ret
	MOVB AX, 2(DI)(BX*1)

ret:
	RET

bigloop_sse:
	MOVOU   (SI), X1
	PSHUFB  encodeShuf<>(SB), X1
	MOVOU   X1, X0
	PAND    encodeAnd<>(SB), X0
	PSLLL   $4, X1
	PAND    encodeAnd<>+0x10(SB), X1
	POR     X0, X1
	MOVOU   X1, X0
	PAND    encodeAnd<>+0x20(SB), X0
	PSLLL   $2, X1
	PAND    encodeAnd<>+0x30(SB), X1
	POR     X0, X1
	PSHUFB  encodeShufOut<>(SB), X1
	MOVOU   X1, X2
	PCMPGTB (R14), X2
	MOVOU   X1, X3
	PCMPGTB 16(R14), X3
	MOVOU   X1, X4
	PCMPEQB X1, X4
	MOVOU   X1, X5
	PCMPEQB X1, X5
	MOVOU   X2, X0
	MOVOU   X11, X2

	// PBLENDVB X0, X12, X2
	BYTE  $0x66; BYTE $0x41; BYTE $0x0f; BYTE $0x38; BYTE $0x10; BYTE $0xd4
	MOVOU X3, X0

	// PBLENDVB X0, X13, X2
	BYTE  $0x66; BYTE $0x41; BYTE $0x0f; BYTE $0x38; BYTE $0x10; BYTE $0xd5
	PADDB X2, X1
	MOVOU X4, X0

	// PBLENDVB X0, X14, X1
	BYTE  $0x66; BYTE $0x41; BYTE $0x0f; BYTE $0x38; BYTE $0x10; BYTE $0xce
	MOVOU X5, X0

	// PBLENDVB X0, X15, X1
	BYTE  $0x66; BYTE $0x41; BYTE $0x0f; BYTE $0x38; BYTE $0x10; BYTE $0xcf
	MOVOU X1, (DI)
	SUBQ  $12, BX
	JZ    ret
	ADDQ  $12, SI
	ADDQ  $16, DI
	CMPQ  BX, $16
	JAE   bigloop_sse
	CMPQ  BX, $3
	JB    tail
	JMP   loop_preheader
