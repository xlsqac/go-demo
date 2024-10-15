#include "textflag.h"

TEXT ·bytesToString2(SB), NOSPLIT, $0-40
    MOVQ    bs+0(FP), AX
    MOVQ    bs+8(FP), BX
    MOVQ    AX, r0+24(FP)
    MOVQ    BX, r0+32(FP)

    RET

