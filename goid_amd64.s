#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

// see also: https://golang.org/doc/asm

// func GoID() int64
TEXT ·GoID(SB),NOSPLIT,$0-8
	get_tls(CX)
	MOVQ	g(CX), BX
	MOVQ	g_goid(BX), CX
	MOVQ	CX, ret+0(FP)
	RET

