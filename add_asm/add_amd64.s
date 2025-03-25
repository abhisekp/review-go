#include "textflag.h"

// func Add(a, b int64) int64
TEXT ·Add(SB),NOSPLIT,$0
  MOVQ a+0(FP), AX    // Load first arg (a) into AX
  MOVQ b+8(FP), BX    // Load second arg (b) into BX
  ADDQ BX, AX         // AX = AX + BX
  MOVQ AX, ret+16(FP) // Store result in return slot
  RET
