// Code generated from gen/PPC64.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "math"
import "cmd/internal/objabi"
import "cmd/compile/internal/types"

func rewriteValuePPC64(v *Value) bool {
	switch v.Op {
	case OpAbs:
		v.Op = OpPPC64FABS
		return true
	case OpAdd16:
		v.Op = OpPPC64ADD
		return true
	case OpAdd32:
		v.Op = OpPPC64ADD
		return true
	case OpAdd32F:
		v.Op = OpPPC64FADDS
		return true
	case OpAdd64:
		v.Op = OpPPC64ADD
		return true
	case OpAdd64F:
		v.Op = OpPPC64FADD
		return true
	case OpAdd64carry:
		v.Op = OpPPC64LoweredAdd64Carry
		return true
	case OpAdd8:
		v.Op = OpPPC64ADD
		return true
	case OpAddPtr:
		v.Op = OpPPC64ADD
		return true
	case OpAddr:
		v.Op = OpPPC64MOVDaddr
		return true
	case OpAnd16:
		v.Op = OpPPC64AND
		return true
	case OpAnd32:
		v.Op = OpPPC64AND
		return true
	case OpAnd64:
		v.Op = OpPPC64AND
		return true
	case OpAnd8:
		v.Op = OpPPC64AND
		return true
	case OpAndB:
		v.Op = OpPPC64AND
		return true
	case OpAtomicAdd32:
		v.Op = OpPPC64LoweredAtomicAdd32
		return true
	case OpAtomicAdd64:
		v.Op = OpPPC64LoweredAtomicAdd64
		return true
	case OpAtomicAnd8:
		v.Op = OpPPC64LoweredAtomicAnd8
		return true
	case OpAtomicCompareAndSwap32:
		return rewriteValuePPC64_OpAtomicCompareAndSwap32(v)
	case OpAtomicCompareAndSwap64:
		return rewriteValuePPC64_OpAtomicCompareAndSwap64(v)
	case OpAtomicCompareAndSwapRel32:
		return rewriteValuePPC64_OpAtomicCompareAndSwapRel32(v)
	case OpAtomicExchange32:
		v.Op = OpPPC64LoweredAtomicExchange32
		return true
	case OpAtomicExchange64:
		v.Op = OpPPC64LoweredAtomicExchange64
		return true
	case OpAtomicLoad32:
		return rewriteValuePPC64_OpAtomicLoad32(v)
	case OpAtomicLoad64:
		return rewriteValuePPC64_OpAtomicLoad64(v)
	case OpAtomicLoad8:
		return rewriteValuePPC64_OpAtomicLoad8(v)
	case OpAtomicLoadAcq32:
		return rewriteValuePPC64_OpAtomicLoadAcq32(v)
	case OpAtomicLoadPtr:
		return rewriteValuePPC64_OpAtomicLoadPtr(v)
	case OpAtomicOr8:
		v.Op = OpPPC64LoweredAtomicOr8
		return true
	case OpAtomicStore32:
		return rewriteValuePPC64_OpAtomicStore32(v)
	case OpAtomicStore64:
		return rewriteValuePPC64_OpAtomicStore64(v)
	case OpAtomicStore8:
		return rewriteValuePPC64_OpAtomicStore8(v)
	case OpAtomicStoreRel32:
		return rewriteValuePPC64_OpAtomicStoreRel32(v)
	case OpAvg64u:
		return rewriteValuePPC64_OpAvg64u(v)
	case OpBitLen32:
		return rewriteValuePPC64_OpBitLen32(v)
	case OpBitLen64:
		return rewriteValuePPC64_OpBitLen64(v)
	case OpCeil:
		v.Op = OpPPC64FCEIL
		return true
	case OpClosureCall:
		v.Op = OpPPC64CALLclosure
		return true
	case OpCom16:
		return rewriteValuePPC64_OpCom16(v)
	case OpCom32:
		return rewriteValuePPC64_OpCom32(v)
	case OpCom64:
		return rewriteValuePPC64_OpCom64(v)
	case OpCom8:
		return rewriteValuePPC64_OpCom8(v)
	case OpCondSelect:
		return rewriteValuePPC64_OpCondSelect(v)
	case OpConst16:
		v.Op = OpPPC64MOVDconst
		return true
	case OpConst32:
		v.Op = OpPPC64MOVDconst
		return true
	case OpConst32F:
		v.Op = OpPPC64FMOVSconst
		return true
	case OpConst64:
		v.Op = OpPPC64MOVDconst
		return true
	case OpConst64F:
		v.Op = OpPPC64FMOVDconst
		return true
	case OpConst8:
		v.Op = OpPPC64MOVDconst
		return true
	case OpConstBool:
		v.Op = OpPPC64MOVDconst
		return true
	case OpConstNil:
		return rewriteValuePPC64_OpConstNil(v)
	case OpCopysign:
		return rewriteValuePPC64_OpCopysign(v)
	case OpCtz16:
		return rewriteValuePPC64_OpCtz16(v)
	case OpCtz32:
		return rewriteValuePPC64_OpCtz32(v)
	case OpCtz32NonZero:
		v.Op = OpCtz32
		return true
	case OpCtz64:
		return rewriteValuePPC64_OpCtz64(v)
	case OpCtz64NonZero:
		v.Op = OpCtz64
		return true
	case OpCtz8:
		return rewriteValuePPC64_OpCtz8(v)
	case OpCvt32Fto32:
		return rewriteValuePPC64_OpCvt32Fto32(v)
	case OpCvt32Fto64:
		return rewriteValuePPC64_OpCvt32Fto64(v)
	case OpCvt32Fto64F:
		v.Op = OpCopy
		return true
	case OpCvt32to32F:
		return rewriteValuePPC64_OpCvt32to32F(v)
	case OpCvt32to64F:
		return rewriteValuePPC64_OpCvt32to64F(v)
	case OpCvt64Fto32:
		return rewriteValuePPC64_OpCvt64Fto32(v)
	case OpCvt64Fto32F:
		v.Op = OpPPC64FRSP
		return true
	case OpCvt64Fto64:
		return rewriteValuePPC64_OpCvt64Fto64(v)
	case OpCvt64to32F:
		return rewriteValuePPC64_OpCvt64to32F(v)
	case OpCvt64to64F:
		return rewriteValuePPC64_OpCvt64to64F(v)
	case OpDiv16:
		return rewriteValuePPC64_OpDiv16(v)
	case OpDiv16u:
		return rewriteValuePPC64_OpDiv16u(v)
	case OpDiv32:
		return rewriteValuePPC64_OpDiv32(v)
	case OpDiv32F:
		v.Op = OpPPC64FDIVS
		return true
	case OpDiv32u:
		v.Op = OpPPC64DIVWU
		return true
	case OpDiv64:
		return rewriteValuePPC64_OpDiv64(v)
	case OpDiv64F:
		v.Op = OpPPC64FDIV
		return true
	case OpDiv64u:
		v.Op = OpPPC64DIVDU
		return true
	case OpDiv8:
		return rewriteValuePPC64_OpDiv8(v)
	case OpDiv8u:
		return rewriteValuePPC64_OpDiv8u(v)
	case OpEq16:
		return rewriteValuePPC64_OpEq16(v)
	case OpEq32:
		return rewriteValuePPC64_OpEq32(v)
	case OpEq32F:
		return rewriteValuePPC64_OpEq32F(v)
	case OpEq64:
		return rewriteValuePPC64_OpEq64(v)
	case OpEq64F:
		return rewriteValuePPC64_OpEq64F(v)
	case OpEq8:
		return rewriteValuePPC64_OpEq8(v)
	case OpEqB:
		return rewriteValuePPC64_OpEqB(v)
	case OpEqPtr:
		return rewriteValuePPC64_OpEqPtr(v)
	case OpFMA:
		v.Op = OpPPC64FMADD
		return true
	case OpFloor:
		v.Op = OpPPC64FFLOOR
		return true
	case OpGeq32F:
		return rewriteValuePPC64_OpGeq32F(v)
	case OpGeq64F:
		return rewriteValuePPC64_OpGeq64F(v)
	case OpGetCallerPC:
		v.Op = OpPPC64LoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpPPC64LoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpPPC64LoweredGetClosurePtr
		return true
	case OpGreater32F:
		return rewriteValuePPC64_OpGreater32F(v)
	case OpGreater64F:
		return rewriteValuePPC64_OpGreater64F(v)
	case OpHmul32:
		v.Op = OpPPC64MULHW
		return true
	case OpHmul32u:
		v.Op = OpPPC64MULHWU
		return true
	case OpHmul64:
		v.Op = OpPPC64MULHD
		return true
	case OpHmul64u:
		v.Op = OpPPC64MULHDU
		return true
	case OpInterCall:
		v.Op = OpPPC64CALLinter
		return true
	case OpIsInBounds:
		return rewriteValuePPC64_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValuePPC64_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValuePPC64_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValuePPC64_OpLeq16(v)
	case OpLeq16U:
		return rewriteValuePPC64_OpLeq16U(v)
	case OpLeq32:
		return rewriteValuePPC64_OpLeq32(v)
	case OpLeq32F:
		return rewriteValuePPC64_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValuePPC64_OpLeq32U(v)
	case OpLeq64:
		return rewriteValuePPC64_OpLeq64(v)
	case OpLeq64F:
		return rewriteValuePPC64_OpLeq64F(v)
	case OpLeq64U:
		return rewriteValuePPC64_OpLeq64U(v)
	case OpLeq8:
		return rewriteValuePPC64_OpLeq8(v)
	case OpLeq8U:
		return rewriteValuePPC64_OpLeq8U(v)
	case OpLess16:
		return rewriteValuePPC64_OpLess16(v)
	case OpLess16U:
		return rewriteValuePPC64_OpLess16U(v)
	case OpLess32:
		return rewriteValuePPC64_OpLess32(v)
	case OpLess32F:
		return rewriteValuePPC64_OpLess32F(v)
	case OpLess32U:
		return rewriteValuePPC64_OpLess32U(v)
	case OpLess64:
		return rewriteValuePPC64_OpLess64(v)
	case OpLess64F:
		return rewriteValuePPC64_OpLess64F(v)
	case OpLess64U:
		return rewriteValuePPC64_OpLess64U(v)
	case OpLess8:
		return rewriteValuePPC64_OpLess8(v)
	case OpLess8U:
		return rewriteValuePPC64_OpLess8U(v)
	case OpLoad:
		return rewriteValuePPC64_OpLoad(v)
	case OpLocalAddr:
		return rewriteValuePPC64_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValuePPC64_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValuePPC64_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValuePPC64_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValuePPC64_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValuePPC64_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValuePPC64_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValuePPC64_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValuePPC64_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValuePPC64_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValuePPC64_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValuePPC64_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValuePPC64_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValuePPC64_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValuePPC64_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValuePPC64_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValuePPC64_OpLsh8x8(v)
	case OpMod16:
		return rewriteValuePPC64_OpMod16(v)
	case OpMod16u:
		return rewriteValuePPC64_OpMod16u(v)
	case OpMod32:
		return rewriteValuePPC64_OpMod32(v)
	case OpMod32u:
		return rewriteValuePPC64_OpMod32u(v)
	case OpMod64:
		return rewriteValuePPC64_OpMod64(v)
	case OpMod64u:
		return rewriteValuePPC64_OpMod64u(v)
	case OpMod8:
		return rewriteValuePPC64_OpMod8(v)
	case OpMod8u:
		return rewriteValuePPC64_OpMod8u(v)
	case OpMove:
		return rewriteValuePPC64_OpMove(v)
	case OpMul16:
		v.Op = OpPPC64MULLW
		return true
	case OpMul32:
		v.Op = OpPPC64MULLW
		return true
	case OpMul32F:
		v.Op = OpPPC64FMULS
		return true
	case OpMul64:
		v.Op = OpPPC64MULLD
		return true
	case OpMul64F:
		v.Op = OpPPC64FMUL
		return true
	case OpMul64uhilo:
		v.Op = OpPPC64LoweredMuluhilo
		return true
	case OpMul8:
		v.Op = OpPPC64MULLW
		return true
	case OpNeg16:
		v.Op = OpPPC64NEG
		return true
	case OpNeg32:
		v.Op = OpPPC64NEG
		return true
	case OpNeg32F:
		v.Op = OpPPC64FNEG
		return true
	case OpNeg64:
		v.Op = OpPPC64NEG
		return true
	case OpNeg64F:
		v.Op = OpPPC64FNEG
		return true
	case OpNeg8:
		v.Op = OpPPC64NEG
		return true
	case OpNeq16:
		return rewriteValuePPC64_OpNeq16(v)
	case OpNeq32:
		return rewriteValuePPC64_OpNeq32(v)
	case OpNeq32F:
		return rewriteValuePPC64_OpNeq32F(v)
	case OpNeq64:
		return rewriteValuePPC64_OpNeq64(v)
	case OpNeq64F:
		return rewriteValuePPC64_OpNeq64F(v)
	case OpNeq8:
		return rewriteValuePPC64_OpNeq8(v)
	case OpNeqB:
		v.Op = OpPPC64XOR
		return true
	case OpNeqPtr:
		return rewriteValuePPC64_OpNeqPtr(v)
	case OpNilCheck:
		v.Op = OpPPC64LoweredNilCheck
		return true
	case OpNot:
		return rewriteValuePPC64_OpNot(v)
	case OpOffPtr:
		return rewriteValuePPC64_OpOffPtr(v)
	case OpOr16:
		v.Op = OpPPC64OR
		return true
	case OpOr32:
		v.Op = OpPPC64OR
		return true
	case OpOr64:
		v.Op = OpPPC64OR
		return true
	case OpOr8:
		v.Op = OpPPC64OR
		return true
	case OpOrB:
		v.Op = OpPPC64OR
		return true
	case OpPPC64ADD:
		return rewriteValuePPC64_OpPPC64ADD(v)
	case OpPPC64ADDconst:
		return rewriteValuePPC64_OpPPC64ADDconst(v)
	case OpPPC64AND:
		return rewriteValuePPC64_OpPPC64AND(v)
	case OpPPC64ANDconst:
		return rewriteValuePPC64_OpPPC64ANDconst(v)
	case OpPPC64CMP:
		return rewriteValuePPC64_OpPPC64CMP(v)
	case OpPPC64CMPU:
		return rewriteValuePPC64_OpPPC64CMPU(v)
	case OpPPC64CMPUconst:
		return rewriteValuePPC64_OpPPC64CMPUconst(v)
	case OpPPC64CMPW:
		return rewriteValuePPC64_OpPPC64CMPW(v)
	case OpPPC64CMPWU:
		return rewriteValuePPC64_OpPPC64CMPWU(v)
	case OpPPC64CMPWUconst:
		return rewriteValuePPC64_OpPPC64CMPWUconst(v)
	case OpPPC64CMPWconst:
		return rewriteValuePPC64_OpPPC64CMPWconst(v)
	case OpPPC64CMPconst:
		return rewriteValuePPC64_OpPPC64CMPconst(v)
	case OpPPC64Equal:
		return rewriteValuePPC64_OpPPC64Equal(v)
	case OpPPC64FABS:
		return rewriteValuePPC64_OpPPC64FABS(v)
	case OpPPC64FADD:
		return rewriteValuePPC64_OpPPC64FADD(v)
	case OpPPC64FADDS:
		return rewriteValuePPC64_OpPPC64FADDS(v)
	case OpPPC64FCEIL:
		return rewriteValuePPC64_OpPPC64FCEIL(v)
	case OpPPC64FFLOOR:
		return rewriteValuePPC64_OpPPC64FFLOOR(v)
	case OpPPC64FGreaterEqual:
		return rewriteValuePPC64_OpPPC64FGreaterEqual(v)
	case OpPPC64FGreaterThan:
		return rewriteValuePPC64_OpPPC64FGreaterThan(v)
	case OpPPC64FLessEqual:
		return rewriteValuePPC64_OpPPC64FLessEqual(v)
	case OpPPC64FLessThan:
		return rewriteValuePPC64_OpPPC64FLessThan(v)
	case OpPPC64FMOVDload:
		return rewriteValuePPC64_OpPPC64FMOVDload(v)
	case OpPPC64FMOVDstore:
		return rewriteValuePPC64_OpPPC64FMOVDstore(v)
	case OpPPC64FMOVSload:
		return rewriteValuePPC64_OpPPC64FMOVSload(v)
	case OpPPC64FMOVSstore:
		return rewriteValuePPC64_OpPPC64FMOVSstore(v)
	case OpPPC64FNEG:
		return rewriteValuePPC64_OpPPC64FNEG(v)
	case OpPPC64FSQRT:
		return rewriteValuePPC64_OpPPC64FSQRT(v)
	case OpPPC64FSUB:
		return rewriteValuePPC64_OpPPC64FSUB(v)
	case OpPPC64FSUBS:
		return rewriteValuePPC64_OpPPC64FSUBS(v)
	case OpPPC64FTRUNC:
		return rewriteValuePPC64_OpPPC64FTRUNC(v)
	case OpPPC64GreaterEqual:
		return rewriteValuePPC64_OpPPC64GreaterEqual(v)
	case OpPPC64GreaterThan:
		return rewriteValuePPC64_OpPPC64GreaterThan(v)
	case OpPPC64ISEL:
		return rewriteValuePPC64_OpPPC64ISEL(v)
	case OpPPC64ISELB:
		return rewriteValuePPC64_OpPPC64ISELB(v)
	case OpPPC64LessEqual:
		return rewriteValuePPC64_OpPPC64LessEqual(v)
	case OpPPC64LessThan:
		return rewriteValuePPC64_OpPPC64LessThan(v)
	case OpPPC64MFVSRD:
		return rewriteValuePPC64_OpPPC64MFVSRD(v)
	case OpPPC64MOVBZload:
		return rewriteValuePPC64_OpPPC64MOVBZload(v)
	case OpPPC64MOVBZloadidx:
		return rewriteValuePPC64_OpPPC64MOVBZloadidx(v)
	case OpPPC64MOVBZreg:
		return rewriteValuePPC64_OpPPC64MOVBZreg(v)
	case OpPPC64MOVBreg:
		return rewriteValuePPC64_OpPPC64MOVBreg(v)
	case OpPPC64MOVBstore:
		return rewriteValuePPC64_OpPPC64MOVBstore(v)
	case OpPPC64MOVBstoreidx:
		return rewriteValuePPC64_OpPPC64MOVBstoreidx(v)
	case OpPPC64MOVBstorezero:
		return rewriteValuePPC64_OpPPC64MOVBstorezero(v)
	case OpPPC64MOVDload:
		return rewriteValuePPC64_OpPPC64MOVDload(v)
	case OpPPC64MOVDloadidx:
		return rewriteValuePPC64_OpPPC64MOVDloadidx(v)
	case OpPPC64MOVDstore:
		return rewriteValuePPC64_OpPPC64MOVDstore(v)
	case OpPPC64MOVDstoreidx:
		return rewriteValuePPC64_OpPPC64MOVDstoreidx(v)
	case OpPPC64MOVDstorezero:
		return rewriteValuePPC64_OpPPC64MOVDstorezero(v)
	case OpPPC64MOVHBRstore:
		return rewriteValuePPC64_OpPPC64MOVHBRstore(v)
	case OpPPC64MOVHZload:
		return rewriteValuePPC64_OpPPC64MOVHZload(v)
	case OpPPC64MOVHZloadidx:
		return rewriteValuePPC64_OpPPC64MOVHZloadidx(v)
	case OpPPC64MOVHZreg:
		return rewriteValuePPC64_OpPPC64MOVHZreg(v)
	case OpPPC64MOVHload:
		return rewriteValuePPC64_OpPPC64MOVHload(v)
	case OpPPC64MOVHloadidx:
		return rewriteValuePPC64_OpPPC64MOVHloadidx(v)
	case OpPPC64MOVHreg:
		return rewriteValuePPC64_OpPPC64MOVHreg(v)
	case OpPPC64MOVHstore:
		return rewriteValuePPC64_OpPPC64MOVHstore(v)
	case OpPPC64MOVHstoreidx:
		return rewriteValuePPC64_OpPPC64MOVHstoreidx(v)
	case OpPPC64MOVHstorezero:
		return rewriteValuePPC64_OpPPC64MOVHstorezero(v)
	case OpPPC64MOVWBRstore:
		return rewriteValuePPC64_OpPPC64MOVWBRstore(v)
	case OpPPC64MOVWZload:
		return rewriteValuePPC64_OpPPC64MOVWZload(v)
	case OpPPC64MOVWZloadidx:
		return rewriteValuePPC64_OpPPC64MOVWZloadidx(v)
	case OpPPC64MOVWZreg:
		return rewriteValuePPC64_OpPPC64MOVWZreg(v)
	case OpPPC64MOVWload:
		return rewriteValuePPC64_OpPPC64MOVWload(v)
	case OpPPC64MOVWloadidx:
		return rewriteValuePPC64_OpPPC64MOVWloadidx(v)
	case OpPPC64MOVWreg:
		return rewriteValuePPC64_OpPPC64MOVWreg(v)
	case OpPPC64MOVWstore:
		return rewriteValuePPC64_OpPPC64MOVWstore(v)
	case OpPPC64MOVWstoreidx:
		return rewriteValuePPC64_OpPPC64MOVWstoreidx(v)
	case OpPPC64MOVWstorezero:
		return rewriteValuePPC64_OpPPC64MOVWstorezero(v)
	case OpPPC64MTVSRD:
		return rewriteValuePPC64_OpPPC64MTVSRD(v)
	case OpPPC64MaskIfNotCarry:
		return rewriteValuePPC64_OpPPC64MaskIfNotCarry(v)
	case OpPPC64NotEqual:
		return rewriteValuePPC64_OpPPC64NotEqual(v)
	case OpPPC64OR:
		return rewriteValuePPC64_OpPPC64OR(v)
	case OpPPC64ORN:
		return rewriteValuePPC64_OpPPC64ORN(v)
	case OpPPC64ORconst:
		return rewriteValuePPC64_OpPPC64ORconst(v)
	case OpPPC64ROTL:
		return rewriteValuePPC64_OpPPC64ROTL(v)
	case OpPPC64ROTLW:
		return rewriteValuePPC64_OpPPC64ROTLW(v)
	case OpPPC64SUB:
		return rewriteValuePPC64_OpPPC64SUB(v)
	case OpPPC64XOR:
		return rewriteValuePPC64_OpPPC64XOR(v)
	case OpPPC64XORconst:
		return rewriteValuePPC64_OpPPC64XORconst(v)
	case OpPanicBounds:
		return rewriteValuePPC64_OpPanicBounds(v)
	case OpPopCount16:
		return rewriteValuePPC64_OpPopCount16(v)
	case OpPopCount32:
		return rewriteValuePPC64_OpPopCount32(v)
	case OpPopCount64:
		v.Op = OpPPC64POPCNTD
		return true
	case OpPopCount8:
		return rewriteValuePPC64_OpPopCount8(v)
	case OpRotateLeft16:
		return rewriteValuePPC64_OpRotateLeft16(v)
	case OpRotateLeft32:
		return rewriteValuePPC64_OpRotateLeft32(v)
	case OpRotateLeft64:
		return rewriteValuePPC64_OpRotateLeft64(v)
	case OpRotateLeft8:
		return rewriteValuePPC64_OpRotateLeft8(v)
	case OpRound:
		v.Op = OpPPC64FROUND
		return true
	case OpRound32F:
		v.Op = OpPPC64LoweredRound32F
		return true
	case OpRound64F:
		v.Op = OpPPC64LoweredRound64F
		return true
	case OpRsh16Ux16:
		return rewriteValuePPC64_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValuePPC64_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValuePPC64_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValuePPC64_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValuePPC64_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValuePPC64_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValuePPC64_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValuePPC64_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValuePPC64_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValuePPC64_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValuePPC64_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValuePPC64_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValuePPC64_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValuePPC64_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValuePPC64_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValuePPC64_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValuePPC64_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValuePPC64_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValuePPC64_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValuePPC64_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValuePPC64_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValuePPC64_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValuePPC64_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValuePPC64_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValuePPC64_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValuePPC64_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValuePPC64_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValuePPC64_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValuePPC64_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValuePPC64_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValuePPC64_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValuePPC64_OpRsh8x8(v)
	case OpSignExt16to32:
		v.Op = OpPPC64MOVHreg
		return true
	case OpSignExt16to64:
		v.Op = OpPPC64MOVHreg
		return true
	case OpSignExt32to64:
		v.Op = OpPPC64MOVWreg
		return true
	case OpSignExt8to16:
		v.Op = OpPPC64MOVBreg
		return true
	case OpSignExt8to32:
		v.Op = OpPPC64MOVBreg
		return true
	case OpSignExt8to64:
		v.Op = OpPPC64MOVBreg
		return true
	case OpSlicemask:
		return rewriteValuePPC64_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpPPC64FSQRT
		return true
	case OpStaticCall:
		v.Op = OpPPC64CALLstatic
		return true
	case OpStore:
		return rewriteValuePPC64_OpStore(v)
	case OpSub16:
		v.Op = OpPPC64SUB
		return true
	case OpSub32:
		v.Op = OpPPC64SUB
		return true
	case OpSub32F:
		v.Op = OpPPC64FSUBS
		return true
	case OpSub64:
		v.Op = OpPPC64SUB
		return true
	case OpSub64F:
		v.Op = OpPPC64FSUB
		return true
	case OpSub8:
		v.Op = OpPPC64SUB
		return true
	case OpSubPtr:
		v.Op = OpPPC64SUB
		return true
	case OpTrunc:
		v.Op = OpPPC64FTRUNC
		return true
	case OpTrunc16to8:
		return rewriteValuePPC64_OpTrunc16to8(v)
	case OpTrunc32to16:
		return rewriteValuePPC64_OpTrunc32to16(v)
	case OpTrunc32to8:
		return rewriteValuePPC64_OpTrunc32to8(v)
	case OpTrunc64to16:
		return rewriteValuePPC64_OpTrunc64to16(v)
	case OpTrunc64to32:
		return rewriteValuePPC64_OpTrunc64to32(v)
	case OpTrunc64to8:
		return rewriteValuePPC64_OpTrunc64to8(v)
	case OpWB:
		v.Op = OpPPC64LoweredWB
		return true
	case OpXor16:
		v.Op = OpPPC64XOR
		return true
	case OpXor32:
		v.Op = OpPPC64XOR
		return true
	case OpXor64:
		v.Op = OpPPC64XOR
		return true
	case OpXor8:
		v.Op = OpPPC64XOR
		return true
	case OpZero:
		return rewriteValuePPC64_OpZero(v)
	case OpZeroExt16to32:
		v.Op = OpPPC64MOVHZreg
		return true
	case OpZeroExt16to64:
		v.Op = OpPPC64MOVHZreg
		return true
	case OpZeroExt32to64:
		v.Op = OpPPC64MOVWZreg
		return true
	case OpZeroExt8to16:
		v.Op = OpPPC64MOVBZreg
		return true
	case OpZeroExt8to32:
		v.Op = OpPPC64MOVBZreg
		return true
	case OpZeroExt8to64:
		v.Op = OpPPC64MOVBZreg
		return true
	}
	return false
}
func rewriteValuePPC64_OpAtomicCompareAndSwap32(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicCompareAndSwap32 ptr old new_ mem)
	// result: (LoweredAtomicCas32 [1] ptr old new_ mem)
	for {
		ptr := v_0
		old := v_1
		new_ := v_2
		mem := v_3
		v.reset(OpPPC64LoweredAtomicCas32)
		v.AuxInt = 1
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicCompareAndSwap64(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicCompareAndSwap64 ptr old new_ mem)
	// result: (LoweredAtomicCas64 [1] ptr old new_ mem)
	for {
		ptr := v_0
		old := v_1
		new_ := v_2
		mem := v_3
		v.reset(OpPPC64LoweredAtomicCas64)
		v.AuxInt = 1
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicCompareAndSwapRel32(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicCompareAndSwapRel32 ptr old new_ mem)
	// result: (LoweredAtomicCas32 [0] ptr old new_ mem)
	for {
		ptr := v_0
		old := v_1
		new_ := v_2
		mem := v_3
		v.reset(OpPPC64LoweredAtomicCas32)
		v.AuxInt = 0
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoad32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicLoad32 ptr mem)
	// result: (LoweredAtomicLoad32 [1] ptr mem)
	for {
		ptr := v_0
		mem := v_1
		v.reset(OpPPC64LoweredAtomicLoad32)
		v.AuxInt = 1
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoad64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicLoad64 ptr mem)
	// result: (LoweredAtomicLoad64 [1] ptr mem)
	for {
		ptr := v_0
		mem := v_1
		v.reset(OpPPC64LoweredAtomicLoad64)
		v.AuxInt = 1
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoad8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicLoad8 ptr mem)
	// result: (LoweredAtomicLoad8 [1] ptr mem)
	for {
		ptr := v_0
		mem := v_1
		v.reset(OpPPC64LoweredAtomicLoad8)
		v.AuxInt = 1
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoadAcq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicLoadAcq32 ptr mem)
	// result: (LoweredAtomicLoad32 [0] ptr mem)
	for {
		ptr := v_0
		mem := v_1
		v.reset(OpPPC64LoweredAtomicLoad32)
		v.AuxInt = 0
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoadPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicLoadPtr ptr mem)
	// result: (LoweredAtomicLoadPtr [1] ptr mem)
	for {
		ptr := v_0
		mem := v_1
		v.reset(OpPPC64LoweredAtomicLoadPtr)
		v.AuxInt = 1
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicStore32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicStore32 ptr val mem)
	// result: (LoweredAtomicStore32 [1] ptr val mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpPPC64LoweredAtomicStore32)
		v.AuxInt = 1
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicStore64(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicStore64 ptr val mem)
	// result: (LoweredAtomicStore64 [1] ptr val mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpPPC64LoweredAtomicStore64)
		v.AuxInt = 1
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicStore8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicStore8 ptr val mem)
	// result: (LoweredAtomicStore8 [1] ptr val mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpPPC64LoweredAtomicStore8)
		v.AuxInt = 1
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicStoreRel32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AtomicStoreRel32 ptr val mem)
	// result: (LoweredAtomicStore32 [0] ptr val mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpPPC64LoweredAtomicStore32)
		v.AuxInt = 0
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAvg64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg64u <t> x y)
	// result: (ADD (SRDconst <t> (SUB <t> x y) [1]) y)
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpPPC64ADD)
		v0 := b.NewValue0(v.Pos, OpPPC64SRDconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpPPC64SUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpBitLen32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen32 x)
	// result: (SUB (MOVDconst [32]) (CNTLZW <typ.Int> x))
	for {
		x := v_0
		v.reset(OpPPC64SUB)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 32
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64CNTLZW, typ.Int)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpBitLen64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen64 x)
	// result: (SUB (MOVDconst [64]) (CNTLZD <typ.Int> x))
	for {
		x := v_0
		v.reset(OpPPC64SUB)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 64
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64CNTLZD, typ.Int)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpCom16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com16 x)
	// result: (NOR x x)
	for {
		x := v_0
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCom32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com32 x)
	// result: (NOR x x)
	for {
		x := v_0
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCom64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com64 x)
	// result: (NOR x x)
	for {
		x := v_0
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCom8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com8 x)
	// result: (NOR x x)
	for {
		x := v_0
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCondSelect(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CondSelect x y bool)
	// cond: flagArg(bool) != nil
	// result: (ISEL [2] x y bool)
	for {
		x := v_0
		y := v_1
		bool := v_2
		if !(flagArg(bool) != nil) {
			break
		}
		v.reset(OpPPC64ISEL)
		v.AuxInt = 2
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(bool)
		return true
	}
	// match: (CondSelect x y bool)
	// cond: flagArg(bool) == nil
	// result: (ISEL [2] x y (CMPWconst [0] bool))
	for {
		x := v_0
		y := v_1
		bool := v_2
		if !(flagArg(bool) == nil) {
			break
		}
		v.reset(OpPPC64ISEL)
		v.AuxInt = 2
		v.AddArg(x)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(bool)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVDconst [0])
	for {
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValuePPC64_OpCopysign(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Copysign x y)
	// result: (FCPSGN y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64FCPSGN)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCtz16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz16 x)
	// result: (POPCNTW (MOVHZreg (ANDN <typ.Int16> (ADDconst <typ.Int16> [-1] x) x)))
	for {
		x := v_0
		v.reset(OpPPC64POPCNTW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZreg, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpPPC64ANDN, typ.Int16)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconst, typ.Int16)
		v2.AuxInt = -1
		v2.AddArg(x)
		v1.AddArg(v2)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCtz32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz32 x)
	// cond: objabi.GOPPC64<=8
	// result: (POPCNTW (MOVWZreg (ANDN <typ.Int> (ADDconst <typ.Int> [-1] x) x)))
	for {
		x := v_0
		if !(objabi.GOPPC64 <= 8) {
			break
		}
		v.reset(OpPPC64POPCNTW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZreg, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpPPC64ANDN, typ.Int)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconst, typ.Int)
		v2.AuxInt = -1
		v2.AddArg(x)
		v1.AddArg(v2)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Ctz32 x)
	// result: (CNTTZW (MOVWZreg x))
	for {
		x := v_0
		v.reset(OpPPC64CNTTZW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCtz64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz64 x)
	// cond: objabi.GOPPC64<=8
	// result: (POPCNTD (ANDN <typ.Int64> (ADDconst <typ.Int64> [-1] x) x))
	for {
		x := v_0
		if !(objabi.GOPPC64 <= 8) {
			break
		}
		v.reset(OpPPC64POPCNTD)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDN, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpPPC64ADDconst, typ.Int64)
		v1.AuxInt = -1
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Ctz64 x)
	// result: (CNTTZD x)
	for {
		x := v_0
		v.reset(OpPPC64CNTTZD)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCtz8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz8 x)
	// result: (POPCNTB (MOVBZreg (ANDN <typ.UInt8> (ADDconst <typ.UInt8> [-1] x) x)))
	for {
		x := v_0
		v.reset(OpPPC64POPCNTB)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZreg, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpPPC64ANDN, typ.UInt8)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconst, typ.UInt8)
		v2.AuxInt = -1
		v2.AddArg(x)
		v1.AddArg(v2)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt32Fto32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32Fto32 x)
	// result: (MFVSRD (FCTIWZ x))
	for {
		x := v_0
		v.reset(OpPPC64MFVSRD)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIWZ, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt32Fto64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32Fto64 x)
	// result: (MFVSRD (FCTIDZ x))
	for {
		x := v_0
		v.reset(OpPPC64MFVSRD)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIDZ, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt32to32F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to32F x)
	// result: (FCFIDS (MTVSRD (SignExt32to64 x)))
	for {
		x := v_0
		v.reset(OpPPC64FCFIDS)
		v0 := b.NewValue0(v.Pos, OpPPC64MTVSRD, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt32to64F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to64F x)
	// result: (FCFID (MTVSRD (SignExt32to64 x)))
	for {
		x := v_0
		v.reset(OpPPC64FCFID)
		v0 := b.NewValue0(v.Pos, OpPPC64MTVSRD, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64Fto32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt64Fto32 x)
	// result: (MFVSRD (FCTIWZ x))
	for {
		x := v_0
		v.reset(OpPPC64MFVSRD)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIWZ, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64Fto64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt64Fto64 x)
	// result: (MFVSRD (FCTIDZ x))
	for {
		x := v_0
		v.reset(OpPPC64MFVSRD)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIDZ, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64to32F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt64to32F x)
	// result: (FCFIDS (MTVSRD x))
	for {
		x := v_0
		v.reset(OpPPC64FCFIDS)
		v0 := b.NewValue0(v.Pos, OpPPC64MTVSRD, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64to64F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt64to64F x)
	// result: (FCFID (MTVSRD x))
	for {
		x := v_0
		v.reset(OpPPC64FCFID)
		v0 := b.NewValue0(v.Pos, OpPPC64MTVSRD, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// result: (DIVW (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (DIVWU (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64DIVWU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div32 [a] x y)
	// result: (DIVW x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64DIVW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div64 [a] x y)
	// result: (DIVD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64DIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (DIVW (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (DIVWU (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64DIVWU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// cond: isSigned(x.Type) && isSigned(y.Type)
	// result: (Equal (CMPW (SignExt16to32 x) (SignExt16to32 y)))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			y := v_1
			if !(isSigned(x.Type) && isSigned(y.Type)) {
				continue
			}
			v.reset(OpPPC64Equal)
			v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
			v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
			v1.AddArg(x)
			v0.AddArg(v1)
			v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
			v2.AddArg(y)
			v0.AddArg(v2)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Eq16 x y)
	// result: (Equal (CMPW (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32 x y)
	// result: (Equal (CMPW x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32F x y)
	// result: (Equal (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64 x y)
	// result: (Equal (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64F x y)
	// result: (Equal (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// cond: isSigned(x.Type) && isSigned(y.Type)
	// result: (Equal (CMPW (SignExt8to32 x) (SignExt8to32 y)))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			y := v_1
			if !(isSigned(x.Type) && isSigned(y.Type)) {
				continue
			}
			v.reset(OpPPC64Equal)
			v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
			v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
			v1.AddArg(x)
			v0.AddArg(v1)
			v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
			v2.AddArg(y)
			v0.AddArg(v2)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Eq8 x y)
	// result: (Equal (CMPW (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (ANDconst [1] (EQV x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64ANDconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpPPC64EQV, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (EqPtr x y)
	// result: (Equal (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Geq32F x y)
	// result: (FGreaterEqual (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64FGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Geq64F x y)
	// result: (FGreaterEqual (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64FGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGreater32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Greater32F x y)
	// result: (FGreaterThan (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64FGreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGreater64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Greater64F x y)
	// result: (FGreaterThan (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64FGreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpIsInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsInBounds idx len)
	// result: (LessThan (CMPU idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsNonNil ptr)
	// result: (NotEqual (CMPconst [0] ptr))
	for {
		ptr := v_0
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(ptr)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpIsSliceInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsSliceInBounds idx len)
	// result: (LessEqual (CMPU idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (LessEqual (CMPW (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (LessEqual (CMPWU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32 x y)
	// result: (LessEqual (CMPW x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32F x y)
	// result: (FLessEqual (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64FLessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32U x y)
	// result: (LessEqual (CMPWU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64 x y)
	// result: (LessEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64F x y)
	// result: (FLessEqual (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64FLessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64U x y)
	// result: (LessEqual (CMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (LessEqual (CMPW (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (LessEqual (CMPWU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (LessThan (CMPW (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (LessThan (CMPWU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32 x y)
	// result: (LessThan (CMPW x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32F x y)
	// result: (FLessThan (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64FLessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32U x y)
	// result: (LessThan (CMPWU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64 x y)
	// result: (LessThan (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64F x y)
	// result: (FLessThan (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64FLessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64U x y)
	// result: (LessThan (CMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (LessThan (CMPW (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (LessThan (CMPWU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Load <t> ptr mem)
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitInt(t) && isSigned(t)
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitInt(t) && !isSigned(t)
	// result: (MOVWZload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is16BitInt(t) && isSigned(t)
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is16BitInt(t) && !isSigned(t)
	// result: (MOVHZload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBZload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is8BitInt(t) && isSigned(t)
	// result: (MOVBreg (MOVBZload ptr mem))
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVBreg)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, typ.UInt8)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is8BitInt(t) && !isSigned(t)
	// result: (MOVBZload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (FMOVSload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpPPC64FMOVSload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (FMOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpPPC64FMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpLocalAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LocalAddr {sym} base _)
	// result: (MOVDaddr {sym} base)
	for {
		sym := v.Aux
		base := v_0
		v.reset(OpPPC64MOVDaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValuePPC64_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh16x16 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -16
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x32 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x32 x (MOVDconst [c]))
	// cond: uint32(c) < 16
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh16x32 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -16
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh16x64 x (MOVDconst [c]))
	// cond: uint64(c) < 16
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh16x64 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -16
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh16x8 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -16
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh32x16 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x32 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x32 x (MOVDconst [c]))
	// cond: uint32(c) < 32
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh32x32 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh32x64 x (MOVDconst [c]))
	// cond: uint64(c) < 32
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh32x64 x (AND y (MOVDconst [31])))
	// result: (SLW x (ANDconst <typ.Int32> [31] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			y := v_1_0
			if v_1_1.Op != OpPPC64MOVDconst || v_1_1.AuxInt != 31 {
				continue
			}
			v.reset(OpPPC64SLW)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
			v0.AuxInt = 31
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Lsh32x64 x (ANDconst <typ.Int32> [31] y))
	// result: (SLW x (ANDconst <typ.Int32> [31] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64ANDconst || v_1.Type != typ.Int32 || v_1.AuxInt != 31 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (Lsh32x64 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh32x8 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SLD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh64x16 x y)
	// result: (SLD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x32 x (Const64 [c]))
	// cond: uint32(c) < 64
	// result: (SLDconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x32 x (MOVDconst [c]))
	// cond: uint32(c) < 64
	// result: (SLDconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SLD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh64x32 x y)
	// result: (SLD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x64 x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SLDconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh64x64 x (MOVDconst [c]))
	// cond: uint64(c) < 64
	// result: (SLDconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SLD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh64x64 x (AND y (MOVDconst [63])))
	// result: (SLD x (ANDconst <typ.Int64> [63] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			y := v_1_0
			if v_1_1.Op != OpPPC64MOVDconst || v_1_1.AuxInt != 63 {
				continue
			}
			v.reset(OpPPC64SLD)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
			v0.AuxInt = 63
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Lsh64x64 x (ANDconst <typ.Int64> [63] y))
	// result: (SLD x (ANDconst <typ.Int64> [63] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64ANDconst || v_1.Type != typ.Int64 || v_1.AuxInt != 63 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (Lsh64x64 x y)
	// result: (SLD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SLD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh64x8 x y)
	// result: (SLD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh8x16 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -8
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x32 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x32 x (MOVDconst [c]))
	// cond: uint32(c) < 8
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh8x32 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -8
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh8x64 x (MOVDconst [c]))
	// cond: uint64(c) < 8
	// result: (SLWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh8x64 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -8
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SLW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Lsh8x8 x y)
	// result: (SLW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -8
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (Mod32 (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (Mod32u (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 x y)
	// result: (SUB x (MULLW y (DIVW x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLW, typ.Int32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVW, typ.Int32)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (SUB x (MULLW y (DIVWU x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLW, typ.Int32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVWU, typ.Int32)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64 x y)
	// result: (SUB x (MULLD y (DIVD x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLD, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVD, typ.Int64)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64u x y)
	// result: (SUB x (MULLD y (DIVDU x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLD, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVDU, typ.Int64)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (Mod32 (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (Mod32u (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBZload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpPPC64MOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, typ.UInt8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (MOVHstore dst (MOVHZload src mem) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpPPC64MOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, typ.UInt16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (MOVWstore dst (MOVWZload src mem) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpPPC64MOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVDstore dst (MOVDload src mem) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, typ.Int64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [8] dst src mem)
	// result: (MOVWstore [4] dst (MOVWZload [4] src mem) (MOVWstore dst (MOVWZload src mem) mem))
	for {
		if v.AuxInt != 8 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [3] dst src mem)
	// result: (MOVBstore [2] dst (MOVBZload [2] src mem) (MOVHstore dst (MOVHload src mem) mem))
	for {
		if v.AuxInt != 3 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVHload, typ.Int16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [5] dst src mem)
	// result: (MOVBstore [4] dst (MOVBZload [4] src mem) (MOVWstore dst (MOVWZload src mem) mem))
	for {
		if v.AuxInt != 5 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, typ.UInt8)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [6] dst src mem)
	// result: (MOVHstore [4] dst (MOVHZload [4] src mem) (MOVWstore dst (MOVWZload src mem) mem))
	for {
		if v.AuxInt != 6 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, typ.UInt16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [7] dst src mem)
	// result: (MOVBstore [6] dst (MOVBZload [6] src mem) (MOVHstore [4] dst (MOVHZload [4] src mem) (MOVWstore dst (MOVWZload src mem) mem)))
	for {
		if v.AuxInt != 7 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, typ.UInt8)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHstore, types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVHZload, typ.UInt16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpPPC64MOVWstore, types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 8
	// result: (LoweredMove [s] dst src mem)
	for {
		s := v.AuxInt
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 8) {
			break
		}
		v.reset(OpPPC64LoweredMove)
		v.AuxInt = s
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// cond: isSigned(x.Type) && isSigned(y.Type)
	// result: (NotEqual (CMPW (SignExt16to32 x) (SignExt16to32 y)))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			y := v_1
			if !(isSigned(x.Type) && isSigned(y.Type)) {
				continue
			}
			v.reset(OpPPC64NotEqual)
			v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
			v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
			v1.AddArg(x)
			v0.AddArg(v1)
			v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
			v2.AddArg(y)
			v0.AddArg(v2)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Neq16 x y)
	// result: (NotEqual (CMPW (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32 x y)
	// result: (NotEqual (CMPW x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32F x y)
	// result: (NotEqual (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64 x y)
	// result: (NotEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64F x y)
	// result: (NotEqual (FCMPU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// cond: isSigned(x.Type) && isSigned(y.Type)
	// result: (NotEqual (CMPW (SignExt8to32 x) (SignExt8to32 y)))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			y := v_1
			if !(isSigned(x.Type) && isSigned(y.Type)) {
				continue
			}
			v.reset(OpPPC64NotEqual)
			v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
			v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
			v1.AddArg(x)
			v0.AddArg(v1)
			v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
			v2.AddArg(y)
			v0.AddArg(v2)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Neq8 x y)
	// result: (NotEqual (CMPW (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (NeqPtr x y)
	// result: (NotEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Not x)
	// result: (XORconst [1] x)
	for {
		x := v_0
		v.reset(OpPPC64XORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (OffPtr [off] ptr)
	// result: (ADD (MOVDconst <typ.Int64> [off]) ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		v.reset(OpPPC64ADD)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = off
		v.AddArg(v0)
		v.AddArg(ptr)
		return true
	}
}
func rewriteValuePPC64_OpPPC64ADD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ADD (SLDconst x [c]) (SRDconst x [d]))
	// cond: d == 64-c
	// result: (ROTLconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLDconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpPPC64SRDconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 64-c) {
				continue
			}
			v.reset(OpPPC64ROTLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADD (SLWconst x [c]) (SRWconst x [d]))
	// cond: d == 32-c
	// result: (ROTLWconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLWconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpPPC64SRWconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 32-c) {
				continue
			}
			v.reset(OpPPC64ROTLWconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADD (SLD x (ANDconst <typ.Int64> [63] y)) (SRD x (SUB <typ.UInt> (MOVDconst [64]) (ANDconst <typ.UInt> [63] y))))
	// result: (ROTL x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLD {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpPPC64ANDconst || v_0_1.Type != typ.Int64 || v_0_1.AuxInt != 63 {
				continue
			}
			y := v_0_1.Args[0]
			if v_1.Op != OpPPC64SRD {
				continue
			}
			_ = v_1.Args[1]
			if x != v_1.Args[0] {
				continue
			}
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpPPC64SUB || v_1_1.Type != typ.UInt {
				continue
			}
			_ = v_1_1.Args[1]
			v_1_1_0 := v_1_1.Args[0]
			if v_1_1_0.Op != OpPPC64MOVDconst || v_1_1_0.AuxInt != 64 {
				continue
			}
			v_1_1_1 := v_1_1.Args[1]
			if v_1_1_1.Op != OpPPC64ANDconst || v_1_1_1.Type != typ.UInt || v_1_1_1.AuxInt != 63 || y != v_1_1_1.Args[0] {
				continue
			}
			v.reset(OpPPC64ROTL)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADD (SLW x (ANDconst <typ.Int32> [31] y)) (SRW x (SUB <typ.UInt> (MOVDconst [32]) (ANDconst <typ.UInt> [31] y))))
	// result: (ROTLW x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLW {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpPPC64ANDconst || v_0_1.Type != typ.Int32 || v_0_1.AuxInt != 31 {
				continue
			}
			y := v_0_1.Args[0]
			if v_1.Op != OpPPC64SRW {
				continue
			}
			_ = v_1.Args[1]
			if x != v_1.Args[0] {
				continue
			}
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpPPC64SUB || v_1_1.Type != typ.UInt {
				continue
			}
			_ = v_1_1.Args[1]
			v_1_1_0 := v_1_1.Args[0]
			if v_1_1_0.Op != OpPPC64MOVDconst || v_1_1_0.AuxInt != 32 {
				continue
			}
			v_1_1_1 := v_1_1.Args[1]
			if v_1_1_1.Op != OpPPC64ANDconst || v_1_1_1.Type != typ.UInt || v_1_1_1.AuxInt != 31 || y != v_1_1_1.Args[0] {
				continue
			}
			v.reset(OpPPC64ROTLW)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADD x (MOVDconst [c]))
	// cond: is32Bit(c)
	// result: (ADDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpPPC64MOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpPPC64ADDconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuePPC64_OpPPC64ADDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDconst [c] (ADDconst [d] x))
	// cond: is32Bit(c+d)
	// result: (ADDconst [c+d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpPPC64ADDconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (MOVDaddr [d] {sym} x))
	// result: (MOVDaddr [c+d] {sym} x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		d := v_0.AuxInt
		sym := v_0.Aux
		x := v_0.Args[0]
		v.reset(OpPPC64MOVDaddr)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64AND(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AND x (NOR y y))
	// result: (ANDN x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpPPC64NOR {
				continue
			}
			y := v_1.Args[1]
			if y != v_1.Args[0] {
				continue
			}
			v.reset(OpPPC64ANDN)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (AND (MOVDconst [c]) (MOVDconst [d]))
	// result: (MOVDconst [c&d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64MOVDconst {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpPPC64MOVDconst {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpPPC64MOVDconst)
			v.AuxInt = c & d
			return true
		}
		break
	}
	// match: (AND x (MOVDconst [c]))
	// cond: isU16Bit(c)
	// result: (ANDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpPPC64MOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(isU16Bit(c)) {
				continue
			}
			v.reset(OpPPC64ANDconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND (MOVDconst [c]) y:(MOVWZreg _))
	// cond: c&0xFFFFFFFF == 0xFFFFFFFF
	// result: y
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64MOVDconst {
				continue
			}
			c := v_0.AuxInt
			y := v_1
			if y.Op != OpPPC64MOVWZreg || !(c&0xFFFFFFFF == 0xFFFFFFFF) {
				continue
			}
			v.reset(OpCopy)
			v.Type = y.Type
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (AND (MOVDconst [0xFFFFFFFF]) y:(MOVWreg x))
	// result: (MOVWZreg x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64MOVDconst || v_0.AuxInt != 0xFFFFFFFF {
				continue
			}
			y := v_1
			if y.Op != OpPPC64MOVWreg {
				continue
			}
			x := y.Args[0]
			v.reset(OpPPC64MOVWZreg)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND (MOVDconst [c]) x:(MOVBZload _ _))
	// result: (ANDconst [c&0xFF] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64MOVDconst {
				continue
			}
			c := v_0.AuxInt
			x := v_1
			if x.Op != OpPPC64MOVBZload {
				continue
			}
			_ = x.Args[1]
			v.reset(OpPPC64ANDconst)
			v.AuxInt = c & 0xFF
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuePPC64_OpPPC64ANDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ANDconst [c] (ANDconst [d] x))
	// result: (ANDconst [c&d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64ANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [-1] x)
	// result: x
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [0] _)
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDconst [c] y:(MOVBZreg _))
	// cond: c&0xFF == 0xFF
	// result: y
	for {
		c := v.AuxInt
		y := v_0
		if y.Op != OpPPC64MOVBZreg || !(c&0xFF == 0xFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ANDconst [0xFF] y:(MOVBreg _))
	// result: y
	for {
		if v.AuxInt != 0xFF {
			break
		}
		y := v_0
		if y.Op != OpPPC64MOVBreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ANDconst [c] y:(MOVHZreg _))
	// cond: c&0xFFFF == 0xFFFF
	// result: y
	for {
		c := v.AuxInt
		y := v_0
		if y.Op != OpPPC64MOVHZreg || !(c&0xFFFF == 0xFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ANDconst [0xFFFF] y:(MOVHreg _))
	// result: y
	for {
		if v.AuxInt != 0xFFFF {
			break
		}
		y := v_0
		if y.Op != OpPPC64MOVHreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ANDconst [c] (MOVBreg x))
	// result: (ANDconst [c&0xFF] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64MOVBreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVBZreg x))
	// result: (ANDconst [c&0xFF] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVHreg x))
	// result: (ANDconst [c&0xFFFF] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64MOVHreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFFFF
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVHZreg x))
	// result: (ANDconst [c&0xFFFF] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFFFF
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVWreg x))
	// result: (ANDconst [c&0xFFFFFFFF] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64MOVWreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFFFFFFFF
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVWZreg x))
	// result: (ANDconst [c&0xFFFFFFFF] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFFFFFFFF
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMP(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMP x (MOVDconst [c]))
	// cond: is16Bit(c)
	// result: (CMPconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMP (MOVDconst [c]) y)
	// cond: is16Bit(c)
	// result: (InvertFlags (CMPconst y [c]))
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v_1
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPconst, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x y)
	// cond: x.ID > y.ID
	// result: (InvertFlags (CMP y x))
	for {
		x := v_0
		y := v_1
		if !(x.ID > y.ID) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPU x (MOVDconst [c]))
	// cond: isU16Bit(c)
	// result: (CMPUconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPU (MOVDconst [c]) y)
	// cond: isU16Bit(c)
	// result: (InvertFlags (CMPUconst y [c]))
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v_1
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPUconst, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMPU x y)
	// cond: x.ID > y.ID
	// result: (InvertFlags (CMPU y x))
	for {
		x := v_0
		y := v_1
		if !(x.ID > y.ID) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPUconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMPUconst (MOVDconst [x]) [y])
	// cond: x==y
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x == y) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}
	// match: (CMPUconst (MOVDconst [x]) [y])
	// cond: uint64(x)<uint64(y)
	// result: (FlagLT)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) < uint64(y)) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}
	// match: (CMPUconst (MOVDconst [x]) [y])
	// cond: uint64(x)>uint64(y)
	// result: (FlagGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) > uint64(y)) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPW x (MOVWreg y))
	// result: (CMPW x y)
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64CMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPW (MOVWreg x) y)
	// result: (CMPW x y)
	for {
		if v_0.Op != OpPPC64MOVWreg {
			break
		}
		x := v_0.Args[0]
		y := v_1
		v.reset(OpPPC64CMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPW x (MOVDconst [c]))
	// cond: is16Bit(c)
	// result: (CMPWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPW (MOVDconst [c]) y)
	// cond: is16Bit(c)
	// result: (InvertFlags (CMPWconst y [c]))
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v_1
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWconst, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMPW x y)
	// cond: x.ID > y.ID
	// result: (InvertFlags (CMPW y x))
	for {
		x := v_0
		y := v_1
		if !(x.ID > y.ID) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPWU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPWU x (MOVWZreg y))
	// result: (CMPWU x y)
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64CMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPWU (MOVWZreg x) y)
	// result: (CMPWU x y)
	for {
		if v_0.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_0.Args[0]
		y := v_1
		v.reset(OpPPC64CMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPWU x (MOVDconst [c]))
	// cond: isU16Bit(c)
	// result: (CMPWUconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPWU (MOVDconst [c]) y)
	// cond: isU16Bit(c)
	// result: (InvertFlags (CMPWUconst y [c]))
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v_1
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWUconst, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMPWU x y)
	// cond: x.ID > y.ID
	// result: (InvertFlags (CMPWU y x))
	for {
		x := v_0
		y := v_1
		if !(x.ID > y.ID) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPWUconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMPWUconst (MOVDconst [x]) [y])
	// cond: int32(x)==int32(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}
	// match: (CMPWUconst (MOVDconst [x]) [y])
	// cond: uint32(x)<uint32(y)
	// result: (FlagLT)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}
	// match: (CMPWUconst (MOVDconst [x]) [y])
	// cond: uint32(x)>uint32(y)
	// result: (FlagGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPWconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMPWconst (MOVDconst [x]) [y])
	// cond: int32(x)==int32(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}
	// match: (CMPWconst (MOVDconst [x]) [y])
	// cond: int32(x)<int32(y)
	// result: (FlagLT)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y)) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}
	// match: (CMPWconst (MOVDconst [x]) [y])
	// cond: int32(x)>int32(y)
	// result: (FlagGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y)) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMPconst (MOVDconst [x]) [y])
	// cond: x==y
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x == y) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}
	// match: (CMPconst (MOVDconst [x]) [y])
	// cond: x<y
	// result: (FlagLT)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x < y) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}
	// match: (CMPconst (MOVDconst [x]) [y])
	// cond: x>y
	// result: (FlagGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x > y) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64Equal(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Equal (FlagEQ))
	// result: (MOVDconst [1])
	for {
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (Equal (FlagLT))
	// result: (MOVDconst [0])
	for {
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (FlagGT))
	// result: (MOVDconst [0])
	for {
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (InvertFlags x))
	// result: (Equal x)
	for {
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64Equal)
		v.AddArg(x)
		return true
	}
	// match: (Equal cmp)
	// result: (ISELB [2] (MOVDconst [1]) cmp)
	for {
		cmp := v_0
		v.reset(OpPPC64ISELB)
		v.AuxInt = 2
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(cmp)
		return true
	}
}
func rewriteValuePPC64_OpPPC64FABS(v *Value) bool {
	v_0 := v.Args[0]
	// match: (FABS (FMOVDconst [x]))
	// result: (FMOVDconst [auxFrom64F(math.Abs(auxTo64F(x)))])
	for {
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		x := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = auxFrom64F(math.Abs(auxTo64F(x)))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FADD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FADD (FMUL x y) z)
	// result: (FMADD x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64FMUL {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			z := v_1
			v.reset(OpPPC64FMADD)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	return false
}
func rewriteValuePPC64_OpPPC64FADDS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FADDS (FMULS x y) z)
	// result: (FMADDS x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64FMULS {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			z := v_1
			v.reset(OpPPC64FMADDS)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	return false
}
func rewriteValuePPC64_OpPPC64FCEIL(v *Value) bool {
	v_0 := v.Args[0]
	// match: (FCEIL (FMOVDconst [x]))
	// result: (FMOVDconst [auxFrom64F(math.Ceil(auxTo64F(x)))])
	for {
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		x := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = auxFrom64F(math.Ceil(auxTo64F(x)))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FFLOOR(v *Value) bool {
	v_0 := v.Args[0]
	// match: (FFLOOR (FMOVDconst [x]))
	// result: (FMOVDconst [auxFrom64F(math.Floor(auxTo64F(x)))])
	for {
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		x := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = auxFrom64F(math.Floor(auxTo64F(x)))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FGreaterEqual(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (FGreaterEqual cmp)
	// result: (ISEL [2] (MOVDconst [1]) (ISELB [1] (MOVDconst [1]) cmp) cmp)
	for {
		cmp := v_0
		v.reset(OpPPC64ISEL)
		v.AuxInt = 2
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ISELB, typ.Int32)
		v1.AuxInt = 1
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v2.AuxInt = 1
		v1.AddArg(v2)
		v1.AddArg(cmp)
		v.AddArg(v1)
		v.AddArg(cmp)
		return true
	}
}
func rewriteValuePPC64_OpPPC64FGreaterThan(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (FGreaterThan cmp)
	// result: (ISELB [1] (MOVDconst [1]) cmp)
	for {
		cmp := v_0
		v.reset(OpPPC64ISELB)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(cmp)
		return true
	}
}
func rewriteValuePPC64_OpPPC64FLessEqual(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (FLessEqual cmp)
	// result: (ISEL [2] (MOVDconst [1]) (ISELB [0] (MOVDconst [1]) cmp) cmp)
	for {
		cmp := v_0
		v.reset(OpPPC64ISEL)
		v.AuxInt = 2
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ISELB, typ.Int32)
		v1.AuxInt = 0
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v2.AuxInt = 1
		v1.AddArg(v2)
		v1.AddArg(cmp)
		v.AddArg(v1)
		v.AddArg(cmp)
		return true
	}
}
func rewriteValuePPC64_OpPPC64FLessThan(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (FLessThan cmp)
	// result: (ISELB [0] (MOVDconst [1]) cmp)
	for {
		cmp := v_0
		v.reset(OpPPC64ISELB)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(cmp)
		return true
	}
}
func rewriteValuePPC64_OpPPC64FMOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVDload [off] {sym} ptr (MOVDstore [off] {sym} ptr x _))
	// result: (MTVSRD x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVDstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		if ptr != v_1.Args[0] {
			break
		}
		x := v_1.Args[1]
		v.reset(OpPPC64MTVSRD)
		v.AddArg(x)
		return true
	}
	// match: (FMOVDload [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (FMOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64FMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (FMOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVDstore [off] {sym} ptr (MTVSRD x) mem)
	// result: (MOVDstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MTVSRD {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: is16Bit(off1+off2)
	// result: (FMOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDstore [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (FMOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVSload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVSload [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (FMOVSload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64FMOVSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (FMOVSload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVSstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVSstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: is16Bit(off1+off2)
	// result: (FMOVSstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVSstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSstore [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (FMOVSstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64FMOVSstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FNEG(v *Value) bool {
	v_0 := v.Args[0]
	// match: (FNEG (FABS x))
	// result: (FNABS x)
	for {
		if v_0.Op != OpPPC64FABS {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64FNABS)
		v.AddArg(x)
		return true
	}
	// match: (FNEG (FNABS x))
	// result: (FABS x)
	for {
		if v_0.Op != OpPPC64FNABS {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64FABS)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FSQRT(v *Value) bool {
	v_0 := v.Args[0]
	// match: (FSQRT (FMOVDconst [x]))
	// result: (FMOVDconst [auxFrom64F(math.Sqrt(auxTo64F(x)))])
	for {
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		x := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = auxFrom64F(math.Sqrt(auxTo64F(x)))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FSUB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FSUB (FMUL x y) z)
	// result: (FMSUB x y z)
	for {
		if v_0.Op != OpPPC64FMUL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		z := v_1
		v.reset(OpPPC64FMSUB)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FSUBS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FSUBS (FMULS x y) z)
	// result: (FMSUBS x y z)
	for {
		if v_0.Op != OpPPC64FMULS {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		z := v_1
		v.reset(OpPPC64FMSUBS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FTRUNC(v *Value) bool {
	v_0 := v.Args[0]
	// match: (FTRUNC (FMOVDconst [x]))
	// result: (FMOVDconst [auxFrom64F(math.Trunc(auxTo64F(x)))])
	for {
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		x := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = auxFrom64F(math.Trunc(auxTo64F(x)))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64GreaterEqual(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (GreaterEqual (FlagEQ))
	// result: (MOVDconst [1])
	for {
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqual (FlagLT))
	// result: (MOVDconst [0])
	for {
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterEqual (FlagGT))
	// result: (MOVDconst [1])
	for {
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqual (InvertFlags x))
	// result: (LessEqual x)
	for {
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64LessEqual)
		v.AddArg(x)
		return true
	}
	// match: (GreaterEqual cmp)
	// result: (ISELB [4] (MOVDconst [1]) cmp)
	for {
		cmp := v_0
		v.reset(OpPPC64ISELB)
		v.AuxInt = 4
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(cmp)
		return true
	}
}
func rewriteValuePPC64_OpPPC64GreaterThan(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (GreaterThan (FlagEQ))
	// result: (MOVDconst [0])
	for {
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThan (FlagLT))
	// result: (MOVDconst [0])
	for {
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThan (FlagGT))
	// result: (MOVDconst [1])
	for {
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterThan (InvertFlags x))
	// result: (LessThan x)
	for {
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64LessThan)
		v.AddArg(x)
		return true
	}
	// match: (GreaterThan cmp)
	// result: (ISELB [1] (MOVDconst [1]) cmp)
	for {
		cmp := v_0
		v.reset(OpPPC64ISELB)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(cmp)
		return true
	}
}
func rewriteValuePPC64_OpPPC64ISEL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ISEL [2] x _ (FlagEQ))
	// result: x
	for {
		if v.AuxInt != 2 {
			break
		}
		x := v_0
		if v_2.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ISEL [2] _ y (FlagLT))
	// result: y
	for {
		if v.AuxInt != 2 {
			break
		}
		y := v_1
		if v_2.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ISEL [2] _ y (FlagGT))
	// result: y
	for {
		if v.AuxInt != 2 {
			break
		}
		y := v_1
		if v_2.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ISEL [6] _ y (FlagEQ))
	// result: y
	for {
		if v.AuxInt != 6 {
			break
		}
		y := v_1
		if v_2.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ISEL [6] x _ (FlagLT))
	// result: x
	for {
		if v.AuxInt != 6 {
			break
		}
		x := v_0
		if v_2.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ISEL [6] x _ (FlagGT))
	// result: x
	for {
		if v.AuxInt != 6 {
			break
		}
		x := v_0
		if v_2.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ISEL [0] _ y (FlagEQ))
	// result: y
	for {
		if v.AuxInt != 0 {
			break
		}
		y := v_1
		if v_2.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ISEL [0] _ y (FlagGT))
	// result: y
	for {
		if v.AuxInt != 0 {
			break
		}
		y := v_1
		if v_2.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ISEL [0] x _ (FlagLT))
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		if v_2.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ISEL [5] _ x (FlagEQ))
	// result: x
	for {
		if v.AuxInt != 5 {
			break
		}
		x := v_1
		if v_2.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ISEL [5] _ x (FlagLT))
	// result: x
	for {
		if v.AuxInt != 5 {
			break
		}
		x := v_1
		if v_2.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ISEL [5] y _ (FlagGT))
	// result: y
	for {
		if v.AuxInt != 5 {
			break
		}
		y := v_0
		if v_2.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ISEL [1] _ y (FlagEQ))
	// result: y
	for {
		if v.AuxInt != 1 {
			break
		}
		y := v_1
		if v_2.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ISEL [1] _ y (FlagLT))
	// result: y
	for {
		if v.AuxInt != 1 {
			break
		}
		y := v_1
		if v_2.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ISEL [1] x _ (FlagGT))
	// result: x
	for {
		if v.AuxInt != 1 {
			break
		}
		x := v_0
		if v_2.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ISEL [4] x _ (FlagEQ))
	// result: x
	for {
		if v.AuxInt != 4 {
			break
		}
		x := v_0
		if v_2.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ISEL [4] x _ (FlagGT))
	// result: x
	for {
		if v.AuxInt != 4 {
			break
		}
		x := v_0
		if v_2.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ISEL [4] _ y (FlagLT))
	// result: y
	for {
		if v.AuxInt != 4 {
			break
		}
		y := v_1
		if v_2.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ISEL [n] x y (InvertFlags bool))
	// cond: n%4 == 0
	// result: (ISEL [n+1] x y bool)
	for {
		n := v.AuxInt
		x := v_0
		y := v_1
		if v_2.Op != OpPPC64InvertFlags {
			break
		}
		bool := v_2.Args[0]
		if !(n%4 == 0) {
			break
		}
		v.reset(OpPPC64ISEL)
		v.AuxInt = n + 1
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(bool)
		return true
	}
	// match: (ISEL [n] x y (InvertFlags bool))
	// cond: n%4 == 1
	// result: (ISEL [n-1] x y bool)
	for {
		n := v.AuxInt
		x := v_0
		y := v_1
		if v_2.Op != OpPPC64InvertFlags {
			break
		}
		bool := v_2.Args[0]
		if !(n%4 == 1) {
			break
		}
		v.reset(OpPPC64ISEL)
		v.AuxInt = n - 1
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(bool)
		return true
	}
	// match: (ISEL [n] x y (InvertFlags bool))
	// cond: n%4 == 2
	// result: (ISEL [n] x y bool)
	for {
		n := v.AuxInt
		x := v_0
		y := v_1
		if v_2.Op != OpPPC64InvertFlags {
			break
		}
		bool := v_2.Args[0]
		if !(n%4 == 2) {
			break
		}
		v.reset(OpPPC64ISEL)
		v.AuxInt = n
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(bool)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ISELB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ISELB [0] _ (FlagLT))
	// result: (MOVDconst [1])
	for {
		if v.AuxInt != 0 || v_1.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (ISELB [0] _ (FlagGT))
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 0 || v_1.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ISELB [0] _ (FlagEQ))
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 0 || v_1.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ISELB [1] _ (FlagGT))
	// result: (MOVDconst [1])
	for {
		if v.AuxInt != 1 || v_1.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (ISELB [1] _ (FlagLT))
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 1 || v_1.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ISELB [1] _ (FlagEQ))
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 1 || v_1.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ISELB [2] _ (FlagEQ))
	// result: (MOVDconst [1])
	for {
		if v.AuxInt != 2 || v_1.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (ISELB [2] _ (FlagLT))
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 2 || v_1.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ISELB [2] _ (FlagGT))
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 2 || v_1.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ISELB [4] _ (FlagLT))
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 4 || v_1.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ISELB [4] _ (FlagGT))
	// result: (MOVDconst [1])
	for {
		if v.AuxInt != 4 || v_1.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (ISELB [4] _ (FlagEQ))
	// result: (MOVDconst [1])
	for {
		if v.AuxInt != 4 || v_1.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (ISELB [5] _ (FlagGT))
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 5 || v_1.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ISELB [5] _ (FlagLT))
	// result: (MOVDconst [1])
	for {
		if v.AuxInt != 5 || v_1.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (ISELB [5] _ (FlagEQ))
	// result: (MOVDconst [1])
	for {
		if v.AuxInt != 5 || v_1.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (ISELB [6] _ (FlagEQ))
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 6 || v_1.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ISELB [6] _ (FlagLT))
	// result: (MOVDconst [1])
	for {
		if v.AuxInt != 6 || v_1.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (ISELB [6] _ (FlagGT))
	// result: (MOVDconst [1])
	for {
		if v.AuxInt != 6 || v_1.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (ISELB [n] (MOVDconst [1]) (InvertFlags bool))
	// cond: n%4 == 0
	// result: (ISELB [n+1] (MOVDconst [1]) bool)
	for {
		n := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst || v_0.AuxInt != 1 || v_1.Op != OpPPC64InvertFlags {
			break
		}
		bool := v_1.Args[0]
		if !(n%4 == 0) {
			break
		}
		v.reset(OpPPC64ISELB)
		v.AuxInt = n + 1
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(bool)
		return true
	}
	// match: (ISELB [n] (MOVDconst [1]) (InvertFlags bool))
	// cond: n%4 == 1
	// result: (ISELB [n-1] (MOVDconst [1]) bool)
	for {
		n := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst || v_0.AuxInt != 1 || v_1.Op != OpPPC64InvertFlags {
			break
		}
		bool := v_1.Args[0]
		if !(n%4 == 1) {
			break
		}
		v.reset(OpPPC64ISELB)
		v.AuxInt = n - 1
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(bool)
		return true
	}
	// match: (ISELB [n] (MOVDconst [1]) (InvertFlags bool))
	// cond: n%4 == 2
	// result: (ISELB [n] (MOVDconst [1]) bool)
	for {
		n := v.AuxInt
		if v_0.Op != OpPPC64MOVDconst || v_0.AuxInt != 1 || v_1.Op != OpPPC64InvertFlags {
			break
		}
		bool := v_1.Args[0]
		if !(n%4 == 2) {
			break
		}
		v.reset(OpPPC64ISELB)
		v.AuxInt = n
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(bool)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64LessEqual(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (LessEqual (FlagEQ))
	// result: (MOVDconst [1])
	for {
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqual (FlagLT))
	// result: (MOVDconst [1])
	for {
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqual (FlagGT))
	// result: (MOVDconst [0])
	for {
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessEqual (InvertFlags x))
	// result: (GreaterEqual x)
	for {
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64GreaterEqual)
		v.AddArg(x)
		return true
	}
	// match: (LessEqual cmp)
	// result: (ISELB [5] (MOVDconst [1]) cmp)
	for {
		cmp := v_0
		v.reset(OpPPC64ISELB)
		v.AuxInt = 5
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(cmp)
		return true
	}
}
func rewriteValuePPC64_OpPPC64LessThan(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (LessThan (FlagEQ))
	// result: (MOVDconst [0])
	for {
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThan (FlagLT))
	// result: (MOVDconst [1])
	for {
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessThan (FlagGT))
	// result: (MOVDconst [0])
	for {
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThan (InvertFlags x))
	// result: (GreaterThan x)
	for {
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64GreaterThan)
		v.AddArg(x)
		return true
	}
	// match: (LessThan cmp)
	// result: (ISELB [0] (MOVDconst [1]) cmp)
	for {
		cmp := v_0
		v.reset(OpPPC64ISELB)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(cmp)
		return true
	}
}
func rewriteValuePPC64_OpPPC64MFVSRD(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MFVSRD (FMOVDconst [c]))
	// result: (MOVDconst [c])
	for {
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c
		return true
	}
	// match: (MFVSRD x:(FMOVDload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVDload [off] {sym} ptr mem)
	for {
		x := v_0
		if x.Op != OpPPC64FMOVDload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpPPC64MOVDload, typ.Int64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBZload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBZload [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (MOVBZload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBZload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVBZload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBZload [0] {sym} p:(ADD ptr idx) mem)
	// cond: sym == nil && p.Uses == 1
	// result: (MOVBZloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		p := v_0
		if p.Op != OpPPC64ADD {
			break
		}
		idx := p.Args[1]
		ptr := p.Args[0]
		mem := v_1
		if !(sym == nil && p.Uses == 1) {
			break
		}
		v.reset(OpPPC64MOVBZloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBZloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBZloadidx ptr (MOVDconst [c]) mem)
	// cond: is16Bit(c)
	// result: (MOVBZload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBZloadidx (MOVDconst [c]) ptr mem)
	// cond: is16Bit(c)
	// result: (MOVBZload [c] ptr mem)
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBZreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVBZreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0xFF
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVBZreg (SRWconst [c] (MOVBZreg x)))
	// result: (SRWconst [c] (MOVBZreg x))
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVBZreg (SRWconst [c] x))
	// cond: sizeof(x.Type) == 8
	// result: (SRWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(sizeof(x.Type) == 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg (SRDconst [c] x))
	// cond: c>=56
	// result: (SRDconst [c] x)
	for {
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c >= 56) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg (SRWconst [c] x))
	// cond: c>=24
	// result: (SRWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c >= 24) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg y:(MOVBZreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVBZreg (MOVBreg x))
	// result: (MOVBZreg x)
	for {
		if v_0.Op != OpPPC64MOVBreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg x:(MOVBZload _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg x:(MOVBZloadidx _ _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVBZloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg x:(Arg <t>))
	// cond: is8BitInt(t) && !isSigned(t)
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg (MOVDconst [c]))
	// result: (MOVDconst [int64(uint8(c))])
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVBreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0x7F
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0x7F) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVBreg (SRAWconst [c] (MOVBreg x)))
	// result: (SRAWconst [c] (MOVBreg x))
	for {
		if v_0.Op != OpPPC64SRAWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVBreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVBreg (SRAWconst [c] x))
	// cond: sizeof(x.Type) == 8
	// result: (SRAWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRAWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(sizeof(x.Type) == 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (SRDconst [c] x))
	// cond: c>56
	// result: (SRDconst [c] x)
	for {
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c > 56) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (SRDconst [c] x))
	// cond: c==56
	// result: (SRADconst [c] x)
	for {
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c == 56) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (SRWconst [c] x))
	// cond: c>24
	// result: (SRWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c > 24) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (SRWconst [c] x))
	// cond: c==24
	// result: (SRAWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c == 24) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg y:(MOVBreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVBreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVBreg (MOVBZreg x))
	// result: (MOVBreg x)
	for {
		if v_0.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(Arg <t>))
	// cond: is8BitInt(t) && isSigned(t)
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVDconst [c]))
	// result: (MOVDconst [int64(int8(c))])
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (MOVBstore [off1] {sym} (ADDconst [off2] x) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} x val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVDconst [0]) mem)
	// result: (MOVBstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} p:(ADD ptr idx) val mem)
	// cond: off == 0 && sym == nil && p.Uses == 1
	// result: (MOVBstoreidx ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		p := v_0
		if p.Op != OpPPC64ADD {
			break
		}
		idx := p.Args[1]
		ptr := p.Args[0]
		val := v_1
		mem := v_2
		if !(off == 0 && sym == nil && p.Uses == 1) {
			break
		}
		v.reset(OpPPC64MOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBZreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHZreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWZreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (SRWconst (MOVHreg x) [c]) mem)
	// cond: c <= 8
	// result: (MOVBstore [off] {sym} ptr (SRWconst <typ.UInt32> x [c]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		c := v_1.AuxInt
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVHreg {
			break
		}
		x := v_1_0.Args[0]
		mem := v_2
		if !(c <= 8) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (SRWconst (MOVHZreg x) [c]) mem)
	// cond: c <= 8
	// result: (MOVBstore [off] {sym} ptr (SRWconst <typ.UInt32> x [c]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		c := v_1.AuxInt
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_1_0.Args[0]
		mem := v_2
		if !(c <= 8) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (SRWconst (MOVWreg x) [c]) mem)
	// cond: c <= 24
	// result: (MOVBstore [off] {sym} ptr (SRWconst <typ.UInt32> x [c]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		c := v_1.AuxInt
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1_0.Args[0]
		mem := v_2
		if !(c <= 24) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (SRWconst (MOVWZreg x) [c]) mem)
	// cond: c <= 24
	// result: (MOVBstore [off] {sym} ptr (SRWconst <typ.UInt32> x [c]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		c := v_1.AuxInt
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1_0.Args[0]
		mem := v_2
		if !(c <= 24) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i1] {s} p (SRWconst w [24]) x0:(MOVBstore [i0] {s} p (SRWconst w [16]) mem))
	// cond: !config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)
	// result: (MOVHstore [i0] {s} p (SRWconst <typ.UInt16> w [16]) mem)
	for {
		i1 := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpPPC64SRWconst || v_1.AuxInt != 24 {
			break
		}
		w := v_1.Args[0]
		x0 := v_2
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		mem := x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRWconst || x0_1.AuxInt != 16 || w != x0_1.Args[0] || !(!config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x0.Pos, OpPPC64SRWconst, typ.UInt16)
		v0.AuxInt = 16
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i1] {s} p (SRDconst w [24]) x0:(MOVBstore [i0] {s} p (SRDconst w [16]) mem))
	// cond: !config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)
	// result: (MOVHstore [i0] {s} p (SRWconst <typ.UInt16> w [16]) mem)
	for {
		i1 := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpPPC64SRDconst || v_1.AuxInt != 24 {
			break
		}
		w := v_1.Args[0]
		x0 := v_2
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		mem := x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRDconst || x0_1.AuxInt != 16 || w != x0_1.Args[0] || !(!config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(x0.Pos, OpPPC64SRWconst, typ.UInt16)
		v0.AuxInt = 16
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i1] {s} p (SRWconst w [8]) x0:(MOVBstore [i0] {s} p w mem))
	// cond: !config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)
	// result: (MOVHstore [i0] {s} p w mem)
	for {
		i1 := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpPPC64SRWconst || v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x0 := v_2
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		mem := x0.Args[2]
		if p != x0.Args[0] || w != x0.Args[1] || !(!config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i1] {s} p (SRDconst w [8]) x0:(MOVBstore [i0] {s} p w mem))
	// cond: !config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)
	// result: (MOVHstore [i0] {s} p w mem)
	for {
		i1 := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpPPC64SRDconst || v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x0 := v_2
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		mem := x0.Args[2]
		if p != x0.Args[0] || w != x0.Args[1] || !(!config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i3] {s} p w x0:(MOVBstore [i2] {s} p (SRWconst w [8]) x1:(MOVBstore [i1] {s} p (SRWconst w [16]) x2:(MOVBstore [i0] {s} p (SRWconst w [24]) mem))))
	// cond: !config.BigEndian && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && clobber(x0) && clobber(x1) && clobber(x2)
	// result: (MOVWBRstore (MOVDaddr <typ.Uintptr> [i0] {s} p) w mem)
	for {
		i3 := v.AuxInt
		s := v.Aux
		p := v_0
		w := v_1
		x0 := v_2
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i2 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRWconst || x0_1.AuxInt != 8 || w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpPPC64MOVBstore {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if p != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpPPC64SRWconst || x1_1.AuxInt != 16 || w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		mem := x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpPPC64SRWconst || x2_1.AuxInt != 24 || w != x2_1.Args[0] || !(!config.BigEndian && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpPPC64MOVWBRstore)
		v0 := b.NewValue0(x2.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i1] {s} p w x0:(MOVBstore [i0] {s} p (SRWconst w [8]) mem))
	// cond: !config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)
	// result: (MOVHBRstore (MOVDaddr <typ.Uintptr> [i0] {s} p) w mem)
	for {
		i1 := v.AuxInt
		s := v.Aux
		p := v_0
		w := v_1
		x0 := v_2
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		mem := x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRWconst || x0_1.AuxInt != 8 || w != x0_1.Args[0] || !(!config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVHBRstore)
		v0 := b.NewValue0(x0.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i7] {s} p (SRDconst w [56]) x0:(MOVBstore [i6] {s} p (SRDconst w [48]) x1:(MOVBstore [i5] {s} p (SRDconst w [40]) x2:(MOVBstore [i4] {s} p (SRDconst w [32]) x3:(MOVWstore [i0] {s} p w mem)))))
	// cond: !config.BigEndian && i0%4 == 0 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3)
	// result: (MOVDstore [i0] {s} p w mem)
	for {
		i7 := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpPPC64SRDconst || v_1.AuxInt != 56 {
			break
		}
		w := v_1.Args[0]
		x0 := v_2
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i6 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRDconst || x0_1.AuxInt != 48 || w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpPPC64MOVBstore {
			break
		}
		i5 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if p != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpPPC64SRDconst || x1_1.AuxInt != 40 || w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpPPC64MOVBstore {
			break
		}
		i4 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpPPC64SRDconst || x2_1.AuxInt != 32 || w != x2_1.Args[0] {
			break
		}
		x3 := x2.Args[2]
		if x3.Op != OpPPC64MOVWstore {
			break
		}
		i0 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		mem := x3.Args[2]
		if p != x3.Args[0] || w != x3.Args[1] || !(!config.BigEndian && i0%4 == 0 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3)) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [i7] {s} p w x0:(MOVBstore [i6] {s} p (SRDconst w [8]) x1:(MOVBstore [i5] {s} p (SRDconst w [16]) x2:(MOVBstore [i4] {s} p (SRDconst w [24]) x3:(MOVBstore [i3] {s} p (SRDconst w [32]) x4:(MOVBstore [i2] {s} p (SRDconst w [40]) x5:(MOVBstore [i1] {s} p (SRDconst w [48]) x6:(MOVBstore [i0] {s} p (SRDconst w [56]) mem))))))))
	// cond: !config.BigEndian && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6)
	// result: (MOVDBRstore (MOVDaddr <typ.Uintptr> [i0] {s} p) w mem)
	for {
		i7 := v.AuxInt
		s := v.Aux
		p := v_0
		w := v_1
		x0 := v_2
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i6 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRDconst || x0_1.AuxInt != 8 || w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpPPC64MOVBstore {
			break
		}
		i5 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if p != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpPPC64SRDconst || x1_1.AuxInt != 16 || w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpPPC64MOVBstore {
			break
		}
		i4 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpPPC64SRDconst || x2_1.AuxInt != 24 || w != x2_1.Args[0] {
			break
		}
		x3 := x2.Args[2]
		if x3.Op != OpPPC64MOVBstore {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[2]
		if p != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpPPC64SRDconst || x3_1.AuxInt != 32 || w != x3_1.Args[0] {
			break
		}
		x4 := x3.Args[2]
		if x4.Op != OpPPC64MOVBstore {
			break
		}
		i2 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[2]
		if p != x4.Args[0] {
			break
		}
		x4_1 := x4.Args[1]
		if x4_1.Op != OpPPC64SRDconst || x4_1.AuxInt != 40 || w != x4_1.Args[0] {
			break
		}
		x5 := x4.Args[2]
		if x5.Op != OpPPC64MOVBstore {
			break
		}
		i1 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[2]
		if p != x5.Args[0] {
			break
		}
		x5_1 := x5.Args[1]
		if x5_1.Op != OpPPC64SRDconst || x5_1.AuxInt != 48 || w != x5_1.Args[0] {
			break
		}
		x6 := x5.Args[2]
		if x6.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		mem := x6.Args[2]
		if p != x6.Args[0] {
			break
		}
		x6_1 := x6.Args[1]
		if x6_1.Op != OpPPC64SRDconst || x6_1.AuxInt != 56 || w != x6_1.Args[0] || !(!config.BigEndian && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6)) {
			break
		}
		v.reset(OpPPC64MOVDBRstore)
		v0 := b.NewValue0(x6.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVBstoreidx ptr (MOVDconst [c]) val mem)
	// cond: is16Bit(c)
	// result: (MOVBstore [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v_2
		mem := v_3
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx (MOVDconst [c]) ptr val mem)
	// cond: is16Bit(c)
	// result: (MOVBstore [c] ptr val mem)
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		val := v_2
		mem := v_3
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx [off] {sym} ptr idx (MOVBreg x) mem)
	// result: (MOVBstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVBreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx [off] {sym} ptr idx (MOVBZreg x) mem)
	// result: (MOVBstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx [off] {sym} ptr idx (MOVHreg x) mem)
	// result: (MOVBstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVHreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx [off] {sym} ptr idx (MOVHZreg x) mem)
	// result: (MOVBstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx [off] {sym} ptr idx (MOVWreg x) mem)
	// result: (MOVBstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVWreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx [off] {sym} ptr idx (MOVWZreg x) mem)
	// result: (MOVBstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx [off] {sym} ptr idx (SRWconst (MOVHreg x) [c]) mem)
	// cond: c <= 8
	// result: (MOVBstoreidx [off] {sym} ptr idx (SRWconst <typ.UInt32> x [c]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64SRWconst {
			break
		}
		c := v_2.AuxInt
		v_2_0 := v_2.Args[0]
		if v_2_0.Op != OpPPC64MOVHreg {
			break
		}
		x := v_2_0.Args[0]
		mem := v_3
		if !(c <= 8) {
			break
		}
		v.reset(OpPPC64MOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx [off] {sym} ptr idx (SRWconst (MOVHZreg x) [c]) mem)
	// cond: c <= 8
	// result: (MOVBstoreidx [off] {sym} ptr idx (SRWconst <typ.UInt32> x [c]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64SRWconst {
			break
		}
		c := v_2.AuxInt
		v_2_0 := v_2.Args[0]
		if v_2_0.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_2_0.Args[0]
		mem := v_3
		if !(c <= 8) {
			break
		}
		v.reset(OpPPC64MOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx [off] {sym} ptr idx (SRWconst (MOVWreg x) [c]) mem)
	// cond: c <= 24
	// result: (MOVBstoreidx [off] {sym} ptr idx (SRWconst <typ.UInt32> x [c]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64SRWconst {
			break
		}
		c := v_2.AuxInt
		v_2_0 := v_2.Args[0]
		if v_2_0.Op != OpPPC64MOVWreg {
			break
		}
		x := v_2_0.Args[0]
		mem := v_3
		if !(c <= 24) {
			break
		}
		v.reset(OpPPC64MOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx [off] {sym} ptr idx (SRWconst (MOVWZreg x) [c]) mem)
	// cond: c <= 24
	// result: (MOVBstoreidx [off] {sym} ptr idx (SRWconst <typ.UInt32> x [c]) mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64SRWconst {
			break
		}
		c := v_2.AuxInt
		v_2_0 := v_2.Args[0]
		if v_2_0.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_2_0.Args[0]
		mem := v_3
		if !(c <= 24) {
			break
		}
		v.reset(OpPPC64MOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstorezero [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstorezero [off1] {sym1} p:(MOVDaddr [off2] {sym2} x) mem)
	// cond: canMergeSym(sym1,sym2) && (x.Op != OpSB || p.Uses == 1)
	// result: (MOVBstorezero [off1+off2] {mergeSym(sym1,sym2)} x mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		x := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (x.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDload [off] {sym} ptr (FMOVDstore [off] {sym} ptr x _))
	// result: (MFVSRD x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64FMOVDstore || v_1.AuxInt != off || v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		if ptr != v_1.Args[0] {
			break
		}
		x := v_1.Args[1]
		v.reset(OpPPC64MFVSRD)
		v.AddArg(x)
		return true
	}
	// match: (MOVDload [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1) && (off1+off2)%4 == 0
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1) && (off1+off2)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2) && (off1+off2)%4 == 0
	// result: (MOVDload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) && (off1+off2)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [0] {sym} p:(ADD ptr idx) mem)
	// cond: sym == nil && p.Uses == 1
	// result: (MOVDloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		p := v_0
		if p.Op != OpPPC64ADD {
			break
		}
		idx := p.Args[1]
		ptr := p.Args[0]
		mem := v_1
		if !(sym == nil && p.Uses == 1) {
			break
		}
		v.reset(OpPPC64MOVDloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDloadidx ptr (MOVDconst [c]) mem)
	// cond: is16Bit(c) && c%4 == 0
	// result: (MOVDload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is16Bit(c) && c%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDloadidx (MOVDconst [c]) ptr mem)
	// cond: is16Bit(c) && c%4 == 0
	// result: (MOVDload [c] ptr mem)
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		if !(is16Bit(c) && c%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstore [off] {sym} ptr (MFVSRD x) mem)
	// result: (FMOVDstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MFVSRD {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64FMOVDstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym} (ADDconst [off2] x) val mem)
	// cond: is16Bit(off1+off2) && (off1+off2)%4 == 0
	// result: (MOVDstore [off1+off2] {sym} x val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1+off2) && (off1+off2)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1) && (off1+off2)%4 == 0
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1) && (off1+off2)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off] {sym} ptr (MOVDconst [0]) mem)
	// result: (MOVDstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off] {sym} p:(ADD ptr idx) val mem)
	// cond: off == 0 && sym == nil && p.Uses == 1
	// result: (MOVDstoreidx ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		p := v_0
		if p.Op != OpPPC64ADD {
			break
		}
		idx := p.Args[1]
		ptr := p.Args[0]
		val := v_1
		mem := v_2
		if !(off == 0 && sym == nil && p.Uses == 1) {
			break
		}
		v.reset(OpPPC64MOVDstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstoreidx ptr (MOVDconst [c]) val mem)
	// cond: is16Bit(c) && c%4 == 0
	// result: (MOVDstore [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v_2
		mem := v_3
		if !(is16Bit(c) && c%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstoreidx (MOVDconst [c]) ptr val mem)
	// cond: is16Bit(c) && c%4 == 0
	// result: (MOVDstore [c] ptr val mem)
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		val := v_2
		mem := v_3
		if !(is16Bit(c) && c%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstorezero [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2) && (off1+off2)%4 == 0
	// result: (MOVDstorezero [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) && (off1+off2)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstorezero [off1] {sym1} p:(MOVDaddr [off2] {sym2} x) mem)
	// cond: canMergeSym(sym1,sym2) && (x.Op != OpSB || p.Uses == 1) && (off1+off2)%4 == 0
	// result: (MOVDstorezero [off1+off2] {mergeSym(sym1,sym2)} x mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		x := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (x.Op != OpSB || p.Uses == 1) && (off1+off2)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHBRstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHBRstore {sym} ptr (MOVHreg x) mem)
	// result: (MOVHBRstore {sym} ptr x mem)
	for {
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVHBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHBRstore {sym} ptr (MOVHZreg x) mem)
	// result: (MOVHBRstore {sym} ptr x mem)
	for {
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVHBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHBRstore {sym} ptr (MOVWreg x) mem)
	// result: (MOVHBRstore {sym} ptr x mem)
	for {
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVHBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHBRstore {sym} ptr (MOVWZreg x) mem)
	// result: (MOVHBRstore {sym} ptr x mem)
	for {
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVHBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHZload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHZload [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (MOVHZload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHZload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHZload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHZload [0] {sym} p:(ADD ptr idx) mem)
	// cond: sym == nil && p.Uses == 1
	// result: (MOVHZloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		p := v_0
		if p.Op != OpPPC64ADD {
			break
		}
		idx := p.Args[1]
		ptr := p.Args[0]
		mem := v_1
		if !(sym == nil && p.Uses == 1) {
			break
		}
		v.reset(OpPPC64MOVHZloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHZloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHZloadidx ptr (MOVDconst [c]) mem)
	// cond: is16Bit(c)
	// result: (MOVHZload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHZloadidx (MOVDconst [c]) ptr mem)
	// cond: is16Bit(c)
	// result: (MOVHZload [c] ptr mem)
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHZreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVHZreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0xFFFF
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHZreg (SRWconst [c] (MOVBZreg x)))
	// result: (SRWconst [c] (MOVBZreg x))
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVHZreg (SRWconst [c] (MOVHZreg x)))
	// result: (SRWconst [c] (MOVHZreg x))
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVHZreg (SRWconst [c] x))
	// cond: sizeof(x.Type) <= 16
	// result: (SRWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(sizeof(x.Type) <= 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg (SRDconst [c] x))
	// cond: c>=48
	// result: (SRDconst [c] x)
	for {
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c >= 48) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg (SRWconst [c] x))
	// cond: c>=16
	// result: (SRWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c >= 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg y:(MOVHZreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHZreg y:(MOVBZreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHZreg y:(MOVHBRload _ _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVHBRload {
			break
		}
		_ = y.Args[1]
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHZreg y:(MOVHreg x))
	// result: (MOVHZreg x)
	for {
		y := v_0
		if y.Op != OpPPC64MOVHreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVHZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg x:(MOVBZload _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg x:(MOVBZloadidx _ _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVBZloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg x:(MOVHZload _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVHZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg x:(MOVHZloadidx _ _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVHZloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg x:(Arg <t>))
	// cond: (is8BitInt(t) || is16BitInt(t)) && !isSigned(t)
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !((is8BitInt(t) || is16BitInt(t)) && !isSigned(t)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg (MOVDconst [c]))
	// result: (MOVDconst [int64(uint16(c))])
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHload [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [0] {sym} p:(ADD ptr idx) mem)
	// cond: sym == nil && p.Uses == 1
	// result: (MOVHloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		p := v_0
		if p.Op != OpPPC64ADD {
			break
		}
		idx := p.Args[1]
		ptr := p.Args[0]
		mem := v_1
		if !(sym == nil && p.Uses == 1) {
			break
		}
		v.reset(OpPPC64MOVHloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHloadidx ptr (MOVDconst [c]) mem)
	// cond: is16Bit(c)
	// result: (MOVHload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHloadidx (MOVDconst [c]) ptr mem)
	// cond: is16Bit(c)
	// result: (MOVHload [c] ptr mem)
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVHreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0x7FFF
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0x7FFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHreg (SRAWconst [c] (MOVBreg x)))
	// result: (SRAWconst [c] (MOVBreg x))
	for {
		if v_0.Op != OpPPC64SRAWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVBreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVHreg (SRAWconst [c] (MOVHreg x)))
	// result: (SRAWconst [c] (MOVHreg x))
	for {
		if v_0.Op != OpPPC64SRAWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVHreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVHreg (SRAWconst [c] x))
	// cond: sizeof(x.Type) <= 16
	// result: (SRAWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRAWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(sizeof(x.Type) <= 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (SRDconst [c] x))
	// cond: c>48
	// result: (SRDconst [c] x)
	for {
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c > 48) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (SRDconst [c] x))
	// cond: c==48
	// result: (SRADconst [c] x)
	for {
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c == 48) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (SRWconst [c] x))
	// cond: c>16
	// result: (SRWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c > 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (SRWconst [c] x))
	// cond: c==16
	// result: (SRAWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c == 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg y:(MOVHreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVHreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHreg y:(MOVBreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVBreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHreg y:(MOVHZreg x))
	// result: (MOVHreg x)
	for {
		y := v_0
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHloadidx _ _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVHloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(Arg <t>))
	// cond: (is8BitInt(t) || is16BitInt(t)) && isSigned(t)
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !((is8BitInt(t) || is16BitInt(t)) && isSigned(t)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVDconst [c]))
	// result: (MOVDconst [int64(int16(c))])
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (MOVHstore [off1] {sym} (ADDconst [off2] x) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {sym} x val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVDconst [0]) mem)
	// result: (MOVHstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} p:(ADD ptr idx) val mem)
	// cond: off == 0 && sym == nil && p.Uses == 1
	// result: (MOVHstoreidx ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		p := v_0
		if p.Op != OpPPC64ADD {
			break
		}
		idx := p.Args[1]
		ptr := p.Args[0]
		val := v_1
		mem := v_2
		if !(off == 0 && sym == nil && p.Uses == 1) {
			break
		}
		v.reset(OpPPC64MOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHZreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWZreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [i1] {s} p (SRWconst w [16]) x0:(MOVHstore [i0] {s} p w mem))
	// cond: !config.BigEndian && x0.Uses == 1 && i1 == i0+2 && clobber(x0)
	// result: (MOVWstore [i0] {s} p w mem)
	for {
		i1 := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpPPC64SRWconst || v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x0 := v_2
		if x0.Op != OpPPC64MOVHstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		mem := x0.Args[2]
		if p != x0.Args[0] || w != x0.Args[1] || !(!config.BigEndian && x0.Uses == 1 && i1 == i0+2 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [i1] {s} p (SRDconst w [16]) x0:(MOVHstore [i0] {s} p w mem))
	// cond: !config.BigEndian && x0.Uses == 1 && i1 == i0+2 && clobber(x0)
	// result: (MOVWstore [i0] {s} p w mem)
	for {
		i1 := v.AuxInt
		s := v.Aux
		p := v_0
		if v_1.Op != OpPPC64SRDconst || v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x0 := v_2
		if x0.Op != OpPPC64MOVHstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		mem := x0.Args[2]
		if p != x0.Args[0] || w != x0.Args[1] || !(!config.BigEndian && x0.Uses == 1 && i1 == i0+2 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstoreidx ptr (MOVDconst [c]) val mem)
	// cond: is16Bit(c)
	// result: (MOVHstore [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v_2
		mem := v_3
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreidx (MOVDconst [c]) ptr val mem)
	// cond: is16Bit(c)
	// result: (MOVHstore [c] ptr val mem)
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		val := v_2
		mem := v_3
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreidx [off] {sym} ptr idx (MOVHreg x) mem)
	// result: (MOVHstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVHreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVHstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreidx [off] {sym} ptr idx (MOVHZreg x) mem)
	// result: (MOVHstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVHstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreidx [off] {sym} ptr idx (MOVWreg x) mem)
	// result: (MOVHstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVWreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVHstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreidx [off] {sym} ptr idx (MOVWZreg x) mem)
	// result: (MOVHstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVHstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstorezero [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstorezero [off1] {sym1} p:(MOVDaddr [off2] {sym2} x) mem)
	// cond: canMergeSym(sym1,sym2) && (x.Op != OpSB || p.Uses == 1)
	// result: (MOVHstorezero [off1+off2] {mergeSym(sym1,sym2)} x mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		x := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (x.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWBRstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWBRstore {sym} ptr (MOVWreg x) mem)
	// result: (MOVWBRstore {sym} ptr x mem)
	for {
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVWBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWBRstore {sym} ptr (MOVWZreg x) mem)
	// result: (MOVWBRstore {sym} ptr x mem)
	for {
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVWBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWZload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWZload [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (MOVWZload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWZload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWZload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWZload [0] {sym} p:(ADD ptr idx) mem)
	// cond: sym == nil && p.Uses == 1
	// result: (MOVWZloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		p := v_0
		if p.Op != OpPPC64ADD {
			break
		}
		idx := p.Args[1]
		ptr := p.Args[0]
		mem := v_1
		if !(sym == nil && p.Uses == 1) {
			break
		}
		v.reset(OpPPC64MOVWZloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWZloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWZloadidx ptr (MOVDconst [c]) mem)
	// cond: is16Bit(c)
	// result: (MOVWZload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWZloadidx (MOVDconst [c]) ptr mem)
	// cond: is16Bit(c)
	// result: (MOVWZload [c] ptr mem)
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWZreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVWZreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0xFFFFFFFF
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFFFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(AND (MOVDconst [c]) _))
	// cond: uint64(c) <= 0xFFFFFFFF
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64AND {
			break
		}
		_ = y.Args[1]
		y_0 := y.Args[0]
		y_1 := y.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, y_0, y_1 = _i0+1, y_1, y_0 {
			if y_0.Op != OpPPC64MOVDconst {
				continue
			}
			c := y_0.AuxInt
			if !(uint64(c) <= 0xFFFFFFFF) {
				continue
			}
			v.reset(OpCopy)
			v.Type = y.Type
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (MOVWZreg (SRWconst [c] (MOVBZreg x)))
	// result: (SRWconst [c] (MOVBZreg x))
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVWZreg (SRWconst [c] (MOVHZreg x)))
	// result: (SRWconst [c] (MOVHZreg x))
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVWZreg (SRWconst [c] (MOVWZreg x)))
	// result: (SRWconst [c] (MOVWZreg x))
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVWZreg (SRWconst [c] x))
	// cond: sizeof(x.Type) <= 32
	// result: (SRWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(sizeof(x.Type) <= 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg (SRDconst [c] x))
	// cond: c>=32
	// result: (SRDconst [c] x)
	for {
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c >= 32) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg y:(MOVWZreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVWZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(MOVHZreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(MOVBZreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(MOVHBRload _ _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVHBRload {
			break
		}
		_ = y.Args[1]
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(MOVWBRload _ _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVWBRload {
			break
		}
		_ = y.Args[1]
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(MOVWreg x))
	// result: (MOVWZreg x)
	for {
		y := v_0
		if y.Op != OpPPC64MOVWreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVWZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVBZload _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVBZloadidx _ _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVBZloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVHZload _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVHZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVHZloadidx _ _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVHZloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVWZload _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVWZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(MOVWZloadidx _ _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVWZloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg x:(Arg <t>))
	// cond: (is8BitInt(t) || is16BitInt(t) || is32BitInt(t)) && !isSigned(t)
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !((is8BitInt(t) || is16BitInt(t) || is32BitInt(t)) && !isSigned(t)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWZreg (MOVDconst [c]))
	// result: (MOVDconst [int64(uint32(c))])
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(uint32(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWload [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1) && (off1+off2)%4 == 0
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1) && (off1+off2)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2) && (off1+off2)%4 == 0
	// result: (MOVWload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1+off2) && (off1+off2)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [0] {sym} p:(ADD ptr idx) mem)
	// cond: sym == nil && p.Uses == 1
	// result: (MOVWloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		p := v_0
		if p.Op != OpPPC64ADD {
			break
		}
		idx := p.Args[1]
		ptr := p.Args[0]
		mem := v_1
		if !(sym == nil && p.Uses == 1) {
			break
		}
		v.reset(OpPPC64MOVWloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWloadidx ptr (MOVDconst [c]) mem)
	// cond: is16Bit(c) && c%4 == 0
	// result: (MOVWload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is16Bit(c) && c%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx (MOVDconst [c]) ptr mem)
	// cond: is16Bit(c) && c%4 == 0
	// result: (MOVWload [c] ptr mem)
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		if !(is16Bit(c) && c%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVWreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0xFFFF
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWreg y:(AND (MOVDconst [c]) _))
	// cond: uint64(c) <= 0x7FFFFFFF
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64AND {
			break
		}
		_ = y.Args[1]
		y_0 := y.Args[0]
		y_1 := y.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, y_0, y_1 = _i0+1, y_1, y_0 {
			if y_0.Op != OpPPC64MOVDconst {
				continue
			}
			c := y_0.AuxInt
			if !(uint64(c) <= 0x7FFFFFFF) {
				continue
			}
			v.reset(OpCopy)
			v.Type = y.Type
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (MOVWreg (SRAWconst [c] (MOVBreg x)))
	// result: (SRAWconst [c] (MOVBreg x))
	for {
		if v_0.Op != OpPPC64SRAWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVBreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVWreg (SRAWconst [c] (MOVHreg x)))
	// result: (SRAWconst [c] (MOVHreg x))
	for {
		if v_0.Op != OpPPC64SRAWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVHreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVWreg (SRAWconst [c] (MOVWreg x)))
	// result: (SRAWconst [c] (MOVWreg x))
	for {
		if v_0.Op != OpPPC64SRAWconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64MOVWreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVWreg (SRAWconst [c] x))
	// cond: sizeof(x.Type) <= 32
	// result: (SRAWconst [c] x)
	for {
		if v_0.Op != OpPPC64SRAWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(sizeof(x.Type) <= 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (SRDconst [c] x))
	// cond: c>32
	// result: (SRDconst [c] x)
	for {
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c > 32) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (SRDconst [c] x))
	// cond: c==32
	// result: (SRADconst [c] x)
	for {
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c == 32) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg y:(MOVWreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVWreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWreg y:(MOVHreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVHreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWreg y:(MOVBreg _))
	// result: y
	for {
		y := v_0
		if y.Op != OpPPC64MOVBreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWreg y:(MOVWZreg x))
	// result: (MOVWreg x)
	for {
		y := v_0
		if y.Op != OpPPC64MOVWZreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHload _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHloadidx _ _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVHloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVWload _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVWload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVWloadidx _ _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpPPC64MOVWloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(Arg <t>))
	// cond: (is8BitInt(t) || is16BitInt(t) || is32BitInt(t)) && isSigned(t)
	// result: x
	for {
		x := v_0
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !((is8BitInt(t) || is16BitInt(t) || is32BitInt(t)) && isSigned(t)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (MOVDconst [c]))
	// result: (MOVDconst [int64(int32(c))])
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(int32(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstore [off1] {sym} (ADDconst [off2] x) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} x val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} p:(MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && (ptr.Op != OpSB || p.Uses == 1)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVDconst [0]) mem)
	// result: (MOVWstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} p:(ADD ptr idx) val mem)
	// cond: off == 0 && sym == nil && p.Uses == 1
	// result: (MOVWstoreidx ptr idx val mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		p := v_0
		if p.Op != OpPPC64ADD {
			break
		}
		idx := p.Args[1]
		ptr := p.Args[0]
		val := v_1
		mem := v_2
		if !(off == 0 && sym == nil && p.Uses == 1) {
			break
		}
		v.reset(OpPPC64MOVWstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWZreg x) mem)
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstoreidx ptr (MOVDconst [c]) val mem)
	// cond: is16Bit(c)
	// result: (MOVWstore [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v_2
		mem := v_3
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx (MOVDconst [c]) ptr val mem)
	// cond: is16Bit(c)
	// result: (MOVWstore [c] ptr val mem)
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		val := v_2
		mem := v_3
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx [off] {sym} ptr idx (MOVWreg x) mem)
	// result: (MOVWstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVWreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVWstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx [off] {sym} ptr idx (MOVWZreg x) mem)
	// result: (MOVWstoreidx [off] {sym} ptr idx x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		idx := v_1
		if v_2.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_2.Args[0]
		mem := v_3
		v.reset(OpPPC64MOVWstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstorezero [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstorezero [off1] {sym1} p:(MOVDaddr [off2] {sym2} x) mem)
	// cond: canMergeSym(sym1,sym2) && (x.Op != OpSB || p.Uses == 1)
	// result: (MOVWstorezero [off1+off2] {mergeSym(sym1,sym2)} x mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		p := v_0
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		x := p.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && (x.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MTVSRD(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MTVSRD (MOVDconst [c]))
	// result: (FMOVDconst [c])
	for {
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = c
		return true
	}
	// match: (MTVSRD x:(MOVDload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (FMOVDload [off] {sym} ptr mem)
	for {
		x := v_0
		if x.Op != OpPPC64MOVDload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpPPC64FMOVDload, typ.Float64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MaskIfNotCarry(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MaskIfNotCarry (ADDconstForCarry [c] (ANDconst [d] _)))
	// cond: c < 0 && d > 0 && c + d < 0
	// result: (MOVDconst [-1])
	for {
		if v_0.Op != OpPPC64ADDconstForCarry {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64ANDconst {
			break
		}
		d := v_0_0.AuxInt
		if !(c < 0 && d > 0 && c+d < 0) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = -1
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64NotEqual(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NotEqual (FlagEQ))
	// result: (MOVDconst [0])
	for {
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (NotEqual (FlagLT))
	// result: (MOVDconst [1])
	for {
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (FlagGT))
	// result: (MOVDconst [1])
	for {
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (InvertFlags x))
	// result: (NotEqual x)
	for {
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64NotEqual)
		v.AddArg(x)
		return true
	}
	// match: (NotEqual cmp)
	// result: (ISELB [6] (MOVDconst [1]) cmp)
	for {
		cmp := v_0
		v.reset(OpPPC64ISELB)
		v.AuxInt = 6
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(cmp)
		return true
	}
}
func rewriteValuePPC64_OpPPC64OR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: ( OR (SLDconst x [c]) (SRDconst x [d]))
	// cond: d == 64-c
	// result: (ROTLconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLDconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpPPC64SRDconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 64-c) {
				continue
			}
			v.reset(OpPPC64ROTLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: ( OR (SLWconst x [c]) (SRWconst x [d]))
	// cond: d == 32-c
	// result: (ROTLWconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLWconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpPPC64SRWconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 32-c) {
				continue
			}
			v.reset(OpPPC64ROTLWconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: ( OR (SLD x (ANDconst <typ.Int64> [63] y)) (SRD x (SUB <typ.UInt> (MOVDconst [64]) (ANDconst <typ.UInt> [63] y))))
	// result: (ROTL x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLD {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpPPC64ANDconst || v_0_1.Type != typ.Int64 || v_0_1.AuxInt != 63 {
				continue
			}
			y := v_0_1.Args[0]
			if v_1.Op != OpPPC64SRD {
				continue
			}
			_ = v_1.Args[1]
			if x != v_1.Args[0] {
				continue
			}
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpPPC64SUB || v_1_1.Type != typ.UInt {
				continue
			}
			_ = v_1_1.Args[1]
			v_1_1_0 := v_1_1.Args[0]
			if v_1_1_0.Op != OpPPC64MOVDconst || v_1_1_0.AuxInt != 64 {
				continue
			}
			v_1_1_1 := v_1_1.Args[1]
			if v_1_1_1.Op != OpPPC64ANDconst || v_1_1_1.Type != typ.UInt || v_1_1_1.AuxInt != 63 || y != v_1_1_1.Args[0] {
				continue
			}
			v.reset(OpPPC64ROTL)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: ( OR (SLW x (ANDconst <typ.Int32> [31] y)) (SRW x (SUB <typ.UInt> (MOVDconst [32]) (ANDconst <typ.UInt> [31] y))))
	// result: (ROTLW x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLW {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpPPC64ANDconst || v_0_1.Type != typ.Int32 || v_0_1.AuxInt != 31 {
				continue
			}
			y := v_0_1.Args[0]
			if v_1.Op != OpPPC64SRW {
				continue
			}
			_ = v_1.Args[1]
			if x != v_1.Args[0] {
				continue
			}
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpPPC64SUB || v_1_1.Type != typ.UInt {
				continue
			}
			_ = v_1_1.Args[1]
			v_1_1_0 := v_1_1.Args[0]
			if v_1_1_0.Op != OpPPC64MOVDconst || v_1_1_0.AuxInt != 32 {
				continue
			}
			v_1_1_1 := v_1_1.Args[1]
			if v_1_1_1.Op != OpPPC64ANDconst || v_1_1_1.Type != typ.UInt || v_1_1_1.AuxInt != 31 || y != v_1_1_1.Args[0] {
				continue
			}
			v.reset(OpPPC64ROTLW)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (OR (MOVDconst [c]) (MOVDconst [d]))
	// result: (MOVDconst [c|d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64MOVDconst {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpPPC64MOVDconst {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpPPC64MOVDconst)
			v.AuxInt = c | d
			return true
		}
		break
	}
	// match: (OR x (MOVDconst [c]))
	// cond: isU32Bit(c)
	// result: (ORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpPPC64MOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(isU32Bit(c)) {
				continue
			}
			v.reset(OpPPC64ORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (OR <t> x0:(MOVBZload [i0] {s} p mem) o1:(SLWconst x1:(MOVBZload [i1] {s} p mem) [8]))
	// cond: !config.BigEndian && i1 == i0+1 && x0.Uses ==1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)
	// result: @mergePoint(b,x0,x1) (MOVHZload <t> {s} [i0] p mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x0 := v_0
			if x0.Op != OpPPC64MOVBZload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			o1 := v_1
			if o1.Op != OpPPC64SLWconst || o1.AuxInt != 8 {
				continue
			}
			x1 := o1.Args[0]
			if x1.Op != OpPPC64MOVBZload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpPPC64MOVHZload, t)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (OR <t> x0:(MOVBZload [i0] {s} p mem) o1:(SLDconst x1:(MOVBZload [i1] {s} p mem) [8]))
	// cond: !config.BigEndian && i1 == i0+1 && x0.Uses ==1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)
	// result: @mergePoint(b,x0,x1) (MOVHZload <t> {s} [i0] p mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x0 := v_0
			if x0.Op != OpPPC64MOVBZload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			o1 := v_1
			if o1.Op != OpPPC64SLDconst || o1.AuxInt != 8 {
				continue
			}
			x1 := o1.Args[0]
			if x1.Op != OpPPC64MOVBZload {
				continue
			}
			i1 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpPPC64MOVHZload, t)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = i0
			v0.Aux = s
			v0.AddArg(p)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (OR <t> x0:(MOVBZload [i1] {s} p mem) o1:(SLWconst x1:(MOVBZload [i0] {s} p mem) [8]))
	// cond: !config.BigEndian && i1 == i0+1 && x0.Uses ==1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)
	// result: @mergePoint(b,x0,x1) (MOVHBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x0 := v_0
			if x0.Op != OpPPC64MOVBZload {
				continue
			}
			i1 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			o1 := v_1
			if o1.Op != OpPPC64SLWconst || o1.AuxInt != 8 {
				continue
			}
			x1 := o1.Args[0]
			if x1.Op != OpPPC64MOVBZload {
				continue
			}
			i0 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpPPC64MOVHBRload, t)
			v.reset(OpCopy)
			v.AddArg(v0)
			v1 := b.NewValue0(x1.Pos, OpPPC64MOVDaddr, typ.Uintptr)
			v1.AuxInt = i0
			v1.Aux = s
			v1.AddArg(p)
			v0.AddArg(v1)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (OR <t> x0:(MOVBZload [i1] {s} p mem) o1:(SLDconst x1:(MOVBZload [i0] {s} p mem) [8]))
	// cond: !config.BigEndian && i1 == i0+1 && x0.Uses ==1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)
	// result: @mergePoint(b,x0,x1) (MOVHBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x0 := v_0
			if x0.Op != OpPPC64MOVBZload {
				continue
			}
			i1 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			o1 := v_1
			if o1.Op != OpPPC64SLDconst || o1.AuxInt != 8 {
				continue
			}
			x1 := o1.Args[0]
			if x1.Op != OpPPC64MOVBZload {
				continue
			}
			i0 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpPPC64MOVHBRload, t)
			v.reset(OpCopy)
			v.AddArg(v0)
			v1 := b.NewValue0(x1.Pos, OpPPC64MOVDaddr, typ.Uintptr)
			v1.AuxInt = i0
			v1.Aux = s
			v1.AddArg(p)
			v0.AddArg(v1)
			v0.AddArg(mem)
			return true
		}
		break
	}
	// match: (OR <t> s0:(SLWconst x0:(MOVBZload [i1] {s} p mem) [n1]) s1:(SLWconst x1:(MOVBZload [i0] {s} p mem) [n2]))
	// cond: !config.BigEndian && i1 == i0+1 && n1%8 == 0 && n2 == n1+8 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1)
	// result: @mergePoint(b,x0,x1) (SLDconst <t> (MOVHBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem) [n1])
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s0 := v_0
			if s0.Op != OpPPC64SLWconst {
				continue
			}
			n1 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpPPC64MOVBZload {
				continue
			}
			i1 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			s1 := v_1
			if s1.Op != OpPPC64SLWconst {
				continue
			}
			n2 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpPPC64MOVBZload {
				continue
			}
			i0 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(!config.BigEndian && i1 == i0+1 && n1%8 == 0 && n2 == n1+8 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpPPC64SLDconst, t)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = n1
			v1 := b.NewValue0(x1.Pos, OpPPC64MOVHBRload, t)
			v2 := b.NewValue0(x1.Pos, OpPPC64MOVDaddr, typ.Uintptr)
			v2.AuxInt = i0
			v2.Aux = s
			v2.AddArg(p)
			v1.AddArg(v2)
			v1.AddArg(mem)
			v0.AddArg(v1)
			return true
		}
		break
	}
	// match: (OR <t> s0:(SLDconst x0:(MOVBZload [i1] {s} p mem) [n1]) s1:(SLDconst x1:(MOVBZload [i0] {s} p mem) [n2]))
	// cond: !config.BigEndian && i1 == i0+1 && n1%8 == 0 && n2 == n1+8 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1)
	// result: @mergePoint(b,x0,x1) (SLDconst <t> (MOVHBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem) [n1])
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s0 := v_0
			if s0.Op != OpPPC64SLDconst {
				continue
			}
			n1 := s0.AuxInt
			x0 := s0.Args[0]
			if x0.Op != OpPPC64MOVBZload {
				continue
			}
			i1 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			s1 := v_1
			if s1.Op != OpPPC64SLDconst {
				continue
			}
			n2 := s1.AuxInt
			x1 := s1.Args[0]
			if x1.Op != OpPPC64MOVBZload {
				continue
			}
			i0 := x1.AuxInt
			if x1.Aux != s {
				continue
			}
			_ = x1.Args[1]
			if p != x1.Args[0] || mem != x1.Args[1] || !(!config.BigEndian && i1 == i0+1 && n1%8 == 0 && n2 == n1+8 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1)) {
				continue
			}
			b = mergePoint(b, x0, x1)
			v0 := b.NewValue0(x1.Pos, OpPPC64SLDconst, t)
			v.reset(OpCopy)
			v.AddArg(v0)
			v0.AuxInt = n1
			v1 := b.NewValue0(x1.Pos, OpPPC64MOVHBRload, t)
			v2 := b.NewValue0(x1.Pos, OpPPC64MOVDaddr, typ.Uintptr)
			v2.AuxInt = i0
			v2.Aux = s
			v2.AddArg(p)
			v1.AddArg(v2)
			v1.AddArg(mem)
			v0.AddArg(v1)
			return true
		}
		break
	}
	// match: (OR <t> s1:(SLWconst x2:(MOVBZload [i3] {s} p mem) [24]) o0:(OR <t> s0:(SLWconst x1:(MOVBZload [i2] {s} p mem) [16]) x0:(MOVHZload [i0] {s} p mem)))
	// cond: !config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses ==1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)
	// result: @mergePoint(b,x0,x1,x2) (MOVWZload <t> {s} [i0] p mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s1 := v_0
			if s1.Op != OpPPC64SLWconst || s1.AuxInt != 24 {
				continue
			}
			x2 := s1.Args[0]
			if x2.Op != OpPPC64MOVBZload {
				continue
			}
			i3 := x2.AuxInt
			s := x2.Aux
			mem := x2.Args[1]
			p := x2.Args[0]
			o0 := v_1
			if o0.Op != OpPPC64OR || o0.Type != t {
				continue
			}
			_ = o0.Args[1]
			o0_0 := o0.Args[0]
			o0_1 := o0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o0_0, o0_1 = _i1+1, o0_1, o0_0 {
				s0 := o0_0
				if s0.Op != OpPPC64SLWconst || s0.AuxInt != 16 {
					continue
				}
				x1 := s0.Args[0]
				if x1.Op != OpPPC64MOVBZload {
					continue
				}
				i2 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				x0 := o0_1
				if x0.Op != OpPPC64MOVHZload {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[1]
				if p != x0.Args[0] || mem != x0.Args[1] || !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
					continue
				}
				b = mergePoint(b, x0, x1, x2)
				v0 := b.NewValue0(x0.Pos, OpPPC64MOVWZload, t)
				v.reset(OpCopy)
				v.AddArg(v0)
				v0.AuxInt = i0
				v0.Aux = s
				v0.AddArg(p)
				v0.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (OR <t> s1:(SLDconst x2:(MOVBZload [i3] {s} p mem) [24]) o0:(OR <t> s0:(SLDconst x1:(MOVBZload [i2] {s} p mem) [16]) x0:(MOVHZload [i0] {s} p mem)))
	// cond: !config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses ==1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)
	// result: @mergePoint(b,x0,x1,x2) (MOVWZload <t> {s} [i0] p mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s1 := v_0
			if s1.Op != OpPPC64SLDconst || s1.AuxInt != 24 {
				continue
			}
			x2 := s1.Args[0]
			if x2.Op != OpPPC64MOVBZload {
				continue
			}
			i3 := x2.AuxInt
			s := x2.Aux
			mem := x2.Args[1]
			p := x2.Args[0]
			o0 := v_1
			if o0.Op != OpPPC64OR || o0.Type != t {
				continue
			}
			_ = o0.Args[1]
			o0_0 := o0.Args[0]
			o0_1 := o0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o0_0, o0_1 = _i1+1, o0_1, o0_0 {
				s0 := o0_0
				if s0.Op != OpPPC64SLDconst || s0.AuxInt != 16 {
					continue
				}
				x1 := s0.Args[0]
				if x1.Op != OpPPC64MOVBZload {
					continue
				}
				i2 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				x0 := o0_1
				if x0.Op != OpPPC64MOVHZload {
					continue
				}
				i0 := x0.AuxInt
				if x0.Aux != s {
					continue
				}
				_ = x0.Args[1]
				if p != x0.Args[0] || mem != x0.Args[1] || !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
					continue
				}
				b = mergePoint(b, x0, x1, x2)
				v0 := b.NewValue0(x0.Pos, OpPPC64MOVWZload, t)
				v.reset(OpCopy)
				v.AddArg(v0)
				v0.AuxInt = i0
				v0.Aux = s
				v0.AddArg(p)
				v0.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (OR <t> s1:(SLWconst x2:(MOVBZload [i0] {s} p mem) [24]) o0:(OR <t> s0:(SLWconst x1:(MOVBZload [i1] {s} p mem) [16]) x0:(MOVHBRload <t> (MOVDaddr <typ.Uintptr> [i2] {s} p) mem)))
	// cond: !config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)
	// result: @mergePoint(b,x0,x1,x2) (MOVWBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s1 := v_0
			if s1.Op != OpPPC64SLWconst || s1.AuxInt != 24 {
				continue
			}
			x2 := s1.Args[0]
			if x2.Op != OpPPC64MOVBZload {
				continue
			}
			i0 := x2.AuxInt
			s := x2.Aux
			mem := x2.Args[1]
			p := x2.Args[0]
			o0 := v_1
			if o0.Op != OpPPC64OR || o0.Type != t {
				continue
			}
			_ = o0.Args[1]
			o0_0 := o0.Args[0]
			o0_1 := o0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o0_0, o0_1 = _i1+1, o0_1, o0_0 {
				s0 := o0_0
				if s0.Op != OpPPC64SLWconst || s0.AuxInt != 16 {
					continue
				}
				x1 := s0.Args[0]
				if x1.Op != OpPPC64MOVBZload {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				x0 := o0_1
				if x0.Op != OpPPC64MOVHBRload || x0.Type != t {
					continue
				}
				_ = x0.Args[1]
				x0_0 := x0.Args[0]
				if x0_0.Op != OpPPC64MOVDaddr || x0_0.Type != typ.Uintptr {
					continue
				}
				i2 := x0_0.AuxInt
				if x0_0.Aux != s || p != x0_0.Args[0] || mem != x0.Args[1] || !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
					continue
				}
				b = mergePoint(b, x0, x1, x2)
				v0 := b.NewValue0(x0.Pos, OpPPC64MOVWBRload, t)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x0.Pos, OpPPC64MOVDaddr, typ.Uintptr)
				v1.AuxInt = i0
				v1.Aux = s
				v1.AddArg(p)
				v0.AddArg(v1)
				v0.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (OR <t> s1:(SLDconst x2:(MOVBZload [i0] {s} p mem) [24]) o0:(OR <t> s0:(SLDconst x1:(MOVBZload [i1] {s} p mem) [16]) x0:(MOVHBRload <t> (MOVDaddr <typ.Uintptr> [i2] {s} p) mem)))
	// cond: !config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)
	// result: @mergePoint(b,x0,x1,x2) (MOVWBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s1 := v_0
			if s1.Op != OpPPC64SLDconst || s1.AuxInt != 24 {
				continue
			}
			x2 := s1.Args[0]
			if x2.Op != OpPPC64MOVBZload {
				continue
			}
			i0 := x2.AuxInt
			s := x2.Aux
			mem := x2.Args[1]
			p := x2.Args[0]
			o0 := v_1
			if o0.Op != OpPPC64OR || o0.Type != t {
				continue
			}
			_ = o0.Args[1]
			o0_0 := o0.Args[0]
			o0_1 := o0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o0_0, o0_1 = _i1+1, o0_1, o0_0 {
				s0 := o0_0
				if s0.Op != OpPPC64SLDconst || s0.AuxInt != 16 {
					continue
				}
				x1 := s0.Args[0]
				if x1.Op != OpPPC64MOVBZload {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				x0 := o0_1
				if x0.Op != OpPPC64MOVHBRload || x0.Type != t {
					continue
				}
				_ = x0.Args[1]
				x0_0 := x0.Args[0]
				if x0_0.Op != OpPPC64MOVDaddr || x0_0.Type != typ.Uintptr {
					continue
				}
				i2 := x0_0.AuxInt
				if x0_0.Aux != s || p != x0_0.Args[0] || mem != x0.Args[1] || !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
					continue
				}
				b = mergePoint(b, x0, x1, x2)
				v0 := b.NewValue0(x0.Pos, OpPPC64MOVWBRload, t)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x0.Pos, OpPPC64MOVDaddr, typ.Uintptr)
				v1.AuxInt = i0
				v1.Aux = s
				v1.AddArg(p)
				v0.AddArg(v1)
				v0.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (OR <t> x0:(MOVBZload [i3] {s} p mem) o0:(OR <t> s0:(SLWconst x1:(MOVBZload [i2] {s} p mem) [8]) s1:(SLWconst x2:(MOVHBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem) [16])))
	// cond: !config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)
	// result: @mergePoint(b,x0,x1,x2) (MOVWBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x0 := v_0
			if x0.Op != OpPPC64MOVBZload {
				continue
			}
			i3 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			o0 := v_1
			if o0.Op != OpPPC64OR || o0.Type != t {
				continue
			}
			_ = o0.Args[1]
			o0_0 := o0.Args[0]
			o0_1 := o0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o0_0, o0_1 = _i1+1, o0_1, o0_0 {
				s0 := o0_0
				if s0.Op != OpPPC64SLWconst || s0.AuxInt != 8 {
					continue
				}
				x1 := s0.Args[0]
				if x1.Op != OpPPC64MOVBZload {
					continue
				}
				i2 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				s1 := o0_1
				if s1.Op != OpPPC64SLWconst || s1.AuxInt != 16 {
					continue
				}
				x2 := s1.Args[0]
				if x2.Op != OpPPC64MOVHBRload || x2.Type != t {
					continue
				}
				_ = x2.Args[1]
				x2_0 := x2.Args[0]
				if x2_0.Op != OpPPC64MOVDaddr || x2_0.Type != typ.Uintptr {
					continue
				}
				i0 := x2_0.AuxInt
				if x2_0.Aux != s || p != x2_0.Args[0] || mem != x2.Args[1] || !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
					continue
				}
				b = mergePoint(b, x0, x1, x2)
				v0 := b.NewValue0(x2.Pos, OpPPC64MOVWBRload, t)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x2.Pos, OpPPC64MOVDaddr, typ.Uintptr)
				v1.AuxInt = i0
				v1.Aux = s
				v1.AddArg(p)
				v0.AddArg(v1)
				v0.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (OR <t> x0:(MOVBZload [i3] {s} p mem) o0:(OR <t> s0:(SLDconst x1:(MOVBZload [i2] {s} p mem) [8]) s1:(SLDconst x2:(MOVHBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem) [16])))
	// cond: !config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)
	// result: @mergePoint(b,x0,x1,x2) (MOVWBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x0 := v_0
			if x0.Op != OpPPC64MOVBZload {
				continue
			}
			i3 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			o0 := v_1
			if o0.Op != OpPPC64OR || o0.Type != t {
				continue
			}
			_ = o0.Args[1]
			o0_0 := o0.Args[0]
			o0_1 := o0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o0_0, o0_1 = _i1+1, o0_1, o0_0 {
				s0 := o0_0
				if s0.Op != OpPPC64SLDconst || s0.AuxInt != 8 {
					continue
				}
				x1 := s0.Args[0]
				if x1.Op != OpPPC64MOVBZload {
					continue
				}
				i2 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				s1 := o0_1
				if s1.Op != OpPPC64SLDconst || s1.AuxInt != 16 {
					continue
				}
				x2 := s1.Args[0]
				if x2.Op != OpPPC64MOVHBRload || x2.Type != t {
					continue
				}
				_ = x2.Args[1]
				x2_0 := x2.Args[0]
				if x2_0.Op != OpPPC64MOVDaddr || x2_0.Type != typ.Uintptr {
					continue
				}
				i0 := x2_0.AuxInt
				if x2_0.Aux != s || p != x2_0.Args[0] || mem != x2.Args[1] || !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
					continue
				}
				b = mergePoint(b, x0, x1, x2)
				v0 := b.NewValue0(x2.Pos, OpPPC64MOVWBRload, t)
				v.reset(OpCopy)
				v.AddArg(v0)
				v1 := b.NewValue0(x2.Pos, OpPPC64MOVDaddr, typ.Uintptr)
				v1.AuxInt = i0
				v1.Aux = s
				v1.AddArg(p)
				v0.AddArg(v1)
				v0.AddArg(mem)
				return true
			}
		}
		break
	}
	// match: (OR <t> s2:(SLDconst x2:(MOVBZload [i3] {s} p mem) [32]) o0:(OR <t> s1:(SLDconst x1:(MOVBZload [i2] {s} p mem) [40]) s0:(SLDconst x0:(MOVHBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem) [48])))
	// cond: !config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)
	// result: @mergePoint(b,x0,x1,x2) (SLDconst <t> (MOVWBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem) [32])
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s2 := v_0
			if s2.Op != OpPPC64SLDconst || s2.AuxInt != 32 {
				continue
			}
			x2 := s2.Args[0]
			if x2.Op != OpPPC64MOVBZload {
				continue
			}
			i3 := x2.AuxInt
			s := x2.Aux
			mem := x2.Args[1]
			p := x2.Args[0]
			o0 := v_1
			if o0.Op != OpPPC64OR || o0.Type != t {
				continue
			}
			_ = o0.Args[1]
			o0_0 := o0.Args[0]
			o0_1 := o0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o0_0, o0_1 = _i1+1, o0_1, o0_0 {
				s1 := o0_0
				if s1.Op != OpPPC64SLDconst || s1.AuxInt != 40 {
					continue
				}
				x1 := s1.Args[0]
				if x1.Op != OpPPC64MOVBZload {
					continue
				}
				i2 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				s0 := o0_1
				if s0.Op != OpPPC64SLDconst || s0.AuxInt != 48 {
					continue
				}
				x0 := s0.Args[0]
				if x0.Op != OpPPC64MOVHBRload || x0.Type != t {
					continue
				}
				_ = x0.Args[1]
				x0_0 := x0.Args[0]
				if x0_0.Op != OpPPC64MOVDaddr || x0_0.Type != typ.Uintptr {
					continue
				}
				i0 := x0_0.AuxInt
				if x0_0.Aux != s || p != x0_0.Args[0] || mem != x0.Args[1] || !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)) {
					continue
				}
				b = mergePoint(b, x0, x1, x2)
				v0 := b.NewValue0(x0.Pos, OpPPC64SLDconst, t)
				v.reset(OpCopy)
				v.AddArg(v0)
				v0.AuxInt = 32
				v1 := b.NewValue0(x0.Pos, OpPPC64MOVWBRload, t)
				v2 := b.NewValue0(x0.Pos, OpPPC64MOVDaddr, typ.Uintptr)
				v2.AuxInt = i0
				v2.Aux = s
				v2.AddArg(p)
				v1.AddArg(v2)
				v1.AddArg(mem)
				v0.AddArg(v1)
				return true
			}
		}
		break
	}
	// match: (OR <t> s2:(SLDconst x2:(MOVBZload [i0] {s} p mem) [56]) o0:(OR <t> s1:(SLDconst x1:(MOVBZload [i1] {s} p mem) [48]) s0:(SLDconst x0:(MOVHBRload <t> (MOVDaddr <typ.Uintptr> [i2] {s} p) mem) [32])))
	// cond: !config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)
	// result: @mergePoint(b,x0,x1,x2) (SLDconst <t> (MOVWBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem) [32])
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s2 := v_0
			if s2.Op != OpPPC64SLDconst || s2.AuxInt != 56 {
				continue
			}
			x2 := s2.Args[0]
			if x2.Op != OpPPC64MOVBZload {
				continue
			}
			i0 := x2.AuxInt
			s := x2.Aux
			mem := x2.Args[1]
			p := x2.Args[0]
			o0 := v_1
			if o0.Op != OpPPC64OR || o0.Type != t {
				continue
			}
			_ = o0.Args[1]
			o0_0 := o0.Args[0]
			o0_1 := o0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o0_0, o0_1 = _i1+1, o0_1, o0_0 {
				s1 := o0_0
				if s1.Op != OpPPC64SLDconst || s1.AuxInt != 48 {
					continue
				}
				x1 := s1.Args[0]
				if x1.Op != OpPPC64MOVBZload {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				s0 := o0_1
				if s0.Op != OpPPC64SLDconst || s0.AuxInt != 32 {
					continue
				}
				x0 := s0.Args[0]
				if x0.Op != OpPPC64MOVHBRload || x0.Type != t {
					continue
				}
				_ = x0.Args[1]
				x0_0 := x0.Args[0]
				if x0_0.Op != OpPPC64MOVDaddr || x0_0.Type != typ.Uintptr {
					continue
				}
				i2 := x0_0.AuxInt
				if x0_0.Aux != s || p != x0_0.Args[0] || mem != x0.Args[1] || !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)) {
					continue
				}
				b = mergePoint(b, x0, x1, x2)
				v0 := b.NewValue0(x0.Pos, OpPPC64SLDconst, t)
				v.reset(OpCopy)
				v.AddArg(v0)
				v0.AuxInt = 32
				v1 := b.NewValue0(x0.Pos, OpPPC64MOVWBRload, t)
				v2 := b.NewValue0(x0.Pos, OpPPC64MOVDaddr, typ.Uintptr)
				v2.AuxInt = i0
				v2.Aux = s
				v2.AddArg(p)
				v1.AddArg(v2)
				v1.AddArg(mem)
				v0.AddArg(v1)
				return true
			}
		}
		break
	}
	// match: (OR <t> s6:(SLDconst x7:(MOVBZload [i7] {s} p mem) [56]) o5:(OR <t> s5:(SLDconst x6:(MOVBZload [i6] {s} p mem) [48]) o4:(OR <t> s4:(SLDconst x5:(MOVBZload [i5] {s} p mem) [40]) o3:(OR <t> s3:(SLDconst x4:(MOVBZload [i4] {s} p mem) [32]) x0:(MOVWZload {s} [i0] p mem)))))
	// cond: !config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses ==1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber (s6) && clobber(o3) && clobber(o4) && clobber(o5)
	// result: @mergePoint(b,x0,x4,x5,x6,x7) (MOVDload <t> {s} [i0] p mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s6 := v_0
			if s6.Op != OpPPC64SLDconst || s6.AuxInt != 56 {
				continue
			}
			x7 := s6.Args[0]
			if x7.Op != OpPPC64MOVBZload {
				continue
			}
			i7 := x7.AuxInt
			s := x7.Aux
			mem := x7.Args[1]
			p := x7.Args[0]
			o5 := v_1
			if o5.Op != OpPPC64OR || o5.Type != t {
				continue
			}
			_ = o5.Args[1]
			o5_0 := o5.Args[0]
			o5_1 := o5.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o5_0, o5_1 = _i1+1, o5_1, o5_0 {
				s5 := o5_0
				if s5.Op != OpPPC64SLDconst || s5.AuxInt != 48 {
					continue
				}
				x6 := s5.Args[0]
				if x6.Op != OpPPC64MOVBZload {
					continue
				}
				i6 := x6.AuxInt
				if x6.Aux != s {
					continue
				}
				_ = x6.Args[1]
				if p != x6.Args[0] || mem != x6.Args[1] {
					continue
				}
				o4 := o5_1
				if o4.Op != OpPPC64OR || o4.Type != t {
					continue
				}
				_ = o4.Args[1]
				o4_0 := o4.Args[0]
				o4_1 := o4.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, o4_0, o4_1 = _i2+1, o4_1, o4_0 {
					s4 := o4_0
					if s4.Op != OpPPC64SLDconst || s4.AuxInt != 40 {
						continue
					}
					x5 := s4.Args[0]
					if x5.Op != OpPPC64MOVBZload {
						continue
					}
					i5 := x5.AuxInt
					if x5.Aux != s {
						continue
					}
					_ = x5.Args[1]
					if p != x5.Args[0] || mem != x5.Args[1] {
						continue
					}
					o3 := o4_1
					if o3.Op != OpPPC64OR || o3.Type != t {
						continue
					}
					_ = o3.Args[1]
					o3_0 := o3.Args[0]
					o3_1 := o3.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, o3_0, o3_1 = _i3+1, o3_1, o3_0 {
						s3 := o3_0
						if s3.Op != OpPPC64SLDconst || s3.AuxInt != 32 {
							continue
						}
						x4 := s3.Args[0]
						if x4.Op != OpPPC64MOVBZload {
							continue
						}
						i4 := x4.AuxInt
						if x4.Aux != s {
							continue
						}
						_ = x4.Args[1]
						if p != x4.Args[0] || mem != x4.Args[1] {
							continue
						}
						x0 := o3_1
						if x0.Op != OpPPC64MOVWZload {
							continue
						}
						i0 := x0.AuxInt
						if x0.Aux != s {
							continue
						}
						_ = x0.Args[1]
						if p != x0.Args[0] || mem != x0.Args[1] || !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
							continue
						}
						b = mergePoint(b, x0, x4, x5, x6, x7)
						v0 := b.NewValue0(x0.Pos, OpPPC64MOVDload, t)
						v.reset(OpCopy)
						v.AddArg(v0)
						v0.AuxInt = i0
						v0.Aux = s
						v0.AddArg(p)
						v0.AddArg(mem)
						return true
					}
				}
			}
		}
		break
	}
	// match: (OR <t> s0:(SLDconst x0:(MOVBZload [i0] {s} p mem) [56]) o0:(OR <t> s1:(SLDconst x1:(MOVBZload [i1] {s} p mem) [48]) o1:(OR <t> s2:(SLDconst x2:(MOVBZload [i2] {s} p mem) [40]) o2:(OR <t> s3:(SLDconst x3:(MOVBZload [i3] {s} p mem) [32]) x4:(MOVWBRload <t> (MOVDaddr <typ.Uintptr> [i4] p) mem)))))
	// cond: !config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)
	// result: @mergePoint(b,x0,x1,x2,x3,x4) (MOVDBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s0 := v_0
			if s0.Op != OpPPC64SLDconst || s0.AuxInt != 56 {
				continue
			}
			x0 := s0.Args[0]
			if x0.Op != OpPPC64MOVBZload {
				continue
			}
			i0 := x0.AuxInt
			s := x0.Aux
			mem := x0.Args[1]
			p := x0.Args[0]
			o0 := v_1
			if o0.Op != OpPPC64OR || o0.Type != t {
				continue
			}
			_ = o0.Args[1]
			o0_0 := o0.Args[0]
			o0_1 := o0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o0_0, o0_1 = _i1+1, o0_1, o0_0 {
				s1 := o0_0
				if s1.Op != OpPPC64SLDconst || s1.AuxInt != 48 {
					continue
				}
				x1 := s1.Args[0]
				if x1.Op != OpPPC64MOVBZload {
					continue
				}
				i1 := x1.AuxInt
				if x1.Aux != s {
					continue
				}
				_ = x1.Args[1]
				if p != x1.Args[0] || mem != x1.Args[1] {
					continue
				}
				o1 := o0_1
				if o1.Op != OpPPC64OR || o1.Type != t {
					continue
				}
				_ = o1.Args[1]
				o1_0 := o1.Args[0]
				o1_1 := o1.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, o1_0, o1_1 = _i2+1, o1_1, o1_0 {
					s2 := o1_0
					if s2.Op != OpPPC64SLDconst || s2.AuxInt != 40 {
						continue
					}
					x2 := s2.Args[0]
					if x2.Op != OpPPC64MOVBZload {
						continue
					}
					i2 := x2.AuxInt
					if x2.Aux != s {
						continue
					}
					_ = x2.Args[1]
					if p != x2.Args[0] || mem != x2.Args[1] {
						continue
					}
					o2 := o1_1
					if o2.Op != OpPPC64OR || o2.Type != t {
						continue
					}
					_ = o2.Args[1]
					o2_0 := o2.Args[0]
					o2_1 := o2.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, o2_0, o2_1 = _i3+1, o2_1, o2_0 {
						s3 := o2_0
						if s3.Op != OpPPC64SLDconst || s3.AuxInt != 32 {
							continue
						}
						x3 := s3.Args[0]
						if x3.Op != OpPPC64MOVBZload {
							continue
						}
						i3 := x3.AuxInt
						if x3.Aux != s {
							continue
						}
						_ = x3.Args[1]
						if p != x3.Args[0] || mem != x3.Args[1] {
							continue
						}
						x4 := o2_1
						if x4.Op != OpPPC64MOVWBRload || x4.Type != t {
							continue
						}
						_ = x4.Args[1]
						x4_0 := x4.Args[0]
						if x4_0.Op != OpPPC64MOVDaddr || x4_0.Type != typ.Uintptr {
							continue
						}
						i4 := x4_0.AuxInt
						if p != x4_0.Args[0] || mem != x4.Args[1] || !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
							continue
						}
						b = mergePoint(b, x0, x1, x2, x3, x4)
						v0 := b.NewValue0(x4.Pos, OpPPC64MOVDBRload, t)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(x4.Pos, OpPPC64MOVDaddr, typ.Uintptr)
						v1.AuxInt = i0
						v1.Aux = s
						v1.AddArg(p)
						v0.AddArg(v1)
						v0.AddArg(mem)
						return true
					}
				}
			}
		}
		break
	}
	// match: (OR <t> x7:(MOVBZload [i7] {s} p mem) o5:(OR <t> s6:(SLDconst x6:(MOVBZload [i6] {s} p mem) [8]) o4:(OR <t> s5:(SLDconst x5:(MOVBZload [i5] {s} p mem) [16]) o3:(OR <t> s4:(SLDconst x4:(MOVBZload [i4] {s} p mem) [24]) s0:(SLWconst x3:(MOVWBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem) [32])))))
	// cond: !config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)
	// result: @mergePoint(b,x3,x4,x5,x6,x7) (MOVDBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x7 := v_0
			if x7.Op != OpPPC64MOVBZload {
				continue
			}
			i7 := x7.AuxInt
			s := x7.Aux
			mem := x7.Args[1]
			p := x7.Args[0]
			o5 := v_1
			if o5.Op != OpPPC64OR || o5.Type != t {
				continue
			}
			_ = o5.Args[1]
			o5_0 := o5.Args[0]
			o5_1 := o5.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o5_0, o5_1 = _i1+1, o5_1, o5_0 {
				s6 := o5_0
				if s6.Op != OpPPC64SLDconst || s6.AuxInt != 8 {
					continue
				}
				x6 := s6.Args[0]
				if x6.Op != OpPPC64MOVBZload {
					continue
				}
				i6 := x6.AuxInt
				if x6.Aux != s {
					continue
				}
				_ = x6.Args[1]
				if p != x6.Args[0] || mem != x6.Args[1] {
					continue
				}
				o4 := o5_1
				if o4.Op != OpPPC64OR || o4.Type != t {
					continue
				}
				_ = o4.Args[1]
				o4_0 := o4.Args[0]
				o4_1 := o4.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, o4_0, o4_1 = _i2+1, o4_1, o4_0 {
					s5 := o4_0
					if s5.Op != OpPPC64SLDconst || s5.AuxInt != 16 {
						continue
					}
					x5 := s5.Args[0]
					if x5.Op != OpPPC64MOVBZload {
						continue
					}
					i5 := x5.AuxInt
					if x5.Aux != s {
						continue
					}
					_ = x5.Args[1]
					if p != x5.Args[0] || mem != x5.Args[1] {
						continue
					}
					o3 := o4_1
					if o3.Op != OpPPC64OR || o3.Type != t {
						continue
					}
					_ = o3.Args[1]
					o3_0 := o3.Args[0]
					o3_1 := o3.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, o3_0, o3_1 = _i3+1, o3_1, o3_0 {
						s4 := o3_0
						if s4.Op != OpPPC64SLDconst || s4.AuxInt != 24 {
							continue
						}
						x4 := s4.Args[0]
						if x4.Op != OpPPC64MOVBZload {
							continue
						}
						i4 := x4.AuxInt
						if x4.Aux != s {
							continue
						}
						_ = x4.Args[1]
						if p != x4.Args[0] || mem != x4.Args[1] {
							continue
						}
						s0 := o3_1
						if s0.Op != OpPPC64SLWconst || s0.AuxInt != 32 {
							continue
						}
						x3 := s0.Args[0]
						if x3.Op != OpPPC64MOVWBRload || x3.Type != t {
							continue
						}
						_ = x3.Args[1]
						x3_0 := x3.Args[0]
						if x3_0.Op != OpPPC64MOVDaddr || x3_0.Type != typ.Uintptr {
							continue
						}
						i0 := x3_0.AuxInt
						if x3_0.Aux != s || p != x3_0.Args[0] || mem != x3.Args[1] || !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
							continue
						}
						b = mergePoint(b, x3, x4, x5, x6, x7)
						v0 := b.NewValue0(x3.Pos, OpPPC64MOVDBRload, t)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(x3.Pos, OpPPC64MOVDaddr, typ.Uintptr)
						v1.AuxInt = i0
						v1.Aux = s
						v1.AddArg(p)
						v0.AddArg(v1)
						v0.AddArg(mem)
						return true
					}
				}
			}
		}
		break
	}
	// match: (OR <t> x7:(MOVBZload [i7] {s} p mem) o5:(OR <t> s6:(SLDconst x6:(MOVBZload [i6] {s} p mem) [8]) o4:(OR <t> s5:(SLDconst x5:(MOVBZload [i5] {s} p mem) [16]) o3:(OR <t> s4:(SLDconst x4:(MOVBZload [i4] {s} p mem) [24]) s0:(SLDconst x3:(MOVWBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem) [32])))))
	// cond: !config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)
	// result: @mergePoint(b,x3,x4,x5,x6,x7) (MOVDBRload <t> (MOVDaddr <typ.Uintptr> [i0] {s} p) mem)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x7 := v_0
			if x7.Op != OpPPC64MOVBZload {
				continue
			}
			i7 := x7.AuxInt
			s := x7.Aux
			mem := x7.Args[1]
			p := x7.Args[0]
			o5 := v_1
			if o5.Op != OpPPC64OR || o5.Type != t {
				continue
			}
			_ = o5.Args[1]
			o5_0 := o5.Args[0]
			o5_1 := o5.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, o5_0, o5_1 = _i1+1, o5_1, o5_0 {
				s6 := o5_0
				if s6.Op != OpPPC64SLDconst || s6.AuxInt != 8 {
					continue
				}
				x6 := s6.Args[0]
				if x6.Op != OpPPC64MOVBZload {
					continue
				}
				i6 := x6.AuxInt
				if x6.Aux != s {
					continue
				}
				_ = x6.Args[1]
				if p != x6.Args[0] || mem != x6.Args[1] {
					continue
				}
				o4 := o5_1
				if o4.Op != OpPPC64OR || o4.Type != t {
					continue
				}
				_ = o4.Args[1]
				o4_0 := o4.Args[0]
				o4_1 := o4.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, o4_0, o4_1 = _i2+1, o4_1, o4_0 {
					s5 := o4_0
					if s5.Op != OpPPC64SLDconst || s5.AuxInt != 16 {
						continue
					}
					x5 := s5.Args[0]
					if x5.Op != OpPPC64MOVBZload {
						continue
					}
					i5 := x5.AuxInt
					if x5.Aux != s {
						continue
					}
					_ = x5.Args[1]
					if p != x5.Args[0] || mem != x5.Args[1] {
						continue
					}
					o3 := o4_1
					if o3.Op != OpPPC64OR || o3.Type != t {
						continue
					}
					_ = o3.Args[1]
					o3_0 := o3.Args[0]
					o3_1 := o3.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, o3_0, o3_1 = _i3+1, o3_1, o3_0 {
						s4 := o3_0
						if s4.Op != OpPPC64SLDconst || s4.AuxInt != 24 {
							continue
						}
						x4 := s4.Args[0]
						if x4.Op != OpPPC64MOVBZload {
							continue
						}
						i4 := x4.AuxInt
						if x4.Aux != s {
							continue
						}
						_ = x4.Args[1]
						if p != x4.Args[0] || mem != x4.Args[1] {
							continue
						}
						s0 := o3_1
						if s0.Op != OpPPC64SLDconst || s0.AuxInt != 32 {
							continue
						}
						x3 := s0.Args[0]
						if x3.Op != OpPPC64MOVWBRload || x3.Type != t {
							continue
						}
						_ = x3.Args[1]
						x3_0 := x3.Args[0]
						if x3_0.Op != OpPPC64MOVDaddr || x3_0.Type != typ.Uintptr {
							continue
						}
						i0 := x3_0.AuxInt
						if x3_0.Aux != s || p != x3_0.Args[0] || mem != x3.Args[1] || !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
							continue
						}
						b = mergePoint(b, x3, x4, x5, x6, x7)
						v0 := b.NewValue0(x3.Pos, OpPPC64MOVDBRload, t)
						v.reset(OpCopy)
						v.AddArg(v0)
						v1 := b.NewValue0(x3.Pos, OpPPC64MOVDaddr, typ.Uintptr)
						v1.AuxInt = i0
						v1.Aux = s
						v1.AddArg(p)
						v0.AddArg(v1)
						v0.AddArg(mem)
						return true
					}
				}
			}
		}
		break
	}
	return false
}
func rewriteValuePPC64_OpPPC64ORN(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ORN x (MOVDconst [-1]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst || v_1.AuxInt != -1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ORconst [c] (ORconst [d] x))
	// result: (ORconst [c|d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64ORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpPPC64ORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	// match: (ORconst [-1] _)
	// result: (MOVDconst [-1])
	for {
		if v.AuxInt != -1 {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ROTL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ROTL x (MOVDconst [c]))
	// result: (ROTLconst x [c&63])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64ROTLconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ROTLW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ROTLW x (MOVDconst [c]))
	// result: (ROTLWconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64ROTLWconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64SUB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUB x (MOVDconst [c]))
	// cond: is32Bit(-c)
	// result: (ADDconst [-c] x)
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(-c)) {
			break
		}
		v.reset(OpPPC64ADDconst)
		v.AuxInt = -c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64XOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (XOR (SLDconst x [c]) (SRDconst x [d]))
	// cond: d == 64-c
	// result: (ROTLconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLDconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpPPC64SRDconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 64-c) {
				continue
			}
			v.reset(OpPPC64ROTLconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XOR (SLWconst x [c]) (SRWconst x [d]))
	// cond: d == 32-c
	// result: (ROTLWconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLWconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpPPC64SRWconst {
				continue
			}
			d := v_1.AuxInt
			if x != v_1.Args[0] || !(d == 32-c) {
				continue
			}
			v.reset(OpPPC64ROTLWconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XOR (SLD x (ANDconst <typ.Int64> [63] y)) (SRD x (SUB <typ.UInt> (MOVDconst [64]) (ANDconst <typ.UInt> [63] y))))
	// result: (ROTL x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLD {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpPPC64ANDconst || v_0_1.Type != typ.Int64 || v_0_1.AuxInt != 63 {
				continue
			}
			y := v_0_1.Args[0]
			if v_1.Op != OpPPC64SRD {
				continue
			}
			_ = v_1.Args[1]
			if x != v_1.Args[0] {
				continue
			}
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpPPC64SUB || v_1_1.Type != typ.UInt {
				continue
			}
			_ = v_1_1.Args[1]
			v_1_1_0 := v_1_1.Args[0]
			if v_1_1_0.Op != OpPPC64MOVDconst || v_1_1_0.AuxInt != 64 {
				continue
			}
			v_1_1_1 := v_1_1.Args[1]
			if v_1_1_1.Op != OpPPC64ANDconst || v_1_1_1.Type != typ.UInt || v_1_1_1.AuxInt != 63 || y != v_1_1_1.Args[0] {
				continue
			}
			v.reset(OpPPC64ROTL)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (XOR (SLW x (ANDconst <typ.Int32> [31] y)) (SRW x (SUB <typ.UInt> (MOVDconst [32]) (ANDconst <typ.UInt> [31] y))))
	// result: (ROTLW x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64SLW {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpPPC64ANDconst || v_0_1.Type != typ.Int32 || v_0_1.AuxInt != 31 {
				continue
			}
			y := v_0_1.Args[0]
			if v_1.Op != OpPPC64SRW {
				continue
			}
			_ = v_1.Args[1]
			if x != v_1.Args[0] {
				continue
			}
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpPPC64SUB || v_1_1.Type != typ.UInt {
				continue
			}
			_ = v_1_1.Args[1]
			v_1_1_0 := v_1_1.Args[0]
			if v_1_1_0.Op != OpPPC64MOVDconst || v_1_1_0.AuxInt != 32 {
				continue
			}
			v_1_1_1 := v_1_1.Args[1]
			if v_1_1_1.Op != OpPPC64ANDconst || v_1_1_1.Type != typ.UInt || v_1_1_1.AuxInt != 31 || y != v_1_1_1.Args[0] {
				continue
			}
			v.reset(OpPPC64ROTLW)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (XOR (MOVDconst [c]) (MOVDconst [d]))
	// result: (MOVDconst [c^d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpPPC64MOVDconst {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpPPC64MOVDconst {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpPPC64MOVDconst)
			v.AuxInt = c ^ d
			return true
		}
		break
	}
	// match: (XOR x (MOVDconst [c]))
	// cond: isU32Bit(c)
	// result: (XORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpPPC64MOVDconst {
				continue
			}
			c := v_1.AuxInt
			if !(isU32Bit(c)) {
				continue
			}
			v.reset(OpPPC64XORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuePPC64_OpPPC64XORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (XORconst [c] (XORconst [d] x))
	// result: (XORconst [c^d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpPPC64XORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpPPC64XORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	// match: (XORconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPanicBounds(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicBoundsA [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpPPC64LoweredPanicBoundsA)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicBoundsB [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpPPC64LoweredPanicBoundsB)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicBoundsC [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpPPC64LoweredPanicBoundsC)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPopCount16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount16 x)
	// result: (POPCNTW (MOVHZreg x))
	for {
		x := v_0
		v.reset(OpPPC64POPCNTW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpPopCount32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount32 x)
	// result: (POPCNTW (MOVWZreg x))
	for {
		x := v_0
		v.reset(OpPPC64POPCNTW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpPopCount8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount8 x)
	// result: (POPCNTB (MOVBZreg x))
	for {
		x := v_0
		v.reset(OpPPC64POPCNTB)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVDconst [c]))
	// result: (Or16 (Lsh16x64 <t> x (MOVDconst [c&15])) (Rsh16Ux64 <t> x (MOVDconst [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = c & 15
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux64, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v3.AuxInt = -c & 15
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValuePPC64_OpRotateLeft32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RotateLeft32 x (MOVDconst [c]))
	// result: (ROTLWconst [c&31] x)
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64ROTLWconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	// match: (RotateLeft32 x y)
	// result: (ROTLW x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64ROTLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpRotateLeft64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RotateLeft64 x (MOVDconst [c]))
	// result: (ROTLconst [c&63] x)
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64ROTLconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	// match: (RotateLeft64 x y)
	// result: (ROTL x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64ROTL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVDconst [c]))
	// result: (Or8 (Lsh8x64 <t> x (MOVDconst [c&7])) (Rsh8Ux64 <t> x (MOVDconst [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = c & 7
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux64, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v3.AuxInt = -c & 7
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValuePPC64_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVHZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16Ux16 x y)
	// result: (SRW (ZeroExt16to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SRWconst (ZeroExt16to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux32 x (MOVDconst [c]))
	// cond: uint32(c) < 16
	// result: (SRWconst (ZeroExt16to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVHZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16Ux32 x y)
	// result: (SRW (ZeroExt16to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRWconst (ZeroExt16to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh16Ux64 x (MOVDconst [c]))
	// cond: uint64(c) < 16
	// result: (SRWconst (ZeroExt16to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVHZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16Ux64 x y)
	// result: (SRW (ZeroExt16to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVHZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16Ux8 x y)
	// result: (SRW (ZeroExt16to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVHreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16x16 x y)
	// result: (SRAW (SignExt16to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SRAWconst (SignExt16to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x32 x (MOVDconst [c]))
	// cond: uint32(c) < 16
	// result: (SRAWconst (SignExt16to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVHreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16x32 x y)
	// result: (SRAW (SignExt16to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRAWconst (SignExt16to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (SRAWconst (SignExt16to32 x) [63])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (MOVDconst [c]))
	// cond: uint64(c) < 16
	// result: (SRAWconst (SignExt16to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVHreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16x64 x y)
	// result: (SRAW (SignExt16to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVHreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh16x8 x y)
	// result: (SRAW (SignExt16to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32Ux16 x y)
	// result: (SRW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SRWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux32 x (MOVDconst [c]))
	// cond: uint32(c) < 32
	// result: (SRWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32Ux32 x y)
	// result: (SRW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh32Ux64 x (MOVDconst [c]))
	// cond: uint64(c) < 32
	// result: (SRWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32Ux64 x (AND y (MOVDconst [31])))
	// result: (SRW x (ANDconst <typ.Int32> [31] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			y := v_1_0
			if v_1_1.Op != OpPPC64MOVDconst || v_1_1.AuxInt != 31 {
				continue
			}
			v.reset(OpPPC64SRW)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
			v0.AuxInt = 31
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Rsh32Ux64 x (ANDconst <typ.UInt> [31] y))
	// result: (SRW x (ANDconst <typ.UInt> [31] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64ANDconst || v_1.Type != typ.UInt || v_1.AuxInt != 31 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32Ux64 x (SUB <typ.UInt> (MOVDconst [32]) (ANDconst <typ.UInt> [31] y)))
	// result: (SRW x (SUB <typ.UInt> (MOVDconst [32]) (ANDconst <typ.UInt> [31] y)))
	for {
		x := v_0
		if v_1.Op != OpPPC64SUB || v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst || v_1_0.AuxInt != 32 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst || v_1_1.Type != typ.UInt || v_1_1.AuxInt != 31 {
			break
		}
		y := v_1_1.Args[0]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 32
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 31
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32Ux64 x (SUB <typ.UInt> (MOVDconst [32]) (AND <typ.UInt> y (MOVDconst [31]))))
	// result: (SRW x (SUB <typ.UInt> (MOVDconst [32]) (ANDconst <typ.UInt> [31] y)))
	for {
		x := v_0
		if v_1.Op != OpPPC64SUB || v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst || v_1_0.AuxInt != 32 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND || v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		v_1_1_1 := v_1_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_1_0, v_1_1_1 = _i0+1, v_1_1_1, v_1_1_0 {
			y := v_1_1_0
			if v_1_1_1.Op != OpPPC64MOVDconst || v_1_1_1.AuxInt != 31 {
				continue
			}
			v.reset(OpPPC64SRW)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
			v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
			v1.AuxInt = 32
			v0.AddArg(v1)
			v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
			v2.AuxInt = 31
			v2.AddArg(y)
			v0.AddArg(v2)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Rsh32Ux64 x y)
	// result: (SRW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32Ux8 x y)
	// result: (SRW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32x16 x y)
	// result: (SRAW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x32 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SRAWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x32 x (MOVDconst [c]))
	// cond: uint32(c) < 32
	// result: (SRAWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32x32 x y)
	// result: (SRAW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRAWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (SRAWconst x [63])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = 63
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x (MOVDconst [c]))
	// cond: uint64(c) < 32
	// result: (SRAWconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32x64 x (AND y (MOVDconst [31])))
	// result: (SRAW x (ANDconst <typ.Int32> [31] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			y := v_1_0
			if v_1_1.Op != OpPPC64MOVDconst || v_1_1.AuxInt != 31 {
				continue
			}
			v.reset(OpPPC64SRAW)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
			v0.AuxInt = 31
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Rsh32x64 x (ANDconst <typ.UInt> [31] y))
	// result: (SRAW x (ANDconst <typ.UInt> [31] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64ANDconst || v_1.Type != typ.UInt || v_1.AuxInt != 31 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x64 x (SUB <typ.UInt> (MOVDconst [32]) (ANDconst <typ.UInt> [31] y)))
	// result: (SRAW x (SUB <typ.UInt> (MOVDconst [32]) (ANDconst <typ.UInt> [31] y)))
	for {
		x := v_0
		if v_1.Op != OpPPC64SUB || v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst || v_1_0.AuxInt != 32 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst || v_1_1.Type != typ.UInt || v_1_1.AuxInt != 31 {
			break
		}
		y := v_1_1.Args[0]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 32
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 31
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x64 x (SUB <typ.UInt> (MOVDconst [32]) (AND <typ.UInt> y (MOVDconst [31]))))
	// result: (SRAW x (SUB <typ.UInt> (MOVDconst [32]) (ANDconst <typ.UInt> [31] y)))
	for {
		x := v_0
		if v_1.Op != OpPPC64SUB || v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst || v_1_0.AuxInt != 32 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND || v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		v_1_1_1 := v_1_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_1_0, v_1_1_1 = _i0+1, v_1_1_1, v_1_1_0 {
			y := v_1_1_0
			if v_1_1_1.Op != OpPPC64MOVDconst || v_1_1_1.AuxInt != 31 {
				continue
			}
			v.reset(OpPPC64SRAW)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
			v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
			v1.AuxInt = 32
			v0.AddArg(v1)
			v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
			v2.AuxInt = 31
			v2.AddArg(y)
			v0.AddArg(v2)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Rsh32x64 x y)
	// result: (SRAW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh32x8 x y)
	// result: (SRAW x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64Ux16 x y)
	// result: (SRD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux32 x (Const64 [c]))
	// cond: uint32(c) < 64
	// result: (SRDconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux32 x (MOVDconst [c]))
	// cond: uint32(c) < 64
	// result: (SRDconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64Ux32 x y)
	// result: (SRD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux64 x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRDconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh64Ux64 x (MOVDconst [c]))
	// cond: uint64(c) < 64
	// result: (SRDconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64Ux64 x (AND y (MOVDconst [63])))
	// result: (SRD x (ANDconst <typ.Int64> [63] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			y := v_1_0
			if v_1_1.Op != OpPPC64MOVDconst || v_1_1.AuxInt != 63 {
				continue
			}
			v.reset(OpPPC64SRD)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
			v0.AuxInt = 63
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Rsh64Ux64 x (ANDconst <typ.UInt> [63] y))
	// result: (SRD x (ANDconst <typ.UInt> [63] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64ANDconst || v_1.Type != typ.UInt || v_1.AuxInt != 63 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64Ux64 x (SUB <typ.UInt> (MOVDconst [64]) (ANDconst <typ.UInt> [63] y)))
	// result: (SRD x (SUB <typ.UInt> (MOVDconst [64]) (ANDconst <typ.UInt> [63] y)))
	for {
		x := v_0
		if v_1.Op != OpPPC64SUB || v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst || v_1_0.AuxInt != 64 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst || v_1_1.Type != typ.UInt || v_1_1.AuxInt != 63 {
			break
		}
		y := v_1_1.Args[0]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 64
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 63
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64Ux64 x (SUB <typ.UInt> (MOVDconst [64]) (AND <typ.UInt> y (MOVDconst [63]))))
	// result: (SRD x (SUB <typ.UInt> (MOVDconst [64]) (ANDconst <typ.UInt> [63] y)))
	for {
		x := v_0
		if v_1.Op != OpPPC64SUB || v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst || v_1_0.AuxInt != 64 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND || v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		v_1_1_1 := v_1_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_1_0, v_1_1_1 = _i0+1, v_1_1_1, v_1_1_0 {
			y := v_1_1_0
			if v_1_1_1.Op != OpPPC64MOVDconst || v_1_1_1.AuxInt != 63 {
				continue
			}
			v.reset(OpPPC64SRD)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
			v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
			v1.AuxInt = 64
			v0.AddArg(v1)
			v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
			v2.AuxInt = 63
			v2.AddArg(y)
			v0.AddArg(v2)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Rsh64Ux64 x y)
	// result: (SRD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64Ux8 x y)
	// result: (SRD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64x16 x y)
	// result: (SRAD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x32 x (Const64 [c]))
	// cond: uint32(c) < 64
	// result: (SRADconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x32 x (MOVDconst [c]))
	// cond: uint32(c) < 64
	// result: (SRADconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64x32 x y)
	// result: (SRAD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x64 x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRADconst x [63])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = 63
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 x (MOVDconst [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst x [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64x64 x (AND y (MOVDconst [63])))
	// result: (SRAD x (ANDconst <typ.Int64> [63] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			y := v_1_0
			if v_1_1.Op != OpPPC64MOVDconst || v_1_1.AuxInt != 63 {
				continue
			}
			v.reset(OpPPC64SRAD)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
			v0.AuxInt = 63
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Rsh64x64 x (ANDconst <typ.UInt> [63] y))
	// result: (SRAD x (ANDconst <typ.UInt> [63] y))
	for {
		x := v_0
		if v_1.Op != OpPPC64ANDconst || v_1.Type != typ.UInt || v_1.AuxInt != 63 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64x64 x (SUB <typ.UInt> (MOVDconst [64]) (ANDconst <typ.UInt> [63] y)))
	// result: (SRAD x (SUB <typ.UInt> (MOVDconst [64]) (ANDconst <typ.UInt> [63] y)))
	for {
		x := v_0
		if v_1.Op != OpPPC64SUB || v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst || v_1_0.AuxInt != 64 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst || v_1_1.Type != typ.UInt || v_1_1.AuxInt != 63 {
			break
		}
		y := v_1_1.Args[0]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 64
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 63
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64x64 x (SUB <typ.UInt> (MOVDconst [64]) (AND <typ.UInt> y (MOVDconst [63]))))
	// result: (SRAD x (SUB <typ.UInt> (MOVDconst [64]) (ANDconst <typ.UInt> [63] y)))
	for {
		x := v_0
		if v_1.Op != OpPPC64SUB || v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst || v_1_0.AuxInt != 64 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND || v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		v_1_1_1 := v_1_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_1_0, v_1_1_1 = _i0+1, v_1_1_1, v_1_1_0 {
			y := v_1_1_0
			if v_1_1_1.Op != OpPPC64MOVDconst || v_1_1_1.AuxInt != 63 {
				continue
			}
			v.reset(OpPPC64SRAD)
			v.AddArg(x)
			v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
			v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
			v1.AuxInt = 64
			v0.AddArg(v1)
			v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
			v2.AuxInt = 63
			v2.AddArg(y)
			v0.AddArg(v2)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Rsh64x64 x y)
	// result: (SRAD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAD x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Rsh64x8 x y)
	// result: (SRAD x (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVBZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8Ux16 x y)
	// result: (SRW (ZeroExt8to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SRWconst (ZeroExt8to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux32 x (MOVDconst [c]))
	// cond: uint32(c) < 8
	// result: (SRWconst (ZeroExt8to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVBZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8Ux32 x y)
	// result: (SRW (ZeroExt8to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRWconst (ZeroExt8to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh8Ux64 x (MOVDconst [c]))
	// cond: uint64(c) < 8
	// result: (SRWconst (ZeroExt8to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVBZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8Ux64 x y)
	// result: (SRW (ZeroExt8to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRW (MOVBZreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8Ux8 x y)
	// result: (SRW (ZeroExt8to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVBreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8x16 x y)
	// result: (SRAW (SignExt8to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt16to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SRAWconst (SignExt8to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x32 x (MOVDconst [c]))
	// cond: uint32(c) < 8
	// result: (SRAWconst (SignExt8to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x32 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVBreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8x32 x y)
	// result: (SRAW (SignExt8to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt32to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRAWconst (SignExt8to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (SRAWconst (SignExt8to32 x) [63])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (MOVDconst [c]))
	// cond: uint64(c) < 8
	// result: (SRAWconst (SignExt8to32 x) [c])
	for {
		x := v_0
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVBreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8x64 x y)
	// result: (SRAW (SignExt8to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 x y)
	// cond: shiftIsBounded(v)
	// result: (SRAW (MOVBreg x) y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
	// match: (Rsh8x8 x y)
	// result: (SRAW (SignExt8to32 x) (ORN y <typ.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt8to64 y)))))
	for {
		x := v_0
		y := v_1
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SRADconst (NEG <t> x) [63])
	for {
		t := v.Type
		x := v_0
		v.reset(OpPPC64SRADconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpPPC64NEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (FMOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && is32BitFloat(val.Type)
	// result: (FMOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (FMOVSstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpPPC64FMOVSstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && (is64BitInt(val.Type) || isPtr(val.Type))
	// result: (MOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && (is64BitInt(val.Type) || isPtr(val.Type))) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && is32BitInt(val.Type)
	// result: (MOVWstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && is32BitInt(val.Type)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpTrunc16to8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc16to8 <t> x)
	// cond: isSigned(t)
	// result: (MOVBreg x)
	for {
		t := v.Type
		x := v_0
		if !(isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (Trunc16to8 x)
	// result: (MOVBZreg x)
	for {
		x := v_0
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc32to16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc32to16 <t> x)
	// cond: isSigned(t)
	// result: (MOVHreg x)
	for {
		t := v.Type
		x := v_0
		if !(isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (Trunc32to16 x)
	// result: (MOVHZreg x)
	for {
		x := v_0
		v.reset(OpPPC64MOVHZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc32to8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc32to8 <t> x)
	// cond: isSigned(t)
	// result: (MOVBreg x)
	for {
		t := v.Type
		x := v_0
		if !(isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (Trunc32to8 x)
	// result: (MOVBZreg x)
	for {
		x := v_0
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc64to16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc64to16 <t> x)
	// cond: isSigned(t)
	// result: (MOVHreg x)
	for {
		t := v.Type
		x := v_0
		if !(isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (Trunc64to16 x)
	// result: (MOVHZreg x)
	for {
		x := v_0
		v.reset(OpPPC64MOVHZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc64to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc64to32 <t> x)
	// cond: isSigned(t)
	// result: (MOVWreg x)
	for {
		t := v.Type
		x := v_0
		if !(isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (Trunc64to32 x)
	// result: (MOVWZreg x)
	for {
		x := v_0
		v.reset(OpPPC64MOVWZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc64to8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc64to8 <t> x)
	// cond: isSigned(t)
	// result: (MOVBreg x)
	for {
		t := v.Type
		x := v_0
		if !(isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (Trunc64to8 x)
	// result: (MOVBZreg x)
	for {
		x := v_0
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v_1
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Zero [1] destptr mem)
	// result: (MOVBstorezero destptr mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpPPC64MOVBstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] destptr mem)
	// result: (MOVHstorezero destptr mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpPPC64MOVHstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [3] destptr mem)
	// result: (MOVBstorezero [2] destptr (MOVHstorezero destptr mem))
	for {
		if v.AuxInt != 3 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = 2
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHstorezero, types.TypeMem)
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [4] destptr mem)
	// result: (MOVWstorezero destptr mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpPPC64MOVWstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [5] destptr mem)
	// result: (MOVBstorezero [4] destptr (MOVWstorezero destptr mem))
	for {
		if v.AuxInt != 5 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = 4
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, types.TypeMem)
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [6] destptr mem)
	// result: (MOVHstorezero [4] destptr (MOVWstorezero destptr mem))
	for {
		if v.AuxInt != 6 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = 4
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, types.TypeMem)
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [7] destptr mem)
	// result: (MOVBstorezero [6] destptr (MOVHstorezero [4] destptr (MOVWstorezero destptr mem)))
	for {
		if v.AuxInt != 7 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = 6
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHstorezero, types.TypeMem)
		v0.AuxInt = 4
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, types.TypeMem)
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [8] {t} destptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVDstorezero destptr mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		destptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [8] destptr mem)
	// result: (MOVWstorezero [4] destptr (MOVWstorezero [0] destptr mem))
	for {
		if v.AuxInt != 8 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = 4
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [12] {t} destptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstorezero [8] destptr (MOVDstorezero [0] destptr mem))
	for {
		if v.AuxInt != 12 {
			break
		}
		t := v.Aux
		destptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = 8
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [16] {t} destptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVDstorezero [8] destptr (MOVDstorezero [0] destptr mem))
	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		destptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 8
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [24] {t} destptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVDstorezero [16] destptr (MOVDstorezero [8] destptr (MOVDstorezero [0] destptr mem)))
	for {
		if v.AuxInt != 24 {
			break
		}
		t := v.Aux
		destptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 16
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, types.TypeMem)
		v0.AuxInt = 8
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [32] {t} destptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVDstorezero [24] destptr (MOVDstorezero [16] destptr (MOVDstorezero [8] destptr (MOVDstorezero [0] destptr mem))))
	for {
		if v.AuxInt != 32 {
			break
		}
		t := v.Aux
		destptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 24
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, types.TypeMem)
		v0.AuxInt = 16
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, types.TypeMem)
		v2.AuxInt = 0
		v2.AddArg(destptr)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [s] ptr mem)
	// result: (LoweredZero [s] ptr mem)
	for {
		s := v.AuxInt
		ptr := v_0
		mem := v_1
		v.reset(OpPPC64LoweredZero)
		v.AuxInt = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteBlockPPC64(b *Block) bool {
	switch b.Kind {
	case BlockPPC64EQ:
		// match: (EQ (CMPconst [0] (ANDconst [c] x)) yes no)
		// result: (EQ (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64EQ)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPWconst [0] (ANDconst [c] x)) yes no)
		// result: (EQ (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64EQ)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpPPC64FlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (EQ (FlagLT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpPPC64FlagLT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (EQ (FlagGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpPPC64FlagGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (EQ (InvertFlags cmp) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpPPC64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockPPC64EQ)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (CMPconst [0] (ANDconst [c] x)) yes no)
		// result: (EQ (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64EQ)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPWconst [0] (ANDconst [c] x)) yes no)
		// result: (EQ (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64EQ)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] z:(AND x y)) yes no)
		// cond: z.Uses == 1
		// result: (EQ (ANDCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64AND {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64EQ)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (CMPconst [0] z:(OR x y)) yes no)
		// cond: z.Uses == 1
		// result: (EQ (ORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64OR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64EQ)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (CMPconst [0] z:(XOR x y)) yes no)
		// cond: z.Uses == 1
		// result: (EQ (XORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64XOR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64EQ)
				v0 := b.NewValue0(v_0.Pos, OpPPC64XORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
	case BlockPPC64GE:
		// match: (GE (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpPPC64FlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GE (FlagLT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpPPC64FlagLT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GE (FlagGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpPPC64FlagGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GE (InvertFlags cmp) yes no)
		// result: (LE cmp yes no)
		for b.Controls[0].Op == OpPPC64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockPPC64LE)
			b.AddControl(cmp)
			return true
		}
		// match: (GE (CMPconst [0] (ANDconst [c] x)) yes no)
		// result: (GE (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64GE)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPWconst [0] (ANDconst [c] x)) yes no)
		// result: (GE (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64GE)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] z:(AND x y)) yes no)
		// cond: z.Uses == 1
		// result: (GE (ANDCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64AND {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64GE)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (GE (CMPconst [0] z:(OR x y)) yes no)
		// cond: z.Uses == 1
		// result: (GE (ORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64OR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64GE)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (GE (CMPconst [0] z:(XOR x y)) yes no)
		// cond: z.Uses == 1
		// result: (GE (XORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64XOR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64GE)
				v0 := b.NewValue0(v_0.Pos, OpPPC64XORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
	case BlockPPC64GT:
		// match: (GT (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpPPC64FlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagLT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpPPC64FlagLT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpPPC64FlagGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GT (InvertFlags cmp) yes no)
		// result: (LT cmp yes no)
		for b.Controls[0].Op == OpPPC64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockPPC64LT)
			b.AddControl(cmp)
			return true
		}
		// match: (GT (CMPconst [0] (ANDconst [c] x)) yes no)
		// result: (GT (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64GT)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPWconst [0] (ANDconst [c] x)) yes no)
		// result: (GT (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64GT)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] z:(AND x y)) yes no)
		// cond: z.Uses == 1
		// result: (GT (ANDCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64AND {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64GT)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (GT (CMPconst [0] z:(OR x y)) yes no)
		// cond: z.Uses == 1
		// result: (GT (ORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64OR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64GT)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (GT (CMPconst [0] z:(XOR x y)) yes no)
		// cond: z.Uses == 1
		// result: (GT (XORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64XOR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64GT)
				v0 := b.NewValue0(v_0.Pos, OpPPC64XORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
	case BlockIf:
		// match: (If (Equal cc) yes no)
		// result: (EQ cc yes no)
		for b.Controls[0].Op == OpPPC64Equal {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockPPC64EQ)
			b.AddControl(cc)
			return true
		}
		// match: (If (NotEqual cc) yes no)
		// result: (NE cc yes no)
		for b.Controls[0].Op == OpPPC64NotEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockPPC64NE)
			b.AddControl(cc)
			return true
		}
		// match: (If (LessThan cc) yes no)
		// result: (LT cc yes no)
		for b.Controls[0].Op == OpPPC64LessThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockPPC64LT)
			b.AddControl(cc)
			return true
		}
		// match: (If (LessEqual cc) yes no)
		// result: (LE cc yes no)
		for b.Controls[0].Op == OpPPC64LessEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockPPC64LE)
			b.AddControl(cc)
			return true
		}
		// match: (If (GreaterThan cc) yes no)
		// result: (GT cc yes no)
		for b.Controls[0].Op == OpPPC64GreaterThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockPPC64GT)
			b.AddControl(cc)
			return true
		}
		// match: (If (GreaterEqual cc) yes no)
		// result: (GE cc yes no)
		for b.Controls[0].Op == OpPPC64GreaterEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockPPC64GE)
			b.AddControl(cc)
			return true
		}
		// match: (If (FLessThan cc) yes no)
		// result: (FLT cc yes no)
		for b.Controls[0].Op == OpPPC64FLessThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockPPC64FLT)
			b.AddControl(cc)
			return true
		}
		// match: (If (FLessEqual cc) yes no)
		// result: (FLE cc yes no)
		for b.Controls[0].Op == OpPPC64FLessEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockPPC64FLE)
			b.AddControl(cc)
			return true
		}
		// match: (If (FGreaterThan cc) yes no)
		// result: (FGT cc yes no)
		for b.Controls[0].Op == OpPPC64FGreaterThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockPPC64FGT)
			b.AddControl(cc)
			return true
		}
		// match: (If (FGreaterEqual cc) yes no)
		// result: (FGE cc yes no)
		for b.Controls[0].Op == OpPPC64FGreaterEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockPPC64FGE)
			b.AddControl(cc)
			return true
		}
		// match: (If cond yes no)
		// result: (NE (CMPWconst [0] cond) yes no)
		for {
			cond := b.Controls[0]
			b.Reset(BlockPPC64NE)
			v0 := b.NewValue0(cond.Pos, OpPPC64CMPWconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(cond)
			b.AddControl(v0)
			return true
		}
	case BlockPPC64LE:
		// match: (LE (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpPPC64FlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LE (FlagLT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpPPC64FlagLT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LE (FlagGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpPPC64FlagGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LE (InvertFlags cmp) yes no)
		// result: (GE cmp yes no)
		for b.Controls[0].Op == OpPPC64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockPPC64GE)
			b.AddControl(cmp)
			return true
		}
		// match: (LE (CMPconst [0] (ANDconst [c] x)) yes no)
		// result: (LE (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64LE)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPWconst [0] (ANDconst [c] x)) yes no)
		// result: (LE (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64LE)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] z:(AND x y)) yes no)
		// cond: z.Uses == 1
		// result: (LE (ANDCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64AND {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64LE)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (LE (CMPconst [0] z:(OR x y)) yes no)
		// cond: z.Uses == 1
		// result: (LE (ORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64OR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64LE)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (LE (CMPconst [0] z:(XOR x y)) yes no)
		// cond: z.Uses == 1
		// result: (LE (XORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64XOR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64LE)
				v0 := b.NewValue0(v_0.Pos, OpPPC64XORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
	case BlockPPC64LT:
		// match: (LT (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpPPC64FlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LT (FlagLT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpPPC64FlagLT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LT (FlagGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpPPC64FlagGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LT (InvertFlags cmp) yes no)
		// result: (GT cmp yes no)
		for b.Controls[0].Op == OpPPC64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockPPC64GT)
			b.AddControl(cmp)
			return true
		}
		// match: (LT (CMPconst [0] (ANDconst [c] x)) yes no)
		// result: (LT (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64LT)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPWconst [0] (ANDconst [c] x)) yes no)
		// result: (LT (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64LT)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] z:(AND x y)) yes no)
		// cond: z.Uses == 1
		// result: (LT (ANDCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64AND {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64LT)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (LT (CMPconst [0] z:(OR x y)) yes no)
		// cond: z.Uses == 1
		// result: (LT (ORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64OR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64LT)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (LT (CMPconst [0] z:(XOR x y)) yes no)
		// cond: z.Uses == 1
		// result: (LT (XORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64XOR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64LT)
				v0 := b.NewValue0(v_0.Pos, OpPPC64XORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
	case BlockPPC64NE:
		// match: (NE (CMPWconst [0] (Equal cc)) yes no)
		// result: (EQ cc yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64Equal {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockPPC64EQ)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPWconst [0] (NotEqual cc)) yes no)
		// result: (NE cc yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64NotEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockPPC64NE)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPWconst [0] (LessThan cc)) yes no)
		// result: (LT cc yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64LessThan {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockPPC64LT)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPWconst [0] (LessEqual cc)) yes no)
		// result: (LE cc yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64LessEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockPPC64LE)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPWconst [0] (GreaterThan cc)) yes no)
		// result: (GT cc yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64GreaterThan {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockPPC64GT)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPWconst [0] (GreaterEqual cc)) yes no)
		// result: (GE cc yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64GreaterEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockPPC64GE)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPWconst [0] (FLessThan cc)) yes no)
		// result: (FLT cc yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64FLessThan {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockPPC64FLT)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPWconst [0] (FLessEqual cc)) yes no)
		// result: (FLE cc yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64FLessEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockPPC64FLE)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPWconst [0] (FGreaterThan cc)) yes no)
		// result: (FGT cc yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64FGreaterThan {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockPPC64FGT)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPWconst [0] (FGreaterEqual cc)) yes no)
		// result: (FGE cc yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64FGreaterEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockPPC64FGE)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPconst [0] (ANDconst [c] x)) yes no)
		// result: (NE (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64NE)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPWconst [0] (ANDconst [c] x)) yes no)
		// result: (NE (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64NE)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (NE (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpPPC64FlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (NE (FlagLT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpPPC64FlagLT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (NE (FlagGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpPPC64FlagGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (NE (InvertFlags cmp) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpPPC64InvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockPPC64NE)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (CMPconst [0] (ANDconst [c] x)) yes no)
		// result: (NE (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64NE)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPWconst [0] (ANDconst [c] x)) yes no)
		// result: (NE (ANDCCconst [c] x) yes no)
		for b.Controls[0].Op == OpPPC64CMPWconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0_0.AuxInt
			x := v_0_0.Args[0]
			b.Reset(BlockPPC64NE)
			v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCCconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] z:(AND x y)) yes no)
		// cond: z.Uses == 1
		// result: (NE (ANDCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64AND {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64NE)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ANDCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (CMPconst [0] z:(OR x y)) yes no)
		// cond: z.Uses == 1
		// result: (NE (ORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64OR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64NE)
				v0 := b.NewValue0(v_0.Pos, OpPPC64ORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (CMPconst [0] z:(XOR x y)) yes no)
		// cond: z.Uses == 1
		// result: (NE (XORCC x y) yes no)
		for b.Controls[0].Op == OpPPC64CMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			z := v_0.Args[0]
			if z.Op != OpPPC64XOR {
				break
			}
			_ = z.Args[1]
			z_0 := z.Args[0]
			z_1 := z.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, z_0, z_1 = _i0+1, z_1, z_0 {
				x := z_0
				y := z_1
				if !(z.Uses == 1) {
					continue
				}
				b.Reset(BlockPPC64NE)
				v0 := b.NewValue0(v_0.Pos, OpPPC64XORCC, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
	}
	return false
}
